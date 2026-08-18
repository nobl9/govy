package govy_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"testing"
	"unsafe"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonpath"
)

func TestValidatorPlan_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(newPodValidator())
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_pod_json_schema.json")
	actual, err := json.MarshalIndent(schema, "", "  ")
	assert.Require(t, assert.NoError(t, err))

	assert.Equal(t, expected, string(actual))
}

func TestJSONSchema_UnsupportedType(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		run           func() error
		expectedError string
	}{
		"channel": {
			run:           runJSONSchema[chan struct{}],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "chan"`,
		},
		"complex64": {
			run:           runJSONSchema[complex64],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "complex64"`,
		},
		"complex128": {
			run:           runJSONSchema[complex128],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "complex128"`,
		},
		"function": {
			run:           runJSONSchema[func()],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "func"`,
		},
		"unsafe pointer": {
			run:           runJSONSchema[unsafe.Pointer],
			expectedError: `failed to generate JSON Schema: unsupported Go kind "unsafe.Pointer"`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert.EqualError(t, tc.run(), tc.expectedError)
		})
	}
}

func TestJSONSchema_MapItemUsesValueType(t *testing.T) {
	t.Parallel()

	type annotations map[string]string
	type document struct{}
	validator := govy.New(
		govy.ForMap(func(document) annotations { return nil }).
			WithName("annotations").
			RulesForItems(
				govy.NewRule(func(govy.MapItem[string, string]) error { return nil }).
					WithDescription("key and value must differ"),
			),
	)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_map_item_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_FixedArrayIndexes(t *testing.T) {
	t.Parallel()

	type document struct{}
	validator := govy.New(
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("tuple").Index(2).Name("a")),
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("tuple").Index(2).Name("b")),
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("tuple").Index(10).Name("c")),
	)

	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))

	expected := readTestData(t, "expected_fixed_array_indexes_json_schema.json")
	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(schema)))
	assert.Equal(t, expected, actual.String())
}

func TestJSONSchema_ArrayIndexOutsideIntRange(t *testing.T) {
	t.Parallel()

	type document struct{}
	index := uint(math.MaxUint)
	validator := govy.New(
		govy.For(func(document) string { return "" }).
			WithPath(jsonpath.New().Name("tuple").Index(index)),
	)

	_, err := govy.JSONSchema(validator)
	assert.EqualError(t, err, fmt.Sprintf(
		`failed to generate JSON Schema for %q property: array index %d exceeds maximum supported index %d`,
		jsonpath.NewRoot().Name("tuple").Index(index),
		index,
		math.MaxInt-1,
	))
}

func runJSONSchema[T any]() error {
	_, err := govy.JSONSchema(govy.New[T]())
	return err
}
