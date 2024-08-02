package validation

import (
	"github.com/pkg/errors"
)

// For creates a new [PropertyRules] instance for the property
// which value is extracted through [PropertyGetter] function.
func For[T, S any](getter PropertyGetter[T, S]) PropertyRules[T, S] {
	return PropertyRules[T, S]{getter: func(s S) (v T, err error) { return getter(s), nil }}
}

// ForPointer accepts a getter function returning a pointer and wraps its call in order to
// safely extract the value under the pointer or return a zero value for a give type T.
// If required is set to true, the nil pointer value will result in an error and the
// validation will not proceed.
func ForPointer[T, S any](getter PropertyGetter[*T, S]) PropertyRules[T, S] {
	return PropertyRules[T, S]{getter: func(s S) (indirect T, err error) {
		ptr := getter(s)
		if ptr != nil {
			return *ptr, nil
		}
		zv := *new(T)
		return zv, emptyErr{}
	}, isPointer: true}
}

// Transform transforms value from one type to another.
// Value returned by [PropertyGetter] is transformed through [Transformer] function.
// If [Transformer] returns an error, the validation will not proceed and transformation error will be reported.
// [Transformer] is only called if [PropertyGetter] returns a non-zero value.
func Transform[T, N, S any](getter PropertyGetter[T, S], transform Transformer[T, N]) PropertyRules[N, S] {
	typInfo := getTypeInfo[T]()
	return PropertyRules[N, S]{
		transformGetter: func(s S) (transformed N, original any, err error) {
			v := getter(s)
			if isEmptyFunc(v) {
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
func GetSelf[S any]() PropertyGetter[S, S] {
	return func(s S) S { return s }
}

type Transformer[T, N any] func(T) (N, error)

type PropertyGetter[T, S any] func(S) T

type internalPropertyGetter[T, S any] func(S) (v T, err error)
type internalTransformPropertyGetter[T, S any] func(S) (transformed T, original any, err error)
type emptyErr struct{}

func (emptyErr) Error() string { return "" }

// PropertyRules is responsible for validating a single property.
type PropertyRules[T, S any] struct {
	name            string
	getter          internalPropertyGetter[T, S]
	transformGetter internalTransformPropertyGetter[T, S]
	steps           []interface{}
	required        bool
	omitEmpty       bool
	hideValue       bool
	isPointer       bool
	mode            CascadeMode
	examples        []string
	originalType    *typeInfo

	predicateMatcher[S]
}

// Validate validates the property value using provided rules.
// nolint: gocognit
func (r PropertyRules[T, S]) Validate(st S) PropertyErrors {
	var (
		ruleErrors []error
		allErrors  PropertyErrors
	)
	propValue, skip, err := r.getValue(st)
	if err != nil {
		if r.hideValue {
			err = err.HideValue()
		}
		return PropertyErrors{err}
	}
	if skip {
		return nil
	}
	if !r.matchPredicates(st) {
		return nil
	}
	for _, step := range r.steps {
		stepFailed := false
		switch v := step.(type) {
		// Same as Rule[S] as for GetSelf we'd get the same type on T and S.
		case Rule[T]:
			err := v.Validate(propValue)
			if err != nil {
				stepFailed = true
				switch ev := err.(type) {
				case *PropertyError:
					allErrors = append(allErrors, ev.PrependPropertyName(r.name))
				default:
					ruleErrors = append(ruleErrors, err)
				}
			}
		case validatorI[T]:
			err := v.Validate(propValue)
			if err != nil {
				stepFailed = true
				for _, e := range err.Errors {
					allErrors = append(allErrors, e.PrependPropertyName(r.name))
				}
			}
		}
		if stepFailed && r.mode == CascadeModeStop {
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
		return allErrors.Aggregate()
	}
	return nil
}

func (r PropertyRules[T, S]) WithName(name string) PropertyRules[T, S] {
	r.name = name
	return r
}

func (r PropertyRules[T, S]) WithExamples(examples ...string) PropertyRules[T, S] {
	r.examples = append(r.examples, examples...)
	return r
}

func (r PropertyRules[T, S]) Rules(rules ...Rule[T]) PropertyRules[T, S] {
	r.steps = appendSteps(r.steps, rules)
	return r
}

func (r PropertyRules[T, S]) Include(rules ...Validator[T]) PropertyRules[T, S] {
	r.steps = appendSteps(r.steps, rules)
	return r
}

func (r PropertyRules[T, S]) When(predicate Predicate[S], opts ...WhenOptions) PropertyRules[T, S] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

func (r PropertyRules[T, S]) Required() PropertyRules[T, S] {
	r.required = true
	return r
}

func (r PropertyRules[T, S]) OmitEmpty() PropertyRules[T, S] {
	r.omitEmpty = true
	return r
}

func (r PropertyRules[T, S]) HideValue() PropertyRules[T, S] {
	r.hideValue = true
	return r
}

func (r PropertyRules[T, S]) Cascade(mode CascadeMode) PropertyRules[T, S] {
	r.mode = mode
	return r
}

func (r PropertyRules[T, S]) plan(builder planBuilder) {
	builder.propertyPlan.IsOptional = (r.omitEmpty || r.isPointer) && !r.required
	builder.propertyPlan.IsHidden = r.hideValue
	for _, predicate := range r.predicates {
		builder.rulePlan.Conditions = append(builder.rulePlan.Conditions, predicate.description)
	}
	if r.originalType != nil {
		builder.propertyPlan.Type = r.originalType.Name
		builder.propertyPlan.Package = r.originalType.Package
	} else {
		typInfo := getTypeInfo[T]()
		builder.propertyPlan.Type = typInfo.Name
		builder.propertyPlan.Package = typInfo.Package
	}
	builder = builder.appendPath(r.name).setExamples(r.examples...)
	for _, step := range r.steps {
		if p, ok := step.(planner); ok {
			p.plan(builder)
		}
	}
	// If we don't have any rules defined for this property, append it nonetheless.
	// It can be useful when we have things like [WithExamples] or [Required] set.
	if len(r.steps) == 0 {
		*builder.all = append(*builder.all, builder)
	}
}

func appendSteps[T any](slice []interface{}, steps []T) []interface{} {
	for _, step := range steps {
		slice = append(slice, step)
	}
	return slice
}

// getValue extracts the property value from the provided property.
// It returns the value, a flag indicating whether the validation should be skipped, and any errors encountered.
func (r PropertyRules[T, S]) getValue(st S) (v T, skip bool, propErr *PropertyError) {
	var (
		err           error
		originalValue any
	)
	// Extract value from the property through correct getter.
	if r.transformGetter != nil {
		v, originalValue, err = r.transformGetter(st)
	} else {
		v, err = r.getter(st)
	}
	isEmptyError := errors.Is(err, emptyErr{})
	// Any error other than [emptyErr] is considered critical, we don't proceed with validation.
	if err != nil && !isEmptyError {
		var propValue interface{}
		// If the value was transformed, we need to set the property value to the original, pre-transformed one.
		if HasErrorCode(err, ErrorCodeTransform) {
			propValue = originalValue
		} else {
			propValue = v
		}
		return v, false, NewPropertyError(r.name, propValue, err)
	}
	isEmpty := isEmptyError || (!r.isPointer && isEmptyFunc(v))
	// If the value is not empty we simply return it.
	if !isEmpty {
		return v, false, nil
	}
	// If the value is empty and the property is required, we return [ErrorCodeRequired].
	if r.required {
		return v, false, NewPropertyError(r.name, nil, NewRequiredError())
	}
	// If the value is empty and we're skipping empty values or the value is a pointer, we skip the validation.
	if r.omitEmpty || r.isPointer {
		return v, true, nil
	}
	return v, false, nil
}
