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

func TestJSONSchema_ByteSliceLengths(t *testing.T) {
	t.Parallel()
	for name, rule := range map[string]govy.Rule[[]byte]{
		"both limits": rules.SliceLength[[]byte](1, 2),
		"minimum":     rules.SliceMinLength[[]byte](1),
		"maximum":     rules.SliceMaxLength[[]byte](2),
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			v := govy.New(govy.For(govy.GetSelf[[]byte]()).Rules(rule))
			schema, err := govy.JSONSchema(v)
			assert.Require(t, assert.NoError(t, err))
			cases := make([]jsonschematest.Case[[]byte], 0, 4)
			for _, input := range [][]byte{{}, {1}, {1, 2}, {1, 2, 3}} {
				valid := v.Validate(input) == nil
				var difference string
				if !valid {
					difference = "Byte counts do not map to the length of the base64 JSON string."
				}
				cases = append(
					cases,
					jsonschematest.Case[[]byte]{Input: input, Valid: valid, JSONSchemaDifference: difference},
				)
			}
			jsonschematest.Assert(t, schema, "test_data/expected_encoded_bytes.json", cases)
		})
	}
	t.Run("named byte slice", func(t *testing.T) {
		t.Parallel()
		type bytes []byte
		v := govy.New(govy.For(govy.GetSelf[bytes]()).Rules(rules.SliceMinLength[bytes](1)))
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		jsonschematest.Assert(t, schema, "test_data/expected_encoded_bytes.json", []jsonschematest.Case[bytes]{
			{Name: "base64", Input: bytes{1}, Valid: true},
		})
	})
	t.Run("byte elements with JSON methods", func(t *testing.T) {
		t.Parallel()
		v := govy.New(govy.For(govy.GetSelf[[]schemaJSONByte]()).Rules(rules.SliceLength[[]schemaJSONByte](1, 2)))
		schema, err := govy.JSONSchema(v)
		assert.Require(t, assert.NoError(t, err))
		jsonschematest.Assert(
			t,
			schema,
			"test_data/expected_encoded_byte_elements.json",
			[]jsonschematest.Case[[]schemaJSONByte]{
				{Name: "one element", Input: []schemaJSONByte{1}, Valid: true},
				{Name: "too many elements", Input: []schemaJSONByte{1, 2, 3}},
			},
		)
	})
}

type schemaJSONValue struct{ Value int }

func (v schemaJSONValue) MarshalJSON() ([]byte, error) { return json.Marshal(v.Value) }

type schemaTextValue struct{ Value string }

func (v schemaTextValue) MarshalText() ([]byte, error) { return []byte(v.Value), nil }

type schemaPointerTextValue struct{ Value string }

func (v *schemaPointerTextValue) MarshalText() ([]byte, error) { return []byte(v.Value), nil }

type schemaJSONByte byte

func (v schemaJSONByte) MarshalJSON() ([]byte, error) { return json.Marshal(byte(v)) }
