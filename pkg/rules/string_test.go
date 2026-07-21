package rules

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var stringNotEmptyTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"                s", false},
	{"     ", true},
}

func TestStringNotEmpty(t *testing.T) {
	for _, tc := range stringNotEmptyTestCases {
		err := StringNotEmpty().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringNotEmpty))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringNotEmpty(b *testing.B) {
	for _, tc := range stringNotEmptyTestCases {
		rule := StringNotEmpty()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var (
	stringMatchRegexpRegexp    = regexp.MustCompile("[ab]+")
	stringMatchRegexpTestCases = []*struct {
		in            string
		expectedError string
	}{
		{
			in: "ab",
		},
		{
			in:            "cd",
			expectedError: "string must match regular expression: '[ab]+'",
		},
	}
)

func TestStringMatchRegexp(t *testing.T) {
	for _, tc := range stringMatchRegexpTestCases {
		err := StringMatchRegexp(stringMatchRegexpRegexp).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMatchRegexp))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringMatchRegexp(b *testing.B) {
	for _, tc := range stringMatchRegexpTestCases {
		rule := StringMatchRegexp(stringMatchRegexpRegexp)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var (
	stringDenyRegexpRegexp    = regexp.MustCompile("[ab]+")
	stringDenyRegexpTestCases = []*struct {
		in            string
		expectedError string
	}{
		{
			in: "cd",
		},
		{
			in:            "ab",
			expectedError: "string must not match regular expression: '[ab]+'",
		},
	}
)

func TestStringDenyRegexp(t *testing.T) {
	for _, tc := range stringDenyRegexpTestCases {
		err := StringDenyRegexp(stringDenyRegexpRegexp).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDenyRegexp))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDenyRegexp(b *testing.B) {
	for _, tc := range stringDenyRegexpTestCases {
		rule := StringDenyRegexp(stringDenyRegexpRegexp)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringDNSLabelTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test", false},
	{"s", false},
	{"test-this", false},
	{"test-1-this", false},
	{"test1-this", false},
	{"123", false},
	{strings.Repeat("l", 63), false},
	{"", true},
	{strings.Repeat("l", 64), true},
	{"tesT", true},
	{"test?", true},
	{"test this", true},
	{"1_2", true},
	{"LOL", true},
	// cspell:enable
}

func TestStringDNSLabel(t *testing.T) {
	for _, tc := range stringDNSLabelTestCases {
		err := StringDNSLabel().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDNSLabel))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDNSLabel(b *testing.B) {
	for _, tc := range stringDNSLabelTestCases {
		rule := StringDNSLabel()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringDNSSubdomainTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"s", false},
	{"sa", false},
	{"a-1", false},
	{"a--2", false},
	{"a-b-c", false},
	{"a--b--c", false},
	{"0", false},
	{"a.1", false},
	{"a.b", false},
	{"1.b", false},
	{"a.b.c", false},
	{"a.1.c", false},
	{"aa.bb", false},
	{"1.2.3.4", false},
	{"1a.2b.3c.4d", false},
	{"a--b--c.123", false},
	{strings.Repeat("l", 253), false},
	{"", true},
	{" ", true},
	{strings.Repeat("l", 254), true},
	{"tesT", true},
	{"test?", true},
	{"test this", true},
	{"1_2", true},
	{"L", true},
	{"a@b", true},
	{"-", true},
	{"a-", true},
	{"0-", true},
	{"-b", true},
	{"-1", true},
	{"A.1", true},
	{".2.3.4", true},
	{"1a.2B.3c.4d", true},
	{"a--b--c.", true},
	// cspell:enable
}

func TestStringDNSSubdomain(t *testing.T) {
	for _, tc := range stringDNSSubdomainTestCases {
		err := StringDNSSubdomain().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDNSSubdomain))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDNSSubdomain(b *testing.B) {
	for _, tc := range stringDNSSubdomainTestCases {
		rule := StringDNSSubdomain()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringUUIDValidInputs = map[string]string{
	"RFC format example": "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
	"nil UUID":           "00000000-0000-0000-0000-000000000000",
	"max UUID uppercase": "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF",
	"version 1":          "C232AB00-9414-11EC-B3C8-9F6BDECED846",
	"version 3":          "5df41881-3aed-3515-88a7-2f4a814cf09e",
	"version 4":          "919108f7-52d1-4320-9bac-f847db4148a8",
	"version 5":          "2ed6657d-e927-568b-95e1-2665a8aea6a2",
	"version 6":          "1EC9414C-232A-6B00-B3C8-9F6BDECED846",
	"version 7":          "017F22E2-79B0-7CC3-98C4-DC0C0C07398F",
	"version 8":          "2489E9AD-2EE2-8E00-8EC9-32D5F69181C0",
}

var stringUUIDInvalidInputs = map[string]string{
	"empty":                  "",
	"URN representation":     "urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
	"brace representation":   "{f81d4fae-7dec-11d0-a765-00a0c91e6bf6}",
	"compact representation": "f81d4fae7dec11d0a76500a0c91e6bf6",
	"leading space":          " f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
	"trailing newline":       "f81d4fae-7dec-11d0-a765-00a0c91e6bf6\n",
	"too short":              "f81d4fae-7dec-11d0-a765-00a0c91e6bf",
	"too long":               "f81d4fae-7dec-11d0-a765-00a0c91e6bf60",
	"underscore separators":  "f81d4fae_7dec_11d0_a765_00a0c91e6bf6",
	"non-hex character":      "g81d4fae-7dec-11d0-a765-00a0c91e6bf6",
}

func TestStringUUID(t *testing.T) {
	testStringFormatIDRule(
		t,
		StringUUID(),
		ErrorCodeStringUUID,
		"string must match regular expression: '^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$' (e.g. '00000000-0000-0000-0000-000000000000', 'e190c630-8873-11ee-b9d1-0242ac120002', '79258D24-01A7-47E5-ACBB-7E762DE52298'); expected RFC-4122 compliant UUID string",
		stringUUIDValidInputs,
		stringUUIDInvalidInputs,
	)
}

func BenchmarkStringUUID(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUID(), stringUUIDValidInputs, stringUUIDInvalidInputs)
}

var stringUUIDRFC4122ValidInputs = map[string]string{
	"RFC 4122 Appendix B version 1": "7d444840-9dc0-11d1-b245-5ffdce74fad2",
	"RFC 4122 Appendix B version 3": "e902893a-9d22-3c7e-a7b8-d6e313b71d9f",
	"version 1":                     "C232AB00-9414-11EC-B3C8-9F6BDECED846",
	"version 2":                     "f81d4fae-7dec-21d0-a765-00a0c91e6bf6",
	"version 3":                     "5df41881-3aed-3515-88a7-2f4a814cf09e",
	"version 4":                     "919108f7-52d1-4320-9bac-f847db4148a8",
	"version 5":                     "2ed6657d-e927-568b-95e1-2665a8aea6a2",
	"IETF variant lower bound":      "f81d4fae-7dec-11d0-8765-00a0c91e6bf6",
	"IETF variant upper bound":      "f81d4fae-7dec-11d0-b765-00a0c91e6bf6",
}

var stringUUIDRFC4122InvalidInputs = map[string]string{
	"empty":                  "",
	"nil UUID":               "00000000-0000-0000-0000-000000000000",
	"max UUID":               "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF",
	"version 0":              "f81d4fae-7dec-01d0-a765-00a0c91e6bf6",
	"version 6":              "1EC9414C-232A-6B00-B3C8-9F6BDECED846",
	"version 7":              "017F22E2-79B0-7CC3-98C4-DC0C0C07398F",
	"version 8":              "2489E9AD-2EE2-8E00-8EC9-32D5F69181C0",
	"version 9":              "f81d4fae-7dec-91d0-a765-00a0c91e6bf6",
	"version F":              "f81d4fae-7dec-f1d0-a765-00a0c91e6bf6",
	"non-IETF variant lower": "f81d4fae-7dec-11d0-7765-00a0c91e6bf6",
	"non-IETF variant upper": "f81d4fae-7dec-11d0-c765-00a0c91e6bf6",
	"future variant lower":   "00000000-0000-4000-E000-000000000000",
	"future variant upper":   "00000000-0000-4000-F000-000000000000",
	"URN representation":     "urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
	"brace representation":   "{f81d4fae-7dec-11d0-a765-00a0c91e6bf6}",
	"compact representation": "f81d4fae7dec11d0a76500a0c91e6bf6",
	"leading space":          " f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
	"trailing newline":       "f81d4fae-7dec-11d0-a765-00a0c91e6bf6\n",
	"too short":              "f81d4fae-7dec-11d0-a765-00a0c91e6bf",
	"too long":               "f81d4fae-7dec-11d0-a765-00a0c91e6bf60",
	"underscore separators":  "f81d4fae_7dec_11d0_a765_00a0c91e6bf6",
	"non-hex character":      "g81d4fae-7dec-11d0-a765-00a0c91e6bf6",
}

func TestStringUUIDRFC4122(t *testing.T) {
	testStringFormatIDRule(
		t,
		StringUUIDRFC4122(),
		ErrorCodeStringUUIDRFC4122,
		"string must be a valid Universally Unique Identifier (UUID) as defined by RFC 4122",
		stringUUIDRFC4122ValidInputs,
		stringUUIDRFC4122InvalidInputs,
	)
}

func BenchmarkStringUUIDRFC4122(b *testing.B) {
	benchmarkStringFormatIDRule(
		b,
		StringUUIDRFC4122(),
		stringUUIDRFC4122ValidInputs,
		stringUUIDRFC4122InvalidInputs,
	)
}

var stringUUIDv3ValidInputs = map[string]string{
	"official vector, IETF variant lower bound": "5df41881-3aed-3515-88a7-2f4a814cf09e",
	"uppercase":                "5DF41881-3AED-3515-88A7-2F4A814CF09E",
	"mixed case":               "5dF41881-3aED-3515-88A7-2f4A814cF09E",
	"IETF variant upper bound": "5df41881-3aed-3515-b8a7-2f4a814cf09e",
}

var stringUUIDv3InvalidInputs = map[string]string{
	"empty":                  "",
	"adjacent version 2":     "5df41881-3aed-2515-88a7-2f4a814cf09e",
	"adjacent version 4":     "5df41881-3aed-4515-88a7-2f4a814cf09e",
	"non-IETF variant lower": "5df41881-3aed-3515-78a7-2f4a814cf09e",
	"non-IETF variant upper": "5df41881-3aed-3515-c8a7-2f4a814cf09e",
	"URN representation":     "urn:uuid:5df41881-3aed-3515-88a7-2f4a814cf09e",
	"brace representation":   "{5df41881-3aed-3515-88a7-2f4a814cf09e}",
	"compact representation": "5df418813aed351588a72f4a814cf09e",
	"leading space":          " 5df41881-3aed-3515-88a7-2f4a814cf09e",
	"trailing newline":       "5df41881-3aed-3515-88a7-2f4a814cf09e\n",
	"too short":              "5df41881-3aed-3515-88a7-2f4a814cf09",
	"too long":               "5df41881-3aed-3515-88a7-2f4a814cf09e0",
	"underscore separators":  "5df41881_3aed_3515_88a7_2f4a814cf09e",
	"non-hex character":      "gdf41881-3aed-3515-88a7-2f4a814cf09e",
}

func TestStringUUIDv3(t *testing.T) {
	testStringFormatIDRule(
		t,
		StringUUIDv3(),
		ErrorCodeStringUUIDv3,
		"string must be a valid version 3 Universally Unique Identifier (UUID) as defined by RFC 4122",
		stringUUIDv3ValidInputs,
		stringUUIDv3InvalidInputs,
	)
}

func BenchmarkStringUUIDv3(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv3(), stringUUIDv3ValidInputs, stringUUIDv3InvalidInputs)
}

var stringUUIDv4ValidInputs = map[string]string{
	"official vector":          "919108f7-52d1-4320-9bac-f847db4148a8",
	"uppercase":                "919108F7-52D1-4320-9BAC-F847DB4148A8",
	"mixed case":               "919108F7-52d1-4320-9bAc-F847dB4148a8",
	"IETF variant lower bound": "919108f7-52d1-4320-8bac-f847db4148a8",
	"IETF variant upper bound": "919108f7-52d1-4320-bbac-f847db4148a8",
}

var stringUUIDv4InvalidInputs = map[string]string{
	"empty":                  "",
	"adjacent version 3":     "919108f7-52d1-3320-9bac-f847db4148a8",
	"adjacent version 5":     "919108f7-52d1-5320-9bac-f847db4148a8",
	"non-IETF variant lower": "919108f7-52d1-4320-7bac-f847db4148a8",
	"non-IETF variant upper": "919108f7-52d1-4320-cbac-f847db4148a8",
	"URN representation":     "urn:uuid:919108f7-52d1-4320-9bac-f847db4148a8",
	"brace representation":   "{919108f7-52d1-4320-9bac-f847db4148a8}",
	"compact representation": "919108f752d143209bacf847db4148a8",
	"leading space":          " 919108f7-52d1-4320-9bac-f847db4148a8",
	"trailing newline":       "919108f7-52d1-4320-9bac-f847db4148a8\n",
	"too short":              "919108f7-52d1-4320-9bac-f847db4148a",
	"too long":               "919108f7-52d1-4320-9bac-f847db4148a80",
	"underscore separators":  "919108f7_52d1_4320_9bac_f847db4148a8",
	"non-hex character":      "g19108f7-52d1-4320-9bac-f847db4148a8",
}

func TestStringUUIDv4(t *testing.T) {
	testStringFormatIDRule(
		t,
		StringUUIDv4(),
		ErrorCodeStringUUIDv4,
		"string must be a valid version 4 Universally Unique Identifier (UUID) as defined by RFC 4122",
		stringUUIDv4ValidInputs,
		stringUUIDv4InvalidInputs,
	)
}

func BenchmarkStringUUIDv4(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv4(), stringUUIDv4ValidInputs, stringUUIDv4InvalidInputs)
}

var stringUUIDv5ValidInputs = map[string]string{
	"official vector":          "2ed6657d-e927-568b-95e1-2665a8aea6a2",
	"uppercase":                "2ED6657D-E927-568B-95E1-2665A8AEA6A2",
	"mixed case":               "2Ed6657D-e927-568b-95E1-2665a8AeA6a2",
	"IETF variant lower bound": "2ed6657d-e927-568b-85e1-2665a8aea6a2",
	"IETF variant upper bound": "2ed6657d-e927-568b-b5e1-2665a8aea6a2",
}

var stringUUIDv5InvalidInputs = map[string]string{
	"empty":                  "",
	"adjacent version 4":     "2ed6657d-e927-468b-95e1-2665a8aea6a2",
	"adjacent version 6":     "2ed6657d-e927-668b-95e1-2665a8aea6a2",
	"non-IETF variant lower": "2ed6657d-e927-568b-75e1-2665a8aea6a2",
	"non-IETF variant upper": "2ed6657d-e927-568b-c5e1-2665a8aea6a2",
	"URN representation":     "urn:uuid:2ed6657d-e927-568b-95e1-2665a8aea6a2",
	"brace representation":   "{2ed6657d-e927-568b-95e1-2665a8aea6a2}",
	"compact representation": "2ed6657de927568b95e12665a8aea6a2",
	"leading space":          " 2ed6657d-e927-568b-95e1-2665a8aea6a2",
	"trailing newline":       "2ed6657d-e927-568b-95e1-2665a8aea6a2\n",
	"too short":              "2ed6657d-e927-568b-95e1-2665a8aea6a",
	"too long":               "2ed6657d-e927-568b-95e1-2665a8aea6a20",
	"underscore separators":  "2ed6657d_e927_568b_95e1_2665a8aea6a2",
	"non-hex character":      "ged6657d-e927-568b-95e1-2665a8aea6a2",
}

func TestStringUUIDv5(t *testing.T) {
	testStringFormatIDRule(
		t,
		StringUUIDv5(),
		ErrorCodeStringUUIDv5,
		"string must be a valid version 5 Universally Unique Identifier (UUID) as defined by RFC 4122",
		stringUUIDv5ValidInputs,
		stringUUIDv5InvalidInputs,
	)
}

func BenchmarkStringUUIDv5(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv5(), stringUUIDv5ValidInputs, stringUUIDv5InvalidInputs)
}

// cspell:ignore ttttttttttrrrrrrrrrrrrrrrr

// stringULIDValidInputs includes every concrete valid ULID from ulid/spec at
// revision d0c7170df4517939e70129b4d6462cc162f2d5bf and every concrete ULID
// from ulid/javascript's tests at revision 11c2067821ee19e4dc787ca4e0125a025485edc6.
//
// The specification's `ttttttttttrrrrrrrrrrrrrrrr` representation describes
// the field layout; it is not a concrete ULID and is intentionally excluded.
var stringULIDValidInputs = map[string]string{
	// cspell:disable
	"spec introductory example":      "01ARZ3NDEKTSV4RRFFQ69G5FAV",
	"spec layout example":            "01AN4Z07BY79KA1307SR9X4MV3",
	"spec monotonic first":           "01BX5ZZKBKACTAV9WEVGEMMVRY",
	"spec monotonic second":          "01BX5ZZKBKACTAV9WEVGEMMVRZ",
	"spec monotonic third":           "01BX5ZZKBKACTAV9WEVGEMMVS0",
	"spec monotonic fourth":          "01BX5ZZKBKACTAV9WEVGEMMVS1",
	"spec near-overflow X":           "01BX5ZZKBKZZZZZZZZZZZZZZZX",
	"spec near-overflow Y":           "01BX5ZZKBKZZZZZZZZZZZZZZZY",
	"spec near-overflow Z":           "01BX5ZZKBKZZZZZZZZZZZZZZZZ",
	"spec maximum value":             "7ZZZZZZZZZZZZZZZZZZZZZZZZZ",
	"JavaScript decode-time example": "01ARYZ6S41TSV4RRFFQ69G5FAV",
	"JavaScript monotonic first":     "01ARYZ6S41YYYYYYYYYYYYYYYY",
	"JavaScript monotonic second":    "01ARYZ6S41YYYYYYYYYYYYYYYZ",
	"JavaScript monotonic third":     "01ARYZ6S41YYYYYYYYYYYYYYZ0",
	"JavaScript monotonic fourth":    "01ARYZ6S41YYYYYYYYYYYYYYZ1",
	"JavaScript next millisecond":    "01ARYZ6S42YYYYYYYYYYYYYYYY",
	"derived minimum value":          "00000000000000000000000000",
	"derived lowercase example":      "01arz3ndektsv4rrffq69g5fav",
	"derived lowercase maximum":      "7zzzzzzzzzzzzzzzzzzzzzzzzz",
	// cspell:enable
}

var stringULIDInvalidInputs = map[string]string{
	"empty":                        "",
	"too short":                    "01ARZ3NDEKTSV4RRFFQ69G5FA",
	"too long":                     "01ARZ3NDEKTSV4RRFFQ69G5FAV0",
	"minimum overflow":             "80000000000000000000000000",
	"overflow with maximum suffix": "8ZZZZZZZZZZZZZZZZZZZZZZZZZ",
	"all Z":                        "ZZZZZZZZZZZZZZZZZZZZZZZZZZ",
	"forbidden uppercase I":        "01ARZ3NDEKTSV4RRFFQ69G5FIV",
	"forbidden uppercase L":        "01ARZ3NDEKTSV4RRFFQ69G5FLV",
	"forbidden uppercase O":        "01ARZ3NDEKTSV4RRFFQ69G5FOV",
	"forbidden uppercase U":        "01ARZ3NDEKTSV4RRFFQ69G5FUV",
	"forbidden lowercase I":        "01arz3ndektsv4rrffq69g5fiv",
	"forbidden lowercase L":        "01arz3ndektsv4rrffq69g5flv",
	"forbidden lowercase O":        "01arz3ndektsv4rrffq69g5fov",
	"forbidden lowercase U":        "01arz3ndektsv4rrffq69g5fuv",
	"leading space":                " 1ARZ3NDEKTSV4RRFFQ69G5FAV",
	"trailing space":               "01ARZ3NDEKTSV4RRFFQ69G5FA ",
	"leading newline":              "\n1ARZ3NDEKTSV4RRFFQ69G5FAV",
	"trailing newline":             "01ARZ3NDEKTSV4RRFFQ69G5FA\n",
	"hyphen":                       "01ARZ3NDEKTSV4RRFFQ69G5F-V",
	"full-width characters":        "０１ＡＮ４Ｚ０７ＢＹ７９ＫＡ１３０７ＳＲ９Ｘ４ＭＶ３",
}

func TestStringULID(t *testing.T) {
	testStringFormatIDRule(
		t,
		StringULID(),
		ErrorCodeStringULID,
		"string must be a valid Universally Unique Lexicographically Sortable Identifier (ULID)",
		stringULIDValidInputs,
		stringULIDInvalidInputs,
	)
}

func BenchmarkStringULID(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringULID(), stringULIDValidInputs, stringULIDInvalidInputs)
}

var stringASCIITestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"", false},
	{"foobar", false},
	{"0987654321", false},
	{"test@example.com", false},
	{"1234abcDEF", false},
	{"newline\n", false},
	{"\x19test\x7F", false},
	{"ｆｏｏbar", true},
	{"ｘｙｚ０９８", true},
	{"１２３456", true},
	{"ｶﾀｶﾅ", true},
	// cspell:enable
}

func TestStringASCII(t *testing.T) {
	for _, tc := range stringASCIITestCases {
		err := StringASCII().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringASCII))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringASCII(b *testing.B) {
	for _, tc := range stringASCIITestCases {
		rule := StringASCII()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringEmailTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test@mail.com", false},
	{"Dörte@Sörensen.example.com", false},
	{"θσερ@εχαμπλε.ψομ", false},
	{"юзер@екзампл.ком", false},
	{"उपयोगकर्ता@उदाहरण.कॉम", false},
	{"用户@例子.广告", false},
	{`"test test"@email.com`, false},
	{"mail@domain_with_underscores.org", false},
	{"test@email", false},
	{"test@t", false},
	{"", true},
	{"test@", true},
	{"test", true},
	{"test@email.", true},
	{"@email.com", true},
	{`"@email.com`, true},
	// cspell:enable
}

func TestStringEmail(t *testing.T) {
	for _, tc := range stringEmailTestCases {
		err := StringEmail().Validate(tc.in)
		if tc.shouldFail {
			assert.ErrorContains(t, err, "string must be a valid email address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEmail))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringEmail(b *testing.B) {
	for _, tc := range stringEmailTestCases {
		rule := StringEmail()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func TestStringURL(t *testing.T) {
	for _, tc := range urlTestCases {
		err := StringURL().Validate(tc.url)
		if tc.shouldFail {
			assert.Require(t, assert.Error(t, err))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringURL))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("failed to parse url", func(t *testing.T) {
		err := StringURL().Validate("http://\x1f")
		assert.ErrorContains(
			t,
			err,
			"failed to parse URL: parse \"http://\\x1f\": net/url: invalid control character in URL",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringURL))
	})
}

func BenchmarkStringURL(b *testing.B) {
	for _, tc := range urlTestCases {
		rule := StringURL()
		for range b.N {
			_ = rule.Validate(tc.url)
		}
	}
}

var stringMACTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"3D:F2:C9:A6:B3:4F", false},
	{"00:25:96:FF:FE:12:34:56", false},
	{"3D-F2-C9-A6-B3:4F", true},
	{"123", true},
	{"", true},
	{"abacaba", true},
	{"0025:96FF:FE12:3456", true},
	// cspell:enable
}

func TestStringMAC(t *testing.T) {
	for _, tc := range stringMACTestCases {
		err := StringMAC().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid MAC address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMAC))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringMAC(b *testing.B) {
	for _, tc := range stringMACTestCases {
		rule := StringMAC()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringIPTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"10.0.0.1", false},
	{"172.16.0.1", false},
	{"192.168.0.1", false},
	{"192.168.255.254", false},
	{"172.16.255.254", false},
	{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
	{"2001:cdba:0:0:0:0:3257:9652", false},
	{"2001:cdba::3257:9652", false},
	{"", true},
	{"172.16.256.255", true},
	{"192.168.255.256", true},
	// cspell:enable
}

func TestStringIP(t *testing.T) {
	for _, tc := range stringIPTestCases {
		err := StringIP().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IP address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIP))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringIP(b *testing.B) {
	for _, tc := range stringIPTestCases {
		rule := StringIP()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringIPv4TestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"10.0.0.1", false},
	{"172.16.0.1", false},
	{"192.168.0.1", false},
	{"192.168.255.254", false},
	{"172.16.255.254", false},
	{"192.168.255.256", true},
	{"172.16.256.255", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652", true},
	{"2001:cdba:0:0:0:0:3257:9652", true},
	{"2001:cdba::3257:9652", true},
	// cspell:enable
}

func TestStringIPv4(t *testing.T) {
	for _, tc := range stringIPv4TestCases {
		err := StringIPv4().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IPv4 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIPv4))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringIPv4(b *testing.B) {
	for _, tc := range stringIPv4TestCases {
		rule := StringIPv4()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringIPv6TestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
	{"2001:cdba:0:0:0:0:3257:9652", false},
	{"2001:cdba::3257:9652", false},
	{"10.0.0.1", true},
	{"172.16.0.1", true},
	{"192.168.0.1", true},
	{"192.168.255.254", true},
	{"192.168.255.256", true},
	{"172.16.255.254", true},
	{"172.16.256.255", true},
	// cspell:enable
}

func TestStringIPv6(t *testing.T) {
	for _, tc := range stringIPv6TestCases {
		err := StringIPv6().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IPv6 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIPv6))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringIPv6(b *testing.B) {
	for _, tc := range stringIPv6TestCases {
		rule := StringIPv6()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringCIDRTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"10.0.0.0/0", false},
	{"10.0.0.1/8", false},
	{"172.16.0.1/16", false},
	{"192.168.0.1/24", false},
	{"192.168.255.254/24", false},
	{"172.16.255.254/16", false},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/64", false},
	{"2001:cdba:0:0:0:0:3257:9652/32", false},
	{"2001:cdba::3257:9652/16", false},
	{"192.168.255.254/48", true},
	{"192.168.255.256/24", true},
	{"172.16.256.255/16", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/256", true},
	// cspell:enable
}

func TestStringCIDR(t *testing.T) {
	for _, tc := range stringCIDRTestCases {
		err := StringCIDR().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IP address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDR))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringCIDR(b *testing.B) {
	for _, tc := range stringCIDRTestCases {
		rule := StringCIDR()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringCIDRv4TestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"0.0.0.0/0", false},
	{"10.0.0.0/8", false},
	{"172.16.0.0/16", false},
	{"192.168.0.0/24", false},
	{"172.16.0.0/16", false},
	{"192.168.255.0/24", false},
	{"10.0.0.0/0", true},
	{"10.0.0.1/8", true},
	{"172.16.0.1/16", true},
	{"192.168.0.1/24", true},
	{"192.168.255.254/24", true},
	{"192.168.255.254/48", true},
	{"192.168.255.256/24", true},
	{"172.16.255.254/16", true},
	{"172.16.256.255/16", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/64", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/256", true},
	{"2001:cdba:0:0:0:0:3257:9652/32", true},
	{"2001:cdba::3257:9652/16", true},
	{"172.56.1.0/16", true},
	// cspell:enable
}

func TestStringCIDRv4(t *testing.T) {
	for _, tc := range stringCIDRv4TestCases {
		err := StringCIDRv4().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IPv4 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDRv4))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringCIDRv4(b *testing.B) {
	for _, tc := range stringCIDRv4TestCases {
		rule := StringCIDRv4()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringCIDRv6TestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"2001:cdba:0000:0000:0000:0000:3257:9652/64", false},
	{"2001:cdba:0:0:0:0:3257:9652/32", false},
	{"2001:cdba::3257:9652/16", false},
	{"10.0.0.0/0", true},
	{"10.0.0.1/8", true},
	{"172.16.0.1/16", true},
	{"192.168.0.1/24", true},
	{"192.168.255.254/24", true},
	{"192.168.255.254/48", true},
	{"192.168.255.256/24", true},
	{"172.16.255.254/16", true},
	{"172.16.256.255/16", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/256", true},
	// cspell:enable
}

func TestStringCIDRv6(t *testing.T) {
	for _, tc := range stringCIDRv6TestCases {
		err := StringCIDRv6().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IPv6 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDRv6))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringCIDRv6(b *testing.B) {
	for _, tc := range stringCIDRv6TestCases {
		rule := StringCIDRv6()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringJSONTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{`{"foo": "bar"}`, false},
	{`{}`, false},
	{`[]`, false},
	{"{]}", true},
	{"", true},
	{"yaml: ok", true},
}

func TestStringJSON(t *testing.T) {
	for _, tc := range stringJSONTestCases {
		err := StringJSON().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringJSON))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringJSON(b *testing.B) {
	for _, tc := range stringJSONTestCases {
		rule := StringJSON()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var validSemverTestCases = []string{
	"0.0.4",
	"1.2.3",
	"10.20.30",
	"1.1.2-prerelease+meta",
	"1.1.2+meta",
	"1.1.2+meta-valid",
	"1.0.0-alpha",
	"1.0.0-beta",
	"1.0.0-alpha.beta",
	"1.0.0-alpha.beta.1",
	"1.0.0-alpha.1",
	"1.0.0-alpha0.valid",
	"1.0.0-alpha.0valid",
	"1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay",
	"1.0.0-rc.1+build.1",
	"2.0.0-rc.1+build.123",
	"1.2.3-beta",
	"10.2.3-DEV-SNAPSHOT",
	"1.2.3-SNAPSHOT-123",
	"1.0.0",
	"2.0.0",
	"1.1.7",
	"2.0.0+build.1848",
	"2.0.1-alpha.1227",
	"1.0.0-alpha+beta",
	"1.2.3----RC-SNAPSHOT.12.9.1--.12+788",
	"1.2.3----R-S.12.9.1--.12+meta",
	"1.2.3----RC-SNAPSHOT.12.9.1--.12",
	"1.0.0+0.build.1-rc.10000aaa-kk-0.1",
	"99999999999999999999999.999999999999999999.99999999999999999",
	"1.0.0-0A.is.legal",
	"0.1.0",
	"1.0.0+20130313144700",
	"1.0.0-beta+exp.sha.5114f85",
	"2.7.3-rc.1+build.11.e0f985a",
}

var invalidSemverTestCases = []string{
	"1",
	"1.2",
	"1.2.3-0123",
	"1.2.3-0123.0123",
	"1.1.2+.123",
	"+invalid",
	"-invalid",
	"-invalid+invalid",
	"-invalid.01",
	"alpha",
	"alpha.beta",
	"alpha.beta.1",
	"alpha.1",
	"alpha+beta",
	"alpha_beta",
	"alpha.",
	"alpha..",
	"beta",
	"1.0.0-alpha_beta",
	"-alpha.",
	"1.0.0-alpha..",
	"1.0.0-alpha..1",
	"1.0.0-alpha...1",
	"1.0.0-alpha....1",
	"1.0.0-alpha.....1",
	"1.0.0-alpha......1",
	"1.0.0-alpha.......1",
	"01.1.1",
	"1.01.1",
	"1.1.01",
	"1.2",
	"1.2.3.DEV",
	"1.2-SNAPSHOT",
	"1.2.31.2.3----RC-SNAPSHOT.12.09.1--..12+788",
	"1.2-RC-SNAPSHOT",
	"-1.0.3-gamma+b7718",
	"+justmeta",
	"9.8.7+meta+meta",
	"9.8.7-whatever+meta+meta",
	"99999999999999999999999.999999999999999999.99999999999999999----RC-SNAPSHOT.12.09.1--------------------------------..12",
	"",
	"1.2.3.4",
	"01.2.3",
	"1.02.3",
	"1.2.03",
	"1.2.3-",
	"1.2.3-01",
	"v1.2.3",
	"1.2.3+build..1",
}

func TestStringSemver(t *testing.T) {
	rule := StringSemver()
	t.Run("valid versions", func(t *testing.T) {
		for _, version := range validSemverTestCases {
			t.Run(fmt.Sprintf("%q", version), func(t *testing.T) {
				assert.NoError(t, rule.Validate(version))
			})
		}
	})
	t.Run("invalid versions", func(t *testing.T) {
		for _, version := range invalidSemverTestCases {
			t.Run(fmt.Sprintf("%q", version), func(t *testing.T) {
				err := rule.Validate(version)
				assert.EqualError(t, err, "string must be a valid semantic version")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSemver))
			})
		}
	})
}

func BenchmarkStringSemver(b *testing.B) {
	rule := StringSemver()
	for b.Loop() {
		_ = rule.Validate("2.7.3-rc.1+build.11.e0f985a")
	}
}

var stringCVETestCases = map[string]struct {
	in            string
	expectedError string
}{
	"four digit sequence": {
		in: "CVE-1999-0001",
	},
	"four digit year before 1999": {
		in: "CVE-1998-0001",
	},
	"zero sequence": {
		in: "CVE-2021-0000",
	},
	"sequence with leading zero": {
		in: "CVE-2014-0160",
	},
	"five digit sequence with leading zero": {
		in: "CVE-2021-00001",
	},
	"sequence with two leading zeroes": {
		in: "CVE-2021-0990",
	},
	"five digit sequence": {
		in: "CVE-2021-44228",
	},
	"long sequence": {
		in: "CVE-2024-12345",
	},
	"nineteen digit sequence": {
		in: "CVE-2024-1234567890123456789",
	},
	"empty": {
		in:            "",
		expectedError: "string must be a valid CVE ID",
	},
	"lowercase prefix": {
		in:            "cve-2021-44228",
		expectedError: "string must be a valid CVE ID",
	},
	"short sequence": {
		in:            "CVE-2021-123",
		expectedError: "string must be a valid CVE ID",
	},
	"letters in sequence": {
		in:            "CVE-2021-ABCD",
		expectedError: "string must be a valid CVE ID",
	},
	"five digit year": {
		in:            "CVE-10000-0001",
		expectedError: "string must be a valid CVE ID",
	},
	"twenty digit sequence": {
		in:            "CVE-2024-12345678901234567890",
		expectedError: "string must be a valid CVE ID",
	},
}

func TestStringCVE(t *testing.T) {
	for name, tt := range stringCVETestCases {
		t.Run(name, func(t *testing.T) {
			err := StringCVE().Validate(tt.in)
			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCVE))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

var stringE164TestCases = map[string]struct {
	in            string
	expectedError string
}{
	"minimum length": {
		in: "+12",
	},
	"maximum length": {
		in: "+123456789012345",
	},
	"common US number": {
		in: "+14155552671",
	},
	"missing plus sign": {
		in:            "14155552671",
		expectedError: "string must be a valid E.164 phone number",
	},
	"starts with zero": {
		in:            "+0123456789",
		expectedError: "string must be a valid E.164 phone number",
	},
	"too short": {
		in:            "+1",
		expectedError: "string must be a valid E.164 phone number",
	},
	"too long": {
		in:            "+1234567890123456",
		expectedError: "string must be a valid E.164 phone number",
	},
	"contains spaces": {
		in:            "+1 4155552671",
		expectedError: "string must be a valid E.164 phone number",
	},
	"contains punctuation": {
		in:            "+1-415-555-2671",
		expectedError: "string must be a valid E.164 phone number",
	},
	"empty": {
		expectedError: "string must be a valid E.164 phone number",
	},
}

func TestStringE164(t *testing.T) {
	t.Parallel()

	for name, tt := range stringE164TestCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := StringE164().Validate(tt.in)
			if tt.expectedError != "" {
				assert.Require(t, assert.Error(t, err))
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringE164))
				return
			}
			assert.NoError(t, err)
		})
	}
}

func BenchmarkStringE164(b *testing.B) {
	tests := map[string]string{
		"valid":   "+14155552671",
		"invalid": "+1-415-555-2671",
	}

	for name, in := range tests {
		b.Run(name, func(b *testing.B) {
			rule := StringE164()
			for b.Loop() {
				_ = rule.Validate(in)
			}
		})
	}
}

func BenchmarkStringCVE(b *testing.B) {
	rule := StringCVE()
	for b.Loop() {
		_ = rule.Validate("CVE-2021-44228")
	}
}

var stringContainsTestCases = []*struct {
	in            string
	substrings    []string
	expectedError string
}{
	{
		in:         "",
		substrings: []string{""},
	},
	{
		in:         "this",
		substrings: []string{"his"},
	},
	{
		in:         "this",
		substrings: []string{"this"},
	},
	{
		in:         "this",
		substrings: []string{"th", "is"},
	},
	{
		in:            "one",
		substrings:    []string{"th"},
		expectedError: "string must contain the following substrings: 'th'",
	},
	{
		in:            "this",
		substrings:    []string{"th", "ht"},
		expectedError: "string must contain the following substrings: 'th', 'ht'",
	},
	{
		in:            "tha",
		substrings:    []string{"that"},
		expectedError: "string must contain the following substrings: 'that'",
	},
}

func TestStringContains(t *testing.T) {
	for _, tc := range stringContainsTestCases {
		err := StringContains(tc.substrings...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringContains))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringContains(b *testing.B) {
	for _, tc := range stringContainsTestCases {
		rule := StringContains(tc.substrings...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringExcludesTestCases = []*struct {
	in            string
	substrings    []string
	expectedError string
}{
	{
		in:         "one",
		substrings: []string{"th"},
	},
	{
		in:         "this",
		substrings: []string{"tho", "ht"},
	},
	{
		in:         "tha",
		substrings: []string{"that"},
	},
	{
		in:            "",
		substrings:    []string{""},
		expectedError: "string must not contain any of the following substrings: ''",
	},
	{
		in:            "this",
		substrings:    []string{"his"},
		expectedError: "string must not contain any of the following substrings: 'his'",
	},
	{
		in:            "this",
		substrings:    []string{"this"},
		expectedError: "string must not contain any of the following substrings: 'this'",
	},
	{
		in:            "this",
		substrings:    []string{"th", "is"},
		expectedError: "string must not contain any of the following substrings: 'th', 'is'",
	},
}

func TestStringExcludes(t *testing.T) {
	for _, tc := range stringExcludesTestCases {
		err := StringExcludes(tc.substrings...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringExcludes))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringExcludes(b *testing.B) {
	for _, tc := range stringExcludesTestCases {
		rule := StringExcludes(tc.substrings...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringStartsWithTestCases = []*struct {
	in            string
	prefixes      []string
	expectedError string
}{
	{
		in:       "this",
		prefixes: []string{"th"},
	},
	{
		in:       "this",
		prefixes: []string{"is", "th"},
	},
	{
		in:            "one",
		prefixes:      []string{"th"},
		expectedError: "string must start with 'th' prefix",
	},
	{
		in:            "one",
		prefixes:      []string{"th", "ht"},
		expectedError: "string must start with one of the following prefixes: 'th', 'ht'",
	},
}

func TestStringStartsWith(t *testing.T) {
	for _, tc := range stringStartsWithTestCases {
		err := StringStartsWith(tc.prefixes...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringStartsWith))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringStartsWith(b *testing.B) {
	for _, tc := range stringStartsWithTestCases {
		rule := StringStartsWith(tc.prefixes...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringEndsWithTestCases = []*struct {
	in            string
	suffixes      []string
	expectedError string
}{
	{
		in:       "this",
		suffixes: []string{"is"},
	},
	{
		in:       "this",
		suffixes: []string{"th", "is"},
	},
	{
		in:            "one",
		suffixes:      []string{"th"},
		expectedError: "string must end with 'th' suffix",
	},
	{
		in:            "one",
		suffixes:      []string{"th", "ht"},
		expectedError: "string must end with one of the following suffixes: 'th', 'ht'",
	},
}

func TestStringEndsWith(t *testing.T) {
	for _, tc := range stringEndsWithTestCases {
		err := StringEndsWith(tc.suffixes...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEndsWith))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringEndsWith(b *testing.B) {
	for _, tc := range stringEndsWithTestCases {
		rule := StringEndsWith(tc.suffixes...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringTitleTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"", true},
	{"a", true},
	{"A", false},
	{" aaa aaa aaa ", true},
	{" Aaa Aaa Aaa ", false},
	{"123a456", true},
	{"double-blind", true},
	{"Double-Blind", false},
	{"ÿøû", true},
	{"Ÿøû", false},
	{"with_underscore", true},
	{"With_underscore", false},
	{"unicode \xe2\x80\xa8 line separator", true},
	{"Unicode \xe2\x80\xa8 Line Separator", false},
	// cspell:enable
}

func TestStringTitle(t *testing.T) {
	for _, tc := range stringTitleTestCases {
		err := StringTitle().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "each word in a string must start with a capital letter")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringTitle))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringTitle(b *testing.B) {
	for _, tc := range stringTitleTestCases {
		rule := StringTitle()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var (
	errGitRefEmpty           = errors.New("git reference must not be empty")
	errGitRefEndsWithDot     = errors.New("git reference must not end with a '.'")
	errGitRefAtLeastOneSlash = errors.New("git reference must contain at least one '/'")
	errGitRefEmptyPart       = errors.New("git reference must not have empty parts")
	errGitRefStartsWithDash  = errors.New("git branch and tag references must not start with '-'")
	errGitRefForbiddenChars  = errors.New("git reference contains forbidden characters")
)

var stringGitRefTestCases = []*struct {
	in          string
	expectedErr error
}{
	{"refs/heads/master", nil},
	{"refs/notes/commits", nil},
	{"refs/tags/this@", nil},
	{"refs/remotes/origin/master", nil},
	{"HEAD", nil},
	{"refs/tags/v3.1.1", nil},
	{"refs/pulls/1/head", nil},
	{"refs/pulls/1/merge", nil},
	{"refs/pulls/1/abc.123", nil},
	{"refs/pulls", nil},
	{"refs/-", nil},
	{"refs", errGitRefAtLeastOneSlash},
	{"refs/", errGitRefEmptyPart},
	{"refs//", errGitRefEmptyPart},
	{"refs/heads/\\", errGitRefForbiddenChars},
	{"refs/heads/\\foo", errGitRefForbiddenChars},
	{"refs/heads/\\foo/bar", errGitRefForbiddenChars},
	{"abc", errGitRefAtLeastOneSlash},
	{"", errGitRefEmpty},
	{"refs/heads/ ", errGitRefForbiddenChars},
	{"refs/heads/ /", errGitRefForbiddenChars},
	{"refs/heads/ /foo", errGitRefForbiddenChars},
	{"refs/heads/.", errGitRefEndsWithDot},
	{"refs/heads/..", errGitRefEndsWithDot},
	{"refs/heads/foo..", errGitRefEndsWithDot},
	{"refs/heads/foo.lock", errGitRefForbiddenChars},
	{"refs/heads/foo@{bar}", errGitRefForbiddenChars},
	{"refs/heads/foo@{", errGitRefForbiddenChars},
	{"refs/heads/foo[", errGitRefForbiddenChars},
	{"refs/heads/foo~", errGitRefForbiddenChars},
	{"refs/heads/foo^", errGitRefForbiddenChars},
	{"refs/heads/foo:", errGitRefForbiddenChars},
	{"refs/heads/foo?", errGitRefForbiddenChars},
	{"refs/heads/foo*", errGitRefForbiddenChars},
	{"refs/heads/foo[bar", errGitRefForbiddenChars},
	{"refs/heads/foo\t", errGitRefForbiddenChars},
	{"refs/heads/@", errGitRefForbiddenChars},
	{"refs/heads/@{bar}", errGitRefForbiddenChars},
	{"refs/heads/\n", errGitRefForbiddenChars},
	{"refs/heads/-foo", errGitRefStartsWithDash},
	{"refs/heads/foo..bar", errGitRefForbiddenChars},
	{"refs/heads/-", errGitRefStartsWithDash},
	{"refs/tags/-", errGitRefStartsWithDash},
	{"refs/tags/-foo", errGitRefStartsWithDash},
}

func TestStringGitRef(t *testing.T) {
	for _, tc := range stringGitRefTestCases {
		t.Run(tc.in, func(t *testing.T) {
			err := StringGitRef().Validate(tc.in)
			if tc.expectedErr != nil {
				assert.ErrorContains(t, err, tc.expectedErr.Error())
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringGitRef))
				assert.ErrorContains(
					t,
					err,
					"see https://git-scm.com/docs/git-check-ref-format for more information on Git reference naming rules",
				)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringGitRef(b *testing.B) {
	for _, tc := range stringGitRefTestCases {
		rule := StringGitRef()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func prepareFileSystemTests(t testing.TB) (root string) {
	t.Helper()
	root = t.TempDir()
	t.Setenv("HOME", root)
	for _, path := range []struct {
		path  string
		perm  os.FileMode
		isDir bool
	}{
		{"file1", 0o755, false},
		{"dir1", 0o755, true},
		{"dir1/file2", 0o755, false},
		{"dir-no-perm", 0o000, true},
		{"dir1/file-no-perm", 0o000, false},
	} {
		if path.isDir {
			err := os.MkdirAll(filepath.Join(root, path.path), path.perm)
			assert.Require(t, assert.NoError(t, err))
		} else {
			err := os.WriteFile(filepath.Join(root, path.path), []byte{}, path.perm)
			assert.Require(t, assert.NoError(t, err))
		}
	}
	return root
}

type stringFileSystemPathTestCase struct {
	in          string
	expectedErr error
}

func getStringFileSystemPathTestCases(root string) []*stringFileSystemPathTestCase {
	addRoot := func(path string) string {
		// We're not using filepath.Join because it cleans the path.
		return root + string(filepath.Separator) + path
	}
	return []*stringFileSystemPathTestCase{
		{"~/dir1", nil},
		{"~/dir1/", nil},
		{addRoot("dir1"), nil},
		{addRoot("dir1/file2"), nil},
		{"~/dir1/file2", nil},
		{addRoot("dir1/file2/.."), nil},
		{"~/dir1/file2/..", nil},
		{"~/dir1/file2/../../", nil},
		{addRoot("."), nil},
		{addRoot("./"), nil},
		{addRoot("./file1"), nil},
		{addRoot("dir-no-perm"), nil},
		{addRoot("dir1/file-no-perm"), nil},
		{addRoot("dir1/file2/"), syscall.ENOTDIR},
		{"~/dir1/../file1/", syscall.ENOTDIR},
		{addRoot("non-existing-dir"), errFilePathNotExists},
		{"", errFilePathEmpty},
		{"	", errFilePathEmpty},
	}
}

func TestStringFileSystemPath(t *testing.T) {
	root := prepareFileSystemTests(t)
	for _, tc := range getStringFileSystemPathTestCases(root) {
		err := StringFileSystemPath().Validate(tc.in)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error())
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringFileSystemPath))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringFileSystemPath(b *testing.B) {
	root := prepareFileSystemTests(b)
	testCases := getStringFileSystemPathTestCases(root)
	for _, tc := range testCases {
		rule := StringFileSystemPath()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func getStringFilePathTestCases(root string) []*stringFileSystemPathTestCase {
	addRoot := func(path string) string {
		// We're not using filepath.Join because it cleans the path.
		return root + string(filepath.Separator) + path
	}
	return []*stringFileSystemPathTestCase{
		{addRoot("dir1/file2"), nil},
		{"~/dir1/file2", nil},
		{addRoot("./file1"), nil},
		{addRoot("dir1/file-no-perm"), nil},
		{addRoot("dir-no-perm"), errFilePathNotFile},
		{addRoot("dir1"), errFilePathNotFile},
		{addRoot("dir1/file2/.."), errFilePathNotFile},
		{addRoot("."), errFilePathNotFile},
		{addRoot("./"), errFilePathNotFile},
		{"~/dir1/file2/..", errFilePathNotFile},
		{"~/dir1/file2/../../", errFilePathNotFile},
		{"~/dir1", errFilePathNotFile},
		{"~/dir1/", errFilePathNotFile},
		{addRoot("dir1/file2/"), syscall.ENOTDIR},
		{"~/dir1/../file1/", syscall.ENOTDIR},
		{addRoot("non-existing-dir"), errFilePathNotExists},
		{"", errFilePathEmpty},
		{"	", errFilePathEmpty},
	}
}

func TestStringFilePath(t *testing.T) {
	root := prepareFileSystemTests(t)
	for _, tc := range getStringFilePathTestCases(root) {
		err := StringFilePath().Validate(tc.in)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error())
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringFilePath))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringFilePath(b *testing.B) {
	root := prepareFileSystemTests(b)
	testCases := getStringFilePathTestCases(root)
	for _, tc := range testCases {
		rule := StringFilePath()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func getStringDirPathTestCases(root string) []*stringFileSystemPathTestCase {
	addRoot := func(path string) string {
		// We're not using filepath.Join because it cleans the path.
		return root + string(filepath.Separator) + path
	}
	return []*stringFileSystemPathTestCase{
		{addRoot("dir1"), nil},
		{addRoot("dir1/file2/.."), nil},
		{addRoot("."), nil},
		{addRoot("./"), nil},
		{"~/dir1/file2/..", nil},
		{"~/dir1/file2/../../", nil},
		{"~/dir1", nil},
		{"~/dir1/", nil},
		{addRoot("dir-no-perm"), nil},
		{addRoot("dir1/file-no-perm"), errFilePathNotDir},
		{addRoot("dir1/file2"), errFilePathNotDir},
		{"~/dir1/file2", errFilePathNotDir},
		{addRoot("./file1"), errFilePathNotDir},
		{addRoot("dir1/file2/"), syscall.ENOTDIR},
		{"~/dir1/../file1/", syscall.ENOTDIR},
		{addRoot("non-existing-dir"), errFilePathNotExists},
		{"", errFilePathEmpty},
		{"	", errFilePathEmpty},
	}
}

func TestStringDirPath(t *testing.T) {
	root := prepareFileSystemTests(t)
	for _, tc := range getStringDirPathTestCases(root) {
		err := StringDirPath().Validate(tc.in)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error())
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDirPath))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDirPath(b *testing.B) {
	root := prepareFileSystemTests(b)
	testCases := getStringDirPathTestCases(root)
	for _, tc := range testCases {
		rule := StringDirPath()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

// test cases copied from Go's [filepath] standard library.
var stringMatchFileSystemPathTestCases = []*struct {
	pattern, in string
	shouldFail  bool
}{
	// cspell:disable
	{"abc", "abc", false},
	{"*", "abc", false},
	{"*c", "abc", false},
	{"a*", "a", false},
	{"a*", "abc", false},
	{"a*/b", "abc/b", false},
	{"a*b*c*d*e*/f", "axbxcxdxe/f", false},
	{"a*b*c*d*e*/f", "axbxcxdxexxx/f", false},
	{"a*b?c*x", "abxbbxdbxebxczzx", false},
	{"ab[c]", "abc", false},
	{"ab[b-d]", "abc", false},
	{"ab[^e-g]", "abc", false},
	{"a\\*b", "a*b", false},
	{"a?b", "a☺b", false},
	{"a[^a]b", "a☺b", false},
	{"[a-ζ]*", "α", false},
	{"[\\]a]", "]", false},
	{"[\\-]", "-", false},
	{"*x", "xxx", false},
	{"[x\\-]", "x", false},
	{"[x\\-]", "-", false},
	{"[\\-x]", "x", false},
	{"[\\-x]", "-", false},
	{"a*/b", "a/c/b", true},
	{"ab[e-g]", "abc", true},
	{"ab[^c]", "abc", true},
	{"a*", "ab/c", true},
	{"a*b*c*d*e*/f", "axbxcxdxe/xxx/f", true},
	{"a*b*c*d*e*/f", "axbxcxdxexxx/fff", true},
	{"a*b?c*x", "abxbbxdbxebxczzy", true},
	{"ab[^b-d]", "abc", true},
	{"a???b", "a☺b", true},
	{"a\\*b", "ab", true},
	{"a[^a][^a][^a]b", "a☺b", true},
	{"*[a-ζ]", "A", true},
	{"a?b", "a/b", true},
	{"a*b", "a/b", true},
	{"[x\\-]", "z", true},
	{"[\\-x]", "a", true},
	{"[]a]", "]", true},
	{"[-]", "-", true},
	{"[x-]", "x", true},
	{"[x-]", "-", true},
	{"[x-]", "z", true},
	{"[-x]", "x", true},
	{"[-x]", "-", true},
	{"[-x]", "a", true},
	{"\\", "a", true},
	{"[a-b-c]", "a", true},
	{"[", "a", true},
	{"[^", "a", true},
	{"[^bc", "a", true},
	{"a[", "a", true},
	{"a[", "ab", true},
	{"a[", "x", true},
	{"a/b[", "x", true},
	// cspell:enable
}

func TestStringMatchFileSystemPath(t *testing.T) {
	for _, tc := range stringMatchFileSystemPathTestCases {
		err := StringMatchFileSystemPath(tc.pattern).Validate(tc.in)
		if tc.shouldFail {
			if !strings.Contains(err.Error(), "string must match file path pattern") &&
				!strings.Contains(err.Error(), filepath.ErrBadPattern.Error()) {
				assert.Fail(t, "unexpected error: %v", err)
			}
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMatchFileSystemPath))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringMatchFileSystemPath(b *testing.B) {
	for _, tc := range stringMatchFileSystemPathTestCases {
		rule := StringMatchFileSystemPath(tc.pattern)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

// test cases copied from Go's [regexp] standard library.
var stringRegexpTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{``, false},
	{`.`, false},
	{`^.$`, false},
	{`a`, false},
	{`a*`, false},
	{`a+`, false},
	{`a?`, false},
	{`a|b`, false},
	{`a*|b*`, false},
	{`(a*|b)(c*|d)`, false},
	{`[a-z]`, false},
	{`[a-abc-c\-\]\[]`, false},
	{`[a-z]+`, false},
	{`[abc]`, false},
	{`[^1234]`, false},
	{`[^\n]`, false},
	{`\!\\`, false},
	{`*`, true},
	{`+`, true},
	{`?`, true},
	{`(abc`, true},
	{`abc)`, true},
	{`x[a-z`, true},
	{`[z-a]`, true},
	{`abc\`, true},
	{`a**`, true},
	{`a*+`, true},
	{`\x`, true},
	{strings.Repeat(`\pL`, 27000), true},
	// cspell:enable
}

func TestStringRegexp(t *testing.T) {
	for _, tc := range stringRegexpTestCases {
		err := StringRegexp().Validate(tc.in)
		if tc.shouldFail {
			assert.ErrorContains(t, err, "string must be a valid regular expression")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringRegexp))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringRegexp(b *testing.B) {
	for _, tc := range stringRegexpTestCases {
		rule := StringRegexp()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

type stringCrontabTestCase struct {
	in         string
	shouldFail bool
}

func getStringCronTestCases() []*stringCrontabTestCase {
	testCases := []*stringCrontabTestCase{
		{"@annually", false},
		{"@yearly", false},
		{"@monthly", false},
		{"@weekly", false},
		{"@daily", false},
		{"@hourly", false},
		{"@reboot", false},
		{"* * * * *", false},
		{"* * * JAN,MAY,DEC *", false},
		{"* * * JAN-DEC *", false},
		{"* * * FEB-MAY/2 *", false},
		{"* * * fEb-may/10 *", false},
		{"* * * SEP-SEP/2 *", false},
		{"* * * JAN-1 *", false},
		{"* * * JAN-12 *", false},
		{"* * * 1-DEC *", false},
		{"* * * * FRI-7", false},
		{"* * * * 2-WED", false},
		{"* * * * THU-FRI", false},
		{"* * * * TUE-THU/10", false},
		{"* * * * SUN-MON", false},
		{"* * * * WED-3", false},
		{"* * * * THU,FRI,MON", false},
		{"* * * * *", false},
		{"", true},
		{"  @hourly", true},
		{"1h @every", true},
		{"@every 1Y", true},
		{"wrong", true},
		{"@minutely", true},
		{"@every 1h", true},
		{"@every 1h30m10ts", true},
		{"a * * * *", true},
		{"1 b * * *", true},
		{"1 1 c * *", true},
		{"1 1 1 d *", true},
		{"1 1 1 1 e", true},
		{"* * * MAZ *", true},
		{"* * * MAY-FEB/2 *", true},
		{"* * * MAY-2 *", true},
		{"* * * 2-JAN *", true},
		{"* * * FEB-JUN/-10 *", true},
		{"* * * JAN,MAY,DEZ *", true},
		{"* * * * MOZ", true},
		{"* * * * MON-SUN", true},
		{"* * * * 7-FRI", true},
		{"* * * * WED-2", true},
		{"* * * * MON-FRI/-10", true},
		{"* * * * THU,FRI,MOZ", true},
	}
	createCron := func(n int, format string, a ...any) string {
		fields := strings.Fields("* * * * *")
		fields[n] = fmt.Sprintf(format, a...)
		return strings.Join(fields, " ")
	}
	for _, field := range []struct {
		n, lower, upper int
	}{
		{0, 0, 59},
		{1, 0, 23},
		{2, 1, 31},
		{3, 1, 12},
		{4, 0, 7},
	} {
		getRandom := func() int {
			return field.lower + rand.Intn(field.upper-field.lower)
		}
		testCases = append(testCases,
			&stringCrontabTestCase{createCron(field.n, "%d", getRandom()), false},
			&stringCrontabTestCase{createCron(field.n, "%d", field.lower), false},
			&stringCrontabTestCase{createCron(field.n, "%d", field.upper), false},
			&stringCrontabTestCase{createCron(field.n, "%d,%d", field.lower, field.upper), false},
			&stringCrontabTestCase{createCron(field.n, "%d,%d", field.upper, field.lower), false},
			&stringCrontabTestCase{createCron(field.n, "%d-%d", field.lower, field.upper), false},
			&stringCrontabTestCase{createCron(field.n, "%d-%d/10", field.lower, field.upper), false},
			&stringCrontabTestCase{createCron(field.n, "*/10"), false},
			&stringCrontabTestCase{createCron(field.n, "%d", field.lower-1), true},
			&stringCrontabTestCase{createCron(field.n, "%d", field.upper+1), true},
			&stringCrontabTestCase{createCron(field.n, "%d,", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "%d,%d", field.lower, field.upper+1), true},
			&stringCrontabTestCase{createCron(field.n, "%d,%d", field.lower-1, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d/10", getRandom()), true},
			&stringCrontabTestCase{createCron(field.n, "%d,%d/10", field.lower, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "a"), true},
			&stringCrontabTestCase{createCron(field.n, "%d,a", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "a,%d", field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d-", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "%d-/", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "-/"), true},
			&stringCrontabTestCase{createCron(field.n, "%d-%d/", field.lower, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d-%d/a", field.lower, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d-%d/-10", field.lower, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d-*/10", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "*-*/10"), true},
			&stringCrontabTestCase{createCron(field.n, "*-%d/10", field.upper), true},
		)
	}
	for month := range crontabMonthsMap {
		testCases = append(testCases, &stringCrontabTestCase{createCron(3, "%s", month), false})
	}
	for day := range crontabDaysMap {
		// Skip special cases for Sunday.
		if strings.Contains(day, "-") {
			continue
		}
		testCases = append(testCases, &stringCrontabTestCase{createCron(4, "%s", day), false})
	}
	return testCases
}

func TestStringCrontab(t *testing.T) {
	for _, tc := range getStringCronTestCases() {
		t.Run(tc.in, func(t *testing.T) {
			err := StringCrontab().Validate(tc.in)
			if tc.shouldFail {
				assert.ErrorContains(t, err, "string must be a valid cron schedule expression")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCrontab))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringCrontab(b *testing.B) {
	for range b.N {
		testCases := getStringCronTestCases()
		for _, tc := range testCases {
			_ = StringCrontab().Validate(tc.in)
		}
	}
}

var stringDateTimeTestCases = []*struct {
	layout string
	in     string
	errMsg string
}{
	{time.RFC3339, "2024-01-01T15:00:00Z", ""},
	{time.RFC3339, "2024-01-01T15:00:00+01:00", ""},
	{time.DateTime, "2024-01-01 15:00:00", ""},
	{time.DateOnly, "2024-01-01", ""},
	{time.TimeOnly, "15:00:00", ""},
	{
		"invalid-layout",
		"2024-01-01T15:00:00Z",
		"string must be a valid date and time in 'invalid-layout' format",
	},
	{
		time.RFC3339,
		"2024-01-01 15:00:00Z",
		"string must be a valid date and time in '2006-01-02T15:04:05Z07:00' format",
	},
	{
		"15:04",
		"15:00:00",
		"string must be a valid date and time in '15:04'",
	},
}

func TestStringDateTime(t *testing.T) {
	for _, tc := range stringDateTimeTestCases {
		t.Run(tc.layout+tc.in, func(t *testing.T) {
			err := StringDateTime(tc.layout).Validate(tc.in)
			if tc.errMsg != "" {
				assert.ErrorContains(t, err, tc.errMsg)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDateTime))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringDateTime(b *testing.B) {
	for _, tc := range stringDateTimeTestCases {
		rule := StringDateTime(tc.layout)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringTimeZoneTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"UTC", false},
	{"America/New_York", false},
	{"Europe/Warsaw", false},
	{"", true},
	{"Local", true},
	{"America/New_Yorker", true},
	{"x/x", true},
	{"America/Warsaw", true},
}

func TestStringTimeZone(t *testing.T) {
	for _, tc := range stringTimeZoneTestCases {
		t.Run(tc.in, func(t *testing.T) {
			err := StringTimeZone().Validate(tc.in)
			if tc.shouldFail {
				assert.ErrorContains(t, err, "string must be a valid IANA Time Zone Database code")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringTimeZone))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringTimeZone(b *testing.B) {
	for _, tc := range stringDateTimeTestCases {
		rule := StringTimeZone()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringAlphaTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test", false},
	{"tEsT", false},
	{"s", false},
	{"LOL", false},
	{"test-this", true},
	{" test", true},
	{"  ", true},
	{" ", true},
	{"test1", true},
	{"tęst", true},
	// cspell:enable
}

func TestStringAlpha(t *testing.T) {
	for _, tc := range stringAlphaTestCases {
		err := StringAlpha().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAlpha))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringAlpha(b *testing.B) {
	for _, tc := range stringAlphaTestCases {
		rule := StringAlpha()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringAlphanumericTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test", false},
	{"tEsT", false},
	{"s", false},
	{"4", false},
	{"LOL", false},
	{"test1", false},
	{"-921", true},
	{"test-this", true},
	{" test", true},
	{" 1", true},
	{"  ", true},
	{" ", true},
	{"tęst", true},
	{"tęst1", true},
	// cspell:enable
}

func TestStringAlphanumeric(t *testing.T) {
	for _, tc := range stringAlphanumericTestCases {
		err := StringAlphanumeric().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAlphanumeric))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringAlphanumeric(b *testing.B) {
	for _, tc := range stringAlphanumericTestCases {
		rule := StringAlphanumeric()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringAlphaUnicodeTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test", false},
	{"tEsT", false},
	{"s", false},
	{"LOL", false},
	{"tęst", false},
	{"汉语", false},
	{"一二三", false},
	{"test-this", true},
	{" test", true},
	{"  ", true},
	{" ", true},
	{"test1", true},
	{"汉语!", true},
	{"1汉语", true},
	// cspell:enable
}

func TestStringAlphaUnicode(t *testing.T) {
	for _, tc := range stringAlphaUnicodeTestCases {
		err := StringAlphaUnicode().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAlphaUnicode))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringAlphaUnicode(b *testing.B) {
	for _, tc := range stringAlphaUnicodeTestCases {
		rule := StringAlphaUnicode()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringAlphanumericUnicodeTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test", false},
	{"tEsT", false},
	{"s", false},
	{"5", false},
	{"LOL", false},
	{"tęst", false},
	{"汉语", false},
	{"1汉语", false},
	{"test1", false},
	{"tęst1", false},
	{"一二三", false},
	{"-550", true},
	{"test-this", true},
	{" test", true},
	{"  ", true},
	{" ", true},
	{"汉语!", true},
	{"-921", true},
	{" 1", true},
	// cspell:enable
}

func TestStringAlphanumericUnicode(t *testing.T) {
	for _, tc := range stringAlphanumericUnicodeTestCases {
		err := StringAlphanumericUnicode().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAlphanumericUnicode))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringAlphanumericUnicode(b *testing.B) {
	for _, tc := range stringAlphanumericUnicodeTestCases {
		rule := StringAlphanumericUnicode()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringFQDNTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test.example.com", false},
	{"example.com", false},
	{"example24.com", false},
	{"test.example24.com", false},
	{"test24.example24.com", false},
	{"test.example.com.", false},
	{"example.com.", false},
	{"example24.com.", false},
	{"test.example24.com.", false},
	{"test24.example24.com.", false},
	{"24.example24.com", false},
	{"test.24.example.com", false},
	{"test24.example24.com..", true},
	{"example", true},
	{"192.168.0.1", true},
	{"email@example.com", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652", true},
	{"2001:cdba:0:0:0:0:3257:9652", true},
	{"2001:cdba::3257:9652", true},
	{"", true},
	// cspell:enable
}

func TestStringFQDN(t *testing.T) {
	for _, tc := range stringFQDNTestCases {
		err := StringFQDN().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringFQDN))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringFQDN(b *testing.B) {
	for _, tc := range stringFQDNTestCases {
		rule := StringFQDN()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var (
	errK8sQualifiedNameEmptyPrefixPart = errors.New("prefix part must not be empty")
	errK8sQualifiedNamePrefixLength    = errors.New("prefix part length must be less than or equal to 253")
	errK8sQualifiedNamePrefixRegexp    = errors.New(
		`prefix part string must match regular expression: '^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$'`,
	)
	errK8sQualifiedNameTooManyParts   = errors.New("qualified name must have at most 2 parts separated by a '/'")
	errK8sQualifiedNameEmptyNamePart  = errors.New("name part must not be empty")
	errK8sQualifiedNameNamePartLength = errors.New("name part length must be less than or equal to 63")
	errK8sQualifiedNameNamePartRegexp = errors.New(
		"name part string must match regular expression: '^([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$'",
	)
)

var stringK8sQualifiedNameTestCases = []*struct {
	in          string
	expectedErr error
}{
	// cspell:disable
	{"simple", nil},
	{"now-with-dashes", nil},
	{"1-starts-with-num", nil},
	{"1234", nil},
	{"simple/simple", nil},
	{"now-with-dashes/simple", nil},
	{"now-with-dashes/now-with-dashes", nil},
	{"now.with.dots/simple", nil},
	{"now-with.dashes-and.dots/simple", nil},
	{"1-num.2-num/3-num", nil},
	{"1234/5678", nil},
	{"1.2.3.4/5678", nil},
	{"Uppercase_Is_OK_123", nil},
	{"example.com/Uppercase_Is_OK_123", nil},
	{"requests.storage-foo", nil},
	{strings.Repeat("a", 63), nil},
	{strings.Repeat("a", 253) + "/" + strings.Repeat("b", 63), nil},
	// BAD
	{"/", errK8sQualifiedNameEmptyPrefixPart},
	{"nospecialchars%^=@", errK8sQualifiedNameNamePartRegexp},
	{"cantendwithadash-", errK8sQualifiedNameNamePartRegexp},
	{"-cantstartwithadash-", errK8sQualifiedNameNamePartRegexp},
	{"example.com/abc$", errK8sQualifiedNameNamePartRegexp},
	{"only/one/slash", errK8sQualifiedNameTooManyParts},
	{"Example.com/abc", errK8sQualifiedNamePrefixRegexp},
	{"example_com/abc", errK8sQualifiedNamePrefixRegexp},
	{"example.com/", errK8sQualifiedNameEmptyNamePart},
	{"/simple", errK8sQualifiedNameEmptyPrefixPart},
	{"not.Valid/simple", errK8sQualifiedNamePrefixRegexp},
	{strings.Repeat("a", 64), errK8sQualifiedNameNamePartLength},
	{strings.Repeat("a", 254) + "/abc", errK8sQualifiedNamePrefixLength},
	{strings.Repeat("a", 253) + "/" + strings.Repeat("b", 64), errors.New("length must be between 1 and 317")},
	// cspell:enable
}

func TestStringKubernetesQualifiedName(t *testing.T) {
	for _, tc := range stringK8sQualifiedNameTestCases {
		t.Run(tc.in, func(t *testing.T) {
			err := StringKubernetesQualifiedName().Validate(tc.in)
			if tc.expectedErr != nil {
				assert.ErrorContains(t, err, tc.expectedErr.Error())
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringKubernetesQualifiedName))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringKubernetesQualifiedName(b *testing.B) {
	for _, tc := range stringK8sQualifiedNameTestCases {
		rule := StringKubernetesQualifiedName()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func testStringFormatIDRule(
	t *testing.T,
	rule govy.Rule[string],
	errorCode govy.ErrorCode,
	expectedError string,
	validInputs map[string]string,
	invalidInputs map[string]string,
) {
	t.Helper()
	t.Run("valid", func(t *testing.T) {
		for name, input := range validInputs {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid", func(t *testing.T) {
		for name, input := range invalidInputs {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(t, err, expectedError)
				assert.True(t, govy.HasErrorCode(err, errorCode))
			})
		}
	})
}

func benchmarkStringFormatIDRule(
	b *testing.B,
	rule govy.Rule[string],
	validInputs map[string]string,
	invalidInputs map[string]string,
) {
	b.Helper()
	for b.Loop() {
		for _, input := range validInputs {
			_ = rule.Validate(input)
		}
		for _, input := range invalidInputs {
			_ = rule.Validate(input)
		}
	}
}
