package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
)

type lengthTestCase[T any] struct {
	value                T
	minLen               int
	maxLen               int
	expectedError        string
	jsonSchemaDifference string
}

var stringLengthTestCases = []*lengthTestCase[string]{
	{value: "test", minLen: 4, maxLen: 4},
	{value: "test", minLen: 4, maxLen: 6},
	{value: "test", minLen: 2, maxLen: 4},
	{value: "test", minLen: 2, maxLen: 6},
	{value: "test", minLen: 5, maxLen: 6, expectedError: "length must be between 5 and 6"},
	{value: "test", minLen: 1, maxLen: 3, expectedError: "length must be between 1 and 3"},
	{value: "", minLen: 0, maxLen: 0},
	{value: "a", minLen: 0, maxLen: 0, expectedError: "length must be between 0 and 0"},
	{value: "€🤖e\u0301", minLen: 4, maxLen: 4},
}

func TestStringLength(t *testing.T) {
	for _, tc := range stringLengthTestCases {
		err := StringLength(tc.minLen, tc.maxLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringLength))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("panic if minLen is greater than maxLen", func(t *testing.T) {
		assert.Panic(t,
			func() { StringLength(10, 5) },
			"minLen '10' is greater than maxLen '5'")
	})
}

func TestStringLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "string_length", StringLength, stringLengthTestCases)
}

func BenchmarkStringLength(b *testing.B) {
	for _, tc := range stringLengthTestCases {
		rule := StringLength(tc.minLen, tc.maxLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var stringMinLengthTestCases = []*lengthTestCase[string]{
	{value: "test", minLen: 0},
	{value: "test", minLen: 4},
	{value: "test", minLen: 5, expectedError: "length must be greater than or equal to 5"},
	{value: "test", minLen: 10, expectedError: "length must be greater than or equal to 10"},
	{value: "", minLen: 0},
	{value: "€🤖e\u0301", minLen: 4},
	{value: "🤖ab", minLen: 4, expectedError: "length must be greater than or equal to 4"},
}

func TestStringMinLength(t *testing.T) {
	for _, tc := range stringMinLengthTestCases {
		err := StringMinLength(tc.minLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMinLength))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringMinLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "string_min_length", func(minLen, _ int) govy.Rule[string] {
		return StringMinLength(minLen)
	}, stringMinLengthTestCases)
}

func BenchmarkStringMinLength(b *testing.B) {
	for _, tc := range stringMinLengthTestCases {
		rule := StringMinLength(tc.minLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var stringMaxLengthTestCases = []*lengthTestCase[string]{
	{value: "test", maxLen: 10},
	{value: "test", maxLen: 4},
	{value: "test", maxLen: 3, expectedError: "length must be less than or equal to 3"},
	{value: "test", maxLen: 0, expectedError: "length must be less than or equal to 0"},
	{value: "", maxLen: 0},
	{value: "€🤖e\u0301", maxLen: 4},
	{value: "🤖ab", maxLen: 3},
}

func TestStringMaxLength(t *testing.T) {
	for _, tc := range stringMaxLengthTestCases {
		err := StringMaxLength(tc.maxLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMaxLength))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringMaxLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "string_max_length", func(_, maxLen int) govy.Rule[string] {
		return StringMaxLength(maxLen)
	}, stringMaxLengthTestCases)
}

func BenchmarkStringMaxLength(b *testing.B) {
	for _, tc := range stringMaxLengthTestCases {
		rule := StringMaxLength(tc.maxLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var sliceLengthTestCases = []*lengthTestCase[[]string]{
	{value: []string{"a", "b", "c"}, minLen: 3, maxLen: 3},
	{value: []string{"a", "b", "c"}, minLen: 1, maxLen: 4},
	{value: []string{"a", "b", "c"}, minLen: 3, maxLen: 10},
	{value: []string{"a", "b", "c"}, minLen: 0, maxLen: 3},
	{value: []string{"a", "b", "c"}, minLen: 4, maxLen: 10, expectedError: "length must be between 4 and 10"},
	{value: []string{"a", "b", "c"}, minLen: 1, maxLen: 2, expectedError: "length must be between 1 and 2"},
	{value: []string{}, minLen: 0, maxLen: 0},
	{value: []string{"a"}, minLen: 0, maxLen: 0, expectedError: "length must be between 0 and 0"},
	{
		value: nil, minLen: 0, maxLen: 0,
		jsonSchemaDifference: "A nil slice marshals as null, but the generated schema requires an array.",
	},
	{value: nil, minLen: 1, maxLen: 2, expectedError: "length must be between 1 and 2"},
}

func TestSliceLength(t *testing.T) {
	for _, tc := range sliceLengthTestCases {
		err := SliceLength[[]string](tc.minLen, tc.maxLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceLength))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("panic if minLen is greater than maxLen", func(t *testing.T) {
		assert.Panic(t,
			func() { SliceLength[[]string](10, 5) },
			"minLen '10' is greater than maxLen '5'")
	})
}

func TestSliceLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "slice_length", SliceLength[[]string], sliceLengthTestCases)
}

func BenchmarkSliceLength(b *testing.B) {
	for _, tc := range sliceLengthTestCases {
		rule := SliceLength[[]string](tc.minLen, tc.maxLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var sliceMinLengthTestCases = []*lengthTestCase[[]string]{
	{value: []string{"a", "b", "c"}, minLen: 0},
	{value: []string{"a", "b", "c"}, minLen: 3},
	{value: []string{"a", "b", "c"}, minLen: 4, expectedError: "length must be greater than or equal to 4"},
	{value: []string{"a", "b", "c"}, minLen: 10, expectedError: "length must be greater than or equal to 10"},
	{value: []string{}, minLen: 0},
	{
		value: nil, minLen: 0,
		jsonSchemaDifference: "A nil slice marshals as null, but the generated schema requires an array.",
	},
	{value: nil, minLen: 3, expectedError: "length must be greater than or equal to 3"},
}

func TestSliceMinLength(t *testing.T) {
	for _, tc := range sliceMinLengthTestCases {
		err := SliceMinLength[[]string](tc.minLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceMinLength))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestSliceMinLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "slice_min_length", func(minLen, _ int) govy.Rule[[]string] {
		return SliceMinLength[[]string](minLen)
	}, sliceMinLengthTestCases)
}

func BenchmarkSliceMinLength(b *testing.B) {
	for _, tc := range sliceMinLengthTestCases {
		rule := SliceMinLength[[]string](tc.minLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var sliceMaxLengthTestCases = []*lengthTestCase[[]string]{
	{value: []string{"a", "b", "c"}, maxLen: 10},
	{value: []string{"a", "b", "c"}, maxLen: 3},
	{value: []string{"a", "b", "c"}, maxLen: 2, expectedError: "length must be less than or equal to 2"},
	{value: []string{"a", "b", "c"}, maxLen: 0, expectedError: "length must be less than or equal to 0"},
	{value: []string{}, maxLen: 0},
	{
		value: nil, maxLen: 0,
		jsonSchemaDifference: "A nil slice marshals as null, but the generated schema requires an array.",
	},
}

func TestSliceMaxLength(t *testing.T) {
	for _, tc := range sliceMaxLengthTestCases {
		err := SliceMaxLength[[]string](tc.maxLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceMaxLength))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestSliceMaxLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "slice_max_length", func(_, maxLen int) govy.Rule[[]string] {
		return SliceMaxLength[[]string](maxLen)
	}, sliceMaxLengthTestCases)
}

func BenchmarkSliceMaxLength(b *testing.B) {
	for _, tc := range sliceMaxLengthTestCases {
		rule := SliceMaxLength[[]string](tc.maxLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var mapLengthTestCases = []*lengthTestCase[map[string]string]{
	{value: map[string]string{"a": "b", "c": "d"}, minLen: 0, maxLen: 2},
	{value: map[string]string{"a": "b", "c": "d"}, minLen: 1, maxLen: 3},
	{
		value:  map[string]string{"a": "b", "c": "d"},
		minLen: 2,
		maxLen: 2,
	},
	{
		value:         map[string]string{"a": "b", "c": "d"},
		minLen:        3,
		maxLen:        4,
		expectedError: "length must be between 3 and 4",
	},
	{
		value:         map[string]string{"a": "b", "c": "d"},
		minLen:        1,
		maxLen:        1,
		expectedError: "length must be between 1 and 1",
	},
	{value: map[string]string{}, minLen: 0, maxLen: 0},
	{value: map[string]string{"a": "b"}, minLen: 0, maxLen: 0, expectedError: "length must be between 0 and 0"},
	{
		value: nil, minLen: 0, maxLen: 0,
		jsonSchemaDifference: "A nil map marshals as null, but the generated schema requires an object.",
	},
	{value: nil, minLen: 1, maxLen: 1, expectedError: "length must be between 1 and 1"},
}

func TestMapLength(t *testing.T) {
	for _, tc := range mapLengthTestCases {
		err := MapLength[map[string]string](tc.minLen, tc.maxLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMapLength))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("panic if minLen is greater than maxLen", func(t *testing.T) {
		assert.Panic(t,
			func() { MapLength[map[string]string](10, 5) },
			"minLen '10' is greater than maxLen '5'")
	})
}

func TestMapLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "map_length", MapLength[map[string]string], mapLengthTestCases)
}

func BenchmarkMapLength(b *testing.B) {
	for _, tc := range mapLengthTestCases {
		rule := MapLength[map[string]string](tc.minLen, tc.maxLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var mapMinLengthTestCases = []*lengthTestCase[map[string]string]{
	{value: map[string]string{"a": "b", "c": "d"}, minLen: 0},
	{value: map[string]string{"a": "b", "c": "d"}, minLen: 2},
	{
		value:         map[string]string{"a": "b", "c": "d"},
		minLen:        3,
		expectedError: "length must be greater than or equal to 3",
	},
	{
		value:         map[string]string{"a": "b", "c": "d"},
		minLen:        10,
		expectedError: "length must be greater than or equal to 10",
	},
	{value: map[string]string{}, minLen: 0},
	{
		value: nil, minLen: 0,
		jsonSchemaDifference: "A nil map marshals as null, but the generated schema requires an object.",
	},
	{value: nil, minLen: 2, expectedError: "length must be greater than or equal to 2"},
}

func TestMapMinLength(t *testing.T) {
	for _, tc := range mapMinLengthTestCases {
		err := MapMinLength[map[string]string](tc.minLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMapMinLength))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestMapMinLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "map_min_length", func(minLen, _ int) govy.Rule[map[string]string] {
		return MapMinLength[map[string]string](minLen)
	}, mapMinLengthTestCases)
}

func BenchmarkMapMinLength(b *testing.B) {
	for _, tc := range mapMinLengthTestCases {
		rule := MapMinLength[map[string]string](tc.minLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var mapMaxLengthTestCases = []*lengthTestCase[map[string]string]{
	{value: map[string]string{"a": "b", "c": "d"}, maxLen: 10},
	{value: map[string]string{"a": "b", "c": "d"}, maxLen: 2},
	{
		value:         map[string]string{"a": "b", "c": "d"},
		maxLen:        1,
		expectedError: "length must be less than or equal to 1",
	},
	{
		value:         map[string]string{"a": "b", "c": "d"},
		maxLen:        0,
		expectedError: "length must be less than or equal to 0",
	},
	{value: map[string]string{}, maxLen: 0},
	{
		value: nil, maxLen: 0,
		jsonSchemaDifference: "A nil map marshals as null, but the generated schema requires an object.",
	},
}

func TestMapMaxLength(t *testing.T) {
	for _, tc := range mapMaxLengthTestCases {
		err := MapMaxLength[map[string]string](tc.maxLen).Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMapMaxLength))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestMapMaxLength_JSONSchema(t *testing.T) {
	t.Parallel()
	assertLengthJSONSchema(t, "map_max_length", func(_, maxLen int) govy.Rule[map[string]string] {
		return MapMaxLength[map[string]string](maxLen)
	}, mapMaxLengthTestCases)
}

func BenchmarkMapMaxLength(b *testing.B) {
	for _, tc := range mapMaxLengthTestCases {
		rule := MapMaxLength[map[string]string](tc.maxLen)
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

func TestLengthRules_NegativeLimits(t *testing.T) {
	tests := []struct {
		name    string
		newRule func()
	}{
		{name: "StringLength", newRule: func() { StringLength(-1, 0) }},
		{name: "StringMinLength", newRule: func() { StringMinLength(-1) }},
		{name: "StringMaxLength", newRule: func() { StringMaxLength(-1) }},
		{name: "SliceLength", newRule: func() { SliceLength[[]string](-1, 0) }},
		{name: "SliceMinLength", newRule: func() { SliceMinLength[[]string](-1) }},
		{name: "SliceMaxLength", newRule: func() { SliceMaxLength[[]string](-1) }},
		{name: "MapLength", newRule: func() { MapLength[map[string]string](-1, 0) }},
		{name: "MapMinLength", newRule: func() { MapMinLength[map[string]string](-1) }},
		{name: "MapMaxLength", newRule: func() { MapMaxLength[map[string]string](-1) }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Panic(t, tc.newRule, "length limit '-1' is less than 0")
		})
	}
}

func assertLengthJSONSchema[T any](
	t *testing.T,
	name string,
	newRule func(int, int) govy.Rule[T],
	tests []*lengthTestCase[T],
) {
	t.Helper()

	groups := make(map[[2]int][]jsonschematest.Case[T])
	for i, tc := range tests {
		limits := [2]int{tc.minLen, tc.maxLen}
		groups[limits] = append(groups[limits], jsonschematest.Case[T]{
			Name:                 fmt.Sprintf("%d/%v", i, tc.value),
			Input:                tc.value,
			Valid:                tc.expectedError == "",
			JSONSchemaDifference: tc.jsonSchemaDifference,
		})
	}
	for limits, cases := range groups {
		t.Run(fmt.Sprintf("%d_%d", limits[0], limits[1]), func(t *testing.T) {
			t.Parallel()
			rule := newRule(limits[0], limits[1])
			schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[T]()).Rules(rule)))
			assert.Require(t, assert.NoError(t, err))
			fixture := fmt.Sprintf("testdata/jsonschema/expected_%s_%d_%d.json", name, limits[0], limits[1])
			jsonschematest.Assert(t, schema, fixture, cases)
		})
	}
}
