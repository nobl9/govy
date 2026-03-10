package rules

import (
	"cmp"
	"fmt"
	"reflect"
	"strings"

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
		})).
		WithPlanModifiers(govy.RulePlanModifierValidValues(compared))
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
func GT[T cmp.Ordered](compared T) govy.Rule[T] {
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
func GTE[T cmp.Ordered](compared T) govy.Rule[T] {
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
func LT[T cmp.Ordered](compared T) govy.Rule[T] {
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
func LTE[T cmp.Ordered](compared T) govy.Rule[T] {
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
func EqualProperties[T, P any](compare ComparisonFunc[T], getters map[string]func(parent P) T) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.EqualPropertiesTemplate)

	sortedKeys := collections.SortedKeys(getters)
	return govy.NewRule(func(parent P) error {
		if len(getters) < 2 {
			return nil
		}
		var (
			i         = 0
			lastValue T
			lastProp  string
		)
		for _, prop := range sortedKeys {
			v := getters[prop](parent)
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

// GTProperties ensures the first property's value is greater than the second property's value.
// It works with [cmp.Ordered] types (int, float64, string, etc.).
func GTProperties[T cmp.Ordered, P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.GTPropertiesTemplate)

	return govy.NewRule(func(parent P) error {
		v1 := firstGetter(parent)
		v2 := secondGetter(parent)
		if cmp.Compare(v1, v2) <= 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Custom: comparePropertiesTemplateVars{
					FirstProperty:  firstName,
					SecondProperty: secondName,
				},
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeGTProperties).
		WithMessageTemplate(tpl).
		WithDescription(fmt.Sprintf("'%s' must be greater than '%s'", firstName, secondName))
}

// GTEProperties ensures the first property's value is greater than or equal to the second property's value.
// It works with [cmp.Ordered] types (int, float64, string, etc.).
func GTEProperties[T cmp.Ordered, P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.GTEPropertiesTemplate)

	return govy.NewRule(func(parent P) error {
		v1 := firstGetter(parent)
		v2 := secondGetter(parent)
		if cmp.Compare(v1, v2) < 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Custom: comparePropertiesTemplateVars{
					FirstProperty:  firstName,
					SecondProperty: secondName,
				},
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeGTEProperties).
		WithMessageTemplate(tpl).
		WithDescription(fmt.Sprintf("'%s' must be greater than or equal to '%s'", firstName, secondName))
}

// LTProperties ensures the first property's value is less than the second property's value.
// It works with [cmp.Ordered] types (int, float64, string, etc.).
func LTProperties[T cmp.Ordered, P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.LTPropertiesTemplate)

	return govy.NewRule(func(parent P) error {
		v1 := firstGetter(parent)
		v2 := secondGetter(parent)
		if cmp.Compare(v1, v2) >= 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Custom: comparePropertiesTemplateVars{
					FirstProperty:  firstName,
					SecondProperty: secondName,
				},
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeLTProperties).
		WithMessageTemplate(tpl).
		WithDescription(fmt.Sprintf("'%s' must be less than '%s'", firstName, secondName))
}

// LTEProperties ensures the first property's value is less than or equal to the second property's value.
// It works with [cmp.Ordered] types (int, float64, string, etc.).
func LTEProperties[T cmp.Ordered, P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.LTEPropertiesTemplate)

	return govy.NewRule(func(parent P) error {
		v1 := firstGetter(parent)
		v2 := secondGetter(parent)
		if cmp.Compare(v1, v2) > 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Custom: comparePropertiesTemplateVars{
					FirstProperty:  firstName,
					SecondProperty: secondName,
				},
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeLTEProperties).
		WithMessageTemplate(tpl).
		WithDescription(fmt.Sprintf("'%s' must be less than or equal to '%s'", firstName, secondName))
}

// Comparable is an interface for types that implement a Compare method.
// The Compare method should return:
//   - a negative value if the receiver is less than the argument
//   - zero if they are equal
//   - a positive value if the receiver is greater than the argument
//
// This method is implemented by types like [time.Time].
type Comparable[T any] interface {
	Compare(T) int
}

type comparePropertiesTemplateVars struct {
	FirstProperty  string
	SecondProperty string
}

// GTComparableProperties ensures the first property's value is greater than the second property's value.
// It works with types that implement the [Comparable] interface (types with a Compare method like [time.Time]).
func GTComparableProperties[T Comparable[T], P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.GTPropertiesTemplate)

	return govy.NewRule(func(parent P) error {
		v1 := firstGetter(parent)
		v2 := secondGetter(parent)
		if v1.Compare(v2) <= 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Custom: comparePropertiesTemplateVars{
					FirstProperty:  firstName,
					SecondProperty: secondName,
				},
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeGTProperties).
		WithMessageTemplate(tpl).
		WithDescription(fmt.Sprintf("'%s' must be greater than '%s'", firstName, secondName))
}

// GTEComparableProperties ensures the first property's value is greater than or equal to the second property's value.
// It works with types that implement the [Comparable] interface (types with a Compare method like [time.Time]).
func GTEComparableProperties[T Comparable[T], P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.GTEPropertiesTemplate)

	return govy.NewRule(func(parent P) error {
		v1 := firstGetter(parent)
		v2 := secondGetter(parent)
		if v1.Compare(v2) < 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Custom: comparePropertiesTemplateVars{
					FirstProperty:  firstName,
					SecondProperty: secondName,
				},
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeGTEProperties).
		WithMessageTemplate(tpl).
		WithDescription(fmt.Sprintf("'%s' must be greater than or equal to '%s'", firstName, secondName))
}

// LTComparableProperties ensures the first property's value is less than the second property's value.
// It works with types that implement the [Comparable] interface (types with a Compare method like [time.Time]).
func LTComparableProperties[T Comparable[T], P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.LTPropertiesTemplate)

	return govy.NewRule(func(parent P) error {
		v1 := firstGetter(parent)
		v2 := secondGetter(parent)
		if v1.Compare(v2) >= 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Custom: comparePropertiesTemplateVars{
					FirstProperty:  firstName,
					SecondProperty: secondName,
				},
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeLTProperties).
		WithMessageTemplate(tpl).
		WithDescription(fmt.Sprintf("'%s' must be less than '%s'", firstName, secondName))
}

// LTEComparableProperties ensures the first property's value is less than or equal to the second property's value.
// It works with types that implement the [Comparable] interface (types with a Compare method like [time.Time]).
func LTEComparableProperties[T Comparable[T], P any](
	firstName string,
	firstGetter func(parent P) T,
	secondName string,
	secondGetter func(parent P) T,
) govy.Rule[P] {
	tpl := messagetemplates.Get(messagetemplates.LTEPropertiesTemplate)

	return govy.NewRule(func(parent P) error {
		v1 := firstGetter(parent)
		v2 := secondGetter(parent)
		if v1.Compare(v2) > 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				Custom: comparePropertiesTemplateVars{
					FirstProperty:  firstName,
					SecondProperty: secondName,
				},
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeLTEProperties).
		WithMessageTemplate(tpl).
		WithDescription(fmt.Sprintf("'%s' must be less than or equal to '%s'", firstName, secondName))
}
