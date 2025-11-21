package jsonpath

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Type int

const (
	TypeRoot Type = iota
	TypeProperty
	TypeIndex
)

type Segment interface {
	fmt.Stringer
	Type() Type
}

type RootSegment struct{}

func (RootSegment) String() string {
	return "$"
}

func (RootSegment) Type() Type {
	return TypeRoot
}

type PropertySegment struct {
	Name string
}

func (s PropertySegment) String() string {
	const escapedChars = ".[] '\t\n\r"
	shouldWrap := s.Name == "" || strings.ContainsAny(s.Name, escapedChars)
	name := escapeCharacters(s.Name)
	if shouldWrap {
		return "['" + name + "']"
	}
	return name
}

func (PropertySegment) Type() Type {
	return TypeProperty
}

type IndexSegment struct {
	Index int
}

func (s IndexSegment) String() string {
	return fmt.Sprintf("[%d]", s.Index)
}

func (IndexSegment) Type() Type {
	return TypeIndex
}

type Path struct {
	segments []Segment
}

func (p Path) Segments() []Segment {
	return p.segments
}

// Append creates a new Path with the given segments appended to the current path.
func (p Path) Append(segments ...Segment) Path {
	newSegments := make([]Segment, len(p.segments)+len(segments))
	copy(newSegments, p.segments)
	copy(newSegments[len(p.segments):], segments)
	return Path{segments: newSegments}
}

// Parse parses a JSONPath string.
func Parse(pathStr string) (Path, error) {
	if pathStr == "" {
		return Path{}, errors.New("empty path")
	}

	path := Path{
		segments: []Segment{},
	}

	i := 0
	if strings.HasPrefix(pathStr, "$") {
		path.segments = append(path.segments, RootSegment{})
		i = 1
	}

	for i < len(pathStr) {
		switch pathStr[i] {
		case '.':
			i++
			// Check if it's bracket notation after dot
			if i < len(pathStr) && pathStr[i] == '[' {
				seg, consumed, err := parseBracket(pathStr[i:])
				if err != nil {
					return Path{}, err
				}
				path.segments = append(path.segments, seg)
				i += consumed
			} else {
				// Regular property name
				seg, consumed := parseProperty(pathStr[i:])
				path.segments = append(path.segments, seg)
				i += consumed
			}
		case '[':
			seg, consumed, err := parseBracket(pathStr[i:])
			if err != nil {
				return Path{}, err
			}
			path.segments = append(path.segments, seg)
			i += consumed
		default:
			seg, consumed := parseProperty(pathStr[i:])
			path.segments = append(path.segments, seg)
			i += consumed
		}
	}

	return path, nil
}

// String converts path back to string representation.
func (p Path) String() string {
	var sb strings.Builder
	for i, seg := range p.segments {
		switch v := seg.(type) {
		case RootSegment:
			sb.WriteString(v.String())
		case PropertySegment:
			str := v.String()
			if i > 0 && !strings.HasPrefix(str, "['") {
				sb.WriteString(".")
			}
			sb.WriteString(str)
		case IndexSegment:
			sb.WriteString(v.String())
		}
	}
	return sb.String()
}

// parseProperty parses a simple property name (up to . or [)
func parseProperty(s string) (Segment, int) {
	i := 0
	for i < len(s) && s[i] != '.' && s[i] != '[' {
		i++
	}
	return PropertySegment{
		Name: s[:i],
	}, i
}

// parseBracket parses ['string'] or [123]
func parseBracket(s string) (Segment, int, error) {
	if len(s) == 0 || s[0] != '[' {
		return nil, 0, fmt.Errorf("expected [")
	}

	// Check if it's a quoted string
	if len(s) > 1 && s[1] == '\'' {
		// Find the closing quote, handling escapes
		i := 2
		for i < len(s) {
			if s[i] == '\\' && i+1 < len(s) {
				i += 2
				continue
			}
			if s[i] == '\'' {
				if i+1 >= len(s) || s[i+1] != ']' {
					return nil, 0, fmt.Errorf("expected ] after closing quote")
				}
				name := unescapeCharacters(s[2:i])
				return PropertySegment{
					Name: name,
				}, i + 2, nil
			}
			i++
		}
		return nil, 0, fmt.Errorf("unclosed quoted string")
	}

	// Try to parse as index [123]
	closeIdx := strings.Index(s, "]")
	if closeIdx == -1 {
		return nil, 0, fmt.Errorf("unclosed bracket")
	}

	content := s[1:closeIdx]
	idx, err := strconv.Atoi(content)
	if err == nil {
		return IndexSegment{
			Index: idx,
		}, closeIdx + 1, nil
	}

	return nil, 0, fmt.Errorf("invalid bracket content: %s", content)
}

func unescapeCharacters(s string) string {
	if !strings.ContainsRune(s, '\\') {
		return s
	}

	result := strings.Builder{}
	result.Grow(len(s))

	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == '\\' && i+1 < len(s):
			switch s[i+1] {
			case 'n':
				result.WriteByte('\n')
			case 't':
				result.WriteByte('\t')
			case 'r':
				result.WriteByte('\r')
			default:
				result.WriteByte(s[i+1])
			}
			i++
		default:
			result.WriteByte(s[i])
		}
	}

	return result.String()
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

	var (
		buf [64]byte
		t   []byte
	)

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
