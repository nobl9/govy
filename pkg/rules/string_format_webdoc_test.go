package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

var stringJWTTestCases = []*struct {
	in            string
	expectedError string
}{
	// cspell:disable
	{
		in: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9." +
			"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ." +
			"SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	},
	{
		in: "eyJhbGciOiJub25lIn0.e30.",
	},
	{
		in:            "eyJhbGciOiJub25lIn0.e30.c2ln",
		expectedError: `string must be a valid JSON Web Token: JWT signature segment must be empty when alg is "none"`,
	},
	{
		in:            "not-a-jwt",
		expectedError: "string must be a valid JSON Web Token: expected 3 JWT segments",
	},
	{
		in:            "eyJhbGciOiJIUzI1NiJ9.e30.",
		expectedError: `string must be a valid JSON Web Token: JWT signature segment must not be empty unless alg is "none"`,
	},
	{
		in:            "e30.e30.c2ln",
		expectedError: `string must be a valid JSON Web Token: JWT header must contain an "alg" string`,
	},
	{
		in:            "bnVsbA.e30.c2ln",
		expectedError: "string must be a valid JSON Web Token: JWT header segment must contain a JSON object",
	},
	{
		in:            "eyJhbGciOiJIUzI1NiJ9=.e30.c2ln",
		expectedError: "string must be a valid JSON Web Token: JWT header segment must be Base64URL encoded without padding",
	},
	// cspell:enable
}

func TestStringJWT(t *testing.T) {
	for _, tc := range stringJWTTestCases {
		err := StringJWT().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringJWT))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringJWT(b *testing.B) {
	for _, tc := range stringJWTTestCases {
		rule := StringJWT()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringSemverTestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "0.1.0"},
	{in: "1.0.0"},
	{in: "1.0.0-alpha"},
	{in: "1.0.0-alpha.1"},
	{in: "1.0.0+20130313144700"},
	{in: "1.0.0-beta+exp.sha.5114f85"},
	{in: "2.7.3-rc.1+build.11.e0f985a"},
	{in: "", expectedError: "string must be a valid semantic version"},
	{in: "1", expectedError: "string must be a valid semantic version"},
	{in: "1.2", expectedError: "string must be a valid semantic version"},
	{in: "1.2.3.4", expectedError: "string must be a valid semantic version"},
	{in: "01.2.3", expectedError: "string must be a valid semantic version"},
	{in: "1.02.3", expectedError: "string must be a valid semantic version"},
	{in: "1.2.03", expectedError: "string must be a valid semantic version"},
	{in: "1.2.3-", expectedError: "string must be a valid semantic version"},
	{in: "1.2.3-01", expectedError: "string must be a valid semantic version"},
	{in: "v1.2.3", expectedError: "string must be a valid semantic version"},
	{in: "1.2.3+build..1", expectedError: "string must be a valid semantic version"},
}

func TestStringSemver(t *testing.T) {
	for _, tc := range stringSemverTestCases {
		err := StringSemver().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSemver))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringSemver(b *testing.B) {
	for _, tc := range stringSemverTestCases {
		rule := StringSemver()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringCVETestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "CVE-1999-0001"},
	{in: "CVE-2014-0160"},
	{in: "CVE-2021-0100"},
	{in: "CVE-2021-0990"},
	{in: "CVE-2021-44228"},
	{in: "CVE-2024-12345"},
	{in: "", expectedError: "string must be a valid CVE ID"},
	{in: "cve-2021-44228", expectedError: "string must be a valid CVE ID"},
	{in: "CVE-2021-0000", expectedError: "string must be a valid CVE ID"},
	{in: "CVE-2021-00001", expectedError: "string must be a valid CVE ID"},
	{in: "CVE-2021-123", expectedError: "string must be a valid CVE ID"},
	{in: "CVE-1998-0001", expectedError: "string must be a valid CVE ID"},
	{in: "CVE-2021-ABCD", expectedError: "string must be a valid CVE ID"},
	{in: "CVE-10000-0001", expectedError: "string must be a valid CVE ID"},
}

func TestStringCVE(t *testing.T) {
	for _, tc := range stringCVETestCases {
		err := StringCVE().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCVE))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringCVE(b *testing.B) {
	for _, tc := range stringCVETestCases {
		rule := StringCVE()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringE164TestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "+14155552671"},
	{in: "+442071838750"},
	{in: "+33123456789"},
	{in: "+4915123456789"},
	{in: "", expectedError: "string must be a valid E.164 phone number"},
	{in: "14155552671", expectedError: "string must be a valid E.164 phone number"},
	{in: "+014155552671", expectedError: "string must be a valid E.164 phone number"},
	{in: "+1", expectedError: "string must be a valid E.164 phone number"},
	{in: "+1234567890123456", expectedError: "string must be a valid E.164 phone number"},
	{in: "+1 4155552671", expectedError: "string must be a valid E.164 phone number"},
	{in: "+1-415-555-2671", expectedError: "string must be a valid E.164 phone number"},
	{in: "+ABC", expectedError: "string must be a valid E.164 phone number"},
}

func TestStringE164(t *testing.T) {
	for _, tc := range stringE164TestCases {
		err := StringE164().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringE164))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringE164(b *testing.B) {
	for _, tc := range stringE164TestCases {
		rule := StringE164()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringISBNTestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "0-306-40615-2"},
	{in: "0306406152"},
	{in: "0-9752298-0-X"},
	{in: "0 9752298 0 x"},
	{in: "978-0-306-40615-7"},
	{in: "9780306406157"},
	{in: "978-3-16-148410-0"},
	{in: "", expectedError: "string must be a valid ISBN"},
	{in: "0-306-40615-3", expectedError: "string must be a valid ISBN"},
	{in: "978-0-306-40615-8", expectedError: "string must be a valid ISBN"},
	{in: "978030640615X", expectedError: "string must be a valid ISBN"},
	{in: "978--0-306-40615-7", expectedError: "string must be a valid ISBN"},
	{in: "abc", expectedError: "string must be a valid ISBN"},
}

func TestStringISBN(t *testing.T) {
	for _, tc := range stringISBNTestCases {
		err := StringISBN().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISBN(b *testing.B) {
	for _, tc := range stringISBNTestCases {
		rule := StringISBN()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringISBN10TestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "0-306-40615-2"},
	{in: "0306406152"},
	{in: "0-9752298-0-X"},
	{in: "0 9752298 0 x"},
	{in: "", expectedError: "string must be a valid ISBN-10"},
	{in: "0-306-40615-3", expectedError: "string must be a valid ISBN-10"},
	{in: "978-0-306-40615-7", expectedError: "string must be a valid ISBN-10"},
	{in: "9780306406157", expectedError: "string must be a valid ISBN-10"},
	{in: "0-306--40615-2", expectedError: "string must be a valid ISBN-10"},
}

func TestStringISBN10(t *testing.T) {
	for _, tc := range stringISBN10TestCases {
		err := StringISBN10().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN10))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISBN10(b *testing.B) {
	for _, tc := range stringISBN10TestCases {
		rule := StringISBN10()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringISBN13TestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "978-0-306-40615-7"},
	{in: "9780306406157"},
	{in: "978-3-16-148410-0"},
	{in: "979-10-90636-07-1"},
	{in: "", expectedError: "string must be a valid ISBN-13"},
	{in: "0-306-40615-2", expectedError: "string must be a valid ISBN-13"},
	{in: "978-0-306-40615-8", expectedError: "string must be a valid ISBN-13"},
	{in: "978030640615X", expectedError: "string must be a valid ISBN-13"},
	{in: "9770306406157", expectedError: "string must be a valid ISBN-13"},
	{in: "978 0 306 40615 7 ", expectedError: "string must be a valid ISBN-13"},
}

func TestStringISBN13(t *testing.T) {
	for _, tc := range stringISBN13TestCases {
		err := StringISBN13().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN13))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISBN13(b *testing.B) {
	for _, tc := range stringISBN13TestCases {
		rule := StringISBN13()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringISSNTestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "2049-3630"},
	{in: "0378-5955"},
	{in: "2434-561X"},
	{in: "2434-561x"},
	{in: "", expectedError: "string must be a valid ISSN"},
	{in: "20493630", expectedError: "string must be a valid ISSN"},
	{in: "2049-3631", expectedError: "string must be a valid ISSN"},
	{in: "204-93630", expectedError: "string must be a valid ISSN"},
	{in: "2049-36X0", expectedError: "string must be a valid ISSN"},
	{in: "2049-363-", expectedError: "string must be a valid ISSN"},
}

func TestStringISSN(t *testing.T) {
	for _, tc := range stringISSNTestCases {
		err := StringISSN().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISSN))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISSN(b *testing.B) {
	for _, tc := range stringISSNTestCases {
		rule := StringISSN()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}
