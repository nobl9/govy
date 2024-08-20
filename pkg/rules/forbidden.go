package rules

import (
	"errors"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

func Forbidden[T any]() govy.Rule[T] {
	msg := "property is forbidden"
	return govy.NewRule(func(v T) error {
		if internal.IsEmptyFunc(v) {
			return nil
		}
		return errors.New(msg)
	}).
		WithErrorCode(ErrorCodeForbidden).
		WithDescription(msg)
}
