package rules

import (
	"errors"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

// Forbidden ensures the property's value is its type's zero value, i.e. it's empty.
func Forbidden[T any]() govy.Rule[T] {
	msg := "property is forbidden"
	return govy.NewRule(func(v T) error {
		if internal.IsEmpty(v) {
			return nil
		}
		return errors.New(msg)
	}).
		WithErrorCode(ErrorCodeForbidden).
		WithDescription(msg)
}
