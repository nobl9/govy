package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

// cspell:disable
var validBTCAddressTestCases = map[string]string{
	"p2pkh 1AGNa":             "1AGNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62i",
	"p2sh script hash 74f209": "3CMNFxN1oHBc4R1EpboAL5yzHGgE611Xou",
	"p2pkh 1Ax4g":             "1Ax4gZtb7gAit2TivwejZHYtNNLT18PUXJ",
	"p2sh 3QjYX":              "3QjYXhTkvuj8qPaXHTTWb5wjXhdsLAAWVy",
	"p2pkh 1C5bS":             "1C5bSj1iEGUgSTbziymG7Cn18ENQuT36vv",
	"p2sh 3AnNx":              "3AnNxabYGoTxYiTEZwFEnerUoeFXK2Zoks",
	"p2pkh 1Gqk4":             "1Gqk4Tv79P91Cc1STQtU3s1W6277M2CVWu",
	"p2sh 33vt8":              "33vt8ViH5jsr115AGkW6cEmEz9MpvJSwDk",
	"p2pkh 1JwMW":             "1JwMWBVLtiqtscbaRHai4pqHokhFCbtoB4",
	"p2sh 3QCzv":              "3QCzvfL4ZRvmJFiWWBVwxfdaNBT8EtxB5y",
	"p2pkh 19dca":             "19dcawoKcZdQz365WpXWMhX6QCUpR9SY4r",
	"p2sh 37Sp6":              "37Sp6Rv3y4kVd1nQ1JV5pfqXccHNyZm1x3",
	"p2pkh 13p1i":             "13p1ijLwsnrcuyqcTvJXkq2ASdXqcnEBLE",
	"p2sh script hash 5ece0c": "3ALJH9Y951VCGcVZYAdpA3KchoP9McEj1G",
	"p2pkh commonly used":     "1BoatSLRHtKNngkdXEeobR76b53LETtpyT",
	"p2sh commonly used":      "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy",
	"p2pkh zero payload":      "1111111111111111111114oLvT2",
	"p2pkh maximum payload":   "1QLbz7JHiBTspS962RLKV8GndWFwi5j6Qr",
	"p2sh zero payload":       "31h1vYVSYuKP6AhS86fbRdMw9XHieotbST",
	"p2sh maximum payload":    "3R2cuenjG5nFubqX9Wzuukdin2YfBbQ6Kw",
}

var invalidBTCAddressTestCases = map[string]string{
	"empty":                   "",
	"p2pkh checksum mutation": "1BoatSLRHtKNngkdXEeobR76b53LETtpyU",
	"p2sh checksum mutation":  "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLz",
	"testnet p2pkh reference": "mipcBbFg9gMiCh81Kj8tqqdgoZub1ZJRfn",
	"testnet p2pkh mo9n":      "mo9ncXisMeAoXwqcV5EWuyncbmCcQN4rVs",
	"testnet p2sh":            "2N2JD6wb56AfK4tfmM6PwdVmoYk2dCKf4Br",
	"unsupported version":     "3R2cuenjG5nFubqX9Wzuukdin2YfLYZyD1",
	"19 byte payload":         "111111111111111111117K4nzc",
	"21 byte payload":         "11111111111111111111116iowaD",
	"maximum textual length":  "11111111111111111111111111111111111",
	"forbidden digit zero":    "1BoatSLRHtKNngkdXEeobR76b53LETtpy0",
	"forbidden uppercase O":   "1BoatSLRHtKNngkdXEeobR76b53LETtpyO",
	"forbidden uppercase I":   "1BoatSLRHtKNngkdXEeobR76b53LETtpyI",
	"forbidden lowercase l":   "1BoatSLRHtKNngkdXEeobR76b53LETtpyl",
	"bech32 address":          "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
	"leading space":           " 1BoatSLRHtKNngkdXEeobR76b53LETtpyT",
	"trailing space":          "1BoatSLRHtKNngkdXEeobR76b53LETtpyT ",
	"trailing newline":        "1BoatSLRHtKNngkdXEeobR76b53LETtpyT\n",
	"Cyrillic A confusable":   "1\u0410GNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62i",
}

func TestStringBTCAddress(t *testing.T) {
	rule := StringBTCAddress()
	t.Run("valid addresses", func(t *testing.T) {
		for name, address := range validBTCAddressTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(address))
			})
		}
	})
	t.Run("invalid addresses", func(t *testing.T) {
		for name, address := range invalidBTCAddressTestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(address)
				assert.EqualError(t, err, "string must be a valid mainnet legacy Bitcoin Base58Check address")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBTCAddress))
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

var validBTCBech32AddressTestCases = map[string]string{
	"bip173 p2wpkh":                "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
	"bip173 p2wsh":                 "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3qccfmv3",
	"bip173 uppercase p2wpkh":      "BC1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
	"bitcoin core p2wpkh vyq0":     "bc1qvyq0cc6rahyvsazfdje0twl7ez82ndmuac2lhv",
	"bitcoin core p2wsh reference": "bc1qyucykdlhp62tezs0hagqury402qwhk589q80tqs5myh3rxq34nwqhkdhv7",
	"bitcoin core p2wpkh hxt0":     "bc1qhxt04s5xnpy0kxw4x99n5hpdf5pmtzpqs52es2",
	"bitcoin core p2wsh gc9l":      "bc1qgc9ljrvdf2e0zg9rmmq86xklqwfys7r6wptjlacdgrcdc7sa6ggqu4rrxf",
	"bitcoin core uppercase p2wsh": "BC1QRP33G0Q5C5TXSP9ARYSRX4K6ZDKFS4NCE4XJ0GDCCCEFVPYSXF3QCCFMV3",
}

var invalidBTCBech32AddressTestCases = map[string]string{
	"empty":                             "",
	"checksum mutation":                 "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t5",
	"testnet hrp":                       "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
	"mixed case":                        "bc1QW508D6QEJXTDG4Y5R3ZARVARY0C5XW7KV8F3T4",
	"bip173 witness v1":                 "bc1pw508d6qejxtdg4y5r3zarvary0c5xw7kw508d6qejxtdg4y5r3zarvary0c5xw7k7grplx",
	"bip350 witness v1":                 "bc1p0xlxvlhemja6c4dqv22uapctqupfhlxm9h8z3k2e72q4k9hcz7vqzk5jj0",
	"unknown hrp":                       "tc1qw508d6qejxtdg4y5r3zarvary0c5xw7kg3g4ty",
	"invalid witness version":           "BC13W508D6QEJXTDG4Y5R3ZARVARY0C5XW7KN40WF2",
	"short witness program":             "bc1qqqqsyydq4q",
	"long witness program":              "bc1qqqqsyqcyq5rqwzqfpg9scrgwpugpzysnzs23v9ccrydpk8qarc0jqwad46q",
	"invalid witness v0 program length": "BC1QR508D6QEJXTDG4Y5R3ZARVARYV98GJ9P",
	"excess zero padding":               "bc1qw508d6qejxtdg4y5r3zarvaryvqkyqvzl",
	"nonzero padding":                   "bc1qrp33g0q5c5txsp9arysrx4k6zdkfs4nce4xj0gdcccefvpysxf3p9waw3r",
	"empty data":                        "bc1gmk9yu",
	"bech32m checksum for witness v0":   "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kemeawh",
	"leading space":                     " bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
	"trailing space":                    "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4 ",
	"trailing newline":                  "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4\n",
	"zero width space":                  "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4\u200b",
}

func TestStringBTCBech32Address(t *testing.T) {
	rule := StringBTCBech32Address()
	t.Run("valid addresses", func(t *testing.T) {
		for name, address := range validBTCBech32AddressTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(address))
			})
		}
	})
	t.Run("invalid addresses", func(t *testing.T) {
		for name, address := range invalidBTCBech32AddressTestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(address)
				assert.EqualError(t, err, "string must be a valid mainnet Bitcoin Bech32 address")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBTCBech32Address))
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

var validETHAddressTestCases = map[string]string{
	"zero":                  "0x0000000000000000000000000000000000000000",
	"all lowercase de709":   "0xde709f2102306220921060314715629080e2fb77",
	"all lowercase 27b1":    "0x27b1fdb04752bbc536007a920d24acb045561c26",
	"all uppercase 529084":  "0x52908400098527886E0F7030069857D2E4169EE7",
	"all uppercase 8617":    "0x8617E340B3D01FA5F11F306F4090FD50E238070D",
	"all uppercase DE709":   "0xDE709F2102306220921060314715629080E2FB77",
	"eip55 5aAeb":           "0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
	"eip55 fB691":           "0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
	"eip55 dbF03":           "0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB",
	"eip55 D1220":           "0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",
	"additional mixed case": "0x5AEDA56215b167893e80B4fE645BA6d5Bab767DE",
}

var invalidETHAddressTestCases = map[string]string{
	"empty":                          "",
	"eip55 529084 case mutation":     "0x52908400098527886e0F7030069857D2E4169EE7",
	"additional mixed case mutation": "0x5AEDA56215b167893e80B4fE645BA6d5Bab767De",
	"eip55 5aAeb case mutation":      "0x5AAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
	"eip55 fB691 case mutation":      "0xfb6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
	"eip55 dbF03 case mutation":      "0xDbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB",
	"eip55 D1220 case mutation":      "0xd1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",
	"missing prefix":                 "de709f2102306220921060314715629080e2fb77",
	"uppercase prefix":               "0Xde709f2102306220921060314715629080e2fb77",
	"double prefix":                  "0x0xde709f2102306220921060314715629080e2fb77",
	"short payload":                  "0xde709f2102306220921060314715629080e2fb7",
	"nonhexadecimal character":       "0xde709f2102306220921060314715629080e2fb7g",
	"leading space":                  " 0xde709f2102306220921060314715629080e2fb77",
	"trailing space":                 "0xde709f2102306220921060314715629080e2fb77 ",
	"trailing newline":               "0xde709f2102306220921060314715629080e2fb77\n",
	"full width x prefix":            "0\uff58de709f2102306220921060314715629080e2fb77",
	"zero width space after prefix":  "0x\u200bde709f2102306220921060314715629080e2fb77",
}

// cspell:enable

func TestStringETHAddress(t *testing.T) {
	rule := StringETHAddress()
	t.Run("valid addresses", func(t *testing.T) {
		for name, address := range validETHAddressTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(address))
			})
		}
	})
	t.Run("invalid addresses", func(t *testing.T) {
		for name, address := range invalidETHAddressTestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(address)
				assert.EqualError(t, err, "string must be a valid Ethereum address (0x plus 40 hexadecimal characters)")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringETHAddress))
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
	testCaseGroups ...map[string]string,
) {
	b.Helper()

	for _, testCases := range testCaseGroups {
		for name, address := range testCases {
			b.Run(name, func(b *testing.B) {
				for b.Loop() {
					_ = rule.Validate(address)
				}
			})
		}
	}
}
