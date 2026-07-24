package rules

import (
	"regexp"
	"strings"
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

func TestStringISBN_VeryLargeInvalid(t *testing.T) {
	input := strings.Repeat("0", 1<<20)
	tests := map[string]struct {
		rule    govy.Rule[string]
		message string
		code    govy.ErrorCode
	}{
		"isbn": {
			rule:    StringISBN(),
			message: "string must be a valid International Standard Book Number (ISBN) in ISBN-10 or ISBN-13 format",
			code:    ErrorCodeStringISBN,
		},
		"isbn-10": {
			rule:    StringISBN10(),
			message: "string must be a valid International Standard Book Number (ISBN) in ISBN-10 format",
			code:    ErrorCodeStringISBN10,
		},
		"isbn-13": {
			rule:    StringISBN13(),
			message: "string must be a valid International Standard Book Number (ISBN) in ISBN-13 format",
			code:    ErrorCodeStringISBN13,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.rule.Validate(input)
			assert.EqualError(t, err, test.message)
			assert.True(t, govy.HasErrorCode(err, test.code))
		})
	}
}

func BenchmarkStringISBN(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISBN(),
		validISBNTestCases,
		invalidISBNTestCases,
	)
}

func BenchmarkStringISBNVeryLargeInvalid(b *testing.B) {
	rule := StringISBN()
	input := strings.Repeat("0", 1<<20)

	for b.Loop() {
		_ = rule.Validate(input)
	}
	b.ReportMetric(1, "validations/op")
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
	benchmarkStringPublicationRule(
		b,
		StringISBN10(),
		validISBN10TestCases,
		invalidISBN10TestCases,
	)
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

func TestISBNPredicatesMatchReference(t *testing.T) {
	tests := map[string]struct {
		predicate func(string) bool
		reference func(string) bool
		inputs    []map[string]string
	}{
		"isbn": {
			predicate: isISBN,
			reference: referenceISBN,
			inputs: []map[string]string{
				validISBNTestCases,
				invalidISBNTestCases,
			},
		},
		"isbn-10": {
			predicate: isISBN10,
			reference: referenceISBN10,
			inputs: []map[string]string{
				validISBN10TestCases,
				invalidISBN10TestCases,
			},
		},
		"isbn-13": {
			predicate: isISBN13,
			reference: referenceISBN13,
			inputs: []map[string]string{
				validISBN13TestCases,
				invalidISBN13TestCases,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testStringPredicateMatchesReference(t, test.predicate, test.reference, test.inputs...)
		})
	}
}

func BenchmarkStringISBN13(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISBN13(),
		validISBN13TestCases,
		invalidISBN13TestCases,
	)
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

func TestISSNPredicateMatchesReference(t *testing.T) {
	format := regexp.MustCompile(`^\d{4}-\d{3}[0-9Xx]$`)
	testStringPredicateMatchesReference(
		t,
		isISSN,
		func(s string) bool {
			return referenceISSN(format, s)
		},
		validISSNTestCases,
		invalidISSNTestCases,
	)
}

func BenchmarkStringISSN(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISSN(),
		validISSNTestCases,
		invalidISSNTestCases,
	)
}

func Benchmark_isISSN(b *testing.B) {
	for b.Loop() {
		for name, input := range validISSNTestCases {
			if !isISSN(input) {
				b.Fatalf("%s: expected valid ISSN", name)
			}
		}
		for name, input := range invalidISSNTestCases {
			if isISSN(input) {
				b.Fatalf("%s: expected invalid ISSN", name)
			}
		}
	}
	b.ReportMetric(float64(len(validISSNTestCases)+len(invalidISSNTestCases)), "validations/op")
}

func benchmarkStringPublicationRule(
	b *testing.B,
	rule govy.Rule[string],
	validInputs map[string]string,
	invalidInputs map[string]string,
) {
	b.Helper()
	for b.Loop() {
		for _, in := range validInputs {
			_ = rule.Validate(in)
		}
		for _, in := range invalidInputs {
			_ = rule.Validate(in)
		}
	}
	b.ReportMetric(float64(len(validInputs)+len(invalidInputs)), "validations/op")
}

type stringPredicate func(string) bool

func testStringPredicateMatchesReference(
	t *testing.T,
	predicate stringPredicate,
	reference stringPredicate,
	inputs ...map[string]string,
) {
	t.Helper()
	for _, inputSet := range inputs {
		for name, input := range inputSet {
			assertStringPredicateMatchesReference(t, predicate, reference, name, input)
			testStringPredicateByteEdits(t, predicate, reference, input)
		}
	}
}

func testStringPredicateByteEdits(
	t *testing.T,
	predicate stringPredicate,
	reference stringPredicate,
	input string,
) {
	t.Helper()
	for position := range len(input) {
		candidate := input[:position] + input[position+1:]
		assertStringPredicateMatchesReference(t, predicate, reference, "byte deletion", candidate)
	}

	replacement := []byte(input)
	for position := range len(replacement) {
		original := replacement[position]
		for value := range 256 {
			replacement[position] = byte(value)
			assertStringPredicateMatchesReference(
				t,
				predicate,
				reference,
				"byte replacement",
				string(replacement),
			)
		}
		replacement[position] = original
	}

	for position := range len(input) + 1 {
		insertion := make([]byte, len(input)+1)
		copy(insertion, input[:position])
		copy(insertion[position+1:], input[position:])
		for value := range 256 {
			insertion[position] = byte(value)
			assertStringPredicateMatchesReference(
				t,
				predicate,
				reference,
				"byte insertion",
				string(insertion),
			)
		}
	}
}

func assertStringPredicateMatchesReference(
	t *testing.T,
	predicate stringPredicate,
	reference stringPredicate,
	name string,
	input string,
) {
	t.Helper()
	got := predicate(input)
	expected := reference(input)
	if got != expected {
		t.Fatalf("%s for %q: got %t, expected %t", name, input, got, expected)
	}
}

func referenceISBN(s string) bool {
	return referenceISBN10(s) || referenceISBN13(s)
}

func referenceISBN10(s string) bool {
	isbn, ok := referenceNormalizeISBN(s)
	if !ok || len(isbn) != isbn10Length {
		return false
	}

	sum := 0
	for i := range isbn10Length - 1 {
		if !isASCIIDigit(isbn[i]) {
			return false
		}
		sum += int(isbn[i]-'0') * (isbn10Length - i)
	}

	switch checkDigit := isbn[isbn10Length-1]; {
	case isASCIIDigit(checkDigit):
		sum += int(checkDigit - '0')
	case checkDigit == 'X' || checkDigit == 'x':
		sum += 10
	default:
		return false
	}
	return sum%11 == 0
}

func referenceISBN13(s string) bool {
	isbn, ok := referenceNormalizeISBN(s)
	if !ok || len(isbn) != isbn13Length ||
		(!strings.HasPrefix(isbn, "978") && !strings.HasPrefix(isbn, "979")) {
		return false
	}

	sum := 0
	for i := range isbn13Length - 1 {
		if !isASCIIDigit(isbn[i]) {
			return false
		}
		weight := 1
		if i%2 != 0 {
			weight = 3
		}
		sum += int(isbn[i]-'0') * weight
	}
	if !isASCIIDigit(isbn[isbn13Length-1]) {
		return false
	}
	return (10-sum%10)%10 == int(isbn[isbn13Length-1]-'0')
}

func referenceNormalizeISBN(s string) (string, bool) {
	if s == "" {
		return "", false
	}

	var builder strings.Builder
	previousWasSeparator := false
	for i := range len(s) {
		switch c := s[i]; {
		case isASCIIDigit(c) || c == 'X' || c == 'x':
			builder.WriteByte(c)
			previousWasSeparator = false
		case c == '-' || c == ' ':
			if i == 0 || i == len(s)-1 || previousWasSeparator {
				return "", false
			}
			previousWasSeparator = true
		default:
			return "", false
		}
	}
	return builder.String(), true
}

func referenceISSN(format *regexp.Regexp, s string) bool {
	if !format.MatchString(s) {
		return false
	}

	issn := strings.ReplaceAll(s, "-", "")
	sum := 0
	for i := range 7 {
		if !isASCIIDigit(issn[i]) {
			return false
		}
		sum += int(issn[i]-'0') * (8 - i)
	}

	switch checkDigit := issn[7]; {
	case isASCIIDigit(checkDigit):
		sum += int(checkDigit - '0')
	case checkDigit == 'X' || checkDigit == 'x':
		sum += 10
	default:
		return false
	}
	return sum%11 == 0
}
