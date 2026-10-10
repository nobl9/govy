package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
)

var requiredTestCases = []*struct {
	in                   any
	shouldFail           bool
	jsonSchemaDifference string
}{
	{1, false, ""},
	{"s", false, ""},
	{0.1, false, ""},
	{[]int{}, false, ""},
	{map[string]int{}, false, ""},
	{nil, true, "JSON Schema requires the property to be present, not non-null."},
	{struct{}{}, true, "JSON Schema requires the property to be present, not a non-zero struct."},
	{"", true, "JSON Schema requires the property to be present, not a non-empty string."},
	{false, true, "JSON Schema requires the property to be present, not true."},
	{0, true, "JSON Schema requires the property to be present, not a non-zero number."},
	{0.0, true, "JSON Schema requires the property to be present, not a non-zero number."},
}

func TestRequired(t *testing.T) {
	for _, tc := range requiredTestCases {
		err := Required[any]().Validate(tc.in)
		if tc.shouldFail {
			assert.Require(t, assert.Error(t, err))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeRequired))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestRequired_JSONSchema(t *testing.T) {
	t.Parallel()
	validator := govy.New(
		govy.For(func(v map[string]any) any { return v["value"] }).
			WithName("value").
			Rules(Required[any]()),
	)
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	cases := make([]jsonschematest.Case[map[string]any], 0, len(requiredTestCases)+1)
	for i, tc := range requiredTestCases {
		cases = append(cases, jsonschematest.Case[map[string]any]{
			Name:                 fmt.Sprintf("present/%d/%T", i, tc.in),
			Input:                map[string]any{"value": tc.in},
			Valid:                !tc.shouldFail,
			JSONSchemaDifference: tc.jsonSchemaDifference,
		})
	}
	cases = append(cases, jsonschematest.Case[map[string]any]{
		Name:  "missing",
		Input: map[string]any{},
		Valid: false,
	})
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_required.json", cases)
}

func BenchmarkRequired(b *testing.B) {
	for _, tc := range requiredTestCases {
		rule := Required[any]()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}
