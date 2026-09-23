package govy_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_EncodedTypes(t *testing.T) {
	t.Parallel()
	t.Run("number", func(t *testing.T) {
		t.Parallel()
		assertEncodedSchema(t, json.Number("1"), json.Number("2"), "number")
	})
	t.Run("number pointer", func(t *testing.T) {
		t.Parallel()
		one, two := json.Number("1"), json.Number("2")
		assertEncodedSchema(t, &one, &two, "number")
	})
	t.Run("time", func(t *testing.T) {
		t.Parallel()
		value := time.Date(2026, time.September, 23, 12, 0, 0, 0, time.UTC)
		assertEncodedSchema(t, value, value.Add(time.Hour), "time")
	})
	t.Run("custom JSON marshaler", func(t *testing.T) {
		t.Parallel()
		assertEncodedSchema(t, schemaJSONValue{Value: 1}, schemaJSONValue{Value: 2}, "custom_json")
	})
	t.Run("text marshaler", func(t *testing.T) {
		t.Parallel()
		assertEncodedSchema(t, schemaTextValue{Value: "a"}, schemaTextValue{Value: "b"}, "text")
	})
	t.Run("pointer text marshaler", func(t *testing.T) {
		t.Parallel()
		assertEncodedSchema(t, &schemaPointerTextValue{Value: "a"}, &schemaPointerTextValue{Value: "b"}, "text")
	})
	t.Run("value with pointer text method", func(t *testing.T) {
		t.Parallel()
		v := govy.New(govy.For(govy.GetSelf[schemaPointerTextValue]()).Rules(rules.Required[schemaPointerTextValue]()))
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		value := schemaPointerTextValue{Value: "a"}
		assert.NoError(t, v.Validate(value))
		jsonschematest.Assert(t, schema, "test_data/expected_encoded_unknown.json", []jsonschematest.Case[any]{
			{Name: "object value", Input: value, Valid: true},
			{Name: "addressable string", Input: &value, Valid: true},
		})
	})
	t.Run("pointer JSON method precedes value text method", func(t *testing.T) {
		t.Parallel()
		v := govy.New(govy.For(govy.GetSelf[schemaMixedEncoding]()).Rules(rules.Required[schemaMixedEncoding]()))
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		value := schemaMixedEncoding{Value: "a"}
		assert.NoError(t, v.Validate(value))
		jsonschematest.Assert(t, schema, "test_data/expected_encoded_unknown.json", []jsonschematest.Case[any]{
			{Name: "text for value", Input: value, Valid: true},
			{Name: "JSON object for pointer", Input: &value, Valid: true},
		})
	})
	t.Run("number map keys", func(t *testing.T) {
		t.Parallel()
		v := govy.New(govy.ForMap(govy.GetSelf[map[json.Number]string]()).RulesForKeys(rules.Required[json.Number]()))
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		value := map[json.Number]string{"1": "ok"}
		assert.NoError(t, v.Validate(value))
		jsonschematest.Assert(
			t,
			schema,
			"test_data/expected_encoded_number_keys.json",
			[]jsonschematest.Case[map[json.Number]string]{
				{Name: "string property name", Input: value, Valid: true},
			},
		)
	})
}

func assertEncodedSchema[T comparable](t *testing.T, valid, invalid T, fixture string) {
	t.Helper()
	v := govy.New(govy.For(govy.GetSelf[T]()).Rules(rules.EQ(valid)))
	schema, err := govy.JSONSchema(v)
	assert.Require(t, assert.NoError(t, err))
	cases := []jsonschematest.Case[T]{
		{Name: "equal", Input: valid, Valid: true},
		{Name: "different", Input: invalid},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, v.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_encoded_"+fixture+".json", cases)
}

type schemaJSONValue struct{ Value int }

func (v schemaJSONValue) MarshalJSON() ([]byte, error) { return json.Marshal(v.Value) }

type schemaTextValue struct{ Value string }

func (v schemaTextValue) MarshalText() ([]byte, error) { return []byte(v.Value), nil }

type schemaPointerTextValue struct{ Value string }

func (v *schemaPointerTextValue) MarshalText() ([]byte, error) { return []byte(v.Value), nil }

type schemaMixedEncoding struct{ Value string }

func (v schemaMixedEncoding) MarshalText() ([]byte, error) { return []byte(v.Value), nil }

func (v *schemaMixedEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"value": v.Value})
}
