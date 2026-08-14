package govy

import (
	"strings"

	"github.com/nobl9/govy/pkg/jsonpath"
)

// ValidatorInterface defines validation entities which group properties,
// such as [Validator].
type ValidatorInterface[T any] interface {
	validationInterface[T]
	isValidator()
}

// PropertyRulesInterface defines validation entities which describe properties,
// such as [PropertyRules], [PropertyRulesForSlice] and [PropertyRulesForMap].
//
// On top of [validationInterface] requirements it specifies internal functions
// which allow interacting with [PropertyRulesInterface] instances like [PropertyRules]
// in an immutable fashion (no pointer receivers).
type PropertyRulesInterface[T any] interface {
	validationInterface[T]
	cascadeInternal(mode CascadeMode) PropertyRulesInterface[T]
	getPath() jsonpath.Path
	inferPathModeInternal(mode InferPathMode) PropertyRulesInterface[T]
	isPropertyRules()
}

// RulesInterface defines validation entities on the validation rule level,
// such as [Rule] or [RuleSet].
type RulesInterface[T any] interface {
	validationInterface[T]
	isRules()
}

// validationInterface is a common interface implemented by all validation entities.
// These include [Validator], [PropertyRules] and [Rule].
type validationInterface[T any] interface {
	Validate(v T, opts ...ValidationOption) error
}

// validationOptions defines optional configuration passed to [validationInterface.Validate] invocations.
type validationOptions struct {
	hideValue bool
}

func newValidationOptions(opts ...ValidationOption) validationOptions {
	vo := validationOptions{}
	for _, opt := range opts {
		vo = opt(vo)
	}
	return vo
}

// ValidationOption applies optional configuration to [Rule.Validate].
type ValidationOption func(options validationOptions) validationOptions

func hideValue() ValidationOption {
	return func(options validationOptions) validationOptions {
		options.hideValue = true
		return options
	}
}

func hideStringValue(message, stringValue string) string {
	if strings.TrimSpace(stringValue) == "" {
		return message
	}
	return strings.ReplaceAll(message, stringValue, hiddenValue)
}
