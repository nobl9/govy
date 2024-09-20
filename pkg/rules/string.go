package rules

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"unicode"

	"github.com/nobl9/govy/pkg/govy"
)

// StringNotEmpty ensures the property's value is not empty.
// The string is considered empty if it contains only whitespace characters.
func StringNotEmpty() govy.Rule[string] {
	msg := "string cannot be empty"
	return govy.NewRule(func(s string) error {
		if len(strings.TrimSpace(s)) == 0 {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringNotEmpty).
		WithDescription(msg)
}

// StringMatchRegexp ensures the property's value matches the regular expression.
// The error message can be enhanced with examples of valid values.
func StringMatchRegexp(re *regexp.Regexp, examples ...string) govy.Rule[string] {
	msg := fmt.Sprintf("string must match regular expression: '%s'", re.String())
	if len(examples) > 0 {
		msg += " " + prettyExamples(examples)
	}
	return govy.NewRule(func(s string) error {
		if !re.MatchString(s) {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMatchRegexp).
		WithDescription(msg)
}

// StringDenyRegexp ensures the property's value does not match the regular expression.
// The error message can be enhanced with examples of invalid values.
func StringDenyRegexp(re *regexp.Regexp, examples ...string) govy.Rule[string] {
	msg := fmt.Sprintf("string must not match regular expression: '%s'", re.String())
	if len(examples) > 0 {
		msg += " " + prettyExamples(examples)
	}
	return govy.NewRule(func(s string) error {
		if re.MatchString(s) {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringDenyRegexp).
		WithDescription(msg)
}

// StringDNSLabel ensures the property's value is a valid DNS label as defined by RFC 1123.
func StringDNSLabel() govy.RuleSet[string] {
	return govy.NewRuleSet(
		StringLength(1, 63),
		StringMatchRegexp(rfc1123DnsLabelRegexp(), "my-name", "123-abc").
			WithDetails("an RFC-1123 compliant label name must consist of lower case alphanumeric characters or '-',"+
				" and must start and end with an alphanumeric character"),
	).WithErrorCode(ErrorCodeStringDNSLabel)
}

// StringEmail ensures the property's value is a valid email address.
// It follows RFC 5322 specification which is more permissive in regards to domain names.
// Ref: https://www.ietf.org/rfc/rfc5322.txt
func StringEmail() govy.Rule[string] {
	msg := "string must be a valid email address"
	return govy.NewRule(func(s string) error {
		if _, err := mail.ParseAddress(s); err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringEmail).
		WithDescription(msg)
}

// StringURL ensures property's value is a valid URL as defined by [url.Parse] function.
// Unlike [URL] it does not impose any additional rules upon parsed [url.URL].
func StringURL() govy.Rule[string] {
	return govy.NewRule(func(s string) error {
		u, err := url.Parse(s)
		if err != nil {
			return fmt.Errorf("failed to parse URL: %w", err)
		}
		return validateURL(u)
	}).
		WithErrorCode(ErrorCodeStringURL).
		WithDescription(urlDescription)
}

// StringMAC ensures property's value is a valid MAC address.
func StringMAC() govy.Rule[string] {
	msg := "string must be a valid MAC address"
	return govy.NewRule(func(s string) error {
		if _, err := net.ParseMAC(s); err != nil {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMAC).
		WithDescription(msg)
}

// StringIP ensures property's value is a valid IP address.
func StringIP() govy.Rule[string] {
	msg := "string must be a valid IP address"
	return govy.NewRule(func(s string) error {
		if ip := net.ParseIP(s); ip == nil {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringIP).
		WithDescription(msg)
}

// StringIPv4 ensures property's value is a valid IPv4 address.
func StringIPv4() govy.Rule[string] {
	msg := "string must be a valid IPv4 address"
	return govy.NewRule(func(s string) error {
		if ip := net.ParseIP(s); ip == nil || ip.To4() == nil {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringIPv4).
		WithDescription(msg)
}

// StringIPv6 ensures property's value is a valid IPv6 address.
func StringIPv6() govy.Rule[string] {
	msg := "string must be a valid IPv6 address"
	return govy.NewRule(func(s string) error {
		if ip := net.ParseIP(s); ip == nil || ip.To4() != nil || len(ip) != net.IPv6len {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringIPv6).
		WithDescription(msg)
}

// StringCIDR ensures property's value is a valid CIDR notation IP address.
func StringCIDR() govy.Rule[string] {
	msg := "string must be a valid CIDR notation IP address"
	return govy.NewRule(func(s string) error {
		if _, _, err := net.ParseCIDR(s); err != nil {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCIDR).
		WithDescription(msg)
}

// StringCIDRv4 ensures property's value is a valid CIDR notation IPv4 address.
func StringCIDRv4() govy.Rule[string] {
	msg := "string must be a valid CIDR notation IPv4 address"
	return govy.NewRule(func(s string) error {
		if ip, ipNet, err := net.ParseCIDR(s); err != nil || ip.To4() == nil || !ipNet.IP.Equal(ip) {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCIDRv4).
		WithDescription(msg)
}

// StringCIDRv4 ensures property's value is a valid CIDR notation IPv6 address.
func StringCIDRv6() govy.Rule[string] {
	msg := "string must be a valid CIDR notation IPv6 address"
	return govy.NewRule(func(s string) error {
		if ip, _, err := net.ParseCIDR(s); err != nil || ip.To4() != nil || len(ip) != net.IPv6len {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCIDRv6).
		WithDescription(msg)
}

// StringUUID ensures property's value is a valid UUID string as defined by RFC 4122.
// It does not enforce a specific UUID version.
// Ref: https://www.ietf.org/rfc/rfc4122.txt
func StringUUID() govy.Rule[string] {
	return StringMatchRegexp(uuidRegexp(),
		"00000000-0000-0000-0000-000000000000",
		"e190c630-8873-11ee-b9d1-0242ac120002",
		"79258D24-01A7-47E5-ACBB-7E762DE52298").
		WithDetails("expected RFC-4122 compliant UUID string").
		WithErrorCode(ErrorCodeStringUUID)
}

// StringASCII ensures property's value contains only ASCII characters.
func StringASCII() govy.Rule[string] {
	return StringMatchRegexp(asciiRegexp()).WithErrorCode(ErrorCodeStringASCII)
}

// StringJSON ensures property's value is a valid JSON literal.
func StringJSON() govy.Rule[string] {
	msg := "string must be a valid JSON"
	return govy.NewRule(func(s string) error {
		if !json.Valid([]byte(s)) {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringJSON).
		WithDescription(msg)
}

// StringContains ensures the property's value contains all the provided substrings.
func StringContains(substrings ...string) govy.Rule[string] {
	msg := "string must contain the following substrings: " + prettyStringList(substrings)
	return govy.NewRule(func(s string) error {
		matched := true
		for _, substr := range substrings {
			if !strings.Contains(s, substr) {
				matched = false
				break
			}
		}
		if !matched {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringContains).
		WithDescription(msg)
}

// StringExcludes ensures the property's value does not contain any of the provided substrings.
func StringExcludes(substrings ...string) govy.Rule[string] {
	msg := "string must not contain any of the following substrings: " + prettyStringList(substrings)
	return govy.NewRule(func(s string) error {
		for _, substr := range substrings {
			if strings.Contains(s, substr) {
				return errors.New(msg)
			}
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringExcludes).
		WithDescription(msg)
}

// StringStartsWith ensures the property's value starts with one of the provided prefixes.
func StringStartsWith(prefixes ...string) govy.Rule[string] {
	var msg string
	if len(prefixes) == 1 {
		msg = fmt.Sprintf("string must start with '%s' prefix", prefixes[0])
	} else {
		msg = "string must start with one of the following prefixes: " + prettyStringList(prefixes)
	}
	return govy.NewRule(func(s string) error {
		matched := false
		for _, prefix := range prefixes {
			if strings.HasPrefix(s, prefix) {
				matched = true
				break
			}
		}
		if !matched {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringStartsWith).
		WithDescription(msg)
}

// StringEndsWith ensures the property's value ends with one of the provided suffixes.
func StringEndsWith(suffixes ...string) govy.Rule[string] {
	var msg string
	if len(suffixes) == 1 {
		msg = fmt.Sprintf("string must end with '%s' suffix", suffixes[0])
	} else {
		msg = "string must end with one of the following suffixes: " + prettyStringList(suffixes)
	}
	return govy.NewRule(func(s string) error {
		matched := false
		for _, suffix := range suffixes {
			if strings.HasSuffix(s, suffix) {
				matched = true
				break
			}
		}
		if !matched {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringEndsWith).
		WithDescription(msg)
}

// StringTitle ensures each word in a string starts with a capital letter.
func StringTitle() govy.Rule[string] {
	msg := "each word in a string must start with a capital letter"
	return govy.NewRule(func(s string) error {
		if len(s) == 0 {
			return errors.New(msg)
		}
		prev := ' '
		for _, r := range s {
			if isStringSeparator(prev) {
				if !unicode.IsUpper(r) && !isStringSeparator(r) {
					return errors.New(msg)
				}
			}
			prev = r
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringTitle).
		WithDescription(msg)
}

func prettyExamples(examples []string) string {
	if len(examples) == 0 {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("(e.g. ")
	prettyStringListBuilder(&b, examples, true)
	b.WriteString(")")
	return b.String()
}

func prettyStringList[T any](values []T) string {
	b := new(strings.Builder)
	prettyStringListBuilder(b, values, true)
	return b.String()
}

func prettyStringListBuilder[T any](b *strings.Builder, values []T, surroundInSingleQuotes bool) {
	b.Grow(len(values))
	for i := range values {
		if i > 0 {
			b.WriteString(", ")
		}
		if surroundInSingleQuotes {
			b.WriteString("'")
		}
		fmt.Fprint(b, values[i])
		if surroundInSingleQuotes {
			b.WriteString("'")
		}
	}
}

// isStringSeparator is directly copied from [strings] package.
func isStringSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}
