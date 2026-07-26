package rules

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	jwtHeaderB64  = "b64"
	jwtHeaderCrit = "crit"
	jwtHeaderAlg  = "alg"
)

func validateJWTUsingJWS(s string) error {
	headerSegment, remaining, ok := strings.Cut(s, ".")
	if !ok {
		return fmt.Errorf("expected exactly 3 JWT segments")
	}
	claimsSegment, signatureSegment, ok := strings.Cut(remaining, ".")
	if !ok || strings.Contains(signatureSegment, ".") {
		return fmt.Errorf("expected exactly 3 JWT segments")
	}

	headers, err := decodeJWTJSONObject(headerSegment, "JWT header")
	if err != nil {
		return err
	}

	alg, err := getJWTAlgorithm(headers)
	if err != nil {
		return err
	}
	if err = validateJWSB64Header(headers); err != nil {
		return err
	}
	if err = validateJWTJSONObject(claimsSegment, "JWT claims set"); err != nil {
		return err
	}
	if alg == "none" {
		if signatureSegment == "" {
			return nil
		}
		return fmt.Errorf(`JWT signature segment must be empty when alg is "none"`)
	}
	if signatureSegment == "" {
		return fmt.Errorf(`JWT signature segment must not be empty unless alg is "none"`)
	}
	if err = validateJWTBase64URLSegment(signatureSegment, "JWT signature"); err != nil {
		return err
	}
	return nil
}

// validateJWSB64Header validates the `b64` JWS Header Parameter for JWTs.
// [RFC 7797 Section 3] defines it as an optional JSON boolean defaulting to
// `true`, [RFC 7797 Section 6] requires a present `b64` to be listed in the
// `crit` array defined by [RFC 7515 Section 4.1.11], and [RFC 7797 Section 7]
// prohibits JWTs from setting `b64` to `false`.
//
// [RFC 7515 Section 4.1.11]: https://datatracker.ietf.org/doc/html/rfc7515#section-4.1.11
// [RFC 7797 Section 3]: https://datatracker.ietf.org/doc/html/rfc7797#section-3
// [RFC 7797 Section 6]: https://datatracker.ietf.org/doc/html/rfc7797#section-6
// [RFC 7797 Section 7]: https://datatracker.ietf.org/doc/html/rfc7797#section-7
func validateJWSB64Header(headers map[string]json.RawMessage) error {
	b64Raw, ok := headers[jwtHeaderB64]
	if !ok {
		return nil
	}

	b64Raw = bytes.TrimSpace(b64Raw)
	if bytes.Equal(b64Raw, []byte("false")) {
		return fmt.Errorf(`JWT header must not set "b64" to false`)
	}
	if !bytes.Equal(b64Raw, []byte("true")) {
		return fmt.Errorf(`JWT header "b64" must be a boolean`)
	}

	rawCritical, ok := headers[jwtHeaderCrit]
	var critical []string
	if !ok || json.Unmarshal(rawCritical, &critical) != nil {
		return fmt.Errorf(`JWT header "crit" must be an array containing "b64" when "b64" is present`)
	}
	for _, name := range critical {
		if name == jwtHeaderB64 {
			return nil
		}
	}
	return fmt.Errorf(`JWT header "crit" must be an array containing "b64" when "b64" is present`)
}

func getJWTAlgorithm(headers map[string]json.RawMessage) (string, error) {
	rawAlgorithm, ok := headers[jwtHeaderAlg]
	if !ok {
		return "", fmt.Errorf(`JWT header must contain an "alg" string`)
	}

	var algorithm string
	if err := json.Unmarshal(rawAlgorithm, &algorithm); err != nil || algorithm == "" {
		return "", fmt.Errorf(`JWT header must contain an "alg" string`)
	}
	for i := range len(algorithm) {
		if algorithm[i] > unicode.MaxASCII {
			return "", fmt.Errorf(`JWT header must contain an "alg" string`)
		}
	}
	return algorithm, nil
}

func decodeJWTJSONObject(segment, segmentName string) (map[string]json.RawMessage, error) {
	decoded, err := decodeJWTJSON(segment, segmentName)
	if err != nil {
		return nil, err
	}
	return unmarshalJWTJSONObject(decoded, segmentName)
}

func validateJWTJSONObject(segment, segmentName string) error {
	decoded, err := decodeJWTJSON(segment, segmentName)
	if err != nil {
		return err
	}
	trimmed := bytes.TrimSpace(decoded)
	if len(trimmed) > 0 && trimmed[0] == '{' && json.Valid(decoded) {
		return nil
	}

	_, err = unmarshalJWTJSONObject(decoded, segmentName)
	return err
}

func unmarshalJWTJSONObject(decoded []byte, segmentName string) (map[string]json.RawMessage, error) {
	var object map[string]json.RawMessage
	if err := json.Unmarshal(decoded, &object); err != nil {
		return nil, fmt.Errorf("%s segment must contain a JSON object: %w", segmentName, err)
	}
	if object == nil {
		return nil, fmt.Errorf("%s segment must contain a JSON object", segmentName)
	}
	return object, nil
}

func decodeJWTJSON(segment, segmentName string) ([]byte, error) {
	decoded, err := decodeJWTBase64URLSegment(segment, segmentName)
	if err != nil {
		return nil, err
	}
	if !utf8.Valid(decoded) {
		return nil, fmt.Errorf("%s segment must contain valid UTF-8 JSON", segmentName)
	}
	return decoded, nil
}

func decodeJWTBase64URLSegment(segment, segmentName string) ([]byte, error) {
	if err := validateJWTBase64URLAlphabet(segment, segmentName); err != nil {
		return nil, err
	}

	decoded, err := base64.RawURLEncoding.DecodeString(segment)
	if err != nil {
		return nil, fmt.Errorf("%s segment must be base64url encoded without padding: %w", segmentName, err)
	}
	return decoded, nil
}

func validateJWTBase64URLSegment(segment, segmentName string) error {
	if err := validateJWTBase64URLAlphabet(segment, segmentName); err != nil {
		return err
	}
	// An unpadded Base64URL string using the allowed alphabet is decodable
	// unless its length is 1 modulo 4.
	if len(segment)%4 == 1 {
		err := base64.CorruptInputError(len(segment) - 1)
		return fmt.Errorf("%s segment must be base64url encoded without padding: %w", segmentName, err)
	}
	return nil
}

func validateJWTBase64URLAlphabet(segment, segmentName string) error {
	if segment == "" {
		return fmt.Errorf("%s segment must not be empty", segmentName)
	}
	for i := range len(segment) {
		c := segment[i]
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') ||
			c == '-' || c == '_' {
			continue
		}
		return fmt.Errorf("%s segment must be base64url encoded without padding", segmentName)
	}
	return nil
}
