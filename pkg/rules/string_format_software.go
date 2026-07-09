package rules

import (
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringSemver ensures the property's value is a valid Semantic Versioning 2.0.0 version.
// It does not accept a leading "v" prefix.
func StringSemver() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSemverTemplate)

	return govy.NewRule(func(s string) error {
		if !semverRegexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringSemver).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid Semantic Versioning 2.0.0 version")
}

// StringCVE ensures the property's value is a valid CVE ID.
// It validates the CVE-YEAR-SEQUENCE syntax only and does not check whether
// the CVE record is assigned, reserved, published, or rejected.
func StringCVE() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringCVETemplate)

	return govy.NewRule(func(s string) error {
		if !cveRegexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCVE).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid CVE ID in CVE-YEAR-SEQUENCE format")
}
