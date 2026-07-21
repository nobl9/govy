package rules

import (
	"os"
	"strings"
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
		"visa":                      "4111111111111111",
		"visa alternate":            "4242424242424242",
		"mastercard":                "5555555555554444",
		"american express":          "378282246310005",
		"discover":                  "6011111111111117",
		"six-series sixteen digits": "6123451234567893",
		"diners club":               "36227206271667",
		"mastercard 2-series":       "2223003122003222",
		"JCB":                       "3566002020360505",
		"UnionPay nineteen digits":  "6205500000000000004",
		"visa decline test number":  "4000111111111115",
		"minimum thirteen digits":   "1000000000009",
		"maximum nineteen digits":   "1000000000000000009",
	},
	invalidInputs: map[string]string{
		"empty":                          "",
		"minimum length minus one":       "100000000008",
		"maximum length plus one":        "10000000000000000008",
		"failed thirteen digit checksum": "1000000000008",
		"failed nineteen digit checksum": "1000000000000000008",
		"all same digits":                "6666666666666",
		"leading space":                  " 1000000000009",
		"trailing space":                 "1000000000009 ",
		"trailing newline":               "1000000000009\n",
		"embedded spaces":                "4111 1111 1111 1111",
		"hyphens":                        "4111-1111-1111-1111",
		"full-width digit":               "10000000000０9",
		"alphabetic character":           "10000000000A9",
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
		"single zero":                "0",
		"four zeroes":                "0000",
		"twenty zeroes":              "00000000000000000000",
		"two digits ending in eight": "18",
		"two digits ending in nine":  "59",
		"sample number":              "79927398713",
		"six-series example":         "6123451234567893",
		"fifteen digit example":      "808401234567893",
		"nineteen digit example":     "6205500000000000004",
	},
	invalidInputs: map[string]string{
		"empty":                       "",
		"failed thirteen digit check": "6123451234567894",
		"failed fifteen digit check":  "808401234567894",
		"slash before zero":           "/0",
		"colon before zero":           ":0",
		"plus before zero":            "+0",
		"trailing space":              "0 ",
		"alphabetic character":        "79927A398713",
		"full-width digits":           "１２",
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

// cspell:disable
var stringBICTestCases = stringPaymentBankingTestCases{
	validInputs: map[string]string{
		"eight characters":                  "DEUTDEFF",
		"eleven characters":                 "DEUTDEFF500",
		"branch placeholder":                "NEDSZAJJXXX",
		"ISO 9362 eight character example":  "ABCDFRPP",
		"ISO 9362 eleven character example": "WG11US335AB",
		"TC68 eight character example":      "ABCDBE22",
		"alphanumeric party prefix":         "A1B2US33XXX",
		"Kosovo bank example":               "EKOMXKPR",
		"Kosovo bank identifier":            "BPBXXKPR",
	},
	invalidInputs: map[string]string{
		"empty":                       "",
		"lowercase":                   "abcdfrpp",
		"digit in country code":       "ABCD3R22",
		"user-assigned country AA":    "ABCDAA22",
		"user-assigned country QM":    "ABCDQM22",
		"user-assigned country QZ":    "ABCDQZ22",
		"user-assigned country XA":    "ABCDXA22",
		"user-assigned country XZ":    "ABCDXZ22",
		"user-assigned country ZZ":    "ABCDZZ22",
		"deleted country AN":          "ABCDAN22",
		"deleted country CS":          "ABCDCS22",
		"seven characters":            "ABCDFR2",
		"nine characters":             "ABCDFR22X",
		"ten characters":              "ABCDFR22XX",
		"twelve characters":           "ABCDFR22XXXX",
		"punctuation in party prefix": "ABC-FR22",
		"punctuation in party suffix": "ABCDFR2-",
		"punctuation in branch":       "ABCDFR22XY-",
		"trailing space":              "ABCDFR22 ",
	},
}

// cspell:enable

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

// cspell:disable
var stringBICISO93622014TestCases = stringPaymentBankingTestCases{
	validInputs: map[string]string{
		"eight characters":                 "DEUTDEFF",
		"eleven characters":                "DEUTDEFF500",
		"branch placeholder":               "NEDSZAJJXXX",
		"ISO 9362 eight character example": "ABCDFRPP",
		"TC68 eight character example":     "ABCDBE22",
		"Kosovo bank example":              "EKOMXKPR",
		"Kosovo bank identifier":           "BPBXXKPR",
	},
	invalidInputs: map[string]string{
		"empty":                       "",
		"lowercase":                   "abcdfrpp",
		"digit in country code":       "ABCD3R22",
		"user-assigned country AA":    "ABCDAA22",
		"user-assigned country QM":    "ABCDQM22",
		"user-assigned country QZ":    "ABCDQZ22",
		"user-assigned country XA":    "ABCDXA22",
		"user-assigned country XZ":    "ABCDXZ22",
		"user-assigned country ZZ":    "ABCDZZ22",
		"deleted country AN":          "ABCDAN22",
		"deleted country CS":          "ABCDCS22",
		"seven characters":            "ABCDFR2",
		"nine characters":             "ABCDFR22X",
		"ten characters":              "ABCDFR22XX",
		"twelve characters":           "ABCDFR22XXXX",
		"punctuation in party prefix": "ABC-FR22",
		"punctuation in party suffix": "ABCDFR2-",
		"punctuation in branch":       "ABCDFR22XY-",
		"trailing space":              "ABCDFR22 ",
	},
}

// cspell:enable

func TestStringBICISO93622014(t *testing.T) {
	assertPaymentBankingRule(
		t,
		StringBICISO93622014(),
		stringBICISO93622014TestCases,
		"string must be a valid ISO 9362:2014 Business Identifier Code (BIC)",
		ErrorCodeStringBICISO93622014,
	)
}

// cspell:disable
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

// cspell:enable

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
		for name, in := range testCases.invalidInputs {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(in)
				assert.Require(t, assert.EqualError(t, err, expectedError))
				assert.Require(t, assert.IsType[*govy.RuleError](t, err))
				ruleError := err.(*govy.RuleError)
				assert.Equal(t, expectedError, ruleError.Description)
				assert.Equal(t, errorCode, ruleError.Code)
			})
		}
	})
}

func readISOAlpha2CountryCodes(t *testing.T) []string {
	t.Helper()
	data, err := os.ReadFile("testdata/iso_3166_1_alpha2_2026-07-21.txt")
	assert.Require(t, assert.NoError(t, err))

	var codes []string
	for line := range strings.Lines(string(data)) {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		codes = append(codes, strings.Fields(line)...)
	}
	return codes
}
