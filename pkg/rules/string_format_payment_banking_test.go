package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestStringCreditCard(t *testing.T) {
	rule := StringCreditCard()
	for name, tc := range stringCreditCardTestCases {
		t.Run(name, func(t *testing.T) {
			err := rule.Validate(tc.in)
			assertPaymentBankingRule(t, err, tc.expectedError, ErrorCodeStringCreditCard)
		})
	}
}

func TestStringLuhnChecksum(t *testing.T) {
	rule := StringLuhnChecksum()
	for name, tc := range stringLuhnChecksumTestCases {
		t.Run(name, func(t *testing.T) {
			err := rule.Validate(tc.in)
			assertPaymentBankingRule(t, err, tc.expectedError, ErrorCodeStringLuhnChecksum)
		})
	}
}

func TestStringBIC(t *testing.T) {
	rule := StringBIC()
	for name, tc := range stringBICTestCases {
		t.Run(name, func(t *testing.T) {
			err := rule.Validate(tc.in)
			assertPaymentBankingRule(t, err, tc.expectedError, ErrorCodeStringBIC)
		})
	}
}

func TestStringBICISO93622014(t *testing.T) {
	rule := StringBICISO93622014()
	for name, tc := range stringBICISO93622014TestCases {
		t.Run(name, func(t *testing.T) {
			err := rule.Validate(tc.in)
			assertPaymentBankingRule(t, err, tc.expectedError, ErrorCodeStringBICISO93622014)
		})
	}
}

func BenchmarkStringCreditCard(b *testing.B) {
	rule := StringCreditCard()
	for b.Loop() {
		_ = rule.Validate("4111111111111111")
	}
}

func BenchmarkStringLuhnChecksum(b *testing.B) {
	rule := StringLuhnChecksum()
	for b.Loop() {
		_ = rule.Validate("79927398713")
	}
}

func BenchmarkStringBIC(b *testing.B) {
	rule := StringBIC()
	for b.Loop() {
		_ = rule.Validate("DEUTDEFF500")
	}
}

func BenchmarkStringBICISO93622014(b *testing.B) {
	rule := StringBICISO93622014()
	for b.Loop() {
		_ = rule.Validate("DEUTDEFF500")
	}
}

type stringPaymentBankingTestCase struct {
	in            string
	expectedError string
}

var stringCreditCardTestCases = map[string]stringPaymentBankingTestCase{
	"visa":             {in: "4111111111111111"},
	"visa alternate":   {in: "4242424242424242"},
	"mastercard":       {in: "5555555555554444"},
	"american express": {in: "378282246310005"},
	"discover":         {in: "6011111111111117"},
	"thirteen digits":  {in: "4222222222222"},
	"nineteen digits":  {in: "4000000000000000006"},
	"failed checksum":  {in: "4111111111111112", expectedError: "string must be a valid credit card number"},
	"spaces":           {in: "4111 1111 1111 1111", expectedError: "string must be a valid credit card number"},
	"hyphens":          {in: "4111-1111-1111-1111", expectedError: "string must be a valid credit card number"},
	"too short":        {in: "411111111111", expectedError: "string must be a valid credit card number"},
	"too long":         {in: "40000000000000000006", expectedError: "string must be a valid credit card number"},
	"all same digits":  {in: "0000000000000", expectedError: "string must be a valid credit card number"},
}

var stringLuhnChecksumTestCases = map[string]stringPaymentBankingTestCase{
	"single zero":   {in: "0"},
	"all zeroes":    {in: "0000"},
	"sample number": {in: "79927398713"},
	"card number":   {in: "4111111111111111"},
	"empty":         {in: "", expectedError: "string must pass the Luhn checksum"},
	"failed check":  {in: "79927398710", expectedError: "string must pass the Luhn checksum"},
	"non digit":     {in: "7992739871A", expectedError: "string must pass the Luhn checksum"},
}

var stringBICTestCases = map[string]stringPaymentBankingTestCase{
	"eight characters":      {in: "DEUTDEFF"},
	"eleven characters":     {in: "DEUTDEFF500"},
	"branch placeholder":    {in: "NEDSZAJJXXX"},
	"kosovo country code":   {in: "DEUTXKFF"},
	"branch XXX":            {in: "DEUTDEFFXXX"},
	"digits in institution": {in: "A1B2US33XXX"},
	"lowercase":             {in: "deutdeff", expectedError: "string must be a valid BIC"},
	"digit in country code": {in: "DEUTD3FF", expectedError: "string must be a valid BIC"},
	"unknown country code":  {in: "DEUTZZFF", expectedError: "string must be a valid BIC"},
	"location first zero":   {in: "DEUTDE0F", expectedError: "string must be a valid BIC"},
	"location first one":    {in: "DEUTDE1F", expectedError: "string must be a valid BIC"},
	"location second O":     {in: "DEUTDEFO", expectedError: "string must be a valid BIC"},
	"branch starts X":       {in: "DEUTDEFFXAA", expectedError: "string must be a valid BIC"},
	"too short":             {in: "DEUTDEF", expectedError: "string must be a valid BIC"},
	"invalid branch length": {in: "DEUTDEFF50", expectedError: "string must be a valid BIC"},
	"too long":              {in: "DEUTDEFF5000", expectedError: "string must be a valid BIC"},
	"punctuation":           {in: "DEUTDEFF50!", expectedError: "string must be a valid BIC"},
}

var stringBICISO93622014TestCases = map[string]stringPaymentBankingTestCase{
	"eight characters":      {in: "DEUTDEFF"},
	"eleven characters":     {in: "DEUTDEFF500"},
	"branch placeholder":    {in: "NEDSZAJJXXX"},
	"kosovo country code":   {in: "DEUTXKFF"},
	"branch XXX":            {in: "DEUTDEFFXXX"},
	"digits in institution": {in: "A1B2US33XXX", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"lowercase":             {in: "deutdeff", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"digit in country code": {in: "DEUTD3FF", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"unknown country code":  {in: "DEUTZZFF", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"location first zero":   {in: "DEUTDE0F", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"location first one":    {in: "DEUTDE1F", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"location second O":     {in: "DEUTDEFO", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"branch starts X":       {in: "DEUTDEFFXAA", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"too short":             {in: "DEUTDEF", expectedError: "string must be a valid ISO 9362:2014 BIC"},
	"punctuation":           {in: "DEUTDEFF50!", expectedError: "string must be a valid ISO 9362:2014 BIC"},
}

func assertPaymentBankingRule(
	t *testing.T,
	err error,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	if expectedError != "" {
		assert.EqualError(t, err, expectedError)
		assert.True(t, govy.HasErrorCode(err, errorCode))
		return
	}
	assert.NoError(t, err)
}
