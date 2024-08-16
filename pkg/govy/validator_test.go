package govy_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nobl9/govy/pkg/govy"
)

func TestValidator(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		r := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewSingleRule(func(v string) error { return nil })),
		)
		errs := r.Validate(mockValidatorStruct{})
		assert.Nil(t, errs)
	})

	t.Run("errors", func(t *testing.T) {
		err1 := errors.New("1")
		err2 := errors.New("2")
		r := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewSingleRule(func(v string) error { return nil })),
			govy.For(func(m mockValidatorStruct) string { return "name" }).
				WithName("test.name").
				Rules(govy.NewSingleRule(func(v string) error { return err1 })),
			govy.For(func(m mockValidatorStruct) string { return "display" }).
				WithName("test.display").
				Rules(govy.NewSingleRule(func(v string) error { return err2 })),
		)
		err := r.Validate(mockValidatorStruct{})
		require.Len(t, err.Errors, 2)
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
				Rules(govy.NewSingleRule(func(v string) error { return errors.New("test") })),
		).
			When(func(validatorStruct mockValidatorStruct) bool { return false })

		errs := r.Validate(mockValidatorStruct{})
		assert.Nil(t, errs)
	})
	t.Run("when condition is met, validate", func(t *testing.T) {
		r := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewSingleRule(func(v string) error { return errors.New("test") })),
		).
			When(func(validatorStruct mockValidatorStruct) bool { return true })

		errs := r.Validate(mockValidatorStruct{})
		require.Len(t, errs.Errors, 1)
		assert.Equal(t, &govy.ValidatorError{Errors: govy.PropertyErrors{
			&govy.PropertyError{
				PropertyName:  "test",
				PropertyValue: "test",
				Errors:        []*govy.RuleError{{Message: "test"}},
			},
		}}, errs)
	})
}

func TestValidatorWithName(t *testing.T) {
	r := govy.New(
		govy.For(func(m mockValidatorStruct) string { return "test" }).
			WithName("test").
			Rules(govy.NewSingleRule(func(v string) error { return errors.New("test") })),
	).WithName("validator")

	err := r.Validate(mockValidatorStruct{})
	assert.EqualError(t, err, `Validation for validator has failed for the following properties:
  - 'test' with value 'test':
    - test`)
}

type mockValidatorStruct struct {
	Field string
}
