package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequired(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for _, v := range []interface{}{
			1,
			"s",
			0.1,
			[]int{},
			map[string]int{},
		} {
			err := Required[any]().Validate(v)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		for _, v := range []interface{}{
			nil,
			struct{}{},
			"",
			false,
			0,
			0.0,
		} {
			err := Required[any]().Validate(v)
			require.Error(t, err)
			assert.True(t, HasErrorCode(err, ErrorCodeRequired))
		}
	})
}
