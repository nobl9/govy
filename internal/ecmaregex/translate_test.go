package ecmaregex

import (
	"errors"
	"regexp/syntax"
	"strings"
	"testing"
	"unicode"

	"github.com/nobl9/govy/internal/assert"
)

func TestTranslate(t *testing.T) {
	t.Parallel()
	testCases := map[string]struct {
		pattern  string
		expected string
	}{
		"empty": {
			pattern:  ``,
			expected: `(?:)`,
		},
		"never matches": {
			pattern:  `[^\x00-\x{10FFFF}]`,
			expected: `[^\s\S]`,
		},
		"literal": {
			pattern:  `abc`,
			expected: `abc`,
		},
		"quoted metacharacters": {
			pattern:  `\Q^$.*+?()[]{}|\E`,
			expected: `\^\$\.\*\+\?\(\)\[\]\{\}\|`,
		},
		"control characters": {
			pattern:  `\a\f\n\r\t\v\x08\x7f`,
			expected: `\u0007\f\n\r\t\v\u0008\u007f`,
		},
		"Unicode literals": {
			pattern:  `é😀`,
			expected: `é😀`,
		},
		"octal escape": {
			pattern:  `\123`,
			expected: `S`,
		},
		"Unicode escape": {
			pattern:  `\x{1F600}`,
			expected: `😀`,
		},
		"digit class": {
			pattern:  `\d`,
			expected: `[0-9]`,
		},
		"negated digit class": {
			pattern:  `\D`,
			expected: `[^0-9]`,
		},
		"space class": {
			pattern:  `\s`,
			expected: `[\t\n\f\r ]`,
		},
		"negated space class": {
			pattern:  `\S`,
			expected: `[^\t\n\f\r ]`,
		},
		"word class": {
			pattern:  `\w`,
			expected: `[A-Za-z0-9_]`,
		},
		"negated word class": {
			pattern:  `\W`,
			expected: `[^A-Za-z0-9_]`,
		},
		"POSIX class": {
			pattern:  `[[:alpha:]]`,
			expected: `[A-Za-z]`,
		},
		"adjacent class members": {
			pattern:  `[ab]`,
			expected: `[ab]`,
		},
		"class ranges": {
			pattern:  `[a-cx-z]`,
			expected: `[a-cx-z]`,
		},
		"negated class": {
			pattern:  `[^a-c]`,
			expected: `[^a-c]`,
		},
		"Unicode class range": {
			pattern:  `[α-ω]`,
			expected: `[α-ω]`,
		},
		"supplementary Unicode class range": {
			pattern:  `[😀-🙏]`,
			expected: `[😀-🙏]`,
		},
		"dot": {
			pattern:  `.`,
			expected: `[^\n]`,
		},
		"dot including newline": {
			pattern:  `(?s:.)`,
			expected: `[\s\S]`,
		},
		"text anchors": {
			pattern:  `^foo$`,
			expected: `^foo(?![\s\S])`,
		},
		"explicit text anchors": {
			pattern:  `\Afoo\z`,
			expected: `^foo(?![\s\S])`,
		},
		"word boundaries": {
			pattern:  `\bfoo\B`,
			expected: `\bfoo\B`,
		},
		"capturing group": {
			pattern:  `(foo)`,
			expected: `(?:foo)`,
		},
		"named group": {
			pattern:  `(?P<value>foo)`,
			expected: `(?:foo)`,
		},
		"alternation": {
			pattern:  `foo|bar`,
			expected: `foo|bar`,
		},
		"alternation in concatenation": {
			pattern:  `(?:foo|bar)baz`,
			expected: `(?:foo|bar)baz`,
		},
		"greedy repetitions": {
			pattern:  `a*b+c?`,
			expected: `a*b+c?`,
		},
		"lazy repetitions": {
			pattern:  `a*?b+?c??`,
			expected: `a*?b+?c??`,
		},
		"exact repetition": {
			pattern:  `a{2}`,
			expected: `a{2}`,
		},
		"bounded repetition": {
			pattern:  `a{2,4}?`,
			expected: `a{2,4}?`,
		},
		"unbounded repetition": {
			pattern:  `a{2,}`,
			expected: `a{2,}`,
		},
		"repeated concatenation": {
			pattern:  `(?:ab)*`,
			expected: `(?:ab)*`,
		},
		"repeated repetition": {
			pattern:  `(?:a*)*`,
			expected: `(?:a*)*`,
		},
		"lazy by default": {
			pattern:  `(?U:a*?b+)`,
			expected: `a*b+?`,
		},
		"case-folded literal": {
			pattern:  `(?i:k)`,
			expected: `[KkK]`,
		},
		"case-folded literal without variants": {
			pattern:  `(?i:1)`,
			expected: `1`,
		},
		"case-folded adjacent runes": {
			pattern:  `(?i:Ā)`,
			expected: `[Āā]`,
		},
		"case-folded literal sequence": {
			pattern:  `(?i:ab)`,
			expected: `[Aa][Bb]`,
		},
		"case-folded class": {
			pattern:  `(?i:[a-c])`,
			expected: `[A-Ca-c]`,
		},
		"literal quoted section": {
			pattern:  `\Qfoo.bar\E`,
			expected: `foo\.bar`,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual, err := Translate(tc.pattern)
			assert.Require(t, assert.NoError(t, err))
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestTranslate_UnicodeProperty(t *testing.T) {
	t.Parallel()
	actual, err := Translate(`\p{Greek}`)
	assert.Require(t, assert.NoError(t, err))
	expectedRegexp, err := syntax.Parse(`\p{Greek}`, syntax.Perl)
	assert.Require(t, assert.NoError(t, err))
	actualRegexp, err := syntax.Parse(actual, syntax.Perl)
	assert.Require(t, assert.NoError(t, err))
	assert.True(t, expectedRegexp.Equal(actualRegexp))
}

func TestTranslate_InvalidPattern(t *testing.T) {
	t.Parallel()
	actual, err := Translate(`[a-`)
	assert.Equal(t, "", actual)
	assert.ErrorContains(t, err, "parse Go regular expression")
}

func TestTranslate_Unsupported(t *testing.T) {
	t.Parallel()
	testCases := map[string]string{
		"multiline beginning anchor": `(?m)^foo`,
		"multiline end anchor":       `(?m)foo$`,
	}
	for name, pattern := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual, err := Translate(pattern)
			assert.Equal(t, "", actual)
			assert.True(t, errors.Is(err, ErrUnsupported))
			assert.ErrorContains(t, err, name)
		})
	}
}

func Test_writeRegexp_UnknownOperation(t *testing.T) {
	t.Parallel()
	var b strings.Builder
	err := writeRegexp(&b, &syntax.Regexp{Op: syntax.Op(255)}, precedenceAlternate)
	assert.True(t, errors.Is(err, ErrUnsupported))
	assert.ErrorContains(t, err, "unknown syntax operation")
}

func Test_writeRegexp_NoMatch(t *testing.T) {
	t.Parallel()
	var b strings.Builder
	err := writeRegexp(&b, &syntax.Regexp{Op: syntax.OpNoMatch}, precedenceAlternate)
	assert.Require(t, assert.NoError(t, err))
	assert.Equal(t, `[^\s\S]`, b.String())
}

func Test_writeRegexp_PropagatesUnsupported(t *testing.T) {
	t.Parallel()
	unsupported := &syntax.Regexp{Op: syntax.OpBeginLine}
	testCases := map[string]*syntax.Regexp{
		"capture": {
			Op:  syntax.OpCapture,
			Sub: []*syntax.Regexp{unsupported},
		},
		"alternation": {
			Op:  syntax.OpAlternate,
			Sub: []*syntax.Regexp{unsupported},
		},
		"simple repetition": {
			Op:  syntax.OpStar,
			Sub: []*syntax.Regexp{unsupported},
		},
		"counted repetition": {
			Op:  syntax.OpRepeat,
			Sub: []*syntax.Regexp{unsupported},
			Min: 1,
			Max: 2,
		},
	}
	for name, re := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var b strings.Builder
			err := writeRegexp(&b, re, precedenceAlternate)
			assert.True(t, errors.Is(err, ErrUnsupported))
		})
	}
}

func Test_writeCharacterClass_FullRuneRange(t *testing.T) {
	t.Parallel()
	var b strings.Builder
	writeCharacterClass(&b, []rune{0, unicode.MaxRune})
	assert.Equal(t, `[\s\S]`, b.String())
}
