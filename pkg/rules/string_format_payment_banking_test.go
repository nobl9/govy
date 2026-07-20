package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

type stringPaymentBankingTestCases struct {
	validInputs   map[string]string
	invalidInputs map[string]string
}

var stringCreditCardTestCases = stringPaymentBankingTestCases{
	validInputs: map[string]string{
		"visa":             "4111111111111111",
		"visa alternate":   "4242424242424242",
		"mastercard":       "5555555555554444",
		"american express": "378282246310005",
		"discover":         "6011111111111117",
		"thirteen digits":  "4222222222222",
		"nineteen digits":  "4000000000000000006",
	},
	invalidInputs: map[string]string{
		"failed checksum": "4111111111111112",
		"spaces":          "4111 1111 1111 1111",
		"hyphens":         "4111-1111-1111-1111",
		"too short":       "411111111111",
		"too long":        "40000000000000000006",
		"all same digits": "0000000000000",
	},
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

func BenchmarkStringCreditCard(b *testing.B) {
	rule := StringCreditCard()
	for b.Loop() {
		_ = rule.Validate("4111111111111111")
	}
}

var stringLuhnChecksumTestCases = stringPaymentBankingTestCases{
	validInputs: map[string]string{
		"single zero":   "0",
		"all zeroes":    "0000",
		"sample number": "79927398713",
		"card number":   "4111111111111111",
	},
	invalidInputs: map[string]string{
		"empty":        "",
		"failed check": "79927398710",
		"non digit":    "7992739871A",
	},
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

func BenchmarkStringLuhnChecksum(b *testing.B) {
	rule := StringLuhnChecksum()
	for b.Loop() {
		_ = rule.Validate("79927398713")
	}
}

var stringBICTestCases = stringPaymentBankingTestCases{
	validInputs: map[string]string{
		"eight characters":      "DEUTDEFF",
		"eleven characters":     "DEUTDEFF500",
		"branch placeholder":    "NEDSZAJJXXX",
		"kosovo country code":   "DEUTXKFF",
		"branch XXX":            "DEUTDEFFXXX",
		"digits in institution": "A1B2US33XXX",
	},
	invalidInputs: map[string]string{
		"lowercase":             "deutdeff",
		"digit in country code": "DEUTD3FF",
		"unknown country code":  "DEUTZZFF",
		"location first zero":   "DEUTDE0F",
		"location first one":    "DEUTDE1F",
		"location second O":     "DEUTDEFO",
		"branch starts X":       "DEUTDEFFXAA",
		"too short":             "DEUTDEF",
		"invalid branch length": "DEUTDEFF50",
		"too long":              "DEUTDEFF5000",
		"punctuation":           "DEUTDEFF50!",
	},
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

func BenchmarkStringBIC(b *testing.B) {
	rule := StringBIC()
	for b.Loop() {
		_ = rule.Validate("DEUTDEFF500")
	}
}

var stringBICISO93622014TestCases = stringPaymentBankingTestCases{
	validInputs: map[string]string{
		"eight characters":    "DEUTDEFF",
		"eleven characters":   "DEUTDEFF500",
		"branch placeholder":  "NEDSZAJJXXX",
		"kosovo country code": "DEUTXKFF",
		"branch XXX":          "DEUTDEFFXXX",
	},
	invalidInputs: map[string]string{
		"digits in institution": "A1B2US33XXX",
		"lowercase":             "deutdeff",
		"digit in country code": "DEUTD3FF",
		"unknown country code":  "DEUTZZFF",
		"location first zero":   "DEUTDE0F",
		"location first one":    "DEUTDE1F",
		"location second O":     "DEUTDEFO",
		"branch starts X":       "DEUTDEFFXAA",
		"too short":             "DEUTDEF",
		"punctuation":           "DEUTDEFF50!",
	},
}

func TestStringBICISO93622014(t *testing.T) {
	assertPaymentBankingRule(
		t,
		StringBICISO93622014(),
		stringBICISO93622014TestCases,
		"string must be a valid ISO 9362:2014 Business Identifier Code (BIC)",
		ErrorCodeStringBICISO93622014,
	)
}

func BenchmarkStringBICISO93622014(b *testing.B) {
	rule := StringBICISO93622014()
	for b.Loop() {
		_ = rule.Validate("DEUTDEFF500")
	}
}

func assertPaymentBankingRule(
	t *testing.T,
	rule govy.Rule[string],
	testCases stringPaymentBankingTestCases,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	t.Run("valid inputs", func(t *testing.T) {
		for name, in := range testCases.validInputs {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(in))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, in := range testCases.invalidInputs {
			err := rule.Validate(in)
			assert.EqualError(t, err, expectedError)
			assert.Require(t, assert.IsType[*govy.RuleError](t, err))
			assert.Equal(t, expectedError, err.(*govy.RuleError).Description)
			assert.True(t, govy.HasErrorCode(err, errorCode))
			break
		}

		for name, in := range testCases.invalidInputs {
			t.Run(name, func(t *testing.T) {
				assert.Error(t, rule.Validate(in))
			})
		}
	})
}
