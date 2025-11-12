package govy_test

import (
	"errors"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govyconfig"
	"github.com/nobl9/govy/pkg/rules"
)

func TestPropertyRulesForMap(t *testing.T) {
	type Labels map[string]string
	type mockStruct struct {
		StringMap map[string]string
		IntMap    map[string]int
		Labels    Labels
	}

	t.Run("no predicates, no error", func(t *testing.T) {
		baseRules := govy.ForMap(func(m mockStruct) map[string]string { return map[string]string{"key": "value"} }).
			WithName("test.path")
		for _, r := range []govy.PropertyRulesForMap[map[string]string, string, string, mockStruct]{
			baseRules.RulesForKeys(govy.NewRule(func(v string) error { return nil })),
			baseRules.RulesForValues(govy.NewRule(func(v string) error { return nil })),
			baseRules.RulesForItems(govy.NewRule(func(v govy.MapItem[string, string]) error { return nil })),
		} {
			err := r.Validate(mockStruct{})
			assert.NoError(t, err)
		}
	})

	t.Run("no predicates, validate", func(t *testing.T) {
		expectedErr := errors.New("ops!")
		baseRules := govy.ForMap(func(m mockStruct) map[string]string { return map[string]string{"key": "value"} }).
			WithName("test.path")
		for name, test := range map[string]struct {
			Rules    govy.PropertyRulesForMap[map[string]string, string, string, mockStruct]
			Expected *govy.PropertyError
		}{
			"keys": {
				Rules: baseRules.RulesForKeys(govy.NewRule(func(v string) error { return expectedErr })),
				Expected: &govy.PropertyError{
					PropertyName:  "test.path.key",
					PropertyValue: "key",
					IsKeyError:    true,
					Errors:        []*govy.RuleError{{Message: expectedErr.Error()}},
				},
			},
			"values": {
				Rules: baseRules.RulesForValues(govy.NewRule(func(v string) error { return expectedErr })),
				Expected: &govy.PropertyError{
					PropertyName:  "test.path.key",
					PropertyValue: "value",
					Errors:        []*govy.RuleError{{Message: expectedErr.Error()}},
				},
			},
			"items": {
				Rules: baseRules.RulesForItems(
					govy.NewRule(func(v govy.MapItem[string, string]) error { return expectedErr }),
				),
				Expected: &govy.PropertyError{
					PropertyName:  "test.path.key",
					PropertyValue: "value",
					Errors:        []*govy.RuleError{{Message: expectedErr.Error()}},
				},
			},
		} {
			t.Run(name, func(t *testing.T) {
				errs := mustPropertyErrors(t, test.Rules.Validate(mockStruct{}))
				assert.Require(t, assert.Len(t, errs, 1))
				assert.Equal(t, test.Expected, errs[0])
			})
		}
	})

	t.Run("predicate matches, don't validate", func(t *testing.T) {
		baseRules := govy.ForMap(func(m mockStruct) map[string]string { return map[string]string{"key": "value"} }).
			WithName("test.path").
			When(func(mockStruct) bool { return true }).
			When(func(mockStruct) bool { return true }).
			When(func(st mockStruct) bool { return len(st.StringMap) == 0 })
		for _, r := range []govy.PropertyRulesForMap[map[string]string, string, string, mockStruct]{
			baseRules.RulesForKeys(govy.NewRule(func(v string) error { return errors.New("ops!") })),
			baseRules.RulesForValues(govy.NewRule(func(v string) error { return errors.New("ops!") })),
			baseRules.RulesForItems(
				govy.NewRule(func(v govy.MapItem[string, string]) error { return errors.New("ops!") }),
			),
		} {
			err := r.Validate(mockStruct{StringMap: map[string]string{"different": "map"}})
			assert.NoError(t, err)
		}
	})

	t.Run("multiple rules for keys, values and items", func(t *testing.T) {
		errRule := errors.New("rule error")
		errKey := errors.New("key error")
		errNestedKey := errors.New("nested key error")
		errValue := errors.New("value error")
		errNestedValue := errors.New("nested value error")
		errItem := errors.New("value item")
		errNestedItem := errors.New("nested item error")
		errNestedRule := errors.New("nested rule error")

		r := govy.ForMap(func(m mockStruct) map[string]string { return m.StringMap }).
			WithName("test.path").
			Rules(govy.NewRule(func(v map[string]string) error { return errRule })).
			RulesForKeys(
				govy.NewRule(func(v string) error { return errKey }),
				govy.NewRule(func(v string) error {
					return govy.NewPropertyError("nested", "nestedKey", errNestedKey)
				}),
			).
			RulesForValues(
				govy.NewRule(func(v string) error { return errValue }),
				govy.NewRule(func(v string) error {
					return govy.NewPropertyError("nested", "nestedValue", errNestedValue)
				}),
			).
			RulesForItems(
				govy.NewRule(func(v govy.MapItem[string, string]) error { return errItem }),
				govy.NewRule(func(v govy.MapItem[string, string]) error {
					return govy.NewPropertyError("nested", "nestedItem", errNestedItem)
				}),
			).
			Rules(govy.NewRule(func(v map[string]string) error {
				return govy.NewPropertyError("nested", "nestedRule", errNestedRule)
			}))

		errs := mustPropertyErrors(t, r.Validate(mockStruct{StringMap: map[string]string{
			"key1":  "value1",
			"key 2": "value2",
		}}))
		assert.Require(t, assert.Len(t, errs, 12))
		assert.ElementsMatch(t, []*govy.PropertyError{
			{
				PropertyName:  "test.path",
				PropertyValue: `{"key 2":"value2","key1":"value1"}`,
				Errors:        []*govy.RuleError{{Message: errRule.Error()}},
			},
			{
				PropertyName:  "test.path.key1",
				PropertyValue: "key1",
				IsKeyError:    true,
				Errors:        []*govy.RuleError{{Message: errKey.Error()}},
			},
			{
				PropertyName:  "test.path.['key 2']",
				PropertyValue: "key 2",
				IsKeyError:    true,
				Errors:        []*govy.RuleError{{Message: errKey.Error()}},
			},
			{
				PropertyName:  "test.path.key1.nested",
				PropertyValue: "nestedKey",
				IsKeyError:    true,
				Errors:        []*govy.RuleError{{Message: errNestedKey.Error()}},
			},
			{
				PropertyName:  "test.path.['key 2'].nested",
				PropertyValue: "nestedKey",
				IsKeyError:    true,
				Errors:        []*govy.RuleError{{Message: errNestedKey.Error()}},
			},
			{
				PropertyName:  "test.path.key1",
				PropertyValue: "value1",
				Errors: []*govy.RuleError{
					{Message: errValue.Error()},
					{Message: errItem.Error()},
				},
			},
			{
				PropertyName:  "test.path.['key 2']",
				PropertyValue: "value2",
				Errors: []*govy.RuleError{
					{Message: errValue.Error()},
					{Message: errItem.Error()},
				},
			},
			{
				PropertyName:  "test.path.key1.nested",
				PropertyValue: "nestedValue",
				Errors:        []*govy.RuleError{{Message: errNestedValue.Error()}},
			},
			{
				PropertyName:  "test.path.['key 2'].nested",
				PropertyValue: "nestedValue",
				Errors:        []*govy.RuleError{{Message: errNestedValue.Error()}},
			},
			{
				PropertyName:  "test.path.key1.nested",
				PropertyValue: "value1",
				Errors:        []*govy.RuleError{{Message: errNestedItem.Error()}},
			},
			{
				PropertyName:  "test.path.['key 2'].nested",
				PropertyValue: "value2",
				Errors:        []*govy.RuleError{{Message: errNestedItem.Error()}},
			},
			{
				PropertyName:  "test.path.nested",
				PropertyValue: "nestedRule",
				Errors:        []*govy.RuleError{{Message: errNestedRule.Error()}},
			},
		}, errs)
	})

	t.Run("cascade mode stop", func(t *testing.T) {
		keyErr := errors.New("key error")
		valueErr := errors.New("value error")
		r := govy.ForMap(func(m mockStruct) map[string]string { return map[string]string{"key": "value"} }).
			WithName("test.path").
			Cascade(govy.CascadeModeStop).
			RulesForValues(govy.NewRule(func(v string) error { return valueErr })).
			RulesForKeys(govy.NewRule(func(v string) error { return keyErr }))
		errs := mustPropertyErrors(t, r.Validate(mockStruct{}))
		assert.Require(t, assert.Len(t, errs, 2))
		assert.ElementsMatch(t, []*govy.PropertyError{
			{
				PropertyName:  "test.path.key",
				PropertyValue: "key",
				IsKeyError:    true,
				Errors:        []*govy.RuleError{{Message: keyErr.Error()}},
			},
			{
				PropertyName:  "test.path.key",
				PropertyValue: "value",
				Errors:        []*govy.RuleError{{Message: valueErr.Error()}},
			},
		}, errs)
	})

	t.Run("include for keys validator", func(t *testing.T) {
		errRule := errors.New("rule error")
		errIncludedKey1 := errors.New("included key 1 error")
		errIncludedKey2 := errors.New("included key 2 error")
		errIncludedValue1 := errors.New("included value 1 error")
		errIncludedValue2 := errors.New("included value 2 error")
		errIncludedItem1 := errors.New("included item 1 error")
		errIncludedItem2 := errors.New("included item 2 error")

		r := govy.ForMap(func(m mockStruct) map[string]int { return m.IntMap }).
			WithName("test.path").
			Rules(govy.NewRule(func(v map[string]int) error { return errRule })).
			IncludeForKeys(govy.New(
				govy.For(func(s string) string { return s }).
					WithName("included_key").
					Rules(
						govy.NewRule(func(v string) error { return errIncludedKey1 }),
						govy.NewRule(func(v string) error { return errIncludedKey2 }),
					),
			)).
			IncludeForValues(govy.New(
				govy.For(func(i int) int { return i }).
					WithName("included_value").
					Rules(
						govy.NewRule(func(v int) error { return errIncludedValue1 }),
						govy.NewRule(func(v int) error { return errIncludedValue2 }),
					),
			)).
			IncludeForItems(govy.New(
				govy.For(func(i govy.MapItem[string, int]) govy.MapItem[string, int] { return i }).
					WithName("included_item").
					Rules(
						govy.NewRule(func(v govy.MapItem[string, int]) error { return errIncludedItem1 }),
						govy.NewRule(func(v govy.MapItem[string, int]) error { return errIncludedItem2 }),
					),
			))

		errs := mustPropertyErrors(t, r.Validate(mockStruct{IntMap: map[string]int{"key": 1}}))
		assert.Require(t, assert.Len(t, errs, 4))
		assert.ElementsMatch(t, []*govy.PropertyError{
			{
				PropertyName:  "test.path",
				PropertyValue: `{"key":1}`,
				Errors:        []*govy.RuleError{{Message: errRule.Error()}},
			},
			{
				PropertyName:  "test.path.key.included_key",
				PropertyValue: "key",
				IsKeyError:    true,
				Errors: []*govy.RuleError{
					{Message: errIncludedKey1.Error()},
					{Message: errIncludedKey2.Error()},
				},
			},
			{
				PropertyName:  "test.path.key.included_value",
				PropertyValue: "1",
				Errors: []*govy.RuleError{
					{Message: errIncludedValue1.Error()},
					{Message: errIncludedValue2.Error()},
				},
			},
			{
				PropertyName:  "test.path.key.included_item",
				PropertyValue: "1",
				Errors: []*govy.RuleError{
					{Message: errIncludedItem1.Error()},
					{Message: errIncludedItem2.Error()},
				},
			},
		}, errs)
	})

	t.Run("include for keys validator, key and value are same type", func(t *testing.T) {
		errRule := errors.New("rule error")
		errIncludedKey1 := errors.New("included key 1 error")
		errIncludedKey2 := errors.New("included key 2 error")
		errIncludedValue1 := errors.New("included value 1 error")
		errIncludedValue2 := errors.New("included value 2 error")
		errIncludedItem1 := errors.New("included item 1 error")
		errIncludedItem2 := errors.New("included item 2 error")

		r := govy.ForMap(func(m mockStruct) map[string]string { return m.StringMap }).
			WithName("test.path").
			Rules(govy.NewRule(func(v map[string]string) error { return errRule })).
			IncludeForKeys(govy.New(
				govy.For(func(s string) string { return s }).
					WithName("included_key").
					Rules(
						govy.NewRule(func(v string) error { return errIncludedKey1 }),
						govy.NewRule(func(v string) error { return errIncludedKey2 }),
					),
			)).
			IncludeForValues(govy.New(
				govy.For(func(i string) string { return i }).
					WithName("included_value").
					Rules(
						govy.NewRule(func(v string) error { return errIncludedValue1 }),
						govy.NewRule(func(v string) error { return errIncludedValue2 }),
					),
			)).
			IncludeForItems(govy.New(
				govy.For(func(i govy.MapItem[string, string]) govy.MapItem[string, string] { return i }).
					WithName("included_item").
					Rules(
						govy.NewRule(func(v govy.MapItem[string, string]) error { return errIncludedItem1 }),
						govy.NewRule(func(v govy.MapItem[string, string]) error { return errIncludedItem2 }),
					),
			))

		errs := mustPropertyErrors(t, r.Validate(mockStruct{StringMap: map[string]string{"key": "1"}}))
		assert.Require(t, assert.Len(t, errs, 4))
		assert.ElementsMatch(t, []*govy.PropertyError{
			{
				PropertyName:  "test.path",
				PropertyValue: `{"key":"1"}`,
				Errors:        []*govy.RuleError{{Message: errRule.Error()}},
			},
			{
				PropertyName:  "test.path.key.included_key",
				PropertyValue: "key",
				IsKeyError:    true,
				Errors: []*govy.RuleError{
					{Message: errIncludedKey1.Error()},
					{Message: errIncludedKey2.Error()},
				},
			},
			{
				PropertyName:  "test.path.key.included_value",
				PropertyValue: "1",
				Errors: []*govy.RuleError{
					{Message: errIncludedValue1.Error()},
					{Message: errIncludedValue2.Error()},
				},
			},
			{
				PropertyName:  "test.path.key.included_item",
				PropertyValue: "1",
				Errors: []*govy.RuleError{
					{Message: errIncludedItem1.Error()},
					{Message: errIncludedItem2.Error()},
				},
			},
		}, errs)
	})

	t.Run("include nested for map", func(t *testing.T) {
		expectedErr := errors.New("oh no!")
		inc := govy.New(
			govy.ForMap(govy.GetSelf[map[string]string]()).
				RulesForValues(govy.NewRule(func(v string) error { return expectedErr })),
		)
		r := govy.For(func(m mockStruct) map[string]string { return m.StringMap }).
			WithName("test.path").
			Include(inc)

		errs := mustPropertyErrors(t, r.Validate(mockStruct{StringMap: map[string]string{"key": "value"}}))
		assert.Require(t, assert.Len(t, errs, 1))
		assert.Equal(t, &govy.PropertyError{
			PropertyName:  "test.path.key",
			PropertyValue: "value",
			Errors:        []*govy.RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("include validator", func(t *testing.T) {
		err1 := errors.New("error1")
		err2 := errors.New("error2")

		r := govy.ForMap(func(m mockStruct) Labels { return m.Labels }).
			WithName("labels").
			Include(govy.New(
				govy.ForMap(govy.GetSelf[Labels]()).
					RulesForValues(govy.NewRule(func(v string) error { return err1 })),
			)).
			RulesForKeys(govy.NewRule(func(v string) error { return err2 }))

		errs := mustPropertyErrors(t, r.Validate(mockStruct{Labels: Labels{"key": "value"}}))
		assert.Require(t, assert.Len(t, errs, 2))
		assert.ElementsMatch(t, []*govy.PropertyError{
			{
				PropertyName:  "labels.key",
				PropertyValue: "key",
				IsKeyError:    true,
				Errors: []*govy.RuleError{
					{Message: err2.Error()},
				},
			},
			{
				PropertyName:  "labels.key",
				PropertyValue: "value",
				Errors: []*govy.RuleError{
					{Message: err1.Error()},
				},
			},
		}, errs)
	})
}

func TestPropertyRulesForMap_InferName(t *testing.T) {
	govyconfig.SetInferNameIncludeTestFiles(true)
	defer govyconfig.SetInferNameIncludeTestFiles(false)

	type Teacher struct {
		Students map[string]int `json:"students"`
	}

	r := govy.ForMap(func(t Teacher) map[string]int { return t.Students }).
		InferName(govy.InferNameModeRuntime).
		RulesForKeys(rules.EQ("John"))
	errs := mustPropertyErrors(t, r.Validate(Teacher{Students: map[string]int{"Luke": 35}}))
	assert.Len(t, errs, 1)
	assert.EqualError(t, errs, "- 'students.Luke' with key 'Luke':\n  - should be equal to 'John'")
}
