package govy

import (
	"fmt"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/jsonpath"
)

// ForMap creates a new [PropertyRulesForMap] instance for a map property
// which value is extracted through [PropertyGetter] function.
func ForMap[M ~map[K]V, K comparable, V, P any](getter PropertyGetter[M, P]) PropertyRulesForMap[M, K, V, P] {
	return PropertyRulesForMap[M, K, V, P]{
		mapRules:      forConstructor(getter),
		forKeyRules:   forConstructorWithoutNameInference(GetSelf[K]()),
		forValueRules: forConstructorWithoutNameInference(GetSelf[V]()),
		forItemRules:  forConstructorWithoutNameInference(GetSelf[MapItem[K, V]]()),
		getter:        getter,
	}
}

// PropertyRulesForMap is responsible for validating a single property.
type PropertyRulesForMap[M ~map[K]V, K comparable, V, P any] struct {
	mapRules      PropertyRules[M, P]
	forKeyRules   PropertyRules[K, K]
	forValueRules PropertyRules[V, V]
	forItemRules  PropertyRules[MapItem[K, V], MapItem[K, V]]
	getter        PropertyGetter[M, P]
	cascadeMode   CascadeMode

	predicateMatcher[P]
}

// MapItem is a tuple container for map's key and value pair.
type MapItem[K comparable, V any] struct {
	Key   K
	Value V
}

// Validate executes each of the rules sequentially and aggregates the encountered errors.
func (r PropertyRulesForMap[M, K, V, P]) Validate(parent P) error {
	if !r.matchPredicates(parent) {
		return nil
	}
	err := r.mapRules.Validate(parent)
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
	for k, v := range r.getter(parent) {
		if err = r.forKeyRules.Validate(k); err != nil {
			if keyErrors, ok := err.(PropertyErrors); ok {
				for _, e := range keyErrors {
					e.IsKeyError = true
					propErrs = append(propErrs, e.prependParentPropertyName(r.getJSONPathForKey(k)))
				}
			} else {
				logWrongErrorType(PropertyErrors{}, err)
			}
		}
		if err = r.forValueRules.Validate(v); err != nil {
			if valueErrors, ok := err.(PropertyErrors); ok {
				for _, e := range valueErrors {
					propErrs = append(propErrs, e.prependParentPropertyName(r.getJSONPathForKey(k)))
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
					propErrs = append(propErrs, e.prependParentPropertyName(r.getJSONPathForKey(k)))
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
func (r PropertyRulesForMap[M, K, V, P]) WithName(name string) PropertyRulesForMap[M, K, V, P] {
	r.mapRules = r.mapRules.WithName(name)
	return r
}

// WithExamples => refer to [PropertyRules.WithExamples] documentation.
func (r PropertyRulesForMap[M, K, V, P]) WithExamples(examples ...string) PropertyRulesForMap[M, K, V, P] {
	r.mapRules = r.mapRules.WithExamples(examples...)
	return r
}

// RulesForKeys adds [Rule] for map's keys.
func (r PropertyRulesForMap[M, K, V, P]) RulesForKeys(
	rules ...RulesInterface[K],
) PropertyRulesForMap[M, K, V, P] {
	r.forKeyRules = r.forKeyRules.Rules(rules...)
	return r
}

// RulesForValues adds [Rule] for map's values.
func (r PropertyRulesForMap[M, K, V, P]) RulesForValues(
	rules ...RulesInterface[V],
) PropertyRulesForMap[M, K, V, P] {
	r.forValueRules = r.forValueRules.Rules(rules...)
	return r
}

// RulesForItems adds [Rule] for [MapItem].
// It allows validating both key and value in conjunction.
func (r PropertyRulesForMap[M, K, V, P]) RulesForItems(
	rules ...RulesInterface[MapItem[K, V]],
) PropertyRulesForMap[M, K, V, P] {
	r.forItemRules = r.forItemRules.Rules(rules...)
	return r
}

// Rules adds [Rule] for the whole map.
func (r PropertyRulesForMap[M, K, V, P]) Rules(rules ...RulesInterface[M]) PropertyRulesForMap[M, K, V, P] {
	r.mapRules = r.mapRules.Rules(rules...)
	return r
}

// When => refer to [PropertyRules.When] documentation.
func (r PropertyRulesForMap[M, K, V, P]) When(
	predicate Predicate[P],
	opts ...WhenOption,
) PropertyRulesForMap[M, K, V, P] {
	r.predicateMatcher = r.when(predicate, opts...)
	return r
}

// Include embeds specified [Validator] and its [PropertyRules] into the property.
func (r PropertyRulesForMap[M, K, V, P]) Include(
	validators ...ValidatorInterface[M],
) PropertyRulesForMap[M, K, V, P] {
	r.mapRules = r.mapRules.Include(validators...)
	return r
}

// IncludeForKeys associates specified [Validator] and its [PropertyRules] with map's keys.
func (r PropertyRulesForMap[M, K, V, P]) IncludeForKeys(
	validators ...ValidatorInterface[K],
) PropertyRulesForMap[M, K, V, P] {
	r.forKeyRules = r.forKeyRules.Include(validators...)
	return r
}

// IncludeForValues associates specified [Validator] and its [PropertyRules] with map's values.
func (r PropertyRulesForMap[M, K, V, P]) IncludeForValues(
	rules ...ValidatorInterface[V],
) PropertyRulesForMap[M, K, V, P] {
	r.forValueRules = r.forValueRules.Include(rules...)
	return r
}

// IncludeForItems associates specified [Validator] and its [PropertyRules] with [MapItem].
// It allows validating both key and value in conjunction.
func (r PropertyRulesForMap[M, K, V, P]) IncludeForItems(
	rules ...ValidatorInterface[MapItem[K, V]],
) PropertyRulesForMap[M, K, V, P] {
	r.forItemRules = r.forItemRules.Include(rules...)
	return r
}

// Cascade => refer to [PropertyRules.Cascade] documentation.
func (r PropertyRulesForMap[M, K, V, P]) Cascade(mode CascadeMode) PropertyRulesForMap[M, K, V, P] {
	r.cascadeMode = mode
	r.mapRules = r.mapRules.Cascade(mode)
	r.forKeyRules = r.forKeyRules.Cascade(mode)
	r.forValueRules = r.forValueRules.Cascade(mode)
	r.forItemRules = r.forItemRules.Cascade(mode)
	return r
}

// InferName => refer to [PropertyRules.InferName] documentation.
func (r PropertyRulesForMap[M, K, V, P]) InferName(mode InferNameMode) PropertyRulesForMap[M, K, V, P] {
	r.mapRules = r.mapRules.InferName(mode)
	return r
}

// cascadeInternal is an internal wrapper around [PropertyRulesForMap.Cascade] which
// fulfills [PropertyRulesInterface] interface.
// If the [CascadeMode] is already set, it won't change it.
func (r PropertyRulesForMap[M, K, V, P]) cascadeInternal(mode CascadeMode) PropertyRulesInterface[P] {
	if r.cascadeMode != 0 {
		return r
	}
	return r.Cascade(mode)
}

// inferNameModeInternal is an internal wrapper around [PropertyRulesForMap.InferName] which
// fulfills [PropertyRulesInterface] interface.
func (r PropertyRulesForMap[M, K, V, P]) inferNameModeInternal(mode InferNameMode) PropertyRulesInterface[P] {
	return r.InferName(mode)
}

// plan constructs a validation plan for the property rules.
func (r PropertyRulesForMap[M, K, V, P]) plan(builder planBuilder) {
	builder = appendPredicatesToPlanBuilder(builder, r.predicates)
	r.mapRules.plan(builder.setExamples(r.mapRules.examples...))
	builder = builder.appendPath(r.mapRules.getName())
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

// getJSONPathForKey returns a JSONPath for the given key.
func (r PropertyRulesForMap[M, K, V, P]) getJSONPathForKey(key any) string {
	return jsonpath.Join(r.mapRules.getName(), jsonpath.EscapeSegment(fmt.Sprint(key)))
}

// getName returns the name of the property.
func (r PropertyRulesForMap[M, K, V, P]) getName() string {
	return r.mapRules.getName()
}

// isPropertyRules implements [PropertyRulesInterface].
func (r PropertyRulesForMap[M, K, V, P]) isPropertyRules() {}
