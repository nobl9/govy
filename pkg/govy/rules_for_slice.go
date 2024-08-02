package validation

import "fmt"

// ForSlice creates a new [PropertyRulesForSlice] instance for a slice property
// which value is extracted through [PropertyGetter] function.
func ForSlice[T, S any](getter PropertyGetter[[]T, S]) PropertyRulesForSlice[T, S] {
	return PropertyRulesForSlice[T, S]{
		sliceRules:   For(GetSelf[[]T]()),
		forEachRules: For(GetSelf[T]()),
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
func (r PropertyRulesForSlice[T, S]) Validate(st S) PropertyErrors {
	if !r.matchPredicates(st) {
		return nil
	}
	v := r.getter(st)
	err := r.sliceRules.Validate(v)
	if r.mode == CascadeModeStop && err != nil {
		return err
	}
	for i, element := range v {
		forEachErr := r.forEachRules.Validate(element)
		if forEachErr == nil {
			continue
		}
		for _, e := range forEachErr {
			e.IsSliceElementError = true
			err = append(err, e.PrependPropertyName(SliceElementName(r.sliceRules.name, i)))
		}
	}
	return err.Aggregate()
}

func (r PropertyRulesForSlice[T, S]) WithName(name string) PropertyRulesForSlice[T, S] {
	r.sliceRules = r.sliceRules.WithName(name)
	return r
}

func (r PropertyRulesForSlice[T, S]) WithExamples(examples ...string) PropertyRulesForSlice[T, S] {
	r.sliceRules = r.sliceRules.WithExamples(examples...)
	return r
}

func (r PropertyRulesForSlice[T, S]) RulesForEach(rules ...Rule[T]) PropertyRulesForSlice[T, S] {
	r.forEachRules = r.forEachRules.Rules(rules...)
	return r
}

func (r PropertyRulesForSlice[T, S]) Rules(rules ...Rule[[]T]) PropertyRulesForSlice[T, S] {
	r.sliceRules = r.sliceRules.Rules(rules...)
	return r
}

func (r PropertyRulesForSlice[T, S]) When(predicate Predicate[S], opts ...WhenOptions) PropertyRulesForSlice[T, S] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

func (r PropertyRulesForSlice[T, S]) IncludeForEach(rules ...Validator[T]) PropertyRulesForSlice[T, S] {
	r.forEachRules = r.forEachRules.Include(rules...)
	return r
}

func (r PropertyRulesForSlice[T, S]) Cascade(mode CascadeMode) PropertyRulesForSlice[T, S] {
	r.mode = mode
	r.sliceRules = r.sliceRules.Cascade(mode)
	r.forEachRules = r.forEachRules.Cascade(mode)
	return r
}

func (r PropertyRulesForSlice[T, S]) plan(builder planBuilder) {
	for _, predicate := range r.predicates {
		builder.rulePlan.Conditions = append(builder.rulePlan.Conditions, predicate.description)
	}
	r.sliceRules.plan(builder.setExamples(r.sliceRules.examples...))
	builder = builder.appendPath(r.sliceRules.name)
	if len(r.forEachRules.steps) > 0 {
		r.forEachRules.plan(builder.appendPath("[*]"))
	}
}

func SliceElementName(sliceName string, index int) string {
	if sliceName == "" {
		return fmt.Sprintf("[%d]", index)
	}
	return fmt.Sprintf("%s[%d]", sliceName, index)
}
