package govy

import (
	"fmt"

	"github.com/nobl9/govy/internal"
)

// ForMap creates a new [PropertyRulesForMap] instance for a map property
// which value is extracted through [PropertyGetter] function.
func ForMap[M ~map[K]V, K comparable, V, S any](getter PropertyGetter[M, S]) PropertyRulesForMap[M, K, V, S] {
	name := inferName()
	return PropertyRulesForMap[M, K, V, S]{
		mapRules:      forConstructor(getter, name),
		forKeyRules:   forConstructor(GetSelf[K](), ""),
		forValueRules: forConstructor(GetSelf[V](), ""),
		forItemRules:  forConstructor(GetSelf[MapItem[K, V]](), ""),
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
func (r PropertyRulesForMap[M, K, V, S]) Validate(st S) error {
	if !r.matchPredicates(st) {
		return nil
	}
	err := r.mapRules.Validate(st)
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
	for k, v := range r.getter(st) {
		if err = r.forKeyRules.Validate(k); err != nil {
			if keyErrors, ok := err.(PropertyErrors); ok {
				for _, e := range keyErrors {
					e.IsKeyError = true
					propErrs = append(propErrs, e.PrependParentPropertyName(MapElementName(r.mapRules.name, k)))
				}
			} else {
				logWrongErrorType(PropertyErrors{}, err)
			}
		}
		if err = r.forValueRules.Validate(v); err != nil {
			if valueErrors, ok := err.(PropertyErrors); ok {
				for _, e := range valueErrors {
					propErrs = append(propErrs, e.PrependParentPropertyName(MapElementName(r.mapRules.name, k)))
				}
			} else {
				logWrongErrorType(PropertyErrors{}, err)
			}
		}
		if err = r.forItemRules.Validate(MapItem[K, V]{Key: k, Value: v}); err != nil {
			if itemErrors, ok := err.(PropertyErrors); ok {
				for _, e := range itemErrors {
					// TODO: Figure out how to handle custom PropertyErrors.
					// Custom errors' value for nested item will be overridden by the actual value.
					e.PropertyValue = internal.PropertyValueString(v)
					propErrs = append(propErrs, e.PrependParentPropertyName(MapElementName(r.mapRules.name, k)))
				}
			} else {
				logWrongErrorType(PropertyErrors{}, err)
			}
		}
	}
	if len(propErrs) > 0 {
		return propErrs.aggregate().sort()
	}
	return nil
}

// WithName => refer to [PropertyRules.When] documentation.
func (r PropertyRulesForMap[M, K, V, S]) WithName(name string) PropertyRulesForMap[M, K, V, S] {
	r.mapRules = r.mapRules.WithName(name)
	return r
}

// WithExamples => refer to [PropertyRules.WithExamples] documentation.
func (r PropertyRulesForMap[M, K, V, S]) WithExamples(examples ...string) PropertyRulesForMap[M, K, V, S] {
	r.mapRules = r.mapRules.WithExamples(examples...)
	return r
}

// RulesForKeys adds [Rule] for map's keys.
func (r PropertyRulesForMap[M, K, V, S]) RulesForKeys(
	rules ...validationInterface[K],
) PropertyRulesForMap[M, K, V, S] {
	r.forKeyRules = r.forKeyRules.Rules(rules...)
	return r
}

// RulesForValues adds [Rule] for map's values.
func (r PropertyRulesForMap[M, K, V, S]) RulesForValues(
	rules ...validationInterface[V],
) PropertyRulesForMap[M, K, V, S] {
	r.forValueRules = r.forValueRules.Rules(rules...)
	return r
}

// RulesForItems adds [Rule] for [MapItem].
// It allows validating both key and value in conjunction.
func (r PropertyRulesForMap[M, K, V, S]) RulesForItems(
	rules ...validationInterface[MapItem[K, V]],
) PropertyRulesForMap[M, K, V, S] {
	r.forItemRules = r.forItemRules.Rules(rules...)
	return r
}

// Rules adds [Rule] for the whole map.
func (r PropertyRulesForMap[M, K, V, S]) Rules(rules ...validationInterface[M]) PropertyRulesForMap[M, K, V, S] {
	r.mapRules = r.mapRules.Rules(rules...)
	return r
}

// When => refer to [PropertyRules.When] documentation.
func (r PropertyRulesForMap[M, K, V, S]) When(
	predicate Predicate[S],
	opts ...WhenOptions,
) PropertyRulesForMap[M, K, V, S] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

// IncludeForKeys associates specified [Validator] and its [PropertyRules] with map's keys.
func (r PropertyRulesForMap[M, K, V, S]) IncludeForKeys(validators ...Validator[K]) PropertyRulesForMap[M, K, V, S] {
	r.forKeyRules = r.forKeyRules.Include(validators...)
	return r
}

// IncludeForValues associates specified [Validator] and its [PropertyRules] with map's values.
func (r PropertyRulesForMap[M, K, V, S]) IncludeForValues(rules ...Validator[V]) PropertyRulesForMap[M, K, V, S] {
	r.forValueRules = r.forValueRules.Include(rules...)
	return r
}

// IncludeForItems associates specified [Validator] and its [PropertyRules] with [MapItem].
// It allows validating both key and value in conjunction.
func (r PropertyRulesForMap[M, K, V, S]) IncludeForItems(
	rules ...Validator[MapItem[K, V]],
) PropertyRulesForMap[M, K, V, S] {
	r.forItemRules = r.forItemRules.Include(rules...)
	return r
}

// Cascade => refer to [PropertyRules.Cascade] documentation.
func (r PropertyRulesForMap[M, K, V, S]) Cascade(mode CascadeMode) PropertyRulesForMap[M, K, V, S] {
	r.mode = mode
	r.mapRules = r.mapRules.Cascade(mode)
	r.forKeyRules = r.forKeyRules.Cascade(mode)
	r.forValueRules = r.forValueRules.Cascade(mode)
	r.forItemRules = r.forItemRules.Cascade(mode)
	return r
}

// cascadeInternal is an internal wrapper around [PropertyRulesForMap.Cascade] which
// fulfills [propertyRulesInterface] interface.
// If the [CascadeMode] is already set, it won't change it.
func (r PropertyRulesForMap[M, K, V, S]) cascadeInternal(mode CascadeMode) propertyRulesInterface[S] {
	if r.mode != 0 {
		return r
	}
	return r.Cascade(mode)
}

// plan constructs a validation plan for the property rules.
func (r PropertyRulesForMap[M, K, V, S]) plan(builder planBuilder) {
	for _, predicate := range r.predicates {
		builder.rulePlan.Conditions = append(builder.rulePlan.Conditions, predicate.description)
	}
	r.mapRules.plan(builder.setExamples(r.mapRules.examples...))
	builder = builder.appendPath(r.mapRules.name)
	// JSON/YAML path for keys uses '~' to extract the keys.
	if len(r.forKeyRules.rules) > 0 {
		r.forKeyRules.plan(builder.appendPath("~"))
	}
	if len(r.forValueRules.rules) > 0 {
		r.forValueRules.plan(builder.appendPath("*"))
	}
	if len(r.forItemRules.rules) > 0 {
		r.forItemRules.plan(builder.appendPath("*"))
	}
}

// MapElementName generates a name for a map element denoted by its key.
func MapElementName(mapName, key any) string {
	if mapName == "" {
		return fmt.Sprintf("%v", key)
	}
	return fmt.Sprintf("%s.%v", mapName, key)
}
