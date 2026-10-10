package ecmaregex

import (
	"errors"
	"fmt"
	"regexp/syntax"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// ErrUnsupported identifies Go regular expression constructs that Translate
// cannot represent with the supported ECMA-262 syntax.
var ErrUnsupported = errors.New("unsupported Go regular expression construct")

type precedence uint8

const (
	precedenceAlternate precedence = iota
	precedenceConcat
	precedenceRepeat
	precedenceAtom
)

var canonicalClasses = []struct {
	body   string
	ranges []rune
}{
	{ranges: []rune{'0', '9'}, body: `0-9`},
	{ranges: []rune{'\t', '\n', '\f', '\r', ' ', ' '}, body: `\t\n\f\r `},
	{ranges: []rune{'0', '9', 'A', 'Z', '_', '_', 'a', 'z'}, body: `A-Za-z0-9_`},
}

// Translate returns an ECMA-262 pattern that preserves the match behavior of
// the Go regular expression pattern.
//
// Translate returns a parse error for invalid Go syntax and [ErrUnsupported]
// for multiline anchors. Other Go flags are encoded in the returned pattern.
func Translate(pattern string) (string, error) {
	re, err := syntax.Parse(pattern, syntax.Perl)
	if err != nil {
		return "", fmt.Errorf("parse Go regular expression: %w", err)
	}
	var b strings.Builder
	if err := writeRegexp(&b, re, precedenceAlternate); err != nil {
		return "", err
	}
	return b.String(), nil
}

func writeRegexp(b *strings.Builder, re *syntax.Regexp, parent precedence) error {
	current := regexpPrecedence(re)
	group := current < parent
	if group {
		b.WriteString(`(?:`)
	}
	if err := writeRegexpBody(b, re); err != nil {
		return err
	}
	if group {
		b.WriteByte(')')
	}
	return nil
}

func writeRegexpBody(b *strings.Builder, re *syntax.Regexp) error {
	switch re.Op {
	case syntax.OpNoMatch:
		b.WriteString(`[^\s\S]`)
	case syntax.OpEmptyMatch:
		b.WriteString(`(?:)`)
	case syntax.OpLiteral:
		writeLiteral(b, re)
	case syntax.OpCharClass:
		writeCharacterClass(b, re.Rune)
	case syntax.OpAnyCharNotNL:
		b.WriteString(`[^\n]`)
	case syntax.OpAnyChar:
		b.WriteString(`[\s\S]`)
	case syntax.OpBeginLine:
		return unsupported("multiline beginning anchor")
	case syntax.OpEndLine:
		return unsupported("multiline end anchor")
	case syntax.OpBeginText:
		b.WriteByte('^')
	case syntax.OpEndText:
		b.WriteString(`(?![\s\S])`)
	case syntax.OpWordBoundary:
		b.WriteString(`\b`)
	case syntax.OpNoWordBoundary:
		b.WriteString(`\B`)
	case syntax.OpCapture:
		b.WriteString(`(?:`)
		if err := writeRegexp(b, re.Sub[0], precedenceAlternate); err != nil {
			return err
		}
		b.WriteByte(')')
	case syntax.OpStar:
		return writeRepetition(b, re, "*")
	case syntax.OpPlus:
		return writeRepetition(b, re, "+")
	case syntax.OpQuest:
		return writeRepetition(b, re, "?")
	case syntax.OpRepeat:
		return writeCountedRepetition(b, re)
	case syntax.OpConcat:
		for _, sub := range re.Sub {
			if err := writeRegexp(b, sub, precedenceConcat); err != nil {
				return err
			}
		}
	case syntax.OpAlternate:
		for i, sub := range re.Sub {
			if i > 0 {
				b.WriteByte('|')
			}
			if err := writeRegexp(b, sub, precedenceAlternate); err != nil {
				return err
			}
		}
	default:
		return unsupported(fmt.Sprintf("unknown syntax operation %d", re.Op))
	}
	return nil
}

func regexpPrecedence(re *syntax.Regexp) precedence {
	switch re.Op {
	case syntax.OpAlternate:
		return precedenceAlternate
	case syntax.OpConcat:
		return precedenceConcat
	case syntax.OpLiteral:
		if len(re.Rune) > 1 {
			return precedenceConcat
		}
		return precedenceAtom
	case syntax.OpStar, syntax.OpPlus, syntax.OpQuest, syntax.OpRepeat:
		return precedenceRepeat
	default:
		return precedenceAtom
	}
}

func writeRepetition(b *strings.Builder, re *syntax.Regexp, quantifier string) error {
	if err := writeRegexp(b, re.Sub[0], precedenceAtom); err != nil {
		return err
	}
	b.WriteString(quantifier)
	writeGreediness(b, re)
	return nil
}

func writeCountedRepetition(b *strings.Builder, re *syntax.Regexp) error {
	if err := writeRegexp(b, re.Sub[0], precedenceAtom); err != nil {
		return err
	}
	b.WriteByte('{')
	b.WriteString(strconv.Itoa(re.Min))
	if re.Max != re.Min {
		b.WriteByte(',')
		if re.Max >= 0 {
			b.WriteString(strconv.Itoa(re.Max))
		}
	}
	b.WriteByte('}')
	writeGreediness(b, re)
	return nil
}

func writeGreediness(b *strings.Builder, re *syntax.Regexp) {
	if re.Flags&syntax.NonGreedy != 0 {
		b.WriteByte('?')
	}
}

func writeLiteral(b *strings.Builder, re *syntax.Regexp) {
	for _, r := range re.Rune {
		if re.Flags&syntax.FoldCase != 0 {
			writeFoldedRune(b, r)
			continue
		}
		writeLiteralRune(b, r)
	}
}

func writeFoldedRune(b *strings.Builder, r rune) {
	runes := []rune{r}
	for next := unicode.SimpleFold(r); next != r; next = unicode.SimpleFold(next) {
		runes = append(runes, next)
	}
	if len(runes) == 1 {
		writeLiteralRune(b, r)
		return
	}
	slices.Sort(runes)
	writeCharacterClass(b, runesToRanges(runes))
}

func runesToRanges(runes []rune) []rune {
	ranges := make([]rune, 0, len(runes)*2)
	for _, r := range runes {
		if len(ranges) == 0 || r > ranges[len(ranges)-1]+1 {
			ranges = append(ranges, r, r)
			continue
		}
		ranges[len(ranges)-1] = r
	}
	return ranges
}

func writeCharacterClass(b *strings.Builder, ranges []rune) {
	if len(ranges) == 0 {
		b.WriteString(`[^\s\S]`)
		return
	}
	if isFullRuneRange(ranges) {
		b.WriteString(`[\s\S]`)
		return
	}
	if writeCanonicalClass(b, ranges) {
		return
	}

	positive := formatCharacterClass(ranges, false)
	negative := formatCharacterClass(complementRanges(ranges), true)
	if len(negative) < len(positive) {
		b.WriteString(negative)
		return
	}
	b.WriteString(positive)
}

func writeCanonicalClass(b *strings.Builder, ranges []rune) bool {
	for _, class := range canonicalClasses {
		if equalRunes(ranges, class.ranges) {
			b.WriteByte('[')
			b.WriteString(class.body)
			b.WriteByte(']')
			return true
		}
		if equalRunes(ranges, complementRanges(class.ranges)) {
			b.WriteString(`[^`)
			b.WriteString(class.body)
			b.WriteByte(']')
			return true
		}
	}
	return false
}

func formatCharacterClass(ranges []rune, negated bool) string {
	var b strings.Builder
	b.WriteByte('[')
	if negated {
		b.WriteByte('^')
	}
	for i := 0; i < len(ranges); i += 2 {
		lo, hi := ranges[i], ranges[i+1]
		writeClassRune(&b, lo)
		if lo == hi {
			continue
		}
		if hi != lo+1 {
			b.WriteByte('-')
		}
		writeClassRune(&b, hi)
	}
	b.WriteByte(']')
	return b.String()
}

func complementRanges(ranges []rune) []rune {
	complement := make([]rune, 0, len(ranges)+2)
	next := rune(0)
	for i := 0; i < len(ranges); i += 2 {
		lo, hi := ranges[i], ranges[i+1]
		if next < lo {
			complement = append(complement, next, lo-1)
		}
		if hi == unicode.MaxRune {
			return complement
		}
		next = hi + 1
	}
	return append(complement, next, unicode.MaxRune)
}

func isFullRuneRange(ranges []rune) bool {
	return len(ranges) == 2 && ranges[0] == 0 && ranges[1] == unicode.MaxRune
}

func equalRunes(left, right []rune) bool {
	if len(left) != len(right) {
		return false
	}
	for i := range left {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func writeLiteralRune(b *strings.Builder, r rune) {
	if unicode.IsPrint(r) {
		if strings.ContainsRune(`\.+*?()|[]{}^$`, r) {
			b.WriteByte('\\')
		}
		b.WriteRune(r)
		return
	}
	writeEscapedRune(b, r)
}

func writeClassRune(b *strings.Builder, r rune) {
	if unicode.IsPrint(r) {
		if strings.ContainsRune(`\[]^-`, r) {
			b.WriteByte('\\')
		}
		b.WriteRune(r)
		return
	}
	writeEscapedRune(b, r)
}

func writeEscapedRune(b *strings.Builder, r rune) {
	switch r {
	case '\f':
		b.WriteString(`\f`)
	case '\n':
		b.WriteString(`\n`)
	case '\r':
		b.WriteString(`\r`)
	case '\t':
		b.WriteString(`\t`)
	case '\v':
		b.WriteString(`\v`)
	default:
		writeUnicodeEscape(b, r)
	}
}

func writeUnicodeEscape(b *strings.Builder, r rune) {
	hex := strconv.FormatInt(int64(r), 16)
	if r <= 0xffff {
		b.WriteString(`\u`)
		b.WriteString(strings.Repeat("0", 4-len(hex)))
		b.WriteString(hex)
		return
	}
	b.WriteString(`\u{`)
	b.WriteString(hex)
	b.WriteByte('}')
}

func unsupported(feature string) error {
	return fmt.Errorf("%w: %s", ErrUnsupported, feature)
}
