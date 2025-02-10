package rules

import (
	"fmt"
	"unicode/utf8"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringLength ensures the string's length is between min and max (closed interval).
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func StringLength(minLen, maxLen int) govy.Rule[string] {
	enforceMinMaxLength(minLen, maxLen)
	tpl := messagetemplates.Get(messagetemplates.LengthTemplate)

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
	tpl := messagetemplates.Get(messagetemplates.MinLengthTemplate)

	return govy.NewRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length < limit {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: limit,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMinLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}))
}

// StringMaxLength ensures the string's length is less than or equal to the limit.
func StringMaxLength(limit int) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.MaxLengthTemplate)

	return govy.NewRule(func(v string) error {
		length := utf8.RuneCountInString(v)
		if length > limit {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: limit,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMaxLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}))
}

// SliceLength ensures the slice's length is between min and max (closed interval).
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func SliceLength[S ~[]E, E any](minLen, maxLen int) govy.Rule[S] {
	enforceMinMaxLength(minLen, maxLen)
	tpl := messagetemplates.Get(messagetemplates.LengthTemplate)

	return govy.NewRule(func(v S) error {
		length := len(v)
		if length < minLen || length > maxLen {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: v,
				MinLength:     minLen,
				MaxLength:     maxLen,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			MinLength: minLen,
			MaxLength: maxLen,
		}))
}

// SliceMinLength ensures the slice's length is greater than or equal to the limit.
func SliceMinLength[S ~[]E, E any](limit int) govy.Rule[S] {
	tpl := messagetemplates.Get(messagetemplates.MinLengthTemplate)

	return govy.NewRule(func(v S) error {
		length := len(v)
		if length < limit {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: limit,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceMinLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}))
}

// SliceMaxLength ensures the slice's length is less than or equal to the limit.
func SliceMaxLength[S ~[]E, E any](limit int) govy.Rule[S] {
	tpl := messagetemplates.Get(messagetemplates.MaxLengthTemplate)

	return govy.NewRule(func(v S) error {
		length := len(v)
		if length > limit {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: limit,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeSliceMaxLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}))
}

// MapLength ensures the map's length is between min and max (closed interval).
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func MapLength[M ~map[K]V, K comparable, V any](minLen, maxLen int) govy.Rule[M] {
	enforceMinMaxLength(minLen, maxLen)
	tpl := messagetemplates.Get(messagetemplates.LengthTemplate)

	return govy.NewRule(func(v M) error {
		length := len(v)
		if length < minLen || length > maxLen {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: v,
				MinLength:     minLen,
				MaxLength:     maxLen,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			MinLength: minLen,
			MaxLength: maxLen,
		}))
}

// MapMinLength ensures the map's length is greater than or equal to the limit.
func MapMinLength[M ~map[K]V, K comparable, V any](limit int) govy.Rule[M] {
	tpl := messagetemplates.Get(messagetemplates.MinLengthTemplate)

	return govy.NewRule(func(v M) error {
		length := len(v)
		if length < limit {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: limit,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapMinLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}))
}

// MapMaxLength ensures the map's length is less than or equal to the limit.
func MapMaxLength[M ~map[K]V, K comparable, V any](limit int) govy.Rule[M] {
	tpl := messagetemplates.Get(messagetemplates.MaxLengthTemplate)

	return govy.NewRule(func(v M) error {
		length := len(v)
		if length > limit {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   v,
				ComparisonValue: limit,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeMapMaxLength).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}))
}

func enforceMinMaxLength(minLen, maxLen int) {
	if minLen > maxLen {
		panic(fmt.Sprintf("minLen '%d' is greater than maxLen '%d'", minLen, maxLen))
	}
}
