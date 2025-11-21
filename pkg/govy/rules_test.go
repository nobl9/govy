package govy_test

import (
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
			WithPath(jsonpath.PropertySegment{Name: "test"}, jsonpath.PropertySegment{Name: "path"}).
			Rules(govy.NewRule(func(v string) error { return nil }))
		err := r.Validate(mockStruct{})
		assert.NoError(t, err)
	})

	t.Run("no predicates, validate", func(t *testing.T) {
		expectedErr := errors.New("ops!")
		r := govy.For(func(m mockStruct) string { return "path" }).
			WithPath(jsonpath.PropertySegment{Name: "test"}, jsonpath.PropertySegment{Name: "path"}).
			Rules(govy.NewRule(func(v string) error { return expectedErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyName:  "test.path",
			PropertyValue: "path",
			Errors:        []*govy.RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("predicate matches, don't validate", func(t *testing.T) {
		r := govy.For(func(m mockStruct) string { return "value" }).
			WithPath(jsonpath.PropertySegment{Name: "test"}, jsonpath.PropertySegment{Name: "path"}).
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
			WithPath(jsonpath.PropertySegment{Name: "test"}, jsonpath.PropertySegment{Name: "path"}).
			Rules(govy.NewRule(func(v string) error { return nil })).
			Rules(govy.NewRule(func(v string) error { return err1 })).
			Rules(govy.NewRule(func(v string) error { return nil })).
			Rules(govy.NewRule(func(v string) error {
				return govy.NewPropertyError("nested", "nestedValue", &govy.RuleError{
					Message: "property is required",
					Code:    rules.ErrorCodeRequired,
				})
			}))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 2))
		assert.ElementsMatch(t, govy.PropertyErrors{
			&govy.PropertyError{
				PropertyName:  "test.path",
				PropertyValue: "value",
				Errors:        []*govy.RuleError{{Message: err1.Error()}},
			},
			&govy.PropertyError{
				PropertyName:  "test.path.nested",
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
			WithPath(jsonpath.PropertySegment{Name: "test"}, jsonpath.PropertySegment{Name: "path"}).
			Cascade(govy.CascadeModeStop).
			Rules(govy.NewRule(func(v string) error { return expectedErr })).
			Rules(govy.NewRule(func(v string) error { return errors.New("no") }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyName:  "test.path",
			PropertyValue: "value",
			Errors:        []*govy.RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("include validator", func(t *testing.T) {
		err1 := errors.New("oh no!")
		err2 := errors.New("included")
		err3 := errors.New("included again")
		r := govy.For(func(m mockStruct) mockStruct { return m }).
			WithPath(jsonpath.PropertySegment{Name: "test"}, jsonpath.PropertySegment{Name: "path"}).
			Rules(govy.NewRule(func(v mockStruct) error { return err1 })).
			Include(govy.New(
				govy.For(func(s mockStruct) string { return "value" }).
					WithName("included").
					Rules(govy.NewRule(func(v string) error { return err2 })).
					Rules(govy.NewRule(func(v string) error {
						return govy.NewPropertyError("nested", "nestedValue", err3)
					})),
			))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 3))
		assert.ElementsMatch(t, govy.PropertyErrors{
			{
				PropertyName: "test.path",
				Errors:       []*govy.RuleError{{Message: err1.Error()}},
			},
			{
				PropertyName:  "test.path.included",
				PropertyValue: "value",
				Errors:        []*govy.RuleError{{Message: err2.Error()}},
			},
			{
				PropertyName:  "test.path.included.nested",
				PropertyValue: "nestedValue",
				Errors:        []*govy.RuleError{{Message: err3.Error()}},
			},
		}, errs)
	})

	t.Run("get self", func(t *testing.T) {
		expectedErrs := errors.New("self error")
		r := govy.For(govy.GetSelf[mockStruct]()).
			WithPath(jsonpath.PropertySegment{Name: "test"}, jsonpath.PropertySegment{Name: "path"}).
			Rules(govy.NewRule(func(v mockStruct) error { return expectedErrs }))
		object := mockStruct{Field: "this"}
		errs := mustPropertyErrors(t, r.Validate(object))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyName:  "test.path",
			PropertyValue: internal.PropertyValueString(object),
			Errors:        []*govy.RuleError{{Message: expectedErrs.Error()}},
		}, errs[0])
	})

	t.Run("hide value", func(t *testing.T) {
		expectedErr := errors.New("oh no! here's the value: 'secret'")
		r := govy.For(func(m mockStruct) string { return "secret" }).
			WithPath(jsonpath.PropertySegment{Name: "test"}, jsonpath.PropertySegment{Name: "path"}).
			HideValue().
			Rules(govy.NewRule(func(v string) error { return expectedErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyName:  "test.path",
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

func TestPropertyRules_InferName(t *testing.T) {
	govyconfig.SetNameInferIncludeTestFiles(true)
	govyconfig.SetNameInferMode(govyconfig.NameInferModeRuntime)
	defer func() {
		govyconfig.SetNameInferIncludeTestFiles(false)
		govyconfig.SetNameInferMode(govyconfig.NameInferModeDisable)
	}()

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
			Rules(rules.EQ("John"))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke"}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'name' with value 'Luke':\n  - should be equal to 'John'")
	})
	t.Run("selector expression getter", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("John"))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke"}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'name' with value 'Luke':\n  - should be equal to 'John'")
	})
	t.Run("nested selector expression getter", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) int { return t.Details.Age.Years }).
			Rules(rules.EQ(29))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Details: Details{Age: Age{Years: 30}}}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'details.age.years' with value '30':\n  - should be equal to '29'")
	})
	t.Run("variable assignment", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) string {
				teacherName := t.Name
				return teacherName
			}).
			Rules(rules.EQ("John"))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke"}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'name' with value 'Luke':\n  - should be equal to 'John'")
	})
	t.Run("nested selector variable assignment", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) int {
				teacherAge := t.Details.Age.Years
				return teacherAge
			}).
			Rules(rules.EQ(29))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Details: Details{Age: Age{Years: 30}}}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'details.age.years' with value '30':\n  - should be equal to '29'")
	})
	t.Run("external function", func(t *testing.T) {
		getter := func(t Teacher) int {
			teacherAge := t.Details.Age.Years
			return teacherAge
		}
		r := govy.
			For(getter).
			Rules(rules.EQ(29))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Details: Details{Age: Age{Years: 30}}}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'details.age.years' with value '30':\n  - should be equal to '29'")
	})
	t.Run("pointer", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) *string { return t.Remarks }).
			Rules(rules.EQ(ptr("No remarks")))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Remarks: ptr("Some remarks")}))
		assert.Len(t, errs, 1)
		assert.ErrorContains(t, errs, "- 'remarks' with value 'Some remarks':\n  - should be equal to '")
	})
	t.Run("multiple return statements, infer from top level", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) *string {
				if t.Remarks == nil {
					return nil
				}
				return t.Remarks
			}).
			Rules(rules.EQ(ptr("No remarks")))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Remarks: ptr("Some remarks")}))
		assert.Len(t, errs, 1)
		assert.ErrorContains(t, errs, "- 'remarks' with value 'Some remarks':\n  - should be equal to '")
	})
	t.Run("multiple return statements, infer from nested if", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) *string {
				if t.Remarks != nil {
					return t.Remarks
				}
				return nil
			}).
			Rules(rules.EQ(ptr("No remarks")))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Name: "Luke", Remarks: ptr("Some remarks")}))
		assert.Len(t, errs, 1)
		assert.ErrorContains(t, errs, "- 'remarks' with value 'Some remarks':\n  - should be equal to '")
	})
	t.Run("no json tag", func(t *testing.T) {
		r := govy.
			For(func(t Teacher) string { return t.Surname }).
			Rules(rules.EQ("Cormack"))
		errs := mustPropertyErrors(t, r.Validate(Teacher{Surname: "Ellis"}))
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, "- 'Surname' with value 'Ellis':\n  - should be equal to 'Cormack'")
	})
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
