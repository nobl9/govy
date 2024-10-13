package rules

import (
	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govy"
)

// Required ensures the property's value is not empty (i.e. it's not its type's zero in).
func Required[T any]() govy.Rule[T] {
	return govy.NewRule(func(v T) error {
		if internal.IsEmpty(v) {
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
