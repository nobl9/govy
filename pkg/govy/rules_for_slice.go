package govy

import (
	"github.com/nobl9/govy/internal/jsonpath"
)

// ForSlice creates a new [PropertyRulesForSlice] instance for a slice property
// which value is extracted through [PropertyGetter] function.
func ForSlice[S ~[]T, T, P any](getter PropertyGetter[S, P]) PropertyRulesForSlice[S, T, P] {
	name := inferName()
	return PropertyRulesForSlice[S, T, P]{
		sliceRules:   forConstructor(GetSelf[S](), name),
		forEachRules: forConstructor(GetSelf[T](), ""),
		getter:       getter,
	}
}

// PropertyRulesForSlice is responsible for validating a single property.
type PropertyRulesForSlice[S ~[]T, T, P any] struct {
	sliceRules   PropertyRules[S, S]
	forEachRules PropertyRules[T, T]
	getter       PropertyGetter[S, P]
	mode         CascadeMode

	predicateMatcher[P]
}

// Validate executes each of the rules sequentially and aggregates the encountered errors.
func (r PropertyRulesForSlice[S, T, P]) Validate(parent P) error {
	if !r.matchPredicates(parent) {
		return nil
	}
	v := r.getter(parent)
	err := r.sliceRules.Validate(v)
	var propErrs PropertyErrors
	if err != nil {
		if r.mode == CascadeModeStop {
			return err
		}
		var ok bool
		propErrs, ok = err.(PropertyErrors)
		if !ok {
			logWrongErrorType(PropertyErrors{}, err)
			return nil
		}
	}
	for i, element := range v {
		err = r.forEachRules.Validate(element)
		if err == nil {
			continue
		}
		forEachErrors, ok := err.(PropertyErrors)
		if !ok {
			logWrongErrorType(PropertyErrors{}, err)
			continue
		}
		for _, e := range forEachErrors {
			e.IsSliceElementError = true
			path := r.getJSONPathForIndex(i)
			propErrs = append(propErrs, e.prependParentPropertyName(path))
		}
	}
	if len(propErrs) > 0 {
		return propErrs.aggregate()
	}
	return nil
}

// WithName => refer to [PropertyRules.WithName] documentation.
func (r PropertyRulesForSlice[S, T, P]) WithName(name string) PropertyRulesForSlice[S, T, P] {
	r.sliceRules = r.sliceRules.WithName(name)
	return r
}

// WithExamples => refer to [PropertyRules.WithExamples] documentation.
func (r PropertyRulesForSlice[S, T, P]) WithExamples(examples ...string) PropertyRulesForSlice[S, T, P] {
	r.sliceRules = r.sliceRules.WithExamples(examples...)
	return r
}

// RulesForEach adds [Rule] for each element of the slice.
func (r PropertyRulesForSlice[S, T, P]) RulesForEach(rules ...rulesInterface[T]) PropertyRulesForSlice[S, T, P] {
	r.forEachRules = r.forEachRules.Rules(rules...)
	return r
}

// Rules adds [Rule] for the whole slice.
func (r PropertyRulesForSlice[S, T, P]) Rules(rules ...rulesInterface[S]) PropertyRulesForSlice[S, T, P] {
	r.sliceRules = r.sliceRules.Rules(rules...)
	return r
}

// When => refer to [PropertyRules.When] documentation.
func (r PropertyRulesForSlice[S, T, P]) When(
	predicate Predicate[P],
	opts ...WhenOptions,
) PropertyRulesForSlice[S, T, P] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

// IncludeForEach associates specified [Validator] and its [PropertyRules] with each element of the slice.
func (r PropertyRulesForSlice[S, T, P]) IncludeForEach(rules ...validatorInterface[T]) PropertyRulesForSlice[S, T, P] {
	r.forEachRules = r.forEachRules.Include(rules...)
	return r
}

// Include embeds specified [Validator] and its [PropertyRules] into the property.
func (r PropertyRulesForSlice[S, T, P]) Include(rules ...validatorInterface[S]) PropertyRulesForSlice[S, T, P] {
	r.sliceRules = r.sliceRules.Include(rules...)
	return r
}

// Cascade => refer to [PropertyRules.Cascade] documentation.
func (r PropertyRulesForSlice[S, T, P]) Cascade(mode CascadeMode) PropertyRulesForSlice[S, T, P] {
	r.mode = mode
	r.sliceRules = r.sliceRules.Cascade(mode)
	r.forEachRules = r.forEachRules.Cascade(mode)
	return r
}

// cascadeInternal is an internal wrapper around [PropertyRulesForMap.Cascade] which
// fulfills [propertyRulesInterface] interface.
// If the [CascadeMode] is already set, it won't change it.
func (r PropertyRulesForSlice[S, T, P]) cascadeInternal(mode CascadeMode) propertyRulesInterface[P] {
	if r.mode != 0 {
		return r
	}
	return r.Cascade(mode)
}

// plan generates a validation plan for the property rules.
func (r PropertyRulesForSlice[S, T, P]) plan(builder planBuilder) {
	builder = appendPredicatesToPlanBuilder(builder, r.predicates)
	r.sliceRules.plan(builder.setExamples(r.sliceRules.examples...))
	builder = builder.appendPath(r.sliceRules.name)
	if len(r.forEachRules.rules) > 0 {
		r.forEachRules.plan(builder.appendPath("[*]"))
	}
}

// getJSONPathForIndex returns a JSONPath for the given index.
func (r PropertyRulesForSlice[S, T, P]) getJSONPathForIndex(index int) string {
	return jsonpath.JoinArray(r.sliceRules.name, jsonpath.NewArrayIndex(index))
}

// isPropertyRules implements [propertyRulesInterface].
func (r PropertyRulesForSlice[S, T, P]) isPropertyRules() {}
