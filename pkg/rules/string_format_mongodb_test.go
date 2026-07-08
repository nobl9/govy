package rules

import (
	"net/url"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

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
			name: "unix socket",
			in:   "mongodb://" + url.PathEscape("/tmp/mongodb-27017.sock"),
		},
		{
			name: "unix socket with database",
			in:   "mongodb://" + url.PathEscape("/var/run/mongodb.sock") + "/admin",
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
			expectedError: err + ": port must contain only digits",
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
			name:          "srv with unix socket",
			in:            "mongodb+srv://" + url.PathEscape("/tmp/mongodb-27017.sock"),
			expectedError: err + ": mongodb+srv host must be a DNS name",
		},
		{
			name:          "unbracketed ipv6",
			in:            "mongodb://2001:db8::1",
			expectedError: err + ": IPv6 hosts must be enclosed in brackets",
		},
		{
			name:          "invalid host label",
			in:            "mongodb://-example.com",
			expectedError: err + ": host must be a valid IP address, DNS name, or URL-encoded Unix domain socket",
		},
		{
			name:          "unix socket missing socket suffix",
			in:            "mongodb://" + url.PathEscape("/tmp/mongodb-27017"),
			expectedError: err + ": host must be a valid IP address, DNS name, or URL-encoded Unix domain socket",
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
