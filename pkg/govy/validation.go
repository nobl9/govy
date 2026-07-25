package govy

import "github.com/nobl9/govy/pkg/jsonpath"

// ValidatorInterface defines validation entities which group properties,
// such as [Validator].
type ValidatorInterface[T any] interface {
	validationInterface[T]
	isValidator()
}

// PropertyRulesInterface defines validation entities which describe properties,
// such as [PropertyRules], [PropertyRulesForSlice] and [PropertyRulesForMap].
//
// GetID exposes the property identifier used by [Validator.RemovePropertiesByID].
// The remaining methods allow the package to interact with property rules
// in an immutable fashion without pointer receivers.
type PropertyRulesInterface[T any] interface {
	validationInterface[T]
	cascadeInternal(mode CascadeMode) PropertyRulesInterface[T]
	GetID() string
	getPath() jsonpath.Path
	inferPathInternal(mode InferPathMode) PropertyRulesInterface[T]
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
	Validate(v T) error
}
