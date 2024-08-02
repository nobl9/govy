package validation

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidator(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		r := New(
			For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(NewSingleRule(func(v string) error { return nil })),
		)
		errs := r.Validate(mockValidatorStruct{})
		assert.Nil(t, errs)
	})

	t.Run("errors", func(t *testing.T) {
		err1 := errors.New("1")
		err2 := errors.New("2")
		r := New(
			For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(NewSingleRule(func(v string) error { return nil })),
			For(func(m mockValidatorStruct) string { return "name" }).
				WithName("test.name").
				Rules(NewSingleRule(func(v string) error { return err1 })),
			For(func(m mockValidatorStruct) string { return "display" }).
				WithName("test.display").
				Rules(NewSingleRule(func(v string) error { return err2 })),
		)
		err := r.Validate(mockValidatorStruct{})
		require.Len(t, err.Errors, 2)
		assert.Equal(t, &ValidatorError{Errors: PropertyErrors{
			&PropertyError{
				PropertyName:  "test.name",
				PropertyValue: "name",
				Errors:        []*RuleError{{Message: err1.Error()}},
			},
			&PropertyError{
				PropertyName:  "test.display",
				PropertyValue: "display",
				Errors:        []*RuleError{{Message: err2.Error()}},
			},
		}}, err)
	})
}

func TestValidatorWhen(t *testing.T) {
	t.Run("when condition is not met, don't validate", func(t *testing.T) {
		r := New(
			For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(NewSingleRule(func(v string) error { return errors.New("test") })),
		).
			When(func(validatorStruct mockValidatorStruct) bool { return false })

		errs := r.Validate(mockValidatorStruct{})
		assert.Nil(t, errs)
	})
	t.Run("when condition is met, validate", func(t *testing.T) {
		r := New(
			For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(NewSingleRule(func(v string) error { return errors.New("test") })),
		).
			When(func(validatorStruct mockValidatorStruct) bool { return true })

		errs := r.Validate(mockValidatorStruct{})
		require.Len(t, errs.Errors, 1)
		assert.Equal(t, &ValidatorError{Errors: PropertyErrors{
			&PropertyError{
				PropertyName:  "test",
				PropertyValue: "test",
				Errors:        []*RuleError{{Message: "test"}},
			},
		}}, errs)
	})
}

func TestValidatorWithName(t *testing.T) {
	r := New(
		For(func(m mockValidatorStruct) string { return "test" }).
			WithName("test").
			Rules(NewSingleRule(func(v string) error { return errors.New("test") })),
	).WithName("validator")

	err := r.Validate(mockValidatorStruct{})
	assert.EqualError(t, err, `Validation for validator has failed for the following properties:
  - 'test' with value 'test':
    - test`)
}

type mockValidatorStruct struct {
	Field string
}
