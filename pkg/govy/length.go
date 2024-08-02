package validation

import (
	"fmt"
	"unicode/utf8"

	"github.com/pkg/errors"
)

func StringLength(lower, upper int) SingleRule[string] {
	msg := fmt.Sprintf("length must be between %d and %d", lower, upper)
	return NewSingleRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length < lower || length > upper {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringLength).
		WithDescription(msg)
}

func StringMinLength(limit int) SingleRule[string] {
	msg := fmt.Sprintf("length must be %s %d", cmpGreaterThanOrEqual, limit)
	return NewSingleRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length < limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMinLength).
		WithDescription(msg)
}

func StringMaxLength(limit int) SingleRule[string] {
	msg := fmt.Sprintf("length must be %s %d", cmpLessThanOrEqual, limit)
	return NewSingleRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length > limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMaxLength).
		WithDescription(msg)
}

func SliceLength[S ~[]E, E any](lower, upper int) SingleRule[S] {
	msg := fmt.Sprintf("length must be between %d and %d", lower, upper)
	return NewSingleRule(func(v S) error {
		length := len(v)
		if length < lower || length > upper {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceLength).
		WithDescription(msg)
}

func SliceMinLength[S ~[]E, E any](limit int) SingleRule[S] {
	msg := fmt.Sprintf("length must be %s %d", cmpGreaterThanOrEqual, limit)
	return NewSingleRule(func(v S) error {
		length := len(v)
		if length < limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceMinLength).
		WithDescription(msg)
}

func SliceMaxLength[S ~[]E, E any](limit int) SingleRule[S] {
	msg := fmt.Sprintf("length must be %s %d", cmpLessThanOrEqual, limit)
	return NewSingleRule(func(v S) error {
		length := len(v)
		if length > limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceMaxLength).
		WithDescription(msg)
}

func MapLength[M ~map[K]V, K comparable, V any](lower, upper int) SingleRule[M] {
	msg := fmt.Sprintf("length must be between %d and %d", lower, upper)
	return NewSingleRule(func(v M) error {
		length := len(v)
		if length < lower || length > upper {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapLength).
		WithDescription(msg)
}

func MapMinLength[M ~map[K]V, K comparable, V any](limit int) SingleRule[M] {
	msg := fmt.Sprintf("length must be %s %d", cmpGreaterThanOrEqual, limit)
	return NewSingleRule(func(v M) error {
		length := len(v)
		if length < limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapMinLength).
		WithDescription(msg)
}

func MapMaxLength[M ~map[K]V, K comparable, V any](limit int) SingleRule[M] {
	msg := fmt.Sprintf("length must be %s %d", cmpLessThanOrEqual, limit)
	return NewSingleRule(func(v M) error {
		length := len(v)
		if length > limit {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapMaxLength).
		WithDescription(msg)
}
