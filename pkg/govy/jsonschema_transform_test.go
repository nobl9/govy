package govy_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_Transform(t *testing.T) {
	t.Parallel()

	type document struct{ Enabled bool }
	getString := func(document) string { return "" }
	unexpectedBuilder := func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return nil, fmt.Errorf("transformed-value builder must not run")
	}
	enabled := any(true)
	validator := govy.New(
		govy.Transform(func(document) []string { return nil }, func(v []string) (int, error) { return len(v), nil }).
			WithName("array").Rules(rules.GT(2)),
		govy.Transform(getString, strconv.Atoi).WithName("conditional").Required().Rules(rules.EQ(12)).
			When(func(v document) bool { return v.Enabled },
				govy.WhenJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
					if !ctx.Path.Equal(jsonpath.NewRoot()) || ctx.Type != jsonschema.TypeObject {
						return nil, fmt.Errorf("condition must describe the original parent")
					}
					return &jsonschema.Schema{
						Properties: map[string]*jsonschema.Schema{"enabled": {Const: &enabled}},
						Required:   []string{"enabled"},
					}, nil
				})),
		govy.Transform(getString, func(v string) (string, error) { return strings.ToLower(v), nil }).
			WithName("normalized").Rules(rules.EQ("admin").WithJSONSchema(unexpectedBuilder)),
		govy.Transform(getString, strconv.Atoi).WithName("number").
			Rules(govy.NewRuleSet(rules.EQ(12), rules.GTE(10))),
		govy.For(func(document) int { return 0 }).WithName("plain").Rules(rules.EQ(12)),
		govy.Transform(getString, strconv.Atoi).WithName("requiredProperty").Required().Rules(rules.EQ(12)),
		govy.Transform(getString, strconv.Atoi).WithName("requiredRule").
			Rules(rules.Required[int]().WithJSONSchema(unexpectedBuilder), rules.EQ(12)),
	)
	tests := []struct {
		name    string
		options []govy.JSONSchemaOption
		fixture string
	}{
		{name: "default", fixture: "expected_transform_json_schema.json"},
		{
			name:    "include omitted rules",
			options: []govy.JSONSchemaOption{govy.JSONSchemaIncludeOmittedRules()},
			fixture: "expected_transform_omitted_rules_json_schema.json",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			schema, err := govy.JSONSchema(validator, tc.options...)
			assert.Require(t, assert.NoError(t, err))

			expected := readTestData(t, tc.fixture)
			var actual bytes.Buffer
			encoder := json.NewEncoder(&actual)
			encoder.SetIndent("", "  ")
			assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
			assert.Equal(t, expected, actual.String())
		})
	}
}

func TestJSONSchema_TransformIncludedValidators(t *testing.T) {
	t.Parallel()

	type decoded struct {
		Name        string
		Nested      string
		Unsupported chan int
	}
	type document struct{}
	unexpectedBuilder := func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return nil, fmt.Errorf("builder for an included transformed value must not run")
	}
	objectValidator := govy.New(
		govy.For(govy.GetSelf[decoded]()).Rules(
			govy.NewRule(func(decoded) error { return nil }).WithErrorCode("decoded").WithJSONSchema(unexpectedBuilder),
		),
		govy.For(func(v decoded) string { return v.Name }).WithName("name").Required().Rules(rules.EQ("admin")),
		govy.Transform(func(v decoded) string { return v.Nested }, strconv.Atoi).
			WithName("nested").Required().Rules(rules.EQ(12)).
			When(func(decoded) bool { return true }, govy.WhenJSONSchema(unexpectedBuilder)),
		govy.For(func(v decoded) chan int { return v.Unsupported }).WithName("unsupported"),
	)
	validator := govy.New(
		govy.Transform(func(document) string { return "" }, func(string) (decoded, error) {
			t.Fatal("schema generation must not execute a transform")
			return decoded{}, nil
		}).WithName("object").Include(objectValidator),
		govy.Transform(func(document) string { return "" }, func(string) ([]int, error) {
			return nil, nil
		}).WithName("slice").Include(govy.New(
			govy.ForSlice(govy.GetSelf[[]int]()).Rules(rules.SliceMinLength[[]int](1)).RulesForEach(rules.GT(0)),
		)),
		govy.Transform(func(document) string { return "" }, func(string) (map[string]int, error) {
			return nil, nil
		}).WithName("map").Include(govy.New(
			govy.ForMap(govy.GetSelf[map[string]int]()).
				RulesForKeys(rules.StringMinLength(1)).
				RulesForValues(rules.GT(0)),
		)),
	)

	schema, err := govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_transform_included_validators_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_TransformInputMapping(t *testing.T) {
	t.Parallel()

	transformed := govy.Transform(govy.GetSelf[string](), strconv.Atoi).Rules(rules.EQ(12))
	input := govy.For(govy.GetSelf[string]()).Rules(
		govy.NewRule(func(string) error { return nil }).WithDescription("describe the original input").
			WithJSONSchema(func(ctx govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
				if !ctx.Path.Equal(jsonpath.NewRoot()) || ctx.Type != jsonschema.TypeString {
					return nil, fmt.Errorf("builder must describe the original string")
				}
				minimum := uint64(1)
				return &jsonschema.Schema{MinLength: &minimum}, nil
			}),
	)
	for name, validator := range map[string]govy.Validator[string]{
		"input mapping first": govy.New(input, transformed),
		"transform first":     govy.New(transformed, input),
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			schema, err := govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
			assert.Require(t, assert.NoError(t, err))

			expected := readTestData(t, "expected_transform_input_mapping_json_schema.json")
			var actual bytes.Buffer
			encoder := json.NewEncoder(&actual)
			encoder.SetIndent("", "  ")
			assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
			assert.Equal(t, expected, actual.String())

			assert.NoError(t, validator.Validate("12"))
			assert.True(t, govy.HasErrorCode(validator.Validate("13"), rules.ErrorCodeEqualTo))
			assert.True(t, govy.HasErrorCode(validator.Validate("invalid"), govy.ErrorCodeTransform))
		})
	}
}

func TestJSONSchema_TransformUnsupportedInput(t *testing.T) {
	t.Parallel()

	validator := govy.New(
		govy.Transform(func(struct{}) complex64 { return 0 }, func(complex64) (string, error) { return "", nil }).
			WithName("value").Rules(rules.EQ("valid")),
	)
	_, err := govy.JSONSchema(validator)
	assert.EqualError(
		t,
		err,
		`failed to generate JSON Schema type info for "$.value" property: unsupported Go kind "complex64"`,
	)
}

func TestPlan_Transform(t *testing.T) {
	t.Parallel()

	type decoded struct{ Count int }
	validator := govy.New(
		govy.Transform(govy.GetSelf[string](), func(v string) (decoded, error) {
			count, err := strconv.Atoi(v)
			return decoded{Count: count}, err
		}).Required().Include(govy.New(
			govy.For(func(v decoded) int { return v.Count }).WithName("count").Required().Rules(rules.GT(0)),
		)),
	)
	expected := readTestData(t, "expected_transform_plan.json")
	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))
	assert.Equal(t, expected, requireJSON(t, plan))

	_, err = govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
	assert.Require(t, assert.NoError(t, err))

	plan, err = govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))
	assert.Equal(t, expected, requireJSON(t, plan))
}
