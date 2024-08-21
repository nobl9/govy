package rules

import (
	"errors"
	"net/url"

	"github.com/nobl9/govy/pkg/govy"
)

// URL ensures the URL is valid.
// The URL must have a scheme (e.g. https://) and contain either host, fragment or opaque data.
func URL() govy.Rule[*url.URL] {
	return govy.NewRule(validateURL).
		WithErrorCode(ErrorCodeURL).
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
