package rules

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

// cspell:disable
var validBTCAddressTestCases = map[string]string{
	// Fourteen historical regressions; not attributed to the pinned Bitcoin Core source.
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
	// Common examples and locally constructed boundary cases.
	"p2pkh commonly used":   "1BoatSLRHtKNngkdXEeobR76b53LETtpyT",
	"p2sh commonly used":    "3J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy",
	"p2pkh zero payload":    "1111111111111111111114oLvT2",
	"p2pkh maximum payload": "1QLbz7JHiBTspS962RLKV8GndWFwi5j6Qr",
	"p2sh zero payload":     "31h1vYVSYuKP6AhS86fbRdMw9XHieotbST",
	"p2sh maximum payload":  "3R2cuenjG5nFubqX9Wzuukdin2YfBbQ6Kw",
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

// The fixtures are exact copies of Bitcoin Core's key_io_valid.json and
// key_io_invalid.json at commit 006f8f7d49a39bc1e3a302269c8ad244f05a209b.
// Their SHA-256 hashes are asserted while loading so corpus changes require an
// explicit source-pin update. Private-key WIF rows are counted but excluded
// because [StringBTCAddress] and [StringBTCBech32Address] validate destinations.
func TestStringBTCAddressRules_BitcoinCoreKeyIO(t *testing.T) {
	const (
		legacyErrorMessage = "string must be a valid mainnet legacy Bitcoin Base58Check address"
		bech32ErrorMessage = "string must be a valid mainnet Bitcoin Bech32 address"
	)

	legacyRule := StringBTCAddress()
	bech32Rule := StringBTCBech32Address()
	valid := loadBitcoinCoreKeyIOValid(t)
	partitions := map[string]int{"valid rows": len(valid)}
	caseFlipped := make(map[string]struct{}, 32)

	for row, vector := range valid {
		if vector.Metadata.IsPrivkey {
			partitions["WIF excluded"]++
			if vector.Metadata.TryCaseFlip {
				t.Errorf("valid fixture row %d: WIF unexpectedly requests case flip", row+1)
			}
			continue
		}

		partitions["destination rows"]++
		kind, witnessVersion := classifyBitcoinCoreDestination(t, vector)
		wantLegacy := kind == "legacy" && vector.Metadata.Chain == "main"
		wantBech32 := kind == "segwit" && vector.Metadata.Chain == "main" && witnessVersion == 0
		assertStringFormatRuleResult(
			t,
			legacyRule,
			vector.Input,
			wantLegacy,
			legacyErrorMessage,
			ErrorCodeStringBTCAddress,
			"valid fixture legacy boundary",
			row+1,
		)
		assertStringFormatRuleResult(
			t,
			bech32Rule,
			vector.Input,
			wantBech32,
			bech32ErrorMessage,
			ErrorCodeStringBTCBech32Address,
			"valid fixture Bech32 boundary",
			row+1,
		)

		switch kind {
		case "legacy":
			partitions["legacy destinations"]++
			if vector.Metadata.Chain == "main" {
				partitions["legacy mainnet"]++
			} else {
				partitions["legacy non-mainnet"]++
			}
		case "segwit":
			partitions["SegWit destinations"]++
			network := "non-mainnet"
			if vector.Metadata.Chain == "main" {
				network = "mainnet"
			}
			version := "v1+"
			if witnessVersion == 0 {
				version = "v0"
			}
			partitions["SegWit "+network+" "+version]++
		}

		if !vector.Metadata.TryCaseFlip {
			continue
		}
		partitions["case flips"]++
		if strings.ToLower(vector.Input) != vector.Input {
			t.Errorf("valid fixture row %d: tryCaseFlip source is not lowercase: %q", row+1, vector.Input)
		}
		uppercase := strings.ToUpper(vector.Input)
		if _, exists := caseFlipped[uppercase]; exists {
			t.Errorf("valid fixture row %d: duplicate uppercase mutation %q", row+1, uppercase)
		}
		caseFlipped[uppercase] = struct{}{}
		assertStringFormatRuleResult(
			t,
			legacyRule,
			uppercase,
			wantLegacy,
			legacyErrorMessage,
			ErrorCodeStringBTCAddress,
			"tryCaseFlip legacy boundary",
			row+1,
		)
		assertStringFormatRuleResult(
			t,
			bech32Rule,
			uppercase,
			wantBech32,
			bech32ErrorMessage,
			ErrorCodeStringBTCBech32Address,
			"tryCaseFlip Bech32 boundary",
			row+1,
		)
	}

	expectedPartitions := map[string]int{
		"valid rows":             70,
		"WIF excluded":           16,
		"destination rows":       54,
		"legacy destinations":    22,
		"legacy mainnet":         6,
		"legacy non-mainnet":     16,
		"SegWit destinations":    32,
		"SegWit mainnet v0":      4,
		"SegWit mainnet v1+":     4,
		"SegWit non-mainnet v0":  12,
		"SegWit non-mainnet v1+": 12,
		"case flips":             32,
	}
	for name, expected := range expectedPartitions {
		if actual := partitions[name]; actual != expected {
			t.Errorf("%s count is %d; want %d", name, actual, expected)
		}
	}

	invalid := loadBitcoinCoreKeyIOInvalid(t)
	for row, input := range invalid {
		assertStringFormatRuleResult(
			t,
			legacyRule,
			input,
			false,
			legacyErrorMessage,
			ErrorCodeStringBTCAddress,
			"invalid fixture legacy boundary",
			row+1,
		)
		assertStringFormatRuleResult(
			t,
			bech32Rule,
			input,
			false,
			bech32ErrorMessage,
			ErrorCodeStringBTCBech32Address,
			"invalid fixture Bech32 boundary",
			row+1,
		)
	}
}

// The fixtures preserve every native SegWit vector from BIP-173 and BIP-350
// at bitcoin/bips commit 8c369ac8e60629ac6c032ffe21bb5ec5b35213d7.
// Each exact source input is classified against [StringBTCBech32Address]'s
// intentionally narrower mainnet, witness-v0, Bech32-only contract.
func TestStringBTCBech32Address_BIPSegWitVectors(t *testing.T) {
	const errorMessage = "string must be a valid mainnet Bitcoin Bech32 address"
	rule := StringBTCBech32Address()
	tests := map[string]struct {
		fixture            string
		fixtureSHA256      string
		sourceFile         string
		sourceFileSHA256   string
		expectedValid      int
		expectedInvalid    int
		expectedPartitions map[string]int
	}{
		"BIP-173": {
			fixture:          "bip173_segwit_vectors.json",
			fixtureSHA256:    "043e7295b39a985e519a720638712e3a6f1883cecb87936f60a136caf7a12c15",
			sourceFile:       "bip-0173.mediawiki",
			sourceFileSHA256: "b19eab06b264bbd15ad7f7d640fa34de8f4a8809fba8a6ecc3e165545b27a543",
			expectedValid:    6,
			expectedInvalid:  10,
			expectedPartitions: map[string]int{
				"source-valid":    6,
				"source-invalid":  10,
				"accepted":        1,
				"non-mainnet":     2,
				"v1+":             3,
				"non-mainnet v1+": 0,
			},
		},
		"BIP-350": {
			fixture:          "bip350_segwit_vectors.json",
			fixtureSHA256:    "1fcd84cc0388ab57c3371158bd59717ce66f853a9fb285497832cbd0fbb417f9",
			sourceFile:       "bip-0350.mediawiki",
			sourceFileSHA256: "63634b06aa8bae88b31929674736e74964c4598684269ac2b0b140c43a7a0dec",
			expectedValid:    8,
			expectedInvalid:  15,
			expectedPartitions: map[string]int{
				"source-valid":    8,
				"source-invalid":  15,
				"accepted":        1,
				"non-mainnet":     2,
				"v1+":             4,
				"non-mainnet v1+": 1,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			fixture := loadBIPSegWitFixture(t, tc.fixture, tc.fixtureSHA256)
			if fixture.Source.Repository != "https://github.com/bitcoin/bips" {
				t.Errorf("source repository is %q", fixture.Source.Repository)
			}
			if fixture.Source.Commit != "8c369ac8e60629ac6c032ffe21bb5ec5b35213d7" {
				t.Errorf("source commit is %q", fixture.Source.Commit)
			}
			if fixture.Source.File != tc.sourceFile {
				t.Errorf("source file is %q; want %q", fixture.Source.File, tc.sourceFile)
			}
			if fixture.Source.FileSHA256 != tc.sourceFileSHA256 {
				t.Errorf("source file SHA-256 is %q; want %q", fixture.Source.FileSHA256, tc.sourceFileSHA256)
			}
			if fixture.Source.ValidCount != tc.expectedValid || fixture.Source.InvalidCount != tc.expectedInvalid {
				t.Errorf(
					"source counts are %d valid/%d invalid; want %d/%d",
					fixture.Source.ValidCount,
					fixture.Source.InvalidCount,
					tc.expectedValid,
					tc.expectedInvalid,
				)
			}

			partitions := make(map[string]int, len(tc.expectedPartitions))
			for row, vector := range fixture.Vectors {
				contract := classifyBIPSegWitVector(t, tc.fixture, row+1, vector)
				if vector.SourceValid {
					partitions["source-valid"]++
					partitions[contract]++
				} else {
					partitions["source-invalid"]++
				}
				assertStringFormatRuleResult(
					t,
					rule,
					vector.Address,
					contract == "accepted",
					errorMessage,
					ErrorCodeStringBTCBech32Address,
					tc.fixture,
					row+1,
				)
			}
			for partition, expected := range tc.expectedPartitions {
				if actual := partitions[partition]; actual != expected {
					t.Errorf("%s count is %d; want %d", partition, actual, expected)
				}
			}
		})
	}
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

type bitcoinCoreKeyIOValidMetadata struct {
	Chain       string `json:"chain"`
	IsPrivkey   bool   `json:"isPrivkey"`
	TryCaseFlip bool   `json:"tryCaseFlip"`
}

type bitcoinCoreKeyIOValidVector struct {
	Input    string
	Script   string
	Metadata bitcoinCoreKeyIOValidMetadata
}

type bipSegWitFixtureSource struct {
	Repository   string `json:"repository"`
	Commit       string `json:"commit"`
	File         string `json:"file"`
	FileSHA256   string `json:"file_sha256"`
	ValidCount   int    `json:"valid_count"`
	InvalidCount int    `json:"invalid_count"`
}

type bipSegWitVector struct {
	Address        string `json:"address"`
	SourceValid    bool   `json:"source_valid"`
	ScriptPubKey   string `json:"script_pub_key"`
	Reason         string `json:"reason"`
	Network        string `json:"network"`
	WitnessVersion *int   `json:"witness_version"`
	Encoding       string `json:"encoding"`
	Contract       string `json:"contract"`
}

type bipSegWitFixture struct {
	Source  bipSegWitFixtureSource `json:"source"`
	Vectors []bipSegWitVector      `json:"vectors"`
}

func loadBitcoinCoreKeyIOValid(t *testing.T) []bitcoinCoreKeyIOValidVector {
	t.Helper()
	data := readPinnedCryptoFixture(
		t,
		"bitcoin_core_key_io_valid.json",
		"90bd1d35d12763e0d00c5400b2c9fe551e532a821e1e466c20cad3aced70a7fe",
	)

	var rows [][]json.RawMessage
	if err := json.Unmarshal(data, &rows); err != nil {
		t.Fatalf("decode Bitcoin Core valid fixture: %v", err)
	}
	if len(rows) != 70 {
		t.Fatalf("Bitcoin Core valid fixture contains %d rows; want 70", len(rows))
	}

	vectors := make([]bitcoinCoreKeyIOValidVector, 0, len(rows))
	seen := make(map[string]int, len(rows))
	for row, fields := range rows {
		if len(fields) != 3 {
			t.Fatalf("Bitcoin Core valid fixture row %d has %d fields; want 3", row+1, len(fields))
		}
		var vector bitcoinCoreKeyIOValidVector
		if err := json.Unmarshal(fields[0], &vector.Input); err != nil {
			t.Fatalf("decode Bitcoin Core valid fixture row %d input: %v", row+1, err)
		}
		if err := json.Unmarshal(fields[1], &vector.Script); err != nil {
			t.Fatalf("decode Bitcoin Core valid fixture row %d script: %v", row+1, err)
		}
		if err := json.Unmarshal(fields[2], &vector.Metadata); err != nil {
			t.Fatalf("decode Bitcoin Core valid fixture row %d metadata: %v", row+1, err)
		}
		if firstRow, exists := seen[vector.Input]; exists {
			t.Fatalf("Bitcoin Core valid fixture row %d duplicates row %d input %q", row+1, firstRow, vector.Input)
		}
		seen[vector.Input] = row + 1
		vectors = append(vectors, vector)
	}
	return vectors
}

func loadBitcoinCoreKeyIOInvalid(t *testing.T) []string {
	t.Helper()
	data := readPinnedCryptoFixture(
		t,
		"bitcoin_core_key_io_invalid.json",
		"c3ca74ddad7c01faaca7c26063537feb73ff2be03eb7fc83a42a655d2979261d",
	)

	var rows [][]json.RawMessage
	if err := json.Unmarshal(data, &rows); err != nil {
		t.Fatalf("decode Bitcoin Core invalid fixture: %v", err)
	}
	if len(rows) != 70 {
		t.Fatalf("Bitcoin Core invalid fixture contains %d rows; want 70", len(rows))
	}

	inputs := make([]string, 0, len(rows))
	seen := make(map[string]int, len(rows))
	for row, fields := range rows {
		if len(fields) != 1 {
			t.Fatalf("Bitcoin Core invalid fixture row %d has %d fields; want 1", row+1, len(fields))
		}
		var input string
		if err := json.Unmarshal(fields[0], &input); err != nil {
			t.Fatalf("decode Bitcoin Core invalid fixture row %d: %v", row+1, err)
		}
		if firstRow, exists := seen[input]; exists {
			t.Fatalf("Bitcoin Core invalid fixture row %d duplicates row %d input %q", row+1, firstRow, input)
		}
		seen[input] = row + 1
		inputs = append(inputs, input)
	}
	return inputs
}

func loadBIPSegWitFixture(t *testing.T, name, expectedSHA256 string) bipSegWitFixture {
	t.Helper()
	data := readPinnedCryptoFixture(t, name, expectedSHA256)
	var fixture bipSegWitFixture
	if err := json.Unmarshal(data, &fixture); err != nil {
		t.Fatalf("decode %s: %v", name, err)
	}
	expectedCount := fixture.Source.ValidCount + fixture.Source.InvalidCount
	if len(fixture.Vectors) != expectedCount {
		t.Fatalf("%s contains %d vectors; source metadata declares %d", name, len(fixture.Vectors), expectedCount)
	}
	seen := make(map[string]int, len(fixture.Vectors))
	for row, vector := range fixture.Vectors {
		if firstRow, exists := seen[vector.Address]; exists {
			t.Fatalf("%s row %d duplicates row %d input %q", name, row+1, firstRow, vector.Address)
		}
		seen[vector.Address] = row + 1
	}
	return fixture
}

func readPinnedCryptoFixture(t *testing.T, name, expectedSHA256 string) []byte {
	t.Helper()
	path := filepath.Join("testdata", name)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	sum := sha256.Sum256(data)
	if actual := hex.EncodeToString(sum[:]); actual != expectedSHA256 {
		t.Fatalf("%s SHA-256 is %s; want %s", path, actual, expectedSHA256)
	}
	return data
}

func classifyBitcoinCoreDestination(
	t *testing.T,
	vector bitcoinCoreKeyIOValidVector,
) (kind string, witnessVersion int) {
	t.Helper()
	script, err := hex.DecodeString(vector.Script)
	if err != nil {
		t.Fatalf("decode script for %q: %v", vector.Input, err)
	}
	if len(script) == 25 && script[0] == 0x76 && script[1] == 0xa9 && script[2] == 0x14 &&
		script[23] == 0x88 && script[24] == 0xac {
		return "legacy", -1
	}
	if len(script) == 23 && script[0] == 0xa9 && script[1] == 0x14 && script[22] == 0x87 {
		return "legacy", -1
	}
	if len(script) < 4 || len(script) != int(script[1])+2 {
		t.Fatalf("unexpected destination script for %q: %s", vector.Input, vector.Script)
	}

	switch opcode := script[0]; {
	case opcode == 0:
		return "segwit", 0
	case opcode >= 0x51 && opcode <= 0x60:
		return "segwit", int(opcode - 0x50)
	default:
		t.Fatalf("unexpected witness version opcode for %q: 0x%02x", vector.Input, opcode)
		return "", -1
	}
}

func classifyBIPSegWitVector(t *testing.T, fixture string, row int, vector bipSegWitVector) string {
	t.Helper()
	if !vector.SourceValid {
		if vector.Contract != "source-invalid" || vector.Reason == "" || vector.ScriptPubKey != "" ||
			vector.WitnessVersion != nil {
			t.Fatalf("%s row %d has inconsistent source-invalid metadata", fixture, row)
		}
		return "source-invalid"
	}
	if vector.ScriptPubKey == "" || vector.Reason != "" || vector.WitnessVersion == nil {
		t.Fatalf("%s row %d has incomplete source-valid metadata", fixture, row)
	}
	if _, err := hex.DecodeString(vector.ScriptPubKey); err != nil {
		t.Fatalf("%s row %d scriptPubKey is not hexadecimal: %v", fixture, row, err)
	}
	expectedEncoding := "bech32"
	if strings.HasPrefix(fixture, "bip350_") && *vector.WitnessVersion > 0 {
		expectedEncoding = "bech32m"
	}
	if vector.Encoding != expectedEncoding {
		t.Fatalf("%s row %d encoding is %q; want %q", fixture, row, vector.Encoding, expectedEncoding)
	}

	var expected string
	switch {
	case vector.Network == "mainnet" && *vector.WitnessVersion == 0 && vector.Encoding == "bech32":
		expected = "accepted"
	case vector.Network == "testnet" && *vector.WitnessVersion == 0 && vector.Encoding == "bech32":
		expected = "non-mainnet"
	case vector.Network == "mainnet" && *vector.WitnessVersion > 0:
		expected = "v1+"
	case vector.Network == "testnet" && *vector.WitnessVersion > 0:
		expected = "non-mainnet v1+"
	default:
		t.Fatalf("%s row %d has unsupported source-valid classification", fixture, row)
	}
	if vector.Contract != expected {
		t.Fatalf("%s row %d contract is %q; want %q", fixture, row, vector.Contract, expected)
	}
	return expected
}

func assertStringFormatRuleResult(
	t *testing.T,
	rule govy.Rule[string],
	input string,
	wantValid bool,
	errorMessage string,
	errorCode govy.ErrorCode,
	source string,
	row int,
) {
	t.Helper()
	err := rule.Validate(input)
	if wantValid {
		if err != nil {
			t.Errorf("%s row %d rejected %q: %v", source, row, input, err)
		}
		return
	}
	if err == nil {
		t.Errorf("%s row %d accepted %q", source, row, input)
		return
	}
	if err.Error() != errorMessage {
		t.Errorf("%s row %d error is %q; want %q", source, row, err.Error(), errorMessage)
	}
	if !govy.HasErrorCode(err, errorCode) {
		t.Errorf("%s row %d error does not contain code %q", source, row, errorCode)
	}
}
