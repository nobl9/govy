package govy_test

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
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
				govy.ForMap(govy.GetSelf[map[string]any]()).RulesForValues(rules.NEQ[any](1)),
			),
			fixture: "mixed",
			cases: []jsonschematest.Case[map[string]any]{
				{Name: "named string remains valid", Input: map[string]any{"name": "ok", "extra": 2}, Valid: true},
				{Name: "wildcard rejects integer", Input: map[string]any{"name": "ok", "extra": 1}},
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

func TestJSONSchema_CustomApplicatorsWithExplicitPaths(t *testing.T) {
	t.Parallel()
	t.Run("named property", func(t *testing.T) {
		t.Parallel()
		rule := govy.NewRule(func(value map[string]int) error {
			for _, item := range value {
				if item < 2 {
					return fmt.Errorf("all values must be at least 2")
				}
			}
			return nil
		}).WithDescription("minimum for all values").
			WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				return &jsonschema.Schema{AdditionalProperties: &jsonschema.Schema{Minimum: "2"}}, nil
			})
		v := govy.New(
			govy.For(govy.GetSelf[map[string]int]()).Rules(rule),
			govy.For(func(value map[string]int) int { return value["name"] }).WithName("name").Rules(rules.GTE(1)),
		)
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		cases := []jsonschematest.Case[map[string]int]{
			{Name: "both constraints pass", Input: map[string]int{"name": 2, "extra": 2}, Valid: true},
			{Name: "custom rule rejects named value", Input: map[string]int{"name": 1, "extra": 2}},
			{Name: "custom rule rejects extra value", Input: map[string]int{"name": 2, "extra": 1}},
		}
		for _, tc := range cases {
			assert.Equal(t, tc.Valid, v.Validate(tc.Input) == nil)
		}
		jsonschematest.Assert(t, schema, "test_data/expected_custom_applicator_named.json", cases)
	})
	t.Run("indexed item", func(t *testing.T) {
		t.Parallel()
		rule := govy.NewRule(func(value []int) error {
			for _, item := range value {
				if item < 2 {
					return fmt.Errorf("all items must be at least 2")
				}
			}
			return nil
		}).WithDescription("minimum for all items").
			WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				return &jsonschema.Schema{Items: &jsonschema.Schema{Minimum: "2"}}, nil
			})
		v := govy.New(
			govy.For(govy.GetSelf[[]int]()).Rules(rule),
			govy.For(func(value []int) int { return value[0] }).WithPath(jsonpath.New().Index(0)).Rules(rules.GTE(1)),
		)
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		cases := []jsonschematest.Case[[]int]{
			{Name: "both constraints pass", Input: []int{2, 2}, Valid: true},
			{Name: "custom rule rejects indexed value", Input: []int{1, 2}},
			{Name: "custom rule rejects remaining value", Input: []int{2, 1}},
		}
		for _, tc := range cases {
			assert.Equal(t, tc.Valid, v.Validate(tc.Input) == nil)
		}
		jsonschematest.Assert(t, schema, "test_data/expected_custom_applicator_indexed.json", cases)
	})
}
