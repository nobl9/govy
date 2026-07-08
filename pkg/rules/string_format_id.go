package rules

import (
	"text/template"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringUUIDRFC4122 ensures the property's value is an RFC 4122 UUID string.
// It requires the canonical 36-character form, a version from 1 through 5, and RFC 4122 variant bits.
func StringUUIDRFC4122() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringUUIDRFC4122Template),
		ErrorCodeStringUUIDRFC4122,
		func(s string) bool { return uuidRFC4122Regexp().MatchString(s) },
	)
}

// StringUUIDv3 ensures the property's value is a version 3 RFC 4122 UUID string.
func StringUUIDv3() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringUUIDv3Template),
		ErrorCodeStringUUIDv3,
		func(s string) bool { return uuidv3Regexp().MatchString(s) },
	)
}

// StringUUIDv4 ensures the property's value is a version 4 RFC 4122 UUID string.
func StringUUIDv4() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringUUIDv4Template),
		ErrorCodeStringUUIDv4,
		func(s string) bool { return uuidv4Regexp().MatchString(s) },
	)
}

// StringUUIDv5 ensures the property's value is a version 5 RFC 4122 UUID string.
func StringUUIDv5() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringUUIDv5Template),
		ErrorCodeStringUUIDv5,
		func(s string) bool { return uuidv5Regexp().MatchString(s) },
	)
}

// StringULID ensures the property's value is a 26-character Crockford-base32 ULID string.
func StringULID() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringULIDTemplate),
		ErrorCodeStringULID,
		isValidULID,
	)
}

// StringCreditCard ensures the property's value is a plausible digit-only credit card number.
// It requires a 13- to 19-digit value that passes the Luhn checksum.
func StringCreditCard() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringCreditCardTemplate),
		ErrorCodeStringCreditCard,
		isValidCreditCard,
	)
}

// StringLuhnChecksum ensures the property's value is a digit-only string that passes the Luhn checksum.
func StringLuhnChecksum() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringLuhnChecksumTemplate),
		ErrorCodeStringLuhnChecksum,
		isLuhnChecksumValid,
	)
}

// StringBIC ensures the property's value matches the current BIC/SWIFT syntax.
func StringBIC() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringBICTemplate),
		ErrorCodeStringBIC,
		func(s string) bool { return bicRegexp().MatchString(s) },
	)
}

// StringBICISO93622014 ensures the property's value matches the ISO 9362:2014 BIC syntax.
func StringBICISO93622014() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringBICISO93622014Template),
		ErrorCodeStringBICISO93622014,
		func(s string) bool { return bicISO93622014Regexp().MatchString(s) },
	)
}

// StringEIN ensures the property's value is a US EIN in NN-NNNNNNN format.
func StringEIN() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringEINTemplate),
		ErrorCodeStringEIN,
		isValidEIN,
	)
}

// StringSSN ensures the property's value is a US SSN in NNN-NN-NNNN format.
// It rejects area, group, and serial number groups that are commonly invalid.
func StringSSN() govy.Rule[string] {
	return newStringFormatIDRule(
		messagetemplates.Get(messagetemplates.StringSSNTemplate),
		ErrorCodeStringSSN,
		isValidSSN,
	)
}

func newStringFormatIDRule(
	tpl *template.Template,
	errorCode govy.ErrorCode,
	validator func(string) bool,
) govy.Rule[string] {
	return govy.NewRule(func(s string) error {
		if !validator(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(errorCode).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func isValidULID(s string) bool {
	if len(s) != 26 {
		return false
	}
	if s[0] < '0' || s[0] > '7' {
		return false
	}
	for _, r := range s {
		switch {
		case r >= '0' && r <= '9':
		case r >= 'A' && r <= 'H':
		case r >= 'J' && r <= 'K':
		case r >= 'M' && r <= 'N':
		case r >= 'P' && r <= 'T':
		case r >= 'V' && r <= 'Z':
		case r >= 'a' && r <= 'h':
		case r >= 'j' && r <= 'k':
		case r >= 'm' && r <= 'n':
		case r >= 'p' && r <= 't':
		case r >= 'v' && r <= 'z':
		default:
			return false
		}
	}
	return true
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
	if len(s) == 0 {
		return false
	}
	var sum int
	double := false
	for i := len(s) - 1; i >= 0; i-- {
		digit := s[i] - '0'
		if digit > 9 {
			return false
		}
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += int(digit)
		double = !double
	}
	return sum%10 == 0
}

func allSameDigit(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func isValidEIN(s string) bool {
	if !einRegexp().MatchString(s) {
		return false
	}
	return isValidEINPrefix(s[:2])
}

func isValidEINPrefix(prefix string) bool {
	switch prefix {
	case "01", "02", "03", "04", "05", "06", "10", "11", "12", "13", "14", "15", "16",
		"20", "21", "22", "23", "24", "25", "26", "27",
		"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
		"40", "41", "42", "43", "44", "45", "46", "47", "48",
		"50", "51", "52", "53", "54", "55", "56", "57", "58", "59",
		"60", "61", "62", "63", "64", "65", "66", "67", "68",
		"71", "72", "73", "74", "75", "76", "77",
		"80", "81", "82", "83", "84", "85", "86", "87", "88",
		"90", "91", "92", "93", "94", "95", "98", "99":
		return true
	default:
		return false
	}
}

func isValidSSN(s string) bool {
	if !ssnRegexp().MatchString(s) {
		return false
	}
	area := s[:3]
	group := s[4:6]
	serial := s[7:]
	return area != "000" &&
		area != "666" &&
		area[0] != '9' &&
		group != "00" &&
		serial != "0000"
}
