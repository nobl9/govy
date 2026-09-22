package rules

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"maps"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
)

const (
	uuidRFC4122Pattern = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89aAbB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
	uuidv3Pattern      = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[89aAbB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
	uuidv4Pattern      = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89aAbB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
	uuidv5Pattern      = `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89aAbB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`
	ulidPattern        = `^[0-7][0-9A-HJKMNP-TV-Za-hjkmnp-tv-z]{25}$`
)

var stringNotEmptyTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"                s", false},
	{"     ", true},
	{"", true},
	{"\t\n\v\f\r", true},
	{"\u0085", true},
	{"\u00a0\u1680\u2000\u200a\u2028\u2029\u202f\u205f\u3000", true},
	{"\u200b", false},
	{"\ufeff", false},
	{"\u0085s\u00a0", false},
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

func TestStringNotEmpty_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringNotEmpty())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_not_empty.json",
		stringBooleanJSONSchemaCases(stringNotEmptyTestCases))
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
			in: "prefix absuffix",
		},
		{
			in: "\na\n",
		},
		{
			in:            "",
			expectedError: "string must match regular expression: '[ab]+'",
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

func TestStringMatchRegexp_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).
		Rules(StringMatchRegexp(stringMatchRegexpRegexp))))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_match_regexp.json",
		stringRegexpJSONSchemaCases(stringMatchRegexpTestCases))
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
			in: "",
		},
		{
			in:            "prefix absuffix",
			expectedError: "string must not match regular expression: '[ab]+'",
		},
		{
			in:            "\na\n",
			expectedError: "string must not match regular expression: '[ab]+'",
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

func TestStringDenyRegexp_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).
		Rules(StringDenyRegexp(stringDenyRegexpRegexp))))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_deny_regexp.json",
		stringRegexpJSONSchemaCases(stringDenyRegexpTestCases))
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
	{"test\n", true},
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

func TestStringDNSLabel_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringDNSLabel())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_dns_label.json",
		stringBooleanJSONSchemaCases(stringDNSLabelTestCases))
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
	{"example.com\n", true},
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

func TestStringDNSSubdomain_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringDNSSubdomain())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_dns_subdomain.json",
		stringBooleanJSONSchemaCases(stringDNSSubdomainTestCases))
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

func TestStringUUID_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringUUID())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_uuid.json",
		stringNamedJSONSchemaCases(stringUUIDValidInputs, stringUUIDInvalidInputs),
	)
}

func BenchmarkStringUUID(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUID(), stringUUIDValidInputs, stringUUIDInvalidInputs)
}

func BenchmarkUUIDPredicate(b *testing.B) {
	benchmarkStringFormatIDPredicate(b, isValidUUID, stringUUIDValidInputs, stringUUIDInvalidInputs)
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

func TestStringUUIDRFC4122_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringUUIDRFC4122())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_uuid_rfc4122.json",
		stringNamedJSONSchemaCases(stringUUIDRFC4122ValidInputs, stringUUIDRFC4122InvalidInputs),
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

func BenchmarkUUIDRFC4122Predicate(b *testing.B) {
	benchmarkStringFormatIDPredicate(
		b,
		isValidUUIDRFC4122,
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

func TestStringUUIDv3_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringUUIDv3())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_uuid_v3.json",
		stringNamedJSONSchemaCases(stringUUIDv3ValidInputs, stringUUIDv3InvalidInputs),
	)
}

func BenchmarkStringUUIDv3(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv3(), stringUUIDv3ValidInputs, stringUUIDv3InvalidInputs)
}

func BenchmarkUUIDv3Predicate(b *testing.B) {
	benchmarkStringFormatIDPredicate(
		b,
		func(s string) bool { return isValidUUIDVersion(s, '3') },
		stringUUIDv3ValidInputs,
		stringUUIDv3InvalidInputs,
	)
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

func TestStringUUIDv4_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringUUIDv4())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_uuid_v4.json",
		stringNamedJSONSchemaCases(stringUUIDv4ValidInputs, stringUUIDv4InvalidInputs),
	)
}

func BenchmarkStringUUIDv4(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv4(), stringUUIDv4ValidInputs, stringUUIDv4InvalidInputs)
}

func BenchmarkUUIDv4Predicate(b *testing.B) {
	benchmarkStringFormatIDPredicate(
		b,
		func(s string) bool { return isValidUUIDVersion(s, '4') },
		stringUUIDv4ValidInputs,
		stringUUIDv4InvalidInputs,
	)
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

func TestStringUUIDv5_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringUUIDv5())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_uuid_v5.json",
		stringNamedJSONSchemaCases(stringUUIDv5ValidInputs, stringUUIDv5InvalidInputs),
	)
}

func BenchmarkStringUUIDv5(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv5(), stringUUIDv5ValidInputs, stringUUIDv5InvalidInputs)
}

func BenchmarkUUIDv5Predicate(b *testing.B) {
	benchmarkStringFormatIDPredicate(
		b,
		func(s string) bool { return isValidUUIDVersion(s, '5') },
		stringUUIDv5ValidInputs,
		stringUUIDv5InvalidInputs,
	)
}

// stringULIDValidInputs includes every concrete valid ULID from ulid/spec at
// revision d0c7170df4517939e70129b4d6462cc162f2d5bf and every concrete ULID
// from ulid/javascript's tests at revision 11c2067821ee19e4dc787ca4e0125a025485edc6.
//
// The specification's `ttttttttttrrrrrrrrrrrrrrrr` representation describes
// the field layout; it is not a concrete ULID and is intentionally excluded.
var stringULIDValidInputs = map[string]string{
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

func TestStringULID_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringULID())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_ulid.json",
		stringNamedJSONSchemaCases(stringULIDValidInputs, stringULIDInvalidInputs),
	)
}

func BenchmarkStringULID(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringULID(), stringULIDValidInputs, stringULIDInvalidInputs)
}

func BenchmarkULIDPredicate(b *testing.B) {
	benchmarkStringFormatIDPredicate(b, isValidULID, stringULIDValidInputs, stringULIDInvalidInputs)
}

func TestStringFormatIDPredicatesMatchRegularExpressions(t *testing.T) {
	tests := map[string]struct {
		pattern       string
		seed          string
		predicate     func(string) bool
		validInputs   map[string]string
		invalidInputs map[string]string
	}{
		"UUID": {
			pattern:       uuidPattern,
			seed:          "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
			predicate:     isValidUUID,
			validInputs:   stringUUIDValidInputs,
			invalidInputs: stringUUIDInvalidInputs,
		},
		"RFC 4122 UUID": {
			pattern:       uuidRFC4122Pattern,
			seed:          "f81d4fae-7dec-11d0-a765-00a0c91e6bf6",
			predicate:     isValidUUIDRFC4122,
			validInputs:   stringUUIDRFC4122ValidInputs,
			invalidInputs: stringUUIDRFC4122InvalidInputs,
		},
		"UUIDv3": {
			pattern:       uuidv3Pattern,
			seed:          "5df41881-3aed-3515-88a7-2f4a814cf09e",
			predicate:     func(s string) bool { return isValidUUIDVersion(s, '3') },
			validInputs:   stringUUIDv3ValidInputs,
			invalidInputs: stringUUIDv3InvalidInputs,
		},
		"UUIDv4": {
			pattern:       uuidv4Pattern,
			seed:          "919108f7-52d1-4320-9bac-f847db4148a8",
			predicate:     func(s string) bool { return isValidUUIDVersion(s, '4') },
			validInputs:   stringUUIDv4ValidInputs,
			invalidInputs: stringUUIDv4InvalidInputs,
		},
		"UUIDv5": {
			pattern:       uuidv5Pattern,
			seed:          "2ed6657d-e927-568b-95e1-2665a8aea6a2",
			predicate:     func(s string) bool { return isValidUUIDVersion(s, '5') },
			validInputs:   stringUUIDv5ValidInputs,
			invalidInputs: stringUUIDv5InvalidInputs,
		},
		"ULID": {
			pattern:       ulidPattern,
			seed:          "01ARZ3NDEKTSV4RRFFQ69G5FAV",
			predicate:     isValidULID,
			validInputs:   stringULIDValidInputs,
			invalidInputs: stringULIDInvalidInputs,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			re := regexp.MustCompile(tc.pattern)
			for _, input := range tc.validInputs {
				assertStringPredicateMatchesRegexp(t, re, tc.predicate, input)
			}
			for _, input := range tc.invalidInputs {
				assertStringPredicateMatchesRegexp(t, re, tc.predicate, input)
			}

			input := []byte(tc.seed)
			for position := range input {
				original := input[position]
				for value := range 256 {
					input[position] = byte(value)
					assertStringPredicateMatchesRegexp(t, re, tc.predicate, string(input))
				}
				input[position] = original
			}
			for position := range len(tc.seed) {
				input := tc.seed[:position] + tc.seed[position+1:]
				assertStringPredicateMatchesRegexp(t, re, tc.predicate, input)
			}
			for position := range len(tc.seed) + 1 {
				input := tc.seed[:position] + "\x00" + tc.seed[position:]
				assertStringPredicateMatchesRegexp(t, re, tc.predicate, input)
			}
		})
	}
}

var stringASCIITestCases = []*struct {
	in         string
	shouldFail bool
}{
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

func TestStringASCII_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringASCII())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_ascii.json",
		stringBooleanJSONSchemaCases(stringASCIITestCases))
}

func BenchmarkStringASCII(b *testing.B) {
	for _, tc := range stringASCIITestCases {
		rule := StringASCII()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringCreditCardTestCases = []jsonschematest.Case[string]{
	{Name: "visa", Input: "4111111111111111", Valid: true},
	{Name: "visa alternate", Input: "4242424242424242", Valid: true},
	{Name: "mastercard", Input: "5555555555554444", Valid: true},
	{Name: "american express", Input: "378282246310005", Valid: true},
	{Name: "discover", Input: "6011111111111117", Valid: true},
	{Name: "six-series sixteen digits", Input: "6123451234567893", Valid: true},
	{Name: "diners club", Input: "36227206271667", Valid: true},
	{Name: "mastercard 2-series", Input: "2223003122003222", Valid: true},
	{Name: "JCB", Input: "3566002020360505", Valid: true},
	{Name: "UnionPay nineteen digits", Input: "6205500000000000004", Valid: true},
	{Name: "visa decline test number", Input: "4000111111111115", Valid: true},
	{Name: "minimum thirteen digits", Input: "1000000000009", Valid: true},
	{Name: "maximum nineteen digits", Input: "1000000000000000009", Valid: true},
	{Name: "empty", Input: ""},
	{Name: "minimum length minus one", Input: "100000000008"},
	{Name: "maximum length plus one", Input: "10000000000000000008"},
	{
		Name: "failed thirteen digit checksum", Input: "1000000000008",
		JSONSchemaDifference: "JSON Schema checks payment-card syntax but does not validate the Luhn checksum.",
	},
	{
		Name: "failed nineteen digit checksum", Input: "1000000000000000008",
		JSONSchemaDifference: "JSON Schema checks payment-card syntax but does not validate the Luhn checksum.",
	},
	{
		Name: "all same digits", Input: "6666666666666",
		JSONSchemaDifference: "JSON Schema does not reject payment-card numbers with identical digits.",
	},
	{Name: "leading space", Input: " 1000000000009"},
	{Name: "trailing space", Input: "1000000000009 "},
	{Name: "trailing newline", Input: "1000000000009\n"},
	{Name: "embedded spaces", Input: "4111 1111 1111 1111"},
	{Name: "hyphens", Input: "4111-1111-1111-1111"},
	{Name: "full-width digit", Input: "10000000000０9"},
	{Name: "alphabetic character", Input: "10000000000A9"},
}

func TestStringCreditCard(t *testing.T) {
	assertPaymentBankingRule(
		t,
		StringCreditCard(),
		stringCreditCardTestCases,
		"string must be a valid payment card number",
		ErrorCodeStringCreditCard,
	)
}

func TestStringCreditCard_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringCreditCard())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_credit_card.json",
		slices.Concat(stringCreditCardTestCases, readPaymentCardProcessorTestCases(t)),
	)
}

func BenchmarkStringCreditCard(b *testing.B) {
	rule := StringCreditCard()
	benchmarkStringPaymentBankingRule(b, rule, stringCreditCardTestCases)
}

var stringLuhnChecksumTestCases = []jsonschematest.Case[string]{
	{Name: "single zero", Input: "0", Valid: true},
	{Name: "four zeroes", Input: "0000", Valid: true},
	{Name: "twenty zeroes", Input: "00000000000000000000", Valid: true},
	{Name: "two digits ending in eight", Input: "18", Valid: true},
	{Name: "two digits ending in nine", Input: "59", Valid: true},
	{Name: "sample number", Input: "79927398713", Valid: true},
	{Name: "six-series example", Input: "6123451234567893", Valid: true},
	{Name: "fifteen digit example", Input: "808401234567893", Valid: true},
	{Name: "nineteen digit example", Input: "6205500000000000004", Valid: true},
	{Name: "empty", Input: ""},
	{
		Name: "failed thirteen digit check", Input: "6123451234567894",
		JSONSchemaDifference: "JSON Schema requires digits but does not validate the Luhn checksum.",
	},
	{
		Name: "failed fifteen digit check", Input: "808401234567894",
		JSONSchemaDifference: "JSON Schema requires digits but does not validate the Luhn checksum.",
	},
	{Name: "slash before zero", Input: "/0"},
	{Name: "colon before zero", Input: ":0"},
	{Name: "plus before zero", Input: "+0"},
	{Name: "trailing space", Input: "0 "},
	{Name: "alphabetic character", Input: "79927A398713"},
	{Name: "full-width digits", Input: "１２"},
}

func TestStringLuhnChecksum(t *testing.T) {
	assertPaymentBankingRule(
		t,
		StringLuhnChecksum(),
		stringLuhnChecksumTestCases,
		"string must pass the Luhn checksum",
		ErrorCodeStringLuhnChecksum,
	)
}

func TestStringLuhnChecksum_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringLuhnChecksum())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_luhn_checksum.json",
		slices.Concat(stringLuhnChecksumTestCases, readPaymentCardProcessorTestCases(t)),
	)
}

func BenchmarkStringLuhnChecksum(b *testing.B) {
	rule := StringLuhnChecksum()
	benchmarkStringPaymentBankingRule(b, rule, stringLuhnChecksumTestCases)
}

func TestStringPaymentCardProcessorFixtures(t *testing.T) {
	rules := []struct {
		name          string
		rule          govy.Rule[string]
		expectedError string
		errorCode     govy.ErrorCode
	}{
		{
			name:          "StringCreditCard",
			rule:          StringCreditCard(),
			expectedError: "string must be a valid payment card number",
			errorCode:     ErrorCodeStringCreditCard,
		},
		{
			name:          "StringLuhnChecksum",
			rule:          StringLuhnChecksum(),
			expectedError: "string must pass the Luhn checksum",
			errorCode:     ErrorCodeStringLuhnChecksum,
		},
	}
	for _, tc := range readPaymentCardProcessorTestCases(t) {
		t.Run(tc.Name, func(t *testing.T) {
			for _, testRule := range rules {
				t.Run(testRule.name, func(t *testing.T) {
					err := testRule.rule.Validate(tc.Input)
					if !tc.Valid {
						assertPaymentBankingRuleError(t, err, testRule.expectedError, testRule.errorCode)
						return
					}
					assert.NoError(t, err)
				})
			}
		})
	}
}

func readPaymentCardProcessorTestCases(t *testing.T) []jsonschematest.Case[string] {
	t.Helper()
	fixtures := []struct {
		name                   string
		path                   string
		source                 string
		sourceSnapshotSHA256   string
		normalizedValuesSHA256 string
		expectedCount          int
	}{
		{
			name:                   "Stripe",
			path:                   "testdata/stripe_test_card_numbers_2026-07-21.txt",
			source:                 "https://docs.stripe.com/testing",
			sourceSnapshotSHA256:   "bcc37c3f58146fcadf2290b67058d66841b9adfbf207a56ae0bef557c675fc4e",
			normalizedValuesSHA256: "00176aa47882a48d54cc1db7cb654536cd8889406d59ce67c88ae00897e54da9",
			expectedCount:          148,
		},
		{
			name:                   "Braintree",
			path:                   "testdata/braintree_test_card_numbers_2026-07-21.txt",
			source:                 "https://developer.paypal.com/braintree/docs/reference/general/testing/php",
			sourceSnapshotSHA256:   "85231b65e2b6fa674d7ada4cd71512d0fe1756398e83564244af1bbf3cb27211",
			normalizedValuesSHA256: "3ff5fb628215867a24156654d3a026a86bc5d2a089846b8de129ff444c0c1db7",
			expectedCount:          44,
		},
		{
			name:                   "Adyen",
			path:                   "testdata/adyen_test_card_numbers_2026-07-21.txt",
			source:                 "https://docs.adyen.com/development-resources/test-cards-and-credentials/test-card-numbers",
			sourceSnapshotSHA256:   "be12e0c40e9d3cabbfe1025d956c13dbd4250dbc6574c23a7a3178ce06eff82e",
			normalizedValuesSHA256: "636787e0d1b9af8aa73287d6e7e058b9f6f499f47756537dbdf9f322c44f23da",
			expectedCount:          86,
		},
	}
	invalidProcessorCardNumber := jsonschematest.Case[string]{
		Input:                "4242424242424241",
		JSONSchemaDifference: "JSON Schema checks payment-card syntax but does not validate the Luhn checksum.",
	}
	const expectedUniqueProcessorCardNumbers = 264
	var cases []jsonschematest.Case[string]
	union := make(map[string]struct{}, expectedUniqueProcessorCardNumbers)
	invalidOccurrences := 0
	for _, fixture := range fixtures {
		rawFixture, inputs := readTestDataFields(t, fixture.path)
		assert.Require(t, assert.Len(t, inputs, fixture.expectedCount))
		assert.True(t, strings.Contains(rawFixture, "# Source: "+fixture.source+"\n"))
		assert.True(t, strings.Contains(
			rawFixture,
			"# Source snapshot SHA-256: "+fixture.sourceSnapshotSHA256+"\n",
		))
		normalizedValues := strings.Join(inputs, "\n") + "\n"
		actualValuesSHA256 := fmt.Sprintf("%x", sha256.Sum256([]byte(normalizedValues)))
		assert.Equal(t, fixture.normalizedValuesSHA256, actualValuesSHA256)

		uniqueInputs := make(map[string]struct{}, len(inputs))
		for _, input := range inputs {
			uniqueInputs[input] = struct{}{}
			union[input] = struct{}{}
			tc := jsonschematest.Case[string]{Input: input, Valid: true}
			if input == invalidProcessorCardNumber.Input {
				invalidOccurrences++
				tc = invalidProcessorCardNumber
			}
			tc.Name = fixture.name + "/" + input
			cases = append(cases, tc)
		}
		assert.Len(t, uniqueInputs, len(inputs))
	}
	assert.Len(t, union, expectedUniqueProcessorCardNumbers)
	assert.Equal(t, 1, invalidOccurrences)
	return cases
}

var stringBICTestCases = []jsonschematest.Case[string]{
	{Name: "eight characters", Input: "DEUTDEFF", Valid: true},
	{Name: "eleven characters", Input: "DEUTDEFF500", Valid: true},
	{Name: "branch placeholder", Input: "NEDSZAJJXXX", Valid: true},
	{Name: "ISO 9362 eight character example", Input: "ABCDFRPP", Valid: true},
	{Name: "ISO 9362 eleven character example", Input: "WG11US335AB", Valid: true},
	{Name: "TC68 eight character example", Input: "ABCDBE22", Valid: true},
	{Name: "TC68 eleven character example", Input: "ABCDBE22XYZ", Valid: true},
	{Name: "alphanumeric party prefix", Input: "A1B2US33XXX", Valid: true},
	{Name: "zero in party suffix", Input: "ABCDUS0A", Valid: true},
	{Name: "one in party suffix", Input: "ABCDUS1A", Valid: true},
	{Name: "letter O in party suffix", Input: "ABCDUSAO", Valid: true},
	{Name: "Kosovo bank example", Input: "EKOMXKPR", Valid: true},
	{Name: "Kosovo bank identifier", Input: "BPBXXKPR", Valid: true},
	{Name: "empty", Input: ""},
	{Name: "lowercase", Input: "abcdfrpp"},
	{Name: "digit in country code", Input: "ABCD3R22"},
	{
		Name: "user-assigned country AA", Input: "ABCDAA22",
		JSONSchemaDifference: "JSON Schema checks BIC syntax but does not enforce the country-code allowlist.",
	},
	{
		Name: "user-assigned country QM", Input: "ABCDQM22",
		JSONSchemaDifference: "JSON Schema checks BIC syntax but does not enforce the country-code allowlist.",
	},
	{
		Name: "user-assigned country QZ", Input: "ABCDQZ22",
		JSONSchemaDifference: "JSON Schema checks BIC syntax but does not enforce the country-code allowlist.",
	},
	{
		Name: "user-assigned country XA", Input: "ABCDXA22",
		JSONSchemaDifference: "JSON Schema checks BIC syntax but does not enforce the country-code allowlist.",
	},
	{
		Name: "user-assigned country XZ", Input: "ABCDXZ22",
		JSONSchemaDifference: "JSON Schema checks BIC syntax but does not enforce the country-code allowlist.",
	},
	{
		Name: "user-assigned country ZZ", Input: "ABCDZZ22",
		JSONSchemaDifference: "JSON Schema checks BIC syntax but does not enforce the country-code allowlist.",
	},
	{
		Name: "deleted country AN", Input: "ABCDAN22",
		JSONSchemaDifference: "JSON Schema checks BIC syntax but does not enforce the country-code allowlist.",
	},
	{
		Name: "deleted country CS", Input: "ABCDCS22",
		JSONSchemaDifference: "JSON Schema checks BIC syntax but does not enforce the country-code allowlist.",
	},
	{Name: "seven characters", Input: "ABCDFR2"},
	{Name: "nine characters", Input: "ABCDFR22X"},
	{Name: "ten characters", Input: "ABCDFR22XX"},
	{Name: "twelve characters", Input: "ABCDFR22XXXX"},
	{Name: "punctuation in party prefix", Input: "ABC-FR22"},
	{Name: "punctuation in party suffix", Input: "ABCDFR2-"},
	{Name: "punctuation in branch", Input: "ABCDFR22XY-"},
	{Name: "trailing space", Input: "ABCDFR22 "},
}

func TestStringBIC(t *testing.T) {
	assertPaymentBankingRule(
		t,
		StringBIC(),
		stringBICTestCases,
		"string must be a valid Business Identifier Code (BIC)",
		ErrorCodeStringBIC,
	)
}

func TestStringBIC_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := slices.Clone(stringBICTestCases)
	for _, countryCode := range append(readISOAlpha2CountryCodes(t), "XK") {
		cases = append(cases, jsonschematest.Case[string]{
			Name: "country code/" + countryCode, Input: "ABCD" + countryCode + "22", Valid: true,
		})
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringBIC())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_bic.json", cases)
}

func BenchmarkStringBIC(b *testing.B) {
	rule := StringBIC()
	benchmarkStringPaymentBankingRule(b, rule, stringBICTestCases)
}

var stringBICISO93622014TestCases = stringBICTestCases

func TestStringBICISO93622014(t *testing.T) {
	assertPaymentBankingRule(
		t,
		StringBICISO93622014(),
		stringBICISO93622014TestCases,
		"string must be a valid ISO 9362:2014 Business Identifier Code (BIC)",
		ErrorCodeStringBICISO93622014,
	)
}

func TestStringBICCountryCodes(t *testing.T) {
	countryCodes := readISOAlpha2CountryCodes(t)
	assert.Require(t, assert.Len(t, countryCodes, 249))
	uniqueCountryCodes := make(map[string]struct{}, len(countryCodes))
	for _, countryCode := range countryCodes {
		uniqueCountryCodes[countryCode] = struct{}{}
	}
	assert.Require(t, assert.Len(t, uniqueCountryCodes, len(countryCodes)))

	rules := map[string]govy.Rule[string]{
		"StringBIC":            StringBIC(),
		"StringBICISO93622014": StringBICISO93622014(),
	}
	for ruleName, rule := range rules {
		t.Run(ruleName, func(t *testing.T) {
			for _, countryCode := range countryCodes {
				t.Run(countryCode, func(t *testing.T) {
					assert.NoError(t, rule.Validate("ABCD"+countryCode+"22"))
				})
			}
			t.Run("Kosovo exception XK", func(t *testing.T) {
				assert.NoError(t, rule.Validate("ABCDXK22"))
			})
		})
	}
}

func TestBICPredicatesMatchRegexpOracle(t *testing.T) {
	oracle := regexp.MustCompile(`^[A-Z0-9]{4}[A-Z]{2}[A-Z0-9]{2}([A-Z0-9]{3})?$`)
	corpora := map[string][]jsonschematest.Case[string]{
		"StringBIC":            stringBICTestCases,
		"StringBICISO93622014": stringBICISO93622014TestCases,
	}
	for corpusName, testCases := range corpora {
		t.Run(corpusName, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.Name, func(t *testing.T) {
					assertBICPredicatesMatchRegexpOracle(t, oracle, tc.Input)
				})
			}
		})
	}

	t.Run("lengths zero through twelve", func(t *testing.T) {
		for length := range 13 {
			input := strings.Repeat("A", length)
			assertBICPredicatesMatchRegexpOracle(t, oracle, input)
		}
	})

	for name, base := range map[string]string{
		"eight characters":  "ABCDUS22",
		"eleven characters": "ABCDUS22XYZ",
	} {
		t.Run(name+" single-byte substitutions", func(t *testing.T) {
			input := []byte(base)
			for position := range len(input) {
				for value := byte(0); ; value++ {
					input[position] = value
					assertBICPredicatesMatchRegexpOracle(t, oracle, string(input))
					if value == 255 {
						break
					}
				}
				input[position] = base[position]
			}
		})
	}

	t.Run("ISO country codes", func(t *testing.T) {
		countryCodes := append(readISOAlpha2CountryCodes(t), "XK")
		for _, countryCode := range countryCodes {
			assertBICPredicatesMatchRegexpOracle(t, oracle, "ABCD"+countryCode+"22")
			assertBICPredicatesMatchRegexpOracle(t, oracle, "ABCD"+countryCode+"22XYZ")
		}
	})
}

func TestStringBICISO93622014_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := slices.Clone(stringBICTestCases)
	for _, countryCode := range append(readISOAlpha2CountryCodes(t), "XK") {
		cases = append(cases, jsonschematest.Case[string]{
			Name: "country code/" + countryCode, Input: "ABCD" + countryCode + "22", Valid: true,
		})
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringBICISO93622014())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_bic_iso93622014.json", cases)
}

func BenchmarkStringBICISO93622014(b *testing.B) {
	rule := StringBICISO93622014()
	benchmarkStringPaymentBankingRule(b, rule, stringBICISO93622014TestCases)
}

const stringMongoDBObjectIDErrorMessage = "string must be a 24-character hexadecimal MongoDB ObjectID"

// The fixed source corpora are copied verbatim from these pinned revisions:
//   - MongoDB BSON ObjectID corpus (3/3 values):
//     https://github.com/mongodb/specifications/blob/b00d61ca19da7d7e25836ec56930048a8d1de501/source/bson-corpus/tests/oid.json
//   - MongoDB Go driver ObjectID tests (6/6 fixed ObjectIDFromHex inputs):
//     https://github.com/mongodb/mongo-go-driver/blob/25724e5ddec775c78be6a4794279068f3de03b1e/bson/objectid_test.go
//
// The driver's TestFromHex_RoundTrip is excluded because NewObjectID generates
// its input dynamically, so the test publishes no stable literal to copy.
var stringMongoDBObjectIDTestCases = map[string]struct {
	in         string
	shouldFail bool
}{
	"standard lowercase":               {in: "507f1f77bcf86cd799439011"},
	"BSON corpus all zeroes":           {in: "000000000000000000000000"},
	"BSON corpus all ones":             {in: "ffffffffffffffffffffffff"},
	"BSON corpus random":               {in: "56e1fc72e0c917e9c4714161"},
	"Go driver Unix epoch timestamp":   {in: "000000001111111111111111"},
	"Go driver signed timestamp limit": {in: "7FFFFFFF1111111111111111"},
	"Go driver timestamp sign bit":     {in: "800000001111111111111111"},
	"Go driver uint32 timestamp limit": {in: "FFFFFFFF1111111111111111"},
	"mixed case":                       {in: "0123456789abcdefABCDEF01"},

	"empty":                   {shouldFail: true},
	"23 characters":           {in: "507f1f77bcf86cd79943901", shouldFail: true},
	"25 characters":           {in: "507f1f77bcf86cd7994390110", shouldFail: true},
	"Go driver invalid hex":   {in: "this is not a valid hex string!", shouldFail: true},
	"Go driver wrong length":  {in: "deadbeef", shouldFail: true},
	"lowercase non-hex digit": {in: "507f1f77bcf86cd79943901g", shouldFail: true},
	"uppercase non-hex digit": {in: "507F1F77BCF86CD79943901G", shouldFail: true},
	"trailing hyphen":         {in: "507f1f77bcf86cd79943901-", shouldFail: true},
	"trailing colon":          {in: "507f1f77bcf86cd79943901:", shouldFail: true},
	"0x prefix":               {in: "0x507f1f77bcf86cd799439011", shouldFail: true},
	"leading space":           {in: " 507f1f77bcf86cd799439011", shouldFail: true},
	"trailing space":          {in: "507f1f77bcf86cd799439011 ", shouldFail: true},
	"trailing newline":        {in: "507f1f77bcf86cd799439011\n", shouldFail: true},
	"embedded newline":        {in: "507f1f77bcf\n6cd799439011", shouldFail: true},
	"full-width zero":         {in: "０123456789abcdefABCDEF01", shouldFail: true},
	"Cyrillic a":              {in: "а123456789abcdefABCDEF01", shouldFail: true},
	"ObjectId wrapper":        {in: `ObjectId("507f1f77bcf86cd799439011")`, shouldFail: true},
	"Extended JSON wrapper":   {in: `{"$oid":"507f1f77bcf86cd799439011"}`, shouldFail: true},
}

func TestStringMongoDBObjectID(t *testing.T) {
	rule := StringMongoDBObjectID()
	for name, tt := range stringMongoDBObjectIDTestCases {
		t.Run(name, func(t *testing.T) {
			err := rule.Validate(tt.in)
			if tt.shouldFail {
				assert.Require(t, assert.Error(t, err))
				assert.EqualError(t, err, stringMongoDBObjectIDErrorMessage)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMongoDBObjectID))
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_isMongoDBObjectID(t *testing.T) {
	oracle := regexp.MustCompile(`^[0-9a-fA-F]{24}$`)
	assertMatchesOracle := func(t *testing.T, input string) {
		t.Helper()
		expected := oracle.MatchString(input)
		actual := isMongoDBObjectID(input)
		if expected != actual {
			t.Fatalf("input %q: expected %t, got %t", input, expected, actual)
		}
	}

	t.Run("shared corpus", func(t *testing.T) {
		for name, tt := range stringMongoDBObjectIDTestCases {
			t.Run(name, func(t *testing.T) {
				assertMatchesOracle(t, tt.in)
			})
		}
	})
	t.Run("lengths from zero through 48", func(t *testing.T) {
		for length := range 49 {
			assertMatchesOracle(t, strings.Repeat("a", length))
		}
	})
	t.Run("every single-byte substitution", func(t *testing.T) {
		candidate := []byte("000000000000000000000000")
		for position := range candidate {
			for value := range 256 {
				candidate[position] = byte(value)
				assertMatchesOracle(t, string(candidate))
			}
			candidate[position] = '0'
		}
	})
}

func TestStringMongoDBObjectID_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make([]jsonschematest.Case[string], 0, len(stringMongoDBObjectIDTestCases))
	for name, tc := range stringMongoDBObjectIDTestCases {
		cases = append(cases, jsonschematest.Case[string]{
			Name:  name,
			Input: tc.in,
			Valid: !tc.shouldFail,
		})
	}
	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringMongoDBObjectID())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_mongodb_object_id.json", cases)
}

func BenchmarkStringMongoDBObjectID(b *testing.B) {
	rule := StringMongoDBObjectID()
	for name, tt := range stringMongoDBObjectIDTestCases {
		b.Run(name, func(b *testing.B) {
			for b.Loop() {
				_ = rule.Validate(tt.in)
			}
		})
	}
}

var stringEmailTestCases = []jsonschematest.Case[string]{
	{Name: "test@mail.com", Input: "test@mail.com", Valid: true},
	{
		Name:                 "Dörte@Sörensen.example.com",
		Input:                "Dörte@Sörensen.example.com",
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format does not accept internationalized mailbox addresses.",
	},
	{
		Name:                 "θσερ@εχαμπλε.ψομ",
		Input:                "θσερ@εχαμπλε.ψομ",
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format does not accept internationalized mailbox addresses.",
	},
	{
		Name:                 "юзер@екзампл.ком",
		Input:                "юзер@екзампл.ком",
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format does not accept internationalized mailbox addresses.",
	},
	{
		Name:                 "उपयोगकर्ता@उदाहरण.कॉम",
		Input:                "उपयोगकर्ता@उदाहरण.कॉम",
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format does not accept internationalized mailbox addresses.",
	},
	{
		Name:                 "用户@例子.广告",
		Input:                "用户@例子.广告",
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format does not accept internationalized mailbox addresses.",
	},
	{
		Name:                 `"test test"@email.com`,
		Input:                `"test test"@email.com`,
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format does not accept quoted local parts.",
	},
	{
		Name:                 "mail@domain_with_underscores.org",
		Input:                "mail@domain_with_underscores.org",
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format does not accept underscores in domain names.",
	},
	{
		Name:                 "test@email",
		Input:                "test@email",
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format requires a dot in the domain name.",
	},
	{
		Name:                 "test@t",
		Input:                "test@t",
		Valid:                true,
		JSONSchemaDifference: "Ajv's email format requires a dot in the domain name.",
	},
	{Name: "empty", Input: ""},
	{Name: "test@", Input: "test@"},
	{Name: "test", Input: "test"},
	{Name: "test@email.", Input: "test@email."},
	{Name: "@email.com", Input: "@email.com"},
	{Name: `"@email.com`, Input: `"@email.com`},
}

func TestStringEmail(t *testing.T) {
	for _, tc := range stringEmailTestCases {
		err := StringEmail().Validate(tc.Input)
		if !tc.Valid {
			assert.ErrorContains(t, err, "string must be a valid email address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEmail))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringEmail_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringEmail())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_email.json", stringEmailTestCases)
}

func BenchmarkStringEmail(b *testing.B) {
	for _, tc := range stringEmailTestCases {
		rule := StringEmail()
		for range b.N {
			_ = rule.Validate(tc.Input)
		}
	}
}

var stringURLParseErrorTestCase = jsonschematest.Case[string]{
	Name:  "control character",
	Input: "http://\x1f",
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
		err := StringURL().Validate(stringURLParseErrorTestCase.Input)
		assert.ErrorContains(
			t,
			err,
			"failed to parse URL: parse \"http://\\x1f\": net/url: invalid control character in URL",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringURL))
	})
}

func TestStringURL_JSONSchema(t *testing.T) {
	t.Parallel()

	urlCases := make([]jsonschematest.Case[string], 0, len(urlTestCases)+1)
	for _, tc := range urlTestCases {
		urlCases = append(urlCases, jsonschematest.Case[string]{
			Name:                 fmt.Sprintf("%q", tc.url),
			Input:                tc.url,
			Valid:                !tc.shouldFail,
			JSONSchemaDifference: tc.jsonSchemaDifference,
		})
	}
	urlCases = append(urlCases, stringURLParseErrorTestCase)

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringURL())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_url.json", urlCases)
}

func BenchmarkStringURL(b *testing.B) {
	for _, tc := range urlTestCases {
		rule := StringURL()
		for range b.N {
			_ = rule.Validate(tc.url)
		}
	}
}

var stringMACTestCases = []jsonschematest.Case[string]{
	{Name: "3D:F2:C9:A6:B3:4F", Input: "3D:F2:C9:A6:B3:4F", Valid: true},
	{Name: "00:25:96:FF:FE:12:34:56", Input: "00:25:96:FF:FE:12:34:56", Valid: true},
	{Name: "3D-F2-C9-A6-B3:4F", Input: "3D-F2-C9-A6-B3:4F"},
	{Name: "123", Input: "123"},
	{Name: "empty", Input: ""},
	{Name: "abacaba", Input: "abacaba"},
	{Name: "0025:96FF:FE12:3456", Input: "0025:96FF:FE12:3456"},
}

func TestStringMAC(t *testing.T) {
	for _, tc := range stringMACTestCases {
		err := StringMAC().Validate(tc.Input)
		if !tc.Valid {
			assert.EqualError(t, err, "string must be a valid MAC address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMAC))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringMAC_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringMAC())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_mac.json", stringMACTestCases)
}

func BenchmarkStringMAC(b *testing.B) {
	for _, tc := range stringMACTestCases {
		rule := StringMAC()
		for range b.N {
			_ = rule.Validate(tc.Input)
		}
	}
}

var stringIPTestCases = []jsonschematest.Case[string]{
	{Name: "10.0.0.1", Input: "10.0.0.1", Valid: true},
	{Name: "172.16.0.1", Input: "172.16.0.1", Valid: true},
	{Name: "192.168.0.1", Input: "192.168.0.1", Valid: true},
	{Name: "192.168.255.254", Input: "192.168.255.254", Valid: true},
	{Name: "172.16.255.254", Input: "172.16.255.254", Valid: true},
	{Name: "2001:cdba:0000:0000:0000:0000:3257:9652", Input: "2001:cdba:0000:0000:0000:0000:3257:9652", Valid: true},
	{Name: "2001:cdba:0:0:0:0:3257:9652", Input: "2001:cdba:0:0:0:0:3257:9652", Valid: true},
	{Name: "2001:cdba::3257:9652", Input: "2001:cdba::3257:9652", Valid: true},
	{Name: "empty", Input: ""},
	{Name: "172.16.256.255", Input: "172.16.256.255"},
	{Name: "192.168.255.256", Input: "192.168.255.256"},
}

func TestStringIP(t *testing.T) {
	for _, tc := range stringIPTestCases {
		err := StringIP().Validate(tc.Input)
		if !tc.Valid {
			assert.EqualError(t, err, "string must be a valid IP address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIP))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringIP_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringIP())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_ip.json", stringIPTestCases)
}

func BenchmarkStringIP(b *testing.B) {
	for _, tc := range stringIPTestCases {
		rule := StringIP()
		for range b.N {
			_ = rule.Validate(tc.Input)
		}
	}
}

var stringIPv4TestCases = []jsonschematest.Case[string]{
	{Name: "10.0.0.1", Input: "10.0.0.1", Valid: true},
	{Name: "172.16.0.1", Input: "172.16.0.1", Valid: true},
	{Name: "192.168.0.1", Input: "192.168.0.1", Valid: true},
	{Name: "192.168.255.254", Input: "192.168.255.254", Valid: true},
	{Name: "172.16.255.254", Input: "172.16.255.254", Valid: true},
	{Name: "192.168.255.256", Input: "192.168.255.256"},
	{Name: "172.16.256.255", Input: "172.16.256.255"},
	{Name: "2001:cdba:0000:0000:0000:0000:3257:9652", Input: "2001:cdba:0000:0000:0000:0000:3257:9652"},
	{Name: "2001:cdba:0:0:0:0:3257:9652", Input: "2001:cdba:0:0:0:0:3257:9652"},
	{Name: "2001:cdba::3257:9652", Input: "2001:cdba::3257:9652"},
}

func TestStringIPv4(t *testing.T) {
	for _, tc := range stringIPv4TestCases {
		err := StringIPv4().Validate(tc.Input)
		if !tc.Valid {
			assert.EqualError(t, err, "string must be a valid IPv4 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIPv4))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringIPv4_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringIPv4())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_ipv4.json", stringIPv4TestCases)
}

func BenchmarkStringIPv4(b *testing.B) {
	for _, tc := range stringIPv4TestCases {
		rule := StringIPv4()
		for range b.N {
			_ = rule.Validate(tc.Input)
		}
	}
}

var stringIPv6TestCases = []jsonschematest.Case[string]{
	{Name: "2001:cdba:0000:0000:0000:0000:3257:9652", Input: "2001:cdba:0000:0000:0000:0000:3257:9652", Valid: true},
	{Name: "2001:cdba:0:0:0:0:3257:9652", Input: "2001:cdba:0:0:0:0:3257:9652", Valid: true},
	{Name: "2001:cdba::3257:9652", Input: "2001:cdba::3257:9652", Valid: true},
	{Name: "10.0.0.1", Input: "10.0.0.1"},
	{Name: "172.16.0.1", Input: "172.16.0.1"},
	{Name: "192.168.0.1", Input: "192.168.0.1"},
	{Name: "192.168.255.254", Input: "192.168.255.254"},
	{Name: "192.168.255.256", Input: "192.168.255.256"},
	{Name: "172.16.255.254", Input: "172.16.255.254"},
	{Name: "172.16.256.255", Input: "172.16.256.255"},
}

func TestStringIPv6(t *testing.T) {
	for _, tc := range stringIPv6TestCases {
		err := StringIPv6().Validate(tc.Input)
		if !tc.Valid {
			assert.EqualError(t, err, "string must be a valid IPv6 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIPv6))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringIPv6_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringIPv6())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_ipv6.json", stringIPv6TestCases)
}

func BenchmarkStringIPv6(b *testing.B) {
	for _, tc := range stringIPv6TestCases {
		rule := StringIPv6()
		for range b.N {
			_ = rule.Validate(tc.Input)
		}
	}
}

var stringCIDRTestCases = []jsonschematest.Case[string]{
	{Name: "10.0.0.0/0", Input: "10.0.0.0/0", Valid: true},
	{Name: "10.0.0.1/8", Input: "10.0.0.1/8", Valid: true},
	{Name: "172.16.0.1/16", Input: "172.16.0.1/16", Valid: true},
	{Name: "192.168.0.1/24", Input: "192.168.0.1/24", Valid: true},
	{Name: "192.168.255.254/24", Input: "192.168.255.254/24", Valid: true},
	{Name: "172.16.255.254/16", Input: "172.16.255.254/16", Valid: true},
	{
		Name:  "2001:cdba:0000:0000:0000:0000:3257:9652/64",
		Input: "2001:cdba:0000:0000:0000:0000:3257:9652/64",
		Valid: true,
	},
	{Name: "2001:cdba:0:0:0:0:3257:9652/32", Input: "2001:cdba:0:0:0:0:3257:9652/32", Valid: true},
	{Name: "2001:cdba::3257:9652/16", Input: "2001:cdba::3257:9652/16", Valid: true},
	{
		Name:                 "192.168.255.254/48",
		Input:                "192.168.255.254/48",
		JSONSchemaDifference: "The schema checks CIDR structure but does not validate the prefix range.",
	},
	{
		Name:                 "192.168.255.256/24",
		Input:                "192.168.255.256/24",
		JSONSchemaDifference: "The schema checks CIDR structure but does not validate address components.",
	},
	{
		Name:                 "172.16.256.255/16",
		Input:                "172.16.256.255/16",
		JSONSchemaDifference: "The schema checks CIDR structure but does not validate address components.",
	},
	{
		Name:                 "2001:cdba:0000:0000:0000:0000:3257:9652/256",
		Input:                "2001:cdba:0000:0000:0000:0000:3257:9652/256",
		JSONSchemaDifference: "The schema checks CIDR structure but does not validate the prefix range.",
	},
}

func TestStringCIDR(t *testing.T) {
	for _, tc := range stringCIDRTestCases {
		err := StringCIDR().Validate(tc.Input)
		if !tc.Valid {
			assert.EqualError(t, err, "string must be a valid CIDR notation IP address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDR))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringCIDR_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringCIDR())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_cidr.json", stringCIDRTestCases)
}

func BenchmarkStringCIDR(b *testing.B) {
	for _, tc := range stringCIDRTestCases {
		rule := StringCIDR()
		for range b.N {
			_ = rule.Validate(tc.Input)
		}
	}
}

var stringCIDRv4TestCases = []jsonschematest.Case[string]{
	{Name: "0.0.0.0/0", Input: "0.0.0.0/0", Valid: true},
	{Name: "10.0.0.0/8", Input: "10.0.0.0/8", Valid: true},
	{Name: "172.16.0.0/16", Input: "172.16.0.0/16", Valid: true},
	{Name: "192.168.0.0/24", Input: "192.168.0.0/24", Valid: true},
	{Name: "172.16.0.0/16", Input: "172.16.0.0/16", Valid: true},
	{Name: "192.168.255.0/24", Input: "192.168.255.0/24", Valid: true},
	{
		Name:                 "10.0.0.0/0",
		Input:                "10.0.0.0/0",
		JSONSchemaDifference: "The schema checks CIDR structure but does not require network alignment.",
	},
	{
		Name:                 "10.0.0.1/8",
		Input:                "10.0.0.1/8",
		JSONSchemaDifference: "The schema checks CIDR structure but does not require network alignment.",
	},
	{
		Name:                 "172.16.0.1/16",
		Input:                "172.16.0.1/16",
		JSONSchemaDifference: "The schema checks CIDR structure but does not require network alignment.",
	},
	{
		Name:                 "192.168.0.1/24",
		Input:                "192.168.0.1/24",
		JSONSchemaDifference: "The schema checks CIDR structure but does not require network alignment.",
	},
	{
		Name:                 "192.168.255.254/24",
		Input:                "192.168.255.254/24",
		JSONSchemaDifference: "The schema checks CIDR structure but does not require network alignment.",
	},
	{
		Name:                 "192.168.255.254/48",
		Input:                "192.168.255.254/48",
		JSONSchemaDifference: "The schema checks CIDR structure but does not validate the prefix range.",
	},
	{
		Name:                 "192.168.255.256/24",
		Input:                "192.168.255.256/24",
		JSONSchemaDifference: "The schema checks CIDR structure but does not validate address components.",
	},
	{
		Name:                 "172.16.255.254/16",
		Input:                "172.16.255.254/16",
		JSONSchemaDifference: "The schema checks CIDR structure but does not require network alignment.",
	},
	{
		Name:                 "172.16.256.255/16",
		Input:                "172.16.256.255/16",
		JSONSchemaDifference: "The schema checks CIDR structure but does not validate address components.",
	},
	{Name: "2001:cdba:0000:0000:0000:0000:3257:9652/64", Input: "2001:cdba:0000:0000:0000:0000:3257:9652/64"},
	{Name: "2001:cdba:0000:0000:0000:0000:3257:9652/256", Input: "2001:cdba:0000:0000:0000:0000:3257:9652/256"},
	{Name: "2001:cdba:0:0:0:0:3257:9652/32", Input: "2001:cdba:0:0:0:0:3257:9652/32"},
	{Name: "2001:cdba::3257:9652/16", Input: "2001:cdba::3257:9652/16"},
	{
		Name:                 "172.56.1.0/16",
		Input:                "172.56.1.0/16",
		JSONSchemaDifference: "The schema checks CIDR structure but does not require network alignment.",
	},
}

func TestStringCIDRv4(t *testing.T) {
	for _, tc := range stringCIDRv4TestCases {
		err := StringCIDRv4().Validate(tc.Input)
		if !tc.Valid {
			assert.EqualError(t, err, "string must be a valid CIDR notation IPv4 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDRv4))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringCIDRv4_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringCIDRv4())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_cidrv4.json", stringCIDRv4TestCases)
}

func BenchmarkStringCIDRv4(b *testing.B) {
	for _, tc := range stringCIDRv4TestCases {
		rule := StringCIDRv4()
		for range b.N {
			_ = rule.Validate(tc.Input)
		}
	}
}

var stringCIDRv6TestCases = []jsonschematest.Case[string]{
	{
		Name:  "2001:cdba:0000:0000:0000:0000:3257:9652/64",
		Input: "2001:cdba:0000:0000:0000:0000:3257:9652/64",
		Valid: true,
	},
	{Name: "2001:cdba:0:0:0:0:3257:9652/32", Input: "2001:cdba:0:0:0:0:3257:9652/32", Valid: true},
	{Name: "2001:cdba::3257:9652/16", Input: "2001:cdba::3257:9652/16", Valid: true},
	{Name: "10.0.0.0/0", Input: "10.0.0.0/0"},
	{Name: "10.0.0.1/8", Input: "10.0.0.1/8"},
	{Name: "172.16.0.1/16", Input: "172.16.0.1/16"},
	{Name: "192.168.0.1/24", Input: "192.168.0.1/24"},
	{Name: "192.168.255.254/24", Input: "192.168.255.254/24"},
	{Name: "192.168.255.254/48", Input: "192.168.255.254/48"},
	{Name: "192.168.255.256/24", Input: "192.168.255.256/24"},
	{Name: "172.16.255.254/16", Input: "172.16.255.254/16"},
	{Name: "172.16.256.255/16", Input: "172.16.256.255/16"},
	{
		Name:                 "2001:cdba:0000:0000:0000:0000:3257:9652/256",
		Input:                "2001:cdba:0000:0000:0000:0000:3257:9652/256",
		JSONSchemaDifference: "The schema checks CIDR structure but does not validate the prefix range.",
	},
}

func TestStringCIDRv6(t *testing.T) {
	for _, tc := range stringCIDRv6TestCases {
		err := StringCIDRv6().Validate(tc.Input)
		if !tc.Valid {
			assert.EqualError(t, err, "string must be a valid CIDR notation IPv6 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDRv6))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringCIDRv6_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringCIDRv6())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_cidrv6.json", stringCIDRv6TestCases)
}

func BenchmarkStringCIDRv6(b *testing.B) {
	for _, tc := range stringCIDRv6TestCases {
		rule := StringCIDRv6()
		for range b.N {
			_ = rule.Validate(tc.Input)
		}
	}
}

var stringJSONTestCases = []jsonschematest.Case[string]{
	{Name: `{"foo": "bar"}`, Input: `{"foo": "bar"}`, Valid: true},
	{Name: `{}`, Input: `{}`, Valid: true},
	{Name: `[]`, Input: `[]`, Valid: true},
	{
		Name:                 "{]}",
		Input:                "{]}",
		JSONSchemaDifference: "contentMediaType is an annotation and does not validate the embedded JSON.",
	},
	{
		Name:                 "empty",
		Input:                "",
		JSONSchemaDifference: "contentMediaType is an annotation and does not validate the embedded JSON.",
	},
	{
		Name:                 "yaml: ok",
		Input:                "yaml: ok",
		JSONSchemaDifference: "contentMediaType is an annotation and does not validate the embedded JSON.",
	},
}

func TestStringJSON(t *testing.T) {
	for _, tc := range stringJSONTestCases {
		err := StringJSON().Validate(tc.Input)
		if !tc.Valid {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringJSON))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringJSON_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringJSON())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_json.json", stringJSONTestCases)
}

func BenchmarkStringJSON(b *testing.B) {
	for _, tc := range stringJSONTestCases {
		rule := StringJSON()
		for range b.N {
			_ = rule.Validate(tc.Input)
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

func TestStringSemver_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make([]jsonschematest.Case[string], 0, len(validSemverTestCases)+len(invalidSemverTestCases))
	for _, input := range validSemverTestCases {
		cases = append(cases, jsonschematest.Case[string]{Name: "valid/" + input, Input: input, Valid: true})
	}
	for _, input := range invalidSemverTestCases {
		cases = append(cases, jsonschematest.Case[string]{Name: "invalid/" + input, Input: input})
	}
	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringSemver())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_semver.json", cases)
}

func BenchmarkStringSemver(b *testing.B) {
	rule := StringSemver()
	for b.Loop() {
		for _, input := range validSemverTestCases {
			_ = rule.Validate(input)
		}
		for _, input := range invalidSemverTestCases {
			_ = rule.Validate(input)
		}
	}
	b.ReportMetric(float64(len(validSemverTestCases)+len(invalidSemverTestCases)), "validations/op")
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

func TestStringCVE_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make([]jsonschematest.Case[string], 0, len(stringCVETestCases))
	for name, tc := range stringCVETestCases {
		cases = append(cases, jsonschematest.Case[string]{Name: name, Input: tc.in, Valid: tc.expectedError == ""})
	}
	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringCVE())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_cve.json", cases)
}

func BenchmarkStringCVE(b *testing.B) {
	rule := StringCVE()
	for b.Loop() {
		for _, tc := range stringCVETestCases {
			_ = rule.Validate(tc.in)
		}
	}
	b.ReportMetric(float64(len(stringCVETestCases)), "validations/op")
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

func TestStringE164_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make([]jsonschematest.Case[string], 0, len(stringE164TestCases))
	for name, tc := range stringE164TestCases {
		cases = append(cases, jsonschematest.Case[string]{Name: name, Input: tc.in, Valid: tc.expectedError == ""})
	}
	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringE164())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_e164.json", cases)
}

func BenchmarkStringE164(b *testing.B) {
	for name, tc := range stringE164TestCases {
		b.Run(name, func(b *testing.B) {
			rule := StringE164()
			for b.Loop() {
				_ = rule.Validate(tc.in)
			}
		})
	}
}

// Corpus sources: RFC 4648 sections 9-10 and golang/go commit
// fefb02adf45c4bcc879bd406a8d61f2a292c26a9 (Go 1.25.5),
// src/encoding/base64/base64_test.go. The tables cover all 17 unique Go
// pairs, all 24 TestDecodeCorrupt inputs, all 11 TestNewLineCharacters
// inputs, and both strict-decoder issue 15656 inputs. Go's duplicate
// "sure." pair is represented once. CR and LF are invalid under this rule's
// RFC 4648 generic profile even though Go's decoder ignores them.
var stringBase64ValidInputs = map[string]jsonschematest.Case[string]{
	"empty":                           {Input: "", Valid: true},
	"rfc 4648 one byte":               {Input: "Zg==", Valid: true},
	"rfc 4648 two bytes":              {Input: "Zm8=", Valid: true},
	"rfc 4648 three bytes":            {Input: "Zm9v", Valid: true},
	"rfc 4648 four bytes":             {Input: "Zm9vYg==", Valid: true},
	"rfc 4648 five bytes":             {Input: "Zm9vYmE=", Valid: true},
	"rfc 4648 six bytes":              {Input: "Zm9vYmFy", Valid: true},
	"rfc 4648 six-byte illustration":  {Input: "FPucA9l+", Valid: true},
	"rfc 4648 five-byte illustration": {Input: "FPucA9k=", Valid: true},
	"rfc 4648 four-byte illustration": {Input: "FPucAw==", Valid: true},
	"single zero byte":                {Input: "AA==", Valid: true},
	"two zero bytes":                  {Input: "AAA=", Valid: true},
	"three zero bytes":                {Input: "AAAA", Valid: true},
	"standard alphabet characters":    {Input: "+/8=", Valid: true},
	"go corpus sure with period":      {Input: "c3VyZS4=", Valid: true},
	"go corpus sure":                  {Input: "c3VyZQ==", Valid: true},
	"go corpus sur":                   {Input: "c3Vy", Valid: true},
	"go corpus su":                    {Input: "c3U=", Valid: true},
	"go corpus eight-byte variant":    {Input: "bGVhc3VyZS4=", Valid: true},
	"go corpus seven-byte variant":    {Input: "ZWFzdXJlLg==", Valid: true},
	"go corpus six-byte variant":      {Input: "YXN1cmUu", Valid: true},
	"go strict canonical pad bits":    {Input: "WvLTlMrX9NpYDQlEIFlnDA==", Valid: true},
}

var stringBase64InvalidInputs = map[string]jsonschematest.Case[string]{
	"noncanonical one-byte pad bits": {
		Input:                "Zh==",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
	"noncanonical two-byte pad bits": {
		Input:                "Zm9=",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
	"invalid alphabet characters":                                 {Input: "!!!!"},
	"padding only":                                                {Input: "===="},
	"one symbol with three padding characters":                    {Input: "x==="},
	"leading padding":                                             {Input: "=AAA"},
	"padding in second position":                                  {Input: "A=AA"},
	"padding in third position":                                   {Input: "AA=A"},
	"data after double padding":                                   {Input: "AA==A"},
	"data after padding":                                          {Input: "AAA=AAAA"},
	"length modulo four is one":                                   {Input: "AAAAA"},
	"six symbols without required padding":                        {Input: "AAAAAA"},
	"one symbol with one padding character":                       {Input: "A="},
	"one symbol with two padding characters":                      {Input: "A=="},
	"two symbols with one padding character":                      {Input: "AA="},
	"six symbols with one padding character":                      {Input: "AAAAAA="},
	"excess padding":                                              {Input: "YWJjZA====="},
	"missing double padding":                                      {Input: "Zg"},
	"missing single padding":                                      {Input: "Zm8"},
	"padding after complete quantum":                              {Input: "Zm9v="},
	"three padding characters after one byte":                     {Input: "Zg==="},
	"concatenated padded values":                                  {Input: "Zg==AA=="},
	"url-safe alphabet":                                           {Input: "-_8="},
	"line feed":                                                   {Input: "Zm9v\n"},
	"carriage return and line feed":                               {Input: "Zm9v\r\n"},
	"embedded carriage return and line feed":                      {Input: "Zm\r\n9v"},
	"newline only":                                                {Input: "\n"},
	"padded data followed by newline":                             {Input: "AAA=\n"},
	"complete quantum followed by newline":                        {Input: "AAAA\n"},
	"invalid symbol followed by newline":                          {Input: "A!\n"},
	"incomplete padding followed by newline":                      {Input: "A=\n"},
	"go newline corpus trailing carriage return":                  {Input: "c3VyZQ==\r"},
	"go newline corpus trailing line feed":                        {Input: "c3VyZQ==\n"},
	"go newline corpus trailing CRLF":                             {Input: "c3VyZQ==\r\n"},
	"go newline corpus embedded CRLF":                             {Input: "c3VyZ\r\nQ=="},
	"go newline corpus interleaved carriage return and line feed": {Input: "c3V\ryZ\nQ=="},
	"go newline corpus interleaved line feed and carriage return": {Input: "c3V\nyZ\rQ=="},
	"go newline corpus before fourth symbol":                      {Input: "c3VyZ\nQ=="},
	"go newline corpus before padding":                            {Input: "c3VyZQ\n=="},
	"go newline corpus between padding":                           {Input: "c3VyZQ=\n="},
	"go newline corpus repeated CRLF between padding":             {Input: "c3VyZQ=\r\n\r\n="},
	"embedded space":                                              {Input: "Zm 9v"},
	"leading space":                                               {Input: " Zm9v"},
	"trailing tab":                                                {Input: "Zm9v\t"},
	"NUL byte":                                                    {Input: "Zm9v\x00"},
	"zero-width space":                                            {Input: "Zm9v\u200b"},
	"non-ASCII letter":                                            {Input: "Zm9vé"},
	"mixed standard and URL-safe alphabet":                        {Input: "AA+_"},
	"go strict noncanonical pad bits": {
		Input:                "WvLTlMrX9NpYDQlEIFlnDB==",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
}

func TestStringBase64(t *testing.T) {
	runStringEncodingRuleTest(
		t,
		StringBase64(),
		ErrorCodeStringBase64,
		"string must be a valid padded Base64 value using the standard alphabet",
		stringBase64ValidInputs,
		stringBase64InvalidInputs,
	)
}

func TestStringBase64_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make([]jsonschematest.Case[string], 0, len(stringBase64ValidInputs)+len(stringBase64InvalidInputs))
	for name, tc := range stringBase64ValidInputs {
		tc.Name = "valid/" + name
		cases = append(cases, tc)
	}
	for name, tc := range stringBase64InvalidInputs {
		tc.Name = "invalid/" + name
		cases = append(cases, tc)
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringBase64())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_base64.json", cases)
}

func BenchmarkStringBase64(b *testing.B) {
	benchmarkStringEncodingRule(
		b,
		StringBase64(),
		stringBase64ValidInputs,
		stringBase64InvalidInputs,
	)
}

// The padded URL-safe corpus applies the same complete Go tables after the
// standard-to-URL alphabet conversion and includes RFC 7515 Appendix C.
var stringBase64URLValidInputs = map[string]jsonschematest.Case[string]{
	"empty":                             {Input: "", Valid: true},
	"rfc 4648 one byte":                 {Input: "Zg==", Valid: true},
	"rfc 4648 two bytes":                {Input: "Zm8=", Valid: true},
	"rfc 4648 three bytes":              {Input: "Zm9v", Valid: true},
	"rfc 4648 four bytes":               {Input: "Zm9vYg==", Valid: true},
	"rfc 4648 five bytes":               {Input: "Zm9vYmE=", Valid: true},
	"rfc 4648 six bytes":                {Input: "Zm9vYmFy", Valid: true},
	"rfc 4648 six-byte URL form":        {Input: "FPucA9l-", Valid: true},
	"rfc 4648 five-byte URL form":       {Input: "FPucA9k=", Valid: true},
	"rfc 4648 four-byte URL form":       {Input: "FPucAw==", Valid: true},
	"single zero byte":                  {Input: "AA==", Valid: true},
	"two zero bytes":                    {Input: "AAA=", Valid: true},
	"three zero bytes":                  {Input: "AAAA", Valid: true},
	"six zero-bit symbols with padding": {Input: "AAAAAA==", Valid: true},
	"rfc 7515 appendix C padded form":   {Input: "A-z_4ME=", Valid: true},
	"url-safe alphabet characters":      {Input: "-_8=", Valid: true},
	"go corpus sure with period":        {Input: "c3VyZS4=", Valid: true},
	"go corpus sure":                    {Input: "c3VyZQ==", Valid: true},
	"go corpus sur":                     {Input: "c3Vy", Valid: true},
	"go corpus su":                      {Input: "c3U=", Valid: true},
	"go corpus eight-byte variant":      {Input: "bGVhc3VyZS4=", Valid: true},
	"go corpus seven-byte variant":      {Input: "ZWFzdXJlLg==", Valid: true},
	"go corpus six-byte variant":        {Input: "YXN1cmUu", Valid: true},
	"go strict canonical pad bits":      {Input: "WvLTlMrX9NpYDQlEIFlnDA==", Valid: true},
}

var stringBase64URLInvalidInputs = map[string]jsonschematest.Case[string]{
	"noncanonical one-byte pad bits": {
		Input:                "Zh==",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
	"noncanonical two-byte pad bits": {
		Input:                "Zm9=",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
	"invalid alphabet characters":                                 {Input: "!!!!"},
	"padding only":                                                {Input: "===="},
	"one symbol with three padding characters":                    {Input: "x==="},
	"leading padding":                                             {Input: "=AAA"},
	"padding in second position":                                  {Input: "A=AA"},
	"padding in third position":                                   {Input: "AA=A"},
	"data after double padding":                                   {Input: "AA==A"},
	"data after padding":                                          {Input: "AAA=AAAA"},
	"length modulo four is one":                                   {Input: "AAAAA"},
	"six symbols without required padding":                        {Input: "AAAAAA"},
	"one symbol with one padding character":                       {Input: "A="},
	"one symbol with two padding characters":                      {Input: "A=="},
	"two symbols with one padding character":                      {Input: "AA="},
	"six symbols with one padding character":                      {Input: "AAAAAA="},
	"excess padding":                                              {Input: "YWJjZA====="},
	"missing double padding":                                      {Input: "Zg"},
	"missing single padding":                                      {Input: "Zm8"},
	"padding after complete quantum":                              {Input: "Zm9v="},
	"three padding characters after one byte":                     {Input: "Zg==="},
	"concatenated padded values":                                  {Input: "Zg==AA=="},
	"standard alphabet":                                           {Input: "+/8="},
	"line feed":                                                   {Input: "Zm9v\n"},
	"carriage return and line feed":                               {Input: "Zm9v\r\n"},
	"embedded carriage return and line feed":                      {Input: "Zm\r\n9v"},
	"newline only":                                                {Input: "\n"},
	"padded data followed by newline":                             {Input: "AAA=\n"},
	"complete quantum followed by newline":                        {Input: "AAAA\n"},
	"invalid symbol followed by newline":                          {Input: "A!\n"},
	"incomplete padding followed by newline":                      {Input: "A=\n"},
	"go newline corpus trailing carriage return":                  {Input: "c3VyZQ==\r"},
	"go newline corpus trailing line feed":                        {Input: "c3VyZQ==\n"},
	"go newline corpus trailing CRLF":                             {Input: "c3VyZQ==\r\n"},
	"go newline corpus embedded CRLF":                             {Input: "c3VyZ\r\nQ=="},
	"go newline corpus interleaved carriage return and line feed": {Input: "c3V\ryZ\nQ=="},
	"go newline corpus interleaved line feed and carriage return": {Input: "c3V\nyZ\rQ=="},
	"go newline corpus before fourth symbol":                      {Input: "c3VyZ\nQ=="},
	"go newline corpus before padding":                            {Input: "c3VyZQ\n=="},
	"go newline corpus between padding":                           {Input: "c3VyZQ=\n="},
	"go newline corpus repeated CRLF between padding":             {Input: "c3VyZQ=\r\n\r\n="},
	"embedded space":                                              {Input: "Zm 9v"},
	"leading space":                                               {Input: " Zm9v"},
	"trailing tab":                                                {Input: "Zm9v\t"},
	"NUL byte":                                                    {Input: "Zm9v\x00"},
	"zero-width space":                                            {Input: "Zm9v\u200b"},
	"non-ASCII letter":                                            {Input: "Zm9vé"},
	"mixed standard and URL-safe alphabet":                        {Input: "AA+_"},
	"go strict noncanonical pad bits": {
		Input:                "WvLTlMrX9NpYDQlEIFlnDB==",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
}

func TestStringBase64URL(t *testing.T) {
	runStringEncodingRuleTest(
		t,
		StringBase64URL(),
		ErrorCodeStringBase64URL,
		"string must be a valid padded URL-safe Base64 value",
		stringBase64URLValidInputs,
		stringBase64URLInvalidInputs,
	)
}

func TestStringBase64URL_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make([]jsonschematest.Case[string], 0, len(stringBase64URLValidInputs)+len(stringBase64URLInvalidInputs))
	for name, tc := range stringBase64URLValidInputs {
		tc.Name = "valid/" + name
		cases = append(cases, tc)
	}
	for name, tc := range stringBase64URLInvalidInputs {
		tc.Name = "invalid/" + name
		cases = append(cases, tc)
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringBase64URL())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_base64_url.json", cases)
}

func BenchmarkStringBase64URL(b *testing.B) {
	benchmarkStringEncodingRule(
		b,
		StringBase64URL(),
		stringBase64URLValidInputs,
		stringBase64URLInvalidInputs,
	)
}

// The raw URL-safe tables contain all 17 unique Go pairs after rawURLRef,
// all 24 exact TestDecodeCorrupt inputs (3 valid and 21 invalid), and all 11
// exact TestNewLineCharacters inputs (all invalid), plus RFC 7515 Appendix C
// and strict canonical pad-bit boundaries. Overlapping literals are represented
// once, and exact corpus inputs are not replaced by rawURLRef transformations.
var stringBase64RawURLValidInputs = map[string]jsonschematest.Case[string]{
	"empty":                            {Input: "", Valid: true},
	"one byte without padding":         {Input: "Zg", Valid: true},
	"two bytes without padding":        {Input: "Zm8", Valid: true},
	"three bytes":                      {Input: "Zm9v", Valid: true},
	"four bytes without padding":       {Input: "Zm9vYg", Valid: true},
	"five bytes without padding":       {Input: "Zm9vYmE", Valid: true},
	"six bytes":                        {Input: "Zm9vYmFy", Valid: true},
	"rfc 4648 six-byte raw URL form":   {Input: "FPucA9l-", Valid: true},
	"rfc 4648 five-byte raw URL form":  {Input: "FPucA9k", Valid: true},
	"rfc 4648 four-byte raw URL form":  {Input: "FPucAw", Valid: true},
	"single zero byte without padding": {Input: "AA", Valid: true},
	"two zero bytes without padding":   {Input: "AAA", Valid: true},
	"three zero bytes":                 {Input: "AAAA", Valid: true},
	"six zero-bit symbols":             {Input: "AAAAAA", Valid: true},
	"rfc 7515 appendix C":              {Input: "A-z_4ME", Valid: true},
	"url-safe alphabet characters":     {Input: "-_8", Valid: true},
	"go corpus sure with period":       {Input: "c3VyZS4", Valid: true},
	"go corpus sure":                   {Input: "c3VyZQ", Valid: true},
	"go corpus sur":                    {Input: "c3Vy", Valid: true},
	"go corpus su":                     {Input: "c3U", Valid: true},
	"go corpus eight-byte variant":     {Input: "bGVhc3VyZS4", Valid: true},
	"go corpus seven-byte variant":     {Input: "ZWFzdXJlLg", Valid: true},
	"go corpus six-byte variant":       {Input: "YXN1cmUu", Valid: true},
	"go strict canonical pad bits":     {Input: "WvLTlMrX9NpYDQlEIFlnDA", Valid: true},
}

var stringBase64RawURLInvalidInputs = map[string]jsonschematest.Case[string]{
	"noncanonical one-byte pad bits": {
		Input:                "Zh",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
	"noncanonical two-byte pad bits": {
		Input:                "Zm9",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
	"length modulo four is one": {
		Input:                "A",
		JSONSchemaDifference: "The schema checks the Base64 alphabet but does not reject impossible unpadded lengths.",
	},
	"longer length modulo four is one": {
		Input:                "AAAAA",
		JSONSchemaDifference: "The schema checks the Base64 alphabet but does not reject impossible unpadded lengths.",
	},
	"single padding character":               {Input: "Zm8="},
	"double padding characters":              {Input: "Zg=="},
	"padding after complete quantum":         {Input: "Zm9v="},
	"standard alphabet":                      {Input: "+/8"},
	"line feed":                              {Input: "Zm9v\n"},
	"carriage return and line feed":          {Input: "Zm9v\r\n"},
	"embedded carriage return and line feed": {Input: "Zm\r\n9v"},
	"embedded space":                         {Input: "Zm 9v"},
	"leading space":                          {Input: " Zm9v"},
	"trailing tab":                           {Input: "Zm9v\t"},
	"NUL byte":                               {Input: "Zm9v\x00"},
	"zero-width space":                       {Input: "Zm9v\u200b"},
	"non-ASCII letter":                       {Input: "Zm9vé"},
	"mixed standard and URL-safe alphabet":   {Input: "AA+_"},
	"go strict noncanonical pad bits": {
		Input:                "WvLTlMrX9NpYDQlEIFlnDB",
		JSONSchemaDifference: "The schema checks Base64 shape but does not require zero padding bits.",
	},
	"go corrupt corpus newline only":                              {Input: "\n"},
	"go corrupt corpus padded data followed by newline":           {Input: "AAA=\n"},
	"go corrupt corpus complete quantum followed by newline":      {Input: "AAAA\n"},
	"go corrupt corpus invalid alphabet":                          {Input: "!!!!"},
	"go corrupt corpus padding only":                              {Input: "===="},
	"go corrupt corpus excess padding after one symbol":           {Input: "x==="},
	"go corrupt corpus leading padding":                           {Input: "=AAA"},
	"go corrupt corpus padding in second position":                {Input: "A=AA"},
	"go corrupt corpus padding in third position":                 {Input: "AA=A"},
	"go corrupt corpus data after double padding":                 {Input: "AA==A"},
	"go corrupt corpus data after padding":                        {Input: "AAA=AAAA"},
	"go corrupt corpus one symbol with padding":                   {Input: "A="},
	"go corrupt corpus one symbol with double padding":            {Input: "A=="},
	"go corrupt corpus two symbols with padding":                  {Input: "AA="},
	"go corrupt corpus two symbols with double padding":           {Input: "AA=="},
	"go corrupt corpus three symbols with padding":                {Input: "AAA="},
	"go corrupt corpus excess padding after six symbols":          {Input: "AAAAAA="},
	"go corrupt corpus excess terminal padding":                   {Input: "YWJjZA====="},
	"go corrupt corpus invalid symbol followed by newline":        {Input: "A!\n"},
	"go corrupt corpus incomplete padding followed by newline":    {Input: "A=\n"},
	"go newline corpus padded base":                               {Input: "c3VyZQ=="},
	"go newline corpus trailing carriage return":                  {Input: "c3VyZQ==\r"},
	"go newline corpus trailing line feed":                        {Input: "c3VyZQ==\n"},
	"go newline corpus trailing CRLF":                             {Input: "c3VyZQ==\r\n"},
	"go newline corpus embedded CRLF":                             {Input: "c3VyZ\r\nQ=="},
	"go newline corpus interleaved carriage return and line feed": {Input: "c3V\ryZ\nQ=="},
	"go newline corpus interleaved line feed and carriage return": {Input: "c3V\nyZ\rQ=="},
	"go newline corpus before fourth symbol":                      {Input: "c3VyZ\nQ=="},
	"go newline corpus before padding":                            {Input: "c3VyZQ\n=="},
	"go newline corpus between padding":                           {Input: "c3VyZQ=\n="},
	"go newline corpus repeated CRLF between padding":             {Input: "c3VyZQ=\r\n\r\n="},
}

func TestStringBase64RawURL(t *testing.T) {
	runStringEncodingRuleTest(
		t,
		StringBase64RawURL(),
		ErrorCodeStringBase64RawURL,
		"string must be a valid URL-safe Base64 value without padding",
		stringBase64RawURLValidInputs,
		stringBase64RawURLInvalidInputs,
	)
}

func TestStringBase64RawURL_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make(
		[]jsonschematest.Case[string],
		0,
		len(stringBase64RawURLValidInputs)+len(stringBase64RawURLInvalidInputs),
	)
	for name, tc := range stringBase64RawURLValidInputs {
		tc.Name = "valid/" + name
		cases = append(cases, tc)
	}
	for name, tc := range stringBase64RawURLInvalidInputs {
		tc.Name = "invalid/" + name
		cases = append(cases, tc)
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringBase64RawURL())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_base64_raw_url.json", cases)
}

func BenchmarkStringBase64RawURL(b *testing.B) {
	benchmarkStringEncodingRule(
		b,
		StringBase64RawURL(),
		stringBase64RawURLValidInputs,
		stringBase64RawURLInvalidInputs,
	)
}

// Corpus sources: RFC 4648 section 10 and golang/go commit
// fefb02adf45c4bcc879bd406a8d61f2a292c26a9 (Go 1.25.5),
// src/encoding/hex/hex_test.go. All 8 Go encode/decode inputs and all 9
// error-table inputs are represented. Empty is invalid under this rule's
// nonempty digit-string contract, while Go's three odd-length error inputs are
// valid here.
var stringHexadecimalValidInputs = map[string]jsonschematest.Case[string]{
	"two digits":                            {Input: "66", Valid: true},
	"four digits":                           {Input: "666F", Valid: true},
	"six digits":                            {Input: "666F6F", Valid: true},
	"eight digits":                          {Input: "666F6F62", Valid: true},
	"ten digits":                            {Input: "666F6F6261", Valid: true},
	"twelve digits":                         {Input: "666F6F626172", Valid: true},
	"single digit":                          {Input: "F", Valid: true},
	"two zero digits":                       {Input: "00", Valid: true},
	"lowercase":                             {Input: "deadbeef", Valid: true},
	"uppercase":                             {Input: "DEADBEEF", Valid: true},
	"lowercase prefix":                      {Input: "0xdeadBEEF", Valid: true},
	"uppercase prefix":                      {Input: "0XABCDEF", Valid: true},
	"lowercase prefix with one digit":       {Input: "0x0", Valid: true},
	"uppercase prefix with one digit":       {Input: "0Xf", Valid: true},
	"hex digits beginning with lowercase b": {Input: "0b1010", Valid: true},
	"hex digits beginning with uppercase B": {Input: "0B1010", Valid: true},
	"go byte range zero through seven":      {Input: "0001020304050607", Valid: true},
	"go byte range eight through fifteen":   {Input: "08090a0b0c0d0e0f", Valid: true},
	"go byte range f0 through f7":           {Input: "f0f1f2f3f4f5f6f7", Valid: true},
	"go byte range f8 through ff":           {Input: "f8f9fafbfcfdfeff", Valid: true},
	"go single byte":                        {Input: "67", Valid: true},
	"go two bytes":                          {Input: "e3a1", Valid: true},
	"go uppercase byte range":               {Input: "F8F9FAFBFCFDFEFF", Valid: true},
	"go odd single digit":                   {Input: "0", Valid: true},
	"go odd numeric sequence":               {Input: "30313", Valid: true},
	"go odd alphabetic sequence":            {Input: "ffeed", Valid: true},
}

var stringHexadecimalInvalidInputs = map[string]jsonschematest.Case[string]{
	"empty":                            {Input: ""},
	"lowercase prefix only":            {Input: "0x"},
	"uppercase prefix only":            {Input: "0X"},
	"repeated prefix":                  {Input: "0x0x1"},
	"prefix after digit":               {Input: "10x1"},
	"minus sign":                       {Input: "-0x1"},
	"plus sign":                        {Input: "+F"},
	"underscore":                       {Input: "dead_beef"},
	"leading whitespace":               {Input: " deadbeef"},
	"trailing whitespace":              {Input: "deadbeef "},
	"invalid digit G":                  {Input: "G"},
	"full-width F":                     {Input: "\uff26"},
	"Arabic-Indic digit":               {Input: "\u0660"},
	"NUL byte":                         {Input: "00\x00"},
	"octal prefix":                     {Input: "0o755"},
	"go invalid first digit":           {Input: "zd4aa"},
	"go invalid final digit":           {Input: "d4aaz"},
	"go invalid second digit":          {Input: "0g"},
	"go invalid trailing digits":       {Input: "00gg"},
	"go control byte":                  {Input: "0\x01"},
	"prefix after zero digit":          {Input: "00x1"},
	"plus before prefix":               {Input: "+0x1"},
	"underscore after prefix":          {Input: "0x1_2"},
	"leading whitespace before prefix": {Input: " 0x1"},
	"trailing whitespace after prefix": {Input: "0x1 "},
	"whitespace after prefix":          {Input: "0x 1"},
	"invalid prefixed digit G":         {Input: "0xG"},
	"full-width digits":                {Input: "１２"},
	"non-ASCII letter":                 {Input: "é"},
	"NUL byte only":                    {Input: "\x00"},
	"audit octal prefix":               {Input: "0o77"},
}

func TestStringHexadecimal(t *testing.T) {
	runStringEncodingRuleTest(
		t,
		StringHexadecimal(),
		ErrorCodeStringHexadecimal,
		"string must be a valid hexadecimal value",
		stringHexadecimalValidInputs,
		stringHexadecimalInvalidInputs,
	)
}

func Test_isHexadecimalMatchesRegexp(t *testing.T) {
	reference := regexp.MustCompile(`^(?:0[xX])?[0-9a-fA-F]+$`)
	check := func(input string) {
		expected := reference.MatchString(input)
		if actual := isHexadecimal(input); actual != expected {
			t.Fatalf("isHexadecimal(%q) = %t, expected %t", input, actual, expected)
		}
	}

	for _, prefix := range []string{"", "0x", "0X"} {
		check(prefix)
		for first := range 256 {
			check(prefix + string([]byte{byte(first)}))
			for second := range 256 {
				check(prefix + string([]byte{byte(first), byte(second)}))
			}
		}
	}
}

func TestStringHexadecimal_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make(
		[]jsonschematest.Case[string],
		0,
		len(stringHexadecimalValidInputs)+len(stringHexadecimalInvalidInputs),
	)
	for name, tc := range stringHexadecimalValidInputs {
		tc.Name = "valid/" + name
		cases = append(cases, tc)
	}
	for name, tc := range stringHexadecimalInvalidInputs {
		tc.Name = "invalid/" + name
		cases = append(cases, tc)
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringHexadecimal())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_hexadecimal.json", cases)
}

func BenchmarkStringHexadecimal(b *testing.B) {
	benchmarkStringEncodingRule(
		b,
		StringHexadecimal(),
		stringHexadecimalValidInputs,
		stringHexadecimalInvalidInputs,
	)
}

func runStringEncodingRuleTest(
	t *testing.T,
	rule govy.Rule[string],
	errorCode govy.ErrorCode,
	expectedError string,
	validInputs map[string]jsonschematest.Case[string],
	invalidInputs map[string]jsonschematest.Case[string],
) {
	t.Helper()
	t.Run("valid", func(t *testing.T) {
		for name, tc := range validInputs {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(tc.Input))
			})
		}
	})
	t.Run("invalid", func(t *testing.T) {
		for name, tc := range invalidInputs {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(tc.Input)
				assert.EqualError(t, err, expectedError)
				assert.True(t, govy.HasErrorCode(err, errorCode))
			})
		}
	})
}

func benchmarkStringEncodingRule(
	b *testing.B,
	rule govy.Rule[string],
	validInputs map[string]jsonschematest.Case[string],
	invalidInputs map[string]jsonschematest.Case[string],
) {
	b.Helper()
	b.Run("valid", func(b *testing.B) {
		benchmarkStringEncodingInputs(b, rule, validInputs)
	})
	b.Run("invalid", func(b *testing.B) {
		benchmarkStringEncodingInputs(b, rule, invalidInputs)
	})
}

func benchmarkStringEncodingInputs(
	b *testing.B,
	rule govy.Rule[string],
	inputs map[string]jsonschematest.Case[string],
) {
	b.Helper()
	for b.Loop() {
		for _, tc := range inputs {
			_ = rule.Validate(tc.Input)
		}
	}
	b.ReportMetric(float64(len(inputs)), "validations/op")
}

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

func TestStringEIN_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringEIN())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_ein.json", taxJSONSchemaCases(
		stringEINPrefixTestCases,
		stringEINAcceptedStructureTestCases,
		stringEINRejectedStructureTestCases,
	))
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

func TestStringSSN_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringSSN())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_ssn.json", taxJSONSchemaCases(
		stringSSNAcceptedStructureTestCases,
		stringSSNRejectedStructureTestCases,
		stringSSNAreaTestCases,
		stringSSNGroupTestCases,
		stringSSNSerialTestCases,
	))
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

// The first seven values are the complete RFC 1321 Appendix A.5 output suite:
// https://www.rfc-editor.org/rfc/rfc1321.html#appendix-A.5
var validMD5TestCases = map[string]string{
	"RFC 1321 empty message":         "d41d8cd98f00b204e9800998ecf8427e",
	"RFC 1321 single letter":         "0cc175b9c0f1b6a831c399e269772661",
	"RFC 1321 abc":                   "900150983cd24fb0d6963f7d28e17f72",
	"RFC 1321 message digest":        "f96b697d7cb7938d525a2f31aaf161d0",
	"RFC 1321 lowercase alphabet":    "c3fcd3d76192e4007dfb496cca67e13b",
	"RFC 1321 alphanumeric":          "d174ab98d277d9f5a5611c2c9f419d9f",
	"RFC 1321 numeric sequence":      "57edf4a22be3c955ac49da2e2107b67a",
	"zero-leading repeated sequence": "0123456789abcdef0123456789abcdef",
}

var invalidMD5TestCases = map[string]string{
	"empty":                              "",
	"all uppercase":                      "D41D8CD98F00B204E9800998ECF8427E",
	"mixed case":                         "d41d8cd98f00b204e9800998ecf8427E",
	"one character short":                "d41d8cd98f00b204e9800998ecf8427",
	"one character long":                 "d41d8cd98f00b204e9800998ecf8427e0",
	"terminal non-hexadecimal character": "d41d8cd98f00b204e9800998ecf8427g",
	"0x prefix":                          "0x0123456789abcdef0123456789abcdef",
	"leading space":                      " 0123456789abcdef0123456789abcdef",
	"trailing space":                     "0123456789abcdef0123456789abcdef ",
	"trailing newline":                   "0123456789abcdef0123456789abcdef\n",
	"embedded hyphen":                    "01234567-9abcdef0123456789abcdef",
	"embedded Unicode letter":            "01234567é9abcdef0123456789abcdef",
}

func TestStringMD5(t *testing.T) {
	rule := StringMD5()
	t.Run("valid digests", func(t *testing.T) {
		for name, digest := range validMD5TestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(digest))
			})
		}
	})
	t.Run("invalid digests", func(t *testing.T) {
		for name, digest := range invalidMD5TestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(digest)
				assert.EqualError(t, err, "string must be a valid lowercase MD5 hexadecimal digest")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMD5))
			})
		}
	})
}

func TestStringMD5_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringMD5())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_md5.json",
		stringNamedJSONSchemaCases(validMD5TestCases, invalidMD5TestCases),
	)
}

func BenchmarkStringMD5(b *testing.B) {
	rule := StringMD5()
	benchmarkStringHashDigestRule(b, rule, validMD5TestCases, invalidMD5TestCases)
}

var validSHA256TestCases = map[string]string{
	"letter-leading":                 "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	"nonzero-digit-leading":          "28969cdfa74a12c82f3bad960b0b000aca2ac329deea5c2328ebc6f2ba9802c1",
	"zero-leading repeated sequence": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
}

var invalidSHA256TestCases = map[string]string{
	"empty":                              "",
	"all uppercase":                      "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
	"mixed case":                         "E3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
	"one character short":                "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85",
	"one character long":                 "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b8550",
	"terminal non-hexadecimal character": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85g",
	"0x prefix":                          "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"leading space":                      " 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"trailing space":                     "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef ",
	"trailing newline":                   "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\n",
	"embedded hyphen":                    "01234567-9abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"embedded Unicode letter":            "01234567é9abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
}

func TestStringSHA256(t *testing.T) {
	rule := StringSHA256()
	t.Run("valid digests", func(t *testing.T) {
		for name, digest := range validSHA256TestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(digest))
			})
		}
	})
	t.Run("invalid digests", func(t *testing.T) {
		for name, digest := range invalidSHA256TestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(digest)
				assert.EqualError(t, err, "string must be a valid lowercase SHA-256 hexadecimal digest")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSHA256))
			})
		}
	})
}

func TestStringSHA256_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := stringNamedJSONSchemaCases(validSHA256TestCases, invalidSHA256TestCases)
	for line, digest := range loadNISTDigestOutputs(t, "nist_sha256_digest_outputs.txt", 729, 64) {
		cases = append(cases, jsonschematest.Case[string]{
			Name:  fmt.Sprintf("NIST/%d", line+1),
			Input: digest,
			Valid: true,
		})
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringSHA256())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_sha256.json", cases)
}

func BenchmarkStringSHA256(b *testing.B) {
	rule := StringSHA256()
	benchmarkStringHashDigestRule(b, rule, validSHA256TestCases, invalidSHA256TestCases)
}

var validSHA384TestCases = map[string]string{
	"digit-leading":                  "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b",
	"letter-leading":                 "b52b72da75d0666379e20f9b4a79c33a329a01f06a2fb7865c9062a28c1de860ba432edfd86b4cb1cb8a75b46076e3b1",
	"zero-leading repeated sequence": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
}

var invalidSHA384TestCases = map[string]string{
	"empty":                              "",
	"all uppercase":                      "38B060A751AC96384CD9327EB1B1E36A21FDB71114BE07434C0CC7BF63F6E1DA274EDEBFE76F65FBD51AD2F14898B95B",
	"mixed case":                         "38B060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b",
	"one character short":                "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95",
	"one character long":                 "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b0",
	"terminal non-hexadecimal character": "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95g",
	"0x prefix":                          "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"leading space":                      " 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"trailing space":                     "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef ",
	"trailing newline":                   "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\n",
	"embedded hyphen":                    "01234567-9abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"embedded Unicode letter":            "01234567é9abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
}

func TestStringSHA384(t *testing.T) {
	rule := StringSHA384()
	t.Run("valid digests", func(t *testing.T) {
		for name, digest := range validSHA384TestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(digest))
			})
		}
	})
	t.Run("invalid digests", func(t *testing.T) {
		for name, digest := range invalidSHA384TestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(digest)
				assert.EqualError(t, err, "string must be a valid lowercase SHA-384 hexadecimal digest")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSHA384))
			})
		}
	})
}

func TestStringSHA384_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := stringNamedJSONSchemaCases(validSHA384TestCases, invalidSHA384TestCases)
	for line, digest := range loadNISTDigestOutputs(t, "nist_sha384_digest_outputs.txt", 857, 96) {
		cases = append(cases, jsonschematest.Case[string]{
			Name:  fmt.Sprintf("NIST/%d", line+1),
			Input: digest,
			Valid: true,
		})
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringSHA384())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_sha384.json", cases)
}

func BenchmarkStringSHA384(b *testing.B) {
	rule := StringSHA384()
	benchmarkStringHashDigestRule(b, rule, validSHA384TestCases, invalidSHA384TestCases)
}

var validSHA512TestCases = map[string]string{
	"letter-leading":                 "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
	"digit-leading":                  "3831a6a6155e509dee59a7f451eb35324d8f8f2df6e3708894740f98fdee23889f4de5adb0c5010dfb555cda77c8ab5dc902094c52de3278f35a75ebc25f093a",
	"zero-leading repeated sequence": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
}

var invalidSHA512TestCases = map[string]string{
	"empty":                              "",
	"all uppercase":                      "CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E",
	"mixed case":                         "Cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
	"one character short":                "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3",
	"one character long":                 "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e0",
	"terminal non-hexadecimal character": "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3g",
	"0x prefix":                          "0x0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"leading space":                      " 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"trailing space":                     "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef ",
	"trailing newline":                   "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef\n",
	"embedded hyphen":                    "01234567-9abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	"embedded Unicode letter":            "01234567é9abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
}

func TestStringSHA512(t *testing.T) {
	rule := StringSHA512()
	t.Run("valid digests", func(t *testing.T) {
		for name, digest := range validSHA512TestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(digest))
			})
		}
	})
	t.Run("invalid digests", func(t *testing.T) {
		for name, digest := range invalidSHA512TestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(digest)
				assert.EqualError(t, err, "string must be a valid lowercase SHA-512 hexadecimal digest")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSHA512))
			})
		}
	})
}

func TestStringSHA512_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := stringNamedJSONSchemaCases(validSHA512TestCases, invalidSHA512TestCases)
	for line, digest := range loadNISTDigestOutputs(t, "nist_sha512_digest_outputs.txt", 857, 128) {
		cases = append(cases, jsonschematest.Case[string]{
			Name:  fmt.Sprintf("NIST/%d", line+1),
			Input: digest,
			Valid: true,
		})
	}

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringSHA512())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_sha512.json", cases)
}

func BenchmarkStringSHA512(b *testing.B) {
	rule := StringSHA512()
	benchmarkStringHashDigestRule(b, rule, validSHA512TestCases, invalidSHA512TestCases)
}

// The fixtures contain every distinct digest output from the NIST
// byte-oriented SHA test vector archive:
// https://csrc.nist.gov/CSRC/media/Projects/Cryptographic-Algorithm-Validation-Program/documents/shs/shabytetestvectors.zip
// The archive is pinned by SHA-256
// 929ef80b7b3418aca026643f6f248815913b60e01741a44bba9e118067f4c9b8.
// Short- and long-message vectors use validation-system revision 11.0
// (2011-03-15); Monte Carlo response and trace vectors use revision 11.1
// (2011-05-11). The fixtures include every distinct MD and MDi output value
// exactly once: 729 SHA-256, 857 SHA-384, and 857 SHA-512 values. M, Msg, and
// Seed fields are input material, not digest outputs. SHA-1, SHA-224,
// SHA-512/224, and SHA-512/256 are excluded because this package does not
// expose corresponding string rules.
func TestStringSHANISTDigestOutputs(t *testing.T) {
	tests := []struct {
		name          string
		fixture       string
		expectedCount int
		digestLength  int
		rule          govy.Rule[string]
	}{
		{
			name:          "SHA-256",
			fixture:       "nist_sha256_digest_outputs.txt",
			expectedCount: 729,
			digestLength:  64,
			rule:          StringSHA256(),
		},
		{
			name:          "SHA-384",
			fixture:       "nist_sha384_digest_outputs.txt",
			expectedCount: 857,
			digestLength:  96,
			rule:          StringSHA384(),
		},
		{
			name:          "SHA-512",
			fixture:       "nist_sha512_digest_outputs.txt",
			expectedCount: 857,
			digestLength:  128,
			rule:          StringSHA512(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			digests := loadNISTDigestOutputs(t, tc.fixture, tc.expectedCount, tc.digestLength)
			for line, digest := range digests {
				if err := tc.rule.Validate(digest); err != nil {
					t.Fatalf("testdata/%s:%d: digest rejected: %v", tc.fixture, line+1, err)
				}
			}
		})
	}
}

func TestStringHashDigestRulesMatchLowercaseHexadecimalLanguage(t *testing.T) {
	tests := []struct {
		name             string
		digestLength     int
		rule             govy.Rule[string]
		validTestCases   map[string]string
		invalidTestCases map[string]string
	}{
		{
			name:             "MD5",
			digestLength:     32,
			rule:             StringMD5(),
			validTestCases:   validMD5TestCases,
			invalidTestCases: invalidMD5TestCases,
		},
		{
			name:             "SHA-256",
			digestLength:     64,
			rule:             StringSHA256(),
			validTestCases:   validSHA256TestCases,
			invalidTestCases: invalidSHA256TestCases,
		},
		{
			name:             "SHA-384",
			digestLength:     96,
			rule:             StringSHA384(),
			validTestCases:   validSHA384TestCases,
			invalidTestCases: invalidSHA384TestCases,
		},
		{
			name:             "SHA-512",
			digestLength:     128,
			rule:             StringSHA512(),
			validTestCases:   validSHA512TestCases,
			invalidTestCases: invalidSHA512TestCases,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			oracle := regexp.MustCompile(fmt.Sprintf(`^[0-9a-f]{%d}$`, tc.digestLength))
			assertMatchesOracle := func(name, in string) {
				t.Helper()
				expected := oracle.MatchString(in)
				actual := tc.rule.Validate(in) == nil
				if actual != expected {
					t.Fatalf("%s: validation result is %t; want %t for %q", name, actual, expected, in)
				}
			}

			for name, in := range tc.validTestCases {
				assertMatchesOracle("valid/"+name, in)
			}
			for name, in := range tc.invalidTestCases {
				assertMatchesOracle("invalid/"+name, in)
			}

			for _, length := range []int{
				0,
				1,
				tc.digestLength - 1,
				tc.digestLength,
				tc.digestLength + 1,
				2 * tc.digestLength,
			} {
				assertMatchesOracle(fmt.Sprintf("length/%d", length), strings.Repeat("0", length))
			}

			digest := []byte(strings.Repeat("0", tc.digestLength))
			for position := range len(digest) {
				for value := range 256 {
					digest[position] = byte(value)
					assertMatchesOracle(
						fmt.Sprintf("position/%d/byte/%d", position, value),
						string(digest),
					)
				}
				digest[position] = '0'
			}
		})
	}
}

func benchmarkStringHashDigestRule(
	b *testing.B,
	rule govy.Rule[string],
	validTestCases map[string]string,
	invalidTestCases map[string]string,
) {
	b.Helper()
	for b.Loop() {
		for _, in := range validTestCases {
			_ = rule.Validate(in)
		}
		for _, in := range invalidTestCases {
			_ = rule.Validate(in)
		}
	}
}

// RFC 7519 sections 3.1, 6.1, A.1, and A.2 provide the four complete JWT
// examples: https://www.rfc-editor.org/rfc/rfc7519.html.
// The b64=false case comes from the immutable RFC 7797 section 7 prohibition:
// https://www.rfc-editor.org/rfc/rfc7797.html#section-7.
const stringJWTJSONSchemaDifference = "JSON Schema's contentMediaType is an annotation and does not validate JWT contents."

type stringJWTTestCase struct {
	in                   string
	expectedErrorDetails string
	jsonSchemaDifference string
}

var stringJWTTestCases = map[string]stringJWTTestCase{
	"signed token": {
		in: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
			"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ." +
			"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	},
	"RFC 7519 signed token": {
		in: "eyJ0eXAiOiJKV1QiLA0KICJhbGciOiJIUzI1NiJ9." +
			"eyJpc3MiOiJqb2UiLA0KICJleHAiOjEzMDA4MTkzODAsDQogImh0dHA6Ly9leGFtcGxlLmNvbS9pc19yb290Ijp0cnVlfQ." +
			"dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk",
	},
	"RFC 7519 unsecured token": {
		in: "eyJhbGciOiJub25lIn0." +
			"eyJpc3MiOiJqb2UiLA0KICJleHAiOjEzMDA4MTkzODAsDQogImh0dHA6Ly9leGFtcGxlLmNvbS9pc19yb290Ijp0cnVlfQ.",
	},
	"whitespace and Unicode JSON": {
		in: "ew0KICJhbGciIDogIm5vbmUiLA0KICJraWQiIDogIs66zrvOtc65zrTOryINCn0." +
			"ewogIm5hbWUiOiAiSsO2aG4g6ZuqIiwKICJhZG1pbiI6IHRydWUKfQ.",
	},
	"minimal unsecured token": {
		in: "eyJhbGciOiJub25lIn0.e30.",
	},
	"empty token": {
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "expected exactly 3 JWT segments",
	},
	"one segment": {
		in:                   "not-a-jwt",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "expected exactly 3 JWT segments",
	},
	"two segments": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.e30",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "expected exactly 3 JWT segments",
	},
	"four segments": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.e30.c2ln.ZXh0cmE",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "expected exactly 3 JWT segments",
	},
	"RFC 7519 encrypted JWE": {
		in: "eyJhbGciOiJSU0ExXzUiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2In0." +
			"QR1Owv2ug2WyPBnbQrRARTeEk9kDO2w8qDcjiHnSJflSdv1iNqhWXaKH4MqAkQtM" +
			"oNfABIPJaZm0HaA415sv3aeuBWnD8J-Ui7Ah6cWafs3ZwwFKDFUUsWHSK-IPKxLG" +
			"TkND09XyjORj_CHAgOPJ-Sd8ONQRnJvWn_hXV1BNMHzUjPyYwEsRhDhzjAD26ima" +
			"sOTsgruobpYGoQcXUwFDn7moXPRfDE8-NoQX7N7ZYMmpUDkR-Cx9obNGwJQ3nM52" +
			"YCitxoQVPzjbl7WBuB7AohdBoZOdZ24WlN1lVIeh8v1K4krB8xgKvRU8kgFrEn_a" +
			"1rZgN5TiysnmzTROF869lQ." +
			"AxY8DCtDaGlsbGljb3RoZQ." +
			"MKOle7UQrG6nSxTLX6Mqwt0orbHvAKeWnDYvpIAeZ72deHxz3roJDXQyhxx0wKaM" +
			"HDjUEOKIwrtkHthpqEanSBNYHZgmNOV7sln1Eu9g3J8." +
			"fiK51VwhsxJ-siBMR-YFiA",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "expected exactly 3 JWT segments",
	},
	"RFC 7519 nested JWT": {
		in: "eyJhbGciOiJSU0ExXzUiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiY3R5IjoiSldU" +
			"In0." +
			"g_hEwksO1Ax8Qn7HoN-BVeBoa8FXe0kpyk_XdcSmxvcM5_P296JXXtoHISr_DD_M" +
			"qewaQSH4dZOQHoUgKLeFly-9RI11TG-_Ge1bZFazBPwKC5lJ6OLANLMd0QSL4fYE" +
			"b9ERe-epKYE3xb2jfY1AltHqBO-PM6j23Guj2yDKnFv6WO72tteVzm_2n17SBFvh" +
			"DuR9a2nHTE67pe0XGBUS_TK7ecA-iVq5COeVdJR4U4VZGGlxRGPLRHvolVLEHx6D" +
			"YyLpw30Ay9R6d68YCLi9FYTq3hIXPK_-dmPlOUlKvPr1GgJzRoeC9G5qCvdcHWsq" +
			"JGTO_z3Wfo5zsqwkxruxwA." +
			"UmVkbW9uZCBXQSA5ODA1Mg." +
			"VwHERHPvCNcHHpTjkoigx3_ExK0Qc71RMEParpatm0X_qpg-w8kozSjfNIPPXiTB" +
			"BLXR65CIPkFqz4l1Ae9w_uowKiwyi9acgVztAi-pSL8GQSXnaamh9kX1mdh3M_TT" +
			"-FZGQFQsFhu0Z72gJKGdfGE-OE7hS1zuBD5oEUfk0Dmb0VzWEzpxxiSSBbBAzP10" +
			"l56pPfAtrjEYw-7ygeMkwBl6Z_mLS6w6xUgKlvW6ULmkV-uLC4FUiyKECK4e3WZY" +
			"Kw1bpgIqGYsw2v_grHjszJZ-_I5uM-9RA8ycX9KqPRp9gc6pXmoU_-27ATs9XCvr" +
			"ZXUtK2902AUzqpeEUJYjWWxSNsS-r1TJ1I-FMJ4XyAiGrfmo9hQPcNBYxPz3GQb2" +
			"8Y5CLSQfNgKSGt0A4isp1hBUXBHAndgtcslt7ZoQJaKe_nNJgNliWtWpJ_ebuOpE" +
			"l8jdhehdccnRMIwAmU1n7SPkmhIl1HlSOpvcvDfhUN5wuqU955vOBvfkBOh5A11U" +
			"zBuo2WlgZ6hYi9-e3w29bR0C2-pp3jbqxEDw3iWaf2dc5b-LnR0FEYXvI_tYk5rd" +
			"_J9N0mg0tQ6RbpxNEMNoA9QWk5lgdPvbh9BaO195abQ." +
			"AVO9iT5AV4CzvDJCdhSFlQ",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "expected exactly 3 JWT segments",
	},
	"empty header segment": {
		in:                   ".e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must not be empty",
	},
	"empty claims set segment": {
		in:                   "eyJhbGciOiJIUzI1NiJ9..c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must not be empty",
	},
	"padded header segment": {
		in:                   "eyJhbGciOiJIUzI1NiJ9=.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must be base64url encoded without padding",
	},
	"padded claims set segment": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.e30=.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must be base64url encoded without padding",
	},
	"padded signature segment": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.e30.c2ln=",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT signature segment must be base64url encoded without padding",
	},
	"illegal alphabet in header segment": {
		in:                   "*.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must be base64url encoded without padding",
	},
	"illegal alphabet in claims set segment": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.*.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must be base64url encoded without padding",
	},
	"illegal alphabet in signature segment": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.e30.*",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT signature segment must be base64url encoded without padding",
	},
	"impossible base64url length in header segment": {
		in:                   "A.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must be base64url encoded without padding: " +
			"illegal base64 data at input byte 0",
	},
	"impossible base64url length in claims set segment": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.A.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must be base64url encoded without padding: " +
			"illegal base64 data at input byte 0",
	},
	"impossible base64url length in signature segment": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.e30.A",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT signature segment must be base64url encoded without padding: " +
			"illegal base64 data at input byte 0",
	},
	"malformed header JSON": {
		in:                   "eyJhbGciOiJIUzI1NiI.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must contain a JSON object: " +
			"unexpected end of JSON input",
	},
	"malformed claims set JSON": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOg.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must contain a JSON object: " +
			"unexpected end of JSON input",
	},
	"header segment is JSON array": {
		in:                   "W10.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must contain a JSON object: " +
			"json: cannot unmarshal array into Go value of type map[string]json.RawMessage",
	},
	"claims set segment is JSON array": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.W10.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must contain a JSON object: " +
			"json: cannot unmarshal array into Go value of type map[string]json.RawMessage",
	},
	"header segment is JSON null": {
		in:                   "bnVsbA.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must contain a JSON object",
	},
	"claims set segment is JSON null": {
		in: "eyJhbGciOiJIUzI1NiJ9." +
			"bnVsbA." +
			"c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must contain a JSON object",
	},
	"missing algorithm": {
		in:                   "e30.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must contain an "alg" string`,
	},
	"empty algorithm": {
		in:                   "eyJhbGciOiIifQ.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must contain an "alg" string`,
	},
	"null algorithm": {
		in:                   "eyJhbGciOm51bGx9.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must contain an "alg" string`,
	},
	"numeric algorithm": {
		in:                   "eyJhbGciOjEyM30.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must contain an "alg" string`,
	},
	"non-ASCII algorithm": {
		in:                   "eyJhbGciOiLimIMifQ.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must contain an "alg" string`,
	},
	"missing signature for signed token": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.e30.",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT signature segment must not be empty unless alg is "none"`,
	},
	"signature present for none algorithm": {
		in:                   "eyJhbGciOiJub25lIn0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT signature segment must be empty when alg is "none"`,
	},
	"leading token whitespace": {
		in:                   " eyJhbGciOiJIUzI1NiJ9.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must be base64url encoded without padding",
	},
	"raw Unicode signature": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.e30.雪",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT signature segment must be base64url encoded without padding",
	},
	"invalid UTF-8 in header JSON": {
		in:                   "eyJhbGciOiJIUzI1NiIsIngiOiL_In0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT header segment must contain valid UTF-8 JSON",
	},
	"invalid UTF-8 in claims set JSON": {
		in:                   "eyJhbGciOiJIUzI1NiJ9.eyJ4Ijoi_yJ9.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must contain valid UTF-8 JSON",
	},
	"derived encoded payload option": {
		in: "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6WyJiNjQiXX0." +
			"e30.c2ln",
	},
	"RFC 7797 unencoded payload option": {
		in: "eyJhbGciOiJIUzI1NiIsImI2NCI6ZmFsc2UsImNyaXQiOlsiYjY0Il19." +
			"e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must not set "b64" to false`,
	},
	"null b64 option": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6bnVsbH0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "b64" must be a boolean`,
	},
	"string b64 option": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6ImZhbHNlIn0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "b64" must be a boolean`,
	},
	"number b64 option": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6MH0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "b64" must be a boolean`,
	},
	"object b64 option": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6e319.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "b64" must be a boolean`,
	},
	"array b64 option": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6W119.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "b64" must be a boolean`,
	},
	"b64 option without crit": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZX0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with null crit": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6bnVsbH0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with string crit": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6ImI2NCJ9.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with number crit": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6MH0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with object crit": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6e319.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with empty crit": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6W119.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with unrelated crit": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6WyJleHAiXX0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with non-string crit member": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6WzBdfQ.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with null crit member": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6W251bGxdfQ.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with boolean crit member": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6W3RydWVdfQ.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with object crit member": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6W3t9XX0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with array crit member": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6W1tdXX0.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"b64 option with mixed crit members": {
		in:                   "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZSwiY3JpdCI6WyJiNjQiLDBdfQ.e30.c2ln",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
}

func TestStringJWT(t *testing.T) {
	const expectedErrorPrefix = "string must be a valid JSON Web Token (JWT): "

	for name, tt := range stringJWTTestCases {
		t.Run(name, func(t *testing.T) {
			err := StringJWT().Validate(tt.in)
			if tt.expectedErrorDetails != "" {
				assert.Require(t, assert.Error(t, err))
				assert.EqualError(t, err, expectedErrorPrefix+tt.expectedErrorDetails)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringJWT))
				return
			}
			assert.NoError(t, err)
		})
	}
}

var stringJWTErrorPrecedenceTestCases = map[string]stringJWTTestCase{
	"algorithm before claims": {
		in:                   "e30.eyJzdWIiOg.*",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must contain an "alg" string`,
	},
	"b64 before claims": {
		in: "eyJhbGciOiJIUzI1NiIsImI2NCI6ZmFsc2UsImNyaXQiOlsiYjY0Il19." +
			"eyJzdWIiOg.*",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must not set "b64" to false`,
	},
	"crit before claims": {
		in: "eyJhbGciOiJIUzI1NiIsImI2NCI6dHJ1ZX0." +
			"eyJzdWIiOg.*",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header "crit" must be an array containing "b64" when "b64" is present`,
	},
	"claims before signature": {
		in: "eyJhbGciOiJIUzI1NiJ9." +
			"eyJzdWIiOg.*",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: "JWT claims set segment must contain a JSON object: " +
			"unexpected end of JSON input",
	},
	"algorithm before signature": {
		in:                   "e30.e30.*",
		jsonSchemaDifference: stringJWTJSONSchemaDifference,
		expectedErrorDetails: `JWT header must contain an "alg" string`,
	},
}

func TestStringJWTErrorPrecedence(t *testing.T) {
	const expectedErrorPrefix = "string must be a valid JSON Web Token (JWT): "

	for name, tt := range stringJWTErrorPrecedenceTestCases {
		t.Run(name, func(t *testing.T) {
			err := StringJWT().Validate(tt.in)
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, expectedErrorPrefix+tt.expectedErrorDetails)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringJWT))
		})
	}
}

func TestStringJWT_JSONSchema(t *testing.T) {
	t.Parallel()

	cases := make([]jsonschematest.Case[string], 0, len(stringJWTTestCases)+len(stringJWTErrorPrecedenceTestCases))
	for group, tests := range map[string]map[string]stringJWTTestCase{
		"tokens":           stringJWTTestCases,
		"error precedence": stringJWTErrorPrecedenceTestCases,
	} {
		for name, tc := range tests {
			cases = append(cases, jsonschematest.Case[string]{
				Name:                 group + "/" + name,
				Input:                tc.in,
				Valid:                tc.expectedErrorDetails == "",
				JSONSchemaDifference: tc.jsonSchemaDifference,
			})
		}
	}
	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringJWT())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_jwt.json", cases)
}

func BenchmarkStringJWT(b *testing.B) {
	rule := StringJWT()

	for b.Loop() {
		for _, tt := range stringJWTTestCases {
			_ = rule.Validate(tt.in)
		}
		for _, tt := range stringJWTErrorPrecedenceTestCases {
			_ = rule.Validate(tt.in)
		}
	}
	b.ReportMetric(float64(len(stringJWTTestCases)+len(stringJWTErrorPrecedenceTestCases)), "validations/op")
}

var stringContainsTestCases = []*struct {
	in            string
	substrings    []string
	expectedError string
}{
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
	{in: "", substrings: []string{"his"}, expectedError: "string must contain the following substrings: 'his'"},
	{in: "is", substrings: []string{"this"}, expectedError: "string must contain the following substrings: 'this'"},
	{
		in:            "th",
		substrings:    []string{"th", "is"},
		expectedError: "string must contain the following substrings: 'th', 'is'",
	},
	{in: "this", substrings: []string{"th"}},
	{in: "th ht", substrings: []string{"th", "ht"}},
	{in: "that", substrings: []string{"that"}},
	{in: "a t.h z", substrings: []string{"t.h"}},
	{in: "a tXh z", substrings: []string{"t.h"}, expectedError: "string must contain the following substrings: 't.h'"},
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
	t.Run("panic if substrings are empty", func(t *testing.T) {
		assert.Panic(t, func() { StringContains() }, "substrings must not be empty")
	})
	t.Run("panic if a substring is empty", func(t *testing.T) {
		assert.Panic(t,
			func() { StringContains("value", "") },
			"substrings must not contain empty strings")
	})
}

func TestStringContains_JSONSchema(t *testing.T) {
	t.Parallel()
	cases := make([]stringTextArgumentCase, 0, len(stringContainsTestCases))
	for _, tc := range stringContainsTestCases {
		cases = append(cases, stringTextArgumentCase{tc.substrings, tc.in, tc.expectedError == ""})
	}
	assertStringArgumentsJSONSchema(t, "string_contains", StringContains, cases)
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
	{
		in:            "this",
		substrings:    []string{"th"},
		expectedError: "string must not contain any of the following substrings: 'th'",
	},
	{
		in:            "tho",
		substrings:    []string{"tho", "ht"},
		expectedError: "string must not contain any of the following substrings: 'tho', 'ht'",
	},
	{
		in:            "that",
		substrings:    []string{"that"},
		expectedError: "string must not contain any of the following substrings: 'that'",
	},
	{in: "one", substrings: []string{"his"}},
	{in: "one", substrings: []string{"this"}},
	{in: "one", substrings: []string{"th", "is"}},
	{in: "a tXh z", substrings: []string{"t.h"}},
	{
		in:            "a t.h z",
		substrings:    []string{"t.h"},
		expectedError: "string must not contain any of the following substrings: 't.h'",
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
	t.Run("panic if substrings are empty", func(t *testing.T) {
		assert.Panic(t, func() { StringExcludes() }, "substrings must not be empty")
	})
	t.Run("panic if a substring is empty", func(t *testing.T) {
		assert.Panic(t,
			func() { StringExcludes("value", "") },
			"substrings must not contain empty strings")
	})
}

func TestStringExcludes_JSONSchema(t *testing.T) {
	t.Parallel()
	cases := make([]stringTextArgumentCase, 0, len(stringExcludesTestCases))
	for _, tc := range stringExcludesTestCases {
		cases = append(cases, stringTextArgumentCase{tc.substrings, tc.in, tc.expectedError == ""})
	}
	assertStringArgumentsJSONSchema(t, "string_excludes", StringExcludes, cases)
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
	{
		in:            "one",
		prefixes:      []string{"is", "th"},
		expectedError: "string must start with one of the following prefixes: 'is', 'th'",
	},
	{in: "this", prefixes: []string{"th", "ht"}},
	{in: ".this", prefixes: []string{".t"}},
	{in: "xthis", prefixes: []string{".t"}, expectedError: "string must start with '.t' prefix"},
	{in: "a\nthis", prefixes: []string{"th"}, expectedError: "string must start with 'th' prefix"},
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
	t.Run("panic if prefixes are empty", func(t *testing.T) {
		assert.Panic(t, func() { StringStartsWith() }, "prefixes must not be empty")
	})
	t.Run("panic if a prefix is empty", func(t *testing.T) {
		assert.Panic(t,
			func() { StringStartsWith("prefix", "") },
			"prefixes must not contain empty strings")
	})
}

func TestStringStartsWith_JSONSchema(t *testing.T) {
	t.Parallel()
	cases := make([]stringTextArgumentCase, 0, len(stringStartsWithTestCases))
	for _, tc := range stringStartsWithTestCases {
		cases = append(cases, stringTextArgumentCase{tc.prefixes, tc.in, tc.expectedError == ""})
	}
	assertStringArgumentsJSONSchema(t, "string_starts_with", StringStartsWith, cases)
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
	{in: "this\n", suffixes: []string{"is"}, expectedError: "string must end with 'is' suffix"},
	{
		in:            "one",
		suffixes:      []string{"th", "is"},
		expectedError: "string must end with one of the following suffixes: 'th', 'is'",
	},
	{in: "with", suffixes: []string{"th"}},
	{in: "light", suffixes: []string{"th", "ht"}},
	{in: "x.h", suffixes: []string{".h"}},
	{in: "xxh", suffixes: []string{".h"}, expectedError: "string must end with '.h' suffix"},
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
	t.Run("panic if suffixes are empty", func(t *testing.T) {
		assert.Panic(t, func() { StringEndsWith() }, "suffixes must not be empty")
	})
	t.Run("panic if a suffix is empty", func(t *testing.T) {
		assert.Panic(t,
			func() { StringEndsWith("suffix", "") },
			"suffixes must not contain empty strings")
	})
}

func TestStringEndsWith_JSONSchema(t *testing.T) {
	t.Parallel()
	cases := make([]stringTextArgumentCase, 0, len(stringEndsWithTestCases))
	for _, tc := range stringEndsWithTestCases {
		cases = append(cases, stringTextArgumentCase{tc.suffixes, tc.in, tc.expectedError == ""})
	}
	assertStringArgumentsJSONSchema(t, "string_ends_with", StringEndsWith, cases)
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
	in                   string
	expectedErr          error
	jsonSchemaDifference string
}{
	{"refs/heads/master", nil, ""},
	{"refs/notes/commits", nil, ""},
	{"refs/tags/this@", nil, ""},
	{"refs/remotes/origin/master", nil, ""},
	{"HEAD", nil, ""},
	{"refs/tags/v3.1.1", nil, ""},
	{"refs/pulls/1/head", nil, ""},
	{"refs/pulls/1/merge", nil, ""},
	{"refs/pulls/1/abc.123", nil, ""},
	{"refs/pulls", nil, ""},
	{"refs/-", nil, ""},
	{"refs", errGitRefAtLeastOneSlash, ""},
	{"refs/", errGitRefEmptyPart, ""},
	{"refs//", errGitRefEmptyPart, ""},
	{"refs/heads/\\", errGitRefForbiddenChars, ""},
	{"refs/heads/\\foo", errGitRefForbiddenChars, ""},
	{"refs/heads/\\foo/bar", errGitRefForbiddenChars, ""},
	{"abc", errGitRefAtLeastOneSlash, ""},
	{"", errGitRefEmpty, ""},
	{"refs/heads/ ", errGitRefForbiddenChars, ""},
	{"refs/heads/ /", errGitRefForbiddenChars, ""},
	{"refs/heads/ /foo", errGitRefForbiddenChars, ""},
	{"refs/heads/.", errGitRefEndsWithDot, "JSON Schema does not reject a trailing dot."},
	{"refs/heads/..", errGitRefEndsWithDot, "JSON Schema does not reject a trailing dot."},
	{"refs/heads/foo..", errGitRefEndsWithDot, "JSON Schema does not reject a trailing dot."},
	{"refs/heads/foo.lock", errGitRefForbiddenChars, "JSON Schema does not reject components that end with .lock."},
	{"refs/heads/foo@{bar}", errGitRefForbiddenChars, "JSON Schema does not reject the @{ sequence."},
	{"refs/heads/foo@{", errGitRefForbiddenChars, "JSON Schema does not reject the @{ sequence."},
	{"refs/heads/foo[", errGitRefForbiddenChars, ""},
	{"refs/heads/foo~", errGitRefForbiddenChars, ""},
	{"refs/heads/foo^", errGitRefForbiddenChars, ""},
	{"refs/heads/foo:", errGitRefForbiddenChars, ""},
	{"refs/heads/foo?", errGitRefForbiddenChars, ""},
	{"refs/heads/foo*", errGitRefForbiddenChars, ""},
	{"refs/heads/foo[bar", errGitRefForbiddenChars, ""},
	{"refs/heads/foo\t", errGitRefForbiddenChars, ""},
	{"refs/heads/@", errGitRefForbiddenChars, "JSON Schema permits @ as a complete component."},
	{"refs/heads/@{bar}", errGitRefForbiddenChars, "JSON Schema does not reject the @{ sequence."},
	{"refs/heads/\n", errGitRefForbiddenChars, ""},
	{
		"refs/heads/-foo",
		errGitRefStartsWithDash,
		"JSON Schema does not reject leading dashes in branch or tag components.",
	},
	{"refs/heads/foo..bar", errGitRefForbiddenChars, "JSON Schema does not reject the .. sequence."},
	{
		"refs/heads/-",
		errGitRefStartsWithDash,
		"JSON Schema does not reject leading dashes in branch or tag components.",
	},
	{"refs/tags/-", errGitRefStartsWithDash, "JSON Schema does not reject leading dashes in branch or tag components."},
	{
		"refs/tags/-foo",
		errGitRefStartsWithDash,
		"JSON Schema does not reject leading dashes in branch or tag components.",
	},
	{"refs/heads/.hidden", errGitRefForbiddenChars, "JSON Schema does not reject components that start with a dot."},
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

func TestStringGitRef_JSONSchema(t *testing.T) {
	t.Parallel()
	cases := make([]jsonschematest.Case[string], 0, len(stringGitRefTestCases))
	for _, tc := range stringGitRefTestCases {
		cases = append(cases, jsonschematest.Case[string]{
			Name: strconv.Quote(tc.in), Input: tc.in, Valid: tc.expectedErr == nil,
			JSONSchemaDifference: tc.jsonSchemaDifference,
		})
	}
	assertStringTextJSONSchema(t, StringGitRef(), "expected_string_git_ref.json", cases)
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

type stringAbsoluteFilePathTestCase struct {
	in           string
	isAbs        bool
	isAbsWindows bool
}

// The source cases include all isabstests and winisabstests inputs from Go 1.26.5:
// https://github.com/golang/go/blob/go1.26.5/src/path/filepath/path_test.go#L1062-L1090.
// The expected result depends on the host operating system.
func getStringAbsoluteFilePathTestCases(root string) []stringAbsoluteFilePathTestCase {
	return []stringAbsoluteFilePathTestCase{
		// Go's isabstests.
		{"", false, false},
		{"/", true, false},
		{"/usr/bin/gcc", true, false},
		{"..", false, false},
		{"/a/../bb", true, false},
		{".", false, false},
		{"./", false, false},
		{"lala", false, false},
		// Go's winisabstests.
		{`C:\`, false, true},
		{`c\`, false, false},
		{`c::`, false, false},
		{`c:`, false, false},
		{`/`, true, false},
		{`\`, false, false},
		{`\Windows`, false, false},
		{`c:a\b`, false, false},
		{`c:\a\b`, false, true},
		{`c:/a/b`, false, true},
		{`\\host\share`, false, true},
		{`\\host\share\`, false, true},
		{`\\host\share\foo`, false, true},
		{`//host/share/foo/bar`, true, true},
		{`\\?\a\b\c`, false, true},
		{`\??\a\b\c`, false, true},
		// Go's TestIsAbs also prefixes each isabstests input with "c:".
		{"c:", false, false},
		{"c:/", false, true},
		{"c:/usr/bin/gcc", false, true},
		{"c:..", false, false},
		{"c:/a/../bb", false, true},
		{"c:.", false, false},
		{"c:./", false, false},
		{"c:lala", false, false},
		// Additional relative paths, whitespace, and paths that need no normalization.
		{"config.yaml", false, false},
		{"dir/config.yaml", false, false},
		{"./config.yaml", false, false},
		{"../config.yaml", false, false},
		{"~/config.yaml", false, false},
		{"$HOME/config.yaml", false, false},
		{"%USERPROFILE%\\config.yaml", false, false},
		{" ", false, false},
		{"\t", false, false},
		{" /config.yaml", false, false},
		{"/dir/./config.yaml", true, false},
		{"/dir//config.yaml", true, false},
		{"/dir with spaces/config.yaml", true, false},
		{"/dir/", true, false},
		{`C:\dir\..\config.yaml`, false, true},
		{`C:\dir with spaces\config.yaml`, false, true},
		{filepath.Join(root, "missing", "config.yaml"), true, true},
	}
}

func TestStringAbsoluteFilePath(t *testing.T) {
	rule := StringAbsoluteFilePath()
	customRule := rule.WithMessageTemplateString("absolute path required: '{{ .PropertyValue }}'")
	for _, tc := range getStringAbsoluteFilePathTestCases(t.TempDir()) {
		t.Run(fmt.Sprintf("%q", tc.in), func(t *testing.T) {
			isAbs := tc.isAbs
			if runtime.GOOS == "windows" {
				isAbs = tc.isAbsWindows
			}
			err := rule.Validate(tc.in)
			if isAbs {
				assert.NoError(t, err)
			} else {
				assert.Require(t, assert.EqualError(t, err, "string must be an absolute file path"))
				assert.Require(t, assert.IsType[*govy.RuleError](t, err))
				assert.Equal(t, "string must be an absolute file path", err.(*govy.RuleError).Description)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAbsoluteFilePath))
				err = customRule.Validate(tc.in)
				assert.EqualError(t, err, fmt.Sprintf("absolute path required: '%s'", tc.in))
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAbsoluteFilePath))
			}
		})
	}
}

func BenchmarkStringAbsoluteFilePath(b *testing.B) {
	rule := StringAbsoluteFilePath()
	testCases := getStringAbsoluteFilePathTestCases(b.TempDir())
	for b.Loop() {
		for _, tc := range testCases {
			_ = rule.Validate(tc.in)
		}
	}
}

// test cases copied from Go's [filepath] standard library.
var stringMatchFileSystemPathTestCases = []*struct {
	pattern, in string
	shouldFail  bool
}{
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
	layout               string
	in                   string
	errMsg               string
	jsonSchemaDifference string
}{
	{time.RFC3339, "2024-01-01T15:00:00Z", "", ""},
	{time.RFC3339, "2024-01-01T15:00:00+01:00", "", ""},
	{time.RFC3339, "2024-02-29T15:00:00Z", "", ""},
	{time.RFC3339Nano, "2024-01-01T15:00:00.123456789Z", "", ""},
	{time.RFC3339Nano, "2024-01-01T15:00:00Z", "", ""},
	{time.DateTime, "2024-01-01 15:00:00", "", ""},
	{time.DateOnly, "2024-01-01", "", ""},
	{time.TimeOnly, "15:00:00", "", ""},
	{
		"invalid-layout",
		"2024-01-01T15:00:00Z",
		"string must be a valid date and time in 'invalid-layout' format",
		"JSON Schema does not constrain unsupported Go time layouts.",
	},
	{
		time.RFC3339,
		"2024-01-01 15:00:00Z",
		"string must be a valid date and time in '2006-01-02T15:04:05Z07:00' format",
		"Ajv accepts a space between the date and time; Go requires T for this layout.",
	},
	{
		time.RFC3339,
		"2024-01-01t15:00:00z",
		"string must be a valid date and time in '2006-01-02T15:04:05Z07:00' format",
		"Ajv accepts lowercase t and z; Go requires uppercase T and Z for this layout.",
	},
	{
		time.RFC3339,
		"2016-12-31T23:59:60Z",
		"string must be a valid date and time in '2006-01-02T15:04:05Z07:00' format",
		"Ajv accepts leap seconds; Go time.Parse rejects them.",
	},
	{
		time.RFC3339,
		"2023-02-29T15:00:00Z",
		"string must be a valid date and time in '2006-01-02T15:04:05Z07:00' format",
		"",
	},
	{
		time.RFC3339Nano,
		"2024-02-30T15:00:00.123456789Z",
		"string must be a valid date and time in '2006-01-02T15:04:05.999999999Z07:00' format",
		"",
	},
	{
		"15:04",
		"15:00:00",
		"string must be a valid date and time in '15:04'",
		"JSON Schema does not constrain unsupported Go time layouts.",
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

func TestStringDateTime_JSONSchema(t *testing.T) {
	t.Parallel()
	casesByLayout := make(map[string][]jsonschematest.Case[string])
	for _, tc := range stringDateTimeTestCases {
		casesByLayout[tc.layout] = append(casesByLayout[tc.layout], jsonschematest.Case[string]{
			Name: strconv.Quote(tc.in), Input: tc.in, Valid: tc.errMsg == "",
			JSONSchemaDifference: tc.jsonSchemaDifference,
		})
	}
	for layout, cases := range casesByLayout {
		t.Run(layout, func(t *testing.T) {
			t.Parallel()
			fixture := "expected_string_date_time_unsupported.json"
			if layout == time.RFC3339 || layout == time.RFC3339Nano {
				fixture = "expected_string_date_time.json"
			}
			assertStringTextJSONSchema(t, StringDateTime(layout), fixture, cases)
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
	{"", false},
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
	{"test\n", true},
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

func TestStringAlpha_JSONSchema(t *testing.T) {
	t.Parallel()
	assertStringTextJSONSchema(t, StringAlpha(), "expected_string_alpha.json",
		stringCharacterJSONSchemaCases(stringAlphaTestCases))
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
	{"", false},
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
	{"test1\n", true},
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

func TestStringAlphanumeric_JSONSchema(t *testing.T) {
	t.Parallel()
	assertStringTextJSONSchema(t, StringAlphanumeric(), "expected_string_alphanumeric.json",
		stringCharacterJSONSchemaCases(stringAlphanumericTestCases))
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
	{"", false},
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
	{"𐐀", false},
	{"汉语\n", true},
	{"e\u0301", true},
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

func TestStringAlphaUnicode_JSONSchema(t *testing.T) {
	t.Parallel()
	assertStringTextJSONSchema(t, StringAlphaUnicode(), "expected_string_alpha_unicode.json",
		stringCharacterJSONSchemaCases(stringAlphaUnicodeTestCases))
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
	{"", true},
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
	{"𐐀𐒠", false},
	{"汉语1\n", true},
	{"e\u0301", true},
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

func TestStringAlphanumericUnicode_JSONSchema(t *testing.T) {
	t.Parallel()
	assertStringTextJSONSchema(t, StringAlphanumericUnicode(), "expected_string_alphanumeric_unicode.json",
		stringCharacterJSONSchemaCases(stringAlphanumericUnicodeTestCases))
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

func TestStringFQDN_JSONSchema(t *testing.T) {
	t.Parallel()
	assertStringTextJSONSchema(t, StringFQDN(), "expected_string_fqdn.json",
		stringCharacterJSONSchemaCases(stringFQDNTestCases))
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
	in                   string
	expectedErr          error
	jsonSchemaDifference string
}{
	{"simple", nil, ""},
	{"now-with-dashes", nil, ""},
	{"1-starts-with-num", nil, ""},
	{"1234", nil, ""},
	{"simple/simple", nil, ""},
	{"now-with-dashes/simple", nil, ""},
	{"now-with-dashes/now-with-dashes", nil, ""},
	{"now.with.dots/simple", nil, ""},
	{"now-with.dashes-and.dots/simple", nil, ""},
	{"1-num.2-num/3-num", nil, ""},
	{"1234/5678", nil, ""},
	{"1.2.3.4/5678", nil, ""},
	{"Uppercase_Is_OK_123", nil, ""},
	{"example.com/Uppercase_Is_OK_123", nil, ""},
	{"requests.storage-foo", nil, ""},
	{strings.Repeat("a", 63), nil, ""},
	{strings.Repeat("a", 253) + "/" + strings.Repeat("b", 63), nil, ""},
	// BAD
	{"/", errK8sQualifiedNameEmptyPrefixPart, ""},
	{"nospecialchars%^=@", errK8sQualifiedNameNamePartRegexp, ""},
	{"cantendwithadash-", errK8sQualifiedNameNamePartRegexp, ""},
	{"-cantstartwithadash-", errK8sQualifiedNameNamePartRegexp, ""},
	{"example.com/abc$", errK8sQualifiedNameNamePartRegexp, ""},
	{"only/one/slash", errK8sQualifiedNameTooManyParts, ""},
	{"Example.com/abc", errK8sQualifiedNamePrefixRegexp, ""},
	{"example_com/abc", errK8sQualifiedNamePrefixRegexp, ""},
	{"example.com/", errK8sQualifiedNameEmptyNamePart, ""},
	{"/simple", errK8sQualifiedNameEmptyPrefixPart, ""},
	{"not.Valid/simple", errK8sQualifiedNamePrefixRegexp, ""},
	{
		strings.Repeat("a", 64),
		errK8sQualifiedNameNamePartLength,
		"JSON Schema limits total length but not the name part separately.",
	},
	{
		strings.Repeat("a", 254) + "/abc",
		errK8sQualifiedNamePrefixLength,
		"JSON Schema limits total length but not the prefix separately.",
	},
	{strings.Repeat("a", 253) + "/" + strings.Repeat("b", 64), errors.New("length must be between 1 and 317"), ""},
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

func TestStringKubernetesQualifiedName_JSONSchema(t *testing.T) {
	t.Parallel()
	cases := make([]jsonschematest.Case[string], 0, len(stringK8sQualifiedNameTestCases))
	for _, tc := range stringK8sQualifiedNameTestCases {
		cases = append(cases, jsonschematest.Case[string]{
			Name: strconv.Quote(tc.in), Input: tc.in, Valid: tc.expectedErr == nil,
			JSONSchemaDifference: tc.jsonSchemaDifference,
		})
	}
	assertStringTextJSONSchema(
		t,
		StringKubernetesQualifiedName(),
		"expected_string_kubernetes_qualified_name.json",
		cases,
	)
}

func BenchmarkStringKubernetesQualifiedName(b *testing.B) {
	for _, tc := range stringK8sQualifiedNameTestCases {
		rule := StringKubernetesQualifiedName()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func benchmarkStringPaymentBankingRule(
	b *testing.B,
	rule govy.Rule[string],
	testCases []jsonschematest.Case[string],
) {
	b.Helper()
	for b.Loop() {
		for _, tc := range testCases {
			_ = rule.Validate(tc.Input)
		}
	}
}

func assertPaymentBankingRule(
	t *testing.T,
	rule govy.Rule[string],
	testCases []jsonschematest.Case[string],
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := rule.Validate(tc.Input)
			if tc.Valid {
				assert.NoError(t, err)
				return
			}
			assertPaymentBankingRuleError(t, err, expectedError, errorCode)
		})
	}
}

func assertPaymentBankingRuleError(
	t *testing.T,
	err error,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	assert.Require(t, assert.EqualError(t, err, expectedError))
	assert.Require(t, assert.IsType[*govy.RuleError](t, err))
	ruleError := err.(*govy.RuleError)
	assert.Equal(t, expectedError, ruleError.Description)
	assert.Equal(t, errorCode, ruleError.Code)
}

func assertBICPredicatesMatchRegexpOracle(
	t *testing.T,
	oracle *regexp.Regexp,
	input string,
) {
	t.Helper()
	expected := oracle.MatchString(input)
	if expected {
		expected = isValidBICCountryCode(input[4:6])
	}
	for name, predicate := range map[string]func(string) bool{
		"isValidBIC":            isValidBIC,
		"isValidBICISO93622014": isValidBICISO93622014,
	} {
		if actual := predicate(input); actual != expected {
			t.Errorf("%s(%q) = %t, regexp oracle = %t", name, input, actual, expected)
		}
	}
}

func readISOAlpha2CountryCodes(t *testing.T) []string {
	t.Helper()
	_, codes := readTestDataFields(t, "testdata/iso_3166_1_alpha2_2026-07-21.txt")
	return codes
}

func readTestDataFields(t *testing.T, path string) (raw string, fields []string) {
	t.Helper()
	data, err := os.ReadFile(path)
	assert.Require(t, assert.NoError(t, err))

	raw = string(data)
	for line := range strings.Lines(raw) {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields = append(fields, strings.Fields(line)...)
	}
	return raw, fields
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
	valid, mixed := stringFormatIDBenchmarkInputs(validInputs, invalidInputs)
	b.Run("valid", func(b *testing.B) {
		for b.Loop() {
			for _, input := range valid {
				_ = rule.Validate(input)
			}
		}
		b.ReportMetric(float64(len(valid)), "validations/op")
	})
	b.Run("mixed", func(b *testing.B) {
		for b.Loop() {
			for _, input := range mixed {
				_ = rule.Validate(input)
			}
		}
		b.ReportMetric(float64(len(mixed)), "validations/op")
	})
}

func benchmarkStringFormatIDPredicate(
	b *testing.B,
	predicate func(string) bool,
	validInputs map[string]string,
	invalidInputs map[string]string,
) {
	b.Helper()
	valid, mixed := stringFormatIDBenchmarkInputs(validInputs, invalidInputs)
	b.Run("valid", func(b *testing.B) {
		for b.Loop() {
			for _, input := range valid {
				_ = predicate(input)
			}
		}
		b.ReportMetric(float64(len(valid)), "validations/op")
	})
	b.Run("mixed", func(b *testing.B) {
		for b.Loop() {
			for _, input := range mixed {
				_ = predicate(input)
			}
		}
		b.ReportMetric(float64(len(mixed)), "validations/op")
	})
}

func stringFormatIDBenchmarkInputs(
	validInputs map[string]string,
	invalidInputs map[string]string,
) (valid, mixed []string) {
	valid = stringMapValues(validInputs)
	invalid := stringMapValues(invalidInputs)
	mixed = make([]string, 0, len(valid)+len(invalid))
	mixed = append(mixed, valid...)
	mixed = append(mixed, invalid...)
	return valid, mixed
}

func stringMapValues(values map[string]string) []string {
	keys := slices.Sorted(maps.Keys(values))
	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, values[key])
	}
	return result
}

func assertStringPredicateMatchesRegexp(
	t *testing.T,
	re *regexp.Regexp,
	predicate func(string) bool,
	input string,
) {
	t.Helper()
	expected := re.MatchString(input)
	if actual := predicate(input); actual != expected {
		t.Fatalf("predicate result for %q: expected %t, got %t", input, expected, actual)
	}
}

func loadNISTDigestOutputs(t *testing.T, fixture string, expectedCount, digestLength int) []string {
	t.Helper()

	path := filepath.Join("testdata", fixture)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if len(data) == 0 || data[len(data)-1] != '\n' {
		t.Fatalf("%s must be non-empty and end with a newline", path)
	}

	digests := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	if len(digests) != expectedCount {
		t.Fatalf("%s contains %d digests; want %d", path, len(digests), expectedCount)
	}

	seen := make(map[string]int, len(digests))
	for line, digest := range digests {
		if len(digest) != digestLength {
			t.Fatalf("%s:%d: digest length is %d; want %d", path, line+1, len(digest), digestLength)
		}
		if firstLine, ok := seen[digest]; ok {
			t.Fatalf("%s:%d: duplicate digest first seen on line %d", path, line+1, firstLine)
		}
		seen[digest] = line + 1
	}
	return digests
}

var isbnTestCases = []jsonschematest.Case[string]{
	{Name: "isbn 10 hyphenated", Input: "0-306-40615-2", Valid: true},
	{Name: "isbn 10 plain", Input: "0306406152", Valid: true},
	{Name: "isbn 10 x check", Input: "0-9752298-0-X", Valid: true},
	{Name: "isbn 10 spaced", Input: "0 9752298 0 x", Valid: true},
	{Name: "isbn 10 library converter numeric", Input: "0394170660", Valid: true},
	{Name: "isbn 10 library converter alternate", Input: "0717941728", Valid: true},
	{Name: "isbn 10 library converter x check", Input: "087779443X", Valid: true},
	{Name: "isbn 10 MARC hyphenated", Input: "0-87068-693-3", Valid: true},
	{Name: "isbn 13 hyphenated", Input: "978-0-306-40615-7", Valid: true},
	{Name: "isbn 13 plain", Input: "9780306406157", Valid: true},
	{Name: "isbn 13 grouped", Input: "978-3-16-148410-0", Valid: true},
	{Name: "isbn 13 agency manual hyphenated", Input: "978-92-95055-12-4", Valid: true},
	{Name: "isbn 13 agency manual spaced", Input: "978 92 95055 12 4", Valid: true},
	{Name: "isbn 13 agency manual compact", Input: "9789295055124", Valid: true},
	{Name: "isbn 13 agency manual hardback", Input: "978-951-45-9693-3", Valid: true},
	{Name: "isbn 13 agency manual paperback", Input: "978-951-45-9694-0", Valid: true},
	{Name: "isbn 13 agency manual PDF", Input: "978-951-45-9695-7", Valid: true},
	{Name: "isbn 13 agency manual EPUB", Input: "978-951-45-9696-4", Valid: true},
	{Name: "isbn 13 library converter first", Input: "9780060723804", Valid: true},
	{Name: "isbn 13 library converter second", Input: "9780060799748", Valid: true},
	{Name: "isbn 13 979 prefix", Input: "979-10-90636-07-1", Valid: true},
	{Name: "empty", Input: ""},
	{
		Name: "isbn 10 failed check", Input: "0-306-40615-3",
		JSONSchemaDifference: "JSON Schema checks ISBN syntax but does not validate the checksum.",
	},
	{
		Name: "isbn 10 x check mutation", Input: "0877794430",
		JSONSchemaDifference: "JSON Schema checks ISBN syntax but does not validate the checksum.",
	},
	{Name: "isbn 10 x in body", Input: "08777X443X"},
	{Name: "isbn 10 x in fourth position", Input: "087X79443X"},
	{Name: "isbn 10 short", Input: "087779443"},
	{Name: "isbn 10 trailing space", Input: "087779443X "},
	{
		Name: "isbn 13 failed check", Input: "978-0-306-40615-8",
		JSONSchemaDifference: "JSON Schema checks ISBN syntax but does not validate the checksum.",
	},
	{
		Name: "isbn 13 manual check mutation", Input: "978-92-95055-12-5",
		JSONSchemaDifference: "JSON Schema checks ISBN syntax but does not validate the checksum.",
	},
	{Name: "isbn 13 checksum valid wrong prefix", Input: "9779295055125"},
	{Name: "isbn 13 x in body", Input: "978-92-X5055-12-4"},
	{Name: "isbn 13 x check", Input: "978-92-95055-12-X"},
	{Name: "isbn 13 en dash separators", Input: "978–92–95055–12–4"},
	{Name: "isbn 13 full width digits", Input: "９７８９２９５０５５１２４"},
	{Name: "isbn 13 display prefix", Input: "ISBN 978-92-95055-12-4"},
	{Name: "repeated separator", Input: "978--0-306-40615-7"},
	{Name: "letters", Input: "abc"},
}

func TestStringISBN(t *testing.T) {
	rule := StringISBN()
	for _, tc := range isbnTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := rule.Validate(tc.Input)
			if tc.Valid {
				assert.NoError(t, err)
				return
			}
			assert.EqualError(
				t,
				err,
				"string must be a valid International Standard Book Number (ISBN) in ISBN-10 or ISBN-13 format",
			)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN))
		})
	}
}

var isbnVeryLargeInvalidTestCase = jsonschematest.Case[string]{
	Name: "very large invalid", Input: strings.Repeat("0", 1<<20),
}

func TestStringISBN_VeryLargeInvalid(t *testing.T) {
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
			err := test.rule.Validate(isbnVeryLargeInvalidTestCase.Input)
			assert.EqualError(t, err, test.message)
			assert.True(t, govy.HasErrorCode(err, test.code))
		})
	}
}

func TestStringISBN_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringISBN())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_isbn.json",
		slices.Concat(isbnTestCases, []jsonschematest.Case[string]{isbnVeryLargeInvalidTestCase}),
	)
}

func BenchmarkStringISBN(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISBN(),
		isbnTestCases,
	)
}

func BenchmarkStringISBNVeryLargeInvalid(b *testing.B) {
	rule := StringISBN()
	for b.Loop() {
		_ = rule.Validate(isbnVeryLargeInvalidTestCase.Input)
	}
	b.ReportMetric(1, "validations/op")
}

var isbn10TestCases = []jsonschematest.Case[string]{
	{Name: "hyphenated", Input: "0-306-40615-2", Valid: true},
	{Name: "plain", Input: "0306406152", Valid: true},
	{Name: "x check", Input: "0-9752298-0-X", Valid: true},
	{Name: "spaced", Input: "0 9752298 0 x", Valid: true},
	{Name: "library converter numeric", Input: "0394170660", Valid: true},
	{Name: "library converter alternate", Input: "0717941728", Valid: true},
	{Name: "library converter x check", Input: "087779443X", Valid: true},
	{Name: "MARC hyphenated", Input: "0-87068-693-3", Valid: true},
	{Name: "empty", Input: ""},
	{
		Name:                 "failed check",
		Input:                "0-306-40615-3",
		JSONSchemaDifference: "JSON Schema checks ISBN-10 syntax but does not validate the checksum.",
	},
	{
		Name:                 "x check mutation",
		Input:                "0877794430",
		JSONSchemaDifference: "JSON Schema checks ISBN-10 syntax but does not validate the checksum.",
	},
	{Name: "x in body", Input: "08777X443X"},
	{Name: "x in fourth position", Input: "087X79443X"},
	{Name: "short", Input: "087779443"},
	{Name: "trailing space", Input: "087779443X "},
	{Name: "isbn 13", Input: "978-0-306-40615-7"},
	{Name: "isbn 13 plain", Input: "9780306406157"},
	{Name: "repeated separator", Input: "0-306--40615-2"},
}

func TestStringISBN10(t *testing.T) {
	rule := StringISBN10()
	for _, tc := range isbn10TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := rule.Validate(tc.Input)
			if tc.Valid {
				assert.NoError(t, err)
				return
			}
			assert.EqualError(
				t,
				err,
				"string must be a valid International Standard Book Number (ISBN) in ISBN-10 format",
			)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN10))
		})
	}
}

func TestStringISBN10_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringISBN10())))
	assert.Require(t, assert.NoError(t, err))
	cases := append([]jsonschematest.Case[string]{isbnVeryLargeInvalidTestCase}, isbn10TestCases...)
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_isbn10.json", cases)
}

func BenchmarkStringISBN10(b *testing.B) {
	rule := StringISBN10()
	for b.Loop() {
		for _, tc := range isbn10TestCases {
			_ = rule.Validate(tc.Input)
		}
	}
	b.ReportMetric(float64(len(isbn10TestCases)), "validations/op")
}

var isbn13TestCases = []jsonschematest.Case[string]{
	{Name: "hyphenated", Input: "978-0-306-40615-7", Valid: true},
	{Name: "plain", Input: "9780306406157", Valid: true},
	{Name: "grouped", Input: "978-3-16-148410-0", Valid: true},
	{Name: "agency manual hyphenated", Input: "978-92-95055-12-4", Valid: true},
	{Name: "agency manual spaced", Input: "978 92 95055 12 4", Valid: true},
	{Name: "agency manual compact", Input: "9789295055124", Valid: true},
	{Name: "agency manual hardback", Input: "978-951-45-9693-3", Valid: true},
	{Name: "agency manual paperback", Input: "978-951-45-9694-0", Valid: true},
	{Name: "agency manual PDF", Input: "978-951-45-9695-7", Valid: true},
	{Name: "agency manual EPUB", Input: "978-951-45-9696-4", Valid: true},
	{Name: "library converter first", Input: "9780060723804", Valid: true},
	{Name: "library converter second", Input: "9780060799748", Valid: true},
	{Name: "979 prefix", Input: "979-10-90636-07-1", Valid: true},
	{Name: "empty", Input: ""},
	{Name: "isbn 10", Input: "0-306-40615-2"},
	{
		Name: "failed check", Input: "978-0-306-40615-8",
		JSONSchemaDifference: "JSON Schema checks ISBN-13 syntax but does not validate the checksum.",
	},
	{
		Name: "manual check mutation", Input: "978-92-95055-12-5",
		JSONSchemaDifference: "JSON Schema checks ISBN-13 syntax but does not validate the checksum.",
	},
	{Name: "checksum valid wrong prefix", Input: "9779295055125"},
	{Name: "x in body", Input: "978-92-X5055-12-4"},
	{Name: "x check", Input: "978-92-95055-12-X"},
	{Name: "en dash separators", Input: "978–92–95055–12–4"},
	{Name: "full width digits", Input: "９７８９２９５０５５１２４"},
	{Name: "display prefix", Input: "ISBN 978-92-95055-12-4"},
	{Name: "invalid prefix", Input: "9770306406157"},
	{Name: "trailing space", Input: "978 0 306 40615 7 "},
}

func TestStringISBN13(t *testing.T) {
	rule := StringISBN13()
	for _, tc := range isbn13TestCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := rule.Validate(tc.Input)
			if tc.Valid {
				assert.NoError(t, err)
				return
			}
			assert.EqualError(
				t,
				err,
				"string must be a valid International Standard Book Number (ISBN) in ISBN-13 format",
			)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN13))
		})
	}
}

func TestISBNPredicatesMatchReference(t *testing.T) {
	tests := map[string]struct {
		predicate func(string) bool
		reference func(string) bool
		inputs    []jsonschematest.Case[string]
	}{
		"isbn": {
			predicate: isISBN,
			reference: referenceISBN,
			inputs:    isbnTestCases,
		},
		"isbn-10": {
			predicate: isISBN10,
			reference: referenceISBN10,
			inputs:    isbn10TestCases,
		},
		"isbn-13": {
			predicate: isISBN13,
			reference: referenceISBN13,
			inputs:    isbn13TestCases,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testStringPredicateMatchesReference(t, test.predicate, test.reference, test.inputs)
		})
	}
}

func TestStringISBN13_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringISBN13())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(
		t,
		schema,
		"testdata/jsonschema/expected_string_isbn13.json",
		slices.Concat(isbn13TestCases, []jsonschematest.Case[string]{isbnVeryLargeInvalidTestCase}),
	)
}

func BenchmarkStringISBN13(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISBN13(),
		isbn13TestCases,
	)
}

// issnTestCases includes exact compact construction examples from the
// [ISSN Manual, May 2025] because StringISSN requires the ASCII-hyphenated form.
// Their derived hyphenated forms are also included as valid inputs.
//
// [ISSN Manual, May 2025]: https://www.issn.org/wp-content/uploads/2025/05/Manual-ISSN_ENG-marc21_May2025.pdf
var issnTestCases = []jsonschematest.Case[string]{
	{Name: "numeric check", Input: "2049-3630", Valid: true},
	{Name: "numeric example", Input: "0378-5955", Valid: true},
	{Name: "uppercase x", Input: "2434-561X", Valid: true},
	{Name: "lowercase x", Input: "2434-561x", Valid: true},
	{Name: "manual numeric", Input: "1106-1111", Valid: true},
	{Name: "manual uppercase x", Input: "1092-003X", Valid: true},
	{Name: "library check digit", Input: "0317-8471", Valid: true},
	{Name: "numeric 2162", Input: "2162-3546", Valid: true},
	{Name: "numeric 1548", Input: "1548-7180", Valid: true},
	{Name: "uppercase x 1204", Input: "1204-539X", Valid: true},
	{Name: "empty", Input: ""},
	{Name: "missing hyphen", Input: "20493630"},
	{Name: "manual compact 2162", Input: "21623546"},
	{Name: "manual compact 1548", Input: "15487180"},
	{
		Name: "failed check", Input: "2049-3631",
		JSONSchemaDifference: "JSON Schema checks ISSN syntax but does not validate the checksum.",
	},
	{
		Name: "numeric check mutation", Input: "1106-1112",
		JSONSchemaDifference: "JSON Schema checks ISSN syntax but does not validate the checksum.",
	},
	{
		Name: "uppercase x check mutation", Input: "1092-0030",
		JSONSchemaDifference: "JSON Schema checks ISSN syntax but does not validate the checksum.",
	},
	{Name: "wrong grouping", Input: "204-93630"},
	{Name: "x before check", Input: "2049-36X0"},
	{Name: "hyphen as check", Input: "2049-363-"},
	{Name: "unicode hyphen", Input: "1106–1111"},
	{Name: "U+2010 hyphen", Input: "1092‐003X"},
	{Name: "space separator", Input: "1106 1111"},
	{Name: "display prefix", Input: "ISSN 1106-1111"},
	{Name: "trailing newline", Input: "1106-1111\n"},
	{Name: "full width digits", Input: "１１０６-１１１１"},
}

func TestStringISSN(t *testing.T) {
	rule := StringISSN()
	for _, tc := range issnTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			err := rule.Validate(tc.Input)
			if tc.Valid {
				assert.NoError(t, err)
				return
			}
			assert.EqualError(t, err, "string must be a valid International Standard Serial Number (ISSN)")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISSN))
		})
	}
}

func TestISSNPredicateMatchesReference(t *testing.T) {
	format := regexp.MustCompile(`^\d{4}-\d{3}[0-9Xx]$`)
	testStringPredicateMatchesReference(
		t,
		isISSN,
		func(s string) bool {
			return referenceISSN(format, s)
		},
		issnTestCases,
	)
}

func TestStringISSN_JSONSchema(t *testing.T) {
	t.Parallel()

	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(StringISSN())))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_string_issn.json", issnTestCases)
}

func BenchmarkStringISSN(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISSN(),
		issnTestCases,
	)
}

func Benchmark_isISSN(b *testing.B) {
	for b.Loop() {
		for _, tc := range issnTestCases {
			if isISSN(tc.Input) != tc.Valid {
				b.Fatalf("%s: expected ISSN validity %t", tc.Name, tc.Valid)
			}
		}
	}
	b.ReportMetric(float64(len(issnTestCases)), "validations/op")
}

func benchmarkStringPublicationRule(
	b *testing.B,
	rule govy.Rule[string],
	cases []jsonschematest.Case[string],
) {
	b.Helper()
	for b.Loop() {
		for _, tc := range cases {
			_ = rule.Validate(tc.Input)
		}
	}
	b.ReportMetric(float64(len(cases)), "validations/op")
}

type stringPredicate func(string) bool

func testStringPredicateMatchesReference(
	t *testing.T,
	predicate stringPredicate,
	reference stringPredicate,
	cases []jsonschematest.Case[string],
) {
	t.Helper()
	for _, tc := range cases {
		assertStringPredicateMatchesReference(t, predicate, reference, tc.Name, tc.Input)
		testStringPredicateByteEdits(t, predicate, reference, tc.Input)
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

func stringBooleanJSONSchemaCases(inputs []*struct {
	in         string
	shouldFail bool
},
) []jsonschematest.Case[string] {
	cases := make([]jsonschematest.Case[string], len(inputs))
	for i, tc := range inputs {
		cases[i] = jsonschematest.Case[string]{Name: strconv.Quote(tc.in), Input: tc.in, Valid: !tc.shouldFail}
	}
	return cases
}

func stringRegexpJSONSchemaCases(inputs []*struct {
	in            string
	expectedError string
},
) []jsonschematest.Case[string] {
	cases := make([]jsonschematest.Case[string], len(inputs))
	for i, tc := range inputs {
		cases[i] = jsonschematest.Case[string]{Name: strconv.Quote(tc.in), Input: tc.in, Valid: tc.expectedError == ""}
	}
	return cases
}

func stringNamedJSONSchemaCases(validInputs, invalidInputs map[string]string) []jsonschematest.Case[string] {
	cases := make([]jsonschematest.Case[string], 0, len(validInputs)+len(invalidInputs))
	for name, input := range validInputs {
		cases = append(cases, jsonschematest.Case[string]{Name: "valid/" + name, Input: input, Valid: true})
	}
	for name, input := range invalidInputs {
		cases = append(cases, jsonschematest.Case[string]{Name: "invalid/" + name, Input: input})
	}
	return cases
}

func taxJSONSchemaCases(groups ...[]stringTaxIDTestCase) []jsonschematest.Case[string] {
	var cases []jsonschematest.Case[string]
	for _, group := range groups {
		for _, tc := range group {
			cases = append(cases, jsonschematest.Case[string]{
				Name:  tc.name,
				Input: tc.in,
				Valid: tc.shouldPass,
			})
		}
	}
	return cases
}

func assertStringTextJSONSchema(
	t *testing.T,
	rule govy.RulesInterface[string],
	fixture string,
	cases []jsonschematest.Case[string],
) {
	t.Helper()
	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(rule)))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, filepath.Join("testdata", "jsonschema", fixture), cases)
}

type stringTextArgumentCase struct {
	arguments []string
	input     string
	valid     bool
}

func assertStringArgumentsJSONSchema(
	t *testing.T,
	fixturePrefix string,
	constructor func(...string) govy.Rule[string],
	cases []stringTextArgumentCase,
) {
	t.Helper()
	type group struct {
		arguments []string
		cases     []jsonschematest.Case[string]
	}
	groups := make(map[string]*group)
	for _, tc := range cases {
		name := strings.Join(tc.arguments, "_")
		if groups[name] == nil {
			groups[name] = &group{arguments: tc.arguments}
		}
		groups[name].cases = append(groups[name].cases, jsonschematest.Case[string]{
			Name: strconv.Quote(tc.input), Input: tc.input, Valid: tc.valid,
		})
	}
	for name, group := range groups {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			fixture := "expected_" + fixturePrefix + "_" + name + ".json"
			assertStringTextJSONSchema(t, constructor(group.arguments...), fixture, group.cases)
		})
	}
}

func stringCharacterJSONSchemaCases(cases []*struct {
	in         string
	shouldFail bool
},
) []jsonschematest.Case[string] {
	result := make([]jsonschematest.Case[string], 0, len(cases))
	for _, tc := range cases {
		result = append(result, jsonschematest.Case[string]{
			Name: strconv.Quote(tc.in), Input: tc.in, Valid: !tc.shouldFail,
		})
	}
	return result
}
