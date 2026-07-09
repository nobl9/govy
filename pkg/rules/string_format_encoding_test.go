package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestStringBase64(t *testing.T) {
	runStringEncodingRuleTest(t, StringBase64(), ErrorCodeStringBase64, map[string]stringEncodingRuleTestCase{
		// cspell:disable
		"empty":             {in: ""},
		"single byte":       {in: "Zg=="},
		"two bytes":         {in: "Zm8="},
		"three bytes":       {in: "Zm9v"},
		"standard alphabet": {in: "+/8="},
		"missing padding": {
			in:            "Zg",
			expectedError: "string must be a valid standard padded base64 value",
		},
		"url alphabet": {
			in:            "-_8=",
			expectedError: "string must be a valid standard padded base64 value",
		},
		"new line": {
			in:            "Zm9v\n",
			expectedError: "string must be a valid standard padded base64 value",
		},
		"extra padding": {
			in:            "Zm9v====",
			expectedError: "string must be a valid standard padded base64 value",
		},
		// cspell:enable
	})
}

func TestStringBase64URL(t *testing.T) {
	runStringEncodingRuleTest(t, StringBase64URL(), ErrorCodeStringBase64URL, map[string]stringEncodingRuleTestCase{
		// cspell:disable
		"empty":        {in: ""},
		"single byte":  {in: "Zg=="},
		"two bytes":    {in: "Zm8="},
		"three bytes":  {in: "Zm9v"},
		"url alphabet": {in: "-_8="},
		"standard alphabet": {
			in:            "+/8=",
			expectedError: "string must be a valid URL-safe padded base64 value",
		},
		"missing padding": {
			in:            "Zm8",
			expectedError: "string must be a valid URL-safe padded base64 value",
		},
		"new line": {
			in:            "Zm9v\n",
			expectedError: "string must be a valid URL-safe padded base64 value",
		},
		"extra padding": {
			in:            "Zm9v====",
			expectedError: "string must be a valid URL-safe padded base64 value",
		},
		// cspell:enable
	})
}

func TestStringBase64RawURL(t *testing.T) {
	runStringEncodingRuleTest(
		t,
		StringBase64RawURL(),
		ErrorCodeStringBase64RawURL,
		map[string]stringEncodingRuleTestCase{
			// cspell:disable
			"empty":        {in: ""},
			"single byte":  {in: "Zg"},
			"two bytes":    {in: "Zm8"},
			"three bytes":  {in: "Zm9v"},
			"url alphabet": {in: "-_8"},
			"padded": {
				in:            "Zg==",
				expectedError: "string must be a valid URL-safe base64 value without padding",
			},
			"standard alphabet": {
				in:            "+/8",
				expectedError: "string must be a valid URL-safe base64 value without padding",
			},
			"invalid length": {
				in:            "A",
				expectedError: "string must be a valid URL-safe base64 value without padding",
			},
			"new line": {
				in:            "Zm9v\n",
				expectedError: "string must be a valid URL-safe base64 value without padding",
			},
			// cspell:enable
		},
	)
}

func TestStringHexadecimal(t *testing.T) {
	runStringEncodingRuleTest(t, StringHexadecimal(), ErrorCodeStringHexadecimal, map[string]stringEncodingRuleTestCase{
		// cspell:disable
		"single digit":     {in: "0"},
		"lowercase":        {in: "deadbeef"},
		"uppercase":        {in: "DEADBEEF"},
		"lowercase prefix": {in: "0xdeadBEEF"},
		"uppercase prefix": {in: "0XABCDEF"},
		"empty": {
			in:            "",
			expectedError: "string must be a valid hexadecimal value",
		},
		"prefix only": {
			in:            "0x",
			expectedError: "string must be a valid hexadecimal value",
		},
		"invalid digit": {
			in:            "0xabcdefg",
			expectedError: "string must be a valid hexadecimal value",
		},
		"minus sign": {
			in:            "-0x1",
			expectedError: "string must be a valid hexadecimal value",
		},
		// cspell:enable
	})
}

func BenchmarkStringBase64(b *testing.B) {
	benchmarkStringEncodingRule(b, StringBase64(), []string{"", "Zg==", "Zm9v", "+/8=", "Zg"})
}

func BenchmarkStringBase64URL(b *testing.B) {
	benchmarkStringEncodingRule(b, StringBase64URL(), []string{"", "Zg==", "Zm9v", "-_8=", "+/8="})
}

func BenchmarkStringBase64RawURL(b *testing.B) {
	benchmarkStringEncodingRule(b, StringBase64RawURL(), []string{"", "Zg", "Zm9v", "-_8", "Zg=="})
}

func BenchmarkStringHexadecimal(b *testing.B) {
	benchmarkStringEncodingRule(b, StringHexadecimal(), []string{"0", "deadbeef", "0XABCDEF", "0x"})
}

type stringEncodingRuleTestCase struct {
	in            string
	expectedError string
}

func runStringEncodingRuleTest(
	t *testing.T,
	rule govy.Rule[string],
	errorCode govy.ErrorCode,
	testCases map[string]stringEncodingRuleTestCase,
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

func benchmarkStringEncodingRule(b *testing.B, rule govy.Rule[string], inputs []string) {
	b.Helper()
	for b.Loop() {
		for _, in := range inputs {
			_ = rule.Validate(in)
		}
	}
}
