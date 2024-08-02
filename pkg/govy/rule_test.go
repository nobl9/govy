package validation

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestSingleRule(t *testing.T) {
	r := NewSingleRule(func(v int) error {
		if v < 0 {
			return errors.Errorf("must be positive")
		}
		return nil
	})

	err := r.Validate(0)
	assert.Nil(t, err)
	err = r.Validate(-1)
	assert.EqualError(t, err, "must be positive")
}

func TestSingleRule_WithErrorCode(t *testing.T) {
	r := NewSingleRule(func(v int) error {
		if v < 0 {
			return errors.Errorf("must be positive")
		}
		return nil
	}).WithErrorCode("test")

	err := r.Validate(0)
	assert.Nil(t, err)
	err = r.Validate(-1)
	assert.EqualError(t, err, "must be positive")
	assert.Equal(t, "test", err.(*RuleError).Code)
}

func TestSingleRule_WithMessage(t *testing.T) {
	for _, test := range []struct {
		Error         string
		Message       string
		Details       string
		ExpectedError string
	}{
		{
			Error:         "this is error",
			Message:       "",
			Details:       "details",
			ExpectedError: "this is error; details",
		},
		{
			Error:         "this is error",
			Message:       "this is message",
			Details:       "",
			ExpectedError: "this is message",
		},
		{
			Error:         "",
			Message:       "message",
			Details:       "details",
			ExpectedError: "message; details",
		},
	} {
		r := NewSingleRule(func(v int) error {
			if v < 0 {
				return errors.Errorf(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithMessage(test.Message).
			WithDetails(test.Details)

		err := r.Validate(0)
		assert.Nil(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, "test", err.(*RuleError).Code)
	}
}

func TestSingleRule_WithDetails(t *testing.T) {
	for _, test := range []struct {
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
	} {
		r := NewSingleRule(func(v int) error {
			if v < 0 {
				return errors.Errorf(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithDetails(test.Details)

		err := r.Validate(0)
		assert.Nil(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, "test", err.(*RuleError).Code)
	}
}
