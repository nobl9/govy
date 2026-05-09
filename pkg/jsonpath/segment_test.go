package jsonpath_test

import (
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/jsonpath"
)

func TestEscapeSegment(t *testing.T) {
	tests := map[string]struct {
		pathSegment string
		expected    string
	}{
		"empty":                     {"", "['']"},
		"not dots":                  {"foo", "foo"},
		"underscore":                {"_", "_"},
		"unicode shorthand":         {"東京", "東京"},
		"digit after first":         {"foo1", "foo1"},
		"digit first":               {"1foo", "['1foo']"},
		"hyphen":                    {"foo-bar", "['foo-bar']"},
		"tilde":                     {"~", "['~']"},
		"one dot":                   {"foo.bar", "['foo.bar']"},
		"two dots":                  {"foo.bar.baz", "['foo.bar.baz']"},
		"single quote":              {"'foo", `['\'foo']`},
		"single quotes":             {"'foo'", `['\'foo\'']`},
		"backslash":                 {`foo\bar`, `['foo\\bar']`},
		"backspace":                 {"\bfoo", "['\\bfoo']"},
		"form feed":                 {"\ffoo", "['\\ffoo']"},
		"null byte":                 {"\x00foo", "['\\u0000foo']"},
		"fake escaped path":         {"['foo.bar']", `['[\'foo.bar\']']`},
		"left bracket":              {"[foo", "['[foo']"},
		"right bracket":             {"foo]", "['foo]']"},
		"both brackets":             {"[foo]", "['[foo]']"},
		"rackets":                   {"[foo.bar]", "['[foo.bar]']"},
		"mixed brackets and quotes": {"'[foo]'", `['\'[foo]\'']`},
		"single whitespace":         {"foo bar", "['foo bar']"},
		"multiple whitespaces":      {"  foo.bar ", "['  foo.bar ']"},
		"tab":                       {"\tfoo", "['\\tfoo']"},
		"newline":                   {"\nfoo", "['\\nfoo']"},
		"carriage return":           {"\rfoo", "['\\rfoo']"},
		"large string with escape":  {"\n" + strings.Repeat("l", 1000), "['\\n" + strings.Repeat("l", 1000) + "']"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := jsonpath.EscapeSegment(tc.pathSegment)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
