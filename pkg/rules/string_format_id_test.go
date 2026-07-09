package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

type stringFormatIDTestCase struct {
	in            string
	expectedError string
}

type stringUUIDTestCase struct {
	in         string
	shouldFail bool
}

// cspell:disable
var stringUUIDTestCases = []*stringUUIDTestCase{
	{in: "00000000-0000-0000-0000-000000000000"},
	{in: "e190c630-8873-11ee-b9d1-0242ac120002"},
	{in: "79258D24-01A7-47E5-ACBB-7E762DE52298"},
	{in: "a987Fbc9-4bed-3078-cf07-9141ba07c9f3"},
	{in: "foobar", shouldFail: true},
	{in: "0987654321", shouldFail: true},
	{in: "AXAXAXAX-AAAA-AAAA-AAAA-AAAAAAAAAAAA", shouldFail: true},
	{in: "00000000-0000-0000-0000-0000000000", shouldFail: true},
	{in: "", shouldFail: true},
	{in: "xxxa987Fbc9-4bed-3078-cf07-9141ba07c9f3", shouldFail: true},
	{in: "a987Fbc9-4bed-3078-cf07-9141ba07c9f3xxx", shouldFail: true},
	{in: "a987Fbc94bed3078cf079141ba07c9f3", shouldFail: true},
	{in: "934859", shouldFail: true},
	{in: "987fbc9-4bed-3078-cf07a-9141ba07c9F3", shouldFail: true},
	{in: "aaaaaaaa-1111-1111-aaaG-111111111111", shouldFail: true},
}

var stringUUIDRFC4122TestCases = []*stringFormatIDTestCase{
	{in: "f81d4fae-7dec-11d0-a765-00a0c91e6bf6"},
	{in: "A987FBC9-4BED-3078-8F07-9141BA07C9F3"},
	{in: "57b73598-8764-4ad0-a76a-679bb6640eb1"},
	{in: "987fbc97-4bed-5078-9f07-9141ba07c9f3"},
	{
		in:            "00000000-0000-0000-0000-000000000000",
		expectedError: "string must be a valid RFC 4122 UUID",
	},
	{
		in:            "a987fbc9-4bed-3078-cf07-9141ba07c9f3",
		expectedError: "string must be a valid RFC 4122 UUID",
	},
	{
		in:            "01890f3e-23b4-7d68-9c2a-8f56c8a87f1b",
		expectedError: "string must be a valid RFC 4122 UUID",
	},
	{
		in:            "a987fbc94bed30788f079141ba07c9f3",
		expectedError: "string must be a valid RFC 4122 UUID",
	},
	{
		in:            "aaaaaaaa-1111-1111-aaaG-111111111111",
		expectedError: "string must be a valid RFC 4122 UUID",
	},
	{
		in:            "",
		expectedError: "string must be a valid RFC 4122 UUID",
	},
}

var stringUUIDv3TestCases = []*stringFormatIDTestCase{
	{in: "a987fbc9-4bed-3078-8f07-9141ba07c9f3"},
	{in: "A987FBC9-4BED-3078-BF07-9141BA07C9F3"},
	{
		in:            "57b73598-8764-4ad0-a76a-679bb6640eb1",
		expectedError: "string must be a valid version 3 RFC 4122 UUID",
	},
	{
		in:            "a987fbc9-4bed-3078-cf07-9141ba07c9f3",
		expectedError: "string must be a valid version 3 RFC 4122 UUID",
	},
	{
		in:            "a987fbc9-4bed-3078-8f07-9141ba07c9fg",
		expectedError: "string must be a valid version 3 RFC 4122 UUID",
	},
}

var stringUUIDv4TestCases = []*stringFormatIDTestCase{
	{in: "57b73598-8764-4ad0-a76a-679bb6640eb1"},
	{in: "57B73598-8764-4AD0-B76A-679BB6640EB1"},
	{
		in:            "a987fbc9-4bed-3078-8f07-9141ba07c9f3",
		expectedError: "string must be a valid version 4 RFC 4122 UUID",
	},
	{
		in:            "57b73598-8764-4ad0-c76a-679bb6640eb1",
		expectedError: "string must be a valid version 4 RFC 4122 UUID",
	},
	{
		in:            "57b73598-8764-4ad0-a76a-679bb6640eg1",
		expectedError: "string must be a valid version 4 RFC 4122 UUID",
	},
}

var stringUUIDv5TestCases = []*stringFormatIDTestCase{
	{in: "987fbc97-4bed-5078-9f07-9141ba07c9f3"},
	{in: "987FBC97-4BED-5078-AF07-9141BA07C9F3"},
	{
		in:            "57b73598-8764-4ad0-a76a-679bb6640eb1",
		expectedError: "string must be a valid version 5 RFC 4122 UUID",
	},
	{
		in:            "987fbc97-4bed-5078-cf07-9141ba07c9f3",
		expectedError: "string must be a valid version 5 RFC 4122 UUID",
	},
	{
		in:            "987fbc97-4bed-5078-9f07-9141ba07c9fg",
		expectedError: "string must be a valid version 5 RFC 4122 UUID",
	},
}

var stringULIDTestCases = []*stringFormatIDTestCase{
	{in: "01ARZ3NDEKTSV4RRFFQ69G5FAV"},
	{in: "01arz3ndektsv4rrffq69g5fav"},
	{in: "7ZZZZZZZZZZZZZZZZZZZZZZZZZ"},
	{
		in:            "01ARZ3NDEKTSV4RRFFQ69G5FA",
		expectedError: "string must be a valid ULID",
	},
	{
		in:            "01ARZ3NDEKTSV4RRFFQ69G5FAV0",
		expectedError: "string must be a valid ULID",
	},
	{
		in:            "8ZZZZZZZZZZZZZZZZZZZZZZZZZ",
		expectedError: "string must be a valid ULID",
	},
	{
		in:            "01ARZ3NDEKTSV4RRFFQ69G5FIV",
		expectedError: "string must be a valid ULID",
	},
	{
		in:            "01ARZ3NDEKTSV4RRFFQ69G5FLV",
		expectedError: "string must be a valid ULID",
	},
	{
		in:            "01ARZ3NDEKTSV4RRFFQ69G5FOV",
		expectedError: "string must be a valid ULID",
	},
	{
		in:            "01ARZ3NDEKTSV4RRFFQ69G5FUV",
		expectedError: "string must be a valid ULID",
	},
	{
		in:            "01ARZ3NDEKTSV4RRFFQ69G5F-V",
		expectedError: "string must be a valid ULID",
	},
}

// cspell:enable

func TestStringUUID(t *testing.T) {
	rule := StringUUID()
	for _, tc := range stringUUIDTestCases {
		err := rule.Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringUUID))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringUUID(b *testing.B) {
	rule := StringUUID()
	for b.Loop() {
		for _, tc := range stringUUIDTestCases {
			_ = rule.Validate(tc.in)
		}
	}
}

func TestStringUUIDRFC4122(t *testing.T) {
	testStringFormatIDRule(t, StringUUIDRFC4122(), ErrorCodeStringUUIDRFC4122, stringUUIDRFC4122TestCases)
}

func BenchmarkStringUUIDRFC4122(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDRFC4122(), stringUUIDRFC4122TestCases)
}

func TestStringUUIDv3(t *testing.T) {
	testStringFormatIDRule(t, StringUUIDv3(), ErrorCodeStringUUIDv3, stringUUIDv3TestCases)
}

func BenchmarkStringUUIDv3(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv3(), stringUUIDv3TestCases)
}

func TestStringUUIDv4(t *testing.T) {
	testStringFormatIDRule(t, StringUUIDv4(), ErrorCodeStringUUIDv4, stringUUIDv4TestCases)
}

func BenchmarkStringUUIDv4(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv4(), stringUUIDv4TestCases)
}

func TestStringUUIDv5(t *testing.T) {
	testStringFormatIDRule(t, StringUUIDv5(), ErrorCodeStringUUIDv5, stringUUIDv5TestCases)
}

func BenchmarkStringUUIDv5(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringUUIDv5(), stringUUIDv5TestCases)
}

func TestStringULID(t *testing.T) {
	testStringFormatIDRule(t, StringULID(), ErrorCodeStringULID, stringULIDTestCases)
}

func BenchmarkStringULID(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringULID(), stringULIDTestCases)
}

func testStringFormatIDRule(
	t *testing.T,
	rule govy.Rule[string],
	errorCode govy.ErrorCode,
	testCases []*stringFormatIDTestCase,
) {
	t.Helper()
	for _, tc := range testCases {
		err := rule.Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, errorCode))
		} else {
			assert.NoError(t, err)
		}
	}
}

func benchmarkStringFormatIDRule(
	b *testing.B,
	rule govy.Rule[string],
	testCases []*stringFormatIDTestCase,
) {
	b.Helper()
	for b.Loop() {
		for _, tc := range testCases {
			_ = rule.Validate(tc.in)
		}
	}
}
