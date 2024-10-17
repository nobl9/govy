package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var stringLengthTestCases = []*struct {
	value         string
	minLen        int
	maxLen        int
	expectedError string
}{
	{value: "test", minLen: 4, maxLen: 4},
	{value: "test", minLen: 4, maxLen: 6},
	{value: "test", minLen: 2, maxLen: 4},
	{value: "test", minLen: 2, maxLen: 6},
	{value: "test", minLen: 5, maxLen: 6, expectedError: "length must be between 5 and 6"},
	{value: "test", minLen: 1, maxLen: 3, expectedError: "length must be between 1 and 3"},
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

func BenchmarkStringLength(b *testing.B) {
	for range b.N {
		for _, tc := range stringLengthTestCases {
			_ = StringLength(tc.minLen, tc.maxLen).Validate(tc.value)
		}
	}
}

var stringMinLengthTestCases = []*struct {
	value         string
	minLen        int
	expectedError string
}{
	{value: "test", minLen: 0},
	{value: "test", minLen: 4},
	{value: "test", minLen: 5, expectedError: "length must be greater than or equal to 5"},
	{value: "test", minLen: 10, expectedError: "length must be greater than or equal to 10"},
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

func BenchmarkStringMinLength(b *testing.B) {
	for range b.N {
		for _, tc := range stringMinLengthTestCases {
			_ = StringMinLength(tc.minLen).Validate(tc.value)
		}
	}
}

var stringMaxLengthTestCases = []*struct {
	value         string
	maxLen        int
	expectedError string
}{
	{value: "test", maxLen: 10},
	{value: "test", maxLen: 4},
	{value: "test", maxLen: 3, expectedError: "length must be less than or equal to 3"},
	{value: "test", maxLen: 0, expectedError: "length must be less than or equal to 0"},
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

func BenchmarkStringMaxLength(b *testing.B) {
	for range b.N {
		for _, tc := range stringMaxLengthTestCases {
			_ = StringMaxLength(tc.maxLen).Validate(tc.value)
		}
	}
}

var sliceLengthTestCases = []*struct {
	value         []string
	minLen        int
	maxLen        int
	expectedError string
}{
	{value: []string{"a", "b", "c"}, minLen: 3, maxLen: 3},
	{value: []string{"a", "b", "c"}, minLen: 1, maxLen: 4},
	{value: []string{"a", "b", "c"}, minLen: 3, maxLen: 10},
	{value: []string{"a", "b", "c"}, minLen: 0, maxLen: 3},
	{value: []string{"a", "b", "c"}, minLen: 4, maxLen: 10, expectedError: "length must be between 4 and 10"},
	{value: []string{"a", "b", "c"}, minLen: 1, maxLen: 2, expectedError: "length must be between 1 and 2"},
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

func BenchmarkSliceLength(b *testing.B) {
	for range b.N {
		for _, tc := range sliceLengthTestCases {
			_ = SliceLength[[]string](tc.minLen, tc.maxLen).Validate(tc.value)
		}
	}
}

var sliceMinLengthTestCases = []*struct {
	value         []string
	minLen        int
	expectedError string
}{
	{value: []string{"a", "b", "c"}, minLen: 0},
	{value: []string{"a", "b", "c"}, minLen: 3},
	{value: []string{"a", "b", "c"}, minLen: 4, expectedError: "length must be greater than or equal to 4"},
	{value: []string{"a", "b", "c"}, minLen: 10, expectedError: "length must be greater than or equal to 10"},
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

func BenchmarkSliceMinLength(b *testing.B) {
	for range b.N {
		for _, tc := range sliceMinLengthTestCases {
			_ = SliceMinLength[[]string](tc.minLen).Validate(tc.value)
		}
	}
}

var sliceMaxLengthTestCases = []*struct {
	value         []string
	maxLen        int
	expectedError string
}{
	{value: []string{"a", "b", "c"}, maxLen: 10},
	{value: []string{"a", "b", "c"}, maxLen: 3},
	{value: []string{"a", "b", "c"}, maxLen: 2, expectedError: "length must be less than or equal to 2"},
	{value: []string{"a", "b", "c"}, maxLen: 0, expectedError: "length must be less than or equal to 0"},
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

func BenchmarkSliceMaxLength(b *testing.B) {
	for range b.N {
		for _, tc := range sliceMaxLengthTestCases {
			_ = SliceMaxLength[[]string](tc.maxLen).Validate(tc.value)
		}
	}
}

var mapLengthTestCases = []*struct {
	value         map[string]string
	minLen        int
	maxLen        int
	expectedError string
}{
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

func BenchmarkMapLength(b *testing.B) {
	for range b.N {
		for _, tc := range mapLengthTestCases {
			_ = MapLength[map[string]string](tc.minLen, tc.maxLen).Validate(tc.value)
		}
	}
}

var mapMinLengthTestCases = []*struct {
	value         map[string]string
	minLen        int
	expectedError string
}{
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

func BenchmarkMapMinLength(b *testing.B) {
	for range b.N {
		for _, tc := range mapMinLengthTestCases {
			_ = MapMinLength[map[string]string](tc.minLen).Validate(tc.value)
		}
	}
}

var mapMaxLengthTestCases = []*struct {
	value         map[string]string
	maxLen        int
	expectedError string
}{
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

func BenchmarkMapMaxLength(b *testing.B) {
	for range b.N {
		for _, tc := range mapMaxLengthTestCases {
			_ = MapMaxLength[map[string]string](tc.maxLen).Validate(tc.value)
		}
	}
}
