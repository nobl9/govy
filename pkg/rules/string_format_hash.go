package rules

import (
	"text/template"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringMD5 ensures the property's value is a lowercase hexadecimal MD5 digest.
func StringMD5() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMD5Template)

	return newStringHashDigestRule(func(s string) bool {
		return md5Regexp().MatchString(s)
	}, ErrorCodeStringMD5, tpl)
}

// StringSHA256 ensures the property's value is a lowercase hexadecimal SHA-256 digest.
func StringSHA256() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSHA256Template)

	return newStringHashDigestRule(func(s string) bool {
		return sha256Regexp().MatchString(s)
	}, ErrorCodeStringSHA256, tpl)
}

// StringSHA384 ensures the property's value is a lowercase hexadecimal SHA-384 digest.
func StringSHA384() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSHA384Template)

	return newStringHashDigestRule(func(s string) bool {
		return sha384Regexp().MatchString(s)
	}, ErrorCodeStringSHA384, tpl)
}

// StringSHA512 ensures the property's value is a lowercase hexadecimal SHA-512 digest.
func StringSHA512() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSHA512Template)

	return newStringHashDigestRule(func(s string) bool {
		return sha512Regexp().MatchString(s)
	}, ErrorCodeStringSHA512, tpl)
}

func newStringHashDigestRule(
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
