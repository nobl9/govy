package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

var validISBNTestCases = map[string]string{
	"isbn 10 hyphenated": "0-306-40615-2",
	"isbn 10 plain":      "0306406152",
	"isbn 10 x check":    "0-9752298-0-X",
	"isbn 10 spaced":     "0 9752298 0 x",
	"isbn 13 hyphenated": "978-0-306-40615-7",
	"isbn 13 plain":      "9780306406157",
	"isbn 13 grouped":    "978-3-16-148410-0",
}

var invalidISBNTestCases = map[string]string{
	"empty":                "",
	"isbn 10 failed check": "0-306-40615-3",
	"isbn 13 failed check": "978-0-306-40615-8",
	"isbn 13 x check":      "978030640615X",
	"repeated separator":   "978--0-306-40615-7",
	"letters":              "abc",
}

func TestStringISBN(t *testing.T) {
	rule := StringISBN()
	t.Run("valid inputs", func(t *testing.T) {
		for name, input := range validISBNTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		err := rule.Validate(invalidISBNTestCases["empty"])
		assert.EqualError(
			t,
			err,
			"string must be a valid International Standard Book Number (ISBN) in ISBN-10 or ISBN-13 format",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN))

		for name, input := range invalidISBNTestCases {
			t.Run(name, func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
	})
}

func BenchmarkStringISBN(b *testing.B) {
	benchmarkStringPublicationRule(b, StringISBN(), []string{
		"0-306-40615-2",
		"978-0-306-40615-7",
		"978--0-306-40615-7",
	})
}

var validISBN10TestCases = map[string]string{
	"hyphenated": "0-306-40615-2",
	"plain":      "0306406152",
	"x check":    "0-9752298-0-X",
	"spaced":     "0 9752298 0 x",
}

var invalidISBN10TestCases = map[string]string{
	"empty":              "",
	"failed check":       "0-306-40615-3",
	"isbn 13":            "978-0-306-40615-7",
	"isbn 13 plain":      "9780306406157",
	"repeated separator": "0-306--40615-2",
}

func TestStringISBN10(t *testing.T) {
	rule := StringISBN10()
	t.Run("valid inputs", func(t *testing.T) {
		for name, input := range validISBN10TestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		err := rule.Validate(invalidISBN10TestCases["empty"])
		assert.EqualError(
			t,
			err,
			"string must be a valid International Standard Book Number (ISBN) in ISBN-10 format",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN10))

		for name, input := range invalidISBN10TestCases {
			t.Run(name, func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
	})
}

func BenchmarkStringISBN10(b *testing.B) {
	benchmarkStringPublicationRule(b, StringISBN10(), []string{
		"0-306-40615-2",
		"0-9752298-0-X",
		"978-0-306-40615-7",
	})
}

var validISBN13TestCases = map[string]string{
	"hyphenated": "978-0-306-40615-7",
	"plain":      "9780306406157",
	"grouped":    "978-3-16-148410-0",
	"979 prefix": "979-10-90636-07-1",
}

var invalidISBN13TestCases = map[string]string{
	"empty":          "",
	"isbn 10":        "0-306-40615-2",
	"failed check":   "978-0-306-40615-8",
	"x check":        "978030640615X",
	"invalid prefix": "9770306406157",
	"trailing space": "978 0 306 40615 7 ",
}

func TestStringISBN13(t *testing.T) {
	rule := StringISBN13()
	t.Run("valid inputs", func(t *testing.T) {
		for name, input := range validISBN13TestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		err := rule.Validate(invalidISBN13TestCases["empty"])
		assert.EqualError(
			t,
			err,
			"string must be a valid International Standard Book Number (ISBN) in ISBN-13 format",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN13))

		for name, input := range invalidISBN13TestCases {
			t.Run(name, func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
	})
}

func BenchmarkStringISBN13(b *testing.B) {
	benchmarkStringPublicationRule(b, StringISBN13(), []string{
		"978-0-306-40615-7",
		"979-10-90636-07-1",
		"0-306-40615-2",
	})
}

var validISSNTestCases = map[string]string{
	"numeric check":   "2049-3630",
	"numeric example": "0378-5955",
	"uppercase x":     "2434-561X",
	"lowercase x":     "2434-561x",
}

var invalidISSNTestCases = map[string]string{
	"empty":           "",
	"missing hyphen":  "20493630",
	"failed check":    "2049-3631",
	"wrong grouping":  "204-93630",
	"x before check":  "2049-36X0",
	"hyphen as check": "2049-363-",
}

func TestStringISSN(t *testing.T) {
	rule := StringISSN()
	t.Run("valid inputs", func(t *testing.T) {
		for name, input := range validISSNTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		err := rule.Validate(invalidISSNTestCases["empty"])
		assert.EqualError(t, err, "string must be a valid International Standard Serial Number (ISSN)")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISSN))

		for name, input := range invalidISSNTestCases {
			t.Run(name, func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
	})
}

func BenchmarkStringISSN(b *testing.B) {
	benchmarkStringPublicationRule(b, StringISSN(), []string{
		"2049-3630",
		"2434-561X",
		"2049-3631",
	})
}

func benchmarkStringPublicationRule(b *testing.B, rule govy.Rule[string], inputs []string) {
	b.Helper()
	for b.Loop() {
		for _, in := range inputs {
			_ = rule.Validate(in)
		}
	}
}
