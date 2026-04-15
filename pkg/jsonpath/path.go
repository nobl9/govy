package jsonpath

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const (
	jsonPathSeparator = '.'
	escapedChars      = string(jsonPathSeparator) + "[]' \t\n\r"
)

// Path is a builder for constructing valid [JSONPath] property paths.
// It ensures proper escaping and formatting of path segments.
// Internally it stores a sequence of typed segments; the string form
// is computed on demand by [Path.String].
//
// Examples:
//
//	jsonpath.Parse("metadata.name")            							// metadata.name
//	jsonpath.New().Name("metadata").Name("name")            // metadata.name
//	jsonpath.New().Name("metadata").Name("labels").Index(0) // metadata.labels[0]
//	jsonpath.New().Name("complex.key")                      // ['complex.key']
//
// [JSONPath]: https://www.rfc-editor.org/rfc/rfc9535.html
type Path struct {
	segments []segment
}

// New creates a new empty [Path].
func New() Path {
	return Path{}
}

// Parse parses a JSONPath string into a structured [Path].
// It handles dotted names, bracket notation, array indices, and wildcards.
// Malformed input is treated gracefully: unparsable content becomes a name segment.
func Parse(s string) Path {
	if s == "" {
		return Path{}
	}
	return Path{segments: parseSegments(s)}
}

// Name appends a named segment to the path, escaping special characters as needed.
func (p Path) Name(name string) Path {
	return p.appendSegment(segment{kind: segmentName, name: name})
}

// Index appends an array index segment to the path.
func (p Path) Index(index uint) Path {
	return p.appendSegment(segment{kind: segmentIndex, index: index})
}

// Key appends a map key segment to the path, escaping special characters as needed.
// Keys are stored as name segments since they render identically.
func (p Path) Key(key any) Path {
	return p.appendSegment(segment{kind: segmentName, name: fmt.Sprint(key)})
}

// Join appends another [Path] to this one.
func (p Path) Join(other Path) Path {
	if len(other.segments) == 0 {
		return p
	}
	if len(p.segments) == 0 {
		return other
	}
	joined := make([]segment, 0, len(p.segments)+len(other.segments))
	joined = append(joined, p.segments...)
	joined = append(joined, other.segments...)
	return Path{segments: joined}
}

// Equal reports whether two paths have identical segments.
func (p Path) Equal(other Path) bool {
	return slices.Equal(p.segments, other.segments)
}

// Compare returns -1, 0, or 1 comparing two paths lexicographically
// by their string representations.
func (p Path) Compare(other Path) int {
	return cmp.Compare(p.String(), other.String())
}

// UnknownIndex appends an unknown array index segment "[]" to the path.
// This is used when the actual index is not statically known.
func (p Path) UnknownIndex() Path {
	return p.appendSegment(segment{kind: segmentUnknownIndex})
}

// IsEmpty returns true if the path contains no segments.
func (p Path) IsEmpty() bool {
	return len(p.segments) == 0
}

// String returns the string representation of the path.
func (p Path) String() string {
	if len(p.segments) == 0 {
		return ""
	}
	var b strings.Builder
	for i, s := range p.segments {
		switch s.kind {
		case segmentName:
			rendered := EscapeSegment(s.name)
			if i > 0 && !strings.HasPrefix(rendered, "[") {
				b.WriteByte(jsonPathSeparator)
			}
			b.WriteString(rendered)
		case segmentIndex:
			b.WriteByte('[')
			b.WriteString(strconv.FormatUint(uint64(s.index), 10))
			b.WriteByte(']')
		case segmentUnknownIndex:
			b.WriteString("[]")
		case segmentWildcard:
			if i > 0 && !strings.HasPrefix(s.name, "[") {
				b.WriteByte(jsonPathSeparator)
			}
			b.WriteString(s.name)
		}
	}
	return b.String()
}

// MarshalText implements [encoding.TextMarshaler].
// This ensures Path serializes as a plain string in JSON struct fields.
func (p Path) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler].
func (p *Path) UnmarshalText(data []byte) error {
	*p = Parse(string(data))
	return nil
}

// appendSegment returns a new Path with the given segment appended.
func (p Path) appendSegment(s segment) Path {
	result := make([]segment, len(p.segments)+1)
	copy(result, p.segments)
	result[len(p.segments)] = s
	return Path{segments: result}
}

// parseSegments tokenizes a JSONPath string into segments.
func parseSegments(s string) []segment {
	var segments []segment
	i := 0
	for i < len(s) {
		switch s[i] {
		case '.':
			i++
		case '[':
			seg, end := parseBracketSegment(s, i)
			segments = append(segments, seg)
			i = end
		case '~':
			segments = append(segments, segment{kind: segmentWildcard, name: "~"})
			i++
		case '*':
			segments = append(segments, segment{kind: segmentWildcard, name: "*"})
			i++
		case '$':
			segments = append(segments, segment{kind: segmentWildcard, name: "$"})
			i++
		default:
			seg, end := parseNameSegment(s, i)
			segments = append(segments, seg)
			i = end
		}
	}
	return segments
}

// parseBracketSegment parses a bracket-delimited segment starting at position i.
// It handles: [N] (index), [] (unknown index), [*] (wildcard), ['name'] (quoted name).
// Returns the segment and the position after the closing ']'.
func parseBracketSegment(s string, i int) (seg segment, end int) {
	start := i + 1 // skip opening '['
	if start >= len(s) {
		return segment{kind: segmentName, name: s[i:]}, len(s)
	}
	// For quoted names, find the matching closing '] after the quote.
	if s[start] == '\'' {
		quoteEnd := findQuotedEnd(s, start)
		if quoteEnd == -1 {
			return segment{kind: segmentName, name: s[i:]}, len(s)
		}
		inner := s[start:quoteEnd]
		name := parseQuotedName(inner)
		// quoteEnd points after closing quote; expect ']' next.
		if quoteEnd < len(s) && s[quoteEnd] == ']' {
			return segment{kind: segmentName, name: name}, quoteEnd + 1
		}
		return segment{kind: segmentName, name: name}, quoteEnd
	}
	// Non-quoted: find closing bracket.
	closeIdx := strings.IndexByte(s[i:], ']')
	if closeIdx == -1 {
		return segment{kind: segmentName, name: s[i:]}, len(s)
	}
	closeIdx += i
	inner := s[i+1 : closeIdx]
	end = closeIdx + 1
	switch inner {
	case "":
		return segment{kind: segmentUnknownIndex}, end
	case "*":
		return segment{kind: segmentWildcard, name: "[*]"}, end
	default:
		v, err := strconv.ParseUint(inner, 10, 64)
		if err == nil {
			return segment{kind: segmentIndex, index: uint(v)}, end
		}
		return segment{kind: segmentName, name: inner}, end
	}
}

// findQuotedEnd finds the position after the closing quote in a single-quoted string.
// It handles backslash-escaped quotes. Returns the position of the closing quote + 1,
// or -1 if no closing quote is found.
func findQuotedEnd(s string, start int) int {
	// start points at the opening quote.
	for j := start + 1; j < len(s); j++ {
		if s[j] == '\\' {
			j++ // skip escaped character
			continue
		}
		if s[j] == '\'' {
			return j + 1
		}
	}
	return -1
}

// parseQuotedName extracts a name from a single-quoted bracket expression.
// Input is the content between [ and ], e.g. "'foo.bar'" or "'foo\'s'".
func parseQuotedName(inner string) string {
	// Strip surrounding quotes.
	if len(inner) >= 2 && inner[0] == '\'' && inner[len(inner)-1] == '\'' {
		inner = inner[1 : len(inner)-1]
	}
	return unescapeCharacters(inner)
}

// parseNameSegment parses a dot-separated name segment starting at position i.
// Returns the segment and the position after the name.
func parseNameSegment(s string, i int) (seg segment, end int) {
	end = i
	for end < len(s) && s[end] != jsonPathSeparator && s[end] != '[' {
		end++
	}
	return segment{kind: segmentName, name: s[i:end]}, end
}

// unescapeCharacters reverses the escaping done by [escapeCharacters].
func unescapeCharacters(s string) string {
	if !strings.ContainsRune(s, '\\') {
		return s
	}
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			i++
			switch s[i] {
			case 'n':
				b.WriteByte('\n')
			case 't':
				b.WriteByte('\t')
			case 'r':
				b.WriteByte('\r')
			default:
				b.WriteByte(s[i])
			}
			continue
		}
		b.WriteByte(s[i])
	}
	return b.String()
}
