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
		rules: rules,
	}
}

// RuleSet allows defining a [Rule] which aggregates multiple sub-rules.
type RuleSet[T any] struct {
	rules []Rule[T]
	mode  CascadeMode
}

// Validate works the same way as [Rule.Validate],
// except each aggregated rule is validated individually.
// The errors are aggregated and returned as a single [RuleSetError]
// which serves as a container for them.
func (r RuleSet[T]) Validate(v T) error {
	var errs RuleSetError
	for i := range r.rules {
		err := r.rules[i].Validate(v)
		if err == nil {
			continue
		}
		switch ev := err.(type) {
		case *RuleError, *PropertyError:
			errs = append(errs, ev)
		default:
			errs = append(errs, &RuleError{
				Message: ev.Error(),
			})
		}
		if r.mode == CascadeModeStop {
			break
		}
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// WithErrorCode sets the error code for each returned [RuleError].
func (r RuleSet[T]) WithErrorCode(code ErrorCode) RuleSet[T] {
	for i := range r.rules {
		r.rules[i].errorCode = r.rules[i].errorCode.Add(code)
	}
	return r
}

// Cascade sets the [CascadeMode] for the rule set,
// which controls the flow of evaluating the validation rules.
func (r RuleSet[T]) Cascade(mode CascadeMode) RuleSet[T] {
	r.mode = mode
	return r
}

func (r RuleSet[T]) plan(builder planBuilder) {
	for _, rule := range r.rules {
		rule.plan(builder)
	}
}

// isRules implements [rulesInterface].
func (r RuleSet[T]) isRules() {}
