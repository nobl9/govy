package validation

import (
	"fmt"
)

// ForMap creates a new [PropertyRulesForMap] instance for a map property
// which value is extracted through [PropertyGetter] function.
func ForMap[M ~map[K]V, K comparable, V, S any](getter PropertyGetter[M, S]) PropertyRulesForMap[M, K, V, S] {
	return PropertyRulesForMap[M, K, V, S]{
		mapRules:      For(getter),
		forKeyRules:   For(GetSelf[K]()),
		forValueRules: For(GetSelf[V]()),
		forItemRules:  For(GetSelf[MapItem[K, V]]()),
		getter:        getter,
	}
}

// PropertyRulesForMap is responsible for validating a single property.
type PropertyRulesForMap[M ~map[K]V, K comparable, V, S any] struct {
	mapRules      PropertyRules[M, S]
	forKeyRules   PropertyRules[K, K]
	forValueRules PropertyRules[V, V]
	forItemRules  PropertyRules[MapItem[K, V], MapItem[K, V]]
	getter        PropertyGetter[M, S]
	mode          CascadeMode

	predicateMatcher[S]
}

// MapItem is a tuple container for map's key and value pair.
type MapItem[K comparable, V any] struct {
	Key   K
	Value V
}

// Validate executes each of the rules sequentially and aggregates the encountered errors.
func (r PropertyRulesForMap[M, K, V, S]) Validate(st S) PropertyErrors {
	if !r.matchPredicates(st) {
		return nil
	}
	err := r.mapRules.Validate(st)
	if r.mode == CascadeModeStop && err != nil {
		return err
	}
	for k, v := range r.getter(st) {
		forKeyErr := r.forKeyRules.Validate(k)
		for _, e := range forKeyErr {
			e.IsKeyError = true
			err = append(err, e.PrependPropertyName(MapElementName(r.mapRules.name, k)))
		}
		forValueErr := r.forValueRules.Validate(v)
		for _, e := range forValueErr {
			err = append(err, e.PrependPropertyName(MapElementName(r.mapRules.name, k)))
		}
		forItemErr := r.forItemRules.Validate(MapItem[K, V]{Key: k, Value: v})
		for _, e := range forItemErr {
			// TODO: Figure out how to handle custom PropertyErrors.
			// Custom errors' value for nested item will be overridden by the actual value.
			e.PropertyValue = propertyValueString(v)
			err = append(err, e.PrependPropertyName(MapElementName(r.mapRules.name, k)))
		}
	}
	return err.Aggregate().Sort()
}

func (r PropertyRulesForMap[M, K, V, S]) WithName(name string) PropertyRulesForMap[M, K, V, S] {
	r.mapRules = r.mapRules.WithName(name)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) WithExamples(examples ...string) PropertyRulesForMap[M, K, V, S] {
	r.mapRules = r.mapRules.WithExamples(examples...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) RulesForKeys(rules ...Rule[K]) PropertyRulesForMap[M, K, V, S] {
	r.forKeyRules = r.forKeyRules.Rules(rules...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) RulesForValues(rules ...Rule[V]) PropertyRulesForMap[M, K, V, S] {
	r.forValueRules = r.forValueRules.Rules(rules...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) RulesForItems(rules ...Rule[MapItem[K, V]]) PropertyRulesForMap[M, K, V, S] {
	r.forItemRules = r.forItemRules.Rules(rules...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) Rules(rules ...Rule[M]) PropertyRulesForMap[M, K, V, S] {
	r.mapRules = r.mapRules.Rules(rules...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) When(
	predicate Predicate[S],
	opts ...WhenOptions,
) PropertyRulesForMap[M, K, V, S] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) IncludeForKeys(validators ...Validator[K]) PropertyRulesForMap[M, K, V, S] {
	r.forKeyRules = r.forKeyRules.Include(validators...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) IncludeForValues(rules ...Validator[V]) PropertyRulesForMap[M, K, V, S] {
	r.forValueRules = r.forValueRules.Include(rules...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) IncludeForItems(
	rules ...Validator[MapItem[K, V]],
) PropertyRulesForMap[M, K, V, S] {
	r.forItemRules = r.forItemRules.Include(rules...)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) Cascade(mode CascadeMode) PropertyRulesForMap[M, K, V, S] {
	r.mode = mode
	r.mapRules = r.mapRules.Cascade(mode)
	r.forKeyRules = r.forKeyRules.Cascade(mode)
	r.forValueRules = r.forValueRules.Cascade(mode)
	r.forItemRules = r.forItemRules.Cascade(mode)
	return r
}

func (r PropertyRulesForMap[M, K, V, S]) plan(builder planBuilder) {
	for _, predicate := range r.predicates {
		builder.rulePlan.Conditions = append(builder.rulePlan.Conditions, predicate.description)
	}
	r.mapRules.plan(builder.setExamples(r.mapRules.examples...))
	builder = builder.appendPath(r.mapRules.name)
	// JSON/YAML path for keys uses '~' to extract the keys.
	if len(r.forKeyRules.steps) > 0 {
		r.forKeyRules.plan(builder.appendPath("~"))
	}
	if len(r.forValueRules.steps) > 0 {
		r.forValueRules.plan(builder.appendPath("*"))
	}
	if len(r.forItemRules.steps) > 0 {
		r.forItemRules.plan(builder.appendPath("*"))
	}
}

func MapElementName(mapName, key any) string {
	if mapName == "" {
		return fmt.Sprintf("%v", key)
	}
	return fmt.Sprintf("%s.%v", mapName, key)
}
