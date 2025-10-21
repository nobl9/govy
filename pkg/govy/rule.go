package govy

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/internal/messagetemplates"
)

// NewRule creates a new [Rule] instance.
func NewRule[T any](validate func(v T) error) Rule[T] {
	return Rule[T]{validate: validate}
}

// RuleToPointer converts an existing [Rule] to its pointer variant.
// It retains all the properties of the original [Rule],
// but modifies its type constraints to work with a pointer to the original type.
// If the validated value is nil, the validation will be skipped.
func RuleToPointer[T any](rule Rule[T]) Rule[*T] {
	return Rule[*T]{
		validate: func(v *T) error {
			if v == nil {
				return nil
			}
			return rule.validate(*v)
		},
		errorCode:       rule.errorCode,
		details:         rule.details,
		message:         rule.message,
		messageTemplate: rule.messageTemplate,
		examples:        rule.examples,
		description:     rule.description,
		planModifiers:   rule.planModifiers,
	}
}

// Rule is the basic validation building block.
// It evaluates the provided validation function and enhances it
// with optional [ErrorCode] and arbitrary details.
type Rule[T any] struct {
	validate        func(v T) error
	errorCode       ErrorCode
	details         string
	message         string
	messageTemplate *template.Template
	examples        []string
	description     string
	planModifiers   []RulePlanModifier
}

// Validate runs validation function on the provided value.
// It can handle different types of errors returned by the function:
//   - [*RuleError], which details and [ErrorCode] are optionally extended with the ones defined by [Rule].
//   - [*PropertyError], for each of its errors their [ErrorCode] is extended with the one defined by [Rule].
//   - [RuleErrorTemplate], if message template was set with [Rule.WithMessageTemplate] or
//     [Rule.WithMessageTemplateString] then the [RuleError.Message] is constructed from the provided template
//     using variables passed inside [RuleErrorTemplate.vars].
//
// By default, it will construct a new [*RuleError].
func (r Rule[T]) Validate(v T) error {
	if err := r.validate(v); err != nil {
		switch ev := err.(type) {
		case *RuleError:
			if len(r.message) > 0 {
				ev.Message = createErrorMessage(r.message, r.details, r.examples)
			}
			ev.Description = r.description
			return ev.AddCode(r.errorCode)
		case *PropertyError:
			for _, e := range ev.Errors {
				_ = e.AddCode(r.errorCode)
			}
			return ev
		case RuleErrorTemplate:
			if r.message != "" {
				break
			}
			if r.messageTemplate == nil {
				panic(fmt.Sprintf("rule returned %T error but message template is not set", ev))
			}
			ev.vars.PropertyValue = v
			ev.vars.Details = r.details
			ev.vars.Examples = r.examples
			var buf bytes.Buffer
			if err = r.messageTemplate.Execute(&buf, ev.vars); err != nil {
				panic(fmt.Sprintf("failed to execute message template: %s", err))
			}
			return &RuleError{
				Message:     buf.String(),
				Code:        r.errorCode,
				Description: r.description,
			}
		}
		msg := err.Error()
		if len(r.message) > 0 {
			msg = r.message
		}
		return &RuleError{
			Message:     createErrorMessage(msg, r.details, r.examples),
			Code:        r.errorCode,
			Description: r.description,
		}
	}
	return nil
}

// WithErrorCode sets the error code for the returned [RuleError].
func (r Rule[T]) WithErrorCode(code ErrorCode) Rule[T] {
	r.errorCode = code
	return r
}

// WithMessage overrides the returned [RuleError] error message.
func (r Rule[T]) WithMessage(format string, a ...any) Rule[T] {
	r.messageTemplate = nil
	if len(a) == 0 {
		r.message = format
	} else {
		r.message = fmt.Sprintf(format, a...)
	}
	return r
}

// WithMessageTemplate overrides the returned [RuleError] error message using provided [template.Template].
func (r Rule[T]) WithMessageTemplate(tpl *template.Template) Rule[T] {
	r.messageTemplate = messagetemplates.AddFunctions(tpl)
	return r
}

// WithMessageTemplateString overrides the returned [RuleError] error message using provided template string.
// The string is parsed into [template.Template], it panics if any error is encountered during parsing.
func (r Rule[T]) WithMessageTemplateString(tplStr string) Rule[T] {
	tpl := messagetemplates.AddFunctions(template.New(""))
	tpl, err := tpl.Parse(tplStr)
	if err != nil {
		panic(fmt.Sprintf("failed to parse message template: %s", err))
	}
	return r.WithMessageTemplate(tpl)
}

// WithDetails adds details to the returned [RuleError] error message.
func (r Rule[T]) WithDetails(format string, a ...any) Rule[T] {
	r.details = fmt.Sprintf(format, a...)
	return r
}

// WithExamples adds examples to the returned [RuleError].
// Each example is converted to a string.
func (r Rule[T]) WithExamples(examples ...T) Rule[T] {
	r.examples = collections.ToStringSlice(examples)
	return r
}

// WithPlanModifiers adds [RulePlanModifier] which allow modifying [RulePlan] calculated for this [Rule].
func (r Rule[T]) WithPlanModifiers(mods ...RulePlanModifier) Rule[T] {
	r.planModifiers = append(r.planModifiers, mods...)
	return r
}

// WithDescription adds a custom description to the rule.
// It is used to enhance the [RulePlan], but otherwise does not appear in standard [RuleError.Error] output.
func (r Rule[T]) WithDescription(description string) Rule[T] {
	r.description = description
	return r
}

// RulePlanModifier allows modifying [RulePlan] calculated when calling [Plan].
type RulePlanModifier func(plan RulePlan) RulePlan

// RulePlanModifierValidValues adds valid values associated with the given [RulePlan].
// These values are not directly availabile through [RulePlan], rather
// they are aggregated and an intersection is calculated for [PropertyPlan].
func RulePlanModifierValidValues[T any](values ...T) RulePlanModifier {
	return func(plan RulePlan) RulePlan {
		plan.values = collections.ToStringSlice(values)
		return plan
	}
}

func (r Rule[T]) plan(builder planBuilder) {
	rulePlan := RulePlan{
		ErrorCode:   r.errorCode,
		Details:     r.details,
		Description: r.description,
		Conditions:  builder.rulePlan.Conditions,
		Examples:    r.examples,
	}
	for _, mod := range r.planModifiers {
		rulePlan = mod(rulePlan)
	}
	builder.rulePlan = rulePlan
	*builder.path = append(*builder.path, builder)
}

func createErrorMessage(message, details string, examples []string) string {
	if message == "" {
		return details
	}
	message += examplesToString(examples)
	if details == "" {
		return message
	}
	return message + "; " + details
}

func examplesToString(examples []string) string {
	if len(examples) == 0 {
		return ""
	}
	b := strings.Builder{}
	b.WriteString(" (e.g. ")
	internal.PrettyStringListBuilder(&b, examples, "'")
	b.WriteString(")")
	return b.String()
}

// isRules implements [rulesInterface].
func (r Rule[T]) isRules() {}
