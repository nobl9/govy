package collections

import "github.com/nobl9/govy/internal/stringconvert"

// GenericToString converts a slice of T to a slice of string.
func GenericToString[T any](s []T) []string {
	return mapSlice(s, func(v T) string { return stringconvert.Convert(v) })
}

// mapSlice applies a mapping function f to each element of the slice (type T)
// and returns a new slice with the results mapped to type N.
func mapSlice[T any, N any](s []T, f func(T) N) []N {
	result := make([]N, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}
