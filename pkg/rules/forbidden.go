package rules

import (
	"github.com/pkg/errors"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

func Forbidden[T any]() govy.SingleRule[T] {
	msg := "property is forbidden"
	return govy.NewSingleRule(func(v T) error {
		if internal.IsEmptyFunc(v) {
			return nil
		}
		return errors.New(msg)
	}).
		WithErrorCode(ErrorCodeForbidden).
		WithDescription(msg)
}
