package rules

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

const uuidPattern = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`

// StringNotEmpty ensures the property's value is not empty.
// The string is considered empty if it contains only whitespace characters.
func StringNotEmpty() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringNonEmptyTemplate)

	return govy.NewRule(func(s string) error {
		if len(strings.TrimSpace(s)) == 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringNotEmpty).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringMatchRegexp ensures the property's value matches the regular expression.
// The error message can be enhanced with examples of valid values.
func StringMatchRegexp(re *regexp.Regexp) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMatchRegexpTemplate)

	return govy.NewRule(func(s string) error {
		if !re.MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: re.String(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMatchRegexp).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: re.String(),
		}))
}

// StringDenyRegexp ensures the property's value does not match the regular expression.
// The error message can be enhanced with examples of invalid values.
func StringDenyRegexp(re *regexp.Regexp) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringDenyRegexpTemplate)

	return govy.NewRule(func(s string) error {
		if re.MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: re.String(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringDenyRegexp).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: re.String(),
		}))
}

// StringDNSLabel ensures the property's value is a valid DNS label as defined by [RFC 1123].
//
// [RFC 1123]: https://www.ietf.org/rfc/rfc1123.txt
func StringDNSLabel() govy.RuleSet[string] {
	return govy.NewRuleSet(
		StringLength(1, 63),
		StringMatchRegexp(rfc1123DnsLabelRegexp()).
			WithDetails("an RFC-1123 compliant label name must consist of lower case alphanumeric characters or '-',"+
				" and must start and end with an alphanumeric character").
			WithExamples("my-name", "123-abc"),
	).
		WithErrorCode(ErrorCodeStringDNSLabel).
		Cascade(govy.CascadeModeStop)
}

// StringDNSSubdomain ensures the property's value is a valid DNS subdomain as defined by [RFC 1123].
//
// [RFC 1123]: https://www.ietf.org/rfc/rfc1123.txt
func StringDNSSubdomain() govy.RuleSet[string] {
	return govy.NewRuleSet(
		StringLength(1, 253),
		StringMatchRegexp(rfc1123DnsSubdomainRegexp()).
			WithDetails("an RFC-1123 compliant subdomain must consist of lower case alphanumeric characters, '-'"+
				" or '.', and must start and end with an alphanumeric character").
			WithExamples("example.com"),
	).
		WithErrorCode(ErrorCodeStringDNSSubdomain).
		Cascade(govy.CascadeModeStop)
}

// StringEmail ensures the property's value is a valid email address.
// It follows [RFC 5322] specification which is more permissive in regards to domain names.
//
// [RFC 5322]: https://www.ietf.org/rfc/rfc5322.txt
func StringEmail() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringEmailTemplate)

	return govy.NewRule(func(s string) error {
		if _, err := mail.ParseAddress(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringEmail).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid email address")
}

// StringURL ensures property's value is a valid URL as defined by [url.Parse] function.
// Unlike [URL] it does not impose any additional rules upon parsed [url.URL].
func StringURL() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.URLTemplate)

	return govy.NewRule(func(s string) error {
		u, err := url.Parse(s)
		if err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         "failed to parse URL: " + err.Error(),
			})
		}
		if err = validateURL(u); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringURL).
		WithMessageTemplate(tpl).
		WithDescription(urlDescription)
}

// StringMAC ensures property's value is a valid MAC address.
func StringMAC() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMACTemplate)

	return govy.NewRule(func(s string) error {
		if _, err := net.ParseMAC(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMAC).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringIP ensures property's value is a valid IP address.
func StringIP() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringIPTemplate)

	return govy.NewRule(func(s string) error {
		if ip := net.ParseIP(s); ip == nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringIP).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringIPv4 ensures property's value is a valid IPv4 address.
func StringIPv4() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringIPv4Template)

	return govy.NewRule(func(s string) error {
		if ip := net.ParseIP(s); ip == nil || ip.To4() == nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringIPv4).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringIPv6 ensures property's value is a valid IPv6 address.
func StringIPv6() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringIPv6Template)

	return govy.NewRule(func(s string) error {
		if ip := net.ParseIP(s); ip == nil || ip.To4() != nil || len(ip) != net.IPv6len {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringIPv6).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringCIDR ensures property's value is a valid CIDR notation IP address.
func StringCIDR() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringCIDRTemplate)

	return govy.NewRule(func(s string) error {
		if _, _, err := net.ParseCIDR(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCIDR).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringCIDRv4 ensures property's value is a valid CIDR notation IPv4 address.
func StringCIDRv4() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringCIDRv4Template)

	return govy.NewRule(func(s string) error {
		if ip, ipNet, err := net.ParseCIDR(s); err != nil || ip.To4() == nil || !ipNet.IP.Equal(ip) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCIDRv4).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringCIDRv6 ensures property's value is a valid CIDR notation IPv6 address.
func StringCIDRv6() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringCIDRv6Template)

	return govy.NewRule(func(s string) error {
		if ip, _, err := net.ParseCIDR(s); err != nil || ip.To4() != nil || len(ip) != net.IPv6len {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCIDRv6).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringEIN ensures the property's value is a United States Employer Identification Number (EIN)
// in NN-NNNNNNN format with a recognized prefix.
func StringEIN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringEINTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidEIN(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringEIN).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func isValidEIN(s string) bool {
	if len(s) != 10 ||
		s[2] != '-' ||
		!isASCIIDigit(s[0]) ||
		!isASCIIDigit(s[1]) ||
		!isASCIIDigit(s[3]) ||
		!isASCIIDigit(s[4]) ||
		!isASCIIDigit(s[5]) ||
		!isASCIIDigit(s[6]) ||
		!isASCIIDigit(s[7]) ||
		!isASCIIDigit(s[8]) ||
		!isASCIIDigit(s[9]) {
		return false
	}
	return isValidEINPrefix(s[:2])
}

func isValidEINPrefix(prefix string) bool {
	switch prefix {
	case "01", "02", "03", "04", "05", "06", "10", "11", "12", "13", "14", "15", "16",
		"20", "21", "22", "23", "24", "25", "26", "27",
		"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
		"40", "41", "42", "43", "44", "45", "46", "47", "48",
		"50", "51", "52", "53", "54", "55", "56", "57", "58", "59",
		"60", "61", "62", "63", "64", "65", "66", "67", "68",
		"71", "72", "73", "74", "75", "76", "77",
		"80", "81", "82", "83", "84", "85", "86", "87", "88",
		"90", "91", "92", "93", "94", "95", "98", "99":
		return true
	default:
		return false
	}
}

// StringSSN ensures the property's value is a United States Social Security Number (SSN)
// in NNN-NN-NNNN format.
// It rejects areas 000, 666, and 900-999, groups 00, and serials 0000.
// Rejecting areas 900-999 excludes every Individual Taxpayer Identification Number (ITIN).
func StringSSN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSSNTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidSSN(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringSSN).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func isValidSSN(s string) bool {
	if len(s) != 11 ||
		s[3] != '-' ||
		s[6] != '-' ||
		!isASCIIDigit(s[0]) ||
		!isASCIIDigit(s[1]) ||
		!isASCIIDigit(s[2]) ||
		!isASCIIDigit(s[4]) ||
		!isASCIIDigit(s[5]) ||
		!isASCIIDigit(s[7]) ||
		!isASCIIDigit(s[8]) ||
		!isASCIIDigit(s[9]) ||
		!isASCIIDigit(s[10]) {
		return false
	}
	return (s[0] != '0' || s[1] != '0' || s[2] != '0') &&
		(s[0] != '6' || s[1] != '6' || s[2] != '6') &&
		s[0] != '9' &&
		(s[4] != '0' || s[5] != '0') &&
		(s[7] != '0' || s[8] != '0' || s[9] != '0' || s[10] != '0')
}

// StringUUID ensures property's value is a valid UUID string as defined by [RFC 4122].
// It does not enforce a specific UUID version.
//
// [RFC 4122]: https://www.ietf.org/rfc/rfc4122.txt
func StringUUID() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMatchRegexpTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidUUID(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: uuidPattern,
			})
		}
		return nil
	}).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: uuidPattern,
		})).
		WithDetails("expected RFC-4122 compliant UUID string").
		WithExamples(
			"00000000-0000-0000-0000-000000000000",
			"e190c630-8873-11ee-b9d1-0242ac120002",
			"79258D24-01A7-47E5-ACBB-7E762DE52298",
		).
		WithErrorCode(ErrorCodeStringUUID)
}

// StringUUIDRFC4122 ensures the property's value is a Universally Unique Identifier (UUID)
// string as defined by RFC 4122.
// It requires the canonical 36-character form, a version from 1 through 5, and RFC 4122 variant bits.
func StringUUIDRFC4122() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringUUIDRFC4122Template)

	return govy.NewRule(func(s string) error {
		if !isValidUUIDRFC4122(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringUUIDRFC4122).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringUUIDv3 ensures the property's value is a version 3 Universally Unique Identifier (UUID)
// string as defined by RFC 4122.
func StringUUIDv3() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringUUIDv3Template)

	return govy.NewRule(func(s string) error {
		if !isValidUUIDVersion(s, '3') {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringUUIDv3).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringUUIDv4 ensures the property's value is a version 4 Universally Unique Identifier (UUID)
// string as defined by RFC 4122.
func StringUUIDv4() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringUUIDv4Template)

	return govy.NewRule(func(s string) error {
		if !isValidUUIDVersion(s, '4') {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringUUIDv4).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringUUIDv5 ensures the property's value is a version 5 Universally Unique Identifier (UUID)
// string as defined by RFC 4122.
func StringUUIDv5() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringUUIDv5Template)

	return govy.NewRule(func(s string) error {
		if !isValidUUIDVersion(s, '5') {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringUUIDv5).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func isValidUUID(s string) bool {
	if len(s) != 36 || s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return false
	}
	return isHexadecimalString(s[:8]) &&
		isHexadecimalString(s[9:13]) &&
		isHexadecimalString(s[14:18]) &&
		isHexadecimalString(s[19:23]) &&
		isHexadecimalString(s[24:])
}

func isValidUUIDRFC4122(s string) bool {
	return len(s) == 36 &&
		s[14] >= '1' && s[14] <= '5' &&
		isUUIDRFC4122Variant(s[19]) &&
		isValidUUID(s)
}

func isValidUUIDVersion(s string, version byte) bool {
	return len(s) == 36 &&
		s[14] == version &&
		isUUIDRFC4122Variant(s[19]) &&
		isValidUUID(s)
}

func isUUIDRFC4122Variant(b byte) bool {
	return b == '8' || b == '9' || b|0x20 == 'a' || b|0x20 == 'b'
}

func isHexadecimalString(s string) bool {
	for i := range len(s) {
		b := s[i]
		lower := b | 0x20
		if (b < '0' || b > '9') && (lower < 'a' || lower > 'f') {
			return false
		}
	}
	return true
}

// StringULID ensures the property's value is a 26-character Crockford Base32
// Universally Unique Lexicographically Sortable Identifier (ULID) string.
func StringULID() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringULIDTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidULID(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringULID).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func isValidULID(s string) bool {
	if len(s) != 26 {
		return false
	}
	if s[0] < '0' || s[0] > '7' {
		return false
	}
	for i := 1; i < len(s); i++ {
		if !isCrockfordBase32Byte(s[i]) {
			return false
		}
	}
	return true
}

func isCrockfordBase32Byte(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	lower := b | 0x20
	return lower >= 'a' && lower <= 'z' &&
		lower != 'i' && lower != 'l' && lower != 'o' && lower != 'u'
}

// StringMongoDBObjectID ensures the property's value is a 24-character
// hexadecimal MongoDB ObjectID.
func StringMongoDBObjectID() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMongoDBObjectIDTemplate)

	return govy.NewRule(func(s string) error {
		if !isMongoDBObjectID(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMongoDBObjectID).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringCreditCard ensures the property's value is a plausible digit-only
// payment card number. It requires a 13- to 19-digit value that passes the
// Luhn checksum and rejects all-same-digit values.
func StringCreditCard() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringCreditCardTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidCreditCard(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCreditCard).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringLuhnChecksum ensures the property's value is a digit-only string that
// passes the Luhn checksum.
func StringLuhnChecksum() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringLuhnChecksumTemplate)

	return govy.NewRule(func(s string) error {
		if !isLuhnChecksumValid(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringLuhnChecksum).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringBIC ensures the property's value matches the current Business
// Identifier Code (BIC) syntax.
func StringBIC() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBICTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidBIC(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBIC).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringBICISO93622014 ensures the property's value matches the ISO 9362:2014
// Business Identifier Code (BIC) syntax.
func StringBICISO93622014() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBICISO93622014Template)

	return govy.NewRule(func(s string) error {
		if !isValidBICISO93622014(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBICISO93622014).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringASCII ensures property's value contains only ASCII characters.
func StringASCII() govy.Rule[string] {
	return StringMatchRegexp(asciiRegexp()).WithErrorCode(ErrorCodeStringASCII)
}

// StringJSON ensures property's value is a valid JSON literal.
func StringJSON() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringJSONTemplate)

	return govy.NewRule(func(s string) error {
		if !json.Valid([]byte(s)) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringJSON).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringE164 ensures the property's value is a valid E.164 phone number.
func StringE164() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringE164Template)

	return govy.NewRule(func(s string) error {
		if !e164Regexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringE164).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

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

// StringBase64 ensures the property's value is a standard padded base64 string.
// It validates input with [base64.StdEncoding].
func StringBase64() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBase64Template)

	return govy.NewRule(func(s string) error {
		if !standardBase64Regexp().MatchString(s) ||
			!decodesBase64(base64.StdEncoding.Strict(), s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBase64).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringBase64URL ensures the property's value is a URL-safe padded base64 string.
// It validates input with [base64.URLEncoding].
func StringBase64URL() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBase64URLTemplate)

	return govy.NewRule(func(s string) error {
		if !base64URLRegexp().MatchString(s) ||
			!decodesBase64(base64.URLEncoding.Strict(), s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBase64URL).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringBase64RawURL ensures the property's value is a URL-safe base64 string without padding.
// It validates input with [base64.RawURLEncoding].
func StringBase64RawURL() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBase64RawURLTemplate)

	return govy.NewRule(func(s string) error {
		if !base64RawURLRegexp().MatchString(s) ||
			!decodesBase64(base64.RawURLEncoding.Strict(), s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBase64RawURL).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringHexadecimal ensures the property's value is a hexadecimal string.
// It allows an optional `0x` or `0X` prefix.
func StringHexadecimal() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringHexadecimalTemplate)

	return govy.NewRule(func(s string) error {
		if !isHexadecimal(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringHexadecimal).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func decodesBase64(encoding *base64.Encoding, s string) bool {
	_, err := encoding.DecodeString(s)
	return err == nil
}

func isHexadecimal(s string) bool {
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}
	if len(s) == 0 {
		return false
	}
	for i := range len(s) {
		c := s[i]
		if ('0' <= c && c <= '9') ||
			('a' <= c && c <= 'f') ||
			('A' <= c && c <= 'F') {
			continue
		}
		return false
	}
	return true
}

// StringMD5 ensures the property's value is a lowercase hexadecimal MD5 digest.
func StringMD5() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMD5Template)

	return govy.NewRule(func(s string) error {
		if !isLowerHexadecimal(s, 32) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMD5).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringSHA256 ensures the property's value is a lowercase hexadecimal SHA-256 digest.
func StringSHA256() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSHA256Template)

	return govy.NewRule(func(s string) error {
		if !isLowerHexadecimal(s, 64) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringSHA256).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringSHA384 ensures the property's value is a lowercase hexadecimal SHA-384 digest.
func StringSHA384() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSHA384Template)

	return govy.NewRule(func(s string) error {
		if !isLowerHexadecimal(s, 96) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringSHA384).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringSHA512 ensures the property's value is a lowercase hexadecimal SHA-512 digest.
func StringSHA512() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSHA512Template)

	return govy.NewRule(func(s string) error {
		if !isLowerHexadecimal(s, 128) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringSHA512).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringJWT ensures the property's value is a JSON Web Token (JWT) represented
// using [JWS Compact Serialization].
// It validates the three base64url-encoded segments, JSON object header,
// JSON object claims set, and required `alg` header.
// JWTs represented using [JWE Compact Serialization] are not accepted.
// It does not verify the signature, algorithm trust, or claim values.
//
// [JWE Compact Serialization]: https://datatracker.ietf.org/doc/html/rfc7516#section-3.1
// [JWS Compact Serialization]: https://datatracker.ietf.org/doc/html/rfc7515#section-3.1
func StringJWT() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringJWTTemplate)

	return govy.NewRule(func(s string) error {
		if err := validateJWTUsingJWS(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringJWT).
		WithMessageTemplate(tpl).
		WithDescription("string must be a JSON Web Token (JWT) using JWS Compact Serialization")
}

// StringContains ensures the property's value contains all the provided substrings.
func StringContains(substrings ...string) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringContainsTemplate)

	return govy.NewRule(func(s string) error {
		matched := true
		for _, substr := range substrings {
			if !strings.Contains(s, substr) {
				matched = false
				break
			}
		}
		if !matched {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: substrings,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringContains).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: substrings,
		}))
}

// StringExcludes ensures the property's value does not contain any of the provided substrings.
func StringExcludes(substrings ...string) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringExcludesTemplate)

	return govy.NewRule(func(s string) error {
		for _, substr := range substrings {
			if strings.Contains(s, substr) {
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					PropertyValue:   s,
					ComparisonValue: substrings,
				})
			}
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringExcludes).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: substrings,
		}))
}

// StringStartsWith ensures the property's value starts with one of the provided prefixes.
func StringStartsWith(prefixes ...string) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringStartsWithTemplate)

	return govy.NewRule(func(s string) error {
		matched := false
		for _, prefix := range prefixes {
			if strings.HasPrefix(s, prefix) {
				matched = true
				break
			}
		}
		if !matched {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: prefixes,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringStartsWith).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: prefixes,
		}))
}

// StringEndsWith ensures the property's value ends with one of the provided suffixes.
func StringEndsWith(suffixes ...string) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringEndsWithTemplate)

	return govy.NewRule(func(s string) error {
		matched := false
		for _, suffix := range suffixes {
			if strings.HasSuffix(s, suffix) {
				matched = true
				break
			}
		}
		if !matched {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: suffixes,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringEndsWith).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: suffixes,
		}))
}

// StringTitle ensures each word in a string starts with a capital letter.
func StringTitle() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringTitleTemplate)

	return govy.NewRule(func(s string) error {
		if len(s) == 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		prev := ' '
		for _, r := range s {
			if isStringSeparator(prev) {
				if !unicode.IsUpper(r) && !isStringSeparator(r) {
					return govy.NewRuleErrorTemplate(govy.TemplateVars{
						PropertyValue: s,
					})
				}
			}
			prev = r
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringTitle).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

type stringGitRefTemplateVars struct {
	GitRefEmpty           bool
	GitRefEndsWithDot     bool
	GitRefAtLeastOneSlash bool
	GitRefEmptyPart       bool
	GitRefStartsWithDash  bool
	GitRefForbiddenChars  bool
}

// StringGitRef ensures a git reference name follows the [git-check-ref-format] rules.
//
// It is important to note that this function does not check if the reference exists in the repository.
// It only checks if the reference name is valid.
// This functions does not support the '--refspec-pattern', '--normalize', and '--allow-onelevel' options.
//
// Git imposes the following rules on how references are named:
//
//  1. They can include slash '/' for hierarchical (directory) grouping, but no
//     slash-separated component can begin with a dot '.' or end with the
//     sequence '.lock'.
//  2. They must contain at least one '/'. This enforces the presence of a
//     category (e.g. 'heads/', 'tags/'), but the actual names are not restricted.
//  3. They cannot have ASCII control characters (i.e. bytes whose values are
//     lower than '\040', or '\177' DEL).
//  4. They cannot have '?', '*', '[', ' ', '~', '^', ', '\t', '\n', '@{', '\\' and '..',
//  5. They cannot begin or end with a slash '/'.
//  6. They cannot end with a '.'.
//  7. They cannot be the single character '@'.
//  8. 'HEAD' is an allowed special name.
//
// Slightly modified version of [go-git] implementation, kudos to the authors!
//
// [git-check-ref-format] :https://git-scm.com/docs/git-check-ref-format
// [go-git]: https://github.com/go-git/go-git/blob/95afe7e1cdf71c59ee8a71971fac71880020a744/plumbing/reference.go#L167
func StringGitRef() govy.Rule[string] {
	type tplVars = stringGitRefTemplateVars
	tpl := messagetemplates.Get(messagetemplates.StringGitRefTemplate)

	return govy.NewRule(func(s string) error {
		if len(s) == 0 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Custom:        tplVars{GitRefEmpty: true},
			})
		}
		if s == "HEAD" {
			return nil
		}
		if strings.HasSuffix(s, ".") {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Custom:        tplVars{GitRefEndsWithDot: true},
			})
		}
		parts := strings.Split(s, "/")
		if len(parts) < 2 {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Custom:        tplVars{GitRefAtLeastOneSlash: true},
			})
		}
		isBranch := strings.HasPrefix(s, "refs/heads/")
		isTag := strings.HasPrefix(s, "refs/tags/")
		for _, part := range parts {
			if len(part) == 0 {
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					PropertyValue: s,
					Custom:        tplVars{GitRefEmptyPart: true},
				})
			}
			if (isBranch || isTag) && strings.HasPrefix(part, "-") {
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					PropertyValue: s,
					Custom:        tplVars{GitRefStartsWithDash: true},
				})
			}
			if part == "@" ||
				strings.HasPrefix(part, ".") ||
				strings.HasSuffix(part, ".lock") ||
				stringContainsGitRefForbiddenChars(part) {
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					PropertyValue: s,
					Custom:        tplVars{GitRefForbiddenChars: true},
				})
			}
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringGitRef).
		WithMessageTemplate(tpl).
		WithDetails("see https://git-scm.com/docs/git-check-ref-format for more information on Git reference naming rules").
		WithDescription("string must be a valid git reference")
}

// StringFileSystemPath ensures the property's value is an existing file system path.
func StringFileSystemPath() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringFileSystemPathTemplate)

	return govy.NewRule(func(s string) error {
		if _, err := osStatFile(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         handleFilePathError(err).Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringFileSystemPath).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringFilePath ensures the property's value is a file system path pointing to an existing file.
func StringFilePath() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringFilePathTemplate)

	return govy.NewRule(func(s string) error {
		info, err := osStatFile(s)
		if err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         handleFilePathError(err).Error(),
			})
		}
		if info.IsDir() {
			return errFilePathNotFile
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringFilePath).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringDirPath ensures the property's value is a file system path pointing to an existing directory.
func StringDirPath() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringDirPathTemplate)

	return govy.NewRule(func(s string) error {
		info, err := osStatFile(s)
		if err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         handleFilePathError(err).Error(),
			})
		}
		if !info.IsDir() {
			return errFilePathNotDir
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringDirPath).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringMatchFileSystemPath ensures the property's value matches the provided file path pattern.
// It uses [filepath.Match] to match the pattern. The native function comes with some limitations,
// most notably it does not support '**' recursive expansion.
// It does not check if the file path exists on the file system.
func StringMatchFileSystemPath(pattern string) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMatchFileSystemPathTemplate)

	return govy.NewRule(func(s string) error {
		ok, err := filepath.Match(pattern, s)
		if err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: pattern,
				Error:           err.Error(),
			})
		}
		if !ok {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: pattern,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMatchFileSystemPath).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: pattern,
		}))
}

// StringRegexp ensures the property's value is a valid regular expression.
// The accepted regular expression syntax must comply to RE2.
// It is described at https://golang.org/s/re2syntax, except for \C.
// For an overview of the syntax, see [regexp/syntax] package.
//
// [regexp/syntax]: https://pkg.go.dev/regexp/syntax
func StringRegexp() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringRegexpTemplate)

	return govy.NewRule(func(s string) error {
		if _, err := regexp.Compile(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringRegexp).
		WithMessageTemplate(tpl).
		WithDetails(`the regular expression syntax must comply to RE2, it is described at https://golang.org/s/re2syntax, except for \C; for an overview of the syntax, see https://pkg.go.dev/regexp/syntax`).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringCrontab ensures the property's value is a valid crontab schedule expression.
// For more details on cron expressions read [crontab manual] and visit [crontab.guru].
//
// [crontab manual]: https://www.man7.org/linux/man-pages/man5/crontab.5.html
// [crontab.guru]: https://crontab.guru
func StringCrontab() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringCrontabTemplate)

	return govy.NewRule(func(s string) error {
		if err := parseCrontab(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCrontab).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringDateTime ensures the property's value is a valid date and time in the specified layout.
//
// The layout must be a valid time format string as defined by [time.Parse],
// an example of which is [time.RFC3339].
func StringDateTime(layout string) govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringDateTimeTemplate)

	return govy.NewRule(func(s string) error {
		if _, err := time.Parse(layout, s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				Error:           err.Error(),
				ComparisonValue: layout,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringDateTime).
		WithMessageTemplate(tpl).
		WithDetails("date and time format follows Go's time layout, see https://pkg.go.dev/time#Layout for more details").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{
			ComparisonValue: layout,
		}))
}

// StringTimeZone ensures the property's value is a valid time zone name which
// uniquely identifies a time zone in the IANA Time Zone database.
// Example: "America/New_York", "Europe/London".
//
// Under the hood [time.LoadLocation] is called to parse the zone.
// The native function allows empty string and 'Local' keyword to be supplied.
// However, these two options are explicitly forbidden by [StringTimeZone].
//
// Furthermore, the time zone data is not readily available in one predefined place.
// [time.LoadLocation] looks for the IANA Time Zone database in specific places,
// please refer to its documentation for more information.
func StringTimeZone() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringTimeZoneTemplate)

	return govy.NewRule(func(s string) error {
		if s == "" || s == "Local" {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		if _, err := time.LoadLocation(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringTimeZone).
		WithMessageTemplate(tpl).
		WithExamples("UTC", "America/New_York", "Europe/Warsaw").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringAlpha ensures the property's value consists only of ASCII letters.
func StringAlpha() govy.Rule[string] {
	return StringMatchRegexp(alphaRegexp()).
		WithErrorCode(ErrorCodeStringAlpha)
}

// StringAlphanumeric ensures the property's value consists only of ASCII letters and numbers.
func StringAlphanumeric() govy.Rule[string] {
	return StringMatchRegexp(alphanumericRegexp()).
		WithErrorCode(ErrorCodeStringAlphanumeric)
}

// StringAlphaUnicode ensures the property's value consists only of Unicode letters.
func StringAlphaUnicode() govy.Rule[string] {
	return StringMatchRegexp(alphaUnicodeRegexp()).
		WithErrorCode(ErrorCodeStringAlphaUnicode)
}

// StringAlphanumericUnicode ensures the property's value consists only of Unicode letters and numbers.
func StringAlphanumericUnicode() govy.Rule[string] {
	return StringMatchRegexp(alphanumericUnicodeRegexp()).
		WithErrorCode(ErrorCodeStringAlphanumericUnicode)
}

// StringFQDN ensures the property's value is a fully qualified domain name (FQDN).
func StringFQDN() govy.Rule[string] {
	return StringMatchRegexp(fqdnRegexp()).
		WithErrorCode(ErrorCodeStringFQDN)
}

type stringKubernetesQualifiedNameTemplateVars struct {
	EmptyPrefixPart bool
	PrefixLength    bool
	PrefixRegexp    bool
	TooManyParts    bool
	EmptyNamePart   bool
	NamePartLength  bool
	NamePartRegexp  bool
}

const (
	maxK8sSubdomainPrefixPartLength = 253
	maxK8sQualifiedNamePartLength   = 63
)

// StringKubernetesQualifiedName ensures the property's value is a valid "qualified name"
// as defined by [Kubernetes validation].
// The qualified name is used in various parts of the Kubernetes system, examples:
//   - annotation names
//   - label names
//
// [Kubernetes validation]: https://github.com/kubernetes/kubernetes/blob/55573a0739785292e62b32a748c0b0735ff963ba/staging/src/k8s.io/apimachinery/pkg/util/validation/validation.go#L41
func StringKubernetesQualifiedName() govy.RuleSet[string] {
	return govy.NewRuleSet(
		StringLength(1, maxK8sSubdomainPrefixPartLength+1+maxK8sQualifiedNamePartLength),
		stringKubernetesQualifiedNameRule(),
	).
		Cascade(govy.CascadeModeStop).
		WithErrorCode(ErrorCodeStringKubernetesQualifiedName)
}

func stringKubernetesQualifiedNameRule() govy.Rule[string] {
	type tplVars = stringKubernetesQualifiedNameTemplateVars
	tpl := messagetemplates.Get(messagetemplates.StringKubernetesQualifiedNameTemplate)

	return govy.NewRule(func(s string) error {
		parts := strings.Split(s, "/")
		var name string
		switch len(parts) {
		case 1:
			name = parts[0]
		case 2:
			var prefix string
			prefix, name = parts[0], parts[1]
			switch {
			case len(prefix) == 0:
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					PropertyValue: s,
					Custom:        tplVars{EmptyPrefixPart: true},
				})
			case len(prefix) > maxK8sSubdomainPrefixPartLength:
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					PropertyValue:   s,
					ComparisonValue: maxK8sSubdomainPrefixPartLength,
					Custom:          tplVars{PrefixLength: true},
				})
			case !rfc1123DnsSubdomainRegexp().MatchString(prefix):
				return govy.NewRuleErrorTemplate(govy.TemplateVars{
					PropertyValue:   s,
					ComparisonValue: rfc1123DnsSubdomainRegexp().String(),
					Custom:          tplVars{PrefixRegexp: true},
				})
			}
		default:
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Custom:        tplVars{TooManyParts: true},
			})
		}

		switch {
		case len(name) == 0:
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Custom:        tplVars{EmptyNamePart: true},
			})
		case len(name) > maxK8sQualifiedNamePartLength:
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: maxK8sQualifiedNamePartLength,
				Custom:          tplVars{NamePartLength: true},
			})
		case !k8sQualifiedNamePartRegexp().MatchString(name):
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   s,
				ComparisonValue: k8sQualifiedNamePartRegexp().String(),
				Custom:          tplVars{NamePartRegexp: true},
			})
		}
		return nil
	}).
		WithMessageTemplate(tpl).
		WithDetails("Kubernetes Qualified Name must consist of alphanumeric characters, '-', '_' or '.', "+
			"and must start and end with an alphanumeric character with an optional DNS subdomain prefix and '/'").
		WithExamples("my.domain/MyName", "MyName", "my.name", "123-abc").
		WithDescription("string must be a Kubernetes Qualified Name")
}

func isValidCreditCard(s string) bool {
	if len(s) < 13 || len(s) > 19 {
		return false
	}
	if allSameDigit(s) {
		return false
	}
	return isLuhnChecksumValid(s)
}

func isLuhnChecksumValid(s string) bool {
	if s == "" {
		return false
	}

	sum := 0
	shouldDouble := false
	for i := len(s) - 1; i >= 0; i-- {
		digit := s[i] - '0'
		if digit > 9 {
			return false
		}
		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += int(digit)
		shouldDouble = !shouldDouble
	}
	return sum%10 == 0
}

func allSameDigit(s string) bool {
	if s == "" {
		return false
	}
	first := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] != first {
			return false
		}
	}
	return true
}

func isValidBIC(s string) bool {
	return isValidBICCode(s)
}

func isValidBICISO93622014(s string) bool {
	return isValidBICCode(s)
}

func isValidBICCode(s string) bool {
	if len(s) != 8 && len(s) != 11 {
		return false
	}
	if !isValidBICCountryCode(s[4:6]) {
		return false
	}
	for i := range 4 {
		if !isUpperASCIIAlphanumeric(s[i]) {
			return false
		}
	}
	for i := 6; i < len(s); i++ {
		if !isUpperASCIIAlphanumeric(s[i]) {
			return false
		}
	}
	return true
}

func isValidBICCountryCode(s string) bool {
	switch s {
	case "XK":
		// XK is included as a common Kosovo exception.
		return true
	case "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ",
		"BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ",
		"CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ",
		"DE", "DJ", "DK", "DM", "DO", "DZ",
		"EC", "EE", "EG", "EH", "ER", "ES", "ET",
		"FI", "FJ", "FK", "FM", "FO", "FR",
		"GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY",
		"HK", "HM", "HN", "HR", "HT", "HU",
		"ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT",
		"JE", "JM", "JO", "JP",
		"KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ",
		"LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY",
		"MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ",
		"NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ",
		"OM",
		"PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY",
		"QA",
		"RE", "RO", "RS", "RU", "RW",
		"SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ",
		"TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ",
		"UA", "UG", "UM", "US", "UY", "UZ",
		"VA", "VC", "VE", "VG", "VI", "VN", "VU",
		"WF", "WS",
		"YE", "YT",
		"ZA", "ZM", "ZW":
		return true
	default:
		return false
	}
}

func isUpperASCIIAlphanumeric(b byte) bool {
	return b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
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

func isMongoDBObjectID(s string) bool {
	if len(s) != 24 {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if '0' <= c && c <= '9' ||
			'a' <= c && c <= 'f' ||
			'A' <= c && c <= 'F' {
			continue
		}
		return false
	}
	return true
}

var gitRefDisallowedStrings = map[rune]struct{}{
	'\\': {}, '?': {}, '*': {}, '[': {}, ' ': {}, '~': {}, '^': {}, ':': {}, '\t': {}, '\n': {},
}

// stringContainsGitRefForbiddenChars is a brute force method to check if a string contains
// any of the Git reference forbidden characters.
func stringContainsGitRefForbiddenChars(s string) bool {
	for i, c := range s {
		if c == '\177' || (c >= '\000' && c <= '\037') {
			return true
		}
		// Check for '..' and '@{'.
		if c == '.' && i < len(s)-1 && s[i+1] == '.' ||
			c == '@' && i < len(s)-1 && s[i+1] == '{' {
			return true
		}
		if _, ok := gitRefDisallowedStrings[c]; !ok {
			continue
		}
		return true
	}
	return false
}

func osStatFile(path string) (os.FileInfo, error) {
	if strings.TrimSpace(path) == "" {
		return nil, errFilePathEmpty
	}
	hasSeparatorSuffix := strings.HasSuffix(path, string(filepath.Separator))
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("failed to get user home directory: %w", err)
		}
		path = home + string(filepath.Separator) + path[1:]
	}
	path = filepath.Clean(path)
	// If the path ends with a separator, we need to add it back after cleaning.
	if hasSeparatorSuffix {
		path += string(filepath.Separator)
	}
	return os.Stat(path)
}

var (
	errFilePathNotExists = errors.New("path does not exist")
	errFilePathNoPerm    = errors.New("permission to inspect path denied")
	errFilePathEmpty     = errors.New("path does not exist")
	errFilePathNotFile   = errors.New("path must point to a file and not to a directory")
	errFilePathNotDir    = errors.New("path must point to a directory and not to a file")
)

func handleFilePathError(err error) error {
	var pathErr *os.PathError
	if !errors.As(err, &pathErr) {
		return err
	}
	if errors.Is(err, os.ErrNotExist) {
		return errFilePathNotExists
	}
	if errors.Is(err, os.ErrPermission) {
		return errFilePathNoPerm
	}
	return err
}

func isASCIIDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLowerHexadecimal(s string, length int) bool {
	if len(s) != length {
		return false
	}
	for i := range len(s) {
		c := s[i]
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			return false
		}
	}
	return true
}
