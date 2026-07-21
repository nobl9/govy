package govy_test

import (
	"bytes"
	cryptorand "crypto/rand"
	"errors"
	"strconv"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govyconfig"
	"github.com/nobl9/govy/pkg/jsonpath"
	"github.com/nobl9/govy/pkg/rules"
)

func TestPropertyRules(t *testing.T) {
	type mockStruct struct {
		Field string
	}

	t.Run("no predicates, no error", func(t *testing.T) {
		r := govy.For(func(m mockStruct) string { return "path" }).
			WithPath(jsonpath.New().Name("test").Name("path")).
			Rules(govy.NewRule(func(v string) error { return nil }))
		err := r.Validate(mockStruct{})
		assert.NoError(t, err)
	})

	t.Run("no predicates, validate", func(t *testing.T) {
		expectedErr := errors.New("ops!")
		r := govy.For(func(m mockStruct) string { return "path" }).
			WithPath(jsonpath.New().Name("test").Name("path")).
			Rules(govy.NewRule(func(v string) error { return expectedErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyPath:  jsonpath.Parse("test.path"),
			PropertyValue: "path",
			Errors:        []*govy.RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("predicate matches, don't validate", func(t *testing.T) {
		r := govy.For(func(m mockStruct) string { return "value" }).
			WithPath(jsonpath.New().Name("test").Name("path")).
			When(func(mockStruct) bool { return true }).
			When(func(mockStruct) bool { return true }).
			When(func(st mockStruct) bool { return st.Field == "" }).
			Rules(govy.NewRule(func(v string) error { return errors.New("ops!") }))
		err := r.Validate(mockStruct{Field: "something"})
		assert.NoError(t, err)
	})

	t.Run("multiple rules", func(t *testing.T) {
		err1 := errors.New("oh no!")
		r := govy.For(func(m mockStruct) string { return "value" }).
			WithPath(jsonpath.New().Name("test").Name("path")).
			Rules(govy.NewRule(func(v string) error { return nil })).
			Rules(govy.NewRule(func(v string) error { return err1 })).
			Rules(govy.NewRule(func(v string) error { return nil })).
			Rules(govy.NewRule(func(v string) error {
				return govy.NewPropertyError(jsonpath.Parse("nested"), "nestedValue", &govy.RuleError{
					Message: "property is required",
					Code:    rules.ErrorCodeRequired,
				})
			}))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 2))
		assert.ElementsMatch(t, govy.PropertyErrors{
			&govy.PropertyError{
				PropertyPath:  jsonpath.Parse("test.path"),
				PropertyValue: "value",
				Errors:        []*govy.RuleError{{Message: err1.Error()}},
			},
			&govy.PropertyError{
				PropertyPath:  jsonpath.Parse("test.path.nested"),
				PropertyValue: "nestedValue",
				Errors: []*govy.RuleError{{
					Message: "property is required",
					Code:    rules.ErrorCodeRequired,
				}},
			},
		}, errs)
	})

	t.Run("cascade mode stop", func(t *testing.T) {
		expectedErr := errors.New("oh no!")
		r := govy.For(func(m mockStruct) string { return "value" }).
			WithPath(jsonpath.New().Name("test").Name("path")).
			Cascade(govy.CascadeModeStop).
			Rules(govy.NewRule(func(v string) error { return expectedErr })).
			Rules(govy.NewRule(func(v string) error { return errors.New("no") }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyPath:  jsonpath.Parse("test.path"),
			PropertyValue: "value",
			Errors:        []*govy.RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("include validator", func(t *testing.T) {
		err1 := errors.New("oh no!")
		err2 := errors.New("included")
		err3 := errors.New("included again")
		r := govy.For(func(m mockStruct) mockStruct { return m }).
			WithPath(jsonpath.New().Name("test").Name("path")).
			Rules(govy.NewRule(func(v mockStruct) error { return err1 })).
			Include(govy.New(
				govy.For(func(s mockStruct) string { return "value" }).
					WithName("included").
					Rules(govy.NewRule(func(v string) error { return err2 })).
					Rules(govy.NewRule(func(v string) error {
						return govy.NewPropertyError(jsonpath.Parse("nested"), "nestedValue", err3)
					})),
			))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 3))
		assert.ElementsMatch(t, govy.PropertyErrors{
			{
				PropertyPath: jsonpath.Parse("test.path"),
				Errors:       []*govy.RuleError{{Message: err1.Error()}},
			},
			{
				PropertyPath:  jsonpath.Parse("test.path.included"),
				PropertyValue: "value",
				Errors:        []*govy.RuleError{{Message: err2.Error()}},
			},
			{
				PropertyPath:  jsonpath.Parse("test.path.included.nested"),
				PropertyValue: "nestedValue",
				Errors:        []*govy.RuleError{{Message: err3.Error()}},
			},
		}, errs)
	})

	t.Run("get self", func(t *testing.T) {
		expectedErrs := errors.New("self error")
		r := govy.For(govy.GetSelf[mockStruct]()).
			WithPath(jsonpath.New().Name("test").Name("path")).
			Rules(govy.NewRule(func(v mockStruct) error { return expectedErrs }))
		object := mockStruct{Field: "this"}
		errs := mustPropertyErrors(t, r.Validate(object))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyPath:  jsonpath.Parse("test.path"),
			PropertyValue: internal.PropertyValueString(object),
			Errors:        []*govy.RuleError{{Message: expectedErrs.Error()}},
		}, errs[0])
	})

	t.Run("WithName escapes dots", func(t *testing.T) {
		expectedErr := errors.New("bad")
		r := govy.For(func(m mockStruct) string { return "value" }).
			WithName("foo.bar").
			Rules(govy.NewRule(func(v string) error { return expectedErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, jsonpath.Parse("['foo.bar']"), errs[0].PropertyPath)
	})

	t.Run("WithName escapes brackets and spaces", func(t *testing.T) {
		expectedErr := errors.New("bad")
		r := govy.For(func(m mockStruct) string { return "value" }).
			WithName("key [0]").
			Rules(govy.NewRule(func(v string) error { return expectedErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, jsonpath.Parse("['key [0]']"), errs[0].PropertyPath)
	})

	t.Run("WithName does not escape simple names", func(t *testing.T) {
		expectedErr := errors.New("bad")
		r := govy.For(func(m mockStruct) string { return "value" }).
			WithName("simpleName").
			Rules(govy.NewRule(func(v string) error { return expectedErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, jsonpath.Parse("simpleName"), errs[0].PropertyPath)
	})

	t.Run("hide value", func(t *testing.T) {
		expectedErr := errors.New("oh no! here's the value: 'secret'")
		r := govy.For(func(m mockStruct) string { return "secret" }).
			WithPath(jsonpath.New().Name("test").Name("path")).
			HideValue().
			Rules(govy.NewRule(func(v string) error { return expectedErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyPath:  jsonpath.Parse("test.path"),
			PropertyValue: "",
			Errors:        []*govy.RuleError{{Message: "oh no! here's the value: '[hidden]'"}},
		}, errs[0])
	})
}

func TestForPointer(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		r := govy.ForPointer(func(s *string) *string { return s }).
			Required()
		err := r.Validate(nil)
		assert.Error(t, err)
	})
	t.Run("non nil pointer", func(t *testing.T) {
		r := govy.ForPointer(func(s *string) *string { return s }).
			Required()
		s := "this string"
		err := r.Validate(&s)
		assert.NoError(t, err)
	})
}

func TestRequiredAndOmitEmpty(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		r := govy.ForPointer(func(s *string) *string { return s }).
			Rules(rules.StringMinLength(10))

		t.Run("implicit omitEmpty", func(t *testing.T) {
			err := r.Validate(nil)
			assert.NoError(t, err)
		})
		t.Run("explicit omitEmpty", func(t *testing.T) {
			err := r.OmitEmpty().Validate(nil)
			assert.NoError(t, err)
		})
		t.Run("required", func(t *testing.T) {
			errs := mustPropertyErrors(t, r.Required().Validate(nil))
			assert.Len(t, errs, 1)
			assert.True(t, govy.HasErrorCode(errs, rules.ErrorCodeRequired))
		})
	})

	t.Run("non empty pointer", func(t *testing.T) {
		r := govy.ForPointer(func(s *string) *string { return s }).
			Rules(rules.StringMinLength(10))

		t.Run("validate", func(t *testing.T) {
			errs := mustPropertyErrors(t, r.Validate(ptr("")))
			assert.Len(t, errs, 1)
			assert.True(t, govy.HasErrorCode(errs, rules.ErrorCodeStringMinLength))
		})
		t.Run("omitEmpty", func(t *testing.T) {
			errs := mustPropertyErrors(t, r.OmitEmpty().Validate(ptr("")))
			assert.Len(t, errs, 1)
			assert.True(t, govy.HasErrorCode(errs, rules.ErrorCodeStringMinLength))
		})
		t.Run("required", func(t *testing.T) {
			errs := mustPropertyErrors(t, r.Required().Validate(ptr("")))
			assert.Len(t, errs, 1)
			assert.True(t, govy.HasErrorCode(errs, rules.ErrorCodeStringMinLength))
		})
	})

	t.Run("required with when condition", func(t *testing.T) {
		r := govy.ForPointer(func(s *string) *string { return s }).
			When(func(s *string) bool { return s != nil }).
			Required().
			Rules(rules.StringMinLength(10))
		err := r.Validate(nil)
		assert.NoError(t, err)
	})
}

func TestTransform(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := govy.Transform(getter, strconv.Atoi).
			Rules(rules.GT(122))
		err := transformed.Validate("123")
		assert.NoError(t, err)
	})
	t.Run("fails validation", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := govy.Transform(getter, strconv.Atoi).
			WithName("prop").
			Rules(rules.GT(123))
		errs := mustPropertyErrors(t, transformed.Validate("123"))
		assert.Len(t, errs, 1)
		assert.True(t, govy.HasErrorCode(errs, rules.ErrorCodeGreaterThan))
	})
	t.Run("zero value with omitEmpty", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := govy.Transform(getter, strconv.Atoi).
			WithName("prop").
			OmitEmpty().
			Rules(rules.GT(123))
		err := transformed.Validate("")
		assert.NoError(t, err)
	})
	t.Run("zero value with required", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := govy.Transform(getter, strconv.Atoi).
			WithName("prop").
			Required().
			Rules(rules.GT(123))
		errs := mustPropertyErrors(t, transformed.Validate(""))
		assert.Len(t, errs, 1)
		assert.True(t, govy.HasErrorCode(errs, rules.ErrorCodeRequired))
	})
	t.Run("skip zero value", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := govy.Transform(getter, strconv.Atoi).
			WithName("prop").
			Rules(rules.GT(123))
		errs := mustPropertyErrors(t, transformed.Validate(""))
		assert.Len(t, errs, 1)
		assert.True(t, govy.HasErrorCode(errs, rules.ErrorCodeGreaterThan))
	})
	t.Run("fails transformation", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := govy.Transform(getter, strconv.Atoi).
			WithName("prop").
			Rules(rules.GT(123))
		errs := mustPropertyErrors(t, transformed.Validate("123z"))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, expectedErrorOutput(t, "property_error_transform.txt"))
		assert.True(t, govy.HasErrorCode(errs, govy.ErrorCodeTransform))
	})
	t.Run("fail transformation with hidden value", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := govy.Transform(getter, strconv.Atoi).
			WithName("prop").
			HideValue().
			Rules(rules.GT(123))
		errs := mustPropertyErrors(t, transformed.Validate("secret!"))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, expectedErrorOutput(t, "property_error_transform_with_hidden_value.txt"))
		assert.True(t, govy.HasErrorCode(errs, govy.ErrorCodeTransform))
	})
}

func TestPropertyRules_InferPath(t *testing.T) {
	govyconfig.SetInferPathIncludeTestFiles(true)
	defer govyconfig.SetInferPathIncludeTestFiles(false)

	type Age struct {
		Years int `json:"years"`
	}
	type Details struct {
		Age Age `json:"age"`
	}
	type Teacher struct {
		Name    string `json:"name,omitempty"`
		Surname string
		Details Details `json:"details"`
		Remarks *string `json:"remarks"`
	}

	t.Run("inline getter", func(t *testing.T) {
		r := govy.For(func(t Teacher) string { return t.Name }).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ("John"))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke"}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'name' with value 'Luke':\n  - must be equal to 'John'")
	})
	t.Run("selector expression getter", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) string { return t.Name }).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ("John"))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke"}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'name' with value 'Luke':\n  - must be equal to 'John'")
	})
	t.Run("nested selector expression getter", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) int { return t.Details.Age.Years }).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ(29))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Details: Details{Age: Age{Years: 30}}}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'details.age.years' with value '30':\n  - must be equal to '29'")
	})
	t.Run("variable assignment", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) string {
				teacherName := t.Name
				return teacherName
			}).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ("John"))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke"}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'name' with value 'Luke':\n  - must be equal to 'John'")
	})
	t.Run("nested selector variable assignment", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) int {
				teacherAge := t.Details.Age.Years
				return teacherAge
			}).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ(29))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Details: Details{Age: Age{Years: 30}}}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'details.age.years' with value '30':\n  - must be equal to '29'")
	})
	t.Run("external function", func(t *testing.T) {
		getter := func(t Teacher) int {
			teacherAge := t.Details.Age.Years
			return teacherAge
		}
		r := govy.
			For(getter).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ(29))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Details: Details{Age: Age{Years: 30}}}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'details.age.years' with value '30':\n  - must be equal to '29'")
	})
	t.Run("pointer", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) *string { return t.Remarks }).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ(ptr("No remarks")))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Remarks: ptr("Some remarks")}))
		assert.Len(t, errs, 1)
		assert.ErrorContains(t, errs, "- 'remarks' with value 'Some remarks':\n  - must be equal to '")
	})
	t.Run("multiple return statements, infer from top level", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) *string {
				if t.Remarks == nil {
					return nil
				}
				return t.Remarks
			}).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ(ptr("No remarks")))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Remarks: ptr("Some remarks")}))
		assert.Len(t, errs, 1)
		assert.ErrorContains(t, errs, "- 'remarks' with value 'Some remarks':\n  - must be equal to '")
	})
	t.Run("multiple return statements, infer from nested if", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) *string {
				if t.Remarks != nil {
					return t.Remarks
				}
				return nil
			}).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ(ptr("No remarks")))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Remarks: ptr("Some remarks")}))
		assert.Len(t, errs, 1)
		assert.ErrorContains(t, errs, "- 'remarks' with value 'Some remarks':\n  - must be equal to '")
	})
	t.Run("no json tag", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) string { return t.Surname }).
			InferPath(govy.InferPathModeRuntime).
			Rules(rules.EQ("Cormack"))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Surname: "Ellis"}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'Surname' with value 'Ellis':\n  - must be equal to 'Cormack'")
	})
}

type mockStruct struct {
	Field string `json:"field"`
}

func TestPropertyRulesWithID(t *testing.T) {
	t.Run("generated IDs remain stable across validation and plan generation", func(t *testing.T) {
		propertyRules := []govy.PropertyRulesInterface[mockStruct]{
			govy.For(func(m mockStruct) string { return m.Field }).
				WithName("scalar").
				Rules(govy.NewRule(func(string) error { return nil }).WithDescription("scalar rule")),
			govy.ForPointer(func(m mockStruct) *string { return &m.Field }).
				WithName("pointer").
				Rules(govy.NewRule(func(string) error { return nil }).WithDescription("pointer rule")),
			govy.Transform(
				func(m mockStruct) string { return m.Field },
				func(value string) (int, error) { return len(value), nil },
			).
				WithName("transformed").
				Rules(govy.NewRule(func(int) error { return nil }).WithDescription("transformed rule")),
			govy.ForSlice(func(m mockStruct) []string { return []string{m.Field} }).
				WithName("slice").
				Rules(govy.NewRule(func([]string) error { return nil }).WithDescription("slice rule")),
			govy.ForMap(func(m mockStruct) map[string]string {
				return map[string]string{"key": m.Field}
			}).
				WithName("map").
				Rules(govy.NewRule(func(map[string]string) error { return nil }).WithDescription("map rule")),
		}
		ids := make([]string, len(propertyRules))
		for i, property := range propertyRules {
			ids[i] = property.GetID()
			assert.True(t, ids[i] != "")
		}

		validator := govy.New(propertyRules...)
		assert.NoError(t, validator.Validate(mockStruct{Field: "value"}))
		plan, err := govy.Plan(validator)
		assert.Require(t, assert.NoError(t, err))
		assert.Len(t, plan.Properties, len(propertyRules))

		for i, property := range propertyRules {
			assert.Equal(t, ids[i], property.GetID())
		}
	})

	t.Run("scalar builders preserve custom ID", func(t *testing.T) {
		included := govy.New(
			govy.For(govy.GetSelf[string]()).
				WithName("value").
				Rules(govy.NewRule(func(string) error { return nil })),
		)
		base := govy.For(func(m mockStruct) string { return m.Field }).WithID("stable")
		derived := []struct {
			name string
			id   string
		}{
			{"WithName", base.WithName("field").GetID()},
			{"WithPath", base.WithPath(jsonpath.New().Name("field")).GetID()},
			{"WithExamples", base.WithExamples("example").GetID()},
			{"Rules", base.Rules(govy.NewRule(func(string) error { return nil })).GetID()},
			{"Include", base.Include(included).GetID()},
			{"When", base.When(func(mockStruct) bool { return true }).GetID()},
			{"Required", base.Required().GetID()},
			{"OmitEmpty", base.OmitEmpty().GetID()},
			{"HideValue", base.HideValue().GetID()},
			{"Cascade", base.Cascade(govy.CascadeModeStop).GetID()},
			{"InferPath", base.InferPath(govy.InferPathModeRuntime).GetID()},
		}

		for _, tc := range derived {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, "stable", base.GetID())
				assert.Equal(t, "stable", tc.id)
			})
		}
	})

	t.Run("slice builders preserve custom ID", func(t *testing.T) {
		wholeSlice := govy.New(
			govy.For(govy.GetSelf[[]string]()).
				WithName("slice").
				Rules(govy.NewRule(func([]string) error { return nil })),
		)
		element := govy.New(
			govy.For(govy.GetSelf[string]()).
				WithName("element").
				Rules(govy.NewRule(func(string) error { return nil })),
		)
		base := govy.ForSlice(func(m mockStruct) []string { return []string{m.Field} }).WithID("stable")
		derived := []struct {
			name string
			id   string
		}{
			{"WithName", base.WithName("items").GetID()},
			{"WithPath", base.WithPath(jsonpath.New().Name("items")).GetID()},
			{"WithExamples", base.WithExamples("example").GetID()},
			{"Rules", base.Rules(govy.NewRule(func([]string) error { return nil })).GetID()},
			{"RulesForEach", base.RulesForEach(govy.NewRule(func(string) error { return nil })).GetID()},
			{"When", base.When(func(mockStruct) bool { return true }).GetID()},
			{"Include", base.Include(wholeSlice).GetID()},
			{"IncludeForEach", base.IncludeForEach(element).GetID()},
			{"Cascade", base.Cascade(govy.CascadeModeStop).GetID()},
			{"InferPath", base.InferPath(govy.InferPathModeRuntime).GetID()},
		}

		for _, tc := range derived {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, "stable", base.GetID())
				assert.Equal(t, "stable", tc.id)
			})
		}
	})

	t.Run("map builders preserve custom ID", func(t *testing.T) {
		wholeMap := govy.New(
			govy.For(govy.GetSelf[map[string]string]()).
				WithName("map").
				Rules(govy.NewRule(func(map[string]string) error { return nil })),
		)
		stringValue := govy.New(
			govy.For(govy.GetSelf[string]()).
				WithName("value").
				Rules(govy.NewRule(func(string) error { return nil })),
		)
		item := govy.New(
			govy.For(func(value govy.MapItem[string, string]) string { return value.Value }).
				WithName("item").
				Rules(govy.NewRule(func(string) error { return nil })),
		)
		base := govy.ForMap(func(m mockStruct) map[string]string {
			return map[string]string{"key": m.Field}
		}).WithID("stable")
		derived := []struct {
			name string
			id   string
		}{
			{"WithName", base.WithName("data").GetID()},
			{"WithPath", base.WithPath(jsonpath.New().Name("data")).GetID()},
			{"WithExamples", base.WithExamples("example").GetID()},
			{"Rules", base.Rules(govy.NewRule(func(map[string]string) error { return nil })).GetID()},
			{"RulesForKeys", base.RulesForKeys(govy.NewRule(func(string) error { return nil })).GetID()},
			{"RulesForValues", base.RulesForValues(govy.NewRule(func(string) error { return nil })).GetID()},
			{
				"RulesForItems",
				base.RulesForItems(govy.NewRule(func(govy.MapItem[string, string]) error { return nil })).GetID(),
			},
			{"When", base.When(func(mockStruct) bool { return true }).GetID()},
			{"Include", base.Include(wholeMap).GetID()},
			{"IncludeForKeys", base.IncludeForKeys(stringValue).GetID()},
			{"IncludeForValues", base.IncludeForValues(stringValue).GetID()},
			{"IncludeForItems", base.IncludeForItems(item).GetID()},
			{"Cascade", base.Cascade(govy.CascadeModeStop).GetID()},
			{"InferPath", base.InferPath(govy.InferPathModeRuntime).GetID()},
		}

		for _, tc := range derived {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, "stable", base.GetID())
				assert.Equal(t, "stable", tc.id)
			})
		}
	})

	t.Run("custom ID overrides create independent copies", func(t *testing.T) {
		testCases := []struct {
			name string
			ids  func() []string
		}{
			{
				name: "scalar",
				ids: func() []string {
					base := govy.For(func(m mockStruct) string { return m.Field })
					baseID := base.GetID()
					first := base.WithID("first")
					second := first.WithID("second")
					return []string{baseID, base.GetID(), first.GetID(), second.GetID()}
				},
			},
			{
				name: "slice",
				ids: func() []string {
					base := govy.ForSlice(func(m mockStruct) []string { return []string{m.Field} })
					baseID := base.GetID()
					first := base.WithID("first")
					second := first.WithID("second")
					return []string{baseID, base.GetID(), first.GetID(), second.GetID()}
				},
			},
			{
				name: "map",
				ids: func() []string {
					base := govy.ForMap(func(m mockStruct) map[string]string {
						return map[string]string{"key": m.Field}
					})
					baseID := base.GetID()
					first := base.WithID("first")
					second := first.WithID("second")
					return []string{baseID, base.GetID(), first.GetID(), second.GetID()}
				},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				ids := tc.ids()
				assert.True(t, ids[0] != "")
				assert.Equal(t, ids[0], ids[1])
				assert.Equal(t, "first", ids[2])
				assert.Equal(t, "second", ids[3])
			})
		}
	})

	t.Run("generated ID belongs to the final derived instance", func(t *testing.T) {
		entropy := make([]byte, 16*32)
		for block := range 32 {
			entropy[block*16+15] = byte(block + 1)
		}
		originalReader := cryptorand.Reader
		cryptorand.Reader = bytes.NewReader(entropy)
		t.Cleanup(func() { cryptorand.Reader = originalReader })

		failingStringRule := govy.NewRule(func(string) error { return errors.New("string") })
		failingSliceRule := govy.NewRule(func([]string) error { return errors.New("slice") })
		failingMapRule := govy.NewRule(func(map[string]string) error { return errors.New("map") })

		scalarBase := govy.For(func(m mockStruct) string { return m.Field }).
			WithName("scalar").
			Rules(failingStringRule)
		scalarDerived := scalarBase.WithExamples("example")
		sliceBase := govy.ForSlice(func(m mockStruct) []string { return []string{m.Field} }).
			WithName("slice").
			Rules(failingSliceRule)
		sliceDerived := sliceBase.WithExamples("example")
		mapBase := govy.ForMap(func(m mockStruct) map[string]string {
			return map[string]string{"key": m.Field}
		}).
			WithName("map").
			Rules(failingMapRule)
		mapDerived := mapBase.WithExamples("example")

		testCases := []struct {
			name      string
			baseID    string
			derivedID string
			property  govy.PropertyRulesInterface[mockStruct]
		}{
			{"scalar", scalarBase.GetID(), scalarDerived.GetID(), scalarDerived},
			{"slice", sliceBase.GetID(), sliceDerived.GetID(), sliceDerived},
			{"map", mapBase.GetID(), mapDerived.GetID(), mapDerived},
		}
		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				assert.True(t, tc.baseID != tc.derivedID)
				validator := govy.New(tc.property)
				assert.Error(t, validator.Validate(mockStruct{}))
				assert.Error(t, validator.RemovePropertiesByID(tc.baseID).Validate(mockStruct{}))
				assert.NoError(t, validator.RemovePropertiesByID(tc.derivedID).Validate(mockStruct{}))
			})
		}
	})
}

func BenchmarkFor(b *testing.B) {
	for b.Loop() {
		_ = govy.For(func(m mockStruct) string { return m.Field })
	}
}

func mustPropertyErrors(t *testing.T, err error) govy.PropertyErrors {
	t.Helper()
	return mustErrorType[govy.PropertyErrors](t, err)
}

func mustErrorType[T error](t *testing.T, err error) T {
	t.Helper()
	assert.Require(t, assert.Error(t, err))
	assert.Require(t, assert.IsType[T](t, err))
	return err.(T)
}

func ptr[T any](v T) *T { return &v }
