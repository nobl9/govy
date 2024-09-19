package govy_test

import (
	"errors"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestValidator(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		r := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return nil })),
		)
		err := r.Validate(mockValidatorStruct{})
		assert.NoError(t, err)
	})

	t.Run("errors", func(t *testing.T) {
		err1 := errors.New("1")
		err2 := errors.New("2")
		r := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return nil })),
			govy.For(func(m mockValidatorStruct) string { return "name" }).
				WithName("test.name").
				Rules(govy.NewRule(func(v string) error { return err1 })),
			govy.For(func(m mockValidatorStruct) string { return "display" }).
				WithName("test.display").
				Rules(govy.NewRule(func(v string) error { return err2 })),
		)
		err := mustValidatorError(t, r.Validate(mockValidatorStruct{}))
		assert.Require(t, assert.Len(t, err.Errors, 2))
		assert.Equal(t, &govy.ValidatorError{Errors: govy.PropertyErrors{
			&govy.PropertyError{
				PropertyName:  "test.name",
				PropertyValue: "name",
				Errors:        []*govy.RuleError{{Message: err1.Error()}},
			},
			&govy.PropertyError{
				PropertyName:  "test.display",
				PropertyValue: "display",
				Errors:        []*govy.RuleError{{Message: err2.Error()}},
			},
		}}, err)
	})
}

func TestValidatorWhen(t *testing.T) {
	t.Run("when condition is not met, don't validate", func(t *testing.T) {
		r := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
		).
			When(func(validatorStruct mockValidatorStruct) bool { return false })

		err := r.Validate(mockValidatorStruct{})
		assert.NoError(t, err)
	})
	t.Run("when condition is met, validate", func(t *testing.T) {
		r := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
		).
			When(func(validatorStruct mockValidatorStruct) bool { return true })

		err := mustValidatorError(t, r.Validate(mockValidatorStruct{}))
		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, &govy.ValidatorError{Errors: govy.PropertyErrors{
			&govy.PropertyError{
				PropertyName:  "test",
				PropertyValue: "test",
				Errors:        []*govy.RuleError{{Message: "test"}},
			},
		}}, err)
	})
}

func TestValidatorWithName(t *testing.T) {
	r := govy.New(
		govy.For(func(m mockValidatorStruct) string { return "test" }).
			WithName("test").
			Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
	).WithName("validator")

	err := r.Validate(mockValidatorStruct{})
	assert.Require(t, assert.Error(t, err))
	assert.EqualError(t, err, `Validation for validator has failed for the following properties:
  - 'test' with value 'test':
    - test`)
}

func TestValidatorInferName(t *testing.T) {
	r := govy.New(
		govy.For(func(m mockValidatorStruct) string { return "test" }).
			WithName("test").
			Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
	).InferName()

	err := r.Validate(mockValidatorStruct{})
	assert.Require(t, assert.Error(t, err))
	assert.EqualError(t, err, `Validation for mockValidatorStruct has failed for the following properties:
  - 'test' with value 'test':
    - test`)
}

func mustValidatorError(t *testing.T, err error) *govy.ValidatorError {
	t.Helper()
	return mustErrorType[*govy.ValidatorError](t, err)
}

type mockValidatorStruct struct {
	Field string
}
