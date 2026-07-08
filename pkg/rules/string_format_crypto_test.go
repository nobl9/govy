package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestStringBTCAddress(t *testing.T) {
	for _, tc := range stringBTCAddressTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			assertStringFormatRule(t, StringBTCAddress(), ErrorCodeStringBTCAddress, tc)
		})
	}
}

func BenchmarkStringBTCAddress(b *testing.B) {
	benchmarkStringFormatRule(b, StringBTCAddress(), stringBTCAddressTestCases())
}

func TestStringBTCBech32Address(t *testing.T) {
	for _, tc := range stringBTCBech32AddressTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			assertStringFormatRule(t, StringBTCBech32Address(), ErrorCodeStringBTCBech32Address, tc)
		})
	}
}

func BenchmarkStringBTCBech32Address(b *testing.B) {
	benchmarkStringFormatRule(b, StringBTCBech32Address(), stringBTCBech32AddressTestCases())
}

func TestStringETHAddress(t *testing.T) {
	for _, tc := range stringETHAddressTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			assertStringFormatRule(t, StringETHAddress(), ErrorCodeStringETHAddress, tc)
		})
	}
}

func BenchmarkStringETHAddress(b *testing.B) {
	benchmarkStringFormatRule(b, StringETHAddress(), stringETHAddressTestCases())
}

type stringFormatRuleTestCase struct {
	name          string
	in            string
	expectedError string
}

func stringBTCAddressTestCases() []stringFormatRuleTestCase {
	const err = "string must be a valid mainnet legacy Bitcoin Base58Check address"
	return []stringFormatRuleTestCase{
		{
			name: "p2pkh",
			in:   "1BoatSLRHtKNngkdXEeobR76b53LETtpyT",
		},
		{
			name: "p2sh",
			in:   "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy",
		},
		{
			name:          "bad checksum",
			in:            "1BoatSLRHtKNngkdXEeobR76b53LETtpyU",
			expectedError: err,
		},
		{
			name:          "testnet address",
			in:            "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
			expectedError: err,
		},
		{
			name:          "invalid base58 character",
			in:            "1BoatSLRHtKNngkdXEeobR76b53LETtpy0",
			expectedError: err,
		},
		{
			name:          "bech32 address",
			in:            "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			expectedError: err,
		},
		{
			name:          "empty",
			expectedError: err,
		},
	}
}

func stringBTCBech32AddressTestCases() []stringFormatRuleTestCase {
	const err = "string must be a valid mainnet Bitcoin Bech32 address"
	return []stringFormatRuleTestCase{
		{
			name: "p2wpkh",
			in:   "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
		},
		{
			name: "p2wsh",
			in:   "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
		},
		{
			name: "uppercase",
			in:   "BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
		},
		{
			name:          "bad checksum",
			in:            "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t5",
			expectedError: err,
		},
		{
			name:          "testnet address",
			in:            "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
			expectedError: err,
		},
		{
			name:          "mixed case",
			in:            "bc1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
			expectedError: err,
		},
		{
			name:          "bech32m taproot address",
			in:            "bc1p5cyxnuxmeuwuvkwfem96llyxfvvwq0rscl0g7k",
			expectedError: err,
		},
		{
			name:          "empty",
			expectedError: err,
		},
	}
}

func stringETHAddressTestCases() []stringFormatRuleTestCase {
	const err = "string must be a valid Ethereum address (0x plus 40 hexadecimal characters)"
	return []stringFormatRuleTestCase{
		{
			name: "zero address",
			in:   "0x0000000000000000000000000000000000000000",
		},
		{
			name: "lowercase",
			in:   "0xde709f2102306220921060314715629080e2fb77",
		},
		{
			name: "mixed case syntax only",
			in:   "0x52908400098527886E0F7030069857D2E4169EE8",
		},
		{
			name:          "missing prefix",
			in:            "de709f2102306220921060314715629080e2fb77",
			expectedError: err,
		},
		{
			name:          "uppercase prefix",
			in:            "0Xde709f2102306220921060314715629080e2fb77",
			expectedError: err,
		},
		{
			name:          "too short",
			in:            "0xde709f2102306220921060314715629080e2fb7",
			expectedError: err,
		},
		{
			name:          "not hex",
			in:            "0xde709f2102306220921060314715629080e2fb7g",
			expectedError: err,
		},
		{
			name:          "empty",
			expectedError: err,
		},
	}
}

func assertStringFormatRule(
	t *testing.T,
	rule govy.Rule[string],
	errorCode govy.ErrorCode,
	tc stringFormatRuleTestCase,
) {
	t.Helper()

	err := rule.Validate(tc.in)
	if tc.expectedError != "" {
		assert.EqualError(t, err, tc.expectedError)
		assert.True(t, govy.HasErrorCode(err, errorCode))
		return
	}
	assert.NoError(t, err)
}

func benchmarkStringFormatRule(
	b *testing.B,
	rule govy.Rule[string],
	testCases []stringFormatRuleTestCase,
) {
	b.Helper()

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for range b.N {
				_ = rule.Validate(tc.in)
			}
		})
	}
}
