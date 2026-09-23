package govy_test

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_CustomObjectApplicators(t *testing.T) {
	t.Parallel()

	contribution := func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return &jsonschema.Schema{
			Properties:           map[string]*jsonschema.Schema{"name": {Type: jsonschema.TypeString}},
			AdditionalProperties: &jsonschema.Schema{Type: jsonschema.TypeInteger},
		}, nil
	}
	matches := func(value map[string]any) bool {
		for key, item := range value {
			if key == "name" {
				if _, ok := item.(string); !ok {
					return false
				}
			} else if _, ok := item.(int); !ok {
				return false
			}
		}
		return true
	}
	rule := govy.NewRule(func(value map[string]any) error {
		if !matches(value) {
			return fmt.Errorf("invalid object")
		}
		return nil
	}).WithDescription("custom object").WithJSONSchema(contribution)
	base := govy.New(govy.For(govy.GetSelf[map[string]any]()).Rules(rule))
	for _, tt := range []struct {
		name      string
		validator govy.Validator[map[string]any]
		fixture   string
		cases     []jsonschematest.Case[map[string]any]
	}{
		{
			name: "custom contribution", validator: base, fixture: "object",
			cases: []jsonschematest.Case[map[string]any]{
				{Name: "named string and integer extra", Input: map[string]any{"name": "ok", "extra": 1}, Valid: true},
				{Name: "invalid extra", Input: map[string]any{"name": "ok", "extra": "wrong"}},
			},
		},
		{
			name: "custom condition",
			validator: govy.New(govy.For(govy.GetSelf[map[string]any]()).Rules(rules.MapMaxLength[map[string]any](2))).
				When(matches, govy.WhenJSONSchema(contribution)),
			fixture: "condition",
			cases: []jsonschematest.Case[map[string]any]{
				{Name: "condition matches", Input: map[string]any{"name": "ok", "a": 1}, Valid: true},
				{Name: "condition enables length check", Input: map[string]any{"name": "ok", "a": 1, "b": 2}},
				{Name: "condition disables length check", Input: map[string]any{"name": 3, "a": 1, "b": 2}, Valid: true},
			},
		},
		{
			name: "custom and generated wildcards",
			validator: govy.New(
				govy.For(govy.GetSelf[map[string]any]()).Rules(rule),
				govy.ForMap(govy.GetSelf[map[string]any]()).RulesForValues(rules.NEQ[any](nil)),
			),
			fixture: "mixed",
			cases: []jsonschematest.Case[map[string]any]{
				{Name: "named string remains valid", Input: map[string]any{"name": "ok", "extra": 1}, Valid: true},
				{Name: "invalid extra", Input: map[string]any{"name": "ok", "extra": "wrong"}},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			schema, err := govy.JSONSchema(tt.validator)
			assert.Require(t, assert.NoError(t, err))
			for _, tc := range tt.cases {
				assert.Equal(t, tc.Valid, tt.validator.Validate(tc.Input) == nil)
			}
			jsonschematest.Assert(t, schema, "test_data/expected_custom_applicator_"+tt.fixture+".json", tt.cases)
		})
	}
}

func TestJSONSchema_CustomArrayApplicators(t *testing.T) {
	t.Parallel()

	rule := govy.NewRule(func(value []any) error {
		for i, item := range value {
			if i == 0 {
				if _, ok := item.(string); !ok {
					return fmt.Errorf("invalid prefix item")
				}
			} else if _, ok := item.(int); !ok {
				return fmt.Errorf("invalid remaining item")
			}
		}
		return nil
	}).WithDescription("custom tuple").
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{
				PrefixItems: []*jsonschema.Schema{{Type: jsonschema.TypeString}},
				Items:       &jsonschema.Schema{Type: jsonschema.TypeInteger},
			}, nil
		})
	validator := govy.New(govy.For(govy.GetSelf[[]any]()).Rules(rule))
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	cases := []jsonschematest.Case[[]any]{
		{Name: "string prefix and integer tail", Input: []any{"ok", 1}, Valid: true},
		{Name: "invalid tail", Input: []any{"ok", "wrong"}},
		{Name: "invalid prefix", Input: []any{1, 2}},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_custom_applicator_array.json", cases)
}
