package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Require fails the test if the provided boolean is false.
// It should be used in conjunction with assert functions.
// Example:
//
//	testutils.Require(t, testutils.AssertError(t, err))
func Require(t *testing.T, isPassing bool) {
	t.Helper()
	if !isPassing {
		t.FailNow()
	}
}

func Equal(t *testing.T, expected, actual interface{}) bool {
	t.Helper()
	if !areEqual(expected, actual) {
		return fail(t, "Expected %v, got %v", expected, actual)
	}
	return true
}

func True(t *testing.T, actual bool) bool {
	t.Helper()
	if !actual {
		return fail(t, "Should be true")
	}
	return true
}

func Len(t *testing.T, object interface{}, length int) bool {
	t.Helper()
	if actual := getLen(object); actual != length {
		return fail(t, "Expected length %d, got %d", length, actual)
	}
	return true
}

func IsType[T any](t *testing.T, object interface{}) bool {
	t.Helper()
	switch object.(type) {
	case T:
		return true
	default:
		return fail(t, "Expected type %T, got %T", *new(T), object)
	}
}

func Error(t *testing.T, err error) bool {
	t.Helper()
	if err == nil {
		return fail(t, "An error is expected but got nil.")
	}
	return true
}

func NoError(t *testing.T, err error) bool {
	t.Helper()
	if err != nil {
		return fail(t, "Unexpected error:\n%+v", err)
	}
	return true
}

func EqualError(t *testing.T, expected error, actual string) bool {
	t.Helper()
	if !Error(t, expected) {
		return false
	}
	if expected.Error() != actual {
		return fail(t, "Expected error message %q, got %q", expected.Error(), actual)
	}
	return true
}

func ErrorContains(t *testing.T, expected error, contains string) bool {
	t.Helper()
	if !Error(t, expected) {
		return false
	}
	if !strings.Contains(expected.Error(), contains) {
		return fail(t, "Expected error message to contain %q, got %q", contains, expected.Error())
	}
	return true
}

func ElementsMatch[T comparable](t *testing.T, expected, actual []T) bool {
	t.Helper()
	if len(expected) != len(actual) {
		return fail(t, "Slices are not equal in length, expected: %d, got: %d", len(expected), len(actual))
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
			return fail(t, "Expected element %v not found in actual slice", e)
		}
	}
	for i := range actual {
		if !actualVisited[i] {
			return fail(t, "Unexpected element %v found in actual slice", actual[i])
		}
	}
	return true
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

func fail(t *testing.T, msg string, a ...interface{}) bool {
	t.Helper()
	t.Errorf(msg, a...)
	return false
}
