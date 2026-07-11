package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestStringISBN(t *testing.T) {
	runStringPublicationRuleTest(t, StringISBN(), ErrorCodeStringISBN, map[string]stringPublicationRuleTestCase{
		"isbn 10 hyphenated": {in: "0-306-40615-2"},
		"isbn 10 plain":      {in: "0306406152"},
		"isbn 10 x check":    {in: "0-9752298-0-X"},
		"isbn 10 spaced":     {in: "0 9752298 0 x"},
		"isbn 13 hyphenated": {in: "978-0-306-40615-7"},
		"isbn 13 plain":      {in: "9780306406157"},
		"isbn 13 grouped":    {in: "978-3-16-148410-0"},
		"empty": {
			in:            "",
			expectedError: "string must be a valid ISBN",
		},
		"isbn 10 failed check": {
			in:            "0-306-40615-3",
			expectedError: "string must be a valid ISBN",
		},
		"isbn 13 failed check": {
			in:            "978-0-306-40615-8",
			expectedError: "string must be a valid ISBN",
		},
		"isbn 13 x check": {
			in:            "978030640615X",
			expectedError: "string must be a valid ISBN",
		},
		"repeated separator": {
			in:            "978--0-306-40615-7",
			expectedError: "string must be a valid ISBN",
		},
		"letters": {
			in:            "abc",
			expectedError: "string must be a valid ISBN",
		},
	})
}

func TestStringISBN10(t *testing.T) {
	runStringPublicationRuleTest(t, StringISBN10(), ErrorCodeStringISBN10, map[string]stringPublicationRuleTestCase{
		"hyphenated": {in: "0-306-40615-2"},
		"plain":      {in: "0306406152"},
		"x check":    {in: "0-9752298-0-X"},
		"spaced":     {in: "0 9752298 0 x"},
		"empty": {
			in:            "",
			expectedError: "string must be a valid ISBN-10",
		},
		"failed check": {
			in:            "0-306-40615-3",
			expectedError: "string must be a valid ISBN-10",
		},
		"isbn 13": {
			in:            "978-0-306-40615-7",
			expectedError: "string must be a valid ISBN-10",
		},
		"isbn 13 plain": {
			in:            "9780306406157",
			expectedError: "string must be a valid ISBN-10",
		},
		"repeated separator": {
			in:            "0-306--40615-2",
			expectedError: "string must be a valid ISBN-10",
		},
	})
}

func TestStringISBN13(t *testing.T) {
	runStringPublicationRuleTest(t, StringISBN13(), ErrorCodeStringISBN13, map[string]stringPublicationRuleTestCase{
		"hyphenated":     {in: "978-0-306-40615-7"},
		"plain":          {in: "9780306406157"},
		"grouped":        {in: "978-3-16-148410-0"},
		"979 prefix":     {in: "979-10-90636-07-1"},
		"empty":          {in: "", expectedError: "string must be a valid ISBN-13"},
		"isbn 10":        {in: "0-306-40615-2", expectedError: "string must be a valid ISBN-13"},
		"failed check":   {in: "978-0-306-40615-8", expectedError: "string must be a valid ISBN-13"},
		"x check":        {in: "978030640615X", expectedError: "string must be a valid ISBN-13"},
		"invalid prefix": {in: "9770306406157", expectedError: "string must be a valid ISBN-13"},
		"trailing space": {in: "978 0 306 40615 7 ", expectedError: "string must be a valid ISBN-13"},
	})
}

func TestStringISSN(t *testing.T) {
	runStringPublicationRuleTest(t, StringISSN(), ErrorCodeStringISSN, map[string]stringPublicationRuleTestCase{
		"numeric check":   {in: "2049-3630"},
		"numeric example": {in: "0378-5955"},
		"uppercase x":     {in: "2434-561X"},
		"lowercase x":     {in: "2434-561x"},
		"empty":           {in: "", expectedError: "string must be a valid ISSN"},
		"missing hyphen":  {in: "20493630", expectedError: "string must be a valid ISSN"},
		"failed check":    {in: "2049-3631", expectedError: "string must be a valid ISSN"},
		"wrong grouping":  {in: "204-93630", expectedError: "string must be a valid ISSN"},
		"x before check":  {in: "2049-36X0", expectedError: "string must be a valid ISSN"},
		"hyphen as check": {in: "2049-363-", expectedError: "string must be a valid ISSN"},
	})
}

func BenchmarkStringISBN(b *testing.B) {
	benchmarkStringPublicationRule(b, StringISBN(), []string{
		"0-306-40615-2",
		"978-0-306-40615-7",
		"978--0-306-40615-7",
	})
}

func BenchmarkStringISBN10(b *testing.B) {
	benchmarkStringPublicationRule(b, StringISBN10(), []string{
		"0-306-40615-2",
		"0-9752298-0-X",
		"978-0-306-40615-7",
	})
}

func BenchmarkStringISBN13(b *testing.B) {
	benchmarkStringPublicationRule(b, StringISBN13(), []string{
		"978-0-306-40615-7",
		"979-10-90636-07-1",
		"0-306-40615-2",
	})
}

func BenchmarkStringISSN(b *testing.B) {
	benchmarkStringPublicationRule(b, StringISSN(), []string{
		"2049-3630",
		"2434-561X",
		"2049-3631",
	})
}

type stringPublicationRuleTestCase struct {
	in            string
	expectedError string
}

func runStringPublicationRuleTest(
	t *testing.T,
	rule govy.Rule[string],
	errorCode govy.ErrorCode,
	testCases map[string]stringPublicationRuleTestCase,
) {
	t.Helper()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := rule.Validate(tc.in)
			if tc.expectedError != "" {
				assert.EqualError(t, err, tc.expectedError)
				assert.True(t, govy.HasErrorCode(err, errorCode))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func benchmarkStringPublicationRule(b *testing.B, rule govy.Rule[string], inputs []string) {
	b.Helper()
	for b.Loop() {
		for _, in := range inputs {
			_ = rule.Validate(in)
		}
	}
}
