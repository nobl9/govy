package govy_test

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_RepeatedConditionTypes(t *testing.T) {
	t.Parallel()

	type document struct {
		Name string
	}
	tests := []struct {
		name  string
		path  string
		build func(govy.Validator[string]) (*jsonschema.Document, error)
	}{
		{name: "root", path: "$", build: func(v govy.Validator[string]) (*jsonschema.Document, error) {
			return govy.JSONSchema(v)
		}},
		{name: "named property", path: "$.name", build: func(v govy.Validator[string]) (*jsonschema.Document, error) {
			return govy.JSONSchema(govy.New(govy.For(func(d document) string { return d.Name }).
				WithName("name").Include(v)))
		}},
		{name: "array item", path: "$[*]", build: func(v govy.Validator[string]) (*jsonschema.Document, error) {
			return govy.JSONSchema(govy.New(govy.ForSlice(govy.GetSelf[[]string]()).IncludeForEach(v)))
		}},
		{name: "map value", path: "$.*", build: func(v govy.Validator[string]) (*jsonschema.Document, error) {
			return govy.JSONSchema(govy.New(govy.ForMap(govy.GetSelf[map[string]string]()).IncludeForValues(v)))
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var contexts []govy.JSONSchemaBuilderContext
			condition := govy.WhenJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				contexts = append(contexts, ctx)
				return &jsonschema.Schema{}, nil
			})
			v := govy.New(govy.For(govy.GetSelf[string]()).Rules(rules.StringMinLength(1))).
				When(func(string) bool { return true }, condition).
				When(func(string) bool { return true }, condition)
			_, err := tt.build(v)
			assert.Require(t, assert.NoError(t, err))
			assert.Require(t, assert.Len(t, contexts, 2))
			for _, ctx := range contexts {
				assert.Equal(t, jsonschema.TypeString, ctx.Type)
				assert.Equal(t, tt.path, ctx.Path.String())
			}
		})
	}
}
