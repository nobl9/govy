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

func mustValidatorError(t *testing.T, err error) *govy.ValidatorError {
	t.Helper()
	return mustErrorType[*govy.ValidatorError](t, err)
}

func mustValidatorErrors(t *testing.T, err error) govy.ValidatorErrors {
	t.Helper()
	return mustErrorType[govy.ValidatorErrors](t, err)
}

func TestValidatorWithID(t *testing.T) {
	t.Run("set custom ID", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return nil })),
		).WithID("custom-validator-id")

		err := v.Validate(mockValidatorStruct{})
		assert.NoError(t, err)
	})

	t.Run("WithID creates a copy", func(t *testing.T) {
		original := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return nil })),
		)
		modified := original.WithID("custom-id")

		assert.NoError(t, original.Validate(mockValidatorStruct{}))
		assert.NoError(t, modified.Validate(mockValidatorStruct{}))
	})
}

func TestValidatorRemovePropertiesByID(t *testing.T) {
	t.Run("remove single property by ID", func(t *testing.T) {
		err1 := errors.New("error1")
		err2 := errors.New("error2")
		prop1 := govy.For(func(m mockValidatorStruct) string { return "test1" }).
			WithName("property1").
			Rules(govy.NewRule(func(v string) error { return err1 }))
		prop2 := govy.For(func(m mockValidatorStruct) string { return "test2" }).
			WithName("property2").
			Rules(govy.NewRule(func(v string) error { return err2 }))
		v := govy.New(prop1, prop2)

		filteredV := v.RemovePropertiesByID(prop1.GetID())
		err := mustValidatorError(t, filteredV.Validate(mockValidatorStruct{}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, "property2", err.Errors[0].PropertyName)
	})

	t.Run("remove multiple properties by ID", func(t *testing.T) {
		err1 := errors.New("error1")
		err2 := errors.New("error2")
		err3 := errors.New("error3")
		prop1 := govy.For(func(m mockValidatorStruct) string { return "test1" }).
			WithName("property1").
			Rules(govy.NewRule(func(v string) error { return err1 }))
		prop2 := govy.For(func(m mockValidatorStruct) string { return "test2" }).
			WithName("property2").
			Rules(govy.NewRule(func(v string) error { return err2 }))
		prop3 := govy.For(func(m mockValidatorStruct) string { return "test3" }).
			WithName("property3").
			Rules(govy.NewRule(func(v string) error { return err3 }))
		v := govy.New(prop1, prop2, prop3)

		filteredV := v.RemovePropertiesByID(prop1.GetID(), prop3.GetID())
		err := mustValidatorError(t, filteredV.Validate(mockValidatorStruct{}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, "property2", err.Errors[0].PropertyName)
	})

	t.Run("remove property by generated ID", func(t *testing.T) {
		err1 := errors.New("error1")
		err2 := errors.New("error2")

		prop1 := govy.For(func(m mockValidatorStruct) string { return "test1" }).
			WithName("property1").
			Rules(govy.NewRule(func(v string) error { return err1 }))

		prop2 := govy.For(func(m mockValidatorStruct) string { return "test2" }).
			WithName("property2").
			Rules(govy.NewRule(func(v string) error { return err2 }))

		v := govy.New(prop1, prop2)

		filteredV := v.RemovePropertiesByID(prop1.GetID())
		err := mustValidatorError(t, filteredV.Validate(mockValidatorStruct{}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, "property2", err.Errors[0].PropertyName)
	})

	t.Run("remove all properties", func(t *testing.T) {
		prop1 := govy.For(func(m mockValidatorStruct) string { return "test1" }).
			WithName("property1").
			Rules(govy.NewRule(func(v string) error { return errors.New("error1") }))
		prop2 := govy.For(func(m mockValidatorStruct) string { return "test2" }).
			WithName("property2").
			Rules(govy.NewRule(func(v string) error { return errors.New("error2") }))
		v := govy.New(prop1, prop2)

		filteredV := v.RemovePropertiesByID(prop1.GetID(), prop2.GetID())
		err := filteredV.Validate(mockValidatorStruct{})

		assert.NoError(t, err)
	})

	t.Run("remove non-existent property", func(t *testing.T) {
		err1 := errors.New("error1")
		prop1 := govy.For(func(m mockValidatorStruct) string { return "test1" }).
			WithName("property1").
			Rules(govy.NewRule(func(v string) error { return err1 }))
		v := govy.New(prop1)

		filteredV := v.RemovePropertiesByID("non-existent-uuid-12345")
		err := mustValidatorError(t, filteredV.Validate(mockValidatorStruct{}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, "property1", err.Errors[0].PropertyName)
	})

	t.Run("remove with empty IDs slice", func(t *testing.T) {
		err1 := errors.New("error1")
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test1" }).
				WithName("property1").
				Rules(govy.NewRule(func(v string) error { return err1 })),
		)

		filteredV := v.RemovePropertiesByID()
		err := mustValidatorError(t, filteredV.Validate(mockValidatorStruct{}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, "property1", err.Errors[0].PropertyName)
	})

	t.Run("original validator is unchanged", func(t *testing.T) {
		err1 := errors.New("error1")
		err2 := errors.New("error2")
		prop1 := govy.For(func(m mockValidatorStruct) string { return "test1" }).
			WithName("property1").
			Rules(govy.NewRule(func(v string) error { return err1 }))
		prop2 := govy.For(func(m mockValidatorStruct) string { return "test2" }).
			WithName("property2").
			Rules(govy.NewRule(func(v string) error { return err2 }))
		v := govy.New(prop1, prop2)

		filteredV := v.RemovePropertiesByID(prop1.GetID())

		originalErr := mustValidatorError(t, v.Validate(mockValidatorStruct{}))
		assert.Len(t, originalErr.Errors, 2)

		filteredErr := mustValidatorError(t, filteredV.Validate(mockValidatorStruct{}))
		assert.Len(t, filteredErr.Errors, 1)
	})

	t.Run("remove nested validators with include", func(t *testing.T) {
		type nested struct {
			Value string
		}
		type parent struct {
			Nested nested
		}

		nestedValidator := govy.New(
			govy.For(func(n nested) string { return n.Value }).
				WithName("value").
				Rules(rules.EQ("expected")),
		).WithID("nested-validator")

		nestedProp := govy.For(func(p parent) nested { return p.Nested }).
			WithName("nested").
			Include(nestedValidator)

		parentValidator := govy.New(nestedProp)

		obj := parent{Nested: nested{Value: "wrong"}}

		err := mustValidatorError(t, parentValidator.Validate(obj))
		assert.Require(t, assert.Len(t, err.Errors, 1))

		filteredValidator := parentValidator.RemovePropertiesByID(nestedProp.GetID())
		filteredErr := filteredValidator.Validate(obj)
		assert.NoError(t, filteredErr)
	})

	t.Run("remove slice property rules", func(t *testing.T) {
		type withSlice struct {
			Items []string
		}

		sliceProp := govy.ForSlice(func(w withSlice) []string { return w.Items }).
			WithName("items").
			Rules(rules.SliceMaxLength[[]string](1))

		v := govy.New(sliceProp)

		obj := withSlice{Items: []string{"a", "b"}}

		err := mustValidatorError(t, v.Validate(obj))
		assert.Require(t, assert.Len(t, err.Errors, 1))

		filteredV := v.RemovePropertiesByID(sliceProp.GetID())
		filteredErr := filteredV.Validate(obj)
		assert.NoError(t, filteredErr)
	})

	t.Run("remove map property rules", func(t *testing.T) {
		type withMap struct {
			Data map[string]string
		}

		mapProp := govy.ForMap(func(w withMap) map[string]string { return w.Data }).
			WithName("data").
			Rules(rules.MapMaxLength[map[string]string](1))

		v := govy.New(mapProp)

		obj := withMap{Data: map[string]string{"a": "1", "b": "2"}}

		err := mustValidatorError(t, v.Validate(obj))
		assert.Require(t, assert.Len(t, err.Errors, 1))

		filteredV := v.RemovePropertiesByID(mapProp.GetID())
		filteredErr := filteredV.Validate(obj)
		assert.NoError(t, filteredErr)
	})
}

type mockValidatorStruct struct {
	Field string
}
