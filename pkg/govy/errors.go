package govy

import (
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/jsonpath"
	"github.com/nobl9/govy/internal/logging"
)

const (
	ErrorCodeSeparator     = string(errorCodeSeparatorRune)
	errorCodeSeparatorRune = ':'
	hiddenValue            = "[hidden]"
)

func NewValidatorError(errs PropertyErrors) *ValidatorError {
	return &ValidatorError{Errors: errs}
}

// ValidatorError is the top-level error type for validation errors.
// It aggregates the property errors of [Validator].
type ValidatorError struct {
	Errors PropertyErrors `json:"errors"`
	// Name is an optional name of the [Validator].
	Name string `json:"name,omitempty"`
	// SliceIndex is set if the error was created by [Validator.ValidateSlice].
	// It contains a 0-based index.
	SliceIndex *int `json:"sliceIndex,omitempty"`
}

// WithName sets the [ValidatorError.Name] field.
func (e *ValidatorError) WithName(name string) *ValidatorError {
	e.Name = name
	return e
}

// Error implements the error interface.
func (e *ValidatorError) Error() string {
	b := strings.Builder{}
	b.WriteString("Validation")
	if e.Name != "" {
		b.WriteString(" for ")
		b.WriteString(e.Name)
	}
	if e.SliceIndex != nil {
		b.WriteString(" at index ")
		b.WriteString(strconv.Itoa(*e.SliceIndex))
	}
	b.WriteString(" has failed")
	if e.hasAtLeastOnePropertyName() {
		b.WriteString(" for the following properties")
	}
	b.WriteString(":\n")
	internal.JoinErrors(&b, e.Errors, strings.Repeat(" ", 2))
	return b.String()
}

func (e *ValidatorError) hasAtLeastOnePropertyName() bool {
	for _, e := range e.Errors {
		if e.PropertyName != "" {
			return true
		}
	}
	return false
}

// ValidatorErrors is a slice of [ValidatorError].
type ValidatorErrors []*ValidatorError

// Error implements the error interface.
func (e ValidatorErrors) Error() string {
	b := strings.Builder{}
	for i, vErr := range e {
		b.WriteString(vErr.Error())
		if i < len(e)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

// PropertyErrors is a slice of [PropertyError].
type PropertyErrors []*PropertyError

// Error implements the error interface.
func (e PropertyErrors) Error() string {
	b := strings.Builder{}
	internal.JoinErrors(&b, e, "")
	return b.String()
}

// HideValue hides the property value from each of the [PropertyError].
func (e PropertyErrors) HideValue() PropertyErrors {
	for _, err := range e {
		_ = err.HideValue()
	}
	return e
}

// sort should be always called after aggregate.
func (e PropertyErrors) sort() PropertyErrors {
	if len(e) == 0 {
		return e
	}
	sort.Slice(e, func(i, j int) bool {
		e1, e2 := e[i], e[j]
		if e1.PropertyName != e2.PropertyName {
			return e1.PropertyName < e2.PropertyName
		}
		if e1.PropertyValue != e2.PropertyValue {
			return e1.PropertyValue < e2.PropertyValue
		}
		if e1.IsKeyError != e2.IsKeyError {
			return e1.IsKeyError
		}
		return e1.IsSliceElementError
	})
	return e
}

// aggregate merges [PropertyError] with according to the [PropertyError.Equal] comparison.
func (e PropertyErrors) aggregate() PropertyErrors {
	if len(e) == 0 {
		return nil
	}
	agg := make(PropertyErrors, 0, len(e))
outer:
	for _, e1 := range e {
		for _, e2 := range agg {
			if e1.Equal(e2) {
				e2.Errors = append(e2.Errors, e1.Errors...)
				continue outer
			}
		}
		agg = append(agg, e1)
	}
	return agg
}

// NewPropertyError constructs new [*PropertyError] instance.
// Property name is assumed to be a valid, escaped JSONPath.
func NewPropertyError(propertyName string, propertyValue any, errs ...error) *PropertyError {
	return &PropertyError{
		PropertyName:  propertyName,
		PropertyValue: internal.PropertyValueString(propertyValue),
		Errors:        unpackRuleErrors(errs, make([]*RuleError, 0, len(errs))),
	}
}

// PropertyError is the error returned by [PropertyRules.Validate].
// It contains property level details along with all the [RuleError] encountered for that property.
type PropertyError struct {
	// PropertyName is a string which should uniquely identify the property within a [Validator] instance.
	// Typically, it is in the form of a JSONPath, constructed from the root of the validated object.
	PropertyName string `json:"propertyName"`
	// PropertyValue is a string representation of the property's value.
	PropertyValue string `json:"propertyValue,omitempty"`
	// IsKeyError is set to true if the error was created through map key validation.
	// PropertyValue in this scenario will be the key value, equal to the last element of PropertyName path.
	IsKeyError bool `json:"isKeyError,omitempty"`
	// IsSliceElementError is set to true if the error was created through slice element validation.
	IsSliceElementError bool `json:"isSliceElementError,omitempty"`
	// Errors are all rule errors reported for this property.
	//
	// Note: You can have multiple [PropertyRules] with the same name and value,
	// in this scenario, all these instances will be aggregated into a single [PropertyError].
	// See [PropertyError.Equal] for details on the equality conditions for [PropertyError].
	Errors []*RuleError `json:"errors"`
}

// Error implements the error interface.
func (e *PropertyError) Error() string {
	b := new(strings.Builder)
	indent := ""
	if e.PropertyName != "" {
		fmt.Fprintf(b, "'%s'", e.PropertyName)
		if e.PropertyValue != "" {
			if e.IsKeyError {
				fmt.Fprintf(b, " with key '%s'", e.PropertyValue)
			} else {
				fmt.Fprintf(b, " with value '%s'", e.PropertyValue)
			}
		}
		b.WriteString(":\n")
		indent = strings.Repeat(" ", 2)
	}
	internal.JoinErrors(b, e.Errors, indent)
	return b.String()
}

// Equal checks if two [PropertyError] are equal.
func (e *PropertyError) Equal(cmp *PropertyError) bool {
	return e.PropertyName == cmp.PropertyName &&
		e.PropertyValue == cmp.PropertyValue &&
		e.IsKeyError == cmp.IsKeyError &&
		e.IsSliceElementError == cmp.IsSliceElementError
}

// HideValue hides the property value from each of the [PropertyError.Errors].
func (e *PropertyError) HideValue() *PropertyError {
	sv := internal.PropertyValueString(e.PropertyValue)
	e.PropertyValue = ""
	for _, err := range e.Errors {
		_ = err.HideValue(sv)
	}
	return e
}

// prependParentPropertyName prepends a given name to the [PropertyError.PropertyName].
// If the name prepended name is a JSONPath, it is assumed to be escaped.
func (e *PropertyError) prependParentPropertyName(name string) *PropertyError {
	switch {
	case e.IsSliceElementError && strings.HasPrefix(e.PropertyName, "["):
		e.PropertyName = jsonpath.JoinArray(name, e.PropertyName)
	default:
		e.PropertyName = jsonpath.Join(name, e.PropertyName)
	}
	return e
}

// NewRuleError creates a new [RuleError] with the given message and optional error codes.
// Error codes are added according to the rules defined by [RuleError.AddCode].
func NewRuleError(message string, codes ...ErrorCode) *RuleError {
	ruleError := &RuleError{Message: message}
	for _, code := range codes {
		ruleError = ruleError.AddCode(code)
	}
	return ruleError
}

// RuleError is the base error associated with a [Rule].
// It is returned by [Rule.Validate].
type RuleError struct {
	Message     string    `json:"error"`
	Code        ErrorCode `json:"code,omitempty"`
	Description string    `json:"description,omitempty"`
}

// Error implements the error interface.
// It simply returns the underlying [RuleError.Message].
func (r *RuleError) Error() string {
	return r.Message
}

// AddCode extends the [RuleError] with the given error code.
// See [ErrorCode.Add] for more details.
func (r *RuleError) AddCode(code ErrorCode) *RuleError {
	r.Code = r.Code.Add(code)
	return r
}

// HideValue replaces all occurrences of a string in the [RuleError.Message] with a '*' characters.
func (r *RuleError) HideValue(stringValue string) *RuleError {
	r.Message = strings.ReplaceAll(r.Message, stringValue, hiddenValue)
	return r
}

// RuleSetError is a container for transferring multiple errors reported by [RuleSet.Validate].
type RuleSetError []error

// Error implements the error interface.
func (r RuleSetError) Error() string {
	b := new(strings.Builder)
	internal.JoinErrors(b, r, "")
	return b.String()
}

// NewRuleErrorTemplate creates a new [RuleErrorTemplate] with the given template variables.
// The variables can be of any type, most commonly it would be a struct or a map.
// These variables are then passed to [template.Template.Execute].
// Example:
//
//	return govy.NewRuleErrorTemplate(map[string]string{
//	  "Name":      "my-property",
//	  "MaxLength": 2,
//	})
//
// For more details on Go templates see: https://pkg.go.dev/text/template.
func NewRuleErrorTemplate(vars TemplateVars) RuleErrorTemplate {
	return RuleErrorTemplate{vars: vars}
}

// RuleErrorTemplate is a container for passing template variables under the guise of an error.
// It's not meant to be used directly as an error but rather
// unpacked by [Rule] in order to create a templated error message.
type RuleErrorTemplate struct {
	vars TemplateVars
}

// Error implements the error interface.
// Since [RuleErrorTemplate] should not be used directly this function returns.
func (e RuleErrorTemplate) Error() string {
	return fmt.Sprintf("%T should not be used directly", e)
}

// TemplateVars lists all the possible variables that can be used by builtin rules' message templates.
// Reuse the variable names to keep the consistency across all the rules.
type TemplateVars struct {
	// Common variables which are available for all the rules.
	PropertyValue any
	Examples      []string
	Details       string

	// Builtin, widely used variables which are available only for selected rules.

	// Error is the dynamic error returned by underlying functions evaluated by the rule,
	// for instance an error returned by [net/url.Parse].
	Error string
	// ComparisonValue is the value defined most commonly during rule creation
	// to which runtime values are compared.
	ComparisonValue any
	MinLength       int
	MaxLength       int

	// Custom variables either provided by the user or case specific,
	// this can be anything, for instance map[string]any or a struct.
	Custom any
}

// HasErrorCode checks if an error contains given [ErrorCode].
// It supports all govy errors.
func HasErrorCode(err error, code ErrorCode) bool {
	switch v := err.(type) {
	case *ValidatorError:
		for _, e := range v.Errors {
			if HasErrorCode(e, code) {
				return true
			}
		}
		return false
	case ValidatorErrors:
		for _, e := range v {
			if HasErrorCode(e, code) {
				return true
			}
		}
		return false
	case PropertyErrors:
		for _, e := range v {
			if HasErrorCode(e, code) {
				return true
			}
		}
		return false
	case *PropertyError:
		for _, e := range v.Errors {
			if HasErrorCode(e, code) {
				return true
			}
		}
	case *RuleError:
		return v.Code.Has(code)
	case RuleSetError:
		for _, e := range v {
			if HasErrorCode(e, code) {
				return true
			}
		}
	}
	return false
}

// unpackRuleErrors unpacks error messages recursively scanning [RuleSetError] if it is detected.
func unpackRuleErrors(errs []error, ruleErrors []*RuleError) []*RuleError {
	for _, err := range errs {
		switch v := err.(type) {
		case RuleSetError:
			ruleErrors = unpackRuleErrors(v, ruleErrors)
		case *RuleError:
			ruleErrors = append(ruleErrors, v)
		default:
			ruleErrors = append(ruleErrors, &RuleError{Message: v.Error()})
		}
	}
	return ruleErrors
}

func logWrongErrorType(expected, actual error) {
	logging.Logger().Error("unexpected error type",
		slog.String("actual_type", fmt.Sprintf("%T", actual)),
		slog.String("expected_type", fmt.Sprintf("%T", expected)))
}
