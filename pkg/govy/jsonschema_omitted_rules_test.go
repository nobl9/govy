package govy_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_OmittedRules(t *testing.T) {
	type item struct{ Value string }
	type document struct{}
	missing := govy.NewRule(func(string) error { return nil }).WithErrorCode("custom")
	mapped := missing.WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return &jsonschema.Schema{Pattern: "mapped"}, nil
	})
	itemValidator := govy.New(
		govy.For(func(v item) string { return v.Value }).WithName("value").Rules(missing),
	)
	validator := govy.New(
		govy.For(govy.GetSelf[document]()).Rules(
			govy.NewRule(func(document) error { return nil }).WithErrorCode("root"),
		),
		govy.For(func(document) string { return "" }).WithName("anonymous").Rules(
			govy.NewRule(func(string) error {
				t.Fatal("schema generation must not execute validation")
				return nil
			}),
		),
		govy.For(func(document) string { return "" }).WithName("timezone").Rules(rules.StringTimeZone()),
		govy.For(func(document) string { return "" }).
			WithName("mixed").
			Rules(govy.NewRuleSet(missing, missing, mapped)),
		govy.ForSlice(func(document) []item { return nil }).WithName("items").IncludeForEach(itemValidator),
		govy.ForMap(func(document) map[string]string { return nil }).WithName("labels").
			RulesForKeys(missing).RulesForValues(missing),
		govy.For(func(document) item { return item{} }).WithName("nested").Include(itemValidator),
	)
	includeOmitted := govy.JSONSchemaIncludeOmittedRules()
	tests := []struct {
		name    string
		options []govy.JSONSchemaOption
		fixture string
	}{
		{
			name:    "disabled by default",
			fixture: "expected_omitted_rules_disabled_json_schema.json",
		},
		{
			name:    "enabled",
			options: []govy.JSONSchemaOption{includeOmitted},
			fixture: "expected_omitted_rules_json_schema.json",
		},
		{
			name:    "reused option",
			options: []govy.JSONSchemaOption{includeOmitted},
			fixture: "expected_omitted_rules_json_schema.json",
		},
		{
			name:    "default after enabled",
			fixture: "expected_omitted_rules_disabled_json_schema.json",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
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

func TestJSONSchema_OmittedConditions(t *testing.T) {
	t.Parallel()

	type document struct{}
	predicate := func(document) bool {
		t.Fatal("schema generation must not execute predicates")
		return false
	}
	builder := func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		t.Fatal("schema generation must not execute builders guarded by an unmapped condition")
		return nil, nil
	}
	rule := govy.NewRule(func(string) error { return nil }).
		WithErrorCode("conditional").WithJSONSchema(builder)
	nested := govy.New(
		govy.For(func(document) string { return "" }).WithName("value").Rules(rule),
	).When(predicate)
	validator := govy.New(
		govy.For(func(document) string { return "" }).WithName("property").
			Required().Rules(rule).When(predicate),
		govy.For(govy.GetSelf[document]()).WithName("nested").OmitEmpty().Include(nested).
			When(predicate, govy.WhenJSONSchema(builder)),
		govy.For(func(document) string { return "" }).WithName("bothMissing").
			Rules(govy.NewRule(func(string) error { return nil }).WithErrorCode("unsupported")).When(predicate),
		govy.For(func(document) string { return "" }).WithName("unconditional").Rules(rules.StringMinLength(1)),
	)

	schema, err := govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_omitted_conditions_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_NoOmittedRules(t *testing.T) {
	t.Parallel()

	type document struct{}
	nilBuilder := func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) { return nil, nil }
	nested := govy.New(
		govy.For(func(document) string { return "" }).WithName("value").Rules(rules.StringMinLength(1)),
	)
	validator := govy.New(
		govy.For(func(document) string { return "" }).WithName("nilRule").Rules(
			govy.NewRule(func(string) error { return nil }).WithErrorCode("nil_rule").WithJSONSchema(nilBuilder),
		),
		govy.For(func(document) string { return "" }).WithName("nilCondition").Rules(
			govy.NewRule(func(string) error { return nil }).WithErrorCode("nil_condition").
				WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
					t.Fatal("a nil condition schema must skip its rule builder")
					return nil, nil
				}),
		).When(func(document) bool { return true }, govy.WhenJSONSchema(nilBuilder)),
		govy.ForPointer(func(document) *string { return nil }).WithName("pointer").Rules(rules.StringMinLength(1)),
		govy.ForPointer(func(document) *document { return nil }).WithName("nestedPointer").Include(nested),
		govy.For(govy.GetSelf[document]()).WithName("nestedEmpty").OmitEmpty().Include(nested),
		govy.For(func(document) string { return "" }).WithName("empty").OmitEmpty(),
		govy.For(func(document) string { return "" }).WithName("required").Required(),
	)

	schema, err := govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_no_omitted_rules_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}
