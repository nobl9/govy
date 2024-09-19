package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestForbidden(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := Forbidden[string]().Validate("")
		assert.Equal(t, nil, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := Forbidden[string]().Validate("test")
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "property is forbidden")
		assert.Equal(t, true, govy.HasErrorCode(err, ErrorCodeForbidden))
	})
}
