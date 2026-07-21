package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

var validISBNTestCases = map[string]string{
	"isbn 10 hyphenated":                  "0-306-40615-2",
	"isbn 10 plain":                       "0306406152",
	"isbn 10 x check":                     "0-9752298-0-X",
	"isbn 10 spaced":                      "0 9752298 0 x",
	"isbn 10 library converter numeric":   "0394170660",
	"isbn 10 library converter alternate": "0717941728",
	"isbn 10 library converter x check":   "087779443X",
	"isbn 10 MARC hyphenated":             "0-87068-693-3",
	"isbn 13 hyphenated":                  "978-0-306-40615-7",
	"isbn 13 plain":                       "9780306406157",
	"isbn 13 grouped":                     "978-3-16-148410-0",
	"isbn 13 agency manual hyphenated":    "978-92-95055-12-4",
	"isbn 13 agency manual spaced":        "978 92 95055 12 4",
	"isbn 13 agency manual compact":       "9789295055124",
	"isbn 13 agency manual hardback":      "978-951-45-9693-3",
	"isbn 13 agency manual paperback":     "978-951-45-9694-0",
	"isbn 13 agency manual PDF":           "978-951-45-9695-7",
	"isbn 13 agency manual EPUB":          "978-951-45-9696-4",
	"isbn 13 library converter first":     "9780060723804",
	"isbn 13 library converter second":    "9780060799748",
	"isbn 13 979 prefix":                  "979-10-90636-07-1",
}

var invalidISBNTestCases = map[string]string{
	"empty":                               "",
	"isbn 10 failed check":                "0-306-40615-3",
	"isbn 10 x check mutation":            "0877794430",
	"isbn 10 x in body":                   "08777X443X",
	"isbn 10 x in fourth position":        "087X79443X",
	"isbn 10 short":                       "087779443",
	"isbn 10 trailing space":              "087779443X ",
	"isbn 13 failed check":                "978-0-306-40615-8",
	"isbn 13 manual check mutation":       "978-92-95055-12-5",
	"isbn 13 checksum valid wrong prefix": "9779295055125",
	"isbn 13 x in body":                   "978-92-X5055-12-4",
	"isbn 13 x check":                     "978-92-95055-12-X",
	"isbn 13 en dash separators":          "978–92–95055–12–4",
	"isbn 13 full width digits":           "９７８９２９５０５５１２４",
	"isbn 13 display prefix":              "ISBN 978-92-95055-12-4",
	"repeated separator":                  "978--0-306-40615-7",
	"letters":                             "abc",
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
		for name, input := range invalidISBNTestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(
					t,
					err,
					"string must be a valid International Standard Book Number (ISBN) in ISBN-10 or ISBN-13 format",
				)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN))
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
	"hyphenated":                  "0-306-40615-2",
	"plain":                       "0306406152",
	"x check":                     "0-9752298-0-X",
	"spaced":                      "0 9752298 0 x",
	"library converter numeric":   "0394170660",
	"library converter alternate": "0717941728",
	"library converter x check":   "087779443X",
	"MARC hyphenated":             "0-87068-693-3",
}

var invalidISBN10TestCases = map[string]string{
	"empty":                "",
	"failed check":         "0-306-40615-3",
	"x check mutation":     "0877794430",
	"x in body":            "08777X443X",
	"x in fourth position": "087X79443X",
	"short":                "087779443",
	"trailing space":       "087779443X ",
	"isbn 13":              "978-0-306-40615-7",
	"isbn 13 plain":        "9780306406157",
	"repeated separator":   "0-306--40615-2",
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
		for name, input := range invalidISBN10TestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(
					t,
					err,
					"string must be a valid International Standard Book Number (ISBN) in ISBN-10 format",
				)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN10))
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
	"hyphenated":               "978-0-306-40615-7",
	"plain":                    "9780306406157",
	"grouped":                  "978-3-16-148410-0",
	"agency manual hyphenated": "978-92-95055-12-4",
	"agency manual spaced":     "978 92 95055 12 4",
	"agency manual compact":    "9789295055124",
	"agency manual hardback":   "978-951-45-9693-3",
	"agency manual paperback":  "978-951-45-9694-0",
	"agency manual PDF":        "978-951-45-9695-7",
	"agency manual EPUB":       "978-951-45-9696-4",
	"library converter first":  "9780060723804",
	"library converter second": "9780060799748",
	"979 prefix":               "979-10-90636-07-1",
}

var invalidISBN13TestCases = map[string]string{
	"empty":                       "",
	"isbn 10":                     "0-306-40615-2",
	"failed check":                "978-0-306-40615-8",
	"manual check mutation":       "978-92-95055-12-5",
	"checksum valid wrong prefix": "9779295055125",
	"x in body":                   "978-92-X5055-12-4",
	"x check":                     "978-92-95055-12-X",
	"en dash separators":          "978–92–95055–12–4",
	"full width digits":           "９７８９２９５０５５１２４",
	"display prefix":              "ISBN 978-92-95055-12-4",
	"invalid prefix":              "9770306406157",
	"trailing space":              "978 0 306 40615 7 ",
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
		for name, input := range invalidISBN13TestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(
					t,
					err,
					"string must be a valid International Standard Book Number (ISBN) in ISBN-13 format",
				)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN13))
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
	"numeric check":       "2049-3630",
	"numeric example":     "0378-5955",
	"uppercase x":         "2434-561X",
	"lowercase x":         "2434-561x",
	"manual numeric":      "1106-1111",
	"manual uppercase x":  "1092-003X",
	"library check digit": "0317-8471",
	"numeric 2162":        "2162-3546",
	"numeric 1548":        "1548-7180",
	"uppercase x 1204":    "1204-539X",
}

// invalidISSNTestCases includes exact compact construction examples from the
// [ISSN Manual, May 2025] because StringISSN requires the ASCII-hyphenated form.
// Their derived hyphenated forms are accepted in validISSNTestCases.
//
// [ISSN Manual, May 2025]: https://www.issn.org/wp-content/uploads/2025/05/Manual-ISSN_ENG-marc21_May2025.pdf
var invalidISSNTestCases = map[string]string{
	"empty":                      "",
	"missing hyphen":             "20493630",
	"manual compact 2162":        "21623546",
	"manual compact 1548":        "15487180",
	"failed check":               "2049-3631",
	"numeric check mutation":     "1106-1112",
	"uppercase x check mutation": "1092-0030",
	"wrong grouping":             "204-93630",
	"x before check":             "2049-36X0",
	"hyphen as check":            "2049-363-",
	"unicode hyphen":             "1106–1111",
	"U+2010 hyphen":              "1092‐003X",
	"space separator":            "1106 1111",
	"display prefix":             "ISSN 1106-1111",
	"trailing newline":           "1106-1111\n",
	"full width digits":          "１１０６-１１１１",
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
		for name, input := range invalidISSNTestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(t, err, "string must be a valid International Standard Serial Number (ISSN)")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISSN))
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
