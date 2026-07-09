package rules

import (
	"encoding/base64"
	"text/template"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringBase64 ensures the property's value is a standard padded base64 string.
// It validates input with [base64.StdEncoding].
func StringBase64() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBase64Template)

	return newStringEncodingRule(func(s string) bool {
		return standardBase64Regexp().MatchString(s) &&
			decodesBase64(base64.StdEncoding.Strict(), s)
	}, ErrorCodeStringBase64, tpl)
}

// StringBase64URL ensures the property's value is a URL-safe padded base64 string.
// It validates input with [base64.URLEncoding].
func StringBase64URL() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBase64URLTemplate)

	return newStringEncodingRule(func(s string) bool {
		return base64URLRegexp().MatchString(s) &&
			decodesBase64(base64.URLEncoding.Strict(), s)
	}, ErrorCodeStringBase64URL, tpl)
}

// StringBase64RawURL ensures the property's value is a URL-safe base64 string without padding.
// It validates input with [base64.RawURLEncoding].
func StringBase64RawURL() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBase64RawURLTemplate)

	return newStringEncodingRule(func(s string) bool {
		return base64RawURLRegexp().MatchString(s) &&
			decodesBase64(base64.RawURLEncoding.Strict(), s)
	}, ErrorCodeStringBase64RawURL, tpl)
}

// StringHexadecimal ensures the property's value is a hexadecimal string.
// It allows an optional `0x` or `0X` prefix.
func StringHexadecimal() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringHexadecimalTemplate)

	return newStringEncodingRule(func(s string) bool {
		return hexadecimalRegexp().MatchString(s)
	}, ErrorCodeStringHexadecimal, tpl)
}

func newStringEncodingRule(
	validate func(string) bool,
	errorCode govy.ErrorCode,
	tpl *template.Template,
) govy.Rule[string] {
	return govy.NewRule(func(s string) error {
		if !validate(s) {
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

func decodesBase64(encoding *base64.Encoding, s string) bool {
	_, err := encoding.DecodeString(s)
	return err == nil
}
