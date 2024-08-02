package validation

import "github.com/pkg/errors"

func Forbidden[T any]() SingleRule[T] {
	msg := "property is forbidden"
	return NewSingleRule(func(v T) error {
		if isEmptyFunc(v) {
			return nil
		}
		return errors.New(msg)
	}).
		WithErrorCode(ErrorCodeForbidden).
		WithDescription(msg)
}
