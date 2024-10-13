package rules

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"

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
