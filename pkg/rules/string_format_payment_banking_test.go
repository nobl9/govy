package rules

import (
	"crypto/sha256"
	"fmt"
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
	benchmarkStringPaymentBankingRule(b, rule, stringCreditCardTestCases)
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
	benchmarkStringPaymentBankingRule(b, rule, stringLuhnChecksumTestCases)
}

func TestStringPaymentCardProcessorFixtures(t *testing.T) {
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

	const (
		invalidProcessorCardNumber         = "4242424242424241"
		expectedUniqueProcessorCardNumbers = 264
	)
	union := make(map[string]struct{}, expectedUniqueProcessorCardNumbers)
	invalidOccurrences := 0
	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
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
				if input == invalidProcessorCardNumber {
					invalidOccurrences++
				}
				t.Run(input, func(t *testing.T) {
					for _, testRule := range rules {
						t.Run(testRule.name, func(t *testing.T) {
							err := testRule.rule.Validate(input)
							if input == invalidProcessorCardNumber {
								assertPaymentBankingRuleError(
									t,
									err,
									testRule.expectedError,
									testRule.errorCode,
								)
								return
							}
							assert.NoError(t, err)
						})
					}
				})
			}
			assert.Len(t, uniqueInputs, len(inputs))
		})
	}
	assert.Len(t, union, expectedUniqueProcessorCardNumbers)
	assert.Equal(t, 1, invalidOccurrences)
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
		"TC68 eleven character example":     "ABCDBE22XYZ",
		"alphanumeric party prefix":         "A1B2US33XXX",
		"zero in party suffix":              "ABCDUS0A",
		"one in party suffix":               "ABCDUS1A",
		"letter O in party suffix":          "ABCDUSAO",
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
	benchmarkStringPaymentBankingRule(b, rule, stringBICTestCases)
}

// cspell:disable
var stringBICISO93622014TestCases = stringPaymentBankingTestCases{
	validInputs: map[string]string{
		"eight characters":                  "DEUTDEFF",
		"eleven characters":                 "DEUTDEFF500",
		"branch placeholder":                "NEDSZAJJXXX",
		"ISO 9362 eight character example":  "ABCDFRPP",
		"ISO 9362 eleven character example": "WG11US335AB",
		"TC68 eight character example":      "ABCDBE22",
		"TC68 eleven character example":     "ABCDBE22XYZ",
		"alphanumeric party prefix":         "A1B2US33XXX",
		"zero in party suffix":              "ABCDUS0A",
		"one in party suffix":               "ABCDUS1A",
		"letter O in party suffix":          "ABCDUSAO",
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
	benchmarkStringPaymentBankingRule(b, rule, stringBICISO93622014TestCases)
}

func benchmarkStringPaymentBankingRule(
	b *testing.B,
	rule govy.Rule[string],
	testCases stringPaymentBankingTestCases,
) {
	b.Helper()
	for b.Loop() {
		for _, in := range testCases.validInputs {
			_ = rule.Validate(in)
		}
		for _, in := range testCases.invalidInputs {
			_ = rule.Validate(in)
		}
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
				assertPaymentBankingRuleError(t, rule.Validate(in), expectedError, errorCode)
			})
		}
	})
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
