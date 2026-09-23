package govy_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
)

func TestJSONSchema_UnlabelledBuilder(t *testing.T) {
	t.Parallel()

	rule := govy.NewRule(func(value int) error {
		if value < 3 {
			return fmt.Errorf("value must be at least 3")
		}
		return nil
	}).WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
		return &jsonschema.Schema{Minimum: "3"}, nil
	})
	validator := govy.New(govy.For(govy.GetSelf[int]()).Rules(rule))
	schema, err := govy.JSONSchema(validator, govy.JSONSchemaIncludeOmittedRules())
	assert.Require(t, assert.NoError(t, err))
	cases := []jsonschematest.Case[int]{
		{Name: "below minimum", Input: 2},
		{Name: "at minimum", Input: 3, Valid: true},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_unlabelled_builder_json_schema.json", cases)
}

func TestJSONSchema_UnlabelledBuilderError(t *testing.T) {
	t.Parallel()

	expected := errors.New("builder failed")
	rule := govy.NewRule(func(int) error { return nil }).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return nil, expected
		})
	_, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[int]()).Rules(rule)))
	assert.True(t, errors.Is(err, expected))
	assert.EqualError(t, err, `failed to build JSON Schema for "$" property and "" rule: builder failed`)
}
