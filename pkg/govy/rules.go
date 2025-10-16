package govy

import (
	"errors"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/typeinfo"
)

// For creates a new [PropertyRules] instance for the property
// which value is extracted through [PropertyGetter] function.
func For[T, P any](getter PropertyGetter[T, P]) PropertyRules[T, P] {
	return forConstructor(getter, inferName())
}

func forConstructor[T, P any](getter PropertyGetter[T, P], name string) PropertyRules[T, P] {
	return PropertyRules[T, P]{
		name:   name,
		getter: func(parent P) (v T, err error) { return getter(parent), nil },
	}
}

// ForPointer accepts a getter function returning a pointer and wraps its call in order to
// safely extract the value under the pointer or return a zero value for a give type T.
// If required is set to true, the nil pointer value will result in an error and the
// validation will not proceed.
func ForPointer[T, P any](getter PropertyGetter[*T, P]) PropertyRules[T, P] {
	return PropertyRules[T, P]{
		name: inferName(),
		getter: func(parent P) (indirect T, err error) {
			ptr := getter(parent)
			if ptr != nil {
				return *ptr, nil
			}
			zv := *new(T)
			return zv, emptyErr{}
		},
		isPointer: true,
	}
}

// Transform transforms value from one type to another.
// Value returned by [PropertyGetter] is transformed through [Transformer] function.
// If [Transformer] returns an error, the validation will not proceed and transformation error will be reported.
// [Transformer] is only called if [PropertyGetter] returns a non-zero value.
func Transform[T, N, P any](getter PropertyGetter[T, P], transform Transformer[T, N]) PropertyRules[N, P] {
	typInfo := typeinfo.Get[T]()
	return PropertyRules[N, P]{
		name: inferName(),
		transformGetter: func(parent P) (transformed N, original any, err error) {
			v := getter(parent)
			if internal.IsEmpty(v) {
				return transformed, nil, emptyErr{}
			}
			transformed, err = transform(v)
			if err != nil {
				return transformed, v, NewRuleError(err.Error(), ErrorCodeTransform)
			}
			return transformed, v, nil
		},
		originalType: &typInfo,
	}
}

// GetSelf is a convenience method for extracting 'self' property of a validated value.
func GetSelf[P any]() PropertyGetter[P, P] {
	return func(parent P) P { return parent }
}

// Transformer is a function that transforms a value of type T to a value of type N.
// If the transformation fails, the function should return an error.
type Transformer[T, N any] func(value T) (N, error)

// PropertyGetter is a function that extracts a property value of type T from a given parent value of type S.
type PropertyGetter[T, P any] func(parent P) T

type (
	internalPropertyGetter[T, P any]          func(P) (v T, err error)
	internalTransformPropertyGetter[T, P any] func(P) (transformed T, original any, err error)
	emptyErr                                  struct{}
)

func (emptyErr) Error() string { return "" }

// PropertyRules is responsible for validating a single property.
// It is a collection of rules, predicates, and other properties that define how the property should be validated.
// It is the middle-level building block of the validation process,
// aggregated by [Validator] and aggregating [Rule].
type PropertyRules[T, P any] struct {
	name            string
	getter          internalPropertyGetter[T, P]
	transformGetter internalTransformPropertyGetter[T, P]
	rules           []validationInterface[T]
	required        bool
	omitEmpty       bool
	hideValue       bool
	isPointer       bool
	mode            CascadeMode
	examples        []string
	originalType    *typeinfo.TypeInfo

	predicateMatcher[P]
}

// Validate validates the property value using provided rules.
func (r PropertyRules[T, P]) Validate(parent P) error {
	if !r.matchPredicates(parent) {
		return nil
	}
	var (
		ruleErrors []error
		allErrors  PropertyErrors
	)
	propValue, skip, propErr := r.getValue(parent)
	if propErr != nil {
		if r.hideValue {
			propErr = propErr.HideValue()
		}
		return PropertyErrors{propErr}
	}
	if skip {
		return nil
	}
	for i := range r.rules {
		err := r.rules[i].Validate(propValue)
		if err == nil {
			continue
		}
		switch errValue := err.(type) {
		// Same as Rule[P] as for GetSelf we'd get the same type on T and S.
		case *PropertyError:
			allErrors = append(allErrors, errValue.prependParentPropertyName(r.name))
		case *ValidatorError:
			for _, e := range errValue.Errors {
				allErrors = append(allErrors, e.prependParentPropertyName(r.name))
			}
		default:
			ruleErrors = append(ruleErrors, err)
		}
		if r.mode == CascadeModeStop {
			break
		}
	}
	if len(ruleErrors) > 0 {
		allErrors = append(allErrors, NewPropertyError(r.name, propValue, ruleErrors...))
	}
	if len(allErrors) > 0 {
		if r.hideValue {
			allErrors = allErrors.HideValue()
		}
		return allErrors.aggregate()
	}
	return nil
}

// WithName sets the name of the property.
// If the name was inferred, it will be overridden.
func (r PropertyRules[T, P]) WithName(name string) PropertyRules[T, P] {
	r.name = name
	return r
}

// WithExamples sets the examples for the property.
func (r PropertyRules[T, P]) WithExamples(examples ...string) PropertyRules[T, P] {
	r.examples = append(r.examples, examples...)
	return r
}

// Rules associates provided [Rule] with the property.
func (r PropertyRules[T, P]) Rules(rules ...rulesInterface[T]) PropertyRules[T, P] {
	for _, rule := range rules {
		r.rules = append(r.rules, rule)
	}
	return r
}

// Include embeds specified [Validator] and its [PropertyRules] into the property.
func (r PropertyRules[T, P]) Include(rules ...validatorInterface[T]) PropertyRules[T, P] {
	for _, rule := range rules {
		r.rules = append(r.rules, rule)
	}
	return r
}

// When defines a [Predicate] which determines when the rules for this property should be evaluated.
// It can be called multiple times to set multiple predicates.
// Additionally, it accepts [WhenOptions] which customizes the behavior of the predicate.
func (r PropertyRules[T, P]) When(predicate Predicate[P], opts ...WhenOptions) PropertyRules[T, P] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

// Required sets the property as required.
// If the property is its type's zero value a [rules.ErrorCodeRequired] will be returned.
func (r PropertyRules[T, P]) Required() PropertyRules[T, P] {
	r.required = true
	return r
}

// OmitEmpty sets the property rules to be omitted if its value is its type's zero value.
func (r PropertyRules[T, P]) OmitEmpty() PropertyRules[T, P] {
	r.omitEmpty = true
	return r
}

// HideValue hides the property value in the error message.
// It's useful when the value is sensitive and should not be exposed.
func (r PropertyRules[T, P]) HideValue() PropertyRules[T, P] {
	r.hideValue = true
	return r
}

// Cascade sets the [CascadeMode] for the property,
// which controls the flow of evaluating the validation rules.
func (r PropertyRules[T, P]) Cascade(mode CascadeMode) PropertyRules[T, P] {
	r.mode = mode
	return r
}

// cascadeInternal is an internal wrapper around [PropertyRules.Cascade] which
// fulfills [propertyRulesInterface] interface.
// If the [CascadeMode] is already set, it won't change it.
func (r PropertyRules[T, P]) cascadeInternal(mode CascadeMode) propertyRulesInterface[P] {
	if r.mode != 0 {
		return r
	}
	return r.Cascade(mode)
}

// plan constructs a validation plan for the property.
func (r PropertyRules[T, P]) plan(builder planBuilder) {
	builder.propertyPlan.IsHidden = r.hideValue
	builder = appendPredicatesToPlanBuilder(builder, r.predicates)
	if r.originalType != nil {
		builder.propertyPlan.TypeInfo = TypeInfo(*r.originalType)
	} else {
		builder.propertyPlan.TypeInfo = TypeInfo(typeinfo.Get[T]())
	}
	builder = builder.appendPath(r.name).setExamples(r.examples...)
	if r.required {
		// Dummy rule to register the property as required.
		NewRule(func(v T) error { return nil }).
			WithDescription(internal.RequiredDescription).
			WithErrorCode(internal.RequiredErrorCode).
			plan(builder)
	} else if r.omitEmpty || r.isPointer {
		// Dummy rule to register the property as optional.
		NewRule(func(v T) error { return nil }).
			WithDescription(internal.OptionalDescription).
			WithErrorCode(internal.OptionalErrorCode).
			plan(builder)
	}
	for _, rule := range r.rules {
		if p, ok := rule.(planner); ok {
			p.plan(builder)
		}
	}
	// If we don't have any rules defined for this property, append it nonetheless.
	// It can be useful when we have things like [WithExamples] or [Required] set.
	if len(r.rules) == 0 {
		*builder.children = append(*builder.children, builder)
	}
}

// getValue extracts the property value from the provided property.
// It returns the value, a flag indicating whether the validation should be skipped, and any errors encountered.
func (r PropertyRules[T, P]) getValue(parent P) (v T, skip bool, propErr *PropertyError) {
	var (
		err           error
		originalValue any
	)
	// Extract value from the property through correct getter.
	if r.transformGetter != nil {
		v, originalValue, err = r.transformGetter(parent)
	} else {
		v, err = r.getter(parent)
	}
	isEmptyError := errors.Is(err, emptyErr{})
	// Any error other than [emptyErr] is considered critical, we don't proceed with validation.
	if err != nil && !isEmptyError {
		var propValue any
		// If the value was transformed, we need to set the property value to the original, pre-transformed one.
		if HasErrorCode(err, ErrorCodeTransform) {
			propValue = originalValue
		} else {
			propValue = v
		}
		return v, false, NewPropertyError(r.name, propValue, err)
	}
	isEmpty := isEmptyError || (!r.isPointer && internal.IsEmpty(v))
	// If the value is not empty we simply return it.
	if !isEmpty {
		return v, false, nil
	}
	// If the value is empty and the property is required, we return [ErrorCodeRequired].
	if r.required {
		return v, false, NewPropertyError(r.name, nil, newRequiredError())
	}
	// If the value is empty and we're skipping empty values or the value is a pointer, we skip the validation.
	if r.omitEmpty || r.isPointer {
		return v, true, nil
	}
	return v, false, nil
}

func newRequiredError() *RuleError {
	return NewRuleError(
		internal.RequiredMessage,
		internal.RequiredErrorCode,
	)
}

// isPropertyRules implements [propertyRulesInterface].
func (r PropertyRules[T, P]) isPropertyRules() {}
