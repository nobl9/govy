package govy_test

import (
	"errors"
	"testing"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govyconfig"
	"github.com/nobl9/govy/pkg/govytest"
	"github.com/nobl9/govy/pkg/jsonpath"
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
				WithPath(jsonpath.New().Name("test").Name("name")).
				Rules(govy.NewRule(func(v string) error { return err1 })),
			govy.For(func(m mockValidatorStruct) string { return "display" }).
				WithPath(jsonpath.New().Name("test").Name("display")).
				Rules(govy.NewRule(func(v string) error { return err2 })),
		)
		err := mustValidatorError(t, v.Validate(mockValidatorStruct{}))
		assert.Require(t, assert.Len(t, err.Errors, 2))
		assert.Equal(t, &govy.ValidatorError{Errors: govy.PropertyErrors{
			&govy.PropertyError{
				PropertyPath:  jsonpath.Parse("test.name"),
				PropertyValue: "name",
				Errors:        []*govy.RuleError{{Message: err1.Error()}},
			},
			&govy.PropertyError{
				PropertyPath:  jsonpath.Parse("test.display"),
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
				PropertyPath:  jsonpath.Parse("test"),
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
	t.Run("custom name function", func(t *testing.T) {
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
	})
	t.Run("type based builtin name function", func(t *testing.T) {
		v := govy.New(
			govy.For(func(t Teacher) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
		).WithNameFunc(govy.NameFuncFromTypeName[Teacher]())

		err := v.Validate(Teacher{})
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, `Validation for Teacher has failed for the following properties:
  - 'test' with value 'test':
    - test`)
	})
	t.Run("setting WithName AFTTER WithNameFunc cancels name function out", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return errors.New("test") })),
		).
			WithNameFunc(govy.NameFuncFromTypeName[mockValidatorStruct]()).
			WithName("myValidator")

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
						PropertyPath: jsonpath.Parse("Field"),
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
						PropertyPath: jsonpath.Parse("Field"),
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
		{PropertyPath: "string", Message: "1"},
		{PropertyPath: "string", Message: "2"},
		{PropertyPath: "slice", Message: "1"},
		{PropertyPath: "slice", Message: "2"},
		{PropertyPath: "map", Message: "1"},
		{PropertyPath: "map", Message: "2"},
	}
	firstErrors := []govytest.ExpectedRuleError{
		{PropertyPath: "string", Message: "1"},
		{PropertyPath: "slice", Message: "1"},
		{PropertyPath: "map", Message: "1"},
	}

	testCases := map[string]struct {
		validator      govy.Validator[mockValidatorStruct]
		expectedErrors []govytest.ExpectedRuleError
	}{
		"mode stop": {
			validator.Cascade(govy.CascadeModeStop),
			[]govytest.ExpectedRuleError{
				{PropertyPath: "string", Message: "1"},
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
				{PropertyPath: "string", Message: "1"},
				{PropertyPath: "string", Message: "2"},
			},
		},
		"mixed inheritance": {
			govy.New(
				propertyRulesNeverFails,
				propertyRules.Cascade(govy.CascadeModeContinue),
				propertyRulesForSlice,
				propertyRulesForMap.Cascade(govy.CascadeModeContinue),
			).Cascade(govy.CascadeModeStop),
			[]govytest.ExpectedRuleError{
				{PropertyPath: "string", Message: "1"},
				{PropertyPath: "string", Message: "2"},
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

func TestValidatorRemovePropertiesByPath(t *testing.T) {
	t.Run("remove single property", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field").
				Rules(rules.EQ("test")),
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("other").
				Rules(rules.EQ("invalid")),
		)
		err := v.Validate(mockValidatorStruct{Field: "invalid"})
		assert.Error(t, err)

		modified := v.RemovePropertiesByPath(jsonpath.New().Name("field"))
		err = modified.Validate(mockValidatorStruct{Field: "invalid"})
		assert.NoError(t, err)
	})

	t.Run("remove multiple properties", func(t *testing.T) {
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
		err := v.Validate(mockValidatorStruct{Field: "valid"})
		assert.Error(t, err)

		modified := v.RemovePropertiesByPath(
			jsonpath.New().Name("field1"),
			jsonpath.New().Name("field2"),
		)
		err = modified.Validate(mockValidatorStruct{Field: "valid"})
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
		err := v.Validate(mockValidatorStruct{Field: "anything"})
		assert.Error(t, err)

		modified := v.RemovePropertiesByPath(
			jsonpath.New().Name("field1"),
			jsonpath.New().Name("field2"),
		)
		err = modified.Validate(mockValidatorStruct{Field: "anything"})
		assert.NoError(t, err)
	})

	t.Run("remove non-existent property", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field").
				Rules(rules.EQ("test")),
		)
		modified := v.RemovePropertiesByPath(jsonpath.New().Name("nonexistent"))
		err := modified.Validate(mockValidatorStruct{Field: "invalid"})
		assert.Error(t, err)
	})

	t.Run("remove with empty paths slice", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field").
				Rules(rules.EQ("test")),
		)
		modified := v.RemovePropertiesByPath()
		err := modified.Validate(mockValidatorStruct{Field: "invalid"})
		assert.Error(t, err)
	})

	t.Run("original validator is unchanged", func(t *testing.T) {
		original := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("field").
				Rules(rules.EQ("test")),
		)
		modified := original.RemovePropertiesByPath(jsonpath.New().Name("field"))

		errOriginal := original.Validate(mockValidatorStruct{Field: "invalid"})
		assert.Error(t, errOriginal)

		errModified := modified.Validate(mockValidatorStruct{Field: "invalid"})
		assert.NoError(t, errModified)
	})

	t.Run("remove slice property rules", func(t *testing.T) {
		v := govy.New(
			govy.ForSlice(func(m mockValidatorStruct) []string { return []string{m.Field} }).
				WithName("items").
				Rules(rules.SliceMaxLength[[]string](0)),
		)
		err := v.Validate(mockValidatorStruct{Field: "test"})
		assert.Error(t, err)

		modified := v.RemovePropertiesByPath(jsonpath.New().Name("items"))
		err = modified.Validate(mockValidatorStruct{Field: "test"})
		assert.NoError(t, err)
	})

	t.Run("remove map property rules", func(t *testing.T) {
		v := govy.New(
			govy.ForMap(func(m mockValidatorStruct) map[string]string {
				return map[string]string{"key": m.Field}
			}).
				WithName("mapping").
				Rules(rules.MapMaxLength[map[string]string](0)),
		)
		err := v.Validate(mockValidatorStruct{Field: "test"})
		assert.Error(t, err)

		modified := v.RemovePropertiesByPath(jsonpath.New().Name("mapping"))
		err = modified.Validate(mockValidatorStruct{Field: "test"})
		assert.NoError(t, err)
	})

	t.Run("remove multi-segment path", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithPath(jsonpath.New().Name("test").Name("field")).
				Rules(rules.EQ("expected")),
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("other").
				Rules(rules.EQ("expected")),
		)
		err := v.Validate(mockValidatorStruct{Field: "actual"})
		assert.Error(t, err)

		modified := v.RemovePropertiesByPath(jsonpath.New().Name("test").Name("field"))
		vErr := mustValidatorError(t, modified.Validate(mockValidatorStruct{Field: "actual"}))
		assert.Require(t, assert.Len(t, vErr.Errors, 1))
		assert.Equal(t, jsonpath.Parse("other"), vErr.Errors[0].PropertyPath)
	})

	t.Run("remove property with special characters in name", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("foo.bar").
				Rules(rules.EQ("expected")),
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("other").
				Rules(rules.EQ("expected")),
		)
		err := v.Validate(mockValidatorStruct{Field: "actual"})
		assert.Error(t, err)

		modified := v.RemovePropertiesByPath(jsonpath.New().Name("foo.bar"))
		vErr := mustValidatorError(t, modified.Validate(mockValidatorStruct{Field: "actual"}))
		assert.Require(t, assert.Len(t, vErr.Errors, 1))
		assert.Equal(t, jsonpath.Parse("other"), vErr.Errors[0].PropertyPath)
	})
}

func TestValidatorInferPath(t *testing.T) {
	govyconfig.SetInferPathIncludeTestFiles(true)
	defer govyconfig.SetInferPathIncludeTestFiles(false)

	type mockInferPathStruct struct {
		Name     string         `json:"name"`
		Students []string       `json:"students"`
		Grades   map[string]int `json:"grades"`
	}

	propertyRules := govy.For(func(m mockInferPathStruct) string { return m.Name }).
		Rules(rules.EQ("expected"))
	propertyRulesForSlice := govy.ForSlice(func(m mockInferPathStruct) []string { return m.Students }).
		RulesForEach(rules.EQ("expected"))
	propertyRulesForMap := govy.ForMap(func(m mockInferPathStruct) map[string]int { return m.Grades }).
		RulesForKeys(rules.EQ("expected"))

	t.Run("propagates mode to all properties", func(t *testing.T) {
		v := govy.New(
			propertyRules,
			propertyRulesForSlice,
			propertyRulesForMap,
		).
			InferPath(govy.InferPathModeRuntime)

		err := v.Validate(mockInferPathStruct{
			Name:     "actual",
			Students: []string{"actual"},
			Grades:   map[string]int{"actual": 1},
		})
		errs := mustValidatorError(t, err)
		govytest.AssertError(
			t,
			errs,
			govytest.ExpectedRuleError{
				PropertyPath: "name",
				Message:      "must be equal to 'expected'",
			},
			govytest.ExpectedRuleError{
				PropertyPath: "students[0]",
				Message:      "must be equal to 'expected'",
			},
			govytest.ExpectedRuleError{
				PropertyPath: "grades.actual",
				IsKeyError:   true,
				Message:      "must be equal to 'expected'",
			},
		)
	})

	t.Run("property mode Runtime is not overridden by validator mode Generate", func(t *testing.T) {
		govyconfig.SetInferredPath(govyconfig.InferredPath{
			// Changed name to differentiate between runtime and generate modes.
			Path: jsonpath.New().Name("generated-students"),
			File: "pkg/govy/validator_test.go",
			Line: 551,
		})

		v := govy.New(
			govy.For(func(m mockInferPathStruct) string { return m.Name }).
				InferPath(govy.InferPathModeDisable).
				Rules(rules.EQ("expected")),
			govy.ForSlice(func(m mockInferPathStruct) []string { return m.Students }).
				InferPath(govy.InferPathModeRuntime).
				RulesForEach(rules.EQ("expected")),
			govy.ForMap(func(m mockInferPathStruct) map[string]int { return m.Grades }).
				// InferPath(govy.InferPathModeRuntime). -- Inference mode should be inherited from Validator.
				RulesForKeys(rules.EQ("expected")),
		).
			InferPath(govy.InferPathModeGenerate)

		err := v.Validate(mockInferPathStruct{
			Name:     "actual",
			Students: []string{"actual"},
			Grades:   map[string]int{"actual": 1},
		})
		errs := mustValidatorError(t, err)
		govytest.AssertError(
			t,
			errs,
			govytest.ExpectedRuleError{
				PropertyPath: "", // Path inference is disabled.
				Message:      "must be equal to 'expected'",
			},
			govytest.ExpectedRuleError{
				PropertyPath: "students[0]", // Runtime mode.
				Message:      "must be equal to 'expected'",
			},
			govytest.ExpectedRuleError{
				PropertyPath: "['generated-students'].actual", // Generate mode.
				IsKeyError:   true,
				Message:      "must be equal to 'expected'",
			},
		)
	})

	t.Run("no mode set on properties - validator can set mode", func(t *testing.T) {
		// Properties don't set any mode (defaults to InferPathModeDisable = 0).
		// Validator.InferPath(InferPathModeRuntime) should propagate to all properties.
		v := govy.New(
			govy.For(func(m mockInferPathStruct) string { return m.Name }).
				Rules(rules.EQ("expected")),
			govy.ForSlice(func(m mockInferPathStruct) []string { return m.Students }).
				RulesForEach(rules.EQ("expected")),
			govy.ForMap(func(m mockInferPathStruct) map[string]int { return m.Grades }).
				RulesForKeys(rules.EQ("expected")),
		).InferPath(govy.InferPathModeRuntime)

		err := v.Validate(mockInferPathStruct{
			Name:     "actual",
			Students: []string{"actual"},
			Grades:   map[string]int{"actual": 1},
		})
		errs := mustValidatorError(t, err)
		// Check that all properties have inferred names (validator's mode was applied).
		govytest.AssertError(
			t,
			errs,
			govytest.ExpectedRuleError{PropertyPath: "name", Message: "must be equal to 'expected'"},
			govytest.ExpectedRuleError{PropertyPath: "students[0]", Message: "must be equal to 'expected'"},
			govytest.ExpectedRuleError{
				PropertyPath: "grades.actual",
				IsKeyError:   true,
				Message:      "must be equal to 'expected'",
			},
		)
	})

	t.Run("preserves property IDs", func(t *testing.T) {
		v := govy.New(
			propertyRules,
			propertyRulesForSlice,
			propertyRulesForMap,
		).
			InferPath(govy.InferPathModeRuntime).
			RemovePropertiesByID(
				propertyRules.GetID(),
				propertyRulesForSlice.GetID(),
				propertyRulesForMap.GetID(),
			)

		assert.NoError(t, v.Validate(mockInferPathStruct{}))
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
		assert.Equal(t, jsonpath.Parse("property2"), err.Errors[0].PropertyPath)
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
		assert.Equal(t, jsonpath.Parse("property2"), err.Errors[0].PropertyPath)
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
		assert.Equal(t, jsonpath.Parse("property2"), err.Errors[0].PropertyPath)
	})

	t.Run("remove property by user-supplied ID", func(t *testing.T) {
		err1 := errors.New("error1")
		err2 := errors.New("error2")

		prop1 := govy.For(func(m mockValidatorStruct) string { return "test1" }).
			WithName("property1").
			WithID("property-1").
			Rules(govy.NewRule(func(v string) error { return err1 }))

		prop2 := govy.For(func(m mockValidatorStruct) string { return "test2" }).
			WithName("property2").
			WithID("property-2").
			Rules(govy.NewRule(func(v string) error { return err2 }))

		v := govy.New(prop1, prop2)

		filteredV := v.RemovePropertiesByID("property-1")
		err := mustValidatorError(t, filteredV.Validate(mockValidatorStruct{}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, jsonpath.Parse("property2"), err.Errors[0].PropertyPath)
	})

	t.Run("remove derived property by generated ID", func(t *testing.T) {
		err1 := errors.New("error1")
		err2 := errors.New("error2")

		base := govy.For(func(m mockValidatorStruct) string { return m.Field })
		prop1 := base.WithName("property1").
			Rules(govy.NewRule(func(v string) error { return err1 }))
		prop2 := base.WithName("property2").
			Rules(govy.NewRule(func(v string) error { return err2 }))

		assert.True(t, prop1.GetID() != prop2.GetID())

		v := govy.New(prop1, prop2)

		filteredV := v.RemovePropertiesByID(prop1.GetID())
		err := mustValidatorError(t, filteredV.Validate(mockValidatorStruct{}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, jsonpath.Parse("property2"), err.Errors[0].PropertyPath)
	})

	t.Run("remove pointer property by generated ID", func(t *testing.T) {
		type withPointers struct {
			Primary   *string
			Secondary *string
		}

		prop1 := govy.ForPointer(func(w withPointers) *string { return w.Primary }).
			WithName("primary").
			Rules(rules.StringMinLength(3))
		prop2 := govy.ForPointer(func(w withPointers) *string { return w.Secondary }).
			WithName("secondary").
			Rules(rules.StringMinLength(3))
		v := govy.New(prop1, prop2)

		filteredV := v.RemovePropertiesByID(prop1.GetID())
		err := mustValidatorError(t, filteredV.Validate(withPointers{
			Primary:   ptr("a"),
			Secondary: ptr("b"),
		}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, jsonpath.Parse("secondary"), err.Errors[0].PropertyPath)
	})

	t.Run("remove transformed property by generated ID", func(t *testing.T) {
		type withTransformed struct {
			Primary   string
			Secondary string
		}
		transform := func(s string) (int, error) { return len(s), nil }

		prop1 := govy.Transform(func(w withTransformed) string { return w.Primary }, transform).
			WithName("primary").
			Rules(rules.GT(1))
		prop2 := govy.Transform(func(w withTransformed) string { return w.Secondary }, transform).
			WithName("secondary").
			Rules(rules.GT(1))
		v := govy.New(prop1, prop2)

		filteredV := v.RemovePropertiesByID(prop1.GetID())
		err := mustValidatorError(t, filteredV.Validate(withTransformed{
			Primary:   "a",
			Secondary: "b",
		}))

		assert.Require(t, assert.Len(t, err.Errors, 1))
		assert.Equal(t, jsonpath.Parse("secondary"), err.Errors[0].PropertyPath)
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
		assert.Equal(t, jsonpath.Parse("property1"), err.Errors[0].PropertyPath)
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
		assert.Equal(t, jsonpath.Parse("property1"), err.Errors[0].PropertyPath)
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
		)

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

func TestValidatorRemovePropertiesByIDDeterminism(t *testing.T) {
	var calls []string
	newProperty := func(id, path string) govy.PropertyRules[string, mockValidatorStruct] {
		return govy.For(func(mockValidatorStruct) string {
			calls = append(calls, "get:"+id)
			return id
		}).
			WithName(path).
			WithID(id).
			Rules(govy.NewRule(func(string) error {
				calls = append(calls, "rule:"+id)
				return errors.New(id)
			}))
	}

	base := govy.New(
		newProperty("a", "a"),
		newProperty("b", "b"),
		newProperty("c", "c"),
	)
	v1 := base.RemovePropertiesByID("b")
	v2 := v1.RemovePropertiesByID("a")
	sibling := base.RemovePropertiesByID("c")

	immutableCases := []struct {
		name          string
		validator     govy.Validator[mockValidatorStruct]
		expectedPaths []string
		expectedCalls []string
	}{
		{"base", base, []string{"a", "b", "c"}, []string{"get:a", "rule:a", "get:b", "rule:b", "get:c", "rule:c"}},
		{"first derivation", v1, []string{"a", "c"}, []string{"get:a", "rule:a", "get:c", "rule:c"}},
		{"second derivation", v2, []string{"c"}, []string{"get:c", "rule:c"}},
		{"sibling derivation", sibling, []string{"a", "b"}, []string{"get:a", "rule:a", "get:b", "rule:b"}},
	}
	for _, tc := range immutableCases {
		t.Run(tc.name, func(t *testing.T) {
			calls = nil
			assert.Equal(t, tc.expectedPaths, validatorErrorPaths(t, tc.validator, mockValidatorStruct{}))
			assert.Equal(t, tc.expectedCalls, calls)
		})
	}

	idempotentCases := []struct {
		name      string
		validator govy.Validator[mockValidatorStruct]
	}{
		{"single removal", base.RemovePropertiesByID("b")},
		{"duplicate arguments", base.RemovePropertiesByID("b", "b")},
		{"repeated calls", base.RemovePropertiesByID("b").RemovePropertiesByID("b")},
		{"missing IDs", base.RemovePropertiesByID("missing", "b", "missing")},
	}
	for _, tc := range idempotentCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, []string{"a", "c"}, validatorErrorPaths(t, tc.validator, mockValidatorStruct{}))
		})
	}

	t.Run("argument order does not affect retained properties", func(t *testing.T) {
		assert.Equal(
			t,
			[]string{"b"},
			validatorErrorPaths(t, base.RemovePropertiesByID("a", "c"), mockValidatorStruct{}),
		)
		assert.Equal(
			t,
			[]string{"b"},
			validatorErrorPaths(t, base.RemovePropertiesByID("c", "a"), mockValidatorStruct{}),
		)
	})

	t.Run("missing and absent arguments preserve the validator", func(t *testing.T) {
		assert.Equal(
			t,
			[]string{"a", "b", "c"},
			validatorErrorPaths(t, base.RemovePropertiesByID("missing"), mockValidatorStruct{}),
		)
		assert.Equal(
			t,
			[]string{"a", "b", "c"},
			validatorErrorPaths(t, base.RemovePropertiesByID(), mockValidatorStruct{}),
		)
	})

	t.Run("IDs are case-sensitive opaque strings", func(t *testing.T) {
		validator := govy.New(
			newProperty("CaseSensitive", "upper"),
			newProperty("casesensitive", "lower"),
			newProperty("tenant/β", "unicode"),
		)

		assert.Equal(
			t,
			[]string{"upper", "unicode"},
			validatorErrorPaths(t, validator.RemovePropertiesByID("casesensitive"), mockValidatorStruct{}),
		)
		assert.Equal(
			t,
			[]string{"upper", "lower"},
			validatorErrorPaths(t, validator.RemovePropertiesByID("tenant/β"), mockValidatorStruct{}),
		)
		assert.Equal(
			t,
			[]string{"upper", "lower", "unicode"},
			validatorErrorPaths(t, validator.RemovePropertiesByID("Casesensitive", "tenant/Β"), mockValidatorStruct{}),
		)
	})
}

func TestValidatorBuildersPreservePropertyIDs(t *testing.T) {
	type fixture struct {
		Scalar string
		Slice  []string
		Map    map[string]string
	}
	type validatorFactory func(customID, explicitMode string) (govy.Validator[fixture], string)

	failingScalarRule := govy.NewRule(func(string) error { return errors.New("scalar") })
	failingSliceRule := govy.NewRule(func([]string) error { return errors.New("slice") })
	failingMapRule := govy.NewRule(func(map[string]string) error { return errors.New("map") })
	factories := []struct {
		name string
		new  validatorFactory
	}{
		{
			name: "scalar",
			new: func(customID, explicitMode string) (govy.Validator[fixture], string) {
				property := govy.For(func(value fixture) string { return value.Scalar }).
					WithName("scalar").
					Rules(failingScalarRule)
				if customID != "" {
					property = property.WithID(customID)
				}
				switch explicitMode {
				case "cascade":
					property = property.Cascade(govy.CascadeModeStop)
				case "infer-path":
					property = property.InferPath(govy.InferPathModeDisable)
				}
				return govy.New(property), property.GetID()
			},
		},
		{
			name: "slice",
			new: func(customID, explicitMode string) (govy.Validator[fixture], string) {
				property := govy.ForSlice(func(value fixture) []string { return value.Slice }).
					WithName("slice").
					Rules(failingSliceRule)
				if customID != "" {
					property = property.WithID(customID)
				}
				switch explicitMode {
				case "cascade":
					property = property.Cascade(govy.CascadeModeStop)
				case "infer-path":
					property = property.InferPath(govy.InferPathModeDisable)
				}
				return govy.New(property), property.GetID()
			},
		},
		{
			name: "map",
			new: func(customID, explicitMode string) (govy.Validator[fixture], string) {
				property := govy.ForMap(func(value fixture) map[string]string { return value.Map }).
					WithName("map").
					Rules(failingMapRule)
				if customID != "" {
					property = property.WithID(customID)
				}
				switch explicitMode {
				case "cascade":
					property = property.Cascade(govy.CascadeModeStop)
				case "infer-path":
					property = property.InferPath(govy.InferPathModeDisable)
				}
				return govy.New(property), property.GetID()
			},
		},
	}
	modifiers := []struct {
		name         string
		explicitMode string
		customOnly   bool
		apply        func(govy.Validator[fixture]) govy.Validator[fixture]
	}{
		{"Cascade inherited", "", false, func(v govy.Validator[fixture]) govy.Validator[fixture] {
			return v.Cascade(govy.CascadeModeContinue)
		}},
		{"Cascade property override", "cascade", false, func(v govy.Validator[fixture]) govy.Validator[fixture] {
			return v.Cascade(govy.CascadeModeContinue)
		}},
		{"InferPath inherited custom ID", "", true, func(v govy.Validator[fixture]) govy.Validator[fixture] {
			return v.InferPath(govy.InferPathModeRuntime)
		}},
		{"InferPath property override", "infer-path", false, func(v govy.Validator[fixture]) govy.Validator[fixture] {
			return v.InferPath(govy.InferPathModeRuntime)
		}},
	}
	identityKinds := []struct {
		name string
		id   string
	}{
		{"generated ID", ""},
		{"custom ID", "stable"},
	}
	value := fixture{
		Scalar: "value",
		Slice:  []string{"value"},
		Map:    map[string]string{"key": "value"},
	}

	for _, factory := range factories {
		for _, modifier := range modifiers {
			for _, identity := range identityKinds {
				if modifier.customOnly && identity.id == "" {
					continue
				}
				t.Run(factory.name+"/"+modifier.name+"/"+identity.name, func(t *testing.T) {
					validator, id := factory.new(identity.id, modifier.explicitMode)
					assert.Error(t, validator.Validate(value))

					derived := modifier.apply(validator)
					assert.Error(t, derived.Validate(value))
					assert.NoError(t, derived.RemovePropertiesByID(id).Validate(value))
					assert.Error(t, validator.Validate(value))
				})
			}
		}
	}
}

func TestValidatorRemovePropertiesByIDPlanAndInferredPath(t *testing.T) {
	govyconfig.SetInferPathIncludeTestFiles(true)
	defer govyconfig.SetInferPathIncludeTestFiles(false)

	type fixture struct {
		A string `json:"a"`
		B string `json:"b"`
	}
	a := govy.For(func(value fixture) string { return value.A }).
		WithID("a").
		InferPath(govy.InferPathModeRuntime).
		Rules(govy.NewRule(func(string) error { return errors.New("a") }).WithDescription("a rule"))
	b := govy.For(func(value fixture) string { return value.B }).
		WithID("b").
		InferPath(govy.InferPathModeRuntime).
		Rules(govy.NewRule(func(string) error { return errors.New("b") }).WithDescription("b rule"))
	aID := a.GetID()
	bID := b.GetID()
	base := govy.New(a, b)

	assert.Equal(t, []string{"a", "b"}, validatorErrorPaths(t, base, fixture{}))
	assert.Equal(t, []string{"$.a", "$.b"}, validatorPlanPaths(t, base))
	assert.Equal(t, aID, a.GetID())
	assert.Equal(t, bID, b.GetID())
	basePlan, err := govy.Plan(base)
	assert.Require(t, assert.NoError(t, err))
	assert.Require(t, assert.Len(t, basePlan.Properties, 2))

	filtered := base.RemovePropertiesByID(aID)
	assert.Equal(t, []string{"b"}, validatorErrorPaths(t, filtered, fixture{}))
	assert.Equal(t, []string{"$.b"}, validatorPlanPaths(t, filtered))
	assert.Equal(t, []string{"$.a", "$.b"}, validatorPlanPaths(t, base))
	filteredPlan, err := govy.Plan(filtered)
	assert.Require(t, assert.NoError(t, err))
	assert.Require(t, assert.Len(t, filteredPlan.Properties, 1))
	assert.Equal(t, basePlan.Properties[1], filteredPlan.Properties[0])
	regeneratedBasePlan, err := govy.Plan(base)
	assert.Require(t, assert.NoError(t, err))
	assert.Equal(t, basePlan, regeneratedBasePlan)

	onlyA := govy.New(a)
	assert.Equal(t, []string{"a"}, validatorErrorPaths(t, onlyA, fixture{}))
	assert.Equal(t, []string{"$.a"}, validatorPlanPaths(t, onlyA))
	removed := onlyA.RemovePropertiesByID(aID)
	assert.NoError(t, removed.Validate(fixture{}))
	assert.Len(t, validatorPlanPaths(t, removed), 0)
}

func TestValidatorRemovePropertiesByIDIncludedValidatorBoundary(t *testing.T) {
	const (
		childID = "child-property"
		outerID = "outer-property"
	)

	t.Run("Include", func(t *testing.T) {
		type child struct{ Value string }
		type parent struct{ Child child }
		childProperty := govy.For(func(value child) string { return value.Value }).
			WithName("value").
			WithID(childID).
			Rules(rules.EQ("expected"))
		childValidator := govy.New(childProperty)
		newParent := func(included govy.Validator[child]) govy.Validator[parent] {
			return govy.New(
				govy.For(func(value parent) child { return value.Child }).
					WithName("child").
					WithID(outerID).
					Include(included),
			)
		}

		assertIncludedValidatorRemovalIsLocal(
			t,
			newParent(childValidator),
			newParent(childValidator.RemovePropertiesByID(childID)),
			parent{Child: child{Value: "actual"}},
			childID,
			outerID,
		)
	})

	t.Run("IncludeForEach", func(t *testing.T) {
		type child struct{ Value string }
		type parent struct{ Children []child }
		childProperty := govy.For(func(value child) string { return value.Value }).
			WithName("value").
			WithID(childID).
			Rules(rules.EQ("expected"))
		childValidator := govy.New(childProperty)
		newParent := func(included govy.Validator[child]) govy.Validator[parent] {
			return govy.New(
				govy.ForSlice(func(value parent) []child { return value.Children }).
					WithName("children").
					WithID(outerID).
					IncludeForEach(included),
			)
		}

		assertIncludedValidatorRemovalIsLocal(
			t,
			newParent(childValidator),
			newParent(childValidator.RemovePropertiesByID(childID)),
			parent{Children: []child{{Value: "actual"}}},
			childID,
			outerID,
		)
	})

	t.Run("IncludeForKeys", func(t *testing.T) {
		type parent struct{ Values map[string]string }
		childProperty := govy.For(govy.GetSelf[string]()).
			WithName("key").
			WithID(childID).
			Rules(rules.EQ("expected"))
		childValidator := govy.New(childProperty)
		newParent := func(child govy.Validator[string]) govy.Validator[parent] {
			return govy.New(
				govy.ForMap(func(value parent) map[string]string { return value.Values }).
					WithName("values").
					WithID(outerID).
					IncludeForKeys(child),
			)
		}

		assertIncludedValidatorRemovalIsLocal(
			t,
			newParent(childValidator),
			newParent(childValidator.RemovePropertiesByID(childID)),
			parent{Values: map[string]string{"actual": "value"}},
			childID,
			outerID,
		)
	})

	t.Run("IncludeForValues", func(t *testing.T) {
		type parent struct{ Values map[string]string }
		childProperty := govy.For(govy.GetSelf[string]()).
			WithName("value").
			WithID(childID).
			Rules(rules.EQ("expected"))
		childValidator := govy.New(childProperty)
		newParent := func(child govy.Validator[string]) govy.Validator[parent] {
			return govy.New(
				govy.ForMap(func(value parent) map[string]string { return value.Values }).
					WithName("values").
					WithID(outerID).
					IncludeForValues(child),
			)
		}

		assertIncludedValidatorRemovalIsLocal(
			t,
			newParent(childValidator),
			newParent(childValidator.RemovePropertiesByID(childID)),
			parent{Values: map[string]string{"key": "actual"}},
			childID,
			outerID,
		)
	})

	t.Run("IncludeForItems", func(t *testing.T) {
		type parent struct{ Values map[string]string }
		childProperty := govy.For(func(value govy.MapItem[string, string]) string { return value.Value }).
			WithName("item").
			WithID(childID).
			Rules(rules.EQ("expected"))
		childValidator := govy.New(childProperty)
		newParent := func(child govy.Validator[govy.MapItem[string, string]]) govy.Validator[parent] {
			return govy.New(
				govy.ForMap(func(value parent) map[string]string { return value.Values }).
					WithName("values").
					WithID(outerID).
					IncludeForItems(child),
			)
		}

		assertIncludedValidatorRemovalIsLocal(
			t,
			newParent(childValidator),
			newParent(childValidator.RemovePropertiesByID(childID)),
			parent{Values: map[string]string{"key": "actual"}},
			childID,
			outerID,
		)
	})
}

func validatorErrorPaths[T any](t *testing.T, validator govy.Validator[T], value T) []string {
	t.Helper()
	err := mustValidatorError(t, validator.Validate(value))
	paths := make([]string, len(err.Errors))
	for i, propertyError := range err.Errors {
		paths[i] = propertyError.PropertyPath.String()
	}
	return paths
}

func validatorPlanPaths[T any](t *testing.T, validator govy.Validator[T]) []string {
	t.Helper()
	plan, err := govy.Plan(validator)
	assert.Require(t, assert.NoError(t, err))
	paths := make([]string, len(plan.Properties))
	for i, property := range plan.Properties {
		paths[i] = property.Path.String()
	}
	return paths
}

func assertIncludedValidatorRemovalIsLocal[T any](
	t *testing.T,
	parent govy.Validator[T],
	parentWithFilteredChild govy.Validator[T],
	value T,
	childID string,
	outerID string,
) {
	t.Helper()
	original := mustValidatorError(t, parent.Validate(value))
	childRemoval := mustValidatorError(t, parent.RemovePropertiesByID(childID).Validate(value))
	assert.Equal(t, original, childRemoval)
	assert.NoError(t, parent.RemovePropertiesByID(outerID).Validate(value))
	assert.NoError(t, parentWithFilteredChild.Validate(value))
	assert.Equal(t, original, mustValidatorError(t, parent.Validate(value)))
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
