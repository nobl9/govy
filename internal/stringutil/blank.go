package stringutil

import (
	"unicode"
	"unicode/utf8"
)

// IsBlank reports whether s is empty or consists entirely of Unicode
// whitespace characters recognized by [unicode.IsSpace].
//
// It can stop at the first non-whitespace character.
// Checking whether [strings.TrimSpace] returns an empty string may
// also inspect the trailing boundary of a nonblank string. Both approaches are
// allocation-free, and all-whitespace inputs require a complete scan either way.
// The custom function presents a minor improvement over the builtin.
//
// [strings.TrimSpace]: https://pkg.go.dev/strings#TrimSpace
func IsBlank(s string) bool {
	for i := 0; i < len(s); {
		c := s[i]
		if c < utf8.RuneSelf {
			switch c {
			case ' ', '\t', '\n', '\r', '\v', '\f':
				i++
				continue
			default:
				return false
			}
		}

		r, size := utf8.DecodeRuneInString(s[i:])
		if !unicode.IsSpace(r) {
			return false
		}
		i += size
	}

	return true
}
