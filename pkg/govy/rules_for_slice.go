package govy

import (
	"github.com/nobl9/govy/pkg/jsonpath"
)

var sliceWildcardPath = jsonpath.Parse("[*]")

// ForSlice creates a new [PropertyRulesForSlice] instance for a slice property
// which value is extracted through [PropertyGetter] function.
func ForSlice[S ~[]T, T, P any](getter PropertyGetter[S, P]) PropertyRulesForSlice[S, T, P] {
	return PropertyRulesForSlice[S, T, P]{
		sliceRules:   forConstructor(GetSelf[S]()),
		forEachRules: forConstructorWithoutPathInference(GetSelf[T]()),
		getter:       getter,
	}
}

// PropertyRulesForSlice is responsible for validating a single property.
type PropertyRulesForSlice[S ~[]T, T, P any] struct {
	sliceRules    PropertyRules[S, S]
	forEachRules  PropertyRules[T, T]
	getter        PropertyGetter[S, P]
	cascadeMode   CascadeMode
	inferPathMode InferPathMode

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
		if r.cascadeMode == CascadeModeStop {
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
			propErrs = append(propErrs, e.prependParentPropertyPath(r.getPathForIndex(i)))
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

// WithPath => refer to [PropertyRules.WithPath] documentation.
func (r PropertyRulesForSlice[S, T, P]) WithPath(path jsonpath.Path) PropertyRulesForSlice[S, T, P] {
	r.sliceRules = r.sliceRules.WithPath(path)
	return r
}

// WithExamples => refer to [PropertyRules.WithExamples] documentation.
func (r PropertyRulesForSlice[S, T, P]) WithExamples(examples ...string) PropertyRulesForSlice[S, T, P] {
	r.sliceRules = r.sliceRules.WithExamples(examples...)
	return r
}

// RulesForEach adds [Rule] for each element of the slice.
func (r PropertyRulesForSlice[S, T, P]) RulesForEach(rules ...RulesInterface[T]) PropertyRulesForSlice[S, T, P] {
	r.forEachRules = r.forEachRules.Rules(rules...)
	return r
}

// Rules adds [Rule] for the whole slice.
func (r PropertyRulesForSlice[S, T, P]) Rules(rules ...RulesInterface[S]) PropertyRulesForSlice[S, T, P] {
	r.sliceRules = r.sliceRules.Rules(rules...)
	return r
}

// When => refer to [PropertyRules.When] documentation.
func (r PropertyRulesForSlice[S, T, P]) When(
	predicate Predicate[P],
	opts ...WhenOption,
) PropertyRulesForSlice[S, T, P] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

// IncludeForEach associates specified [Validator] and its [PropertyRules] with each element of the slice.
func (r PropertyRulesForSlice[S, T, P]) IncludeForEach(rules ...ValidatorInterface[T]) PropertyRulesForSlice[S, T, P] {
	r.forEachRules = r.forEachRules.Include(rules...)
	return r
}

// Include embeds specified [Validator] and its [PropertyRules] into the property.
func (r PropertyRulesForSlice[S, T, P]) Include(rules ...ValidatorInterface[S]) PropertyRulesForSlice[S, T, P] {
	r.sliceRules = r.sliceRules.Include(rules...)
	return r
}

// Cascade => refer to [PropertyRules.Cascade] documentation.
func (r PropertyRulesForSlice[S, T, P]) Cascade(mode CascadeMode) PropertyRulesForSlice[S, T, P] {
	r.cascadeMode = mode
	r.sliceRules = r.sliceRules.Cascade(mode)
	r.forEachRules = r.forEachRules.Cascade(mode)
	return r
}

// InferPath => refer to [PropertyRules.InferPath] documentation.
func (r PropertyRulesForSlice[S, T, P]) InferPath(mode InferPathMode) PropertyRulesForSlice[S, T, P] {
	r.inferPathMode = mode
	r.sliceRules = r.sliceRules.InferPath(mode)
	return r
}

// cascadeInternal is an internal wrapper around [PropertyRulesForSlice.Cascade] which
// fulfills [PropertyRulesInterface] interface.
// If the [CascadeMode] is already set, it won't change it.
func (r PropertyRulesForSlice[S, T, P]) cascadeInternal(mode CascadeMode) PropertyRulesInterface[P] {
	if r.cascadeMode != 0 {
		return r
	}
	return r.Cascade(mode)
}

// inferPathModeInternal is an internal wrapper around [PropertyRulesForSlice.InferPath] which
// fulfills [PropertyRulesInterface] interface.
// If the [InferPathMode] is already set, it won't change it.
func (r PropertyRulesForSlice[S, T, P]) inferPathModeInternal(mode InferPathMode) PropertyRulesInterface[P] {
	if r.inferPathMode != 0 {
		return r
	}
	return r.InferPath(mode)
}

// plan generates a validation plan for the property rules.
func (r PropertyRulesForSlice[S, T, P]) plan(builder planBuilder) {
	builder = appendPredicatesToPlanBuilder(builder, r.predicates)
	r.sliceRules.plan(builder.setExamples(r.sliceRules.examples...))
	builder = builder.appendPath(r.sliceRules.getPath())
	if len(r.forEachRules.rules) > 0 {
		r.forEachRules.plan(builder.appendPath(sliceWildcardPath))
	}
}

func (r PropertyRulesForSlice[S, T, P]) getPathForIndex(index int) jsonpath.Path {
	return r.sliceRules.getPath().Index(uint(index)) // #nosec G115
}

func (r PropertyRulesForSlice[S, T, P]) getPath() jsonpath.Path {
	return r.sliceRules.getPath()
}

// isPropertyRules implements [PropertyRulesInterface].
func (r PropertyRulesForSlice[S, T, P]) isPropertyRules() {}
