package validation

import (
	"strconv"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPropertyRules(t *testing.T) {
	type mockStruct struct {
		Field  string
		Fields []string
	}

	t.Run("no predicates, no error", func(t *testing.T) {
		r := For(func(m mockStruct) string { return "path" }).
			WithName("test.path").
			Rules(NewSingleRule(func(v string) error { return nil }))
		err := r.Validate(mockStruct{})
		assert.Nil(t, err)
	})

	t.Run("no predicates, validate", func(t *testing.T) {
		expectedErr := errors.New("ops!")
		r := For(func(m mockStruct) string { return "path" }).
			WithName("test.path").
			Rules(NewSingleRule(func(v string) error { return expectedErr }))
		errs := r.Validate(mockStruct{})
		require.Len(t, errs, 1)
		assert.Equal(t, &PropertyError{
			PropertyName:  "test.path",
			PropertyValue: "path",
			Errors:        []*RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("predicate matches, don't validate", func(t *testing.T) {
		r := For(func(m mockStruct) string { return "value" }).
			WithName("test.path").
			When(func(mockStruct) bool { return true }).
			When(func(mockStruct) bool { return true }).
			When(func(st mockStruct) bool { return st.Field == "" }).
			Rules(NewSingleRule(func(v string) error { return errors.New("ops!") }))
		err := r.Validate(mockStruct{Field: "something"})
		assert.Nil(t, err)
	})

	t.Run("multiple rules", func(t *testing.T) {
		err1 := errors.New("oh no!")
		r := For(func(m mockStruct) string { return "value" }).
			WithName("test.path").
			Rules(NewSingleRule(func(v string) error { return nil })).
			Rules(NewSingleRule(func(v string) error { return err1 })).
			Rules(NewSingleRule(func(v string) error { return nil })).
			Rules(NewSingleRule(func(v string) error {
				return NewPropertyError("nested", "nestedValue", &RuleError{
					Message: "property is required",
					Code:    ErrorCodeRequired,
				})
			}))
		errs := r.Validate(mockStruct{})
		require.Len(t, errs, 2)
		assert.ElementsMatch(t, PropertyErrors{
			&PropertyError{
				PropertyName:  "test.path",
				PropertyValue: "value",
				Errors:        []*RuleError{{Message: err1.Error()}},
			},
			&PropertyError{
				PropertyName:  "test.path.nested",
				PropertyValue: "nestedValue",
				Errors: []*RuleError{{
					Message: "property is required",
					Code:    ErrorCodeRequired,
				}},
			},
		}, errs)
	})

	t.Run("cascade mode stop", func(t *testing.T) {
		expectedErr := errors.New("oh no!")
		r := For(func(m mockStruct) string { return "value" }).
			WithName("test.path").
			Cascade(CascadeModeStop).
			Rules(NewSingleRule(func(v string) error { return expectedErr })).
			Rules(NewSingleRule(func(v string) error { return errors.New("no") }))
		errs := r.Validate(mockStruct{})
		require.Len(t, errs, 1)
		assert.Equal(t, &PropertyError{
			PropertyName:  "test.path",
			PropertyValue: "value",
			Errors:        []*RuleError{{Message: expectedErr.Error()}},
		}, errs[0])
	})

	t.Run("include validator", func(t *testing.T) {
		err1 := errors.New("oh no!")
		err2 := errors.New("included")
		err3 := errors.New("included again")
		r := For(func(m mockStruct) mockStruct { return m }).
			WithName("test.path").
			Rules(NewSingleRule(func(v mockStruct) error { return err1 })).
			Include(New(
				For(func(s mockStruct) string { return "value" }).
					WithName("included").
					Rules(NewSingleRule(func(v string) error { return err2 })).
					Rules(NewSingleRule(func(v string) error {
						return NewPropertyError("nested", "nestedValue", err3)
					})),
			))
		errs := r.Validate(mockStruct{})
		require.Len(t, errs, 3)
		assert.ElementsMatch(t, PropertyErrors{
			{
				PropertyName: "test.path",
				Errors:       []*RuleError{{Message: err1.Error()}},
			},
			{
				PropertyName:  "test.path.included",
				PropertyValue: "value",
				Errors:        []*RuleError{{Message: err2.Error()}},
			},
			{
				PropertyName:  "test.path.included.nested",
				PropertyValue: "nestedValue",
				Errors:        []*RuleError{{Message: err3.Error()}},
			},
		}, errs)
	})

	t.Run("get self", func(t *testing.T) {
		expectedErrs := errors.New("self error")
		r := For(GetSelf[mockStruct]()).
			WithName("test.path").
			Rules(NewSingleRule(func(v mockStruct) error { return expectedErrs }))
		object := mockStruct{Field: "this"}
		errs := r.Validate(object)
		require.Len(t, errs, 1)
		assert.Equal(t, &PropertyError{
			PropertyName:  "test.path",
			PropertyValue: propertyValueString(object),
			Errors:        []*RuleError{{Message: expectedErrs.Error()}},
		}, errs[0])
	})

	t.Run("hide value", func(t *testing.T) {
		expectedErr := errors.New("oh no! here's the value: 'secret'")
		r := For(func(m mockStruct) string { return "secret" }).
			WithName("test.path").
			HideValue().
			Rules(NewSingleRule(func(v string) error { return expectedErr }))
		errs := r.Validate(mockStruct{})
		require.Len(t, errs, 1)
		assert.Equal(t, &PropertyError{
			PropertyName:  "test.path",
			PropertyValue: "",
			Errors:        []*RuleError{{Message: "oh no! here's the value: '[hidden]'"}},
		}, errs[0])
	})
}

func TestForPointer(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		r := ForPointer(func(s *string) *string { return s })
		v, err := r.getter(nil)
		assert.Equal(t, "", v)
		assert.ErrorIs(t, err, emptyErr{})
	})
	t.Run("non nil pointer", func(t *testing.T) {
		r := ForPointer(func(s *string) *string { return s })
		s := "this string"
		v, err := r.getter(&s)
		assert.Equal(t, s, v)
		assert.NoError(t, err)
	})
}

func TestRequiredAndOmitEmpty(t *testing.T) {
	t.Run("nil pointer", func(t *testing.T) {
		rules := ForPointer(func(s *string) *string { return s }).
			Rules(StringMinLength(10))

		t.Run("implicit omitEmpty", func(t *testing.T) {
			err := rules.Validate(nil)
			assert.Nil(t, err)
		})
		t.Run("explicit omitEmpty", func(t *testing.T) {
			err := rules.OmitEmpty().Validate(nil)
			assert.Nil(t, err)
		})
		t.Run("required", func(t *testing.T) {
			errs := rules.Required().Validate(nil)
			assert.Len(t, errs, 1)
			assert.True(t, HasErrorCode(errs, ErrorCodeRequired))
		})
	})

	t.Run("non empty pointer", func(t *testing.T) {
		rules := ForPointer(func(s *string) *string { return s }).
			Rules(StringMinLength(10))

		t.Run("validate", func(t *testing.T) {
			errs := rules.Validate(ptr(""))
			assert.Len(t, errs, 1)
			assert.True(t, HasErrorCode(errs, ErrorCodeStringMinLength))
		})
		t.Run("omitEmpty", func(t *testing.T) {
			errs := rules.OmitEmpty().Validate(ptr(""))
			assert.Len(t, errs, 1)
			assert.True(t, HasErrorCode(errs, ErrorCodeStringMinLength))
		})
		t.Run("required", func(t *testing.T) {
			errs := rules.Required().Validate(ptr(""))
			assert.Len(t, errs, 1)
			assert.True(t, HasErrorCode(errs, ErrorCodeStringMinLength))
		})
	})
}

func TestTransform(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := Transform(getter, strconv.Atoi).
			Rules(GreaterThan(122))
		errs := transformed.Validate("123")
		assert.Empty(t, errs)
	})
	t.Run("fails validation", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := Transform(getter, strconv.Atoi).
			WithName("prop").
			Rules(GreaterThan(123))
		errs := transformed.Validate("123")
		assert.Len(t, errs, 1)
		assert.True(t, HasErrorCode(errs, ErrorCodeGreaterThan))
	})
	t.Run("zero value with omitEmpty", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := Transform(getter, strconv.Atoi).
			WithName("prop").
			OmitEmpty().
			Rules(GreaterThan(123))
		errs := transformed.Validate("")
		assert.Empty(t, errs)
	})
	t.Run("zero value with required", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := Transform(getter, strconv.Atoi).
			WithName("prop").
			Required().
			Rules(GreaterThan(123))
		errs := transformed.Validate("")
		assert.Len(t, errs, 1)
		assert.True(t, HasErrorCode(errs, ErrorCodeRequired))
	})
	t.Run("skip zero value", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := Transform(getter, strconv.Atoi).
			WithName("prop").
			Rules(GreaterThan(123))
		errs := transformed.Validate("")
		assert.Len(t, errs, 1)
		assert.True(t, HasErrorCode(errs, ErrorCodeGreaterThan))
	})
	t.Run("fails transformation", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := Transform(getter, strconv.Atoi).
			WithName("prop").
			Rules(GreaterThan(123))
		errs := transformed.Validate("123z")
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, expectedErrorOutput(t, "property_error_transform.txt"))
		assert.True(t, HasErrorCode(errs, ErrorCodeTransform))
	})
	t.Run("fail transformation with hidden value", func(t *testing.T) {
		getter := func(s string) string { return s }
		transformed := Transform(getter, strconv.Atoi).
			WithName("prop").
			HideValue().
			Rules(GreaterThan(123))
		errs := transformed.Validate("secret!")
		assert.Len(t, errs, 1)
		assert.EqualError(t, errs, expectedErrorOutput(t, "property_error_transform_with_hidden_value.txt"))
		assert.True(t, HasErrorCode(errs, ErrorCodeTransform))
	})
}

func ptr[T any](v T) *T { return &v }
