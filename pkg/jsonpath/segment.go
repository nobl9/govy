package jsonpath

import "strings"

// segmentKind identifies the type of a path segment.
type segmentKind uint8

const (
	// segmentName is a named path segment or map key, e.g. "metadata".
	segmentName segmentKind = iota
	// segmentRoot is the JSONPath root selector, rendered as $.
	segmentRoot
	// segmentIndex is an array index, e.g. [0].
	segmentIndex
	// segmentUnknownIndex is an unknown array index, rendered as [].
	segmentUnknownIndex
	// segmentWildcard covers [*], *, and ~ selectors.
	segmentWildcard
)

// segment is a single component of a [Path].
type segment struct {
	kind  segmentKind
	name  string // used by [segmentName], [segmentWildcard]
	index uint   // used by [segmentIndex]
}

// EscapeSegment accepts a single named path segment and escapes any special characters.
// Examples:
//
//	EscapeSegment("foo") --> "foo"
//	EscapeSegment("'foo'") --> "['\'foo\'']"
//	EscapeSegment("foo.bar") --> "['foo.bar']"
func EscapeSegment(segment string) string {
	shouldWrap := segment == "" || strings.ContainsAny(segment, escapedChars)
	segment = escapeCharacters(segment)
	if shouldWrap {
		segment = "['" + segment + "']"
	}
	return segment
}

// escapeCharacters has been based on the [net/url] package.
func escapeCharacters(s string) string {
	escapedCount := 0
	for i := range s {
		if shouldEscape(s[i]) {
			escapedCount++
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
	for i := range len(s) {
		switch c := s[i]; {
		case shouldEscape(c):
			t[j] = '\\'
			j++
			switch c {
			case '\n':
				t[j] = 'n'
			case '\t':
				t[j] = 't'
			case '\r':
				t[j] = 'r'
			default:
				t[j] = c
			}
			j++
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

func shouldEscape(r byte) bool {
	switch r {
	case '\'', '\n', '\t', '\r':
		return true
	default:
		return false
	}
}
