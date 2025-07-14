// nolint: lll
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
			out: `Validation for ExpectedRuleError at index 0 has failed:
  - one of [code, containsMessage, message] properties must be set, none was provided`,
		},
		"invalid input - second error": {
			ok:             false,
			inputError:     &govy.ValidatorError{},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo"}, {}},
			out: `Validation for ExpectedRuleError at index 1 has failed:
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
			out: `Validation for ExpectedRuleError has failed:
  - one of [validatorIndex, validatorName] properties must be set, none was provided; The actual error was of type govy.ValidatorErrors.
    In order to match expected error with an actual error produced by a specific govy.Validator instance,
    either the name of the validator, its index (when using ValidateSlice method) or both must be provided.
    Otherwise the tests might produce ambiguous results.`,
		},
		"conflicting errors (names)": {
			ok:             false,
			inputError:     govy.ValidatorErrors{{Name: "foo"}, {Name: "bar"}, {Name: "foo"}},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo", ValidatorName: "this"}},
			out: `Multiple govy.ValidatorErrors errors seem to originate from the same govy.Validator instance.
This will lead to ambiguous test results, as govytest.ExpectedRuleError may be checked multiple times.
FIRST:
{
  "errors": null,
  "name": "foo"
}
SECOND:
{
  "errors": null,
  "name": "foo"
}`,
		},
		"conflicting errors (indexes)": {
			ok:             false,
			inputError:     govy.ValidatorErrors{{SliceIndex: ptr(1)}, {SliceIndex: ptr(1)}, {Name: "foo"}},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo", ValidatorName: "this"}},
			out: `Multiple govy.ValidatorErrors errors seem to originate from the same govy.Validator instance.
This will lead to ambiguous test results, as govytest.ExpectedRuleError may be checked multiple times.
FIRST:
{
  "errors": null,
  "sliceIndex": 1
}
SECOND:
{
  "errors": null,
  "sliceIndex": 1
}`,
		},
		"conflicting errors (names and indexes)": {
			ok: false,
			inputError: govy.ValidatorErrors{
				{Name: "foo", SliceIndex: ptr(1)},
				{Name: "bar"},
				{Name: "foo", SliceIndex: ptr(1)},
			},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo", ValidatorName: "this"}},
			out: `Multiple govy.ValidatorErrors errors seem to originate from the same govy.Validator instance.
This will lead to ambiguous test results, as govytest.ExpectedRuleError may be checked multiple times.
FIRST:
{
  "errors": null,
  "name": "foo",
  "sliceIndex": 1
}
SECOND:
{
  "errors": null,
  "name": "foo",
  "sliceIndex": 1
}`,
		},
		"conflicting errors (empty errors)": {
			ok:             false,
			inputError:     govy.ValidatorErrors{{}, {}},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo", ValidatorName: "this"}},
			out: `Multiple govy.ValidatorErrors errors seem to originate from the same govy.Validator instance.
This will lead to ambiguous test results, as govytest.ExpectedRuleError may be checked multiple times.
FIRST:
{
  "errors": null
}
SECOND:
{
  "errors": null
}`,
		},
		"expected error does not match any ValidatorErrors": {
			ok:             false,
			inputError:     govy.ValidatorErrors{{Name: "foo"}, {Name: "bar"}},
			expectedErrors: []govytest.ExpectedRuleError{{Message: "foo", ValidatorName: "baz"}},
			out: `govytest.ExpectedRuleError did not match any of the govy.ValidatorError.
govytest.ExpectedRuleError must match one of the govy.ValidatorError by either (or both, if both were provided):
- govytest.ExpectedRuleError.ValidatorName == govy.ValidatorError.Name
- govytest.ExpectedRuleError.ValidatorIndex == govy.ValidatorError.SliceIndex
EXPECTED:
{
  "propertyName": "",
  "message": "foo",
  "validatorName": "baz"
}
ACTUAL:
[
  {
    "errors": null,
    "name": "foo"
  },
  {
    "errors": null,
    "name": "bar"
  }
]`,
		},
		"match by name": {
			ok: true,
			inputError: govy.ValidatorErrors{
				{Name: "foo"},
				{Name: "bar", Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				}},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorName: "bar"},
			},
		},
		"match by index": {
			ok: true,
			inputError: govy.ValidatorErrors{
				{SliceIndex: ptr(0)},
				{SliceIndex: ptr(1), Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				}},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorIndex: ptr(1)},
			},
		},
		"match by name and index": {
			ok: true,
			inputError: govy.ValidatorErrors{
				{Name: "foo", SliceIndex: ptr(0)},
				{Name: "bar", SliceIndex: ptr(1), Errors: []*govy.PropertyError{
					{
						PropertyName: "this",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				}},
				{Name: "baz", SliceIndex: ptr(1), Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				}},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorName: "baz", ValidatorIndex: ptr(1)},
			},
		},
		"match by name - fail to match property error": {
			ok: false,
			inputError: govy.ValidatorErrors{
				{Name: "foo"},
				{Name: "bar", Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				}},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorName: "foo"},
			},
			out: `*govy.ValidatorError contains different number of errors than expected, expected: 1, actual: 0.`,
		},
		"match ValidatorError by name": {
			ok: true,
			inputError: &govy.ValidatorError{
				Name: "bar",
				Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorName: "bar"},
			},
		},
		"does not match ValidatorError by name": {
			ok: false,
			inputError: &govy.ValidatorError{
				Name: "bar",
				Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorName: "baz"},
			},
			out: "Expected name 'baz' of *govy.ValidatorError.Name but got 'bar'",
		},
		"match ValidatorError by index": {
			ok: true,
			inputError: &govy.ValidatorError{
				SliceIndex: ptr(1),
				Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorIndex: ptr(1)},
			},
		},
		"does not match ValidatorError by index": {
			ok: false,
			inputError: &govy.ValidatorError{
				SliceIndex: ptr(1),
				Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorIndex: ptr(2)},
			},
			out: "Expected index '2' of *govy.ValidatorError.SliceIndex but got '1'",
		},
		"match ValidatorError by name and index": {
			ok: true,
			inputError: &govy.ValidatorError{
				Name:       "bar",
				SliceIndex: ptr(1),
				Errors: []*govy.PropertyError{
					{
						PropertyName: "that",
						Errors:       []*govy.RuleError{{Message: "test"}},
					},
				},
			},
			expectedErrors: []govytest.ExpectedRuleError{
				{PropertyName: "that", Message: "test", ValidatorName: "bar", ValidatorIndex: ptr(1)},
			},
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
			out: `Validation for ExpectedRuleError has failed:
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

func ptr[T any](v T) *T { return &v }
