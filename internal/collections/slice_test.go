package collections

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestMapSlice(t *testing.T) {
	testCases := []struct {
		f        func() any
		expected any
	}{
		{
			f:        func() any { return MapSlice([]int{1, 2, 4}, strconv.Itoa) },
			expected: []string{"1", "2", "4"},
		},
		{
			f: func() any {
				return MapSlice([]string{"1", "2", "4"}, func(s string) int {
					i, _ := strconv.Atoi(s)
					return i
				})
			},
			expected: []int{1, 2, 4},
		},
		{
			f: func() any {
				return MapSlice([]any{1, "2", []float64{0.64}}, func(v any) string { return fmt.Sprint(v) })
			},
			expected: []string{"1", "2", "[0.64]"},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.f())
	}
}

func TestGenericToAny(t *testing.T) {
	testCases := []struct {
		f        func() any
		expected any
	}{
		{
			f:        func() any { return GenericToAny([]int{1, 2, 4}) },
			expected: []any{1, 2, 4},
		},
		{
			f:        func() any { return GenericToAny([]string{"1", "2", "4"}) },
			expected: []any{"1", "2", "4"},
		},
		{
			f:        func() any { return GenericToAny([]any{1, "2", []float64{0.64}}) },
			expected: []any{1, "2", []float64{0.64}},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.f())
	}
}
