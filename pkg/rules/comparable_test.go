package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestEQ(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := EQ(1.1).Validate(1.1)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := EQ(1.1).Validate(1.3)
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "should be equal to '1.1'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeEqualTo))
	})
}

func TestNEQ(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := NEQ(1.1).Validate(1.3)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := NEQ(1.1).Validate(1.1)
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "should be not equal to '1.1'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeNotEqualTo))
	})
}

func TestGT(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := GT(1).Validate(2)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		for n, v := range map[int]int{1: 1, 4: 2} {
			err := GT(n).Validate(v)
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, fmt.Sprintf("should be greater than '%v'", n))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGreaterThan))
		}
	})
}

func TestGTE(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for n, v := range map[int]int{1: 1, 2: 4} {
			err := GTE(n).Validate(v)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		err := GTE(4).Validate(2)
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "should be greater than or equal to '4'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeGreaterThanOrEqualTo))
	})
}

func TestLT(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := LT(4).Validate(2)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		for n, v := range map[int]int{1: 1, 2: 4} {
			err := LT(n).Validate(v)
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, fmt.Sprintf("should be less than '%v'", n))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLessThan))
		}
	})
}

func TestLTE(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for n, v := range map[int]int{1: 1, 4: 2} {
			err := LTE(n).Validate(v)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		err := LTE(2).Validate(4)
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "should be less than or equal to '2'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeLessThanOrEqualTo))
	})
}
