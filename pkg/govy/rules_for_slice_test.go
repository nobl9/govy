package govy_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govyconfig"
	"github.com/nobl9/govy/pkg/rules"
)

func TestPropertyRulesForEach(t *testing.T) {
	type mockStruct struct {
		Fields []string
	}

	t.Run("no predicates, no error", func(t *testing.T) {
		r := govy.ForSlice(func(m mockStruct) []string { return []string{"path"} }).
			WithName("test.path").
			RulesForEach(govy.NewRule(func(v string) error { return nil }))
		err := r.Validate(mockStruct{})
		assert.NoError(t, err)
	})

	t.Run("no predicates, validate", func(t *testing.T) {
		expectedErr := errors.New("ops!")
		r := govy.ForSlice(func(m mockStruct) []string { return []string{"path"} }).
			WithName("test.path").
			RulesForEach(govy.NewRule(func(v string) error { return expectedErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		require.Len(t, errs, 1)
		assert.Equal(t, &govy.PropertyError{
			PropertyName:        "test.path[0]",
			PropertyValue:       "path",
			IsSliceElementError: true,
			Errors:              []*govy.RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("predicate matches, don't validate", func(t *testing.T) {
		r := govy.ForSlice(func(m mockStruct) []string { return []string{"value"} }).
			WithName("test.path").
			When(func(mockStruct) bool { return true }).
			When(func(mockStruct) bool { return true }).
			When(func(st mockStruct) bool { return len(st.Fields) == 0 }).
			RulesForEach(govy.NewRule(func(v string) error { return errors.New("ops!") }))
		err := r.Validate(mockStruct{Fields: []string{"something"}})
		assert.NoError(t, err)
	})

	t.Run("multiple rules and for each rules", func(t *testing.T) {
		err1 := errors.New("oh no!")
		err2 := errors.New("another error...")
		err3 := errors.New("rule error")
		err4 := errors.New("rule error again")
		r := govy.ForSlice(func(m mockStruct) []string { return m.Fields }).
			WithName("test.path").
			Rules(govy.NewRule(func(v []string) error { return err3 })).
			RulesForEach(
				govy.NewRule(func(v string) error { return err1 }),
				govy.NewRule(func(v string) error {
					return govy.NewPropertyError("nested", "made-up", err2)
				}),
			).
			Rules(govy.NewRule(func(v []string) error {
				return govy.NewPropertyError("nested", "nestedValue", err4)
			}))

		errs := mustPropertyErrors(t, r.Validate(mockStruct{Fields: []string{"1", "2"}}))
		require.Len(t, errs, 6)
		assert.ElementsMatch(t, []*govy.PropertyError{
			{
				PropertyName:  "test.path",
				PropertyValue: `["1","2"]`,
				Errors:        []*govy.RuleError{{Message: err3.Error()}},
			},
			{
				PropertyName:  "test.path.nested",
				PropertyValue: "nestedValue",
				Errors:        []*govy.RuleError{{Message: err4.Error()}},
			},
			{
				PropertyName:        "test.path[0]",
				PropertyValue:       "1",
				IsSliceElementError: true,
				Errors:              []*govy.RuleError{{Message: err1.Error()}},
			},
			{
				PropertyName:        "test.path[1]",
				PropertyValue:       "2",
				IsSliceElementError: true,
				Errors:              []*govy.RuleError{{Message: err1.Error()}},
			},
			{
				PropertyName:        "test.path[0].nested",
				PropertyValue:       "made-up",
				IsSliceElementError: true,
				Errors:              []*govy.RuleError{{Message: err2.Error()}},
			},
			{
				PropertyName:        "test.path[1].nested",
				PropertyValue:       "made-up",
				IsSliceElementError: true,
				Errors:              []*govy.RuleError{{Message: err2.Error()}},
			},
		}, errs)
	})

	t.Run("cascade mode stop", func(t *testing.T) {
		expectedErr := errors.New("oh no!")
		r := govy.ForSlice(func(m mockStruct) []string { return []string{"value"} }).
			WithName("test.path").
			Cascade(govy.CascadeModeStop).
			RulesForEach(govy.NewRule(func(v string) error { return expectedErr })).
			RulesForEach(govy.NewRule(func(v string) error { return errors.New("no") }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		require.Len(t, errs, 1)
		assert.Equal(t, &govy.PropertyError{
			PropertyName:        "test.path[0]",
			PropertyValue:       "value",
			IsSliceElementError: true,
			Errors:              []*govy.RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("include for each validator", func(t *testing.T) {
		err1 := errors.New("oh no!")
		err2 := errors.New("included")
		err3 := errors.New("included again")
		r := govy.ForSlice(func(m mockStruct) []string { return m.Fields }).
			WithName("test.path").
			RulesForEach(govy.NewRule(func(v string) error { return err1 })).
			IncludeForEach(govy.New(
				govy.For(func(s string) string { return "nested" }).
					WithName("included").
					Rules(
						govy.NewRule(func(v string) error { return err2 }),
						govy.NewRule(func(v string) error { return err3 }),
					),
			))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{Fields: []string{"value"}}))
		require.Len(t, errs, 2)
		assert.ElementsMatch(t, []*govy.PropertyError{
			{
				PropertyName:        "test.path[0]",
				PropertyValue:       "value",
				IsSliceElementError: true,
				Errors:              []*govy.RuleError{{Message: err1.Error()}},
			},
			{
				PropertyName:        "test.path[0].included",
				PropertyValue:       "nested",
				IsSliceElementError: true,
				Errors: []*govy.RuleError{
					{Message: err2.Error()},
					{Message: err3.Error()},
				},
			},
		}, errs)
	})

	t.Run("include nested for slice", func(t *testing.T) {
		forEachErr := errors.New("oh no!")
		includedErr := errors.New("oh no!")
		inc := govy.New(
			govy.ForSlice(govy.GetSelf[[]string]()).
				RulesForEach(govy.NewRule(func(v string) error {
					if v == "value1" {
						return forEachErr
					}
					return govy.NewPropertyError("nested", "made-up", includedErr)
				})),
		)
		r := govy.For(func(m mockStruct) []string { return m.Fields }).
			WithName("test.path").
			Include(inc)

		errs := mustPropertyErrors(t, r.Validate(mockStruct{Fields: []string{"value1", "value2"}}))
		require.Len(t, errs, 2)
		assert.ElementsMatch(t, []*govy.PropertyError{
			{
				PropertyName:        "test.path[0]",
				PropertyValue:       "value1",
				IsSliceElementError: true,
				Errors:              []*govy.RuleError{{Message: forEachErr.Error()}},
			},
			{
				PropertyName:        "test.path[1].nested",
				PropertyValue:       "made-up",
				IsSliceElementError: true,
				Errors:              []*govy.RuleError{{Message: includedErr.Error()}},
			},
		}, errs)
	})
}

func TestPropertyRulesForSlice_InferName(t *testing.T) {
	govyconfig.SetNameInferIncludeTestFiles(true)
	govyconfig.SetNameInferMode(govyconfig.NameInferModeRuntime)
	defer func() {
		govyconfig.SetNameInferIncludeTestFiles(false)
		govyconfig.SetNameInferMode(govyconfig.NameInferModeDisable)
	}()

	type Teacher struct {
		Students []string `json:"students"`
	}

	r := govy.ForSlice(func(t Teacher) []string { return t.Students }).
		RulesForEach(rules.EQ("John"))
	errs := mustPropertyErrors(t, r.Validate(Teacher{Students: []string{"Luke"}}))
	assert.Len(t, errs, 1)
	assert.EqualError(t, errs, "- 'students[0]' with value 'Luke':\n  - should be equal to 'John'")
}
