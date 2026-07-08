package rules

import (
	"errors"
	"net/url"
	"strings"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// URLOption configures additional checks for [URL].
type URLOption interface {
	validate(*url.URL) (govy.TemplateVars, bool)
}

type urlOptionFunc func(*url.URL) (govy.TemplateVars, bool)

func (f urlOptionFunc) validate(u *url.URL) (govy.TemplateVars, bool) {
	return f(u)
}

type urlTemplateVars struct {
	Scheme            bool
	HostRequired      bool
	UserInfoForbidden bool
	HostDenyList      bool
	HostAllowList     bool
}

// URLSchemes restricts [URL] to the provided schemes.
func URLSchemes(schemes ...string) URLOption {
	return urlOptionFunc(func(u *url.URL) (govy.TemplateVars, bool) {
		if len(schemes) == 0 || containsString(schemes, u.Scheme) {
			return govy.TemplateVars{}, true
		}
		return govy.TemplateVars{
			PropertyValue:   u,
			ComparisonValue: schemes,
			Custom:          urlTemplateVars{Scheme: true},
		}, false
	})
}

// URLHostRequired requires [URL] values to have a host.
func URLHostRequired() URLOption {
	return urlOptionFunc(func(u *url.URL) (govy.TemplateVars, bool) {
		if u.Hostname() != "" {
			return govy.TemplateVars{}, true
		}
		return govy.TemplateVars{
			PropertyValue: u,
			Custom:        urlTemplateVars{HostRequired: true},
		}, false
	})
}

// URLUserInfoForbidden rejects [URL] values with user information.
func URLUserInfoForbidden() URLOption {
	return urlOptionFunc(func(u *url.URL) (govy.TemplateVars, bool) {
		if u.User == nil {
			return govy.TemplateVars{}, true
		}
		return govy.TemplateVars{
			PropertyValue: u,
			Custom:        urlTemplateVars{UserInfoForbidden: true},
		}, false
	})
}

// URLHostAllowList restricts [URL] values to the provided hostnames.
func URLHostAllowList(hosts ...string) URLOption {
	return urlOptionFunc(func(u *url.URL) (govy.TemplateVars, bool) {
		hostname := u.Hostname()
		if len(hosts) == 0 || containsStringFold(hosts, hostname) {
			return govy.TemplateVars{}, true
		}
		return govy.TemplateVars{
			PropertyValue:   u,
			ComparisonValue: hosts,
			Custom:          urlTemplateVars{HostAllowList: true},
		}, false
	})
}

// URLHostDenyList rejects [URL] values with the provided hostnames.
func URLHostDenyList(hosts ...string) URLOption {
	return urlOptionFunc(func(u *url.URL) (govy.TemplateVars, bool) {
		hostname := u.Hostname()
		if len(hosts) == 0 || !containsStringFold(hosts, hostname) {
			return govy.TemplateVars{}, true
		}
		return govy.TemplateVars{
			PropertyValue:   u,
			ComparisonValue: hosts,
			Custom:          urlTemplateVars{HostDenyList: true},
		}, false
	})
}

// URL ensures the URL is valid.
// The URL must have a scheme (e.g. https://) and contain either host, fragment or opaque data.
func URL(options ...URLOption) govy.Rule[*url.URL] {
	tpl := messagetemplates.Get(messagetemplates.URLTemplate)

	return govy.NewRule(func(v *url.URL) error {
		if err := validateURL(v); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: v,
				Error:         err.Error(),
			})
		}
		for _, option := range options {
			if option == nil {
				continue
			}
			if vars, ok := option.validate(v); !ok {
				return govy.NewRuleErrorTemplate(vars)
			}
		}
		return nil
	}).
		WithErrorCode(ErrorCodeURL).
		WithMessageTemplate(tpl).
		WithDescription(urlDescription)
}

const urlDescription = "valid URL must have a scheme (e.g. https://) and contain either host, fragment or opaque data"

func validateURL(u *url.URL) error {
	if u.Scheme == "" {
		return errors.New("valid URL must have a scheme (e.g. https://)")
	}
	if u.Host == "" && u.Fragment == "" && u.Opaque == "" {
		return errors.New("valid URL must contain either host, fragment or opaque data")
	}
	return nil
}

func containsString(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func containsStringFold(values []string, value string) bool {
	for _, v := range values {
		if strings.EqualFold(v, value) {
			return true
		}
	}
	return false
}
