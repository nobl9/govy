package validation

import (
	"net/url"

	"github.com/pkg/errors"
)

func URL() SingleRule[*url.URL] {
	return NewSingleRule(validateURL).
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
