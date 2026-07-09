package rules

import (
	"text/template"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringCreditCard ensures the property's value is a plausible digit-only
// payment card number. It requires a 13- to 19-digit value that passes the
// Luhn checksum and rejects all-same-digit values.
func StringCreditCard() govy.Rule[string] {
	return newStringPaymentBankingRule(
		messagetemplates.Get(messagetemplates.StringCreditCardTemplate),
		ErrorCodeStringCreditCard,
		isValidCreditCard,
	)
}

// StringLuhnChecksum ensures the property's value is a digit-only string that
// passes the Luhn checksum.
func StringLuhnChecksum() govy.Rule[string] {
	return newStringPaymentBankingRule(
		messagetemplates.Get(messagetemplates.StringLuhnChecksumTemplate),
		ErrorCodeStringLuhnChecksum,
		isLuhnChecksumValid,
	)
}

// StringBIC ensures the property's value matches the current BIC/SWIFT syntax.
func StringBIC() govy.Rule[string] {
	return newStringPaymentBankingRule(
		messagetemplates.Get(messagetemplates.StringBICTemplate),
		ErrorCodeStringBIC,
		bicRegexp().MatchString,
	)
}

// StringBICISO93622014 ensures the property's value matches the ISO 9362:2014
// BIC syntax.
func StringBICISO93622014() govy.Rule[string] {
	return newStringPaymentBankingRule(
		messagetemplates.Get(messagetemplates.StringBICISO93622014Template),
		ErrorCodeStringBICISO93622014,
		bicISO93622014Regexp().MatchString,
	)
}

func newStringPaymentBankingRule(
	tpl *template.Template,
	errorCode govy.ErrorCode,
	validator func(string) bool,
) govy.Rule[string] {
	return govy.NewRule(func(s string) error {
		if validator(s) {
			return nil
		}
		return govy.NewRuleErrorTemplate(govy.TemplateVars{PropertyValue: s})
	}).
		WithErrorCode(errorCode).
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
