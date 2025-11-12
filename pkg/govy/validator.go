package govy

import (
	"fmt"
	"slices"
	"strings"
)

// New creates a new [Validator] aggregating the provided property rules.
func New[T any](props ...propertyRulesInterface[T]) Validator[T] {
	return Validator[T]{
		id:    newInstanceID(),
		props: props,
	}
}

// Validator is the top level validation entity.
// It serves as an aggregator for [PropertyRules].
// Typically, it represents a struct.
type Validator[T any] struct {
	id       instanceID
	props    []propertyRulesInterface[T]
	name     string
	nameFunc func(value T) string
	mode     CascadeMode

	predicateMatcher[T]
}

// WithName when a rule fails will pass the provided name to [ValidatorError.WithName].
func (v Validator[T]) WithName(name string) Validator[T] {
	v.nameFunc = nil
	v.name = name
	return v
}

// WithID sets a unique identifier for this [Validator] instance.
// The identifier can be used to:
//   - Retrieve the validator's ID via [Validator.GetID]
//   - Reference the validator when using [Validator.RemoveProperties]
//
// This is useful when you want explicit control over identifiers
// rather than relying on validator names or auto-generated UUIDs.
func (v Validator[T]) WithID(id string) Validator[T] {
	v.id = v.id.WithUserSuppliedID(id)
	return v
}

// WithNameFunc when a rule fails extracts name from provided function and passes it to [ValidatorError.WithName].
// The function receives validated entity's instance as an argument.
func (v Validator[T]) WithNameFunc(f func(value T) string) Validator[T] {
	v.name = ""
	v.nameFunc = f
	return v
}

// When accepts predicates which will be evaluated BEFORE [Validator] validates ANY rules.
func (v Validator[T]) When(predicate Predicate[T], opts ...WhenOption) Validator[T] {
	v.predicateMatcher = v.when(predicate, opts...)
	return v
}

// InferName will set the name of the [Validator] to its type S.
// If the name was already set through [Validator.WithName], it will not be overridden.
// It does not use the same inference mechanisms as [PropertyRules.InferName],
// it simply checks the [Validator] type parameter using reflection.
func (v Validator[T]) InferName() Validator[T] {
	if v.name != "" {
		return v
	}
	split := strings.Split(fmt.Sprintf("%T", *new(T)), ".")
	if len(split) == 0 {
		return v
	}
	v.name = split[len(split)-1]
	return v
}

// Cascade sets the [CascadeMode] for the validator,
// which controls the flow of evaluating the validation rules.
func (v Validator[T]) Cascade(mode CascadeMode) Validator[T] {
	v.mode = mode
	props := make([]propertyRulesInterface[T], 0, len(v.props))
	for _, prop := range v.props {
		props = append(props, prop.cascadeInternal(mode))
	}
	v.props = props
	return v
}

// RemoveProperties removes any [PropertyRules] or included [Validator]
// which match the provided identifiers.
// It returns a modified [Validator] instance without these rules,
// the original [Validator] is not changed.
//
// Identifiers can be obtained using [PropertyRules.GetID].
// The identifier resolution follows this priority:
//   - User-supplied ID if set
//   - Property name (via [PropertyRules.WithName]) if set
//   - Auto-generated UUID otherwise
func (v Validator[T]) RemoveProperties(ids ...string) Validator[T] {
	if len(ids) == 0 {
		return v
	}
	filtered := make([]propertyRulesInterface[T], 0, len(v.props))
	for _, prop := range v.props {
		if slices.Contains(ids, prop.GetID()) {
			continue
		}
		filtered = append(filtered, prop)
	}
	v.props = filtered
	return v
}

// GetID returns an identifier for this [Validator] instance.
// The identifier is resolved in the following priority order:
//   - User-supplied ID (via [Validator.WithID]) if set
//   - Validator name (via [Validator.WithName]) if set
//   - Auto-generated UUID otherwise
func (v Validator[T]) GetID() string {
	switch {
	case v.id.HasUserSuppliedID():
		return v.id.GetUserSuppliedID()
	case v.name != "":
		return v.name
	default:
		return v.id.GetGeneratedID()
	}
}

// Validate will first evaluate predicates before validating any rules.
// If any predicate does not pass the validation won't be executed (returns nil).
// All errors returned by property rules will be aggregated and wrapped in [ValidatorError].
func (v Validator[T]) Validate(value T) error {
	if !v.matchPredicates(value) {
		return nil
	}
	var allErrors PropertyErrors
	for _, rules := range v.props {
		err := rules.Validate(value)
		if err == nil {
			continue
		}
		pErrs, ok := err.(PropertyErrors)
		if !ok {
			logWrongErrorType(PropertyErrors{}, err)
			continue
		}
		allErrors = append(allErrors, pErrs...)
		if v.mode == CascadeModeStop {
			break
		}
	}
	if len(allErrors) != 0 {
		name := v.name
		if v.nameFunc != nil {
			name = v.nameFunc(value)
		}
		return NewValidatorError(allErrors).WithName(name)
	}
	return nil
}

// ValidateSlice validates a slice of values of the type T.
// Under the hood, [Validator.Validate] is called for each element and the errors
// are aggregated into [ValidatorErrors].
//
// Note: It is designed to be used for validating independent values.
// If you need to validate the slice itself, for instance, to check if it has at most N elements,
// you should use the [Validator] directly in tandem with [ForSlice] and [GetSelf].
func (v Validator[T]) ValidateSlice(values []T) error {
	errs := make(ValidatorErrors, 0)
	for i, value := range values {
		if err := v.Validate(value); err != nil {
			vErr, ok := err.(*ValidatorError)
			if !ok {
				return err
			}
			vErr.SliceIndex = &i
			errs = append(errs, vErr)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

// plan constructs a validation plan for all the properties of the [Validator].
func (v Validator[T]) plan(builder planBuilder) {
	builder = appendPredicatesToPlanBuilder(builder, v.predicates)
	for _, rules := range v.props {
		if p, ok := rules.(planner); ok {
			p.plan(builder)
		}
	}
}

// isValidator implements [validatorInterface].
func (v Validator[T]) isValidator() {}
