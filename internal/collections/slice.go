package collections

// MapSlice applies a mapping function f to each element of the slice (type T)
// and returns a new slice with the results mapped to type N.
func MapSlice[T any, N any](s []T, f func(T) N) []N {
	result := make([]N, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

// GenericToAny converts a slice of T to a slice of any.
func GenericToAny[T any](s []T) []any {
	return MapSlice(s, func(v T) any { return v })
}
