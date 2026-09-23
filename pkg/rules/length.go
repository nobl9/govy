package rules

import (
	"fmt"
	"unicode/utf8"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
)

// StringLength ensures the string's length is between min and max (closed interval).
// It panics if either bound is negative or minLen is greater than maxLen.
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func StringLength(minLen, maxLen int) govy.Rule[string] {
	schemaMinLen, schemaMaxLen := enforceMinMaxLength(minLen, maxLen)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			MinLength: minLen,
			MaxLength: maxLen,
		}).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{
				MinLength: ptr(schemaMinLen),
				MaxLength: ptr(schemaMaxLen),
			}, nil
		})
}

// StringMinLength ensures the string's length is greater than or equal to the limit.
// It panics if limit is negative.
func StringMinLength(limit int) govy.Rule[string] {
	schemaLimit := enforceLength(limit)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{MinLength: ptr(schemaLimit)}, nil
		})
}

// StringMaxLength ensures the string's length is less than or equal to the limit.
// It panics if limit is negative.
func StringMaxLength(limit int) govy.Rule[string] {
	schemaLimit := enforceLength(limit)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{MaxLength: ptr(schemaLimit)}, nil
		})
}

// SliceLength ensures the slice's length is between min and max (closed interval).
// It panics if either bound is negative or minLen is greater than maxLen.
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func SliceLength[S ~[]E, E any](minLen, maxLen int) govy.Rule[S] {
	schemaMinLen, schemaMaxLen := enforceMinMaxLength(minLen, maxLen)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			MinLength: minLen,
			MaxLength: maxLen,
		}).
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			if ctx.Type != jsonschema.TypeArray {
				return nil, nil
			}
			return &jsonschema.Schema{
				MinItems: ptr(schemaMinLen),
				MaxItems: ptr(schemaMaxLen),
			}, nil
		})
}

// SliceMinLength ensures the slice's length is greater than or equal to the limit.
// It panics if limit is negative.
func SliceMinLength[S ~[]E, E any](limit int) govy.Rule[S] {
	schemaLimit := enforceLength(limit)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}).
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			if ctx.Type != jsonschema.TypeArray {
				return nil, nil
			}
			return &jsonschema.Schema{MinItems: ptr(schemaLimit)}, nil
		})
}

// SliceMaxLength ensures the slice's length is less than or equal to the limit.
// It panics if limit is negative.
func SliceMaxLength[S ~[]E, E any](limit int) govy.Rule[S] {
	schemaLimit := enforceLength(limit)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}).
		WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			if ctx.Type != jsonschema.TypeArray {
				return nil, nil
			}
			return &jsonschema.Schema{MaxItems: ptr(schemaLimit)}, nil
		})
}

// MapLength ensures the map's length is between min and max (closed interval).
// It panics if either bound is negative or minLen is greater than maxLen.
//
// The following, additional template variables are supported:
//   - [govy.TemplateVars.MinLength]
//   - [govy.TemplateVars.MaxLength]
func MapLength[M ~map[K]V, K comparable, V any](minLen, maxLen int) govy.Rule[M] {
	schemaMinLen, schemaMaxLen := enforceMinMaxLength(minLen, maxLen)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			MinLength: minLen,
			MaxLength: maxLen,
		}).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{
				MinProperties: ptr(schemaMinLen),
				MaxProperties: ptr(schemaMaxLen),
			}, nil
		})
}

// MapMinLength ensures the map's length is greater than or equal to the limit.
// It panics if limit is negative.
func MapMinLength[M ~map[K]V, K comparable, V any](limit int) govy.Rule[M] {
	schemaLimit := enforceLength(limit)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{MinProperties: ptr(schemaLimit)}, nil
		})
}

// MapMaxLength ensures the map's length is less than or equal to the limit.
// It panics if limit is negative.
func MapMaxLength[M ~map[K]V, K comparable, V any](limit int) govy.Rule[M] {
	schemaLimit := enforceLength(limit)
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
		WithDescriptionTemplate(tpl, govy.TemplateVars{
			ComparisonValue: limit,
		}).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{MaxProperties: ptr(schemaLimit)}, nil
		})
}

func enforceMinMaxLength(minLen, maxLen int) (minLength, maxLength uint64) {
	if minLen > maxLen {
		panic(fmt.Sprintf("minLen '%d' is greater than maxLen '%d'", minLen, maxLen))
	}
	return enforceLength(minLen), enforceLength(maxLen)
}

func enforceLength(limit int) uint64 {
	if limit < 0 {
		panic(fmt.Sprintf("length limit '%d' is less than 0", limit))
	}
	return uint64(limit)
}

func ptr[T any](v T) *T { return &v }
