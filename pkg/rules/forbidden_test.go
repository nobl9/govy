package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nobl9/govy/pkg/govy"
)

func TestForbidden(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := Forbidden[string]().Validate("")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := Forbidden[string]().Validate("test")
		require.Error(t, err)
		assert.EqualError(t, err, "property is forbidden")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeForbidden))
	})
}
