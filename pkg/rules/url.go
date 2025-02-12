package rules

import (
	"net/url"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// URL ensures the URL is valid.
// The URL must have a scheme (e.g. https://) and contain either host, fragment or opaque data.
func URL() govy.Rule[*url.URL] {
	tpl := messagetemplates.Get(messagetemplates.URLTemplate)

	return govy.NewRule(validateURL).
		WithErrorCode(ErrorCodeURL).
		WithMessageTemplate(tpl).
		WithDescription(urlDescription)
}

const urlDescription = "valid URL must have a scheme (e.g. https://) and contain either host, fragment or opaque data"

func validateURL(u *url.URL) error {
	if u.Scheme == "" {
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			PropertyValue: u.String(),
			Error:         "valid URL must have a scheme (e.g. https://)",
		})
	}
	if u.Host == "" && u.Fragment == "" && u.Opaque == "" {
		return govy.NewRuleErrorTemplate(govy.TemplateVars{
			PropertyValue: u.String(),
			Error:         "valid URL must contain either host, fragment or opaque data",
		})
	}
	return nil
}
