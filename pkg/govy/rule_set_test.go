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

func TestRuleSetWithErrorCode(t *testing.T) {
	r := govy.NewRuleSet(
		rules.StringStartsWith("foo"),
		rules.StringEndsWith("bar"),
	).
		WithErrorCode("my-code")

	err := r.Validate("baz_bar")

	assert.Require(t, assert.Error(t, err))
	errs := err.(govy.RuleSetError)
	assert.Require(t, assert.Len(t, errs, 1))
	ruleErr := errs[0].(*govy.RuleError)
	assert.Equal(t, govy.RuleError{
		Message:     "string must start with 'foo' prefix",
		Code:        "my-code:string_starts_with",
		Description: "string must start with 'foo' prefix",
	}, *ruleErr)
}

func TestRuleSetToPointer(t *testing.T) {
	r := govy.NewRuleSet(
		rules.StringStartsWith("foo"),
		rules.StringEndsWith("bar"),
	).
		WithErrorCode("my-code")
	rp := govy.RuleSetToPointer(r)

	t.Run("passes", func(t *testing.T) {
		err := rp.Validate(ptr("foo+bar"))
		assert.NoError(t, err)
	})
	t.Run("skip nil", func(t *testing.T) {
		err := rp.Validate(nil)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := rp.Validate(ptr("baz-bar"))
		assert.Require(t, assert.Error(t, err))
		errs := err.(govy.RuleSetError)
		assert.Require(t, assert.Len(t, errs, 1))
		assert.EqualError(t, errs[0], "string must start with 'foo' prefix")
		assert.Equal(t, govy.ErrorCode("my-code:string_starts_with"), errs[0].(*govy.RuleError).Code)
	})
}

func TestRuleSetCascade(t *testing.T) {
	r := govy.NewRuleSet(
		rules.StringStartsWith("foo"),
		rules.StringEndsWith("bar"),
	)

	prefixErr := &govy.RuleError{
		Message:     "string must start with 'foo' prefix",
		Code:        "string_starts_with",
		Description: "string must start with 'foo' prefix",
	}
	suffixErr := &govy.RuleError{
		Message:     "string must end with 'bar' suffix",
		Code:        "string_ends_with",
		Description: "string must end with 'bar' suffix",
	}

	tests := map[string]struct {
		ruleSet     govy.RuleSet[string]
		expectedErr govy.RuleSetError
	}{
		"default continue": {
			ruleSet:     r,
			expectedErr: govy.RuleSetError{prefixErr, suffixErr},
		},
		"continue mode": {
			ruleSet:     r.Cascade(govy.CascadeModeContinue),
			expectedErr: govy.RuleSetError{prefixErr, suffixErr},
		},
		"stop mode": {
			ruleSet:     r.Cascade(govy.CascadeModeStop),
			expectedErr: govy.RuleSetError{prefixErr},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.ruleSet.Validate("baz_baz")

			assert.Require(t, assert.Error(t, err))
			errs := err.(govy.RuleSetError)
			assert.Require(t, assert.Len(t, errs, len(test.expectedErr)))
			assert.Equal(t, test.expectedErr, errs)
		})
	}
}
