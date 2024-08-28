package govy

import (
	"fmt"
)

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
			ev.Details = r.details
			ev.Description = r.description
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
				Message:     msg,
				Code:        r.errorCode,
				Details:     r.details,
				Description: r.description,
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
