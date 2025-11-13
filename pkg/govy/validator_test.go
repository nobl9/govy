package govy_test

import (
	"errors"
	"testing"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govytest"
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
							Message: internal.RequiredMessage,
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
							Message: internal.RequiredMessage,
							Code:    rules.ErrorCodeRequired,
						}},
					},
				},
			},
		}, errs)
	})
}

func TestValidatorCascade(t *testing.T) {
	propertyRulesNeverFails := govy.For(func(m mockValidatorStruct) string { return "test" }).
		WithName("neverFails").
		Rules(
			govy.NewRule(func(v string) error { return nil }),
		)
	propertyRules := govy.For(func(m mockValidatorStruct) string { return "test" }).
		WithName("string").
		Rules(
			govy.NewRule(func(v string) error { return errors.New("1") }),
			govy.NewRule(func(v string) error { return errors.New("2") }),
		)
	propertyRulesForSlice := govy.ForSlice(func(m mockValidatorStruct) []string { return []string{"test"} }).
		WithName("slice").
		Rules(
			govy.NewRule(func(v []string) error { return errors.New("1") }),
			govy.NewRule(func(v []string) error { return errors.New("2") }),
		)
	propertyRulesForMap := govy.ForMap(
		func(m mockValidatorStruct) map[string]string { return map[string]string{"test": "v"} },
	).
		WithName("map").
		Rules(
			govy.NewRule(func(v map[string]string) error { return errors.New("1") }),
			govy.NewRule(func(v map[string]string) error { return errors.New("2") }),
		)
	validator := govy.New(
		propertyRulesNeverFails,
		propertyRules,
		propertyRulesForSlice,
		propertyRulesForMap,
	)

	allErrors := []govytest.ExpectedRuleError{
		{PropertyName: "string", Message: "1"},
		{PropertyName: "string", Message: "2"},
		{PropertyName: "slice", Message: "1"},
		{PropertyName: "slice", Message: "2"},
		{PropertyName: "map", Message: "1"},
		{PropertyName: "map", Message: "2"},
	}
	firstErrors := []govytest.ExpectedRuleError{
		{PropertyName: "string", Message: "1"},
		{PropertyName: "slice", Message: "1"},
		{PropertyName: "map", Message: "1"},
	}

	testCases := map[string]struct {
		validator      govy.Validator[mockValidatorStruct]
		expectedErrors []govytest.ExpectedRuleError
	}{
		"mode stop": {
			validator.Cascade(govy.CascadeModeStop),
			[]govytest.ExpectedRuleError{
				{PropertyName: "string", Message: "1"},
			},
		},
		"mode continue": {
			validator.Cascade(govy.CascadeModeContinue),
			allErrors,
		},
		"no mode": {
			validator,
			allErrors,
		},
		"no mode and mode overrides": {
			govy.New(
				propertyRules.Cascade(govy.CascadeModeStop),
				propertyRulesForSlice.Cascade(govy.CascadeModeStop),
				propertyRulesForMap.Cascade(govy.CascadeModeStop),
			),
			firstErrors,
		},
		"mode continue and mode overrides": {
			govy.New(
				propertyRulesNeverFails,
				propertyRules.Cascade(govy.CascadeModeStop),
				propertyRulesForSlice.Cascade(govy.CascadeModeStop),
				propertyRulesForMap.Cascade(govy.CascadeModeStop),
			).Cascade(govy.CascadeModeContinue),
			firstErrors,
		},
		"mode stop and mode overrides": {
			govy.New(
				propertyRulesNeverFails,
				propertyRules.Cascade(govy.CascadeModeContinue),
				propertyRulesForSlice.Cascade(govy.CascadeModeContinue),
				propertyRulesForMap.Cascade(govy.CascadeModeContinue),
			).Cascade(govy.CascadeModeStop),
			[]govytest.ExpectedRuleError{
				{PropertyName: "string", Message: "1"},
				{PropertyName: "string", Message: "2"},
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.validator.Validate(mockValidatorStruct{})
			govytest.AssertError(t, err, tc.expectedErrors...)
		})
	}
}

func TestValidatorRemovePropertiesByName(t *testing.T) {
	t.Run("remove single property by name", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field").
				Rules(rules.EQ("test")),
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("other").
				Rules(rules.EQ("invalid")),
		)
		modified := v.RemovePropertiesByName("field")
		err := modified.Validate(mockValidatorStruct{Field: "invalid"})
		assert.NoError(t, err)
	})

	t.Run("remove multiple properties by name", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field1").
				Rules(rules.EQ("test")),
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field2").
				Rules(rules.EQ("test")),
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field3").
				Rules(rules.EQ("valid")),
		)
		modified := v.RemovePropertiesByName("field1", "field2")
		err := modified.Validate(mockValidatorStruct{Field: "valid"})
		assert.NoError(t, err)
	})

	t.Run("remove all properties", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field1").
				Rules(rules.EQ("test")),
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field2").
				Rules(rules.EQ("test")),
		)
		modified := v.RemovePropertiesByName("field1", "field2")
		err := modified.Validate(mockValidatorStruct{Field: "anything"})
		assert.NoError(t, err)
	})

	t.Run("remove non-existent property", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field").
				Rules(rules.EQ("test")),
		)
		modified := v.RemovePropertiesByName("nonexistent")
		err := modified.Validate(mockValidatorStruct{Field: "invalid"})
		assert.Error(t, err)
	})

	t.Run("remove with empty names slice", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field").
				Rules(rules.EQ("test")),
		)
		modified := v.RemovePropertiesByName()
		err := modified.Validate(mockValidatorStruct{Field: "invalid"})
		assert.Error(t, err)
	})

	t.Run("original validator is unchanged", func(t *testing.T) {
		original := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field").
				Rules(rules.EQ("test")),
		)
		modified := original.RemovePropertiesByName("field")

		errOriginal := original.Validate(mockValidatorStruct{Field: "invalid"})
		assert.Error(t, errOriginal)

		errModified := modified.Validate(mockValidatorStruct{Field: "invalid"})
		assert.NoError(t, errModified)
	})

	t.Run("remove slice property rules", func(t *testing.T) {
		v := govy.New(
			govy.ForSlice(func(m mockValidatorStruct) []string { return []string{m.Field} }).
				WithName("items").
				Rules(rules.SliceMaxLength[[]string](5)),
		)
		modified := v.RemovePropertiesByName("items")
		err := modified.Validate(mockValidatorStruct{Field: "test"})
		assert.NoError(t, err)
	})

	t.Run("remove map property rules", func(t *testing.T) {
		v := govy.New(
			govy.ForMap(func(m mockValidatorStruct) map[string]string {
				return map[string]string{"key": m.Field}
			}).
				WithName("mapping").
				Rules(rules.MapMaxLength[map[string]string](5)),
		)
		modified := v.RemovePropertiesByName("mapping")
		err := modified.Validate(mockValidatorStruct{Field: "test"})
		assert.NoError(t, err)
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
