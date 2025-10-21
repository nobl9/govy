package govy_test

import (
	"errors"
	"testing"
	"text/template"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestRule(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New("must be positive")
			}
			return nil
		})

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, "must be positive")
	})
	t.Run("full", func(t *testing.T) {
		r := govy.NewRule(func(v string) error {
			return errors.New("error")
		}).
			WithMessagef("my message %s", "foo").
			WithDetailsf("some details %d", 1).
			WithExamples("foo", "bar")

		err := r.Validate("baz")
		assert.Require(t, assert.Error(t, err))
		assert.Equal(t, "my message foo (e.g. 'foo', 'bar'); some details 1", err.Error())
	})
}

func TestRule_WithErrorCode(t *testing.T) {
	r := govy.NewRule(func(v int) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	}).WithErrorCode("test")

	err := r.Validate(0)
	assert.NoError(t, err)
	err = r.Validate(-1)
	assert.EqualError(t, err, "must be positive")
	assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
}

func TestRule_WithMessage(t *testing.T) {
	tests := []struct {
		Error         string
		Message       string
		ExpectedError string
	}{
		{
			Error:         "this is error",
			Message:       "",
			ExpectedError: "this is error",
		},
		{
			Error:         "this is error",
			Message:       "this is message",
			ExpectedError: "this is message",
		},
		{
			Error:         "",
			Message:       "message",
			ExpectedError: "message",
		},
	}

	for _, test := range tests {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithMessage(test.Message)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
	}
}

func TestRule_WithMessagef(t *testing.T) {
	tests := []struct {
		Error         string
		Message       string
		MessageArgs   []any
		ExpectedError string
	}{
		{
			Error:         "",
			Message:       "message %s %d",
			MessageArgs:   []any{"foo", 1},
			ExpectedError: "message foo 1",
		},
		{
			Error:         "this is error",
			Message:       "message %s",
			MessageArgs:   []any{"bar"},
			ExpectedError: "message bar",
		},
	}

	for _, test := range tests {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithMessagef(test.Message, test.MessageArgs...)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
	}
}

func TestRule_WithDetails(t *testing.T) {
	tests := []struct {
		Error         string
		Details       string
		ExpectedError string
	}{
		{
			Error:         "this is error",
			Details:       "details",
			ExpectedError: "this is error; details",
		},
		{
			Error:         "this is error",
			Details:       "",
			ExpectedError: "this is error",
		},
		{
			Error:         "",
			Details:       "details",
			ExpectedError: "details",
		},
	}

	for _, test := range tests {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithDetails(test.Details)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
	}
}

func TestRule_WithDetailsf(t *testing.T) {
	tests := []struct {
		Error         string
		Details       string
		DetailsArgs   []any
		ExpectedError string
	}{
		{
			Error:         "",
			Details:       "details %s %d",
			DetailsArgs:   []any{"foo", 1},
			ExpectedError: "details foo 1",
		},
		{
			Error:         "this is error",
			Details:       "details %s",
			DetailsArgs:   []any{"bar"},
			ExpectedError: "this is error; details bar",
		},
	}

	for _, test := range tests {
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithDetailsf(test.Details, test.DetailsArgs...)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, govy.ErrorCode("test"), err.(*govy.RuleError).Code)
	}
}

func TestRule_WithDescription(t *testing.T) {
	r := govy.NewRule(func(v int) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	}).
		WithErrorCode("test").
		WithDetails("some details").
		WithDescription("the integer must be positive")

	err := r.Validate(-1)
	assert.Require(t, assert.Error(t, err))
	assert.Equal(t, &govy.RuleError{
		Message:     "must be positive; some details",
		Code:        "test",
		Description: "the integer must be positive",
	}, err)
}

func TestRule_WithExamples(t *testing.T) {
	r := govy.NewRule(func(v string) error {
		if v != "foo" && v != "bar" {
			return errors.New("must be foo or bar")
		}
		return nil
	}).
		WithErrorCode("test").
		WithDetails("some details").
		WithExamples("foo", "bar").
		WithDescription("string must be foo or bar")

	err := r.Validate("baz")
	assert.Require(t, assert.Error(t, err))
	assert.Equal(t, &govy.RuleError{
		Message:     "must be foo or bar (e.g. 'foo', 'bar'); some details",
		Code:        "test",
		Description: "string must be foo or bar",
	}, err)
}

func TestRule_WithMessageTemplate(t *testing.T) {
	tpl, err := template.New("").Parse("This is an {{ .Error }}")
	assert.Require(t, assert.NoError(t, err))

	rule := govy.NewRule(func(v string) error {
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			Error: "error",
		})
	}).
		WithErrorCode("my-code").
		WithExamples("This").
		WithMessageTemplate(tpl)

	err = rule.Validate("")
	assert.Require(t, assert.Error(t, err))
	assert.Equal(t, &govy.RuleError{
		Message: "This is an error",
		Code:    "my-code",
	}, err)
}

func TestRule_WithMessageTemplateString(t *testing.T) {
	rule := govy.NewRule(func(v string) error {
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			Error: "error",
		})
	}).
		WithErrorCode("my-code").
		WithMessageTemplateString("This is an {{ .Error }}")

	err := rule.Validate("")
	assert.Require(t, assert.Error(t, err))
	assert.Equal(t, &govy.RuleError{
		Message: "This is an error",
		Code:    "my-code",
	}, err)
}

func TestRuleToPointer(t *testing.T) {
	r := govy.NewRule(func(v int) error {
		if v < 0 {
			return errors.New("must be positive")
		}
		return nil
	}).
		WithErrorCode("test")
	rp := govy.RuleToPointer(r)
	t.Run("passes", func(t *testing.T) {
		err := rp.Validate(ptr(0))
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := rp.Validate(ptr(-1))
		assert.Require(t, assert.Error(t, err))
		t.Log(err)
		assert.Equal(t, govy.RuleError{
			Message: "must be positive",
			Code:    "test",
		}, *err.(*govy.RuleError))
	})
}
