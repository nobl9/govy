package uuid

import (
	"regexp"
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestGenerateUUID(t *testing.T) {
	t.Run("generates valid UUID format", func(t *testing.T) {
		id := GenerateUUID()

		// UUID v4 format:
		pattern := `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
		matched, err := regexp.MatchString(pattern, id)
		assert.NoError(t, err)
		assert.True(t, matched)
	})

	t.Run("generates unique IDs", func(t *testing.T) {
		ids := make(map[string]bool)
		iterations := 1000

		for range iterations {
			id := GenerateUUID()

			assert.False(t, ids[id])
			ids[id] = true
		}

		assert.Equal(t, iterations, len(ids))
	})

	t.Run("has correct length", func(t *testing.T) {
		id := GenerateUUID()

		expectedLen := 36 // 32 hex chars + 4 hyphens
		assert.Equal(t, expectedLen, len(id))
	})

	t.Run("has correct version bits", func(t *testing.T) {
		id := GenerateUUID()
		// Version should be 4 (at position 14)
		assert.Equal(t, byte('4'), id[14])
	})

	t.Run("has correct variant bits", func(t *testing.T) {
		id := GenerateUUID()
		// Variant should be 8, 9, a, or b (at position 19)
		variant := id[19]
		assert.True(t, variant == '8' || variant == '9' || variant == 'a' || variant == 'b')
	})
}
