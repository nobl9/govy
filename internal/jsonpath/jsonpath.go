package jsonpath

import (
	"strconv"
	"strings"
)

const (
	jsonPathSeparator = "."
	escapedChars      = jsonPathSeparator + "[] \t\n"
	slashEscapedChars = ""
)

// EscapeSegment accepts a path segment (not the entire path!) and escapes any special characters.
// Examples:
//
//	EscapeSegment("foo") --> "foo"
//	EscapeSegment("'foo'") --> "\'foo\'"
//	EscapeSegment("foo.bar") --> "['foo.bar']"
func EscapeSegment(segment string) string {
	shouldWrap := segment == "" || strings.ContainsAny(segment, escapedChars)
	segment = escapeCharacters(segment)
	if shouldWrap {
		segment = "['" + segment + "']"
	}
	return segment
}

// Join extends the JSONPath with a new segment.
// The segment can be a path in of itself, the segment is assumed to be escaped with [EscapeSegment].
// Example:
//
//	Join("foo.bar", "baz") --> "foo.bar.baz"
//	Join("foo.bar", "baz.foo") --> "foo.bar.baz.foo"
func Join(path, segment string) string {
	return joinPaths(path, segment, jsonPathSeparator)
}

// JoinArray extends the JSONPath with a new array segment.
// Example:
//
//	JoinArray("foo.bar", "[2]") --> "foo.bar[2]"
func JoinArray(path, segment string) string {
	return joinPaths(path, segment, "")
}

// NewArrayIndex creates a new array index path segment for the given index.
// Example:
//
//	NewArrayIndex(2) --> "[2]"
func NewArrayIndex(index int) string {
	return "[" + strconv.Itoa(index) + "]"
}

func joinPaths(pre, post, sep string) string {
	if pre == "" {
		return post
	}
	if post == "" {
		return pre
	}
	return pre + sep + post
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
