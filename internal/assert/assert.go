package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Fail fails the test with the provided message.
func Fail(t testing.TB, msg string, a ...interface{}) bool {
	t.Helper()
	t.Errorf(msg, a...)
	return false
}

// Require fails the test if the provided boolean is false.
// It should be used in conjunction with assert functions.
// Example:
//
//	assert.Require(t, assert.AssertError(t, err))
func Require(t testing.TB, isPassing bool) {
	t.Helper()
	if !isPassing {
		t.FailNow()
	}
}

// Equal fails the test if the expected and actual values are not equal.
func Equal(t testing.TB, expected, actual interface{}) bool {
	t.Helper()
	if !areEqual(expected, actual) {
		return Fail(t, "Expected: %v\nActual: %v", expected, actual)
	}
	return true
}

// True fails the test if the actual value is not true.
func True(t testing.TB, actual bool) bool {
	t.Helper()
	if !actual {
		return Fail(t, "Should be true")
	}
	return true
}

// False fails the test if the actual value is not false.
func False(t testing.TB, actual bool) bool {
	t.Helper()
	if actual {
		return Fail(t, "Should be false")
	}
	return true
}

// Len fails the test if the object is not of the expected length.
func Len(t testing.TB, object interface{}, length int) bool {
	t.Helper()
	if actual := getLen(object); actual != length {
		return Fail(t, "Expected length: %d\nActual: %d", length, actual)
	}
	return true
}

// IsType fails the test if the object is not of the expected type.
// The expected type is specified using a type parameter.
func IsType[T any](t testing.TB, object interface{}) bool {
	t.Helper()
	switch object.(type) {
	case T:
		return true
	default:
		return Fail(t, "Expected type: %T\nActual: %T", *new(T), object)
	}
}

// Error fails the test if the error is nil.
func Error(t testing.TB, err error) bool {
	t.Helper()
	if err == nil {
		return Fail(t, "An error is expected but actual nil.")
	}
	return true
}

// NoError fails the test if the error is not nil.
func NoError(t testing.TB, err error) bool {
	t.Helper()
	if err != nil {
		return Fail(t, "Unexpected error:\n%+v", err)
	}
	return true
}

// EqualError fails the test if the expected error is not equal to the actual error message.
func EqualError(t testing.TB, err error, expected string) bool {
	t.Helper()
	if !Error(t, err) {
		return false
	}
	if err.Error() != expected {
		return Fail(t, "Expected error message: %q\nActual: %q", expected, err.Error())
	}
	return true
}

// ErrorContains fails the test if the expected error does not contain the provided string.
func ErrorContains(t testing.TB, err error, contains string) bool {
	t.Helper()
	if !Error(t, err) {
		return false
	}
	if !strings.Contains(err.Error(), contains) {
    return Fail(t, "Expected error message to contain: %q\nActual: %q", contains, err.Error())
	}
	return true
}

// ElementsMatch fails the test if the expected and actual slices do not have the same elements.
func ElementsMatch[T comparable](t testing.TB, expected, actual []T) bool {
	t.Helper()
	if len(expected) != len(actual) {
		return Fail(t, "Slices are not equal in length, expected: %d, actual: %d", len(expected), len(actual))
	}

	actualVisited := make([]bool, len(actual))
	for _, e := range expected {
		found := false
		for j, a := range actual {
			if actualVisited[j] {
				continue
			}
			if areEqual(e, a) {
				actualVisited[j] = true
				found = true
				break
			}
		}
		if !found {
			return Fail(t, "Expected element %v not found in actual slice", e)
		}
	}
	for i := range actual {
		if !actualVisited[i] {
			return Fail(t, "Unexpected element %v found in actual slice", actual[i])
		}
	}
	return true
}

// Panic checks that the function panics with the expected message.
func Panic(t testing.TB, f func(), expected string) (result bool) {
	t.Helper()

	defer func() {
		r := recover()
		if r == nil {
			result = Fail(t, "Function did not panic")
			return
		}
		result = Equal(t, expected, r)
	}()

	f()
	return false
}

func areEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	if !reflect.DeepEqual(expected, actual) {
		return false
	}
	return true
}

func getLen(v interface{}) int {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Map, reflect.String:
		return rv.Len()
	default:
		return -1
	}
}
