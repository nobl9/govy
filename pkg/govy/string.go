package validation

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

func StringNotEmpty() SingleRule[string] {
	msg := "string cannot be empty"
	return NewSingleRule(func(s string) error {
		if len(strings.TrimSpace(s)) == 0 {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringNotEmpty).
		WithDescription(msg)
}

func StringMatchRegexp(re *regexp.Regexp, examples ...string) SingleRule[string] {
	msg := fmt.Sprintf("string must match regular expression: '%s'", re.String())
	if len(examples) > 0 {
		msg += " " + prettyExamples(examples)
	}
	return NewSingleRule(func(s string) error {
		if !re.MatchString(s) {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMatchRegexp).
		WithDescription(msg)
}

func StringDenyRegexp(re *regexp.Regexp, examples ...string) SingleRule[string] {
	msg := fmt.Sprintf("string must not match regular expression: '%s'", re.String())
	if len(examples) > 0 {
		msg += " " + prettyExamples(examples)
	}
	return NewSingleRule(func(s string) error {
		if re.MatchString(s) {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringDenyRegexp).
		WithDescription(msg)
}

var dns1123SubdomainRegexp = regexp.MustCompile("^[a-z0-9]([-a-z0-9]*[a-z0-9])?$")

func StringIsDNSSubdomain() RuleSet[string] {
	return NewRuleSet(
		StringLength(1, 63),
		StringMatchRegexp(dns1123SubdomainRegexp, "my-name", "123-abc").
			WithDetails("a DNS-1123 compliant name must consist of lower case alphanumeric characters or '-',"+
				" and must start and end with an alphanumeric character"),
	).WithErrorCode(ErrorCodeStringIsDNSSubdomain)
}

var validUUIDRegex = regexp.
	MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

func StringUUID() SingleRule[string] {
	return StringMatchRegexp(validUUIDRegex,
		"00000000-0000-0000-0000-000000000000",
		"e190c630-8873-11ee-b9d1-0242ac120002",
		"79258D24-01A7-47E5-ACBB-7E762DE52298").
		WithDetails("expected RFC-4122 compliant UUID string").
		WithErrorCode(ErrorCodeStringUUID)
}

var asciiRegexp = regexp.MustCompile("^[\x00-\x7F]*$")

func StringASCII() SingleRule[string] {
	return StringMatchRegexp(asciiRegexp).WithErrorCode(ErrorCodeStringASCII)
}

func StringDescription() SingleRule[string] {
	return StringLength(0, 1050).WithErrorCode(ErrorCodeStringDescription)
}

func StringURL() SingleRule[string] {
	return NewSingleRule(func(v string) error {
		u, err := url.Parse(v)
		if err != nil {
			return errors.Wrap(err, "failed to parse URL")
		}
		return validateURL(u)
	}).
		WithErrorCode(ErrorCodeStringURL).
		WithDescription(urlDescription)
}

func StringJSON() SingleRule[string] {
	msg := "string must be a valid JSON"
	return NewSingleRule(func(s string) error {
		if !json.Valid([]byte(s)) {
			return errors.New(msg)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringJSON).
		WithDescription(msg)
}

func StringContains(substrings ...string) SingleRule[string] {
	msg := "string must contain the following substrings: " + prettyStringList(substrings)
	return NewSingleRule(func(s string) error {
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

func StringStartsWith(prefixes ...string) SingleRule[string] {
	var msg string
	if len(prefixes) == 1 {
		msg = fmt.Sprintf("string must start with '%s' prefix", prefixes[0])
	} else {
		msg = "string must start with one of the following prefixes: " + prettyStringList(prefixes)
	}
	return NewSingleRule(func(s string) error {
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
