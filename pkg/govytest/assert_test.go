package govytest_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govytest"
)

func TestAssertNoError(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		mt := new(mockTestingT)
		ok := govytest.AssertNoError(mt, nil)
		assert.True(t, ok)
	})
	t.Run("generic error", func(t *testing.T) {
		mt := new(mockTestingT)
		ok := govytest.AssertNoError(mt, errors.New("this"))
		assert.False(t, ok)
		assert.Equal(t, "Received unexpected error:\nthis", mt.recordedError)
	})
	t.Run("validator error", func(t *testing.T) {
		mt := new(mockTestingT)
		ok := govytest.AssertNoError(mt, &govy.ValidatorError{Name: "Service"})
		assert.False(t, ok)
		assert.Equal(t, `Received unexpected error:
{
  "errors": null,
  "name": "Service"
}`, mt.recordedError)
	})
}

func TestAssertError(t *testing.T) {
	tests := map[string]struct {
		ok             bool
		inputError     error
		expectedErrors []govytest.ExpectedRuleError
		out            string
	}{
		"nil error": {
			ok:  false,
			out: "Input error should not be nil.",
		},
		"invalid input": {
			ok:             false,
			inputError:     &govy.ValidatorError{},
			expectedErrors: []govytest.ExpectedRuleError{{}},
			out: `Validation for ExpectedRuleError at index 0 has failed for the following properties:
  - one of [code, containsMessage, message] properties must be set, none was provided`,
		},
		"invalid input - second error": {
			ok:             false,
			inputError:     &govy.ValidatorError{},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo"}, {}},
			out: `Validation for ExpectedRuleError at index 1 has failed for the following properties:
  - one of [code, containsMessage, message] properties must be set, none was provided`,
		},
		"no expected errors": {
			ok:             false,
			inputError:     &govy.ValidatorError{},
			expectedErrors: []govytest.ExpectedRuleError{},
			out:            "[]govytest.ExpectedRuleError must not be empty.",
		},
		"wrong type of error": {
			ok:             false,
			inputError:     errors.New("foo!"),
			expectedErrors: []govytest.ExpectedRuleError{{PropertyName: "this", Message: "test"}},
			out: "Input error should be of type *govy.ValidatorError or govy.ValidatorErrors," +
				" but was of type *errors.errorString.\nError: foo!",
		},
		"errors count mismatch": {
			ok: false,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{Errors: []*govy.RuleError{{}, {}}},
			}},
			expectedErrors: []govytest.ExpectedRuleError{{PropertyName: "this", Message: "test"}},
			out:            "*govy.ValidatorError contains different number of errors than expected, expected: 1, actual: 2.",
		},
		"no matches": {
			ok: false,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test"}},
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", Message: "test"},
			},
			out: `Expected error was not found.
EXPECTED:
{
  "propertyName": "this",
  "message": "test"
}
ACTUAL:
[
  {
    "propertyName": "that",
    "errors": [
      {
        "error": "test"
      }
    ]
  }
]`,
		},
		"match on message": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}, {Message: "test1"}},
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", Message: "test1"},
				{PropertyName: "this", Message: "test2"},
				{PropertyName: "that", Message: "test3"},
			},
		},
		"match on code": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Code: "test3"}},
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Code: "test2"}, {Code: "test1"}},
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", Code: "test1"},
				{PropertyName: "this", Code: "test2"},
				{PropertyName: "that", Code: "test3"},
			},
		},
		"match on message contains": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}, {Message: "test1"}},
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", ContainsMessage: "test"},
				{PropertyName: "this", ContainsMessage: "test"},
				{PropertyName: "that", ContainsMessage: "test"},
			},
		},
		"match on message and code": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3", Code: "code3"}},
				},
				{
					PropertyName: "this",
					Errors: []*govy.RuleError{
						{Message: "test2", Code: "code2"},
						{Message: "test1", Code: "code1"},
					},
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", Message: "test1", Code: "code1"},
				{PropertyName: "this", Message: "test2", Code: "code2"},
				{PropertyName: "that", Message: "test3", Code: "code3"},
			},
		},
		"fail to match on message and code": {
			ok: false,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3", Code: "code3"}},
				},
				{
					PropertyName: "this",
					Errors: []*govy.RuleError{
						{Message: "test2", Code: "code2"},
						{Message: "test1", Code: "code1"},
					},
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", Message: "test1", Code: "code1"},
				{PropertyName: "this", Message: "test2", Code: "code2"},
				{PropertyName: "that", Message: "test3", Code: "code4"},
			},
			out: `Expected error was not found.
EXPECTED:
{
  "propertyName": "that",
  "code": "code4",
  "message": "test3"
}
ACTUAL:
[
  {
    "propertyName": "that",
    "errors": [
      {
        "error": "test3",
        "code": "code3"
      }
    ]
  },
  {
    "propertyName": "this",
    "errors": [
      {
        "error": "test2",
        "code": "code2"
      },
      {
        "error": "test1",
        "code": "code1"
      }
    ]
  }
]`,
		},
		"match on message, code and message contains": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3", Code: "code3"}},
				},
				{
					PropertyName: "this",
					Errors: []*govy.RuleError{
						{Message: "test2", Code: "code2"},
						{Message: "test1", Code: "code1"},
					},
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", Message: "test1", Code: "code1", ContainsMessage: "test"},
				{PropertyName: "this", Message: "test2", Code: "code2", ContainsMessage: "test"},
				{PropertyName: "that", Message: "test3", Code: "code3", ContainsMessage: "test"},
			},
		},
		"error was matched multiple times": {
			ok: false,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}},
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", ContainsMessage: "test"},
				{PropertyName: "this", ContainsMessage: "test"},
			},
			out: "Actual error was matched multiple times. Provide a more specific govytest.ExpectedRuleError list.",
		},
		"matched key error": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
					IsKeyError:   true,
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}},
					IsKeyError:   true,
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", Message: "test2", IsKeyError: true},
				{PropertyName: "that", Message: "test3", IsKeyError: true},
			},
		},
		"failed to match key error": {
			ok: false,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
					IsKeyError:   false,
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}},
					IsKeyError:   true,
				},
			}},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "this", Message: "test2", IsKeyError: true},
				{PropertyName: "that", Message: "test3", IsKeyError: true},
			},
			out: `Expected error was not found.
EXPECTED:
{
  "propertyName": "that",
  "message": "test3",
  "isKeyError": true
}
ACTUAL:
[
  {
    "propertyName": "that",
    "errors": [
      {
        "error": "test3"
      }
    ]
  },
  {
    "propertyName": "this",
    "isKeyError": true,
    "errors": [
      {
        "error": "test2"
      }
    ]
  }
]`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mt := new(mockTestingT)
			ok := govytest.AssertError(mt, tc.inputError, tc.expectedErrors...)
			if tc.ok {
				assert.True(t, ok)
			} else {
				assert.Require(t, assert.False(t, ok))
				assert.Equal(t, tc.out, mt.recordedError)
			}
		})
	}
}

func TestAssertError_ValidatorErrors(t *testing.T) {
	tests := map[string]struct {
		ok             bool
		inputError     error
		expectedErrors []govytest.ExpectedRuleError
		out            string
	}{
		"empty errors": {
			ok:             false,
			inputError:     govy.ValidatorErrors{},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo"}},
			out:            `govy.ValidatorErrors must not be empty.`,
		},
		"no index or name": {
			ok:             false,
			inputError:     govy.ValidatorErrors{{}},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo"}},
			out: `Validation for ExpectedRuleError has failed for the following properties:
  - one of [validatorIndex, validatorName] properties must be set, none was provided; 
    The actual error was of type govy.ValidatorErrors.
    in order to match expected error with an actual error produced by a specific govy.Validator instance,
    either the name of the validator, its index (when using ValidateSlice method) or both must be provided.
    Otherwise the tests might produce ambiguous results.`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mt := new(mockTestingT)
			ok := govytest.AssertError(mt, tc.inputError, tc.expectedErrors...)
			if tc.ok {
				assert.True(t, ok)
			} else {
				assert.Require(t, assert.False(t, ok))
				assert.Equal(t, tc.out, mt.recordedError)
			}
		})
	}
}

func TestAssertErrorContains(t *testing.T) {
	tests := map[string]struct {
		ok            bool
		inputError    error
		expectedError govytest.ExpectedRuleError
		out           string
	}{
		"nil error": {
			ok:  false,
			out: "Input error should not be nil.",
		},
		"invalid input": {
			ok:            false,
			inputError:    &govy.ValidatorError{},
			expectedError: govytest.ExpectedRuleError{},
			out: `Validation for ExpectedRuleError has failed for the following properties:
  - one of [code, containsMessage, message] properties must be set, none was provided`,
		},
		"wrong type of error": {
			ok:            false,
			inputError:    errors.New("foo!"),
			expectedError: govytest.ExpectedRuleError{PropertyName: "this", Message: "test"},
			out: "Input error should be of type *govy.ValidatorError or govy.ValidatorErrors," +
				" but was of type *errors.errorString.\nError: foo!",
		},
		"no matches": {
			ok: false,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test"}},
				},
			}},
			expectedError: govytest.ExpectedRuleError{PropertyName: "this", Message: "test"},
			out: `Expected error was not found.
EXPECTED:
{
  "propertyName": "this",
  "message": "test"
}
ACTUAL:
[
  {
    "propertyName": "that",
    "errors": [
      {
        "error": "test"
      }
    ]
  }
]`,
		},
		"match on message": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}, {Message: "test1"}},
				},
			}},
			expectedError: govytest.ExpectedRuleError{PropertyName: "this", Message: "test1"},
		},
		"match on code": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Code: "test3"}},
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Code: "test2"}, {Code: "test1"}},
				},
			}},
			expectedError: govytest.ExpectedRuleError{PropertyName: "this", Code: "test1"},
		},
		"match on message contains": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}, {Message: "test1"}},
				},
			}},
			expectedError: govytest.ExpectedRuleError{PropertyName: "this", ContainsMessage: "test"},
		},
		"match on message and code": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3", Code: "code3"}},
				},
				{
					PropertyName: "this",
					Errors: []*govy.RuleError{
						{Message: "test2", Code: "code2"},
						{Message: "test1", Code: "code1"},
					},
				},
			}},
			expectedError: govytest.ExpectedRuleError{PropertyName: "this", Message: "test1", Code: "code1"},
		},
		"fail to match on message and code": {
			ok: false,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3", Code: "code3"}},
				},
				{
					PropertyName: "this",
					Errors: []*govy.RuleError{
						{Message: "test2", Code: "code2"},
						{Message: "test1", Code: "code1"},
					},
				},
			}},
			expectedError: govytest.ExpectedRuleError{PropertyName: "that", Message: "test3", Code: "code4"},
			out: `Expected error was not found.
EXPECTED:
{
  "propertyName": "that",
  "code": "code4",
  "message": "test3"
}
ACTUAL:
[
  {
    "propertyName": "that",
    "errors": [
      {
        "error": "test3",
        "code": "code3"
      }
    ]
  },
  {
    "propertyName": "this",
    "errors": [
      {
        "error": "test2",
        "code": "code2"
      },
      {
        "error": "test1",
        "code": "code1"
      }
    ]
  }
]`,
		},
		"match on message, code and message contains": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3", Code: "code3"}},
				},
				{
					PropertyName: "this",
					Errors: []*govy.RuleError{
						{Message: "test2", Code: "code2"},
						{Message: "test1", Code: "code1"},
					},
				},
			}},
			expectedError: govytest.ExpectedRuleError{
				PropertyName:    "this",
				Message:         "test1",
				Code:            "code1",
				ContainsMessage: "test",
			},
		},
		"matched key error": {
			ok: true,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
					IsKeyError:   true,
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}},
					IsKeyError:   true,
				},
			}},
			expectedError: govytest.ExpectedRuleError{PropertyName: "that", Message: "test3", IsKeyError: true},
		},
		"failed to match key error": {
			ok: false,
			inputError: &govy.ValidatorError{Errors: []*govy.PropertyError{
				{
					PropertyName: "that",
					Errors:       []*govy.RuleError{{Message: "test3"}},
					IsKeyError:   false,
				},
				{
					PropertyName: "this",
					Errors:       []*govy.RuleError{{Message: "test2"}},
					IsKeyError:   true,
				},
			}},
			expectedError: govytest.ExpectedRuleError{PropertyName: "that", Message: "test3", IsKeyError: true},
			out: `Expected error was not found.
EXPECTED:
{
  "propertyName": "that",
  "message": "test3",
  "isKeyError": true
}
ACTUAL:
[
  {
    "propertyName": "that",
    "errors": [
      {
        "error": "test3"
      }
    ]
  },
  {
    "propertyName": "this",
    "isKeyError": true,
    "errors": [
      {
        "error": "test2"
      }
    ]
  }
]`,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			mt := new(mockTestingT)
			ok := govytest.AssertErrorContains(mt, tc.inputError, tc.expectedError)
			if tc.ok {
				assert.True(t, ok)
			} else {
				assert.Require(t, assert.False(t, ok))
				assert.Equal(t, tc.out, mt.recordedError)
			}
		})
	}
}

type mockTestingT struct {
	recordedError string
}

func (m *mockTestingT) Errorf(format string, args ...any) {
	m.recordedError = fmt.Sprintf(format, args...)
}

func (m *mockTestingT) Error(args ...any) {
	m.recordedError = fmt.Sprint(args...)
}

func (m *mockTestingT) Helper() {}
