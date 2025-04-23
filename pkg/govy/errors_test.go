package govy_test

import (
	"embed"
	"errors"
	"fmt"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

//go:embed test_data
var errorsTestData embed.FS

func TestValidatorError(t *testing.T) {
	tests := map[string]*govy.ValidatorError{
		"no_name": {
			Errors: govy.PropertyErrors{
				{
					PropertyName:  "this",
					PropertyValue: "123",
					Errors:        []*govy.RuleError{{Message: "this is an error"}},
				},
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "that is an error"}},
				},
			},
		},
		"with_name": {
			Name: "Teacher",
			Errors: govy.PropertyErrors{
				{
					PropertyName:  "this",
					PropertyValue: "123",
					Errors:        []*govy.RuleError{{Message: "this is an error"}},
				},
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "that is an error"}},
				},
			},
		},
		"prop_no_name": {
			Errors: govy.PropertyErrors{
				{
					Errors: []*govy.RuleError{{Message: "no name"}},
				},
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "that is an error"}},
				},
			},
		},
	}

	for name, err := range tests {
		t.Run(name, func(t *testing.T) {
			assert.EqualError(t, err, expectedErrorOutput(t, fmt.Sprintf("validator_error_%s.txt", name)))
		})
	}
}

func TestValidatorErrors(t *testing.T) {
	err := govy.ValidatorErrors{
		{
			Name: "Teacher",
			Errors: govy.PropertyErrors{
				{
					PropertyName:  "this",
					PropertyValue: "123",
					Errors:        []*govy.RuleError{{Message: "this is an error"}},
				},
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "that is an error"}},
				},
			},
			SliceIndex: ptr(0),
		},
		{
			Errors: govy.PropertyErrors{
				{
					PropertyName:  "this",
					PropertyValue: "123",
					Errors:        []*govy.RuleError{{Message: "this is an error"}},
				},
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "that is an error"}},
				},
			},
			SliceIndex: ptr(1),
		},
		{
			Errors: govy.PropertyErrors{
				{
					Errors: []*govy.RuleError{{Message: "no name"}},
				},
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "that is an error"}},
				},
			},
			SliceIndex: ptr(3),
		},
	}
	assert.EqualError(t, err, expectedErrorOutput(t, "validator_errors.txt"))
}

func TestNewPropertyError(t *testing.T) {
	t.Run("string value", func(t *testing.T) {
		err := govy.NewPropertyError("name", "value",
			&govy.RuleError{Message: "top", Code: "1"},
			govy.RuleSetError{
				&govy.RuleError{Message: "rule1", Code: "2"},
				&govy.RuleError{Message: "rule2", Code: "3"},
			},
			&govy.RuleError{Message: "top", Code: "4"},
		)
		assert.Equal(t, &govy.PropertyError{
			PropertyName:  "name",
			PropertyValue: "value",
			Errors: []*govy.RuleError{
				{Message: "top", Code: "1"},
				{Message: "rule1", Code: "2"},
				{Message: "rule2", Code: "3"},
				{Message: "top", Code: "4"},
			},
		}, err)
	})

	tests := map[string]struct {
		InputValue    any
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
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := govy.NewPropertyError(
				"name",
				tc.InputValue,
				&govy.RuleError{Message: "msg"})
			assert.Equal(t, &govy.PropertyError{
				PropertyName:  "name",
				PropertyValue: tc.ExpectedValue,
				Errors:        []*govy.RuleError{{Message: "msg"}},
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
	tests := map[string]any{
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
	}

	for name, value := range tests {
		t.Run(name, func(t *testing.T) {
			err := &govy.PropertyError{
				PropertyName:  "metadata.name",
				PropertyValue: internal.PropertyValueString(value),
				Errors: []*govy.RuleError{
					{Message: "what a shame this happened"},
					{Message: "this is outrageous..."},
					{Message: "here's another error"},
				},
			}
			assert.EqualError(t, err, expectedErrorOutput(t, fmt.Sprintf("property_error_%s.txt", name)))
		})
	}

	t.Run("no name provided", func(t *testing.T) {
		err := &govy.PropertyError{
			Errors: []*govy.RuleError{
				{Message: "what a shame this happened"},
				{Message: "this is outrageous..."},
				{Message: "here's another error"},
			},
		}
		assert.EqualError(t, err, expectedErrorOutput(t, "property_error_no_name.txt"))
	})
}

func TestRuleError(t *testing.T) {
	tests := []struct {
		RuleError    *govy.RuleError
		InputCode    govy.ErrorCode
		ExpectedCode govy.ErrorCode
	}{
		{
			RuleError: govy.NewRuleError("test"),
		},
		{
			RuleError:    govy.NewRuleError("test", "code"),
			ExpectedCode: "code",
		},
		{
			RuleError:    govy.NewRuleError("test"),
			InputCode:    "code",
			ExpectedCode: "code",
		},
		{
			RuleError:    govy.NewRuleError("test", "original"),
			InputCode:    "added",
			ExpectedCode: "added:original",
		},
		{
			RuleError:    govy.NewRuleError("test", "code-1", "code-2"),
			ExpectedCode: "code-2:code-1",
		},
		{
			RuleError:    govy.NewRuleError("test", "original-1", "original-2"),
			InputCode:    "added",
			ExpectedCode: "added:original-2:original-1",
		},
	}

	for _, tc := range tests {
		result := tc.RuleError.AddCode(tc.InputCode)
		assert.Equal(t, tc.RuleError.Message, result.Message)
		assert.Equal(t, tc.ExpectedCode, result.Code)
	}
}

func TestMultiRuleError(t *testing.T) {
	err := govy.RuleSetError{
		errors.New("this is just a test!"),
		errors.New("another error..."),
		errors.New("that is just fatal."),
	}
	assert.EqualError(t, err, expectedErrorOutput(t, "multi_error.txt"))
}

func TestHasErrorCode(t *testing.T) {
	tests := []struct {
		Error        error
		Code         govy.ErrorCode
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
			Error:        &govy.RuleError{Code: "another"},
			Code:         "code",
			HasErrorCode: false,
		},
		{
			Error:        &govy.RuleError{Code: "another:this"},
			Code:         "code",
			HasErrorCode: false,
		},
		{
			Error:        &govy.RuleError{Code: "another:code:this"},
			Code:         "code",
			HasErrorCode: true,
		},
		{
			Error:        govy.RuleSetError{&govy.RuleError{Code: "another:code:this"}},
			Code:         "code",
			HasErrorCode: true,
		},
		{
			Error: govy.RuleSetError{
				&govy.RuleError{Code: "another:this"},
				&govy.RuleError{Code: "another:code:this"},
			},
			Code:         "code",
			HasErrorCode: true,
		},
		{
			Error:        govy.RuleSetError{&govy.RuleError{Code: "code:this"}},
			Code:         "that",
			HasErrorCode: false,
		},
		{
			Error:        &govy.PropertyError{Errors: []*govy.RuleError{{Code: "another"}}},
			Code:         "code",
			HasErrorCode: false,
		},
		{
			Error: &govy.PropertyError{
				Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}},
			},
			Code:         "code",
			HasErrorCode: true,
		},
		{
			Error: govy.PropertyErrors{
				{Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}}},
			},
			Code:         "code",
			HasErrorCode: true,
		},
		{
			Error: govy.PropertyErrors{
				{Errors: []*govy.RuleError{{Code: "this"}, {}, {Code: "this:code"}}},
				{Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "code:this"}}},
			},
			Code:         "another",
			HasErrorCode: true,
		},
		{
			Error: govy.PropertyErrors{
				{Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}}},
			},
			Code:         "that",
			HasErrorCode: false,
		},
		{
			Error: &govy.ValidatorError{
				Errors: govy.PropertyErrors{
					{Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}}},
				},
			},
			Code:         "code",
			HasErrorCode: true,
		},
		{
			Error: &govy.ValidatorError{
				Errors: govy.PropertyErrors{
					{Errors: []*govy.RuleError{{Code: "this"}, {}, {Code: "this:code"}}},
					{Errors: []*govy.RuleError{{Code: "that:another"}, {}, {Code: "code:this"}}},
				},
			},
			Code:         "that",
			HasErrorCode: true,
		},
		{
			Error: &govy.ValidatorError{
				Errors: govy.PropertyErrors{
					{Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}}},
				},
			},
			Code:         "that",
			HasErrorCode: false,
		},
		{
			Error: govy.ValidatorErrors{
				{
					Errors: govy.PropertyErrors{
						{Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}}},
					},
				},
			},
			Code:         "that",
			HasErrorCode: false,
		},
		{
			Error: govy.ValidatorErrors{
				{
					Errors: govy.PropertyErrors{
						{Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}}},
						{Errors: []*govy.RuleError{{Code: "this"}, {}, {Code: "another:code:this"}}},
					},
				},
				{
					Errors: govy.PropertyErrors{
						{Errors: []*govy.RuleError{{Code: "this:another"}, {}, {Code: "another:code:this"}}},
						{Errors: []*govy.RuleError{{Code: "this:that"}, {}, {Code: "another:code:this"}}},
					},
				},
			},
			Code:         "that",
			HasErrorCode: true,
		},
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.HasErrorCode, govy.HasErrorCode(tc.Error, tc.Code))
		})
	}
}

func TestNewRuleErrorTemplate(t *testing.T) {
	rule := govy.NewRule(func(s string) error {
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			Custom: map[string]string{"This": "that"},
		})
	}).
		WithMessageTemplateString("{{ .Custom.This }} is an error!")

	err := rule.Validate("test")
	assert.EqualError(t, err, "that is an error!")
}

func TestRuleErrorTemplate_Error(t *testing.T) {
	err := govy.NewRuleErrorTemplate(govy.TemplateVars{
		Examples: []string{"this", "that"},
	})
	assert.EqualError(t, err, "govy.RuleErrorTemplate should not be used directly")
}

func expectedErrorOutput(t *testing.T, name string) string {
	t.Helper()
	data, err := errorsTestData.ReadFile(filepath.Join("test_data", name))
	assert.Require(t, assert.NoError(t, err))
	return string(data)
}
