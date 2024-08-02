package validation

import (
	"embed"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed test_data
var errorsTestData embed.FS

func TestValidatorError(t *testing.T) {
	for name, err := range map[string]*ValidatorError{
		"no_name": {
			Errors: PropertyErrors{
				{
					PropertyName:  "this",
					PropertyValue: "123",
					Errors:        []*RuleError{{Message: "this is an error"}},
				},
				{
					PropertyName: "that",
					Errors:       []*RuleError{{Message: "that is an error"}},
				},
			},
		},
		"with_name": {
			Name: "Teacher",
			Errors: PropertyErrors{
				{
					PropertyName:  "this",
					PropertyValue: "123",
					Errors:        []*RuleError{{Message: "this is an error"}},
				},
				{
					PropertyName: "that",
					Errors:       []*RuleError{{Message: "that is an error"}},
				},
			},
		},
		"prop_no_name": {
			Errors: PropertyErrors{
				{
					Errors: []*RuleError{{Message: "no name"}},
				},
				{
					PropertyName: "that",
					Errors:       []*RuleError{{Message: "that is an error"}},
				},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			assert.EqualError(t, err, expectedErrorOutput(t, fmt.Sprintf("validator_error_%s.txt", name)))
		})
	}
}

func TestNewPropertyError(t *testing.T) {
	t.Run("string value", func(t *testing.T) {
		err := NewPropertyError("name", "value",
			&RuleError{Message: "top", Code: "1"},
			ruleSetError{
				&RuleError{Message: "rule1", Code: "2"},
				&RuleError{Message: "rule2", Code: "3"},
			},
			&RuleError{Message: "top", Code: "4"},
		)
		assert.Equal(t, &PropertyError{
			PropertyName:  "name",
			PropertyValue: "value",
			Errors: []*RuleError{
				{Message: "top", Code: "1"},
				{Message: "rule1", Code: "2"},
				{Message: "rule2", Code: "3"},
				{Message: "top", Code: "4"},
			},
		}, err)
	})
	for name, test := range map[string]struct {
		InputValue    interface{}
		ExpectedValue string
	}{
		"map": {
			InputValue:    map[string]string{"key": "value"},
			ExpectedValue: `{"key":"value"}`,
		},
		"struct": {
			InputValue: struct {
				V string `json:"that"`
			}{
				V: "this",
			},
			ExpectedValue: `{"that":"this"}`,
		},
		"slice": {
			InputValue:    []string{"value"},
			ExpectedValue: `["value"]`,
		},
		"integer": {
			InputValue:    0,
			ExpectedValue: "0",
		},
		"float": {
			InputValue:    10.1,
			ExpectedValue: "10.1",
		},
		"boolean": {
			InputValue:    false,
			ExpectedValue: "false",
		},
		"pointer": {
			InputValue:    ptr(10.2),
			ExpectedValue: "10.2",
		},
		"initialized nil": {
			InputValue:    func() *float64 { return nil }(),
			ExpectedValue: "",
		},
		"nil": {
			InputValue:    nil,
			ExpectedValue: "",
		},
		"blank lines": {
			InputValue:    ` 		SELECT value FROM my-table WHERE value = "abc"    `,
			ExpectedValue: `SELECT value FROM my-table WHERE value = "abc"`,
		},
		"multiline": {
			InputValue: `
SELECT value FROM
my-table WHERE value = "abc"
`,
			ExpectedValue: "SELECT value FROM\\nmy-table WHERE value = \"abc\"",
		},
		"carriage return": {
			InputValue:    "return\rcarriage",
			ExpectedValue: "return\\rcarriage",
		},
	} {
		t.Run(name, func(t *testing.T) {
			err := NewPropertyError(
				"name",
				test.InputValue,
				&RuleError{Message: "msg"})
			assert.Equal(t, &PropertyError{
				PropertyName:  "name",
				PropertyValue: test.ExpectedValue,
				Errors:        []*RuleError{{Message: "msg"}},
			}, err)
		})
	}
}

type (
	stringerWithTags struct {
		This string `json:"this"`
		That string `json:"THAT"`
	}
	stringerWithoutTags struct {
		This string
		That string
	}
)

func (s stringerWithTags) String() string    { return s.This + "_" + s.That }
func (s stringerWithoutTags) String() string { return s.This + "_" + s.That }

func TestPropertyError(t *testing.T) {
	for name, value := range map[string]interface{}{
		"string": "default",
		"slice":  []string{"this", "that"},
		"map":    map[string]string{"this": "that"},
		"struct": struct {
			This string `json:"this"`
			That string `json:"THAT"`
		}{This: "this", That: "that"},
		"stringer_with_json_tags": stringerWithTags{
			This: "this", That: "that",
		},
		"stringer_without_json_tags": stringerWithoutTags{
			This: "this", That: "that",
		},
		"stringer_pointer": &stringerWithoutTags{
			This: "this", That: "that",
		},
	} {
		t.Run(name, func(t *testing.T) {
			err := &PropertyError{
				PropertyName:  "metadata.name",
				PropertyValue: propertyValueString(value),
				Errors: []*RuleError{
					{Message: "what a shame this happened"},
					{Message: "this is outrageous..."},
					{Message: "here's another error"},
				},
			}
			assert.EqualError(t, err, expectedErrorOutput(t, fmt.Sprintf("property_error_%s.txt", name)))
		})
	}
	t.Run("no name provided", func(t *testing.T) {
		err := &PropertyError{
			Errors: []*RuleError{
				{Message: "what a shame this happened"},
				{Message: "this is outrageous..."},
				{Message: "here's another error"},
			},
		}
		assert.EqualError(t, err, expectedErrorOutput(t, "property_error_no_name.txt"))
	})
}

func TestPropertyError_PrependPropertyName(t *testing.T) {
	for _, test := range []struct {
		PropertyError *PropertyError
		InputName     string
		ExpectedName  string
	}{
		{
			PropertyError: &PropertyError{},
		},
		{
			PropertyError: &PropertyError{PropertyName: "test"},
			ExpectedName:  "test",
		},
		{
			PropertyError: &PropertyError{},
			InputName:     "new",
			ExpectedName:  "new",
		},
		{
			PropertyError: &PropertyError{PropertyName: "original"},
			InputName:     "added",
			ExpectedName:  "added.original",
		},
	} {
		assert.Equal(t, test.ExpectedName, test.PropertyError.PrependPropertyName(test.InputName).PropertyName)
	}
}

func TestRuleError(t *testing.T) {
	for _, test := range []struct {
		RuleError    *RuleError
		InputCode    ErrorCode
		ExpectedCode ErrorCode
	}{
		{
			RuleError: NewRuleError("test"),
		},
		{
			RuleError:    NewRuleError("test", "code"),
			ExpectedCode: "code",
		},
		{
			RuleError:    NewRuleError("test"),
			InputCode:    "code",
			ExpectedCode: "code",
		},
		{
			RuleError:    NewRuleError("test", "original"),
			InputCode:    "added",
			ExpectedCode: "added:original",
		},
		{
			RuleError:    NewRuleError("test", "code-1", "code-2"),
			ExpectedCode: "code-2:code-1",
		},
		{
			RuleError:    NewRuleError("test", "original-1", "original-2"),
			InputCode:    "added",
			ExpectedCode: "added:original-2:original-1",
		},
	} {
		result := test.RuleError.AddCode(test.InputCode)
		assert.Equal(t, test.RuleError.Message, result.Message)
		assert.Equal(t, test.ExpectedCode, result.Code)
	}
}

func TestMultiRuleError(t *testing.T) {
	err := ruleSetError{
		errors.New("this is just a test!"),
		errors.New("another error..."),
		errors.New("that is just fatal."),
	}
	assert.EqualError(t, err, expectedErrorOutput(t, "multi_error.txt"))
}

func TestHasErrorCode(t *testing.T) {
	for _, test := range []struct {
		Error        error
		Code         ErrorCode
		HasErrorCode bool
	}{
		{
			Error:        nil,
			Code:         "code",
			HasErrorCode: false,
		},
		{
			Error:        errors.New("code"),
			Code:         "code",
			HasErrorCode: false,
		},
		{
			Error:        &RuleError{Code: "another"},
			Code:         "code",
			HasErrorCode: false,
		},
		{
			Error:        &RuleError{Code: "another:this"},
			Code:         "code",
			HasErrorCode: false,
		},
		{
			Error:        &RuleError{Code: "another:code:this"},
			Code:         "code",
			HasErrorCode: true,
		},
		{
			Error:        &PropertyError{Errors: []*RuleError{{Code: "another"}}},
			Code:         "code",
			HasErrorCode: false,
		},
		{
			Error:        &PropertyError{Errors: []*RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}}},
			Code:         "code",
			HasErrorCode: true,
		},
	} {
		assert.Equal(t, test.HasErrorCode, HasErrorCode(test.Error, test.Code))
	}
}

func expectedErrorOutput(t *testing.T, name string) string {
	t.Helper()
	data, err := errorsTestData.ReadFile(filepath.Join("test_data", name))
	require.NoError(t, err)
	return string(data)
}
