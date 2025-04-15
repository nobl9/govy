package collections

import "github.com/nobl9/govy/internal/stringconvert"

// ToStringSlice converts a slice of T to a slice of strings.
func ToStringSlice[T any](s []T) []string {
	return MapSlice(s, func(v T) string { return stringconvert.Format(v) })
}

// MapSlice applies a mapping function f to each element of the slice (type T)
// and returns a new slice with the results mapped to type N.
func MapSlice[T, N any](s []T, f func(T) N) []N {
	result := make([]N, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}
