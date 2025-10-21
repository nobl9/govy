package govy

import (
	"fmt"
	"strings"
)

// New creates a new [Validator] aggregating the provided property rules.
func New[T any](props ...propertyRulesInterface[T]) Validator[T] {
	return Validator[T]{props: props}
}

// Validator is the top level validation entity.
// It serves as an aggregator for [PropertyRules].
// Typically, it represents a struct.
type Validator[T any] struct {
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
