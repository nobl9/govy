package govy

import (
	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/jsonpath"
)

var (
	mapKeyWildcardPath   = jsonpath.ParsePath("~")
	mapValueWildcardPath = jsonpath.ParsePath("*")
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
	inferPathMode InferPathMode

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
		keyPath := r.getPathForKey(k)
		if err = r.forKeyRules.Validate(k); err != nil {
			if keyErrors, ok := err.(PropertyErrors); ok {
				for _, e := range keyErrors {
					e.IsKeyError = true
					propErrs = append(propErrs, e.prependParentPropertyPath(keyPath))
				}
			} else {
				logWrongErrorType(PropertyErrors{}, err)
			}
		}
		if err = r.forValueRules.Validate(v); err != nil {
			if valueErrors, ok := err.(PropertyErrors); ok {
				for _, e := range valueErrors {
					propErrs = append(propErrs, e.prependParentPropertyPath(keyPath))
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
					propErrs = append(propErrs, e.prependParentPropertyPath(keyPath))
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

// WithName => refer to [PropertyRules.WithName] documentation.
func (r PropertyRulesForMap[M, K, V, P]) WithName(name string) PropertyRulesForMap[M, K, V, P] {
	r.mapRules = r.mapRules.WithName(name)
	return r
}

// WithPath => refer to [PropertyRules.WithPath] documentation.
func (r PropertyRulesForMap[M, K, V, P]) WithPath(path Path) PropertyRulesForMap[M, K, V, P] {
	r.mapRules = r.mapRules.WithPath(path)
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

// InferPath => refer to [PropertyRules.InferPath] documentation.
func (r PropertyRulesForMap[M, K, V, P]) InferPath(mode InferPathMode) PropertyRulesForMap[M, K, V, P] {
	r.inferPathMode = mode
	r.mapRules = r.mapRules.InferPath(mode)
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

// inferPathModeInternal is an internal wrapper around [PropertyRulesForMap.InferPath] which
// fulfills [PropertyRulesInterface] interface.
// If the [InferPathMode] is already set, it won't change it.
func (r PropertyRulesForMap[M, K, V, P]) inferPathModeInternal(mode InferPathMode) PropertyRulesInterface[P] {
	if r.inferPathMode != 0 {
		return r
	}
	return r.InferPath(mode)
}

// plan constructs a validation plan for the property rules.
func (r PropertyRulesForMap[M, K, V, P]) plan(builder planBuilder) {
	builder = appendPredicatesToPlanBuilder(builder, r.predicates)
	r.mapRules.plan(builder.setExamples(r.mapRules.examples...))
	builder = builder.appendPath(r.mapRules.getPath())
	if len(r.forKeyRules.rules) > 0 {
		r.forKeyRules.plan(builder.appendPath(mapKeyWildcardPath))
	}
	if len(r.forValueRules.rules) > 0 {
		r.forValueRules.plan(builder.appendPath(mapValueWildcardPath))
	}
	if len(r.forItemRules.rules) > 0 {
		r.forItemRules.plan(builder.appendPath(mapValueWildcardPath))
	}
}

func (r PropertyRulesForMap[M, K, V, P]) getPathForKey(key any) Path {
	return r.mapRules.getPath().Key(key)
}

func (r PropertyRulesForMap[M, K, V, P]) getPath() Path {
	return r.mapRules.getPath()
}

// isPropertyRules implements [PropertyRulesInterface].
func (r PropertyRulesForMap[M, K, V, P]) isPropertyRules() {}
