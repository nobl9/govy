package govy_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nobl9/govy/pkg/govy"
)

func TestRule(t *testing.T) {
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
	assert.Equal(t, "test", err.(*govy.RuleError).Code)
}

func TestRule_WithMessage(t *testing.T) {
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
		r := govy.NewRule(func(v int) error {
			if v < 0 {
				return errors.New(test.Error)
			}
			return nil
		}).
			WithErrorCode("test").
			WithMessage(test.Message).
			WithDetails(test.Details)

		err := r.Validate(0)
		assert.NoError(t, err)
		err = r.Validate(-1)
		assert.EqualError(t, err, test.ExpectedError)
		assert.Equal(t, "test", err.(*govy.RuleError).Code)
	}
}

func TestRule_WithDetails(t *testing.T) {
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
		assert.Equal(t, "test", err.(*govy.RuleError).Code)
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
	require.Error(t, err)
	assert.Equal(t, &govy.RuleError{
		Message:     "must be positive",
		Code:        "test",
		Details:     "some details",
		Description: "the integer must be positive",
	}, err)
}
