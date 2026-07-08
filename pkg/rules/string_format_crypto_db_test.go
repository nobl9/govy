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

func TestStringMongoDBObjectID(t *testing.T) {
	for _, tc := range stringMongoDBObjectIDTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			assertStringFormatRule(t, StringMongoDBObjectID(), ErrorCodeStringMongoDBObjectID, tc)
		})
	}
}

func BenchmarkStringMongoDBObjectID(b *testing.B) {
	benchmarkStringFormatRule(b, StringMongoDBObjectID(), stringMongoDBObjectIDTestCases())
}

func TestStringMongoDBConnectionString(t *testing.T) {
	for _, tc := range stringMongoDBConnectionStringTestCases() {
		t.Run(tc.name, func(t *testing.T) {
			assertStringFormatRule(t, StringMongoDBConnectionString(), ErrorCodeStringMongoDBConnectionString, tc)
		})
	}
}

func BenchmarkStringMongoDBConnectionString(b *testing.B) {
	benchmarkStringFormatRule(
		b,
		StringMongoDBConnectionString(),
		stringMongoDBConnectionStringTestCases(),
	)
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

func stringMongoDBObjectIDTestCases() []stringFormatRuleTestCase {
	const err = "string must be a 24-character lowercase hexadecimal MongoDB ObjectID"
	return []stringFormatRuleTestCase{
		{
			name: "object id",
			in:   "507f1f77bcf86cd799439011",
		},
		{
			name: "zero value",
			in:   "000000000000000000000000",
		},
		{
			name:          "uppercase hex",
			in:            "507F1F77BCF86CD799439011",
			expectedError: err,
		},
		{
			name:          "too short",
			in:            "507f1f77bcf86cd79943901",
			expectedError: err,
		},
		{
			name:          "too long",
			in:            "507f1f77bcf86cd7994390110",
			expectedError: err,
		},
		{
			name:          "not hex",
			in:            "507f1f77bcf86cd79943901g",
			expectedError: err,
		},
		{
			name:          "empty",
			expectedError: err,
		},
	}
}

func stringMongoDBConnectionStringTestCases() []stringFormatRuleTestCase {
	const err = "string must be a valid MongoDB connection string"
	return []stringFormatRuleTestCase{
		{
			name: "single host",
			in:   "mongodb://localhost",
		},
		{
			name: "host with port",
			in:   "mongodb://localhost:27017",
		},
		{
			name: "replica set",
			in:   "mongodb://user:pass@db1.example.com:27017,db2.example.com:27018/admin?replicaSet=rs0&authSource=admin",
		},
		{
			name: "ipv4",
			in:   "mongodb://127.0.0.1:27017/test",
		},
		{
			name: "ipv6",
			in:   "mongodb://[2001:db8::1]:27017/test",
		},
		{
			name: "srv",
			in:   "mongodb+srv://cluster0.example.com/test?retryWrites=true&w=majority",
		},
		{
			name:          "empty",
			expectedError: err + ": scheme must be mongodb:// or mongodb+srv://",
		},
		{
			name:          "unsupported scheme",
			in:            "http://localhost",
			expectedError: err + ": scheme must be mongodb:// or mongodb+srv://",
		},
		{
			name:          "missing host",
			in:            "mongodb://",
			expectedError: err + ": host must not be empty",
		},
		{
			name:          "empty host in list",
			in:            "mongodb://db1.example.com,,db2.example.com",
			expectedError: err + ": host must not be empty",
		},
		{
			name:          "bad port",
			in:            "mongodb://localhost:abc",
			expectedError: err + ": failed to parse URL: parse \"mongodb://localhost:abc\": invalid port \":abc\" after host",
		},
		{
			name:          "port out of range",
			in:            "mongodb://localhost:65536",
			expectedError: err + ": port must be between 0 and 65535",
		},
		{
			name:          "srv with multiple hosts",
			in:            "mongodb+srv://db1.example.com,db2.example.com",
			expectedError: err + ": mongodb+srv connection string must contain exactly one host",
		},
		{
			name:          "srv with port",
			in:            "mongodb+srv://cluster0.example.com:27017",
			expectedError: err + ": mongodb+srv host must not include a port",
		},
		{
			name:          "unbracketed ipv6",
			in:            "mongodb://2001:db8::1",
			expectedError: err + ": IPv6 hosts must be enclosed in brackets",
		},
		{
			name:          "invalid host label",
			in:            "mongodb://-example.com",
			expectedError: err + ": host must be a valid IP address or DNS name",
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
