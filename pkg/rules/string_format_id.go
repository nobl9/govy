package rules

import (
	"text/template"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringUUID ensures property's value is a valid UUID string as defined by [RFC 4122].
// It does not enforce a specific UUID version.
//
// [RFC 4122]: https://www.ietf.org/rfc/rfc4122.txt
func StringUUID() govy.Rule[string] {
	return StringMatchRegexp(uuidRegexp()).
		WithDetails("expected RFC-4122 compliant UUID string").
		WithExamples(
			"00000000-0000-0000-0000-000000000000",
			"e190c630-8873-11ee-b9d1-0242ac120002",
			"79258D24-01A7-47E5-ACBB-7E762DE52298",
		).
		WithErrorCode(ErrorCodeStringUUID)
}

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
