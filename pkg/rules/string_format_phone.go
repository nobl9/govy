package rules

import (
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringE164 ensures the property's value is a valid E.164 phone number.
func StringE164() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringE164Template)

	return govy.NewRule(func(s string) error {
		if !e164Regexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringE164).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}
