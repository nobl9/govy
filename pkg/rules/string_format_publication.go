package rules

import (
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

const (
	isbn10Length = 10
	isbn13Length = 13
	issnLength   = 9
)

// StringISBN ensures the property's value is a valid International Standard
// Book Number (ISBN) in ISBN-10 or ISBN-13 format.
// It accepts digits separated by single spaces or hyphens and validates the check digit.
// It does not validate ISBN registration-group or publisher ranges.
func StringISBN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISBNTemplate)

	return govy.NewRule(func(s string) error {
		if !isISBN(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISBN).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid International Standard Book Number (ISBN) in ISBN-10 or ISBN-13 format")
}

// StringISBN10 ensures the property's value is a valid International Standard
// Book Number (ISBN) in ISBN-10 format.
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
		WithDescription("string must be a valid International Standard Book Number (ISBN) in ISBN-10 format")
}

// StringISBN13 ensures the property's value is a valid International Standard
// Book Number (ISBN) in ISBN-13 format.
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
		WithDescription("string must be a valid International Standard Book Number (ISBN) in ISBN-13 format")
}

// StringISSN ensures the property's value is a valid International Standard
// Serial Number (ISSN).
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
		WithDescription("string must be a valid hyphenated International Standard Serial Number (ISSN)")
}

func isISBN(s string) bool {
	isbn, length, ok := normalizeISBN(s, isbn13Length)
	if !ok {
		return false
	}

	switch length {
	case isbn10Length:
		return validISBN10(&isbn)
	case isbn13Length:
		return validISBN13(&isbn)
	default:
		return false
	}
}

func isISBN10(s string) bool {
	isbn, length, ok := normalizeISBN(s, isbn10Length)
	if !ok || length != isbn10Length {
		return false
	}
	return validISBN10(&isbn)
}

func validISBN10(isbn *[isbn13Length]byte) bool {
	sum := 0
	for i, digit := range isbn[:isbn10Length-1] {
		if !isASCIIDigit(digit) {
			return false
		}
		sum += int(digit-'0') * (isbn10Length - i)
	}

	switch checkDigit := isbn[isbn10Length-1]; {
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
	isbn, length, ok := normalizeISBN(s, isbn13Length)
	if !ok || length != isbn13Length {
		return false
	}
	return validISBN13(&isbn)
}

func validISBN13(isbn *[isbn13Length]byte) bool {
	if isbn[0] != '9' || isbn[1] != '7' || (isbn[2] != '8' && isbn[2] != '9') {
		return false
	}

	sum := 0
	for i, digit := range isbn[:isbn13Length-1] {
		if !isASCIIDigit(digit) {
			return false
		}
		weight := 1
		if i%2 != 0 {
			weight = 3
		}
		sum += int(digit-'0') * weight
	}
	if !isASCIIDigit(isbn[isbn13Length-1]) {
		return false
	}
	return (10-sum%10)%10 == int(isbn[isbn13Length-1]-'0')
}

func normalizeISBN(s string, maximumDigits int) (isbn [isbn13Length]byte, length int, ok bool) {
	if len(s) == 0 || len(s) > maximumDigits*2-1 {
		return isbn, 0, false
	}

	previousWasSeparator := false
	for i := range len(s) {
		switch c := s[i]; {
		case isASCIIDigit(c) || c == 'X' || c == 'x':
			if length == maximumDigits {
				return isbn, 0, false
			}
			isbn[length] = c
			length++
			previousWasSeparator = false
		case c == '-' || c == ' ':
			if i == 0 || i == len(s)-1 || previousWasSeparator {
				return isbn, 0, false
			}
			previousWasSeparator = true
		default:
			return isbn, 0, false
		}
	}
	return isbn, length, true
}

func isISSN(s string) bool {
	if len(s) != issnLength || s[4] != '-' {
		return false
	}

	sum := 0
	for i := range 4 {
		if !isASCIIDigit(s[i]) {
			return false
		}
		sum += int(s[i]-'0') * (8 - i)
	}
	for i := 5; i < 8; i++ {
		if !isASCIIDigit(s[i]) {
			return false
		}
		sum += int(s[i]-'0') * (9 - i)
	}
	switch checkDigit := s[issnLength-1]; {
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
