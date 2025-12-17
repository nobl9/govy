package govy

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
	getName() string
	inferNameModeInternal(mode InferNameMode) PropertyRulesInterface[T]
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
