package validation

import (
	"reflect"
)

func Required[T any]() SingleRule[T] {
	return NewSingleRule(func(v T) error {
		if isEmptyFunc(v) {
			return NewRequiredError()
		}
		return nil
	}).
		WithErrorCode(ErrorCodeRequired).
		WithDescription("property is required")
}

// isEmptyFunc checks only the types which it makes sense for.
// It's hard to consider 0 an empty value for anything really.
func isEmptyFunc(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == 0 || rv.IsZero()
}
