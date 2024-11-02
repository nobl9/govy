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
func StringDNSLabel() govy.Rule[string] {
	return StringMatchRegexp(rfc1123DnsLabelRegexp(), "my-name", "123-abc").
		WithDetails("an RFC-1123 compliant label name must consist of lower case alphanumeric characters or '-'," +
			" and must start and end with an alphanumeric character").
		WithErrorCode(ErrorCodeStringDNSLabel)
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

// StringGitRef errors.
var (
	errGitRefEmpty           = errors.New("git reference cannot be empty")
	errGitRefEndsWithDot     = errors.New("git reference must not end with a '.'")
	errGitRefAtLeastOneSlash = errors.New("git reference must contain at least one '/'")
	errGitRefEmptyPart       = errors.New("git reference must not have empty parts")
	errGitRefStartsWithDash  = errors.New("git branch and tag references must not start with '-'")
	errGitRefForbiddenChars  = errors.New("git reference contains forbidden characters")
)

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
	msg := "string must be a valid git reference"
	return govy.NewRule(func(s string) error {
		if len(s) == 0 {
			return errGitRefEmpty
		}
		if s == "HEAD" {
			return nil
		}
		if strings.HasSuffix(s, ".") {
			return errGitRefEndsWithDot
		}
		parts := strings.Split(s, "/")
		if len(parts) < 2 {
			return errGitRefAtLeastOneSlash
		}
		isBranch := strings.HasPrefix(s, "refs/heads/")
		isTag := strings.HasPrefix(s, "refs/tags/")
		for _, part := range parts {
			if len(part) == 0 {
				return errGitRefEmptyPart
			}
			if (isBranch || isTag) && strings.HasPrefix(part, "-") {
				return errGitRefStartsWithDash
			}
			if part == "@" ||
				strings.HasPrefix(part, ".") ||
				strings.HasSuffix(part, ".lock") ||
				stringContainsGitRefForbiddenChars(part) {
				return errGitRefForbiddenChars
			}
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringGitRef).
		WithDetails("see https://git-scm.com/docs/git-check-ref-format for more information on Git reference naming rules").
		WithDescription(msg)
}

// StringFileSystemPath ensures the property's value is an existing file system path.
func StringFileSystemPath() govy.Rule[string] {
	msg := "string must be an existing file system path"
	return govy.NewRule(func(s string) error {
		if _, err := osStatFile(s); err != nil {
			return handleFilePathError(err, msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringFileSystemPath).
		WithDescription(msg)
}

// StringFilePath ensures the property's value is a file system path pointing to an existing file.
func StringFilePath() govy.Rule[string] {
	msg := "string must be a file system path to an existing file"
	return govy.NewRule(func(s string) error {
		info, err := osStatFile(s)
		if err != nil {
			return handleFilePathError(err, msg)
		}
		if info.IsDir() {
			return errFilePathNotFile
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringFilePath).
		WithDescription(msg)
}

// StringDirPath ensures the property's value is a file system path pointing to an existing directory.
func StringDirPath() govy.Rule[string] {
	msg := "string must be a file system path to an existing directory"
	return govy.NewRule(func(s string) error {
		info, err := osStatFile(s)
		if err != nil {
			return handleFilePathError(err, msg)
		}
		if !info.IsDir() {
			return errFilePathNotDir
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringDirPath).
		WithDescription(msg)
}

// StringMatchFileSystemPath ensures the property's value matches the provided file path pattern.
// It uses [filepath.Match] to match the pattern. The native function comes with some limitations,
// most notably it does not support '**' recursive expansion.
// It does not check if the file path exists on the file system.
func StringMatchFileSystemPath(pattern string) govy.Rule[string] {
	msg := fmt.Sprintf("string must match file path pattern: '%s'", pattern)
	return govy.NewRule(func(s string) error {
		ok, err := filepath.Match(pattern, s)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMatchFileSystemPath).
		WithDescription(msg)
}

// StringRegexp ensures the property's value is a valid regular expression.
// The accepted regular expression syntax must comply to RE2.
// It is described at https://golang.org/s/re2syntax, except for \C.
// For an overview of the syntax, see [regexp/syntax] package.
//
// [regexp/syntax]: https://pkg.go.dev/regexp/syntax
func StringRegexp() govy.Rule[string] {
	msg := "string must be a valid regular expression"
	return govy.NewRule(func(s string) error {
		if _, err := regexp.Compile(s); err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringRegexp).
		// nolint: lll
		WithDetails(`the regular expression syntax must comply to RE2, it is described at https://golang.org/s/re2syntax, except for \C; for an overview of the syntax, see https://pkg.go.dev/regexp/syntax`).
		WithDescription(msg)
}

// StringCrontab ensures the property's value is a valid crontab schedule expression.
// For more details on cron expressions read [crontab manual] and visit [crontab.guru].
//
// [crontab manual]: https://www.man7.org/linux/man-pages/man5/crontab.5.html
// [crontab.guru]: https://crontab.guru
func StringCrontab() govy.Rule[string] {
	msg := "string must be a valid cron schedule expression"
	return govy.NewRule(parseCrontab).
		WithMessage(msg).
		WithErrorCode(ErrorCodeStringCrontab)
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

func handleFilePathError(err error, msg string) error {
	var pathErr *os.PathError
	if !errors.As(err, &pathErr) {
		return err
	}
	if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("%s: %w", msg, errFilePathNotExists)
	}
	if errors.Is(err, os.ErrPermission) {
		return fmt.Errorf("%s: %w", msg, errFilePathNoPerm)
	}
	return fmt.Errorf("%s: %w", msg, err)
}
