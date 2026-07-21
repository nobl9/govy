package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestStringEIN(t *testing.T) {
	t.Parallel()

	rule := StringEIN()
	accepted := map[string]string{
		"lowest recognized prefix":  "01-0000001",
		"representative prefix":     "12-3456789",
		"highest recognized prefix": "99-9999999",
	}
	rejected := map[string]string{
		"empty input":          "",
		"zero prefix":          "00-0000000",
		"unrecognized prefix":  "07-3456789",
		"missing separator":    "123456789",
		"one-digit prefix":     "1-3456789",
		"three-digit prefix":   "012-3456789",
		"short serial":         "12-345678",
		"long serial":          "12-34567890",
		"letter prefix":        "AB-3456789",
		"letter serial":        "12-345678A",
		"space separator":      "12 3456789",
		"underscore separator": "12_3456789",
		"en dash separator":    "12–3456789",
		"double separator":     "12--3456789",
		"leading whitespace":   " 12-3456789",
		"trailing whitespace":  "12-3456789 ",
		"full-width digits":    "１２-３４５６７８９",
		"trailing newline":     "12-3456789\n",
	}

	t.Run("accepted structures", func(t *testing.T) {
		t.Parallel()

		for name, in := range accepted {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				assert.NoError(t, rule.Validate(in))
			})
		}
	})
	t.Run("rejected structures", func(t *testing.T) {
		t.Parallel()

		for name, in := range rejected {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				err := rule.Validate(in)
				assert.EqualError(t, err, "string must be a valid Employer Identification Number (EIN)")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEIN))
			})
		}
	})
}

func Test_isValidEINPrefix(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		prefixes []string
		expected bool
	}{
		"recognized prefixes": {
			prefixes: []string{
				"01", "02", "03", "04", "05", "06",
				"10", "11", "12", "13", "14", "15", "16",
				"20", "21", "22", "23", "24", "25", "26", "27",
				"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
				"40", "41", "42", "43", "44", "45", "46", "47", "48",
				"50", "51", "52", "53", "54", "55", "56", "57", "58", "59",
				"60", "61", "62", "63", "64", "65", "66", "67", "68",
				"71", "72", "73", "74", "75", "76", "77",
				"80", "81", "82", "83", "84", "85", "86", "87", "88",
				"90", "91", "92", "93", "94", "95", "98", "99",
			},
			expected: true,
		},
		"unrecognized prefixes": {
			prefixes: []string{
				"00", "07", "08", "09", "17", "18", "19", "28", "29",
				"49", "69", "70", "78", "79", "89", "96", "97",
			},
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			for _, prefix := range test.prefixes {
				t.Run(prefix, func(t *testing.T) {
					t.Parallel()

					assert.Equal(t, test.expected, isValidEINPrefix(prefix))
				})
			}
		})
	}
}

func BenchmarkStringEIN(b *testing.B) {
	rule := StringEIN()
	for b.Loop() {
		_ = rule.Validate("12-3456789")
	}
}

func TestStringSSN(t *testing.T) {
	t.Parallel()

	rule := StringSSN()
	accepted := map[string]string{
		"lowest structural fields":  "001-01-0001",
		"area below 666":            "665-99-9999",
		"area above 666":            "667-01-0001",
		"area 772":                  "772-01-0001",
		"area 800":                  "800-01-0001",
		"highest structural fields": "899-99-9999",
		"representative structure":  "123-45-6789",
	}
	rejected := map[string]string{
		"empty input":         "",
		"zero area":           "000-01-0001",
		"area 666":            "666-01-0001",
		"area 900":            "900-01-0001",
		"area 901":            "901-01-0001",
		"area 998":            "998-01-0001",
		"area 999":            "999-01-0001",
		"zero group":          "001-00-0001",
		"zero serial":         "001-01-0000",
		"short area":          "01-01-0001",
		"long area":           "0001-01-0001",
		"short group":         "001-1-0001",
		"long group":          "001-001-0001",
		"short serial":        "001-01-001",
		"long serial":         "001-01-00001",
		"letter area":         "00A-01-0001",
		"letter group":        "001-0A-0001",
		"letter serial":       "001-01-000A",
		"missing separators":  "001010001",
		"slash separators":    "001/01/0001",
		"en dash separators":  "001–01–0001",
		"space separators":    "001 01 0001",
		"leading whitespace":  " 001-01-0001",
		"trailing whitespace": "001-01-0001 ",
		"full-width digits":   "００１-０１-０００１",
		"trailing newline":    "001-01-0001\n",
		"EIN-shaped input":    "12-3456789",
	}

	t.Run("accepted structures", func(t *testing.T) {
		t.Parallel()

		for name, in := range accepted {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				assert.NoError(t, rule.Validate(in))
			})
		}
	})
	t.Run("rejected structures", func(t *testing.T) {
		t.Parallel()

		for name, in := range rejected {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				err := rule.Validate(in)
				assert.EqualError(t, err, "string must be a valid Social Security Number (SSN)")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSSN))
			})
		}
	})
}

func Test_isValidSSN_StructuralFields(t *testing.T) {
	t.Parallel()

	t.Run("areas 000 through 999", func(t *testing.T) {
		t.Parallel()

		for area := range 1_000 {
			ssn := fmt.Sprintf("%03d-01-0001", area)
			expected := area != 0 && area != 666 && area < 900
			if actual := isValidSSN(ssn); actual != expected {
				t.Errorf("%q: expected validity %t, actual %t", ssn, expected, actual)
			}
		}
	})
	t.Run("groups 00 through 99", func(t *testing.T) {
		t.Parallel()

		for group := range 100 {
			ssn := fmt.Sprintf("001-%02d-0001", group)
			expected := group != 0
			if actual := isValidSSN(ssn); actual != expected {
				t.Errorf("%q: expected validity %t, actual %t", ssn, expected, actual)
			}
		}
	})
	t.Run("serials 0000 through 9999", func(t *testing.T) {
		t.Parallel()

		for serial := range 10_000 {
			ssn := fmt.Sprintf("001-01-%04d", serial)
			expected := serial != 0
			if actual := isValidSSN(ssn); actual != expected {
				t.Errorf("%q: expected validity %t, actual %t", ssn, expected, actual)
			}
		}
	})
}

func BenchmarkStringSSN(b *testing.B) {
	rule := StringSSN()
	for b.Loop() {
		_ = rule.Validate("123-45-6789")
	}
}
