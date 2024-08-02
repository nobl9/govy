package validation

import (
	"fmt"
)

// Rule is the interface for all validation rules.
type Rule[T any] interface {
	Validate(v T) error
}

// NewSingleRule creates a new [SingleRule] instance.
func NewSingleRule[T any](validate func(v T) error) SingleRule[T] {
	return SingleRule[T]{validate: validate}
}

// SingleRule is the basic validation building block.
// It evaluates the provided validation function and enhances it
// with optional [ErrorCode] and arbitrary details.
type SingleRule[T any] struct {
	validate    func(v T) error
	errorCode   ErrorCode
	details     string
	message     string
	description string
}

// Validate runs validation function on the provided value.
// It can handle different types of errors returned by the function:
//   - *[RuleError], which details and [ErrorCode] are optionally extended with the ones defined by [SingleRule].
//   - *[PropertyError], for each of its errors their [ErrorCode] is extended with the one defined by [SingleRule].
//
// By default, it will construct a new RuleError.
func (r SingleRule[T]) Validate(v T) error {
	if err := r.validate(v); err != nil {
		switch ev := err.(type) {
		case *RuleError:
			if len(r.message) > 0 {
				ev.Message = r.message
			}
			ev.Message = addDetailsToMessage(ev.Message, r.details)
			return ev.AddCode(r.errorCode)
		case *PropertyError:
			for _, e := range ev.Errors {
				_ = e.AddCode(r.errorCode)
			}
			return ev
		default:
			msg := ev.Error()
			if len(r.message) > 0 {
				msg = r.message
			}
			return &RuleError{
				Message: addDetailsToMessage(msg, r.details),
				Code:    r.errorCode,
			}
		}
	}
	return nil
}

// WithErrorCode sets the error code for the returned [RuleError].
func (r SingleRule[T]) WithErrorCode(code ErrorCode) SingleRule[T] {
	r.errorCode = code
	return r
}

// WithMessage overrides the returned [RuleError] error message with message.
func (r SingleRule[T]) WithMessage(format string, a ...any) SingleRule[T] {
	if len(a) == 0 {
		r.message = format
	} else {
		r.message = fmt.Sprintf(format, a...)
	}
	return r
}

// WithDetails adds details to the returned [RuleError] error message.
func (r SingleRule[T]) WithDetails(format string, a ...any) SingleRule[T] {
	if len(a) == 0 {
		r.details = format
	} else {
		r.details = fmt.Sprintf(format, a...)
	}
	return r
}

func (r SingleRule[T]) WithDescription(description string) SingleRule[T] {
	r.description = description
	return r
}

func (r SingleRule[T]) plan(builder planBuilder) {
	builder.rulePlan = RulePlan{
		ErrorCode:   r.errorCode,
		Details:     r.details,
		Description: r.description,
		Conditions:  builder.rulePlan.Conditions,
	}
	*builder.all = append(*builder.all, builder)
}

// NewRuleSet creates a new [RuleSet] instance.
func NewRuleSet[T any](rules ...Rule[T]) RuleSet[T] {
	return RuleSet[T]{rules: rules}
}

// RuleSet allows defining [Rule] which aggregates multiple sub-rules.
type RuleSet[T any] struct {
	rules     []Rule[T]
	errorCode ErrorCode
	details   string
}

// Validate works the same way as [SingleRule.Validate],
// except each aggregated rule is validated individually.
// The errors are aggregated and returned as a single error which serves as a container for them.
func (r RuleSet[T]) Validate(v T) error {
	var errs ruleSetError
	for i := range r.rules {
		if err := r.rules[i].Validate(v); err != nil {
			switch ev := err.(type) {
			case *RuleError:
				ev.Message = addDetailsToMessage(ev.Message, r.details)
				errs = append(errs, ev.AddCode(r.errorCode))
			case *PropertyError:
				for _, e := range ev.Errors {
					_ = e.AddCode(r.errorCode)
				}
				errs = append(errs, ev)
			default:
				errs = append(errs, &RuleError{
					Message: addDetailsToMessage(ev.Error(), r.details),
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

// WithDetails adds details to each returned [RuleError] error message.
func (r RuleSet[T]) WithDetails(format string, a ...any) RuleSet[T] {
	if len(a) == 0 {
		r.details = format
	} else {
		r.details = fmt.Sprintf(format, a...)
	}
	return r
}

func (r RuleSet[T]) plan(builder planBuilder) {
	for _, rule := range r.rules {
		if p, ok := rule.(planner); ok {
			p.plan(builder)
		}
	}
}

func addDetailsToMessage(msg, details string) string {
	if details == "" {
		return msg
	}
	if msg == "" {
		return details
	}
	return msg + "; " + details
}
