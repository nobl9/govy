package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
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
			assert.Require(t, assert.Error(t, err))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeRequired))
		}
	})
}
