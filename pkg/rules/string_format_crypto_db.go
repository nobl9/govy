package rules

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"math/big"
	"net"
	"net/url"
	"strconv"
	"strings"

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

// StringETHAddress ensures the property's value is an Ethereum address in
// hexadecimal form, with a 0x prefix followed by 40 hexadecimal characters.
// It does not verify EIP-55 mixed-case checksum casing.
func StringETHAddress() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringETHAddressTemplate)

	return govy.NewRule(func(s string) error {
		if !ethAddressRegexp().MatchString(s) {
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

// StringMongoDBObjectID ensures the property's value is a 24-character
// lowercase hexadecimal MongoDB ObjectID.
func StringMongoDBObjectID() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMongoDBObjectIDTemplate)

	return govy.NewRule(func(s string) error {
		if !mongoDBObjectIDRegexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMongoDBObjectID).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringMongoDBConnectionString ensures the property's value is a MongoDB
// connection string with a mongodb:// or mongodb+srv:// scheme and valid,
// non-empty host entries.
func StringMongoDBConnectionString() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringMongoDBConnectionStringTemplate)

	return govy.NewRule(func(s string) error {
		if err := validateMongoDBConnectionString(s); err != nil {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
				Error:         err.Error(),
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringMongoDBConnectionString).
		WithMessageTemplate(tpl).
		WithDescription("MongoDB connection string must use mongodb:// or mongodb+srv:// and contain valid, non-empty hosts")
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

func validateMongoDBConnectionString(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return errors.New("failed to parse URL: " + err.Error())
	}

	switch u.Scheme {
	case "mongodb", "mongodb+srv":
	default:
		return errors.New("scheme must be mongodb:// or mongodb+srv://")
	}
	if u.Host == "" {
		return errors.New("host must not be empty")
	}

	hosts := strings.Split(u.Host, ",")
	if u.Scheme == "mongodb+srv" && len(hosts) != 1 {
		return errors.New("mongodb+srv connection string must contain exactly one host")
	}
	for _, host := range hosts {
		if err = validateMongoDBHost(host, u.Scheme == "mongodb+srv"); err != nil {
			return err
		}
	}
	return nil
}

func validateMongoDBHost(host string, srv bool) error {
	if host == "" {
		return errors.New("host must not be empty")
	}
	if strings.ContainsAny(host, " \t\r\n") {
		return errors.New("host must not contain whitespace")
	}
	if strings.HasPrefix(host, "[") {
		hostname, port, err := splitMongoDBBracketedHost(host)
		if err != nil {
			return err
		}
		if srv && port != "" {
			return errors.New("mongodb+srv host must not include a port")
		}
		if net.ParseIP(hostname) == nil || !strings.Contains(hostname, ":") {
			return errors.New("host must be a valid IP address or DNS name")
		}
		return nil
	}
	if strings.ContainsAny(host, "[]") {
		return errors.New("host contains malformed IPv6 brackets")
	}
	if strings.Count(host, ":") > 1 {
		return errors.New("IPv6 hosts must be enclosed in brackets")
	}

	hostname, port, hasPort := strings.Cut(host, ":")
	if hostname == "" {
		return errors.New("host must not be empty")
	}
	if hasPort {
		if srv {
			return errors.New("mongodb+srv host must not include a port")
		}
		if err := validateMongoDBPort(port); err != nil {
			return err
		}
	}
	if !isValidMongoDBHostname(hostname) {
		return errors.New("host must be a valid IP address or DNS name")
	}
	return nil
}

func splitMongoDBBracketedHost(host string) (hostname, port string, err error) {
	end := strings.IndexByte(host, ']')
	if end < 0 {
		return "", "", errors.New("host contains malformed IPv6 brackets")
	}

	hostname = host[1:end]
	if hostname == "" {
		return "", "", errors.New("host must not be empty")
	}
	rest := host[end+1:]
	if rest == "" {
		return hostname, "", nil
	}
	if !strings.HasPrefix(rest, ":") {
		return "", "", errors.New("host contains malformed IPv6 brackets")
	}

	port = rest[1:]
	if err := validateMongoDBPort(port); err != nil {
		return "", "", err
	}
	return hostname, port, nil
}

func validateMongoDBPort(port string) error {
	if port == "" {
		return errors.New("port must not be empty")
	}
	for _, r := range port {
		if r < '0' || r > '9' {
			return errors.New("port must contain only digits")
		}
	}
	n, err := strconv.Atoi(port)
	if err != nil || n > 65535 {
		return errors.New("port must be between 0 and 65535")
	}
	return nil
}

func isValidMongoDBHostname(hostname string) bool {
	if hostname == "" || len(hostname) > 253 {
		return false
	}
	if ip := net.ParseIP(hostname); ip != nil {
		return true
	}

	hostname = strings.TrimSuffix(strings.ToLower(hostname), ".")
	if hostname == "" {
		return false
	}
	for _, label := range strings.Split(hostname, ".") {
		if !mongoDBHostLabelRegexp().MatchString(label) {
			return false
		}
	}
	return true
}
