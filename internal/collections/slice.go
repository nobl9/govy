package collections

import (
	"cmp"
	"slices"

	"github.com/nobl9/govy/internal/stringconvert"
)

// ToStringSlice converts a slice of T to a slice of strings.
func ToStringSlice[T any](s []T) []string {
	return mapSlice(s, func(v T) string { return stringconvert.Format(v) })
}

// mapSlice applies a mapping function f to each element of the slice (type T)
// and returns a new slice with the results mapped to type N.
func mapSlice[T, N any](s []T, f func(T) N) []N {
	result := make([]N, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

// Intersection returns the set intersection between provided slices.
func Intersection[T cmp.Ordered](s ...[]T) []T {
	if len(s) == 0 {
		return nil
	}
	if len(s) == 1 {
		return s[0]
	}

	elements := make(map[T]int, len(s[0]))
	for _, v := range s[0] {
		elements[v] = 1
	}

	for i := 1; i < len(s); i++ {
		seen := make(map[T]bool)
		for _, v := range s[i] {
			if _, exists := elements[v]; exists && !seen[v] {
				elements[v]++
				seen[v] = true
			}
		}
	}

	var result []T
	for item, count := range elements {
		if count == len(s) {
			result = append(result, item)
		}
	}
	slices.Sort(result)
	return result
}

// EqualSlices compares two slices, checking if every value under each index is equal in both slices.
// Each slice's elements must satisfy [comparable] interface.
func EqualSlices[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
