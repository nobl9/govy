package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEqualTo(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := EqualTo(1.1).Validate(1.1)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := EqualTo(1.1).Validate(1.3)
		require.Error(t, err)
		assert.EqualError(t, err, "should be equal to '1.1'")
		assert.True(t, HasErrorCode(err, ErrorCodeEqualTo))
	})
}

func TestNotEqualTo(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := NotEqualTo(1.1).Validate(1.3)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := NotEqualTo(1.1).Validate(1.1)
		require.Error(t, err)
		assert.EqualError(t, err, "should be not equal to '1.1'")
		assert.True(t, HasErrorCode(err, ErrorCodeNotEqualTo))
	})
}

func TestGreaterThan(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := GreaterThan(1).Validate(2)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		for n, v := range map[int]int{1: 1, 4: 2} {
			err := GreaterThan(n).Validate(v)
			require.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("should be greater than '%v'", n))
			assert.True(t, HasErrorCode(err, ErrorCodeGreaterThan))
		}
	})
}

func TestGreaterThanOrEqual(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for n, v := range map[int]int{1: 1, 2: 4} {
			err := GreaterThanOrEqualTo(n).Validate(v)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		err := GreaterThanOrEqualTo(4).Validate(2)
		require.Error(t, err)
		assert.EqualError(t, err, "should be greater than or equal to '4'")
		assert.True(t, HasErrorCode(err, ErrorCodeGreaterThanOrEqualTo))
	})
}

func TestLessThan(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := LessThan(4).Validate(2)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		for n, v := range map[int]int{1: 1, 2: 4} {
			err := LessThan(n).Validate(v)
			require.Error(t, err)
			assert.EqualError(t, err, fmt.Sprintf("should be less than '%v'", n))
			assert.True(t, HasErrorCode(err, ErrorCodeLessThan))
		}
	})
}

func TestLessThanOrEqual(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for n, v := range map[int]int{1: 1, 4: 2} {
			err := LessThanOrEqualTo(n).Validate(v)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		err := LessThanOrEqualTo(2).Validate(4)
		require.Error(t, err)
		assert.EqualError(t, err, "should be less than or equal to '2'")
		assert.True(t, HasErrorCode(err, ErrorCodeLessThanOrEqualTo))
	})
}
