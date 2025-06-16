package govy

import (
	"github.com/nobl9/govy/internal/jsonpath"
)

// ForSlice creates a new [PropertyRulesForSlice] instance for a slice property
// which value is extracted through [PropertyGetter] function.
func ForSlice[T, S any](getter PropertyGetter[[]T, S]) PropertyRulesForSlice[T, S] {
	name := inferName()
	return PropertyRulesForSlice[T, S]{
		sliceRules:   forConstructor(GetSelf[[]T](), name),
		forEachRules: forConstructor(GetSelf[T](), ""),
		getter:       getter,
	}
}

// PropertyRulesForSlice is responsible for validating a single property.
type PropertyRulesForSlice[T, S any] struct {
	sliceRules   PropertyRules[[]T, []T]
	forEachRules PropertyRules[T, T]
	getter       PropertyGetter[[]T, S]
	mode         CascadeMode

	predicateMatcher[S]
}

// Validate executes each of the rules sequentially and aggregates the encountered errors.
func (r PropertyRulesForSlice[T, S]) Validate(st S) error {
	if !r.matchPredicates(st) {
		return nil
	}
	v := r.getter(st)
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
func (r PropertyRulesForSlice[T, S]) WithName(name string) PropertyRulesForSlice[T, S] {
	r.sliceRules = r.sliceRules.WithName(name)
	return r
}

// WithExamples => refer to [PropertyRules.WithExamples] documentation.
func (r PropertyRulesForSlice[T, S]) WithExamples(examples ...string) PropertyRulesForSlice[T, S] {
	r.sliceRules = r.sliceRules.WithExamples(examples...)
	return r
}

// RulesForEach adds [Rule] for each element of the slice.
func (r PropertyRulesForSlice[T, S]) RulesForEach(rules ...validationInterface[T]) PropertyRulesForSlice[T, S] {
	r.forEachRules = r.forEachRules.Rules(rules...)
	return r
}

// Rules adds [Rule] for the whole slice.
func (r PropertyRulesForSlice[T, S]) Rules(rules ...validationInterface[[]T]) PropertyRulesForSlice[T, S] {
	r.sliceRules = r.sliceRules.Rules(rules...)
	return r
}

// When => refer to [PropertyRules.When] documentation.
func (r PropertyRulesForSlice[T, S]) When(predicate Predicate[S], opts ...WhenOptions) PropertyRulesForSlice[T, S] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

// IncludeForEach associates specified [Validator] and its [PropertyRules] with each element of the slice.
func (r PropertyRulesForSlice[T, S]) IncludeForEach(rules ...Validator[T]) PropertyRulesForSlice[T, S] {
	r.forEachRules = r.forEachRules.Include(rules...)
	return r
}

// Cascade => refer to [PropertyRules.Cascade] documentation.
func (r PropertyRulesForSlice[T, S]) Cascade(mode CascadeMode) PropertyRulesForSlice[T, S] {
	r.mode = mode
	r.sliceRules = r.sliceRules.Cascade(mode)
	r.forEachRules = r.forEachRules.Cascade(mode)
	return r
}

// cascadeInternal is an internal wrapper around [PropertyRulesForMap.Cascade] which
// fulfills [propertyRulesInterface] interface.
// If the [CascadeMode] is already set, it won't change it.
func (r PropertyRulesForSlice[T, S]) cascadeInternal(mode CascadeMode) propertyRulesInterface[S] {
	if r.mode != 0 {
		return r
	}
	return r.Cascade(mode)
}

// plan generates a validation plan for the property rules.
func (r PropertyRulesForSlice[T, S]) plan(builder planBuilder) {
	builder = appendPredicatesToPlanBuilder(builder, r.predicates)
	r.sliceRules.plan(builder.setExamples(r.sliceRules.examples...))
	builder = builder.appendPath(r.sliceRules.name)
	if len(r.forEachRules.rules) > 0 {
		r.forEachRules.plan(builder.appendPath("[*]"))
	}
}

// getJSONPathForIndex returns a JSONPath for the given index.
func (r PropertyRulesForSlice[T, S]) getJSONPathForIndex(index int) string {
	return jsonpath.JoinArray(r.sliceRules.name, jsonpath.NewArrayIndex(index))
}
