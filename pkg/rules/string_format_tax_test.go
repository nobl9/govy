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

type stringTaxIDTestCase struct {
	name       string
	in         string
	shouldPass bool
}

// stringEINRecognizedPrefixes is the complete set of distinct prefixes
// published by the IRS "Valid EINs" page, updated 2026-04-09.
// https://www.irs.gov/businesses/small-businesses-self-employed/valid-eins
var stringEINRecognizedPrefixes = map[string]struct{}{
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

var (
	stringEINPrefixTestCases            = generateStringEINPrefixTestCases()
	stringEINAcceptedStructureTestCases = []stringTaxIDTestCase{
		{name: "lowest recognized prefix", in: "01-0000001", shouldPass: true},
		{name: "highest recognized prefix", in: "99-9999999", shouldPass: true},
	}
	stringEINRejectedStructureTestCases = []stringTaxIDTestCase{
		{name: "empty input", in: ""},
		{name: "missing separator", in: "123456789"},
		{name: "one-digit prefix", in: "1-3456789"},
		{name: "three-digit prefix", in: "012-3456789"},
		{name: "short serial", in: "12-345678"},
		{name: "long serial", in: "12-34567890"},
		{name: "letter prefix", in: "AB-3456789"},
		{name: "letter serial", in: "12-345678A"},
		{name: "space separator", in: "12 3456789"},
		{name: "underscore separator", in: "12_3456789"},
		{name: "en dash separator", in: "12–3456789"},
		{name: "double separator", in: "12--3456789"},
		{name: "leading whitespace", in: " 12-3456789"},
		{name: "trailing whitespace", in: "12-3456789 "},
		{name: "full-width digits", in: "１２-３４５６７８９"},
		{name: "trailing newline", in: "12-3456789\n"},
	}
)

// IRS IRM 3.13.5.21 (2022-01-01) lists the never-issued examples below.
// https://www.irs.gov/irm/part3/irm_03-013-005
var (
	stringSSNAcceptedStructureTestCases = []stringTaxIDTestCase{
		{name: "lowest structural fields", in: "001-01-0001", shouldPass: true},
		{name: "area below 666", in: "665-99-9999", shouldPass: true},
		{name: "area above 666", in: "667-01-0001", shouldPass: true},
		{name: "area 772", in: "772-01-0001", shouldPass: true},
		{name: "area 800", in: "800-01-0001", shouldPass: true},
		{name: "highest structural fields", in: "899-99-9999", shouldPass: true},
		{name: "IRS never-issued example 111 is structurally valid", in: "111-11-1111", shouldPass: true},
		{name: "IRS never-issued example 222 is structurally valid", in: "222-22-2222", shouldPass: true},
		{name: "IRS never-issued example 777 is structurally valid", in: "777-77-7777", shouldPass: true},
		{name: "IRS never-issued example 123 is structurally valid", in: "123-45-6789", shouldPass: true},
	}
	stringSSNRejectedStructureTestCases = []stringTaxIDTestCase{
		{name: "empty input", in: ""},
		{name: "zero area", in: "000-01-0001"},
		{name: "IRS never-issued area 666", in: "666-66-6666"},
		{name: "area 900", in: "900-01-0001"},
		{name: "area 901", in: "901-01-0001"},
		{name: "area 998", in: "998-01-0001"},
		{name: "area 999", in: "999-01-0001"},
		{name: "zero group", in: "001-00-0001"},
		{name: "zero serial", in: "001-01-0000"},
		{name: "short area", in: "01-01-0001"},
		{name: "long area", in: "0001-01-0001"},
		{name: "short group", in: "001-1-0001"},
		{name: "long group", in: "001-001-0001"},
		{name: "short serial", in: "001-01-001"},
		{name: "long serial", in: "001-01-00001"},
		{name: "letter area", in: "00A-01-0001"},
		{name: "letter group", in: "001-0A-0001"},
		{name: "letter serial", in: "001-01-000A"},
		{name: "missing separators", in: "001010001"},
		{name: "slash separators", in: "001/01/0001"},
		{name: "en dash separators", in: "001–01–0001"},
		{name: "space separators", in: "001 01 0001"},
		{name: "leading whitespace", in: " 001-01-0001"},
		{name: "trailing whitespace", in: "001-01-0001 "},
		{name: "full-width digits", in: "００１-０１-０００１"},
		{name: "trailing newline", in: "001-01-0001\n"},
		{name: "EIN-shaped input", in: "12-3456789"},
	}
	stringSSNAreaTestCases   = generateStringSSNAreaTestCases()
	stringSSNGroupTestCases  = generateStringSSNGroupTestCases()
	stringSSNSerialTestCases = generateStringSSNSerialTestCases()
)

func TestStringEIN(t *testing.T) {
	t.Parallel()

	rule := StringEIN()
	assert.Require(t, assert.Len(t, stringEINRecognizedPrefixes, 83))

	t.Run("IRS prefix corpus", func(t *testing.T) {
		t.Parallel()

		recognizedCount := 0
		unrecognizedCount := 0
		for _, tc := range stringEINPrefixTestCases {
			if tc.shouldPass {
				recognizedCount++
			} else {
				unrecognizedCount++
			}
			t.Run(tc.name, func(t *testing.T) {
				assertTaxIDRuleValidity(
					t,
					rule,
					tc.in,
					tc.shouldPass,
					stringEINErrorMessage,
					ErrorCodeStringEIN,
				)
			})
		}
		assert.Equal(t, 83, recognizedCount)
		assert.Equal(t, 17, unrecognizedCount)
	})

	t.Run("accepted structures", func(t *testing.T) {
		t.Parallel()

		for _, tc := range stringEINAcceptedStructureTestCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				assertTaxIDRuleValidity(
					t,
					rule,
					tc.in,
					tc.shouldPass,
					stringEINErrorMessage,
					ErrorCodeStringEIN,
				)
			})
		}
	})
	t.Run("rejected structures", func(t *testing.T) {
		t.Parallel()

		for _, tc := range stringEINRejectedStructureTestCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				assertTaxIDRuleValidity(
					t,
					rule,
					tc.in,
					tc.shouldPass,
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
	t.Run("accepted structures", func(t *testing.T) {
		t.Parallel()

		for _, tc := range stringSSNAcceptedStructureTestCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				assertTaxIDRuleValidity(
					t,
					rule,
					tc.in,
					tc.shouldPass,
					stringSSNErrorMessage,
					ErrorCodeStringSSN,
				)
			})
		}
	})
	t.Run("rejected structures", func(t *testing.T) {
		t.Parallel()

		for _, tc := range stringSSNRejectedStructureTestCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				assertTaxIDRuleValidity(
					t,
					rule,
					tc.in,
					tc.shouldPass,
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
		for _, tc := range stringSSNAreaTestCases {
			if tc.shouldPass {
				validCount++
			} else {
				invalidCount++
			}
			t.Run(tc.name, func(t *testing.T) {
				assertTaxIDRuleValidity(
					t,
					rule,
					tc.in,
					tc.shouldPass,
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
		for _, tc := range stringSSNGroupTestCases {
			if tc.shouldPass {
				validCount++
			} else {
				invalidCount++
			}
			t.Run(tc.name, func(t *testing.T) {
				assertTaxIDRuleValidity(
					t,
					rule,
					tc.in,
					tc.shouldPass,
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
		for _, tc := range stringSSNSerialTestCases {
			if tc.shouldPass {
				validCount++
			} else {
				invalidCount++
			}
			t.Run(tc.name, func(t *testing.T) {
				assertTaxIDRuleValidity(
					t,
					rule,
					tc.in,
					tc.shouldPass,
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
	benchmarkStringTaxIDRule(
		b,
		rule,
		stringEINPrefixTestCases,
		stringEINAcceptedStructureTestCases,
		stringEINRejectedStructureTestCases,
	)
}

func BenchmarkStringSSN(b *testing.B) {
	rule := StringSSN()
	benchmarkStringTaxIDRule(
		b,
		rule,
		stringSSNAcceptedStructureTestCases,
		stringSSNRejectedStructureTestCases,
		stringSSNAreaTestCases,
		stringSSNGroupTestCases,
		stringSSNSerialTestCases,
	)
}

func benchmarkStringTaxIDRule(
	b *testing.B,
	rule govy.Rule[string],
	testCaseGroups ...[]stringTaxIDTestCase,
) {
	b.Helper()
	for b.Loop() {
		for _, testCases := range testCaseGroups {
			for _, tc := range testCases {
				_ = rule.Validate(tc.in)
			}
		}
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

func generateStringEINPrefixTestCases() []stringTaxIDTestCase {
	testCases := make([]stringTaxIDTestCase, 0, 100)
	for prefixNumber := range 100 {
		prefix := fmt.Sprintf("%02d", prefixNumber)
		_, shouldPass := stringEINRecognizedPrefixes[prefix]
		testCases = append(testCases, stringTaxIDTestCase{
			name:       prefix,
			in:         prefix + "-3456789",
			shouldPass: shouldPass,
		})
	}
	return testCases
}

func generateStringSSNAreaTestCases() []stringTaxIDTestCase {
	testCases := make([]stringTaxIDTestCase, 0, 1_000)
	for area := range 1_000 {
		ssn := fmt.Sprintf("%03d-01-0001", area)
		testCases = append(testCases, stringTaxIDTestCase{
			name:       ssn,
			in:         ssn,
			shouldPass: area != 0 && area != 666 && area < 900,
		})
	}
	return testCases
}

func generateStringSSNGroupTestCases() []stringTaxIDTestCase {
	testCases := make([]stringTaxIDTestCase, 0, 100)
	for group := range 100 {
		ssn := fmt.Sprintf("001-%02d-0001", group)
		testCases = append(testCases, stringTaxIDTestCase{
			name:       ssn,
			in:         ssn,
			shouldPass: group != 0,
		})
	}
	return testCases
}

func generateStringSSNSerialTestCases() []stringTaxIDTestCase {
	testCases := make([]stringTaxIDTestCase, 0, 10_000)
	for serial := range 10_000 {
		ssn := fmt.Sprintf("001-01-%04d", serial)
		testCases = append(testCases, stringTaxIDTestCase{
			name:       ssn,
			in:         ssn,
			shouldPass: serial != 0,
		})
	}
	return testCases
}
