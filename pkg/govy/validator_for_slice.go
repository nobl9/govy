package govy

// NewForSlice creates a new [ValidatorForSlice].
// It stores the provided [Validator] which is used by [ValidatorForSlice.Validate]
// to validate each slice element.
func NewForSlice[S any](validator Validator[S]) ValidatorForSlice[S] {
	return ValidatorForSlice[S]{validator: validator}
}

// ValidatorForSlice is used to validate a slice of values.
// The rules for each slice element are described by the underlying [Validator].
//
// Note: It is designed to be used for validating independent values.
// If you need to validate the slice itself, for instance, to check if it has at most N elements,
// you should use the [Validator] directly in tandem with [ForSlice] and [GetSelf].
type ValidatorForSlice[S any] struct {
	validator Validator[S]
}

// Validate validates all elements of the provided slice using the underlying [Validator].
// All [ValidatorError] returned by the underlying [Validator] will be aggregated and wrapped in [ValidatorErrors].
func (v ValidatorForSlice[S]) Validate(st []S) error {
	errs := make(ValidatorErrors, 0)
	for i, s := range st {
		if err := v.validator.Validate(s); err != nil {
			vErr, ok := err.(*ValidatorError)
			if !ok {
				return err
			}
			vErr.SliceIndex = &i
			errs = append(errs, vErr)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errs
}

// plan constructs a validation plan for all the properties of the underlying [Validator].
// It appends the path with '[*]' to indicate a slice validation.
func (v ValidatorForSlice[S]) plan(builder planBuilder) {
	builder = builder.appendPath("[*]")
	v.validator.plan(builder)
}
