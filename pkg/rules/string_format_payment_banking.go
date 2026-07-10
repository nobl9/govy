package rules

import (
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringCreditCard ensures the property's value is a plausible digit-only
// payment card number. It requires a 13- to 19-digit value that passes the
// Luhn checksum and rejects all-same-digit values.
func StringCreditCard() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringCreditCardTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidCreditCard(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCreditCard).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringLuhnChecksum ensures the property's value is a digit-only string that
// passes the Luhn checksum.
func StringLuhnChecksum() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringLuhnChecksumTemplate)

	return govy.NewRule(func(s string) error {
		if !isLuhnChecksumValid(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringLuhnChecksum).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringBIC ensures the property's value matches the current BIC/SWIFT syntax.
func StringBIC() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBICTemplate)

	return govy.NewRule(func(s string) error {
		if !bicRegexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBIC).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringBICISO93622014 ensures the property's value matches the ISO 9362:2014
// BIC syntax.
func StringBICISO93622014() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBICISO93622014Template)

	return govy.NewRule(func(s string) error {
		if !bicISO93622014Regexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBICISO93622014).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func isValidCreditCard(s string) bool {
	if len(s) < 13 || len(s) > 19 {
		return false
	}
	if allSameDigit(s) {
		return false
	}
	return isLuhnChecksumValid(s)
}

func isLuhnChecksumValid(s string) bool {
	if s == "" {
		return false
	}

	sum := 0
	shouldDouble := false
	for i := len(s) - 1; i >= 0; i-- {
		digit := s[i] - '0'
		if digit > 9 {
			return false
		}
		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += int(digit)
		shouldDouble = !shouldDouble
	}
	return sum%10 == 0
}

func allSameDigit(s string) bool {
	if s == "" {
		return false
	}
	first := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != first {
			return false
		}
	}
	return true
}
