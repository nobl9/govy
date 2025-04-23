package collections

import (
	"cmp"
	"maps"
	"slices"
)

// SortedKeys returns a sorted slice of keys of the input map.
// The keys must meet [constraints.Ordered] type constraint.
func SortedKeys[M ~map[K]V, K cmp.Ordered, V any](m M) []K {
	keys := maps.Keys(m)
	return slices.Sorted(keys)
}
