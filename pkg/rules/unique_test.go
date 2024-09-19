package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestSliceUnique(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := SliceUnique(HashFuncSelf[string]()).Validate([]string{"a", "b", "c"})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := SliceUnique(HashFuncSelf[string]()).Validate([]string{"a", "b", "c", "b"})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "elements are not unique, 2nd and 4th elements collide")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceUnique))
	})
	t.Run("fails with constraint", func(t *testing.T) {
		err := SliceUnique(HashFuncSelf[string](), "values must be unique").
			Validate([]string{"a", "a", "c", "b"})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "elements are not unique, 1st and 2nd elements collide "+
			"based on constraints: values must be unique")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceUnique))
	})
	t.Run("fails with constraints", func(t *testing.T) {
		err := SliceUnique(HashFuncSelf[string](), "constraint 1", "constraint 2").
			Validate([]string{"a", "c", "c", "b"})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "elements are not unique, 2nd and 3rd elements collide "+
			"based on constraints: constraint 1, constraint 2")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceUnique))
	})
}
