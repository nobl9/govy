package govytest

import (
	"encoding/json"
	"strings"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

// testingT is an interface that is compatible with *testing.T.
// It is used to make the functions in this package testable.
type testingT interface {
	Errorf(format string, args ...any)
	Error(args ...any)
	Helper()
}

// ExpectedRuleError defines the expectations for the asserted error.
// Its fields are used to find and match an actual [govy.RuleError].
type ExpectedRuleError struct {
	// Optional. Matched against [govy.PropertyError.PropertyName].
	// It should be only left empty if the validate property has no name.
	PropertyName string `json:"propertyName"`
	// Optional. Matched against [govy.RuleError.Code].
	Code govy.ErrorCode `json:"code,omitempty"`
	// Optional. Matched against [govy.RuleError.Message].
	Message string `json:"message,omitempty"`
	// Optional. Matched against [govy.RuleError.Message] (partial).
	ContainsMessage string `json:"containsMessage,omitempty"`
	// Optional. Matched against [govy.PropertyError.IsKeyError].
	IsKeyError bool `json:"isKeyError,omitempty"`
}

// expectedRuleErrorValidation defines the validation rules for [ExpectedRuleError].
var expectedRuleErrorValidation = govy.New(
	govy.For(govy.GetSelf[ExpectedRuleError]()).
		Rules(rules.OneOfProperties(map[string]func(e ExpectedRuleError) any{
			"code":            func(e ExpectedRuleError) any { return e.Code },
			"message":         func(e ExpectedRuleError) any { return e.Message },
			"containsMessage": func(e ExpectedRuleError) any { return e.ContainsMessage },
		})),
).InferName()

// Validate checks if the [ExpectedRuleError] is valid.
func (e ExpectedRuleError) Validate() error {
	return expectedRuleErrorValidation.Validate(e)
}

// AssertNoError asserts that the provided error is nil.
// If the error is not nil and of type [govy.ValidatorError] it will try
// encoding it to JSON and pretty printing the encountered error.
//
// It returns true if the error is nil, false otherwise.
func AssertNoError(t testingT, err error) bool {
	t.Helper()
	if err == nil {
		return true
	}
	errMsg := err.Error()
	if vErr, ok := err.(*govy.ValidatorError); ok {
		encErr, _ := json.MarshalIndent(vErr, "", "  ")
		errMsg = string(encErr)
	}
	t.Errorf("Received unexpected error:\n%+s", errMsg)
	return false
}

// AssertError asserts that the given error has:
//   - type equal to [*govy.ValidatorError]
//   - the expected number of [govy.RuleError] equal to the number of provided [ExpectedRuleError]
//   - at least one error which matches each of the provided [ExpectedRuleError]
//
// [ExpectedRuleError] and actual error are considered equal if they have the same property name and:
//   - [ExpectedRuleError.Code] is equal to [govy.RuleError.Code]
//   - [ExpectedRuleError.Message] is equal to [govy.RuleError.Message]
//   - [ExpectedRuleError.ContainsMessage] is part of [govy.RuleError.Message]
//
// At least one of the above must be set for [ExpectedRuleError]
// and once set, it will need to match the actual error.
//
// If [ExpectedRuleError.IsKeyError] is provided it will be required to match
// the actual [govy.PropertyError.IsKeyError].
//
// It returns true if the error matches the expectations, false otherwise.
func AssertError(
	t testingT,
	err error,
	expectedErrors ...ExpectedRuleError,
) bool {
	t.Helper()
	return assertError(t, true, err, expectedErrors...)
}

// AssertErrorContains asserts that the given error has:
//   - type equal to [*govy.ValidatorError]
//   - at least one error which matches the provided [ExpectedRuleError]
//
// Unlike [AssertError], it checks only a single error.
// The actual error may contain other errors, If you want to match them all, use [AssertError].
//
// [ExpectedRuleError] and actual error are considered equal if they have the same property name and:
//   - [ExpectedRuleError.Code] is equal to [govy.RuleError.Code]
//   - [ExpectedRuleError.Message] is equal to [govy.RuleError.Message]
//   - [ExpectedRuleError.ContainsMessage] is part of [govy.RuleError.Message]
//
// At least one of the above must be set for [ExpectedRuleError]
// and once set, it will need to match the actual error.
//
// If [ExpectedRuleError.IsKeyError] is provided it will be required to match
// the actual [govy.PropertyError.IsKeyError].
//
// It returns true if the error matches the expectations, false otherwise.
func AssertErrorContains(
	t testingT,
	err error,
	expectedError ExpectedRuleError,
) bool {
	t.Helper()
	return assertError(t, false, err, expectedError)
}

func assertError(
	t testingT,
	countErrors bool,
	err error,
	expectedErrors ...ExpectedRuleError,
) bool {
	t.Helper()

	if !validateExpectedErrors(t, expectedErrors...) {
		return false
	}
	validatorErr, ok := assertValidatorError(t, err)
	if !ok {
		return false
	}
	if countErrors {
		if !assertErrorsCount(t, validatorErr, len(expectedErrors)) {
			return false
		}
	}
	matched := make(matchedErrors, len(expectedErrors))
	for _, expected := range expectedErrors {
		if !assertErrorMatches(t, validatorErr, expected, matched) {
			return false
		}
	}
	return true
}

func validateExpectedErrors(t testingT, expectedErrors ...ExpectedRuleError) bool {
	t.Helper()
	if len(expectedErrors) == 0 {
		t.Errorf("%T must not be empty.", expectedErrors)
		return false
	}
	for _, expected := range expectedErrors {
		if err := expected.Validate(); err != nil {
			t.Error(err.Error())
			return false
		}
	}
	return true
}

func assertValidatorError(t testingT, err error) (*govy.ValidatorError, bool) {
	t.Helper()

	if err == nil {
		t.Errorf("Input error should not be nil.")
		return nil, false
	}
	validatorErr, ok := err.(*govy.ValidatorError)
	if !ok {
		t.Errorf("Input error should be of type %T.", &govy.ValidatorError{})
	}
	return validatorErr, ok
}

func assertErrorsCount(
	t testingT,
	validatorErr *govy.ValidatorError,
	expectedErrorsCount int,
) bool {
	t.Helper()

	actualErrorsCount := 0
	for _, actual := range validatorErr.Errors {
		actualErrorsCount += len(actual.Errors)
	}
	if expectedErrorsCount != actualErrorsCount {
		t.Errorf("%T contains different number of errors than expected, expected: %d, actual: %d.",
			validatorErr, expectedErrorsCount, actualErrorsCount)
		return false
	}
	return true
}

type matchedErrors map[int]map[int]struct{}

func (m matchedErrors) Add(propertyErrorIdx, ruleErrorIdx int) bool {
	if _, ok := m[propertyErrorIdx]; !ok {
		m[propertyErrorIdx] = make(map[int]struct{})
	}
	_, ok := m[propertyErrorIdx][ruleErrorIdx]
	m[propertyErrorIdx][ruleErrorIdx] = struct{}{}
	return ok
}

func assertErrorMatches(
	t testingT,
	validatorErr *govy.ValidatorError,
	expected ExpectedRuleError,
	matched matchedErrors,
) bool {
	t.Helper()

	multiMatch := false
	for i, actual := range validatorErr.Errors {
		if actual.PropertyName != expected.PropertyName {
			continue
		}
		if expected.IsKeyError != actual.IsKeyError {
			continue
		}
		for j, actualRuleErr := range actual.Errors {
			actualMessage := actualRuleErr.Error()
			matchedCtr := 0
			if expected.Message == "" || expected.Message == actualMessage {
				matchedCtr++
			}
			if expected.ContainsMessage == "" ||
				strings.Contains(actualMessage, expected.ContainsMessage) {
				matchedCtr++
			}
			if expected.Code == "" ||
				expected.Code == actualRuleErr.Code ||
				govy.HasErrorCode(actualRuleErr, expected.Code) {
				matchedCtr++
			}
			if matchedCtr == 3 {
				if matched.Add(i, j) {
					multiMatch = true
					continue
				}
				return true
			}
		}
	}

	if multiMatch {
		t.Errorf("Actual error was matched multiple times. Consider providing a more specific %T list.", expected)
		return false
	}
	encExpected, _ := json.MarshalIndent(expected, "", "  ")
	encActual, _ := json.MarshalIndent(validatorErr.Errors, "", "  ")
	t.Errorf("Expected error was not found.\nEXPECTED:\n%s\nACTUAL:\n%s",
		string(encExpected), string(encActual))
	return false
}
