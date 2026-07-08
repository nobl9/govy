package rules

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// cspell:words base64url

var (
	semverRegexp = lazyRegexCompile(
		`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[A-Za-z-][0-9A-Za-z-]*)(?:\.(?:0|[1-9]\d*|\d*[A-Za-z-][0-9A-Za-z-]*))*))?(?:\+([0-9A-Za-z-]+(?:\.[0-9A-Za-z-]+)*))?$`,
	)
	cveRegexp  = lazyRegexCompile(`^CVE-(1999|2[0-9]{3})-(000[1-9]|00[1-9][0-9]|0[1-9][0-9]{2}|[1-9][0-9]{3,})$`)
	e164Regexp = lazyRegexCompile(`^\+[1-9][0-9]{1,14}$`)
	issnRegexp = lazyRegexCompile(`^[0-9]{4}-[0-9]{3}[0-9Xx]$`)
)

// StringJWT ensures the property's value is a compact JSON Web Token (JWT).
// It validates the three Base64URL-encoded segments, the JSON object header,
// the JSON object claims set, and the required "alg" header.
// It does not verify the signature, algorithm trust, or claim values.
func StringJWT() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringJWTTemplate)

	return govy.NewRule(func(s string) error {
		if err := validateJWT(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringJWT).
		WithMessageTemplate(tpl).
		WithDescription("string must be a compact JSON Web Token (JWT) with three Base64URL-encoded segments")
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

// StringE164 ensures the property's value is a valid E.164 phone number.
// It accepts normalized international notation with a leading "+" and up to 15 digits.
// It does not check whether the number is assigned or reachable.
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
		WithDescription("string must be a valid E.164 phone number with a leading '+' and up to 15 digits")
}

// StringISBN ensures the property's value is a valid ISBN-10 or ISBN-13.
// It accepts digits separated by single spaces or hyphens and validates the check digit.
// It does not validate ISBN registration-group or publisher ranges.
func StringISBN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISBNTemplate)

	return govy.NewRule(func(s string) error {
		if !isISBN10(s) && !isISBN13(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISBN).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid ISBN-10 or ISBN-13")
}

// StringISBN10 ensures the property's value is a valid ISBN-10.
// It accepts digits separated by single spaces or hyphens and validates the check digit.
// It does not validate ISBN registration-group or publisher ranges.
func StringISBN10() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISBN10Template)

	return govy.NewRule(func(s string) error {
		if !isISBN10(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISBN10).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid ISBN-10")
}

// StringISBN13 ensures the property's value is a valid ISBN-13.
// It accepts digits separated by single spaces or hyphens and validates the check digit.
// It does not validate ISBN registration-group or publisher ranges.
func StringISBN13() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISBN13Template)

	return govy.NewRule(func(s string) error {
		if !isISBN13(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISBN13).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid ISBN-13")
}

// StringISSN ensures the property's value is a valid ISSN.
// It accepts 4 digits, a hyphen, 3 digits, and a final check character.
func StringISSN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISSNTemplate)

	return govy.NewRule(func(s string) error {
		if !isISSN(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISSN).
		WithMessageTemplate(tpl).
		WithDescription("string must be a valid hyphenated ISSN")
}

func validateJWT(s string) error {
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return fmt.Errorf("expected 3 JWT segments")
	}

	header, err := decodeJWTJSONObject(parts[0], "JWT header")
	if err != nil {
		return err
	}
	if _, err = decodeJWTJSONObject(parts[1], "JWT claims set"); err != nil {
		return err
	}

	alg, err := getJWTAlgorithm(header)
	if err != nil {
		return err
	}
	if alg == "none" {
		if parts[2] == "" {
			return nil
		}
		return fmt.Errorf(`JWT signature segment must be empty when alg is "none"`)
	}
	if parts[2] == "" {
		return fmt.Errorf(`JWT signature segment must not be empty unless alg is "none"`)
	}
	if _, err = decodeJWTBase64URLSegment(parts[2], "JWT signature"); err != nil {
		return err
	}
	return nil
}

func getJWTAlgorithm(header map[string]json.RawMessage) (string, error) {
	rawAlgorithm, ok := header["alg"]
	if !ok {
		return "", fmt.Errorf(`JWT header must contain an "alg" string`)
	}

	var algorithm string
	if err := json.Unmarshal(rawAlgorithm, &algorithm); err != nil || algorithm == "" {
		return "", fmt.Errorf(`JWT header must contain an "alg" string`)
	}
	return algorithm, nil
}

func decodeJWTJSONObject(segment, segmentName string) (map[string]json.RawMessage, error) {
	decoded, err := decodeJWTBase64URLSegment(segment, segmentName)
	if err != nil {
		return nil, err
	}

	var object map[string]json.RawMessage
	if err = json.Unmarshal(decoded, &object); err != nil {
		return nil, fmt.Errorf("%s segment must contain a JSON object: %w", segmentName, err)
	}
	if object == nil {
		return nil, fmt.Errorf("%s segment must contain a JSON object", segmentName)
	}
	return object, nil
}

func decodeJWTBase64URLSegment(segment, segmentName string) ([]byte, error) {
	if segment == "" {
		return nil, fmt.Errorf("%s segment must not be empty", segmentName)
	}
	if strings.Contains(segment, "=") {
		return nil, fmt.Errorf("%s segment must be Base64URL encoded without padding", segmentName)
	}
	for i := 0; i < len(segment); i++ {
		c := segment[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') ||
			c == '-' || c == '_' {
			continue
		}
		return nil, fmt.Errorf("%s segment must be Base64URL encoded without padding", segmentName)
	}

	decoded, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return nil, fmt.Errorf("%s segment must be Base64URL encoded without padding: %w", segmentName, err)
	}
	return decoded, nil
}

func isISBN10(s string) bool {
	isbn, ok := normalizeISBN(s)
	if !ok || len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 9; i++ {
		if !isASCIIDigit(isbn[i]) {
			return false
		}
		sum += int(isbn[i]-'0') * (10 - i)
	}

	switch checkDigit := isbn[9]; {
	case isASCIIDigit(checkDigit):
		sum += int(checkDigit - '0')
	case checkDigit == 'X' || checkDigit == 'x':
		sum += 10
	default:
		return false
	}
	return sum%11 == 0
}

func isISBN13(s string) bool {
	isbn, ok := normalizeISBN(s)
	if !ok || len(isbn) != 13 || (!strings.HasPrefix(isbn, "978") && !strings.HasPrefix(isbn, "979")) {
		return false
	}

	sum := 0
	for i := 0; i < 12; i++ {
		if !isASCIIDigit(isbn[i]) {
			return false
		}
		weight := 1
		if i%2 != 0 {
			weight = 3
		}
		sum += int(isbn[i]-'0') * weight
	}
	if !isASCIIDigit(isbn[12]) {
		return false
	}
	return (10-sum%10)%10 == int(isbn[12]-'0')
}

func normalizeISBN(s string) (string, bool) {
	if s == "" {
		return "", false
	}

	var builder strings.Builder
	previousWasSeparator := false
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case isASCIIDigit(c) || c == 'X' || c == 'x':
			builder.WriteByte(c)
			previousWasSeparator = false
		case c == '-' || c == ' ':
			if i == 0 || i == len(s)-1 || previousWasSeparator {
				return "", false
			}
			previousWasSeparator = true
		default:
			return "", false
		}
	}
	return builder.String(), true
}

func isISSN(s string) bool {
	if !issnRegexp().MatchString(s) {
		return false
	}

	issn := strings.ReplaceAll(s, "-", "")
	sum := 0
	for i := 0; i < 7; i++ {
		if !isASCIIDigit(issn[i]) {
			return false
		}
		sum += int(issn[i]-'0') * (8 - i)
	}

	switch checkDigit := issn[7]; {
	case isASCIIDigit(checkDigit):
		sum += int(checkDigit - '0')
	case checkDigit == 'X' || checkDigit == 'x':
		sum += 10
	default:
		return false
	}
	return sum%11 == 0
}

func isASCIIDigit(c byte) bool {
	return c >= '0' && c <= '9'
}
