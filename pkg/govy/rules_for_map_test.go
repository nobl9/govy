package validation

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPropertyRulesForMap(t *testing.T) {
	type mockStruct struct {
		StringMap map[string]string
		IntMap    map[string]int
	}

	t.Run("no predicates, no error", func(t *testing.T) {
		baseRules := ForMap(func(m mockStruct) map[string]string { return map[string]string{"key": "value"} }).
			WithName("test.path")
		for _, r := range []PropertyRulesForMap[map[string]string, string, string, mockStruct]{
			baseRules.RulesForKeys(NewSingleRule(func(v string) error { return nil })),
			baseRules.RulesForValues(NewSingleRule(func(v string) error { return nil })),
			baseRules.RulesForItems(NewSingleRule(func(v MapItem[string, string]) error { return nil })),
		} {
			errs := r.Validate(mockStruct{})
			assert.Nil(t, errs)
		}
	})

	t.Run("no predicates, validate", func(t *testing.T) {
		expectedErr := errors.New("ops!")
		baseRules := ForMap(func(m mockStruct) map[string]string { return map[string]string{"key": "value"} }).
			WithName("test.path")
		for name, test := range map[string]struct {
			Rules    PropertyRulesForMap[map[string]string, string, string, mockStruct]
			Expected *PropertyError
		}{
			"keys": {
				Rules: baseRules.RulesForKeys(NewSingleRule(func(v string) error { return expectedErr })),
				Expected: &PropertyError{
					PropertyName:  "test.path.key",
					PropertyValue: "key",
					IsKeyError:    true,
					Errors:        []*RuleError{{Message: expectedErr.Error()}},
				},
			},
			"values": {
				Rules: baseRules.RulesForValues(NewSingleRule(func(v string) error { return expectedErr })),
				Expected: &PropertyError{
					PropertyName:  "test.path.key",
					PropertyValue: "value",
					Errors:        []*RuleError{{Message: expectedErr.Error()}},
				},
			},
			"items": {
				Rules: baseRules.RulesForItems(NewSingleRule(func(v MapItem[string, string]) error { return expectedErr })),
				Expected: &PropertyError{
					PropertyName:  "test.path.key",
					PropertyValue: "value",
					Errors:        []*RuleError{{Message: expectedErr.Error()}},
				},
			},
		} {
			t.Run(name, func(t *testing.T) {
				errs := test.Rules.Validate(mockStruct{})
				require.Len(t, errs, 1)
				assert.Equal(t, test.Expected, errs[0])
			})
		}
	})

	t.Run("predicate matches, don't validate", func(t *testing.T) {
		baseRules := ForMap(func(m mockStruct) map[string]string { return map[string]string{"key": "value"} }).
			WithName("test.path").
			When(func(mockStruct) bool { return true }).
			When(func(mockStruct) bool { return true }).
			When(func(st mockStruct) bool { return len(st.StringMap) == 0 })
		for _, r := range []PropertyRulesForMap[map[string]string, string, string, mockStruct]{
			baseRules.RulesForKeys(NewSingleRule(func(v string) error { return errors.New("ops!") })),
			baseRules.RulesForValues(NewSingleRule(func(v string) error { return errors.New("ops!") })),
			baseRules.RulesForItems(NewSingleRule(func(v MapItem[string, string]) error { return errors.New("ops!") })),
		} {
			errs := r.Validate(mockStruct{StringMap: map[string]string{"different": "map"}})
			assert.Nil(t, errs)
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

		r := ForMap(func(m mockStruct) map[string]string { return m.StringMap }).
			WithName("test.path").
			Rules(NewSingleRule(func(v map[string]string) error { return errRule })).
			RulesForKeys(
				NewSingleRule(func(v string) error { return errKey }),
				NewSingleRule(func(v string) error {
					return NewPropertyError("nested", "nestedKey", errNestedKey)
				}),
			).
			RulesForValues(
				NewSingleRule(func(v string) error { return errValue }),
				NewSingleRule(func(v string) error {
					return NewPropertyError("nested", "nestedValue", errNestedValue)
				}),
			).
			RulesForItems(
				NewSingleRule(func(v MapItem[string, string]) error { return errItem }),
				NewSingleRule(func(v MapItem[string, string]) error {
					return NewPropertyError("nested", "nestedItem", errNestedItem)
				}),
			).
			Rules(NewSingleRule(func(v map[string]string) error {
				return NewPropertyError("nested", "nestedRule", errNestedRule)
			}))

		errs := r.Validate(mockStruct{StringMap: map[string]string{
			"key1": "value1",
			"key2": "value2",
		}})
		require.Len(t, errs, 12)
		assert.ElementsMatch(t, []*PropertyError{
			{
				PropertyName:  "test.path",
				PropertyValue: `{"key1":"value1","key2":"value2"}`,
				Errors:        []*RuleError{{Message: errRule.Error()}},
			},
			{
				PropertyName:  "test.path.key1",
				PropertyValue: "key1",
				IsKeyError:    true,
				Errors:        []*RuleError{{Message: errKey.Error()}},
			},
			{
				PropertyName:  "test.path.key2",
				PropertyValue: "key2",
				IsKeyError:    true,
				Errors:        []*RuleError{{Message: errKey.Error()}},
			},
			{
				PropertyName:  "test.path.key1.nested",
				PropertyValue: "nestedKey",
				IsKeyError:    true,
				Errors:        []*RuleError{{Message: errNestedKey.Error()}},
			},
			{
				PropertyName:  "test.path.key2.nested",
				PropertyValue: "nestedKey",
				IsKeyError:    true,
				Errors:        []*RuleError{{Message: errNestedKey.Error()}},
			},
			{
				PropertyName:  "test.path.key1",
				PropertyValue: "value1",
				Errors: []*RuleError{
					{Message: errValue.Error()},
					{Message: errItem.Error()},
				},
			},
			{
				PropertyName:  "test.path.key2",
				PropertyValue: "value2",
				Errors: []*RuleError{
					{Message: errValue.Error()},
					{Message: errItem.Error()},
				},
			},
			{
				PropertyName:  "test.path.key1.nested",
				PropertyValue: "nestedValue",
				Errors:        []*RuleError{{Message: errNestedValue.Error()}},
			},
			{
				PropertyName:  "test.path.key2.nested",
				PropertyValue: "nestedValue",
				Errors:        []*RuleError{{Message: errNestedValue.Error()}},
			},
			{
				PropertyName:  "test.path.key1.nested",
				PropertyValue: "value1",
				Errors:        []*RuleError{{Message: errNestedItem.Error()}},
			},
			{
				PropertyName:  "test.path.key2.nested",
				PropertyValue: "value2",
				Errors:        []*RuleError{{Message: errNestedItem.Error()}},
			},
			{
				PropertyName:  "test.path.nested",
				PropertyValue: "nestedRule",
				Errors:        []*RuleError{{Message: errNestedRule.Error()}},
			},
		}, errs)
	})

	t.Run("cascade mode stop", func(t *testing.T) {
		keyErr := errors.New("key error")
		valueErr := errors.New("value error")
		r := ForMap(func(m mockStruct) map[string]string { return map[string]string{"key": "value"} }).
			WithName("test.path").
			Cascade(CascadeModeStop).
			RulesForValues(NewSingleRule(func(v string) error { return valueErr })).
			RulesForKeys(NewSingleRule(func(v string) error { return keyErr }))
		errs := r.Validate(mockStruct{})
		require.Len(t, errs, 2)
		assert.ElementsMatch(t, []*PropertyError{
			{
				PropertyName:  "test.path.key",
				PropertyValue: "key",
				IsKeyError:    true,
				Errors:        []*RuleError{{Message: keyErr.Error()}},
			},
			{
				PropertyName:  "test.path.key",
				PropertyValue: "value",
				Errors:        []*RuleError{{Message: valueErr.Error()}},
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

		r := ForMap(func(m mockStruct) map[string]int { return m.IntMap }).
			WithName("test.path").
			Rules(NewSingleRule(func(v map[string]int) error { return errRule })).
			IncludeForKeys(New(
				For(func(s string) string { return s }).
					WithName("included_key").
					Rules(
						NewSingleRule(func(v string) error { return errIncludedKey1 }),
						NewSingleRule(func(v string) error { return errIncludedKey2 }),
					),
			)).
			IncludeForValues(New(
				For(func(i int) int { return i }).
					WithName("included_value").
					Rules(
						NewSingleRule(func(v int) error { return errIncludedValue1 }),
						NewSingleRule(func(v int) error { return errIncludedValue2 }),
					),
			)).
			IncludeForItems(New(
				For(func(i MapItem[string, int]) MapItem[string, int] { return i }).
					WithName("included_item").
					Rules(
						NewSingleRule(func(v MapItem[string, int]) error { return errIncludedItem1 }),
						NewSingleRule(func(v MapItem[string, int]) error { return errIncludedItem2 }),
					),
			))

		errs := r.Validate(mockStruct{IntMap: map[string]int{"key": 1}})
		require.Len(t, errs, 4)
		assert.ElementsMatch(t, []*PropertyError{
			{
				PropertyName:  "test.path",
				PropertyValue: `{"key":1}`,
				Errors:        []*RuleError{{Message: errRule.Error()}},
			},
			{
				PropertyName:  "test.path.key.included_key",
				PropertyValue: "key",
				IsKeyError:    true,
				Errors: []*RuleError{
					{Message: errIncludedKey1.Error()},
					{Message: errIncludedKey2.Error()},
				},
			},
			{
				PropertyName:  "test.path.key.included_value",
				PropertyValue: "1",
				Errors: []*RuleError{
					{Message: errIncludedValue1.Error()},
					{Message: errIncludedValue2.Error()},
				},
			},
			{
				PropertyName:  "test.path.key.included_item",
				PropertyValue: "1",
				Errors: []*RuleError{
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

		r := ForMap(func(m mockStruct) map[string]string { return m.StringMap }).
			WithName("test.path").
			Rules(NewSingleRule(func(v map[string]string) error { return errRule })).
			IncludeForKeys(New(
				For(func(s string) string { return s }).
					WithName("included_key").
					Rules(
						NewSingleRule(func(v string) error { return errIncludedKey1 }),
						NewSingleRule(func(v string) error { return errIncludedKey2 }),
					),
			)).
			IncludeForValues(New(
				For(func(i string) string { return i }).
					WithName("included_value").
					Rules(
						NewSingleRule(func(v string) error { return errIncludedValue1 }),
						NewSingleRule(func(v string) error { return errIncludedValue2 }),
					),
			)).
			IncludeForItems(New(
				For(func(i MapItem[string, string]) MapItem[string, string] { return i }).
					WithName("included_item").
					Rules(
						NewSingleRule(func(v MapItem[string, string]) error { return errIncludedItem1 }),
						NewSingleRule(func(v MapItem[string, string]) error { return errIncludedItem2 }),
					),
			))

		errs := r.Validate(mockStruct{StringMap: map[string]string{"key": "1"}})
		require.Len(t, errs, 4)
		assert.ElementsMatch(t, []*PropertyError{
			{
				PropertyName:  "test.path",
				PropertyValue: `{"key":"1"}`,
				Errors:        []*RuleError{{Message: errRule.Error()}},
			},
			{
				PropertyName:  "test.path.key.included_key",
				PropertyValue: "key",
				IsKeyError:    true,
				Errors: []*RuleError{
					{Message: errIncludedKey1.Error()},
					{Message: errIncludedKey2.Error()},
				},
			},
			{
				PropertyName:  "test.path.key.included_value",
				PropertyValue: "1",
				Errors: []*RuleError{
					{Message: errIncludedValue1.Error()},
					{Message: errIncludedValue2.Error()},
				},
			},
			{
				PropertyName:  "test.path.key.included_item",
				PropertyValue: "1",
				Errors: []*RuleError{
					{Message: errIncludedItem1.Error()},
					{Message: errIncludedItem2.Error()},
				},
			},
		}, errs)
	})

	t.Run("include nested for map", func(t *testing.T) {
		expectedErr := errors.New("oh no!")
		inc := New(
			ForMap(GetSelf[map[string]string]()).
				RulesForValues(NewSingleRule(func(v string) error { return expectedErr })),
		)
		r := For(func(m mockStruct) map[string]string { return m.StringMap }).
			WithName("test.path").
			Include(inc)

		errs := r.Validate(mockStruct{StringMap: map[string]string{"key": "value"}})
		require.Len(t, errs, 1)
		assert.Equal(t, &PropertyError{
			PropertyName:  "test.path.key",
			PropertyValue: "value",
			Errors:        []*RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})
}
