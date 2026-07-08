package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestStringBase64(t *testing.T) {
	runStringFormatRuleTest(t, StringBase64(), ErrorCodeStringBase64, map[string]stringFormatRuleTestCase{
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
	runStringFormatRuleTest(t, StringBase64URL(), ErrorCodeStringBase64URL, map[string]stringFormatRuleTestCase{
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
	runStringFormatRuleTest(t, StringBase64RawURL(), ErrorCodeStringBase64RawURL, map[string]stringFormatRuleTestCase{
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
	})
}

func TestStringHexadecimal(t *testing.T) {
	runStringFormatRuleTest(t, StringHexadecimal(), ErrorCodeStringHexadecimal, map[string]stringFormatRuleTestCase{
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

func TestStringMD5(t *testing.T) {
	runStringFormatRuleTest(t, StringMD5(), ErrorCodeStringMD5, map[string]stringFormatRuleTestCase{
		// cspell:disable
		"lowercase digest": {in: "d41d8cd98f00b204e9800998ecf8427e"},
		"uppercase digest": {
			in:            "D41D8CD98F00B204E9800998ECF8427E",
			expectedError: "string must be a valid lowercase MD5 hexadecimal digest",
		},
		"too short": {
			in:            "d41d8cd98f00b204e9800998ecf8427",
			expectedError: "string must be a valid lowercase MD5 hexadecimal digest",
		},
		"non hexadecimal": {
			in:            "d41d8cd98f00b204e9800998ecf8427g",
			expectedError: "string must be a valid lowercase MD5 hexadecimal digest",
		},
		// cspell:enable
	})
}

func TestStringSHA256(t *testing.T) {
	runStringFormatRuleTest(t, StringSHA256(), ErrorCodeStringSHA256, map[string]stringFormatRuleTestCase{
		// cspell:disable
		"lowercase digest": {in: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		"uppercase digest": {
			in:            "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
			expectedError: "string must be a valid lowercase SHA-256 hexadecimal digest",
		},
		"too short": {
			in:            "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85",
			expectedError: "string must be a valid lowercase SHA-256 hexadecimal digest",
		},
		"non hexadecimal": {
			in:            "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b85g",
			expectedError: "string must be a valid lowercase SHA-256 hexadecimal digest",
		},
		// cspell:enable
	})
}

func TestStringSHA384(t *testing.T) {
	runStringFormatRuleTest(t, StringSHA384(), ErrorCodeStringSHA384, map[string]stringFormatRuleTestCase{
		// cspell:disable
		"lowercase digest": {
			in: "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b",
		},
		"uppercase digest": {
			in:            "38B060A751AC96384CD9327EB1B1E36A21FDB71114BE07434C0CC7BF63F6E1DA274EDEBFE76F65FBD51AD2F14898B95B",
			expectedError: "string must be a valid lowercase SHA-384 hexadecimal digest",
		},
		"too short": {
			in:            "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95",
			expectedError: "string must be a valid lowercase SHA-384 hexadecimal digest",
		},
		"non hexadecimal": {
			in:            "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95g",
			expectedError: "string must be a valid lowercase SHA-384 hexadecimal digest",
		},
		// cspell:enable
	})
}

func TestStringSHA512(t *testing.T) {
	runStringFormatRuleTest(t, StringSHA512(), ErrorCodeStringSHA512, map[string]stringFormatRuleTestCase{
		// cspell:disable
		"lowercase digest": {
			in: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
		},
		"uppercase digest": {
			in:            "CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E",
			expectedError: "string must be a valid lowercase SHA-512 hexadecimal digest",
		},
		"too short": {
			in:            "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3",
			expectedError: "string must be a valid lowercase SHA-512 hexadecimal digest",
		},
		"non hexadecimal": {
			in:            "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927dag",
			expectedError: "string must be a valid lowercase SHA-512 hexadecimal digest",
		},
		// cspell:enable
	})
}

func BenchmarkStringBase64(b *testing.B) {
	benchmarkStringFormatRule(b, StringBase64(), []string{"", "Zg==", "Zm9v", "+/8=", "Zg"})
}

func BenchmarkStringBase64URL(b *testing.B) {
	benchmarkStringFormatRule(b, StringBase64URL(), []string{"", "Zg==", "Zm9v", "-_8=", "+/8="})
}

func BenchmarkStringBase64RawURL(b *testing.B) {
	benchmarkStringFormatRule(b, StringBase64RawURL(), []string{"", "Zg", "Zm9v", "-_8", "Zg=="})
}

func BenchmarkStringHexadecimal(b *testing.B) {
	benchmarkStringFormatRule(b, StringHexadecimal(), []string{"0", "deadbeef", "0XABCDEF", "0x"})
}

func BenchmarkStringMD5(b *testing.B) {
	benchmarkStringFormatRule(b, StringMD5(), []string{
		"d41d8cd98f00b204e9800998ecf8427e",
		"D41D8CD98F00B204E9800998ECF8427E",
	})
}

func BenchmarkStringSHA256(b *testing.B) {
	benchmarkStringFormatRule(b, StringSHA256(), []string{
		"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		"E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
	})
}

func BenchmarkStringSHA384(b *testing.B) {
	benchmarkStringFormatRule(b, StringSHA384(), []string{
		"38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b",
		"38B060A751AC96384CD9327EB1B1E36A21FDB71114BE07434C0CC7BF63F6E1DA274EDEBFE76F65FBD51AD2F14898B95B",
	})
}

func BenchmarkStringSHA512(b *testing.B) {
	benchmarkStringFormatRule(b, StringSHA512(), []string{
		"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
		"CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E",
	})
}

type stringFormatRuleTestCase struct {
	in            string
	expectedError string
}

func runStringFormatRuleTest(
	t *testing.T,
	rule govy.Rule[string],
	errorCode govy.ErrorCode,
	testCases map[string]stringFormatRuleTestCase,
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

func benchmarkStringFormatRule(b *testing.B, rule govy.Rule[string], inputs []string) {
	b.Helper()
	for b.Loop() {
		for _, in := range inputs {
			_ = rule.Validate(in)
		}
	}
}
