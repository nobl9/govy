package rules

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/nobl9/govy/pkg/govy"
)

// StringLength ensures the string's length is between min and max (closed interval).
func StringLength(lower, upper int) govy.Rule[string] {
	msg := "length must be between {{ .MinLength }} and {{ .MaxLength }}"
	tpl := getMessageTemplate("StringLength", msg)

	return govy.NewRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length < lower || length > upper {
			return returnTemplatedError(tpl, func() templateVariables[string] {
				return templateVariables[string]{
					PropertyValue: v,
					MinLength:     lower,
					MaxLength:     upper,
				}
			})
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

// SliceLength ensures the slice's length is between min and max (closed interval).
func SliceLength[S ~[]E, E any](minLen, maxLen int) govy.Rule[S] {
	msg := fmt.Sprintf("length must be between %d and %d", minLen, maxLen)
	return govy.NewRule(func(v S) error {
		length := len(v)
		if length < minLen || length > maxLen {
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

// MapLength ensures the map's length is between min and max (closed interval).
func MapLength[M ~map[K]V, K comparable, V any](minLen, maxLen int) govy.Rule[M] {
	msg := fmt.Sprintf("length must be between %d and %d", minLen, maxLen)
	return govy.NewRule(func(v M) error {
		length := len(v)
		if length < minLen || length > maxLen {
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
