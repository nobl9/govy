package govy

import "fmt"

type propertyRulesI[S any] interface {
	Validate(s S) PropertyErrors
}

// New creates a new [Validator] aggregating the provided property rules.
func New[S any](props ...propertyRulesI[S]) Validator[S] {
	return Validator[S]{props: props}
}

// Validator is the top level validation entity.
// It serves as an aggregator for [PropertyRules].
// Typically, it represents a struct.
type Validator[S any] struct {
	props []propertyRulesI[S]
	name  string

	predicateMatcher[S]
}

// WithName when a rule fails will pass the provided name to [ValidatorError.WithName].
func (v Validator[S]) WithName(name string) Validator[S] {
	v.name = name
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
	v.name = fmt.Sprintf("%T", *new(S))
	return v
}

// Validate will first evaluate predicates before validating any rules.
// If any predicate does not pass the validation won't be executed (returns nil).
// All errors returned by property rules will be aggregated and wrapped in [ValidatorError].
func (v Validator[S]) Validate(st S) *ValidatorError {
	if !v.matchPredicates(st) {
		return nil
	}
	var allErrors PropertyErrors
	for _, rules := range v.props {
		if errs := rules.Validate(st); len(errs) > 0 {
			allErrors = append(allErrors, errs...)
		}
	}
	if len(allErrors) != 0 {
		return NewValidatorError(allErrors).WithName(v.name)
	}
	return nil
}

// plan constructs a validation plan for all the properties of the [Validator].
func (v Validator[S]) plan(path planBuilder) {
	for _, predicate := range v.predicates {
		path.rulePlan.Conditions = append(path.rulePlan.Conditions, predicate.description)
	}
	for _, rules := range v.props {
		if p, ok := rules.(planner); ok {
			p.plan(path)
		}
	}
}
