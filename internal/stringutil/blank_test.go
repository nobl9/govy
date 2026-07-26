package stringutil

import (
	"runtime"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestIsBlank(t *testing.T) {
	t.Parallel()

	// Unicode 15.0.0 White_Space property:
	// https://www.unicode.org/Public/15.0.0/ucd/PropList.txt
	testCases := map[string]string{
		"empty string":              "",
		"character tabulation":      "\u0009",
		"line feed":                 "\u000A",
		"line tabulation":           "\u000B",
		"form feed":                 "\u000C",
		"carriage return":           "\u000D",
		"space":                     "\u0020",
		"next line":                 "\u0085",
		"no-break space":            "\u00A0",
		"ogham space mark":          "\u1680",
		"en quad":                   "\u2000",
		"em quad":                   "\u2001",
		"en space":                  "\u2002",
		"em space":                  "\u2003",
		"three-per-em space":        "\u2004",
		"four-per-em space":         "\u2005",
		"six-per-em space":          "\u2006",
		"figure space":              "\u2007",
		"punctuation space":         "\u2008",
		"thin space":                "\u2009",
		"hair space":                "\u200A",
		"line separator":            "\u2028",
		"paragraph separator":       "\u2029",
		"narrow no-break space":     "\u202F",
		"medium mathematical space": "\u205F",
		"ideographic space":         "\u3000",
		"mixed whitespace":          " \t\n\r\v\f\u0085\u00A0\u1680\u2000\u2028\u2029\u202F\u205F\u3000",
	}
	for name, input := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.True(t, IsBlank(input))
		})
	}
}

func TestIsBlank_NonBlank(t *testing.T) {
	t.Parallel()

	testCases := map[string]string{
		"ASCII letter":               "a",
		"ASCII punctuation":          ".",
		"non-space control":          "\x00",
		"zero-width space":           "\u200B",
		"word joiner":                "\u2060",
		"byte order mark":            "\uFEFF",
		"replacement character":      "\uFFFD",
		"invalid UTF-8":              "\xff",
		"nonblank after whitespace":  " \u3000a",
		"nonblank before whitespace": "a \u3000",
		"nonblank amid whitespace":   "\u3000a ",
		"emoji":                      "x ☺",
	}
	for name, input := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.False(t, IsBlank(input))
		})
	}
}

func BenchmarkIsBlank(b *testing.B) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "empty", input: "", expected: true},
		{name: "ASCII text", input: "govy", expected: false},
		{name: "ASCII whitespace", input: strings.Repeat(" ", 128), expected: true},
		{
			name:     "ASCII padded text",
			input:    strings.Repeat(" ", 64) + "govy" + strings.Repeat(" ", 64),
			expected: false,
		},
		{name: "Unicode text", input: "Ｇｏｖｙ", expected: false},
		{name: "Unicode whitespace", input: strings.Repeat("\u3000", 128), expected: true},
		{
			name:     "Unicode padded text",
			input:    strings.Repeat("\u3000", 64) + "Ｇｏｖｙ" + strings.Repeat("\u3000", 64),
			expected: false,
		},
	}
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.Run("IsBlank", func(b *testing.B) {
				if actual := IsBlank(tc.input); actual != tc.expected {
					b.Fatalf("IsBlank(%q) = %t, expected %t", tc.input, actual, tc.expected)
				}
				var result bool
				for b.Loop() {
					result = IsBlank(tc.input)
				}
				runtime.KeepAlive(result)
			})
			b.Run("TrimSpace", func(b *testing.B) {
				if actual := strings.TrimSpace(tc.input) == ""; actual != tc.expected {
					b.Fatalf("strings.TrimSpace(%q) == empty is %t, expected %t", tc.input, actual, tc.expected)
				}
				var result bool
				for b.Loop() {
					result = strings.TrimSpace(tc.input) == ""
				}
				runtime.KeepAlive(result)
			})
		})
	}
}
