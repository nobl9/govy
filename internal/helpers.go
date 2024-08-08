package internal

import "reflect"

const RequiredErrorMessage = "property is required but was empty"

const RequiredErrorCodeString = "required"

// IsEmptyFunc verifies if the value is zero value of its type.
func IsEmptyFunc(v interface{}) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == 0 || rv.IsZero()
}
