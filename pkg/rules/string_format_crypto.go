package rules

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"strings"

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

// StringBTCAddress ensures the property's value is a mainnet legacy Bitcoin
// Base58Check address.
// It validates the Base58Check checksum and accepts P2PKH and P2SH version bytes.
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
// Bech32 address as defined by BIP-173.
// It validates the Bech32 checksum and accepts native v0 SegWit addresses.
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
// Mixed-case addresses must satisfy the EIP-55 checksum.
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
	if !ok || len(decoded) != 25 {
		return false
	}
	if decoded[0] != btcP2PKHVersion && decoded[0] != btcP2SHVersion {
		return false
	}

	checksum := bitcoinChecksum(decoded[:21])
	return bytes.Equal(checksum, decoded[21:])
}

func isETHAddress(s string) bool {
	if !ethAddressRegexp().MatchString(s) {
		return false
	}

	payload := s[len("0x"):]
	if !hasMixedCase(payload) {
		return true
	}
	return hasValidETHChecksum(payload)
}

func hasValidETHChecksum(payload string) bool {
	hash := keccak256([]byte(strings.ToLower(payload)))
	for i := range len(payload) {
		b := payload[i]
		if b >= '0' && b <= '9' {
			continue
		}

		uppercase := ethChecksumNibble(hash, i) >= 8
		if uppercase != (b >= 'A' && b <= 'F') {
			return false
		}
	}
	return true
}

func keccak256(data []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	_, _ = h.Write(data)
	return h.Sum(nil)
}

func ethChecksumNibble(hash []byte, index int) byte {
	if index%2 == 0 {
		return hash[index/2] >> 4
	}
	return hash[index/2] & 0x0f
}

func decodeBitcoinBase58(s string) ([]byte, bool) {
	var decoded big.Int
	base := big.NewInt(58)
	for i := range len(s) {
		digit := strings.IndexByte(bitcoinBase58Alphabet, s[i])
		if digit < 0 {
			return nil, false
		}

		decoded.Mul(&decoded, base)
		decoded.Add(&decoded, big.NewInt(int64(digit)))
	}

	leadingZeroes := 0
	for leadingZeroes < len(s) && s[leadingZeroes] == '1' {
		leadingZeroes++
	}
	return append(make([]byte, leadingZeroes), decoded.Bytes()...), true
}

func bitcoinChecksum(payload []byte) []byte {
	first := sha256.Sum256(payload)
	second := sha256.Sum256(first[:])
	return second[:4]
}

func isBTCBech32Address(s string) bool {
	if !btcBech32AddressRegexp().MatchString(s) || hasMixedCase(s) {
		return false
	}

	s = strings.ToLower(s)
	separator := strings.LastIndexByte(s, '1')
	if separator != len("bc") {
		return false
	}

	data := make([]byte, 0, len(s)-separator-1)
	for i := separator + 1; i < len(s); i++ {
		value := strings.IndexByte(bech32Charset, s[i])
		if value < 0 {
			return false
		}
		data = append(data, byteFromInt(value))
	}
	if !bech32VerifyChecksum("bc", data) || len(data) <= 6 {
		return false
	}

	payload := data[:len(data)-6]
	if len(payload) == 0 || payload[0] != 0 {
		return false
	}
	program, ok := convertBits(payload[1:], 5, 8, false)
	return ok && (len(program) == 20 || len(program) == 32)
}

func hasMixedCase(s string) bool {
	return strings.ToLower(s) != s && strings.ToUpper(s) != s
}

func bech32VerifyChecksum(hrp string, data []byte) bool {
	values := make([]byte, 0, len(hrp)*2+1+len(data))
	for i := range len(hrp) {
		values = append(values, hrp[i]>>5)
	}
	values = append(values, 0)
	for i := range len(hrp) {
		values = append(values, hrp[i]&31)
	}
	values = append(values, data...)
	return bech32Polymod(values) == 1
}

func bech32Polymod(values []byte) uint32 {
	generators := [...]uint32{0x3b6a57b2, 0x26508e6d, 0x1ea119fa, 0x3d4233dd, 0x2a1462b3}
	chk := uint32(1)
	for _, value := range values {
		top := chk >> 25
		chk = (chk&0x1ffffff)<<5 ^ uint32(value)
		for i, generator := range generators {
			if (top>>i)&1 == 1 {
				chk ^= generator
			}
		}
	}
	return chk
}

func convertBits(data []byte, fromBits, toBits uint, pad bool) ([]byte, bool) {
	acc := 0
	bits := uint(0)
	maxValue := (1 << toBits) - 1
	maxAccumulator := (1 << (fromBits + toBits - 1)) - 1
	converted := make([]byte, 0, len(data)*int(fromBits)/int(toBits))

	for _, value := range data {
		if uint(value)>>fromBits != 0 {
			return nil, false
		}
		acc = ((acc << fromBits) | int(value)) & maxAccumulator
		bits += fromBits
		for bits >= toBits {
			bits -= toBits
			converted = append(converted, byteFromInt((acc>>bits)&maxValue))
		}
	}

	if pad {
		if bits > 0 {
			converted = append(converted, byteFromInt((acc<<(toBits-bits))&maxValue))
		}
		return converted, true
	}
	if bits >= fromBits || ((acc<<(toBits-bits))&maxValue) != 0 {
		return nil, false
	}
	return converted, true
}

func byteFromInt(n int) byte {
	if n < 0 || n > 255 {
		panic("integer out of byte range")
	}
	return byte(n)
}
