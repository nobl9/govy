package jsonpath

import (
	"fmt"
	"unicode/utf8"
)

// segmentKind identifies the type of a path segment.
type segmentKind uint8

const (
	// segmentName is a named path segment or map key, e.g. "metadata".
	segmentName segmentKind = iota
	// segmentRoot is the JSONPath root selector, rendered as `$`.
	segmentRoot
	// segmentIndex is an array index, e.g. `[0]`.
	segmentIndex
	// segmentUnknownIndex is an unknown array index, rendered as `[]`.
	segmentUnknownIndex
	// segmentValueWildcard represents value wildcard selectors: `*` and `[*]`.
	segmentValueWildcard
	// segmentKeyWildcard represents the govy map key wildcard selector `*~`.
	segmentKeyWildcard
)

// segment is a single component of a [Path].
type segment struct {
	kind  segmentKind
	name  string // used by [segmentName], [segmentValueWildcard]
	index uint   // used by [segmentIndex]
}

// EscapeSegment accepts a single named path segment and escapes any special characters.
// Examples:
//
//	EscapeSegment("foo") --> "foo"
//	EscapeSegment("'foo'") --> "['\'foo\'']"
//	EscapeSegment("foo.bar") --> "['foo.bar']"
func EscapeSegment(segment string) string {
	shouldWrap := !isMemberNameShorthand(segment)
	segment = escapeCharacters(segment)
	if shouldWrap {
		segment = "['" + segment + "']"
	}
	return segment
}

// escapeCharacters has been based on the [net/url] package.
func escapeCharacters(s string) string {
	escapedCount := 0
	for _, r := range s {
		if shouldEscape(r) {
			escapedCount += escapeSize(r) - utf8.RuneLen(r)
		}
	}
	if escapedCount == 0 {
		return s
	}

	var buf [64]byte
	var t []byte

	required := len(s) + escapedCount
	if required <= len(buf) {
		t = buf[:required]
	} else {
		t = make([]byte, required)
	}

	j := 0
	for _, r := range s {
		switch {
		case shouldEscape(r):
			t[j] = '\\'
			j++
			switch r {
			case '\b':
				t[j] = 'b'
			case '\f':
				t[j] = 'f'
			case '\n':
				t[j] = 'n'
			case '\t':
				t[j] = 't'
			case '\r':
				t[j] = 'r'
			case '\'', '\\':
				t[j] = byte(r)
			default:
				copy(t[j:], fmt.Sprintf("u%04X", r))
				j += 4
			}
			j++
		default:
			j += utf8.EncodeRune(t[j:], r)
		}
	}
	return string(t)
}

func shouldEscape(r rune) bool {
	switch r {
	case '\'', '\\', '\b', '\f', '\n', '\t', '\r':
		return true
	default:
		return r >= 0 && r < ' '
	}
}

func escapeSize(r rune) int {
	switch r {
	case '\'', '\\', '\b', '\f', '\n', '\t', '\r':
		return 2
	default:
		return 6
	}
}

func isMemberNameShorthand(s string) bool {
	if s == "" {
		return false
	}
	i := 0
	for _, r := range s {
		if i == 0 {
			if !isNameFirst(r) {
				return false
			}
		} else if !isNameChar(r) {
			return false
		}
		i++
	}
	return true
}

func isNameChar(r rune) bool {
	return isNameFirst(r) || r >= '0' && r <= '9'
}

func isNameFirst(r rune) bool {
	switch {
	case r >= 'A' && r <= 'Z':
		return true
	case r >= 'a' && r <= 'z':
		return true
	case r == '_':
		return true
	case r >= 0x80 && r <= 0xD7FF:
		return true
	case r >= 0xE000 && r <= 0x10FFFF:
		return true
	default:
		return false
	}
}
