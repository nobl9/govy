package rules

import (
	"regexp"
	"sync"
)

// Define all regular expressions here:
var (
	// Ref: https://www.ietf.org/rfc/rfc4122.txt
	uuidRegexp  = lazyRegexCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)
	asciiRegexp = lazyRegexCompile(`^[\x00-\x7F]*$`)
	// Ref: https://www.rfc-editor.org/rfc/rfc4648.html
	standardBase64Regexp = lazyRegexCompile(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`)
	base64URLRegexp      = lazyRegexCompile(`^(?:[A-Za-z0-9_-]{4})*(?:[A-Za-z0-9_-]{2}==|[A-Za-z0-9_-]{3}=)?$`)
	base64RawURLRegexp   = lazyRegexCompile(`^[A-Za-z0-9_-]*$`)
	hexadecimalRegexp    = lazyRegexCompile(`^(?:0[xX])?[0-9a-fA-F]+$`)
	// Ref: https://www.ietf.org/rfc/rfc1123.txt
	rfc1123DnsLabelRegexp      = lazyRegexCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`)
	rfc1123DnsSubdomainRegexp  = lazyRegexCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`)
	k8sQualifiedNamePartRegexp = lazyRegexCompile(`^([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`)
	alphaRegexp                = lazyRegexCompile(`^[a-zA-Z]*$`)
	alphanumericRegexp         = lazyRegexCompile(`^[a-zA-Z0-9]*$`)
	alphaUnicodeRegexp         = lazyRegexCompile(`^[\p{L}]*$`)
	alphanumericUnicodeRegexp  = lazyRegexCompile(`^[\p{L}\p{N}]+$`)
	fqdnRegexp                 = lazyRegexCompile(
		`^([a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})(\.[a-zA-Z0-9]{1}[a-zA-Z0-9-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$`,
	)
)

// lazyRegexCompile returns a function that compiles the regular expression
// once, when the function is called for the first time.
// If the function is never called, the regular expression is never compiled,
// thus saving on performance.
//
// All regular expression literals should be compiled using this function.
//
// Credits: https://github.com/go-playground/validator/commit/2e1df48b5ab876bdd461bdccc51d109389e7572f
func lazyRegexCompile(str string) func() *regexp.Regexp {
	var (
		regex *regexp.Regexp
		once  sync.Once
	)
	return func() *regexp.Regexp {
		once.Do(func() {
			regex = regexp.MustCompile(str)
		})
		return regex
	}
}
