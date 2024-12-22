package govy

import (
	"fmt"
	"strings"
)

// validationInterface is a common interface implemented by all validation entities.
// These include [Validator], [PropertyRules] and [Rule].
type validationInterface[T any] interface {
	Validate(s T) error
}

// propertyRulesInterface is an internal interface which further limits
// what [New] constructor and [Validator] can accept as property rules.
//
// On top of [validationInterface] requirements it specifies internal functions
// which allow interacting with [propertyRulesInterface] instances like [PropertyRules]
// in an immutable fashion (no pointer receivers).
type propertyRulesInterface[T any] interface {
	validationInterface[T]
	cascadeInternal(mode CascadeMode) propertyRulesInterface[T]
}

// New creates a new [Validator] aggregating the provided property rules.
func New[S any](props ...propertyRulesInterface[S]) Validator[S] {
	return Validator[S]{props: props}
}

// Validator is the top level validation entity.
// It serves as an aggregator for [PropertyRules].
// Typically, it represents a struct.
type Validator[S any] struct {
	props    []propertyRulesInterface[S]
	name     string
	nameFunc func(S) string
	mode     CascadeMode

	predicateMatcher[S]
}

// WithName when a rule fails will pass the provided name to [ValidatorError.WithName].
func (v Validator[S]) WithName(name string) Validator[S] {
	v.nameFunc = nil
	v.name = name
	return v
}

// WithNameFunc when a rule fails extracts name from provided function and passes it to [ValidatorError.WithName].
// The function receives validated entity's instance as an argument.
func (v Validator[S]) WithNameFunc(f func(s S) string) Validator[S] {
	v.name = ""
	v.nameFunc = f
	return v
}

// When accepts predicates which will be evaluated BEFORE [Validator] validates ANY rules.
func (v Validator[S]) When(predicate Predicate[S], opts ...WhenOptions) Validator[S] {
	v.predicateMatcher = v.when(predicate, opts...)
	return v
}

// InferName will set the name of the [Validator] to its type S.
// If the name was already set through [Validator.WithName], it will not be overridden.
// It does not use the same inference mechanisms as [PropertyRules.InferName],
// it simply checks the [Validator] type parameter using reflection.
func (v Validator[S]) InferName() Validator[S] {
	if v.name != "" {
		return v
	}
	split := strings.Split(fmt.Sprintf("%T", *new(S)), ".")
	if len(split) == 0 {
		return v
	}
	v.name = split[len(split)-1]
	return v
}

// Cascade sets the [CascadeMode] for the validator,
// which controls the flow of evaluating the validation rules.
func (v Validator[S]) Cascade(mode CascadeMode) Validator[S] {
	v.mode = mode
	props := make([]propertyRulesInterface[S], 0, len(v.props))
	for _, prop := range v.props {
		props = append(props, prop.cascadeInternal(mode))
	}
	v.props = props
	return v
}

// Validate will first evaluate predicates before validating any rules.
// If any predicate does not pass the validation won't be executed (returns nil).
// All errors returned by property rules will be aggregated and wrapped in [ValidatorError].
func (v Validator[S]) Validate(st S) error {
	if !v.matchPredicates(st) {
		return nil
	}
	var allErrors PropertyErrors
	for _, rules := range v.props {
		err := rules.Validate(st)
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
			name = v.nameFunc(st)
		}
		return NewValidatorError(allErrors).WithName(name)
	}
	return nil
}

// ValidatorSlice is used to validate a slice of values of the type S.
// Under the hood [Validator.Validate] is called for each element and the errors
// are aggregated into [ValidatorErrors].
//
// Note: It is designed to be used for validating independent values.
// If you need to validate the slice itself, for instance, to check if it has at most N elements,
// you should use the [Validator] directly in tandem with [ForSlice] and [GetSelf].
func (v Validator[S]) ValidateSlice(s []S) error {
	errs := make(ValidatorErrors, 0)
	for i, st := range s {
		if err := v.Validate(st); err != nil {
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
func (v Validator[S]) plan(builder planBuilder) {
	for _, predicate := range v.predicates {
		builder.rulePlan.Conditions = append(builder.rulePlan.Conditions, predicate.description)
	}
	for _, rules := range v.props {
		if p, ok := rules.(planner); ok {
			p.plan(builder)
		}
	}
}
