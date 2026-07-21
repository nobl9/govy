package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

const (
	stringEINErrorMessage = "string must be a valid Employer Identification Number (EIN)"
	stringSSNErrorMessage = "string must be a valid Social Security Number (SSN)"
)

func TestStringEIN(t *testing.T) {
	t.Parallel()

	rule := StringEIN()
	// This is the complete set of distinct prefixes published by the IRS
	// "Valid EINs" page, updated 2026-04-09.
	// https://www.irs.gov/businesses/small-businesses-self-employed/valid-eins
	recognizedPrefixes := map[string]struct{}{
		"01": {}, "02": {}, "03": {}, "04": {}, "05": {}, "06": {},
		"10": {}, "11": {}, "12": {}, "13": {}, "14": {}, "15": {}, "16": {},
		"20": {}, "21": {}, "22": {}, "23": {}, "24": {}, "25": {}, "26": {}, "27": {},
		"30": {}, "31": {}, "32": {}, "33": {}, "34": {}, "35": {}, "36": {}, "37": {}, "38": {}, "39": {},
		"40": {}, "41": {}, "42": {}, "43": {}, "44": {}, "45": {}, "46": {}, "47": {}, "48": {},
		"50": {}, "51": {}, "52": {}, "53": {}, "54": {}, "55": {}, "56": {}, "57": {}, "58": {}, "59": {},
		"60": {}, "61": {}, "62": {}, "63": {}, "64": {}, "65": {}, "66": {}, "67": {}, "68": {},
		"71": {}, "72": {}, "73": {}, "74": {}, "75": {}, "76": {}, "77": {},
		"80": {}, "81": {}, "82": {}, "83": {}, "84": {}, "85": {}, "86": {}, "87": {}, "88": {},
		"90": {}, "91": {}, "92": {}, "93": {}, "94": {}, "95": {}, "98": {}, "99": {},
	}
	assert.Require(t, assert.Len(t, recognizedPrefixes, 83))

	t.Run("IRS prefix corpus", func(t *testing.T) {
		t.Parallel()

		recognizedCount := 0
		unrecognizedCount := 0
		for prefixNumber := range 100 {
			prefix := fmt.Sprintf("%02d", prefixNumber)
			_, shouldPass := recognizedPrefixes[prefix]
			if shouldPass {
				recognizedCount++
			} else {
				unrecognizedCount++
			}
			t.Run(prefix, func(t *testing.T) {
				assertTaxIDRuleValidity(
					t,
					rule,
					prefix+"-3456789",
					shouldPass,
					stringEINErrorMessage,
					ErrorCodeStringEIN,
				)
			})
		}
		assert.Equal(t, 83, recognizedCount)
		assert.Equal(t, 17, unrecognizedCount)
	})

	accepted := map[string]string{
		"lowest recognized prefix":  "01-0000001",
		"highest recognized prefix": "99-9999999",
	}
	rejected := map[string]string{
		"empty input":          "",
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

				assertTaxIDRuleValidity(
					t,
					rule,
					in,
					true,
					stringEINErrorMessage,
					ErrorCodeStringEIN,
				)
			})
		}
	})
	t.Run("rejected structures", func(t *testing.T) {
		t.Parallel()

		for name, in := range rejected {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				assertTaxIDRuleValidity(
					t,
					rule,
					in,
					false,
					stringEINErrorMessage,
					ErrorCodeStringEIN,
				)
			})
		}
	})
}

func TestStringSSN(t *testing.T) {
	t.Parallel()

	rule := StringSSN()
	// IRS IRM 3.13.5.21 (2022-01-01) lists these never-issued examples.
	// https://www.irs.gov/irm/part3/irm_03-013-005
	accepted := map[string]string{
		"lowest structural fields":  "001-01-0001",
		"area below 666":            "665-99-9999",
		"area above 666":            "667-01-0001",
		"area 772":                  "772-01-0001",
		"area 800":                  "800-01-0001",
		"highest structural fields": "899-99-9999",
		"IRS never-issued example 111 is structurally valid": "111-11-1111",
		"IRS never-issued example 222 is structurally valid": "222-22-2222",
		"IRS never-issued example 777 is structurally valid": "777-77-7777",
		"IRS never-issued example 123 is structurally valid": "123-45-6789",
	}
	rejected := map[string]string{
		"empty input":               "",
		"zero area":                 "000-01-0001",
		"IRS never-issued area 666": "666-66-6666",
		"area 900":                  "900-01-0001",
		"area 901":                  "901-01-0001",
		"area 998":                  "998-01-0001",
		"area 999":                  "999-01-0001",
		"zero group":                "001-00-0001",
		"zero serial":               "001-01-0000",
		"short area":                "01-01-0001",
		"long area":                 "0001-01-0001",
		"short group":               "001-1-0001",
		"long group":                "001-001-0001",
		"short serial":              "001-01-001",
		"long serial":               "001-01-00001",
		"letter area":               "00A-01-0001",
		"letter group":              "001-0A-0001",
		"letter serial":             "001-01-000A",
		"missing separators":        "001010001",
		"slash separators":          "001/01/0001",
		"en dash separators":        "001–01–0001",
		"space separators":          "001 01 0001",
		"leading whitespace":        " 001-01-0001",
		"trailing whitespace":       "001-01-0001 ",
		"full-width digits":         "００１-０１-０００１",
		"trailing newline":          "001-01-0001\n",
		"EIN-shaped input":          "12-3456789",
	}

	t.Run("accepted structures", func(t *testing.T) {
		t.Parallel()

		for name, in := range accepted {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				assertTaxIDRuleValidity(
					t,
					rule,
					in,
					true,
					stringSSNErrorMessage,
					ErrorCodeStringSSN,
				)
			})
		}
	})
	t.Run("rejected structures", func(t *testing.T) {
		t.Parallel()

		for name, in := range rejected {
			t.Run(name, func(t *testing.T) {
				t.Parallel()

				assertTaxIDRuleValidity(
					t,
					rule,
					in,
					false,
					stringSSNErrorMessage,
					ErrorCodeStringSSN,
				)
			})
		}
	})
}

func TestStringSSN_StructuralFields(t *testing.T) {
	t.Parallel()

	rule := StringSSN()
	// SSA POMS RM 10201.035 (2011-06-23) defines these structural exclusions.
	// https://secure.ssa.gov/poms.nsf/lnx/0110201035
	t.Run("areas 000 through 999", func(t *testing.T) {
		t.Parallel()

		validCount := 0
		invalidCount := 0
		for area := range 1_000 {
			ssn := fmt.Sprintf("%03d-01-0001", area)
			shouldPass := area != 0 && area != 666 && area < 900
			if shouldPass {
				validCount++
			} else {
				invalidCount++
			}
			t.Run(ssn, func(t *testing.T) {
				assertTaxIDRuleValidity(
					t,
					rule,
					ssn,
					shouldPass,
					stringSSNErrorMessage,
					ErrorCodeStringSSN,
				)
			})
		}
		assert.Equal(t, 898, validCount)
		assert.Equal(t, 102, invalidCount)
	})
	t.Run("groups 00 through 99", func(t *testing.T) {
		t.Parallel()

		validCount := 0
		invalidCount := 0
		for group := range 100 {
			ssn := fmt.Sprintf("001-%02d-0001", group)
			shouldPass := group != 0
			if shouldPass {
				validCount++
			} else {
				invalidCount++
			}
			t.Run(ssn, func(t *testing.T) {
				assertTaxIDRuleValidity(
					t,
					rule,
					ssn,
					shouldPass,
					stringSSNErrorMessage,
					ErrorCodeStringSSN,
				)
			})
		}
		assert.Equal(t, 99, validCount)
		assert.Equal(t, 1, invalidCount)
	})
	t.Run("serials 0000 through 9999", func(t *testing.T) {
		t.Parallel()

		validCount := 0
		invalidCount := 0
		for serial := range 10_000 {
			ssn := fmt.Sprintf("001-01-%04d", serial)
			shouldPass := serial != 0
			if shouldPass {
				validCount++
			} else {
				invalidCount++
			}
			t.Run(ssn, func(t *testing.T) {
				assertTaxIDRuleValidity(
					t,
					rule,
					ssn,
					shouldPass,
					stringSSNErrorMessage,
					ErrorCodeStringSSN,
				)
			})
		}
		assert.Equal(t, 9_999, validCount)
		assert.Equal(t, 1, invalidCount)
	})
}

func BenchmarkStringEIN(b *testing.B) {
	rule := StringEIN()
	for b.Loop() {
		_ = rule.Validate("12-3456789")
	}
}

func BenchmarkStringSSN(b *testing.B) {
	rule := StringSSN()
	for b.Loop() {
		_ = rule.Validate("123-45-6789")
	}
}

func assertTaxIDRuleValidity(
	t *testing.T,
	rule govy.Rule[string],
	in string,
	shouldPass bool,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	err := rule.Validate(in)
	if shouldPass {
		assert.NoError(t, err)
		return
	}
	assert.EqualError(t, err, expectedError)
	assert.True(t, govy.HasErrorCode(err, errorCode))
}
