package rules

import (
	"errors"
	"net/url"
	"slices"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// URLOption configures additional checks for [URL].
type URLOption interface {
	apply(*urlRuleOptions)
}

type urlOptionFunc func(*urlRuleOptions)

func (f urlOptionFunc) apply(options *urlRuleOptions) {
	f(options)
}

type urlRuleOptions struct {
	schemes       []string
	hostRequired  bool
	forbidUser    bool
	hostAllowList []string
	hostDenyList  []string
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
	return urlOptionFunc(func(options *urlRuleOptions) {
		options.schemes = slices.Clone(schemes)
	})
}

// URLHostRequired requires [URL] values to have a host.
func URLHostRequired() URLOption {
	return urlOptionFunc(func(options *urlRuleOptions) {
		options.hostRequired = true
	})
}

// URLUserInfoForbidden rejects [URL] values with user information.
func URLUserInfoForbidden() URLOption {
	return urlOptionFunc(func(options *urlRuleOptions) {
		options.forbidUser = true
	})
}

// URLHostAllowList restricts [URL] values to the provided hostnames.
func URLHostAllowList(hosts ...string) URLOption {
	return urlOptionFunc(func(options *urlRuleOptions) {
		options.hostAllowList = slices.Clone(hosts)
	})
}

// URLHostDenyList rejects [URL] values with the provided hostnames.
// Denied entries take precedence over entries defined with [URLHostAllowList].
func URLHostDenyList(hosts ...string) URLOption {
	return urlOptionFunc(func(options *urlRuleOptions) {
		options.hostDenyList = slices.Clone(hosts)
	})
}

// URL ensures the URL is valid.
// The URL must have a scheme (e.g. https://) and contain either host, fragment or opaque data.
func URL(options ...URLOption) govy.Rule[*url.URL] {
	tpl := messagetemplates.Get(messagetemplates.URLTemplate)
	ruleOptions := newURLRuleOptions(options...)

	return govy.NewRule(func(v *url.URL) error {
		if err := validateURL(v); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: v,
				Error:         err.Error(),
			})
		}
		if vars, ok := validateURLOptions(v, ruleOptions); !ok {
			return govy.NewRuleErrorTemplate(vars)
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

func newURLRuleOptions(options ...URLOption) urlRuleOptions {
	var ruleOptions urlRuleOptions
	for _, option := range options {
		if option != nil {
			option.apply(&ruleOptions)
		}
	}
	return ruleOptions
}

func validateURLOptions(u *url.URL, options urlRuleOptions) (govy.TemplateVars, bool) {
	if len(options.schemes) > 0 && !containsString(options.schemes, u.Scheme) {
		return govy.TemplateVars{
			PropertyValue:   u,
			ComparisonValue: options.schemes,
			Custom:          urlTemplateVars{Scheme: true},
		}, false
	}

	hostname := u.Hostname()
	if options.hostRequired && hostname == "" {
		return govy.TemplateVars{
			PropertyValue: u,
			Custom:        urlTemplateVars{HostRequired: true},
		}, false
	}
	if options.forbidUser && u.User != nil {
		return govy.TemplateVars{
			PropertyValue: u,
			Custom:        urlTemplateVars{UserInfoForbidden: true},
		}, false
	}
	if len(options.hostDenyList) > 0 && containsString(options.hostDenyList, hostname) {
		return govy.TemplateVars{
			PropertyValue:   u,
			ComparisonValue: options.hostDenyList,
			Custom:          urlTemplateVars{HostDenyList: true},
		}, false
	}
	if len(options.hostAllowList) > 0 && !containsString(options.hostAllowList, hostname) {
		return govy.TemplateVars{
			PropertyValue:   u,
			ComparisonValue: options.hostAllowList,
			Custom:          urlTemplateVars{HostAllowList: true},
		}, false
	}
	return govy.TemplateVars{}, true
}

func containsString(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}
