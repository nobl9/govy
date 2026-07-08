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

// cspell:disable
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

var stringCreditCardTestCases = []*stringFormatIDTestCase{
	{in: "4111111111111111"},
	{in: "4242424242424242"},
	{in: "5555555555554444"},
	{in: "378282246310005"},
	{in: "6011111111111117"},
	{in: "4222222222222"},
	{in: "4000000000000000006"},
	{
		in:            "4111111111111112",
		expectedError: "string must be a valid credit card number",
	},
	{
		in:            "4111 1111 1111 1111",
		expectedError: "string must be a valid credit card number",
	},
	{
		in:            "4111-1111-1111-1111",
		expectedError: "string must be a valid credit card number",
	},
	{
		in:            "411111111111",
		expectedError: "string must be a valid credit card number",
	},
	{
		in:            "40000000000000000006",
		expectedError: "string must be a valid credit card number",
	},
	{
		in:            "0000000000000",
		expectedError: "string must be a valid credit card number",
	},
}

var stringLuhnChecksumTestCases = []*stringFormatIDTestCase{
	{in: "0"},
	{in: "0000"},
	{in: "79927398713"},
	{in: "4111111111111111"},
	{
		in:            "",
		expectedError: "string must pass the Luhn checksum",
	},
	{
		in:            "79927398710",
		expectedError: "string must pass the Luhn checksum",
	},
	{
		in:            "7992739871A",
		expectedError: "string must pass the Luhn checksum",
	},
}

var stringBICTestCases = []*stringFormatIDTestCase{
	{in: "DEUTDEFF"},
	{in: "DEUTDEFF500"},
	{in: "NEDSZAJJXXX"},
	{in: "A1B2US33XXX"},
	{
		in:            "deutdeff",
		expectedError: "string must be a valid BIC",
	},
	{
		in:            "DEUTD3FF",
		expectedError: "string must be a valid BIC",
	},
	{
		in:            "DEUTDEF",
		expectedError: "string must be a valid BIC",
	},
	{
		in:            "DEUTDEFF50",
		expectedError: "string must be a valid BIC",
	},
	{
		in:            "DEUTDEFF5000",
		expectedError: "string must be a valid BIC",
	},
	{
		in:            "DEUTDEFF50!",
		expectedError: "string must be a valid BIC",
	},
}

var stringBICISO93622014TestCases = []*stringFormatIDTestCase{
	{in: "DEUTDEFF"},
	{in: "DEUTDEFF500"},
	{in: "NEDSZAJJXXX"},
	{
		in:            "A1B2US33XXX",
		expectedError: "string must be a valid ISO 9362:2014 BIC",
	},
	{
		in:            "deutdeff",
		expectedError: "string must be a valid ISO 9362:2014 BIC",
	},
	{
		in:            "DEUTD3FF",
		expectedError: "string must be a valid ISO 9362:2014 BIC",
	},
	{
		in:            "DEUTDEF",
		expectedError: "string must be a valid ISO 9362:2014 BIC",
	},
	{
		in:            "DEUTDEFF50!",
		expectedError: "string must be a valid ISO 9362:2014 BIC",
	},
}

var stringEINTestCases = []*stringFormatIDTestCase{
	{in: "12-3456789"},
	{in: "99-3456789"},
	{
		in:            "00-0000000",
		expectedError: "string must be a valid EIN",
	},
	{
		in:            "07-3456789",
		expectedError: "string must be a valid EIN",
	},
	{
		in:            "123456789",
		expectedError: "string must be a valid EIN",
	},
	{
		in:            "1-23456789",
		expectedError: "string must be a valid EIN",
	},
	{
		in:            "12-345678",
		expectedError: "string must be a valid EIN",
	},
	{
		in:            "12-34567890",
		expectedError: "string must be a valid EIN",
	},
	{
		in:            "AB-3456789",
		expectedError: "string must be a valid EIN",
	},
	{
		in:            "12_3456789",
		expectedError: "string must be a valid EIN",
	},
}

var stringSSNTestCases = []*stringFormatIDTestCase{
	{in: "123-45-6789"},
	{in: "899-99-9999"},
	{in: "001-01-0001"},
	{
		in:            "000-45-6789",
		expectedError: "string must be a valid SSN",
	},
	{
		in:            "666-45-6789",
		expectedError: "string must be a valid SSN",
	},
	{
		in:            "900-45-6789",
		expectedError: "string must be a valid SSN",
	},
	{
		in:            "999-45-6789",
		expectedError: "string must be a valid SSN",
	},
	{
		in:            "123-00-6789",
		expectedError: "string must be a valid SSN",
	},
	{
		in:            "123-45-0000",
		expectedError: "string must be a valid SSN",
	},
	{
		in:            "123456789",
		expectedError: "string must be a valid SSN",
	},
	{
		in:            "12A-45-6789",
		expectedError: "string must be a valid SSN",
	},
}

// cspell:enable

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

func TestStringCreditCard(t *testing.T) {
	testStringFormatIDRule(t, StringCreditCard(), ErrorCodeStringCreditCard, stringCreditCardTestCases)
}

func BenchmarkStringCreditCard(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringCreditCard(), stringCreditCardTestCases)
}

func TestStringLuhnChecksum(t *testing.T) {
	testStringFormatIDRule(t, StringLuhnChecksum(), ErrorCodeStringLuhnChecksum, stringLuhnChecksumTestCases)
}

func BenchmarkStringLuhnChecksum(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringLuhnChecksum(), stringLuhnChecksumTestCases)
}

func TestStringBIC(t *testing.T) {
	testStringFormatIDRule(t, StringBIC(), ErrorCodeStringBIC, stringBICTestCases)
}

func BenchmarkStringBIC(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringBIC(), stringBICTestCases)
}

func TestStringBICISO93622014(t *testing.T) {
	testStringFormatIDRule(t, StringBICISO93622014(), ErrorCodeStringBICISO93622014, stringBICISO93622014TestCases)
}

func BenchmarkStringBICISO93622014(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringBICISO93622014(), stringBICISO93622014TestCases)
}

func TestStringEIN(t *testing.T) {
	testStringFormatIDRule(t, StringEIN(), ErrorCodeStringEIN, stringEINTestCases)
}

func BenchmarkStringEIN(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringEIN(), stringEINTestCases)
}

func TestStringSSN(t *testing.T) {
	testStringFormatIDRule(t, StringSSN(), ErrorCodeStringSSN, stringSSNTestCases)
}

func BenchmarkStringSSN(b *testing.B) {
	benchmarkStringFormatIDRule(b, StringSSN(), stringSSNTestCases)
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
	for _, tc := range testCases {
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}
