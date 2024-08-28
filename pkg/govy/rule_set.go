package govy

// NewRuleSet creates a new [RuleSet] instance.
func NewRuleSet[T any](rules ...validationInterface[T]) RuleSet[T] {
	return RuleSet[T]{rules: rules}
}

// RuleSet allows defining a [Rule] which aggregates multiple sub-rules.
type RuleSet[T any] struct {
	rules     []validationInterface[T]
	errorCode ErrorCode
}

// Validate works the same way as [Rule.Validate],
// except each aggregated rule is validated individually.
// The errors are aggregated and returned as a single error which serves as a container for them.
func (r RuleSet[T]) Validate(v T) error {
	var errs RuleSetError
	for i := range r.rules {
		if err := r.rules[i].Validate(v); err != nil {
			switch ev := err.(type) {
			case *RuleError:
				errs = append(errs, ev.AddCode(r.errorCode))
			case *PropertyError:
				for _, e := range ev.Errors {
					_ = e.AddCode(r.errorCode)
				}
				errs = append(errs, ev)
			default:
				errs = append(errs, &RuleError{
					Message: ev.Error(),
					Code:    r.errorCode,
				})
			}
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// WithErrorCode sets the error code for each returned [RuleError].
func (r RuleSet[T]) WithErrorCode(code ErrorCode) RuleSet[T] {
	r.errorCode = code
	return r
}

func (r RuleSet[T]) plan(builder planBuilder) {
	for _, rule := range r.rules {
		if p, ok := rule.(planner); ok {
			p.plan(builder)
		}
	}
}
