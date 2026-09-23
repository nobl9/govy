package rules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"regexp/syntax"
	"strings"

	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/internal/ecmaregex"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
)

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

func jsonSchemaRequiredAlternatives(names []string) []*jsonschema.Schema {
	alternatives := make([]*jsonschema.Schema, 0, len(names))
	for _, name := range names {
		alternatives = append(alternatives, &jsonschema.Schema{Required: []string{name}})
	}
	return alternatives
}

func jsonSchemaStringEnum(lookup func() map[string]struct{}) govy.JSONSchemaBuilder {
	return func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		codes := collections.SortedKeys(lookup())
		values := make([]any, len(codes))
		for i, code := range codes {
			values[i] = code
		}
		return &jsonschema.Schema{Enum: values}, nil
	}
}

func jsonSchemaFormat(format jsonschema.Format) govy.JSONSchemaBuilder {
	return func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return &jsonschema.Schema{Format: format}, nil
	}
}

func jsonSchemaPattern(pattern string) govy.JSONSchemaBuilder {
	return func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		translated, err := ecmaregex.Translate(pattern)
		if err != nil {
			return nil, fmt.Errorf("translate Go regular expression %q: %w", pattern, err)
		}
		return &jsonschema.Schema{Pattern: translated}, nil
	}
}

func jsonSchemaCompiledPattern(pattern string) govy.JSONSchemaBuilder {
	return func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		perl, err := syntax.Parse(pattern, syntax.Perl)
		if err != nil {
			return nil, fmt.Errorf("parse Go regular expression %q: %w", pattern, err)
		}
		// Regexp.String does not preserve the compilation mode. Only translate
		// patterns whose meaning is independent of that missing information.
		posix, err := syntax.Parse(pattern, syntax.POSIX)
		if err == nil && !perl.Equal(posix) {
			return nil, fmt.Errorf(
				"regular expression %q has different Perl and POSIX meanings; provide an explicit WithJSONSchema mapping",
				pattern,
			)
		}
		return jsonSchemaPattern(pattern)(ctx)
	}
}

func jsonSchemaAffixPattern(values []string, prefix, suffix string) string {
	escaped := make([]string, 0, len(values))
	for _, value := range values {
		escaped = append(escaped, regexp.QuoteMeta(value))
	}
	return prefix + "(?:" + strings.Join(escaped, "|") + ")" + suffix
}
