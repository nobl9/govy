package govy_test

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_RequiredErrorCodes(t *testing.T) {
	t.Parallel()

	type document struct {
		Name string `json:"name,omitempty"`
	}
	for name, rule := range map[string]govy.RulesInterface[string]{
		"original": rules.Required[string](),
		"replaced": rules.Required[string]().WithErrorCode("custom_required"),
		"grouped":  govy.NewRuleSet(rules.Required[string]()).WithErrorCode("group"),
		"custom presence marker": rules.StringMinLength(1).
			WithPlanModifiers(govy.RulePlanModifierRequired()).
			WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) { return nil, nil }),
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			validator := govy.New(govy.For(func(d document) string { return d.Name }).
				WithName("name").Rules(rule))
			schema, err := govy.JSONSchema(validator)
			assert.Require(t, assert.NoError(t, err))
			cases := []jsonschematest.Case[document]{
				{Name: "missing", Input: document{}},
				{Name: "present", Input: document{Name: "ok"}, Valid: true},
			}
			for _, tc := range cases {
				assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
			}
			jsonschematest.Assert(t, schema, "test_data/expected_required_codes_json_schema.json", cases)
		})
	}
	t.Run("unrelated rule reuses required code", func(t *testing.T) {
		t.Parallel()
		validator := govy.New(govy.For(func(d document) string { return d.Name }).
			WithName("name").Rules(rules.StringMinLength(0).WithErrorCode("required")))
		schema, err := govy.JSONSchema(validator)
		assert.Require(t, assert.NoError(t, err))
		assert.Len(t, schema.Required, 0)
		assert.NoError(t, validator.Validate(document{}))
	})
}
