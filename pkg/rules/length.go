package rules

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringLength ensures the string's length is between min and max (closed interval).
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength].
//   - [govy.TemplateVars.MaxLength].
func StringLength(minLen, maxLen int) govy.Rule[string] {
	enforceMinMaxLength(minLen, maxLen)
	tpl := messagetemplates.Get(messagetemplates.StringLengthTemplate)

	return govy.NewRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length < minLen || length > maxLen {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: v,
				MinLength:     minLen,
				MaxLength:     maxLen,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			MinLength: minLen,
			MaxLength: maxLen,
		}))
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
	enforceMinMaxLength(minLen, maxLen)
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
	enforceMinMaxLength(minLen, maxLen)
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

func enforceMinMaxLength(minLen, maxLen int) {
	if minLen > maxLen {
		panic(fmt.Sprintf("minLen '%d' is greater than maxLen '%d'", minLen, maxLen))
	}
}
