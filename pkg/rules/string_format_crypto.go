package rules

import (
	"crypto/sha256"

	"golang.org/x/crypto/sha3"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

const (
	bitcoinBase58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	bech32Charset         = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

	btcP2PKHVersion byte = 0x00
	btcP2SHVersion  byte = 0x05
)

var (
	bitcoinBase58Values = newASCIIValueTable(bitcoinBase58Alphabet)
	bech32CharsetValues = newASCIIValueTable(bech32Charset)
)

// StringBTCAddress ensures the property's value is a mainnet legacy Bitcoin
// Base58Check address.
// It validates the Base58Check checksum and accepts pay-to-public-key-hash
// (P2PKH) and pay-to-script-hash (P2SH) version bytes.
func StringBTCAddress() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBTCAddressTemplate)

	return govy.NewRule(func(s string) error {
		if !isBTCAddress(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBTCAddress).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringBTCBech32Address ensures the property's value is a mainnet Bitcoin
// Bech32 address as defined by Bitcoin Improvement Proposal 173 (BIP-173).
// It validates the Bech32 checksum and accepts native version 0 Segregated
// Witness (SegWit) addresses.
// It does not accept Bech32m addresses such as Taproot addresses.
func StringBTCBech32Address() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBTCBech32AddressTemplate)

	return govy.NewRule(func(s string) error {
		if !isBTCBech32Address(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBTCBech32Address).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringETHAddress ensures the property's value is an Ethereum address with a
// 0x prefix followed by 40 hexadecimal characters.
// Mixed-case addresses must satisfy the Ethereum Improvement Proposal 55
// (EIP-55) checksum.
// All-lowercase and all-uppercase payloads are accepted as unchecksummed
// addresses.
func StringETHAddress() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringETHAddressTemplate)

	return govy.NewRule(func(s string) error {
		if !isETHAddress(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringETHAddress).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func isBTCAddress(s string) bool {
	if !btcAddressRegexp().MatchString(s) {
		return false
	}

	decoded, ok := decodeBitcoinBase58(s)
	if !ok {
		return false
	}
	if decoded[0] != btcP2PKHVersion && decoded[0] != btcP2SHVersion {
		return false
	}

	checksum := bitcoinChecksum(decoded[:21])
	return checksum[0] == decoded[21] &&
		checksum[1] == decoded[22] &&
		checksum[2] == decoded[23] &&
		checksum[3] == decoded[24]
}

func isETHAddress(s string) bool {
	if len(s) != 42 || s[0] != '0' || s[1] != 'x' {
		return false
	}

	payload := s[len("0x"):]
	var lowercase [40]byte
	var hasLower, hasUpper bool
	for i := range len(payload) {
		switch b := payload[i]; {
		case b >= '0' && b <= '9':
			lowercase[i] = b
		case b >= 'a' && b <= 'f':
			lowercase[i] = b
			hasLower = true
		case b >= 'A' && b <= 'F':
			lowercase[i] = b + ('a' - 'A')
			hasUpper = true
		default:
			return false
		}
	}
	if !hasLower || !hasUpper {
		return true
	}
	return hasValidETHChecksum(payload, lowercase)
}

func hasValidETHChecksum(payload string, lowercase [40]byte) bool {
	h := sha3.NewLegacyKeccak256()
	_, _ = h.Write(lowercase[:])
	var hash [32]byte
	digest := h.Sum(hash[:0])

	for i := range len(payload) {
		b := payload[i]
		if b >= '0' && b <= '9' {
			continue
		}

		uppercase := ethChecksumNibble(digest, i) >= 8
		if uppercase != (b >= 'A' && b <= 'F') {
			return false
		}
	}
	return true
}

func ethChecksumNibble(hash []byte, index int) byte {
	if index%2 == 0 {
		return hash[index/2] >> 4
	}
	return hash[index/2] & 0x0f
}

func decodeBitcoinBase58(s string) ([25]byte, bool) {
	var decoded [25]byte
	for i := range len(s) {
		if s[i] >= byte(len(bitcoinBase58Values)) {
			return decoded, false
		}

		value := bitcoinBase58Values[s[i]]
		if value == 0 {
			return decoded, false
		}

		carry := uint32(value - 1)
		for j := len(decoded) - 1; j >= 0; j-- {
			carry += uint32(decoded[j]) * 58
			decoded[j] = byte(carry & 0xff) //nolint:gosec // The mask bounds the conversion.
			carry >>= 8
		}
		if carry != 0 {
			return decoded, false
		}
	}

	leadingZeroes := 0
	for leadingZeroes < len(s) && s[leadingZeroes] == '1' {
		leadingZeroes++
	}
	firstNonZero := 0
	for firstNonZero < len(decoded) && decoded[firstNonZero] == 0 {
		firstNonZero++
	}
	if leadingZeroes+len(decoded)-firstNonZero != len(decoded) {
		return decoded, false
	}
	return decoded, true
}

func bitcoinChecksum(payload []byte) [4]byte {
	first := sha256.Sum256(payload)
	second := sha256.Sum256(first[:])
	return [4]byte{second[0], second[1], second[2], second[3]}
}

func isBTCBech32Address(s string) bool {
	if len(s) != 42 && len(s) != 62 {
		return false
	}
	if s[0]|('a'-'A') != 'b' || s[1]|('a'-'A') != 'c' || s[2] != '1' {
		return false
	}
	if s[3]|('a'-'A') != 'q' {
		return false
	}

	hasLower := s[0] == 'b' || s[1] == 'c'
	hasUpper := s[0] == 'B' || s[1] == 'C'
	if hasLower && hasUpper {
		return false
	}

	checksum := uint32(1)
	checksum = bech32PolymodStep(checksum, 'b'>>5)
	checksum = bech32PolymodStep(checksum, 'c'>>5)
	checksum = bech32PolymodStep(checksum, 0)
	checksum = bech32PolymodStep(checksum, 'b'&31)
	checksum = bech32PolymodStep(checksum, 'c'&31)

	for i := len("bc1"); i < len(s); i++ {
		b := s[i]
		switch {
		case b >= 'a' && b <= 'z':
			hasLower = true
		case b >= 'A' && b <= 'Z':
			hasUpper = true
			b += 'a' - 'A'
		}
		if hasLower && hasUpper {
			return false
		}
		if b >= byte(len(bech32CharsetValues)) {
			return false
		}

		encoded := bech32CharsetValues[b]
		if encoded == 0 {
			return false
		}
		value := encoded - 1
		if len(s) == 62 && i == len(s)-7 && value&0x0f != 0 {
			return false
		}
		checksum = bech32PolymodStep(checksum, value)
	}
	return checksum == 1
}

func bech32PolymodStep(checksum uint32, value byte) uint32 {
	top := checksum >> 25
	checksum = (checksum&0x1ffffff)<<5 ^ uint32(value)
	if top&1 != 0 {
		checksum ^= 0x3b6a57b2
	}
	if top&2 != 0 {
		checksum ^= 0x26508e6d
	}
	if top&4 != 0 {
		checksum ^= 0x1ea119fa
	}
	if top&8 != 0 {
		checksum ^= 0x3d4233dd
	}
	if top&16 != 0 {
		checksum ^= 0x2a1462b3
	}
	return checksum
}

func newASCIIValueTable(alphabet string) [128]byte {
	var values [128]byte
	for i := range len(alphabet) {
		values[alphabet[i]] = byteFromInt(i + 1)
	}
	return values
}

func byteFromInt(n int) byte {
	if n < 0 || n > 255 {
		panic("integer out of byte range")
	}
	return byte(n)
}
