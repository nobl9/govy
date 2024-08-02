package validation

// CascadeMode defines how validation should behave when an error is encountered.
type CascadeMode uint

const (
	// CascadeModeContinue will continue validation after first error.
	CascadeModeContinue CascadeMode = iota
	// CascadeModeStop will stop validation on first error encountered.
	CascadeModeStop
)
