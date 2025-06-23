package collections

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestToStringSlice(t *testing.T) {
	testCases := []struct {
		f        func() []string
		expected []string
	}{
		{
			f:        func() []string { return ToStringSlice([]int{1, 2, 4}) },
			expected: []string{"1", "2", "4"},
		},
		{
			f:        func() []string { return ToStringSlice([]string{"1", "2", "4"}) },
			expected: []string{"1", "2", "4"},
		},
		{
			f:        func() []string { return ToStringSlice([]any{1, "2", []float64{0.64}}) },
			expected: []string{"1", "2", "[0.64]"},
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.f())
	}
}

func TestMapSlice(t *testing.T) {
	testCases := []struct {
		f        func() any
		expected any
	}{
		{
			f:        func() any { return mapSlice([]int{1, 2, 4}, strconv.Itoa) },
			expected: []string{"1", "2", "4"},
		},
		{
			f: func() any {
				return mapSlice([]string{"1", "2", "4"}, func(s string) int {
					i, _ := strconv.Atoi(s)
					return i
				})
			},
			expected: []int{1, 2, 4},
		},
		{
			f: func() any {
				return mapSlice([]any{1, "2", []float64{0.64}}, func(v any) string { return fmt.Sprint(v) })
			},
			expected: []string{"1", "2", "[0.64]"},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, tc.f())
	}
}

func TestIntersection(t *testing.T) {
	for _, tc := range []struct {
		in       [][]string
		expected []string
	}{
		{
			in:       nil,
			expected: nil,
		},
		{
			in:       [][]string{},
			expected: nil,
		},
		{
			in:       [][]string{{"a", "b", "c"}},
			expected: []string{"a", "b", "c"},
		},
		{
			in:       [][]string{{"a", "b", "c"}, {}, {"b"}},
			expected: nil,
		},
		{
			in:       [][]string{{"a", "b", "c"}, {"b"}, {"b", "c"}},
			expected: []string{"b"},
		},
		{
			in:       [][]string{{"x", "b", "c", "a"}, {"a", "b", "c"}, {"b", "a", "c", "d"}},
			expected: []string{"a", "b", "c"},
		},
	} {
		actual := Intersection(tc.in...)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestEqualSlices(t *testing.T) {
	testCases := []struct {
		s1, s2   []int
		expected bool
	}{
		{s1: nil, s2: nil, expected: true},
		{s1: []int{}, s2: []int{}, expected: true},
		{s1: []int{1, 2, 3}, s2: []int{1, 2, 3}, expected: true},
		{s1: []int{1, 2, 3}, s2: []int{3, 2, 1}, expected: false},
		{s1: []int{1, 2}, s2: []int{1, 2, 3}, expected: false},
		{s1: []int{1, 2, 3}, s2: []int{1, 2}, expected: false},
		{s1: []int{1, 2, 3}, s2: nil, expected: false},
		{s1: nil, s2: []int{1, 2, 3}, expected: false},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.expected, EqualSlices(tc.s1, tc.s2))
	}
}
