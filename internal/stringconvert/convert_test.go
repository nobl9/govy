package stringconvert

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestFormat(t *testing.T) {
	testCases := []struct {
		v        any
		expected string
	}{
		{
			v:        1,
			expected: "1",
		},
		{
			v:        "1",
			expected: "1",
		},
		{
			v:        []int{1, 2, 4},
			expected: "[1, 2, 4]",
		},
		{
			v:        []string{"1", "2", "4"},
			expected: "[1, 2, 4]",
		},
		{
			v:        []any{1, "2", []float64{0.64}},
			expected: "[1, 2, [0.64]]",
		},
		{
			v:        []any{1, "2", []float64{0.64, 0.1, 0.0}},
			expected: "[1, 2, [0.64, 0.1, 0]]",
		},
		{
			v: []struct {
				Foo string `json:"foo"`
				Bar int    `json:"bar"`
			}{
				{"foo", 1},
			},
			expected: `[{"foo":"foo","bar":1}]`,
		},
		{
			v:        map[string]any{"a": 1, "b": "2", "c": []float64{0.64, 0.1, 0.0}},
			expected: `{"a":1,"b":"2","c":[0.64,0.1,0]}`,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, Format(tc.v))
	}
}
