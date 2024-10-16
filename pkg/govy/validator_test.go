package govy_test

import (
	"errors"
	"testing"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func TestValidator(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return nil })),
		)
		err := v.Validate(mockValidatorStruct{})
		assert.NoError(t, err)
	})

	t.Run("errors", func(t *testing.T) {
		err1 := errors.New("1")
		err2 := errors.New("2")
		v := govy.New(
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
		err := mustValidatorError(t, v.Validate(mockValidatorStruct{}))
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
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
		).
			When(func(validatorStruct mockValidatorStruct) bool { return false })

		err := v.Validate(mockValidatorStruct{})
		assert.NoError(t, err)
	})
	t.Run("when condition is met, validate", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
		).
			When(func(validatorStruct mockValidatorStruct) bool { return true })

		err := mustValidatorError(t, v.Validate(mockValidatorStruct{}))
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
	v := govy.New(
		govy.For(func(m mockValidatorStruct) string { return "test" }).
			WithName("test").
			Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
	).WithName("validator")

	err := v.Validate(mockValidatorStruct{})
	assert.Require(t, assert.Error(t, err))
	assert.EqualError(t, err, `Validation for validator has failed for the following properties:
  - 'test' with value 'test':
    - test`)
}

func TestValidatorWithNameFunc(t *testing.T) {
	v := govy.New(
		govy.For(func(m mockValidatorStruct) string { return "test" }).
			WithName("test").
			Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
	).WithNameFunc(func(m mockValidatorStruct) string { return "validator with field: " + m.Field })

	err := v.Validate(mockValidatorStruct{Field: "FIELD"})
	assert.Require(t, assert.Error(t, err))
	assert.EqualError(t, err, `Validation for validator with field: FIELD has failed for the following properties:
  - 'test' with value 'test':
    - test`)
}

func TestValidatorInferName(t *testing.T) {
	t.Run("infer name", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
		).InferName()

		err := v.Validate(mockValidatorStruct{})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, `Validation for mockValidatorStruct has failed for the following properties:
  - 'test' with value 'test':
    - test`)
	})
	t.Run("do not infer name if name was already set", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
		).
			WithName("myValidator").
			InferName()

		err := v.Validate(mockValidatorStruct{})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, `Validation for myValidator has failed for the following properties:
  - 'test' with value 'test':
    - test`)
	})
}

func TestValidatorValidateSlice(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return nil })),
		)
		err := v.ValidateSlice([]mockValidatorStruct{{}})
		assert.NoError(t, err)
	})

	t.Run("errors", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("Field").
				Required(),
		).WithName("mock")
		errs := mustValidatorErrors(t, v.ValidateSlice([]mockValidatorStruct{
			{Field: "0"},
			{},
			{Field: "2"},
			{},
		}))
		assert.Require(t, assert.Len(t, errs, 2))
		assert.Require(t, assert.Len(t, errs[0].Errors, 1))
		assert.Equal(t, govy.ValidatorErrors{
			{
				Name:       "mock",
				SliceIndex: ptr(1),
				Errors: govy.PropertyErrors{
					&govy.PropertyError{
						PropertyName: "Field",
						Errors: []*govy.RuleError{{
							Message: internal.RequiredErrorMessage,
							Code:    rules.ErrorCodeRequired,
						}},
					},
				},
			},
			{
				Name:       "mock",
				SliceIndex: ptr(3),
				Errors: govy.PropertyErrors{
					&govy.PropertyError{
						PropertyName: "Field",
						Errors: []*govy.RuleError{{
							Message: internal.RequiredErrorMessage,
							Code:    rules.ErrorCodeRequired,
						}},
					},
				},
			},
		}, errs)
	})
}

func mustValidatorError(t *testing.T, err error) *govy.ValidatorError {
	t.Helper()
	return mustErrorType[*govy.ValidatorError](t, err)
}

func mustValidatorErrors(t *testing.T, err error) govy.ValidatorErrors {
	t.Helper()
	return mustErrorType[govy.ValidatorErrors](t, err)
}

type mockValidatorStruct struct {
	Field string
}
