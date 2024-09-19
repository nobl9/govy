package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestStringLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringLength(0, 4).Validate("test")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		for minLen, maxLen := range map[int]int{
			0:  2,
			10: 20,
		} {
			err := StringLength(minLen, maxLen).Validate("test")
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, fmt.Sprintf("length must be between %d and %d", minLen, maxLen))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringLength))
		}
	})
}

func TestStringMinLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringMinLength(0).Validate("test")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringMinLength(5).Validate("test")
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "length must be greater than or equal to 5")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMinLength))
	})
}

func TestStringMaxLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringMaxLength(4).Validate("test")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringMaxLength(3).Validate("test")
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "length must be less than or equal to 3")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMaxLength))
	})
}

func TestSliceLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := SliceLength[[]string](0, 1).Validate([]string{"test"})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		for minLen, maxLen := range map[int]int{
			0: 1,
			3: 10,
		} {
			err := SliceLength[[]string](minLen, maxLen).Validate([]string{"test", "test"})
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, fmt.Sprintf("length must be between %d and %d", minLen, maxLen))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceLength))
		}
	})
}

func TestSliceMinLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := SliceMinLength[[]string](1).Validate([]string{"test"})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := SliceMinLength[[]string](2).Validate([]string{"test"})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "length must be greater than or equal to 2")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceMinLength))
	})
}

func TestSliceMaxLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := SliceMaxLength[[]string](1).Validate([]string{"test"})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := SliceMaxLength[[]string](1).Validate([]string{"1", "2"})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "length must be less than or equal to 1")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceMaxLength))
	})
}

func TestMapLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := MapLength[map[string]string](0, 1).Validate(map[string]string{"this": "that"})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		for minLen, maxLen := range map[int]int{
			0: 1,
			3: 10,
		} {
			err := MapLength[map[string]string](minLen, maxLen).Validate(map[string]string{"a": "b", "c": "d"})
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, fmt.Sprintf("length must be between %d and %d", minLen, maxLen))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMapLength))
		}
	})
}

func TestMapMinLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := MapMinLength[map[string]string](1).Validate(map[string]string{"a": "b"})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := MapMinLength[map[string]string](2).Validate(map[string]string{"a": "b"})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "length must be greater than or equal to 2")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeMapMinLength))
	})
}

func TestMapMaxLength(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := MapMaxLength[map[string]string](1).Validate(map[string]string{"a": "b"})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := MapMaxLength[map[string]string](1).Validate(map[string]string{"a": "b", "c": "d"})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "length must be less than or equal to 1")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeMapMaxLength))
	})
}
