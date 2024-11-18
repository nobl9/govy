package govy

// NewRuleSet creates a new [RuleSet] instance.
func NewRuleSet[T any](rules ...Rule[T]) RuleSet[T] {
	return RuleSet[T]{rules: rules}
}

// RuleSetToPointer converts an existing [RuleSet] to its pointer variant.
// It retains all the properties of the original [RuleSet],
// and modifies its type constraints to work with a pointer to the original type.
// It calls [RuleToPointer] for each of the underlying [Rule].
func RuleSetToPointer[T any](ruleSet RuleSet[T]) RuleSet[*T] {
	rules := make([]Rule[*T], 0, len(ruleSet.rules))
	for _, rule := range ruleSet.rules {
		rules = append(rules, RuleToPointer(rule))
	}
	return RuleSet[*T]{
		rules:     rules,
		errorCode: ruleSet.errorCode,
	}
}

// RuleSet allows defining a [Rule] which aggregates multiple sub-rules.
type RuleSet[T any] struct {
	rules     []Rule[T]
	errorCode ErrorCode
	examples  []string
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
		rule.plan(builder)
	}
}
