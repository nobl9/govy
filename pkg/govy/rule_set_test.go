package govy_test

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func TestRuleSet(t *testing.T) {
	r := govy.NewRuleSet(
		rules.StringStartsWith("foo"),
		rules.StringEndsWith("bar"),
	)

	t.Run("passes", func(t *testing.T) {
		err := r.Validate("foo_bar")
		assert.NoError(t, err)
	})
	t.Run("first fails", func(t *testing.T) {
		err := r.Validate("baz_bar")
		assert.Require(t, assert.Error(t, err))
		errs := err.(govy.RuleSetError)
		assert.Require(t, assert.Len(t, errs, 1))
		assert.EqualError(t, errs[0], "string must start with 'foo' prefix")
	})
	t.Run("second fails", func(t *testing.T) {
		err := r.Validate("foo_baz")
		assert.Require(t, assert.Error(t, err))
		errs := err.(govy.RuleSetError)
		assert.Require(t, assert.Len(t, errs, 1))
		assert.EqualError(t, errs[0], "string must end with 'bar' suffix")
	})
	t.Run("both fail", func(t *testing.T) {
		err := r.Validate("baz_baz")
		assert.Require(t, assert.Error(t, err))
		errs := err.(govy.RuleSetError)
		assert.Require(t, assert.Len(t, errs, 2))
		assert.EqualError(t, errs[0], "string must start with 'foo' prefix")
		assert.EqualError(t, errs[1], "string must end with 'bar' suffix")
	})
}
