package rules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/nobl9/govy/internal/ecmaregex"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/jsonschema"
)

func addRequiredJSONSchemaProperty(ctx govy.JSONSchemaBuilderContext) {
	if ctx.Parent == nil || ctx.Segment.Kind() != jsonpath.SegmentName {
		return
	}
	name := ctx.Segment.Name()
	if !slices.Contains(ctx.Parent.Required, name) {
		ctx.Parent.Required = append(ctx.Parent.Required, name)
	}
}

func jsonSchemaValue(value any) (any, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return nil, fmt.Errorf("marshal value: %w", err)
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	var result any
	if err = decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("decode marshaled value: %w", err)
	}
	return result, nil
}

func jsonSchemaValues[T any](values []T) ([]any, error) {
	result := make([]any, 0, len(values))
	for i, value := range values {
		converted, err := jsonSchemaValue(value)
		if err != nil {
			return nil, fmt.Errorf("convert value at index %d: %w", i, err)
		}
		duplicate := false
		for _, existing := range result {
			if reflect.DeepEqual(existing, converted) {
				duplicate = true
				break
			}
		}
		if !duplicate {
			result = append(result, converted)
		}
	}
	return result, nil
}

func addJSONSchemaConst(schema *jsonschema.Schema, value any) {
	if schema.Const == nil {
		schema.Const = ptr(value)
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{Const: ptr(value)})
}

func addJSONSchemaEnum(schema *jsonschema.Schema, values []any) {
	if len(schema.Enum) == 0 {
		schema.Enum = values
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{Enum: values})
}

func addJSONSchemaNot(schema, denied *jsonschema.Schema) {
	if schema.Not == nil {
		schema.Not = denied
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{Not: denied})
}

func addJSONSchemaAnyOf(schema *jsonschema.Schema, alternatives []*jsonschema.Schema) {
	if len(schema.AnyOf) == 0 {
		schema.AnyOf = alternatives
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{AnyOf: alternatives})
}

func addJSONSchemaOneOf(schema *jsonschema.Schema, alternatives []*jsonschema.Schema) {
	if len(schema.OneOf) == 0 {
		schema.OneOf = alternatives
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{OneOf: alternatives})
}

func jsonSchemaRequiredAlternatives(names []string) []*jsonschema.Schema {
	alternatives := make([]*jsonschema.Schema, 0, len(names))
	for _, name := range names {
		alternatives = append(alternatives, &jsonschema.Schema{Required: []string{name}})
	}
	return alternatives
}

func addJSONSchemaPattern(schema *jsonschema.Schema, pattern string) {
	if schema.Pattern == "" {
		schema.Pattern = pattern
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{Pattern: pattern})
}

func addJSONSchemaFormat(schema *jsonschema.Schema, format jsonschema.Format) {
	if schema.Format == "" {
		schema.Format = format
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{Format: format})
}

func jsonSchemaFormat(format jsonschema.Format) govy.JSONSchemaBuilder {
	return func(ctx govy.JSONSchemaBuilderContext) error {
		addJSONSchemaFormat(ctx.Schema, format)
		return nil
	}
}

func jsonSchemaPattern(pattern string) govy.JSONSchemaBuilder {
	return func(ctx govy.JSONSchemaBuilderContext) error {
		translated, err := ecmaregex.Translate(pattern)
		if err != nil {
			return err
		}
		addJSONSchemaPattern(ctx.Schema, translated)
		return nil
	}
}

func jsonSchemaDeniedPattern(pattern string) govy.JSONSchemaBuilder {
	return func(ctx govy.JSONSchemaBuilderContext) error {
		translated, err := ecmaregex.Translate(pattern)
		if err != nil {
			return err
		}
		addJSONSchemaNot(ctx.Schema, &jsonschema.Schema{Pattern: translated})
		return nil
	}
}

func addJSONSchemaMinimum(schema *jsonschema.Schema, number json.Number) {
	if schema.Minimum == "" {
		schema.Minimum = number
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{Minimum: number})
}

func addJSONSchemaMultipleOf(schema *jsonschema.Schema, number json.Number) {
	if schema.MultipleOf == "" {
		schema.MultipleOf = number
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{MultipleOf: number})
}

func addJSONSchemaMaximum(schema *jsonschema.Schema, number json.Number) {
	if schema.Maximum == "" {
		schema.Maximum = number
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{Maximum: number})
}

func addJSONSchemaExclusiveMinimum(schema *jsonschema.Schema, number json.Number) {
	if schema.ExclusiveMinimum == "" {
		schema.ExclusiveMinimum = number
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{ExclusiveMinimum: number})
}

func addJSONSchemaExclusiveMaximum(schema *jsonschema.Schema, number json.Number) {
	if schema.ExclusiveMaximum == "" {
		schema.ExclusiveMaximum = number
		return
	}
	schema.AllOf = append(schema.AllOf, &jsonschema.Schema{ExclusiveMaximum: number})
}

func jsonSchemaAffixPattern(values []string, prefix, suffix string) string {
	escaped := make([]string, 0, len(values))
	for _, value := range values {
		escaped = append(escaped, regexp.QuoteMeta(value))
	}
	return prefix + "(?:" + strings.Join(escaped, "|") + ")" + suffix
}
