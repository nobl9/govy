package rules

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/nobl9/govy/pkg/govy"
)

// StringLength ensures the string's length is between lower and upper bound (closed interval).
func StringLength(lower, upper int) govy.Rule[string] {
	msg := fmt.Sprintf("length must be between %d and %d", lower, upper)
	return govy.NewRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length < lower || length > upper {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringLength).
		WithDescription(msg)
}

// StringMinLength ensures the string's length is greater than or equal to the limit.
func StringMinLength(limit int) govy.Rule[string] {
	msg := fmt.Sprintf("length must be %s %d", cmpGreaterThanOrEqual, limit)
	return govy.NewRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length < limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMinLength).
		WithDescription(msg)
}

// StringMaxLength ensures the string's length is less than or equal to the limit.
func StringMaxLength(limit int) govy.Rule[string] {
	msg := fmt.Sprintf("length must be %s %d", cmpLessThanOrEqual, limit)
	return govy.NewRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length > limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMaxLength).
		WithDescription(msg)
}

// SliceLength ensures the slice's length is between lower and upper bound (closed interval).
func SliceLength[S ~[]E, E any](lower, upper int) govy.Rule[S] {
	msg := fmt.Sprintf("length must be between %d and %d", lower, upper)
	return govy.NewRule(func(v S) error {
		length := len(v)
		if length < lower || length > upper {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceLength).
		WithDescription(msg)
}

// SliceMinLength ensures the slice's length is greater than or equal to the limit.
func SliceMinLength[S ~[]E, E any](limit int) govy.Rule[S] {
	msg := fmt.Sprintf("length must be %s %d", cmpGreaterThanOrEqual, limit)
	return govy.NewRule(func(v S) error {
		length := len(v)
		if length < limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceMinLength).
		WithDescription(msg)
}

// SliceMaxLength ensures the slice's length is less than or equal to the limit.
func SliceMaxLength[S ~[]E, E any](limit int) govy.Rule[S] {
	msg := fmt.Sprintf("length must be %s %d", cmpLessThanOrEqual, limit)
	return govy.NewRule(func(v S) error {
		length := len(v)
		if length > limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceMaxLength).
		WithDescription(msg)
}

// MapLength ensures the map's length is between lower and upper bound (closed interval).
func MapLength[M ~map[K]V, K comparable, V any](lower, upper int) govy.Rule[M] {
	msg := fmt.Sprintf("length must be between %d and %d", lower, upper)
	return govy.NewRule(func(v M) error {
		length := len(v)
		if length < lower || length > upper {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapLength).
		WithDescription(msg)
}

// MapMinLength ensures the map's length is greater than or equal to the limit.
func MapMinLength[M ~map[K]V, K comparable, V any](limit int) govy.Rule[M] {
	msg := fmt.Sprintf("length must be %s %d", cmpGreaterThanOrEqual, limit)
	return govy.NewRule(func(v M) error {
		length := len(v)
		if length < limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapMinLength).
		WithDescription(msg)
}

// MapMaxLength ensures the map's length is less than or equal to the limit.
func MapMaxLength[M ~map[K]V, K comparable, V any](limit int) govy.Rule[M] {
	msg := fmt.Sprintf("length must be %s %d", cmpLessThanOrEqual, limit)
	return govy.NewRule(func(v M) error {
		length := len(v)
		if length > limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapMaxLength).
		WithDescription(msg)
}
