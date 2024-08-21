package govy

import (
	"fmt"

	"github.com/nobl9/govy/internal"
)

// ruleInterface is the interface for all validation rules.
type ruleInterface[T any] interface {
	Validate(v T) error
}

// NewRule creates a new [Rule] instance.
func NewRule[T any](validate func(v T) error) Rule[T] {
	return Rule[T]{validate: validate}
}

// Rule is the basic validation building block.
// It evaluates the provided validation function and enhances it
// with optional [ErrorCode] and arbitrary details.
type Rule[T any] struct {
	validate    func(v T) error
	errorCode   ErrorCode
	details     string
	message     string
	description string
}

// Validate runs validation function on the provided value.
// It can handle different types of errors returned by the function:
//   - *[RuleError], which details and [ErrorCode] are optionally extended with the ones defined by [Rule].
//   - *[PropertyError], for each of its errors their [ErrorCode] is extended with the one defined by [Rule].
//
// By default, it will construct a new RuleError.
func (r Rule[T]) Validate(v T) error {
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
func (r Rule[T]) WithErrorCode(code ErrorCode) Rule[T] {
	r.errorCode = code
	return r
}

// WithMessage overrides the returned [RuleError] error message with message.
func (r Rule[T]) WithMessage(format string, a ...any) Rule[T] {
	if len(a) == 0 {
		r.message = format
	} else {
		r.message = fmt.Sprintf(format, a...)
	}
	return r
}

// WithDetails adds details to the returned [RuleError] error message.
func (r Rule[T]) WithDetails(format string, a ...any) Rule[T] {
	if len(a) == 0 {
		r.details = format
	} else {
		r.details = fmt.Sprintf(format, a...)
	}
	return r
}

// WithDescription adds a custom description to the rule.
// It is used to enhance the [RulePlan], but otherwise does not appear in standard [RuleError.Error] output.
func (r Rule[T]) WithDescription(description string) Rule[T] {
	r.description = description
	return r
}

func (r Rule[T]) plan(builder planBuilder) {
	builder.rulePlan = RulePlan{
		ErrorCode:   r.errorCode,
		Details:     r.details,
		Description: r.description,
		Conditions:  builder.rulePlan.Conditions,
	}
	*builder.children = append(*builder.children, builder)
}

// NewRuleSet creates a new [RuleSet] instance.
func NewRuleSet[T any](rules ...ruleInterface[T]) RuleSet[T] {
	return RuleSet[T]{rules: rules}
}

// RuleSet allows defining a [Rule] which aggregates multiple sub-rules.
type RuleSet[T any] struct {
	rules     []ruleInterface[T]
	errorCode ErrorCode
	details   string
}

// Validate works the same way as [Rule.Validate],
// except each aggregated rule is validated individually.
// The errors are aggregated and returned as a single error which serves as a container for them.
func (r RuleSet[T]) Validate(v T) error {
	var errs internal.RuleSetError
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
