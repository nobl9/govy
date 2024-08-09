package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nobl9/govy/pkg/govy"
)

func TestSliceUnique(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := SliceUnique(HashFuncSelf[string]()).Validate([]string{"a", "b", "c"})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := SliceUnique(HashFuncSelf[string]()).Validate([]string{"a", "b", "c", "b"})
		require.Error(t, err)
		assert.EqualError(t, err, "elements are not unique, index 1 collides with index 3")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceUnique))
	})
	t.Run("fails with constraint", func(t *testing.T) {
		err := SliceUnique(HashFuncSelf[string](), "values must be unique").
			Validate([]string{"a", "b", "c", "b"})
		require.Error(t, err)
		assert.EqualError(t, err, "elements are not unique, index 1 collides with index 3 "+
			"based on constraints: values must be unique")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceUnique))
	})
	t.Run("fails with constraints", func(t *testing.T) {
		err := SliceUnique(HashFuncSelf[string](), "constraint 1", "constraint 2").
			Validate([]string{"a", "b", "c", "b"})
		require.Error(t, err)
		assert.EqualError(t, err, "elements are not unique, index 1 collides with index 3 "+
			"based on constraints: constraint 1, constraint 2")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceUnique))
	})
}
