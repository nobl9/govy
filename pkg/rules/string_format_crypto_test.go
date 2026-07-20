package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

var validBTCAddressTestCases = []string{
	"1BoatSLRHtKNngkdXEeobR76b53LETtpyT",
	"3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy",
}

var invalidBTCAddressTestCases = []string{
	"1BoatSLRHtKNngkdXEeobR76b53LETtpyU",
	"mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
	"1BoatSLRHtKNngkdXEeobR76b53LETtpy0",
	"bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
	"",
}

func TestStringBTCAddress(t *testing.T) {
	rule := StringBTCAddress()
	t.Run("valid addresses", func(t *testing.T) {
		for _, address := range validBTCAddressTestCases {
			t.Run(fmt.Sprintf("%q", address), func(t *testing.T) {
				assert.NoError(t, rule.Validate(address))
			})
		}
	})
	t.Run("invalid addresses", func(t *testing.T) {
		err := rule.Validate(invalidBTCAddressTestCases[0])
		assert.EqualError(t, err, "string must be a valid mainnet legacy Bitcoin Base58Check address")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBTCAddress))

		for _, address := range invalidBTCAddressTestCases {
			t.Run(fmt.Sprintf("%q", address), func(t *testing.T) {
				assert.Error(t, rule.Validate(address))
			})
		}
	})
}

func BenchmarkStringBTCAddress(b *testing.B) {
	benchmarkStringFormatRule(
		b,
		StringBTCAddress(),
		validBTCAddressTestCases,
		invalidBTCAddressTestCases,
	)
}

var validBTCBech32AddressTestCases = []string{
	"bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
	"bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
	"BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
}

var invalidBTCBech32AddressTestCases = []string{
	"bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t5",
	"tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
	"bc1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
	"bc1p5cyxnuxmeuwuvkwfem96llyxfvvwq0rscl0g7k",
	"",
}

func TestStringBTCBech32Address(t *testing.T) {
	rule := StringBTCBech32Address()
	t.Run("valid addresses", func(t *testing.T) {
		for _, address := range validBTCBech32AddressTestCases {
			t.Run(fmt.Sprintf("%q", address), func(t *testing.T) {
				assert.NoError(t, rule.Validate(address))
			})
		}
	})
	t.Run("invalid addresses", func(t *testing.T) {
		err := rule.Validate(invalidBTCBech32AddressTestCases[0])
		assert.EqualError(t, err, "string must be a valid mainnet Bitcoin Bech32 address")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBTCBech32Address))

		for _, address := range invalidBTCBech32AddressTestCases {
			t.Run(fmt.Sprintf("%q", address), func(t *testing.T) {
				assert.Error(t, rule.Validate(address))
			})
		}
	})
}

func BenchmarkStringBTCBech32Address(b *testing.B) {
	benchmarkStringFormatRule(
		b,
		StringBTCBech32Address(),
		validBTCBech32AddressTestCases,
		invalidBTCBech32AddressTestCases,
	)
}

var validETHAddressTestCases = []string{
	"0x0000000000000000000000000000000000000000",
	"0xde709f2102306220921060314715629080e2fb77",
	"0x52908400098527886E0F7030069857D2E4169EE7",
	"0x5AEDA56215b167893e80B4fE645BA6d5Bab767DE",
	"0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
	"0xDE709F2102306220921060314715629080E2FB77",
}

var invalidETHAddressTestCases = []string{
	"0x52908400098527886e0F7030069857D2E4169EE7",
	"0x5AEDA56215b167893e80B4fE645BA6d5Bab767De",
	"de709f2102306220921060314715629080e2fb77",
	"0Xde709f2102306220921060314715629080e2fb77",
	"0xde709f2102306220921060314715629080e2fb7",
	"0xde709f2102306220921060314715629080e2fb7g",
	"",
}

func TestStringETHAddress(t *testing.T) {
	rule := StringETHAddress()
	t.Run("valid addresses", func(t *testing.T) {
		for _, address := range validETHAddressTestCases {
			t.Run(fmt.Sprintf("%q", address), func(t *testing.T) {
				assert.NoError(t, rule.Validate(address))
			})
		}
	})
	t.Run("invalid addresses", func(t *testing.T) {
		err := rule.Validate(invalidETHAddressTestCases[0])
		assert.EqualError(t, err, "string must be a valid Ethereum address (0x plus 40 hexadecimal characters)")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringETHAddress))

		for _, address := range invalidETHAddressTestCases {
			t.Run(fmt.Sprintf("%q", address), func(t *testing.T) {
				assert.Error(t, rule.Validate(address))
			})
		}
	})
}

func BenchmarkStringETHAddress(b *testing.B) {
	benchmarkStringFormatRule(
		b,
		StringETHAddress(),
		validETHAddressTestCases,
		invalidETHAddressTestCases,
	)
}

func benchmarkStringFormatRule(
	b *testing.B,
	rule govy.Rule[string],
	testCaseGroups ...[]string,
) {
	b.Helper()

	for _, testCases := range testCaseGroups {
		for _, address := range testCases {
			b.Run(fmt.Sprintf("%q", address), func(b *testing.B) {
				for b.Loop() {
					_ = rule.Validate(address)
				}
			})
		}
	}
}
