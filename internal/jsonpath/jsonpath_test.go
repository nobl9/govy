package jsonpath_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonpath"
)

func TestEscapeSegment(t *testing.T) {
	tests := map[string]struct {
		pathSegment string
		expected    string
	}{
		"empty":                     {"", "['']"},
		"not dots":                  {"foo", "foo"},
		"one dot":                   {"foo.bar", "['foo.bar']"},
		"two dots":                  {"foo.bar.baz", "['foo.bar.baz']"},
		"single quote":              {"'foo", `['\'foo']`},
		"single quotes":             {"'foo'", `['\'foo\'']`},
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

func TestJoin(t *testing.T) {
	tests := map[string]struct {
		path     string
		segment  string
		expected string
	}{
		"empty path and segment": {
			path:     "",
			segment:  "",
			expected: "",
		},
		"empty path": {
			path:     "",
			segment:  "foo",
			expected: "foo",
		},
		"empty segment": {
			path:     "foo",
			segment:  "",
			expected: "foo",
		},
		"path and segment": {
			path:     "foo",
			segment:  "bar",
			expected: "foo.bar",
		},
		"path and segment path": {
			path:     "foo",
			segment:  "bar.baz",
			expected: "foo.bar.baz",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := jsonpath.Join(tc.path, tc.segment)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestJoinArray(t *testing.T) {
	tests := map[string]struct {
		path     string
		segment  string
		expected string
	}{
		"empty path and segment": {
			path:     "",
			segment:  "",
			expected: "",
		},
		"empty path": {
			path:     "",
			segment:  "[1]",
			expected: "[1]",
		},
		"empty segment": {
			path:     "foo",
			segment:  "",
			expected: "foo",
		},
		"path and segment": {
			path:     "foo",
			segment:  "[1]",
			expected: "foo[1]",
		},
		"path and segment path": {
			path:     "foo",
			segment:  "[2].baz",
			expected: "foo[2].baz",
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := jsonpath.JoinArray(tc.path, tc.segment)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestNewArrayIndex(t *testing.T) {
	tests := []struct {
		index    int
		expected string
	}{
		{0, "[0]"},
		{1, "[1]"},
		{10, "[10]"},
	}
	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := jsonpath.NewArrayIndex(tc.index)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
