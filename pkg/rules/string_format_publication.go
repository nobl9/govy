package rules

import (
	"strings"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringISBN ensures the property's value is a valid ISBN-10 or ISBN-13.
// It accepts digits separated by single spaces or hyphens and validates the check digit.
// It does not validate ISBN registration-group or publisher ranges.
func StringISBN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISBNTemplate)

	return govy.NewRule(func(s string) error {
		if !isISBN10(s) && !isISBN13(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISBN).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid ISBN-10 or ISBN-13")
}

// StringISBN10 ensures the property's value is a valid ISBN-10.
// It accepts digits separated by single spaces or hyphens and validates the check digit.
// It does not validate ISBN registration-group or publisher ranges.
func StringISBN10() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISBN10Template)

	return govy.NewRule(func(s string) error {
		if !isISBN10(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISBN10).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid ISBN-10")
}

// StringISBN13 ensures the property's value is a valid ISBN-13.
// It accepts digits separated by single spaces or hyphens and validates the check digit.
// It does not validate ISBN registration-group or publisher ranges.
func StringISBN13() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISBN13Template)

	return govy.NewRule(func(s string) error {
		if !isISBN13(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISBN13).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid ISBN-13")
}

// StringISSN ensures the property's value is a valid ISSN.
// It accepts 4 digits, a hyphen, 3 digits, and a final check character.
func StringISSN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISSNTemplate)

	return govy.NewRule(func(s string) error {
		if !isISSN(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISSN).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid hyphenated ISSN")
}

func isISBN10(s string) bool {
	isbn, ok := normalizeISBN(s)
	if !ok || len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		if !isASCIIDigit(isbn[i]) {
			return false
		}
		sum += int(isbn[i]-'0') * (10 - i)
	}

	switch checkDigit := isbn[9]; {
	case isASCIIDigit(checkDigit):
		sum += int(checkDigit - '0')
	case checkDigit == 'X' || checkDigit == 'x':
		sum += 10
	default:
		return false
	}
	return sum%11 == 0
}

func isISBN13(s string) bool {
	isbn, ok := normalizeISBN(s)
	if !ok || len(isbn) != 13 || (!strings.HasPrefix(isbn, "978") && !strings.HasPrefix(isbn, "979")) {
		return false
	}

	sum := 0
	for i := 0; i < 12; i++ {
		if !isASCIIDigit(isbn[i]) {
			return false
		}
		weight := 1
		if i%2 != 0 {
			weight = 3
		}
		sum += int(isbn[i]-'0') * weight
	}
	if !isASCIIDigit(isbn[12]) {
		return false
	}
	return (10-sum%10)%10 == int(isbn[12]-'0')
}

func normalizeISBN(s string) (string, bool) {
	if s == "" {
		return "", false
	}

	var builder strings.Builder
	previousWasSeparator := false
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case isASCIIDigit(c) || c == 'X' || c == 'x':
			builder.WriteByte(c)
			previousWasSeparator = false
		case c == '-' || c == ' ':
			if i == 0 || i == len(s)-1 || previousWasSeparator {
				return "", false
			}
			previousWasSeparator = true
		default:
			return "", false
		}
	}
	return builder.String(), true
}

func isISSN(s string) bool {
	if !issnRegexp().MatchString(s) {
		return false
	}

	issn := strings.ReplaceAll(s, "-", "")
	sum := 0
	for i := 0; i < 7; i++ {
		if !isASCIIDigit(issn[i]) {
			return false
		}
		sum += int(issn[i]-'0') * (8 - i)
	}

	switch checkDigit := issn[7]; {
	case isASCIIDigit(checkDigit):
		sum += int(checkDigit - '0')
	case checkDigit == 'X' || checkDigit == 'x':
		sum += 10
	default:
		return false
	}
	return sum%11 == 0
}

func isASCIIDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
