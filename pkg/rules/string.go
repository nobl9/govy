package rules

import (
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

// StringUUID ensures property's value is a valid UUID string as defined by [RFC 4122].
// It does not enforce a specific UUID version.
//
// [RFC 4122]: https://www.ietf.org/rfc/rfc4122.txt
func StringUUID() govy.Rule[string] {
	return StringMatchRegexp(uuidRegexp()).
		WithDetails("expected RFC-4122 compliant UUID string").
		WithExamples(
			"00000000-0000-0000-0000-000000000000",
			"e190c630-8873-11ee-b9d1-0242ac120002",
			"79258D24-01A7-47E5-ACBB-7E762DE52298",
		).
		WithErrorCode(ErrorCodeStringUUID)
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
