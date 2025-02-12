package rules

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"

	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// EQ ensures the property's value is equal to the compared value.
func EQ[T comparable](compared T) govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.EQTemplate)

	return govy.NewRule(func(v T) error {
		if v != compared {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: compared,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeEqualTo).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: compared,
		}))
}

// NEQ ensures the property's value is not equal to the compared value.
func NEQ[T comparable](compared T) govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.NEQTemplate)

	return govy.NewRule(func(v T) error {
		if v == compared {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: compared,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeNotEqualTo).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: compared,
		}))
}

// GT ensures the property's value is greater than the compared value.
func GT[T constraints.Ordered](compared T) govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.GTTemplate)

	return govy.NewRule(func(v T) error {
		if v <= compared {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: compared,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeGreaterThan).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: compared,
		}))
}

// GTE ensures the property's value is greater than or equal to the compared value.
func GTE[T constraints.Ordered](compared T) govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.GTETemplate)

	return govy.NewRule(func(v T) error {
		if v < compared {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: compared,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeGreaterThanOrEqualTo).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: compared,
		}))
}

// LT ensures the property's value is less than the compared value.
func LT[T constraints.Ordered](compared T) govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.LTTemplate)

	return govy.NewRule(func(v T) error {
		if v >= compared {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: compared,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeLessThan).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: compared,
		}))
}

// LTE ensures the property's value is less than or equal to the compared value.
func LTE[T constraints.Ordered](compared T) govy.Rule[T] {
	tpl := messagetemplates.Get(messagetemplates.LTETemplate)

	return govy.NewRule(func(v T) error {
		if v > compared {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: compared,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeLessThanOrEqualTo).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: compared,
		}))
}

// ComparisonFunc defines a shape for a function that compares two values.
// It should return true if the values are equal, false otherwise.
type ComparisonFunc[T any] func(v1, v2 T) bool

// CompareFunc compares two values of the same type.
// The type is constrained by the [comparable] interface.
func CompareFunc[T comparable](v1, v2 T) bool {
	return v1 == v2
}

// CompareDeepEqualFunc compares two values of the same type using [reflect.DeepEqual].
// It is particularly useful when comparing pointers' values.
func CompareDeepEqualFunc[T any](v1, v2 T) bool {
	return reflect.DeepEqual(v1, v2)
}

type equalPropertiesTemplateVars struct {
	FirstNotEqual  string
	SecondNotEqual string
}

// EqualProperties checks if all the specified properties are equal.
// It uses the provided [ComparisonFunc] to compare the values.
// The following built-in comparison functions are available:
//   - [CompareFunc]
//   - [CompareDeepEqualFunc]
//
// If builtin [ComparisonFunc] is not enough, a custom function can be used.
func EqualProperties[S, T any](compare ComparisonFunc[T], getters map[string]func(s S) T) govy.Rule[S] {
	tpl := messagetemplates.Get(messagetemplates.EqualPropertiesTemplate)

	sortedKeys := collections.SortedKeys(getters)
	return govy.NewRule(func(s S) error {
		if len(getters) < 2 {
			return nil
		}
		var (
			i         = 0
			lastValue T
			lastProp  string
		)
		for _, prop := range sortedKeys {
			v := getters[prop](s)
			if i != 0 && !compare(v, lastValue) {
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					ComparisonValue: sortedKeys,
					Custom: equalPropertiesTemplateVars{
						FirstNotEqual:  lastProp,
						SecondNotEqual: prop,
					},
				})
			}
			lastProp = prop
			lastValue = v
			i++
		}
		return nil
	}).
		WithErrorCode(ErrorCodeEqualProperties).
		WithMessageTemplate(tpl).
		WithDescription(func() string {
			return fmt.Sprintf(
				"all of the properties must be equal: %s",
				strings.Join(collections.SortedKeys(getters), ", "),
			)
		}())
}
