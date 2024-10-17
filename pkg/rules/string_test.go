package rules

import (
	"regexp"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var stringNotEmptyTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"                s", false},
	{"     ", true},
}

func TestStringNotEmpty(t *testing.T) {
	for _, tc := range stringNotEmptyTestCases {
		err := StringNotEmpty().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringNotEmpty))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringNotEmpty(b *testing.B) {
	for range b.N {
		for _, tc := range stringNotEmptyTestCases {
			_ = StringNotEmpty().Validate(tc.in)
		}
	}
}

var (
	stringMatchRegexpRegexp    = regexp.MustCompile("[ab]+")
	stringMatchRegexpTestCases = []*struct {
		in            string
		examples      []string
		expectedError string
	}{
		{
			in: "ab",
		},
		{
			in:            "cd",
			expectedError: "string must match regular expression: '[ab]+'",
		},
		{
			in:            "cd",
			examples:      []string{"ab", "a", "b"},
			expectedError: "string must match regular expression: '[ab]+' (e.g. 'ab', 'a', 'b')",
		},
	}
)

func TestStringMatchRegexp(t *testing.T) {
	for _, tc := range stringMatchRegexpTestCases {
		err := StringMatchRegexp(stringMatchRegexpRegexp, tc.examples...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMatchRegexp))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringMatchRegexp(b *testing.B) {
	for range b.N {
		for _, tc := range stringMatchRegexpTestCases {
			_ = StringMatchRegexp(stringMatchRegexpRegexp, tc.examples...).Validate(tc.in)
		}
	}
}

var (
	stringDenyRegexpRegexp    = regexp.MustCompile("[ab]+")
	stringDenyRegexpTestCases = []*struct {
		in            string
		examples      []string
		expectedError string
	}{
		{
			in: "cd",
		},
		{
			in:            "ab",
			expectedError: "string must not match regular expression: '[ab]+'",
		},
		{
			in:            "ab",
			examples:      []string{"ab", "a", "b"},
			expectedError: "string must not match regular expression: '[ab]+' (e.g. 'ab', 'a', 'b')",
		},
	}
)

func TestStringDenyRegexp(t *testing.T) {
	for _, tc := range stringDenyRegexpTestCases {
		err := StringDenyRegexp(stringDenyRegexpRegexp, tc.examples...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDenyRegexp))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDenyRegexp(b *testing.B) {
	for range b.N {
		for _, tc := range stringDenyRegexpTestCases {
			_ = StringDenyRegexp(stringDenyRegexpRegexp, tc.examples...).Validate(tc.in)
		}
	}
}

var stringDNSLabelTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test", false},
	{"s", false},
	{"test-this", false},
	{"test-1-this", false},
	{"test1-this", false},
	{"123", false},
	{strings.Repeat("l", 63), false},
	{"", true},
	{strings.Repeat("l", 64), true},
	{"tesT", true},
	{"test?", true},
	{"test this", true},
	{"1_2", true},
	{"LOL", true},
	// cspell:enable
}

func TestStringDNSLabel(t *testing.T) {
	for _, tc := range stringDNSLabelTestCases {
		err := StringDNSLabel().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDNSLabel))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDNSLabel(b *testing.B) {
	for range b.N {
		for _, tc := range stringDNSLabelTestCases {
			_ = StringDNSLabel().Validate(tc.in)
		}
	}
}

var stringASCIITestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"", false},
	{"foobar", false},
	{"0987654321", false},
	{"test@example.com", false},
	{"1234abcDEF", false},
	{"newline\n", false},
	{"\x19test\x7F", false},
	{"ｆｏｏbar", true},
	{"ｘｙｚ０９８", true},
	{"１２３456", true},
	{"ｶﾀｶﾅ", true},
	// cspell:enable
}

func TestStringASCII(t *testing.T) {
	for _, tc := range stringASCIITestCases {
		err := StringASCII().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringASCII))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringASCII(b *testing.B) {
	for range b.N {
		for _, tc := range stringASCIITestCases {
			_ = StringASCII().Validate(tc.in)
		}
	}
}

var stringUUIDTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"00000000-0000-0000-0000-000000000000", false},
	{"e190c630-8873-11ee-b9d1-0242ac120002", false},
	{"79258D24-01A7-47E5-ACBB-7E762DE52298", false},
	{"a987Fbc9-4bed-3078-cf07-9141ba07c9f3", false},
	{"foobar", true},
	{"0987654321", true},
	{"AXAXAXAX-AAAA-AAAA-AAAA-AAAAAAAAAAAA", true},
	{"00000000-0000-0000-0000-0000000000", true},
	{"", true},
	{"xxxa987Fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	{"a987Fbc9-4bed-3078-cf07-9141ba07c9f3xxx", true},
	{"a987Fbc94bed3078cf079141ba07c9f3", true},
	{"934859", true},
	{"987fbc9-4bed-3078-cf07a-9141ba07c9F3", true},
	{"aaaaaaaa-1111-1111-aaaG-111111111111", true},
	// cspell:enable
}

func TestStringUUID(t *testing.T) {
	for _, tc := range stringUUIDTestCases {
		err := StringUUID().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringUUID))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringUUID(b *testing.B) {
	for range b.N {
		for _, tc := range stringUUIDTestCases {
			_ = StringUUID().Validate(tc.in)
		}
	}
}

var stringEmailTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"test@mail.com", false},
	{"Dörte@Sörensen.example.com", false},
	{"θσερ@εχαμπλε.ψομ", false},
	{"юзер@екзампл.ком", false},
	{"उपयोगकर्ता@उदाहरण.कॉम", false},
	{"用户@例子.广告", false},
	{`"test test"@email.com`, false},
	{"mail@domain_with_underscores.org", false},
	{"test@email", false},
	{"test@t", false},
	{"", true},
	{"test@", true},
	{"test", true},
	{"test@email.", true},
	{"@email.com", true},
	{`"@email.com`, true},
	// cspell:enable
}

func TestStringEmail(t *testing.T) {
	for _, tc := range stringEmailTestCases {
		err := StringEmail().Validate(tc.in)
		if tc.shouldFail {
			assert.ErrorContains(t, err, "string must be a valid email address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEmail))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringEmail(b *testing.B) {
	for range b.N {
		for _, tc := range stringEmailTestCases {
			_ = StringEmail().Validate(tc.in)
		}
	}
}

func TestStringURL(t *testing.T) {
	for _, tc := range urlTestCases {
		err := StringURL().Validate(tc.url)
		if tc.shouldFail {
			assert.Require(t, assert.Error(t, err))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringURL))
		} else {
			assert.NoError(t, err)
		}
	}
	t.Run("failed to parse url", func(t *testing.T) {
		err := StringURL().Validate("http://\x1f")
		assert.ErrorContains(
			t,
			err,
			"failed to parse URL: parse \"http://\\x1f\": net/url: invalid control character in URL",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringURL))
	})
}

func BenchmarkStringURL(b *testing.B) {
	for range b.N {
		for _, tc := range urlTestCases {
			_ = StringURL().Validate(tc.url)
		}
	}
}

var stringMACTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"3D:F2:C9:A6:B3:4F", false},
	{"00:25:96:FF:FE:12:34:56", false},
	{"3D-F2-C9-A6-B3:4F", true},
	{"123", true},
	{"", true},
	{"abacaba", true},
	{"0025:96FF:FE12:3456", true},
	// cspell:enable
}

func TestStringMAC(t *testing.T) {
	for _, tc := range stringMACTestCases {
		err := StringMAC().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid MAC address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMAC))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringMAC(b *testing.B) {
	for range b.N {
		for _, tc := range stringMACTestCases {
			_ = StringMAC().Validate(tc.in)
		}
	}
}

var stringIPTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"10.0.0.1", false},
	{"172.16.0.1", false},
	{"192.168.0.1", false},
	{"192.168.255.254", false},
	{"172.16.255.254", false},
	{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
	{"2001:cdba:0:0:0:0:3257:9652", false},
	{"2001:cdba::3257:9652", false},
	{"", true},
	{"172.16.256.255", true},
	{"192.168.255.256", true},
	// cspell:enable
}

func TestStringIP(t *testing.T) {
	for _, tc := range stringIPTestCases {
		err := StringIP().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IP address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIP))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringIP(b *testing.B) {
	for range b.N {
		for _, tc := range stringIPTestCases {
			_ = StringIP().Validate(tc.in)
		}
	}
}

var stringIPv4TestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"10.0.0.1", false},
	{"172.16.0.1", false},
	{"192.168.0.1", false},
	{"192.168.255.254", false},
	{"172.16.255.254", false},
	{"192.168.255.256", true},
	{"172.16.256.255", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652", true},
	{"2001:cdba:0:0:0:0:3257:9652", true},
	{"2001:cdba::3257:9652", true},
	// cspell:enable
}

func TestStringIPv4(t *testing.T) {
	for _, tc := range stringIPv4TestCases {
		err := StringIPv4().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IPv4 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIPv4))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringIPv4(b *testing.B) {
	for range b.N {
		for _, tc := range stringIPv4TestCases {
			_ = StringIPv4().Validate(tc.in)
		}
	}
}

var stringIPv6TestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
	{"2001:cdba:0:0:0:0:3257:9652", false},
	{"2001:cdba::3257:9652", false},
	{"10.0.0.1", true},
	{"172.16.0.1", true},
	{"192.168.0.1", true},
	{"192.168.255.254", true},
	{"192.168.255.256", true},
	{"172.16.255.254", true},
	{"172.16.256.255", true},
	// cspell:enable
}

func TestStringIPv6(t *testing.T) {
	for _, tc := range stringIPv6TestCases {
		err := StringIPv6().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IPv6 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIPv6))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringIPv6(b *testing.B) {
	for range b.N {
		for _, tc := range stringIPv6TestCases {
			_ = StringIPv6().Validate(tc.in)
		}
	}
}

var stringCIDRTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"10.0.0.0/0", false},
	{"10.0.0.1/8", false},
	{"172.16.0.1/16", false},
	{"192.168.0.1/24", false},
	{"192.168.255.254/24", false},
	{"172.16.255.254/16", false},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/64", false},
	{"2001:cdba:0:0:0:0:3257:9652/32", false},
	{"2001:cdba::3257:9652/16", false},
	{"192.168.255.254/48", true},
	{"192.168.255.256/24", true},
	{"172.16.256.255/16", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/256", true},
	// cspell:enable
}

func TestStringCIDR(t *testing.T) {
	for _, tc := range stringCIDRTestCases {
		err := StringCIDR().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IP address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDR))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringCIDR(b *testing.B) {
	for range b.N {
		for _, tc := range stringCIDRTestCases {
			_ = StringCIDR().Validate(tc.in)
		}
	}
}

var stringCIDRv4TestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"0.0.0.0/0", false},
	{"10.0.0.0/8", false},
	{"172.16.0.0/16", false},
	{"192.168.0.0/24", false},
	{"172.16.0.0/16", false},
	{"192.168.255.0/24", false},
	{"10.0.0.0/0", true},
	{"10.0.0.1/8", true},
	{"172.16.0.1/16", true},
	{"192.168.0.1/24", true},
	{"192.168.255.254/24", true},
	{"192.168.255.254/48", true},
	{"192.168.255.256/24", true},
	{"172.16.255.254/16", true},
	{"172.16.256.255/16", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/64", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/256", true},
	{"2001:cdba:0:0:0:0:3257:9652/32", true},
	{"2001:cdba::3257:9652/16", true},
	{"172.56.1.0/16", true},
	// cspell:enable
}

func TestStringCIDRv4(t *testing.T) {
	for _, tc := range stringCIDRv4TestCases {
		err := StringCIDRv4().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IPv4 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDRv4))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringCIDRv4(b *testing.B) {
	for range b.N {
		for _, tc := range stringCIDRv4TestCases {
			_ = StringCIDRv4().Validate(tc.in)
		}
	}
}

var stringCIDRv6TestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"2001:cdba:0000:0000:0000:0000:3257:9652/64", false},
	{"2001:cdba:0:0:0:0:3257:9652/32", false},
	{"2001:cdba::3257:9652/16", false},
	{"10.0.0.0/0", true},
	{"10.0.0.1/8", true},
	{"172.16.0.1/16", true},
	{"192.168.0.1/24", true},
	{"192.168.255.254/24", true},
	{"192.168.255.254/48", true},
	{"192.168.255.256/24", true},
	{"172.16.255.254/16", true},
	{"172.16.256.255/16", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652/256", true},
	// cspell:enable
}

func TestStringCIDRv6(t *testing.T) {
	for _, tc := range stringCIDRv6TestCases {
		err := StringCIDRv6().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IPv6 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDRv6))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringCIDRv6(b *testing.B) {
	for range b.N {
		for _, tc := range stringCIDRv6TestCases {
			_ = StringCIDRv6().Validate(tc.in)
		}
	}
}

var stringJSONTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{`{"foo": "bar"}`, false},
	{`{}`, false},
	{`[]`, false},
	{"{]}", true},
	{"", true},
	{"yaml: ok", true},
}

func TestStringJSON(t *testing.T) {
	for _, tc := range stringJSONTestCases {
		err := StringJSON().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringJSON))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringJSON(b *testing.B) {
	for range b.N {
		for _, tc := range stringJSONTestCases {
			_ = StringJSON().Validate(tc.in)
		}
	}
}

var stringContainsTestCases = []*struct {
	in            string
	substrings    []string
	expectedError string
}{
	{
		in:         "",
		substrings: []string{""},
	},
	{
		in:         "this",
		substrings: []string{"his"},
	},
	{
		in:         "this",
		substrings: []string{"this"},
	},
	{
		in:         "this",
		substrings: []string{"th", "is"},
	},
	{
		in:            "one",
		substrings:    []string{"th"},
		expectedError: "string must contain the following substrings: 'th'",
	},
	{
		in:            "this",
		substrings:    []string{"th", "ht"},
		expectedError: "string must contain the following substrings: 'th', 'ht'",
	},
	{
		in:            "tha",
		substrings:    []string{"that"},
		expectedError: "string must contain the following substrings: 'that'",
	},
}

func TestStringContains(t *testing.T) {
	for _, tc := range stringContainsTestCases {
		err := StringContains(tc.substrings...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringContains))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringContains(b *testing.B) {
	for range b.N {
		for _, tc := range stringContainsTestCases {
			_ = StringContains(tc.substrings...).Validate(tc.in)
		}
	}
}

var stringExcludesTestCases = []*struct {
	in            string
	substrings    []string
	expectedError string
}{
	{
		in:         "one",
		substrings: []string{"th"},
	},
	{
		in:         "this",
		substrings: []string{"tho", "ht"},
	},
	{
		in:         "tha",
		substrings: []string{"that"},
	},
	{
		in:            "",
		substrings:    []string{""},
		expectedError: "string must not contain any of the following substrings: ''",
	},
	{
		in:            "this",
		substrings:    []string{"his"},
		expectedError: "string must not contain any of the following substrings: 'his'",
	},
	{
		in:            "this",
		substrings:    []string{"this"},
		expectedError: "string must not contain any of the following substrings: 'this'",
	},
	{
		in:            "this",
		substrings:    []string{"th", "is"},
		expectedError: "string must not contain any of the following substrings: 'th', 'is'",
	},
}

func TestStringExcludes(t *testing.T) {
	for _, tc := range stringExcludesTestCases {
		err := StringExcludes(tc.substrings...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringExcludes))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringExcludes(b *testing.B) {
	for range b.N {
		for _, tc := range stringExcludesTestCases {
			_ = StringExcludes(tc.substrings...).Validate(tc.in)
		}
	}
}

var stringStartsWithTestCases = []*struct {
	in            string
	prefixes      []string
	expectedError string
}{
	{
		in:       "this",
		prefixes: []string{"th"},
	},
	{
		in:       "this",
		prefixes: []string{"is", "th"},
	},
	{
		in:            "one",
		prefixes:      []string{"th"},
		expectedError: "string must start with 'th' prefix",
	},
	{
		in:            "one",
		prefixes:      []string{"th", "ht"},
		expectedError: "string must start with one of the following prefixes: 'th', 'ht'",
	},
}

func TestStringStartsWith(t *testing.T) {
	for _, tc := range stringStartsWithTestCases {
		err := StringStartsWith(tc.prefixes...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringStartsWith))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringStartsWith(b *testing.B) {
	for range b.N {
		for _, tc := range stringStartsWithTestCases {
			_ = StringStartsWith(tc.prefixes...).Validate(tc.in)
		}
	}
}

var stringEndsWithTestCases = []*struct {
	in            string
	suffixes      []string
	expectedError string
}{
	{
		in:       "this",
		suffixes: []string{"is"},
	},
	{
		in:       "this",
		suffixes: []string{"th", "is"},
	},
	{
		in:            "one",
		suffixes:      []string{"th"},
		expectedError: "string must end with 'th' suffix",
	},
	{
		in:            "one",
		suffixes:      []string{"th", "ht"},
		expectedError: "string must end with one of the following suffixes: 'th', 'ht'",
	},
}

func TestStringEndsWith(t *testing.T) {
	for _, tc := range stringEndsWithTestCases {
		err := StringEndsWith(tc.suffixes...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEndsWith))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringEndsWith(b *testing.B) {
	for range b.N {
		for _, tc := range stringEndsWithTestCases {
			_ = StringEndsWith(tc.suffixes...).Validate(tc.in)
		}
	}
}

var stringTitleTestCases = []*struct {
	in         string
	shouldFail bool
}{
	// cspell:disable
	{"", true},
	{"a", true},
	{"A", false},
	{" aaa aaa aaa ", true},
	{" Aaa Aaa Aaa ", false},
	{"123a456", true},
	{"double-blind", true},
	{"Double-Blind", false},
	{"ÿøû", true},
	{"Ÿøû", false},
	{"with_underscore", true},
	{"With_underscore", false},
	{"unicode \xe2\x80\xa8 line separator", true},
	{"Unicode \xe2\x80\xa8 Line Separator", false},
	// cspell:enable
}

func TestStringTitle(t *testing.T) {
	for _, tc := range stringTitleTestCases {
		err := StringTitle().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "each word in a string must start with a capital letter")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringTitle))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringTitle(b *testing.B) {
	for range b.N {
		for _, tc := range stringTitleTestCases {
			_ = StringTitle().Validate(tc.in)
		}
	}
}

var stringGitRefTestCases = []*struct {
	in          string
	expectedErr error
}{
	{"refs/heads/master", nil},
	{"refs/notes/commits", nil},
	{"refs/tags/this@", nil},
	{"refs/remotes/origin/master", nil},
	{"HEAD", nil},
	{"refs/tags/v3.1.1", nil},
	{"refs/pulls/1/head", nil},
	{"refs/pulls/1/merge", nil},
	{"refs/pulls/1/abc.123", nil},
	{"refs/pulls", nil},
	{"refs/-", nil},
	{"refs", errGitRefAtLeastOneSlash},
	{"refs/", errGitRefEmptyPart},
	{"refs//", errGitRefEmptyPart},
	{"refs/heads/\\", errGitRefForbiddenChars},
	{"refs/heads/\\foo", errGitRefForbiddenChars},
	{"refs/heads/\\foo/bar", errGitRefForbiddenChars},
	{"abc", errGitRefAtLeastOneSlash},
	{"", errGitRefEmpty},
	{"refs/heads/ ", errGitRefForbiddenChars},
	{"refs/heads/ /", errGitRefForbiddenChars},
	{"refs/heads/ /foo", errGitRefForbiddenChars},
	{"refs/heads/.", errGitRefEndsWithDot},
	{"refs/heads/..", errGitRefEndsWithDot},
	{"refs/heads/foo..", errGitRefEndsWithDot},
	{"refs/heads/foo.lock", errGitRefForbiddenChars},
	{"refs/heads/foo@{bar}", errGitRefForbiddenChars},
	{"refs/heads/foo@{", errGitRefForbiddenChars},
	{"refs/heads/foo[", errGitRefForbiddenChars},
	{"refs/heads/foo~", errGitRefForbiddenChars},
	{"refs/heads/foo^", errGitRefForbiddenChars},
	{"refs/heads/foo:", errGitRefForbiddenChars},
	{"refs/heads/foo?", errGitRefForbiddenChars},
	{"refs/heads/foo*", errGitRefForbiddenChars},
	{"refs/heads/foo[bar", errGitRefForbiddenChars},
	{"refs/heads/foo\t", errGitRefForbiddenChars},
	{"refs/heads/@", errGitRefForbiddenChars},
	{"refs/heads/@{bar}", errGitRefForbiddenChars},
	{"refs/heads/\n", errGitRefForbiddenChars},
	{"refs/heads/-foo", errGitRefStartsWithDash},
	{"refs/heads/foo..bar", errGitRefForbiddenChars},
	{"refs/heads/-", errGitRefStartsWithDash},
	{"refs/tags/-", errGitRefStartsWithDash},
	{"refs/tags/-foo", errGitRefStartsWithDash},
}

func TestStringGitRef(t *testing.T) {
	for _, tc := range stringGitRefTestCases {
		t.Run(tc.in, func(t *testing.T) {
			err := StringGitRef().Validate(tc.in)
			if tc.expectedErr != nil {
				assert.ErrorContains(t, err, tc.expectedErr.Error())
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringGitRef))
				assert.Equal(
					t,
					"see https://git-scm.com/docs/git-check-ref-format for more information on Git reference naming rules",
					err.(*govy.RuleError).Details,
				)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringGitRef(b *testing.B) {
	for range b.N {
		for _, tc := range stringGitRefTestCases {
			_ = StringGitRef().Validate(tc.in)
		}
	}
}
