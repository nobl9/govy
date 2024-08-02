package validation

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/exp/constraints"
)

func EqualTo[T comparable](compared T) SingleRule[T] {
	msg := fmt.Sprintf(comparisonFmt, cmpEqualTo, compared)
	return NewSingleRule(func(v T) error {
		if v != compared {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeEqualTo).
		WithDescription(msg)
}

func NotEqualTo[T comparable](compared T) SingleRule[T] {
	msg := fmt.Sprintf(comparisonFmt, cmpNotEqualTo, compared)
	return NewSingleRule(func(v T) error {
		if v == compared {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeNotEqualTo).
		WithDescription(msg)
}

func GreaterThan[T constraints.Ordered](n T) SingleRule[T] {
	return orderedComparisonRule(cmpGreaterThan, n).
		WithErrorCode(ErrorCodeGreaterThan)
}

func GreaterThanOrEqualTo[T constraints.Ordered](n T) SingleRule[T] {
	return orderedComparisonRule(cmpGreaterThanOrEqual, n).
		WithErrorCode(ErrorCodeGreaterThanOrEqualTo)
}

func LessThan[T constraints.Ordered](n T) SingleRule[T] {
	return orderedComparisonRule(cmpLessThan, n).
		WithErrorCode(ErrorCodeLessThan)
}

func LessThanOrEqualTo[T constraints.Ordered](n T) SingleRule[T] {
	return orderedComparisonRule(cmpLessThanOrEqual, n).
		WithErrorCode(ErrorCodeLessThanOrEqualTo)
}

var comparisonFmt = "should be %s '%v'"

func orderedComparisonRule[T constraints.Ordered](op comparisonOperator, compared T) SingleRule[T] {
	msg := fmt.Sprintf(comparisonFmt, op, compared)
	return NewSingleRule(func(v T) error {
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
	//exhaustive: enforce
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
