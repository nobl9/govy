package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/exp/constraints"

	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/pkg/govy"
)

// EQ ensures the property's value is equal to the compared value.
func EQ[T comparable](compared T) govy.Rule[T] {
	msg := fmt.Sprintf(comparisonFmt, cmpEqualTo, compared)
	return govy.NewRule(func(v T) error {
		if v != compared {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeEqualTo).
		WithDescription(msg)
}

// NEQ ensures the property's value is not equal to the compared value.
func NEQ[T comparable](compared T) govy.Rule[T] {
	msg := fmt.Sprintf(comparisonFmt, cmpNotEqualTo, compared)
	return govy.NewRule(func(v T) error {
		if v == compared {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeNotEqualTo).
		WithDescription(msg)
}

// GT ensures the property's value is greater than the compared value.
func GT[T constraints.Ordered](compared T) govy.Rule[T] {
	return orderedComparisonRule(cmpGreaterThan, compared).
		WithErrorCode(ErrorCodeGreaterThan)
}

// GTE ensures the property's value is greater than or equal to the compared value.
func GTE[T constraints.Ordered](compared T) govy.Rule[T] {
	return orderedComparisonRule(cmpGreaterThanOrEqual, compared).
		WithErrorCode(ErrorCodeGreaterThanOrEqualTo)
}

// LT ensures the property's value is less than the compared value.
func LT[T constraints.Ordered](compared T) govy.Rule[T] {
	return orderedComparisonRule(cmpLessThan, compared).
		WithErrorCode(ErrorCodeLessThan)
}

// LTE ensures the property's value is less than or equal to the compared value.
func LTE[T constraints.Ordered](compared T) govy.Rule[T] {
	return orderedComparisonRule(cmpLessThanOrEqual, compared).
		WithErrorCode(ErrorCodeLessThanOrEqualTo)
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

// EqualProperties checks if all of the specified properties are equal.
// It uses the provided [ComparisonFunc] to compare the values.
// The following built-in comparison functions are available:
//   - [CompareFunc]
//   - [CompareDeepEqualFunc]
//
// If builtin [ComparisonFunc] are not enough, a custom function can be used.
func EqualProperties[S, T any](compare ComparisonFunc[T], getters map[string]func(s S) T) govy.Rule[S] {
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
				return fmt.Errorf(
					"all of %s properties must be equal, but '%s' is not equal to '%s'",
					prettyOneOfList(collections.SortedKeys(getters)),
					lastProp,
					prop,
				)
			}
			lastProp = prop
			lastValue = v
			i++
		}
		return nil
	}).
		WithErrorCode(ErrorCodeEqualProperties).
		WithDescription(func() string {
			return fmt.Sprintf(
				"all of the properties must be equal: %s",
				strings.Join(collections.SortedKeys(getters), ", "),
			)
		}())
}

var comparisonFmt = "should be %s '%v'"

func orderedComparisonRule[T constraints.Ordered](op comparisonOperator, compared T) govy.Rule[T] {
	msg := fmt.Sprintf(comparisonFmt, op, compared)
	return govy.NewRule(func(v T) error {
		var passed bool
		switch op {
		case cmpGreaterThan:
			passed = v > compared
		case cmpGreaterThanOrEqual:
			passed = v >= compared
		case cmpLessThan:
			passed = v < compared
		case cmpLessThanOrEqual:
			passed = v <= compared
		default:
			passed = false
		}
		if !passed {
			return errors.New(msg)
		}
		return nil
	}).WithDescription(msg)
}

type comparisonOperator uint8

const (
	cmpEqualTo comparisonOperator = iota
	cmpNotEqualTo
	cmpGreaterThan
	cmpGreaterThanOrEqual
	cmpLessThan
	cmpLessThanOrEqual
)

func (c comparisonOperator) String() string {
	// exhaustive: enforce
	switch c {
	case cmpEqualTo:
		return "equal to"
	case cmpNotEqualTo:
		return "not equal to"
	case cmpGreaterThan:
		return "greater than"
	case cmpGreaterThanOrEqual:
		return "greater than or equal to"
	case cmpLessThan:
		return "less than"
	case cmpLessThanOrEqual:
		return "less than or equal to"
	default:
		return "unknown"
	}
}
