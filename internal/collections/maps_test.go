package collections

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestSortedKeys(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		m := map[int]string{
			3: "c",
			1: "a",
			2: "b",
		}
		keys := SortedKeys(m)
		assert.Equal(t, []int{1, 2, 3}, keys)
	})
	t.Run("strings", func(t *testing.T) {
		m := map[string]int{
			"c": 3,
			"a": 1,
			"b": 2,
		}
		keys := SortedKeys(m)
		assert.Equal(t, []string{"a", "b", "c"}, keys)
	})
}
