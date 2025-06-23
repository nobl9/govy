package govy

// validationInterface is a common interface implemented by all validation entities.
// These include [Validator], [PropertyRules] and [Rule].
type validationInterface[T any] interface {
	Validate(v T) error
}

// validatorInterface defines validation entities which group properties,
// such as [Validator].
type validatorInterface[S any] interface {
	validationInterface[S]
	isValidator()
}

// propertyRulesInterface defines validation entities which describe properties,
// such as [PropertyRules], [PropertyRulesForSlice] and [PropertyRulesForMap].
//
// On top of [validationInterface] requirements it specifies internal functions
// which allow interacting with [propertyRulesInterface] instances like [PropertyRules]
// in an immutable fashion (no pointer receivers).
type propertyRulesInterface[T any] interface {
	validationInterface[T]
	cascadeInternal(mode CascadeMode) propertyRulesInterface[T]
	isPropertyRules()
}

// rulesInterface defines validation entities on the validation rule level,
// such as [Rule] or [RuleSet].
type rulesInterface[T any] interface {
	validationInterface[T]
	isRules()
}
