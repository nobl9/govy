package govytest

import (
	"cmp"
	"encoding/json"
	"maps"
	"slices"
	"strconv"
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
	// It should be only left empty if the validated property has no name.
	PropertyName string `json:"propertyName"`
	// Optional. Matched against [govy.RuleError.Code].
	Code govy.ErrorCode `json:"code,omitempty"`
	// Optional. Matched against [govy.RuleError.Message].
	Message string `json:"message,omitempty"`
	// Optional. Matched against [govy.RuleError.Message] (partial).
	ContainsMessage string `json:"containsMessage,omitempty"`
	// Optional. Matched against [govy.PropertyError.IsKeyError].
	IsKeyError bool `json:"isKeyError,omitempty"`

	// Optional. Matched against [govy.ValidatorError.Name].
	ValidatorName string `json:"validatorName,omitempty"`
	// Optional. Matched against [govy.ValidatorError.SliceIndex].
	ValidatorIndex *int `json:"validatorIndex,omitempty"`
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

// expectedRuleErrorValidationForValidatorErrors defines the validation rules for [ExpectedRuleError]
// when asserting errors in [govy.ValidatorErrors] slice.
var expectedRuleErrorValidationForValidatorErrors = govy.New(
	govy.For(govy.GetSelf[ExpectedRuleError]()).
		Rules(rules.OneOfProperties(map[string]func(e ExpectedRuleError) any{
			"validatorName":  func(e ExpectedRuleError) any { return e.ValidatorName },
			"validatorIndex": func(e ExpectedRuleError) any { return e.ValidatorIndex },
		}).
			WithDetails(
				"The actual error was of type %T."+
					"\n  In order to match expected error with an actual error"+
					" produced by a specific govy.Validator instance,"+
					"\n  either the name of the validator, its index (when using ValidateSlice method) or both must be provided."+
					"\n  Otherwise the tests might produce ambiguous results.",
				govy.ValidatorErrors{},
			)),
	govy.ForPointer(func(e ExpectedRuleError) *int { return e.ValidatorIndex }).
		Rules(rules.GTE(0)),
).InferName()

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
// If the actual error is of type [govy.ValidatorErrors] either one of the following two fields must be provided.
// Otherwise, there's no way to compare the expected and actual errors.
//   - [ExpectedRuleError.ValidatorName] is equal to [govy.ValidatorError.Name]
//   - [ExpectedRuleError.ValidatorIndex] is equal to [govy.ValidatorError.SliceIndex]
//
// In the above case, all [ExpectedRuleError] are aggregated per matching [*govy.ValidatorError]
// and the function runs for every aggregated validator recursively.
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
// If the actual error is of type [govy.ValidatorErrors] the following two fields are also matched:
//   - [ExpectedRuleError.ValidatorName] is equal to [govy.ValidatorError.Name]
//   - [ExpectedRuleError.ValidatorIndex] is equal to [govy.ValidatorError.SliceIndex]
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

	if err == nil {
		t.Errorf("Input error should not be nil.")
		return false
	}

	if !validateExpectedErrors(t, countErrors, expectedErrors...) {
		return false
	}

	switch v := err.(type) {
	case *govy.ValidatorError:
		return assertValidatorError(t, countErrors, v, expectedErrors)
	case govy.ValidatorErrors:
		return assertValidatorErrors(t, countErrors, v, expectedErrors)
	default:
		t.Errorf(
			"Input error should be of type %[1]T or %[2]T, but was of type %[3]T.\nError: %[3]v",
			&govy.ValidatorError{},
			govy.ValidatorErrors{},
			err,
		)
		return false
	}
}

func validateExpectedErrors(t testingT, countErrors bool, expectedErrors ...ExpectedRuleError) bool {
	t.Helper()
	if len(expectedErrors) == 0 {
		t.Errorf("%T must not be empty.", expectedErrors)
		return false
	}
	switch countErrors {
	case true:
		if err := expectedRuleErrorValidation.ValidateSlice(expectedErrors); err != nil {
			t.Error(err.Error())
			return false
		}
	case false:
		if err := expectedRuleErrorValidation.Validate(expectedErrors[0]); err != nil {
			t.Error(err.Error())
			return false
		}
	}
	return true
}

type validatorKey struct {
	Name  string
	Index int
}

func (v validatorKey) Compare(v2 validatorKey) int {
	return cmp.Compare(v.Name+strconv.Itoa(v.Index), v2.Name+strconv.Itoa(v2.Index))
}

func validatorKeyFunc(name string, indexPtr *int) validatorKey {
	index := -1
	if indexPtr != nil {
		index = *indexPtr
	}
	return validatorKey{name, index}
}

func assertValidatorErrors(
	t testingT,
	countErrors bool,
	validatorErrors govy.ValidatorErrors,
	expectedErrors []ExpectedRuleError,
) bool {
	t.Helper()

	if len(validatorErrors) == 0 {
		t.Errorf("%T must not be empty.", validatorErrors)
		return false
	}

	for _, expected := range expectedErrors {
		if err := expectedRuleErrorValidationForValidatorErrors.Validate(expected); err != nil {
			t.Error(err.Error())
			return false
		}
	}

	validators := make(map[validatorKey]*govy.ValidatorError, len(validatorErrors))
	for _, err := range validatorErrors {
		key := validatorKeyFunc(err.Name, err.SliceIndex)
		if previous, ok := validators[key]; ok {
			t.Errorf("Multiple %T errors seem to originate from the same govy.Validator instance."+
				"\nThis will lead to ambiguous test results, as %T may be checked multiple times."+
				"\nFIRST:\n%s"+
				"\nSECOND:\n%s",
				govy.ValidatorErrors{}, ExpectedRuleError{}, mustEncodeJSON(previous), mustEncodeJSON(err))
			return false
		}
		validators[key] = err
	}

	expectedErrorsPerValidator := make(map[validatorKey][]ExpectedRuleError)
	for _, err := range expectedErrors {
		key := validatorKeyFunc(err.ValidatorName, err.ValidatorIndex)
		if _, ok := validators[key]; !ok {
			t.Errorf("%[1]T did not match any of the %[2]T."+
				"\n%[1]T must match one of the %[2]T by either (or both, if both were provided):"+
				"\n- %[1]T.ValidatorName == %[2]T.Name"+
				"\n- %[1]T.ValidatorIndex == %[2]T.SliceIndex"+
				"\nEXPECTED:\n%[3]s"+
				"\nACTUAL:\n%[4]s",
				err, *validatorErrors[0], mustEncodeJSON(err), mustEncodeJSON(validatorErrors))
			return false
		}
		expectedErrorsPerValidator[key] = append(expectedErrorsPerValidator[key], err)
	}

	if len(expectedErrorsPerValidator) == 1 {
		for k, v := range expectedErrorsPerValidator {
			return assertError(t, countErrors, validators[k], v...)
		}
	}

	passed := true
	keys := slices.SortedFunc(
		maps.Keys(expectedErrorsPerValidator),
		func(a, b validatorKey) int { return a.Compare(b) },
	)
	for _, k := range keys {
		ok := assertError(t, countErrors, validators[k], expectedErrorsPerValidator[k]...)
		if !ok {
			passed = false
		}
	}

	return passed
}

func assertValidatorError(
	t testingT,
	countErrors bool,
	validatorErr *govy.ValidatorError,
	expectedErrors []ExpectedRuleError,
) bool {
	t.Helper()

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

	if expected.ValidatorName != "" && expected.ValidatorName != validatorErr.Name {
		t.Errorf("Expected name '%s' of %T.Name but got '%s'", expected.ValidatorName, validatorErr, validatorErr.Name)
		return false
	}
	if expected.ValidatorIndex != nil {
		switch {
		case validatorErr.SliceIndex == nil:
			t.Errorf(
				"Expected index '%d' of %T.SliceIndex but got no index",
				*expected.ValidatorIndex,
				validatorErr,
			)
			return false
		case *expected.ValidatorIndex != *validatorErr.SliceIndex:
			t.Errorf(
				"Expected index '%d' of %T.SliceIndex but got '%d'",
				*expected.ValidatorIndex,
				validatorErr,
				*validatorErr.SliceIndex,
			)
			return false
		}
	}

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
		t.Errorf("Actual error was matched multiple times. Provide a more specific %T list.", expected)
		return false
	}
	t.Errorf("Expected error was not found."+
		"\nEXPECTED:\n%s"+
		"\nACTUAL:\n%s",
		mustEncodeJSON(expected), mustEncodeJSON(validatorErr.Errors))
	return false
}

func mustEncodeJSON(v any) string {
	data, _ := json.MarshalIndent(v, "", "  ")
	return string(data)
}
