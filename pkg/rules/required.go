package rules

import (
	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

func Required[T any]() govy.SingleRule[T] {
	return govy.NewSingleRule(func(v T) error {
		if internal.IsEmptyFunc(v) {
			return govy.NewRuleError(
				internal.RequiredErrorMessage,
				ErrorCodeRequired,
			)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeRequired).
		WithDescription("property is required")
}
