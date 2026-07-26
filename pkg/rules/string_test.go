package rules

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"testing"
	"time"

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
	for _, tc := range stringNotEmptyTestCases {
		rule := StringNotEmpty()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var (
	stringMatchRegexpRegexp    = regexp.MustCompile("[ab]+")
	stringMatchRegexpTestCases = []*struct {
		in            string
		expectedError string
	}{
		{
			in: "ab",
		},
		{
			in:            "cd",
			expectedError: "string must match regular expression: '[ab]+'",
		},
	}
)

func TestStringMatchRegexp(t *testing.T) {
	for _, tc := range stringMatchRegexpTestCases {
		err := StringMatchRegexp(stringMatchRegexpRegexp).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMatchRegexp))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringMatchRegexp(b *testing.B) {
	for _, tc := range stringMatchRegexpTestCases {
		rule := StringMatchRegexp(stringMatchRegexpRegexp)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var (
	stringDenyRegexpRegexp    = regexp.MustCompile("[ab]+")
	stringDenyRegexpTestCases = []*struct {
		in            string
		expectedError string
	}{
		{
			in: "cd",
		},
		{
			in:            "ab",
			expectedError: "string must not match regular expression: '[ab]+'",
		},
	}
)

func TestStringDenyRegexp(t *testing.T) {
	for _, tc := range stringDenyRegexpTestCases {
		err := StringDenyRegexp(stringDenyRegexpRegexp).Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDenyRegexp))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDenyRegexp(b *testing.B) {
	for _, tc := range stringDenyRegexpTestCases {
		rule := StringDenyRegexp(stringDenyRegexpRegexp)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringDNSLabelTestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringDNSLabelTestCases {
		rule := StringDNSLabel()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringDNSSubdomainTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"s", false},
	{"sa", false},
	{"a-1", false},
	{"a--2", false},
	{"a-b-c", false},
	{"a--b--c", false},
	{"0", false},
	{"a.1", false},
	{"a.b", false},
	{"1.b", false},
	{"a.b.c", false},
	{"a.1.c", false},
	{"aa.bb", false},
	{"1.2.3.4", false},
	{"1a.2b.3c.4d", false},
	{"a--b--c.123", false},
	{strings.Repeat("l", 253), false},
	{"", true},
	{" ", true},
	{strings.Repeat("l", 254), true},
	{"tesT", true},
	{"test?", true},
	{"test this", true},
	{"1_2", true},
	{"L", true},
	{"a@b", true},
	{"-", true},
	{"a-", true},
	{"0-", true},
	{"-b", true},
	{"-1", true},
	{"A.1", true},
	{".2.3.4", true},
	{"1a.2B.3c.4d", true},
	{"a--b--c.", true},
}

func TestStringDNSSubdomain(t *testing.T) {
	for _, tc := range stringDNSSubdomainTestCases {
		err := StringDNSSubdomain().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDNSSubdomain))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDNSSubdomain(b *testing.B) {
	for _, tc := range stringDNSSubdomainTestCases {
		rule := StringDNSSubdomain()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringASCIITestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringASCIITestCases {
		rule := StringASCII()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringUUIDTestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringUUIDTestCases {
		rule := StringUUID()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringEmailTestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringEmailTestCases {
		rule := StringEmail()
		for range b.N {
			_ = rule.Validate(tc.in)
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
	for _, tc := range urlTestCases {
		rule := StringURL()
		for range b.N {
			_ = rule.Validate(tc.url)
		}
	}
}

var stringMACTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"3D:F2:C9:A6:B3:4F", false},
	{"00:25:96:FF:FE:12:34:56", false},
	{"3D-F2-C9-A6-B3:4F", true},
	{"123", true},
	{"", true},
	{"abacaba", true},
	{"0025:96FF:FE12:3456", true},
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
	for _, tc := range stringMACTestCases {
		rule := StringMAC()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringIPTestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringIPTestCases {
		rule := StringIP()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringIPv4TestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringIPv4TestCases {
		rule := StringIPv4()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringIPv6TestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringIPv6TestCases {
		rule := StringIPv6()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringCIDRTestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringCIDRTestCases {
		rule := StringCIDR()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringCIDRv4TestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringCIDRv4TestCases {
		rule := StringCIDRv4()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringCIDRv6TestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringCIDRv6TestCases {
		rule := StringCIDRv6()
		for range b.N {
			_ = rule.Validate(tc.in)
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
	for _, tc := range stringJSONTestCases {
		rule := StringJSON()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var validSemverTestCases = []string{
	"0.0.4",
	"1.2.3",
	"10.20.30",
	"1.1.2-prerelease+meta",
	"1.1.2+meta",
	"1.1.2+meta-valid",
	"1.0.0-alpha",
	"1.0.0-beta",
	"1.0.0-alpha.beta",
	"1.0.0-alpha.beta.1",
	"1.0.0-alpha.1",
	"1.0.0-alpha0.valid",
	"1.0.0-alpha.0valid",
	"1.0.0-alpha-a.b-c-somethinglong+build.1-aef.1-its-okay",
	"1.0.0-rc.1+build.1",
	"2.0.0-rc.1+build.123",
	"1.2.3-beta",
	"10.2.3-DEV-SNAPSHOT",
	"1.2.3-SNAPSHOT-123",
	"1.0.0",
	"2.0.0",
	"1.1.7",
	"2.0.0+build.1848",
	"2.0.1-alpha.1227",
	"1.0.0-alpha+beta",
	"1.2.3----RC-SNAPSHOT.12.9.1--.12+788",
	"1.2.3----R-S.12.9.1--.12+meta",
	"1.2.3----RC-SNAPSHOT.12.9.1--.12",
	"1.0.0+0.build.1-rc.10000aaa-kk-0.1",
	"99999999999999999999999.999999999999999999.99999999999999999",
	"1.0.0-0A.is.legal",
	"0.1.0",
	"1.0.0+20130313144700",
	"1.0.0-beta+exp.sha.5114f85",
	"2.7.3-rc.1+build.11.e0f985a",
}

var invalidSemverTestCases = []string{
	"1",
	"1.2",
	"1.2.3-0123",
	"1.2.3-0123.0123",
	"1.1.2+.123",
	"+invalid",
	"-invalid",
	"-invalid+invalid",
	"-invalid.01",
	"alpha",
	"alpha.beta",
	"alpha.beta.1",
	"alpha.1",
	"alpha+beta",
	"alpha_beta",
	"alpha.",
	"alpha..",
	"beta",
	"1.0.0-alpha_beta",
	"-alpha.",
	"1.0.0-alpha..",
	"1.0.0-alpha..1",
	"1.0.0-alpha...1",
	"1.0.0-alpha....1",
	"1.0.0-alpha.....1",
	"1.0.0-alpha......1",
	"1.0.0-alpha.......1",
	"01.1.1",
	"1.01.1",
	"1.1.01",
	"1.2",
	"1.2.3.DEV",
	"1.2-SNAPSHOT",
	"1.2.31.2.3----RC-SNAPSHOT.12.09.1--..12+788",
	"1.2-RC-SNAPSHOT",
	"-1.0.3-gamma+b7718",
	"+justmeta",
	"9.8.7+meta+meta",
	"9.8.7-whatever+meta+meta",
	"99999999999999999999999.999999999999999999.99999999999999999----RC-SNAPSHOT.12.09.1--------------------------------..12",
	"",
	"1.2.3.4",
	"01.2.3",
	"1.02.3",
	"1.2.03",
	"1.2.3-",
	"1.2.3-01",
	"v1.2.3",
	"1.2.3+build..1",
}

func TestStringSemver(t *testing.T) {
	rule := StringSemver()
	t.Run("valid versions", func(t *testing.T) {
		for _, version := range validSemverTestCases {
			t.Run(fmt.Sprintf("%q", version), func(t *testing.T) {
				assert.NoError(t, rule.Validate(version))
			})
		}
	})
	t.Run("invalid versions", func(t *testing.T) {
		for _, version := range invalidSemverTestCases {
			t.Run(fmt.Sprintf("%q", version), func(t *testing.T) {
				err := rule.Validate(version)
				assert.EqualError(t, err, "string must be a valid semantic version")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringSemver))
			})
		}
	})
}

func BenchmarkStringSemver(b *testing.B) {
	rule := StringSemver()
	for b.Loop() {
		_ = rule.Validate("2.7.3-rc.1+build.11.e0f985a")
	}
}

var stringCVETestCases = map[string]struct {
	in            string
	expectedError string
}{
	"four digit sequence": {
		in: "CVE-1999-0001",
	},
	"four digit year before 1999": {
		in: "CVE-1998-0001",
	},
	"zero sequence": {
		in: "CVE-2021-0000",
	},
	"sequence with leading zero": {
		in: "CVE-2014-0160",
	},
	"five digit sequence with leading zero": {
		in: "CVE-2021-00001",
	},
	"sequence with two leading zeroes": {
		in: "CVE-2021-0990",
	},
	"five digit sequence": {
		in: "CVE-2021-44228",
	},
	"long sequence": {
		in: "CVE-2024-12345",
	},
	"nineteen digit sequence": {
		in: "CVE-2024-1234567890123456789",
	},
	"empty": {
		in:            "",
		expectedError: "string must be a valid CVE ID",
	},
	"lowercase prefix": {
		in:            "cve-2021-44228",
		expectedError: "string must be a valid CVE ID",
	},
	"short sequence": {
		in:            "CVE-2021-123",
		expectedError: "string must be a valid CVE ID",
	},
	"letters in sequence": {
		in:            "CVE-2021-ABCD",
		expectedError: "string must be a valid CVE ID",
	},
	"five digit year": {
		in:            "CVE-10000-0001",
		expectedError: "string must be a valid CVE ID",
	},
	"twenty digit sequence": {
		in:            "CVE-2024-12345678901234567890",
		expectedError: "string must be a valid CVE ID",
	},
}

func TestStringCVE(t *testing.T) {
	for name, tt := range stringCVETestCases {
		t.Run(name, func(t *testing.T) {
			err := StringCVE().Validate(tt.in)
			if tt.expectedError != "" {
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCVE))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

var stringE164TestCases = map[string]struct {
	in            string
	expectedError string
}{
	"minimum length": {
		in: "+12",
	},
	"maximum length": {
		in: "+123456789012345",
	},
	"common US number": {
		in: "+14155552671",
	},
	"missing plus sign": {
		in:            "14155552671",
		expectedError: "string must be a valid E.164 phone number",
	},
	"starts with zero": {
		in:            "+0123456789",
		expectedError: "string must be a valid E.164 phone number",
	},
	"too short": {
		in:            "+1",
		expectedError: "string must be a valid E.164 phone number",
	},
	"too long": {
		in:            "+1234567890123456",
		expectedError: "string must be a valid E.164 phone number",
	},
	"contains spaces": {
		in:            "+1 4155552671",
		expectedError: "string must be a valid E.164 phone number",
	},
	"contains punctuation": {
		in:            "+1-415-555-2671",
		expectedError: "string must be a valid E.164 phone number",
	},
	"empty": {
		expectedError: "string must be a valid E.164 phone number",
	},
}

func TestStringE164(t *testing.T) {
	t.Parallel()

	for name, tt := range stringE164TestCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := StringE164().Validate(tt.in)
			if tt.expectedError != "" {
				assert.Require(t, assert.Error(t, err))
				assert.EqualError(t, err, tt.expectedError)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringE164))
				return
			}
			assert.NoError(t, err)
		})
	}
}

func BenchmarkStringE164(b *testing.B) {
	tests := map[string]string{
		"valid":   "+14155552671",
		"invalid": "+1-415-555-2671",
	}

	for name, in := range tests {
		b.Run(name, func(b *testing.B) {
			rule := StringE164()
			for b.Loop() {
				_ = rule.Validate(in)
			}
		})
	}
}

func BenchmarkStringCVE(b *testing.B) {
	rule := StringCVE()
	for b.Loop() {
		_ = rule.Validate("CVE-2021-44228")
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
	for _, tc := range stringContainsTestCases {
		rule := StringContains(tc.substrings...)
		for range b.N {
			_ = rule.Validate(tc.in)
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
	for _, tc := range stringExcludesTestCases {
		rule := StringExcludes(tc.substrings...)
		for range b.N {
			_ = rule.Validate(tc.in)
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
	for _, tc := range stringStartsWithTestCases {
		rule := StringStartsWith(tc.prefixes...)
		for range b.N {
			_ = rule.Validate(tc.in)
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
	for _, tc := range stringEndsWithTestCases {
		rule := StringEndsWith(tc.suffixes...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringTitleTestCases = []*struct {
	in         string
	shouldFail bool
}{
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
	for _, tc := range stringTitleTestCases {
		rule := StringTitle()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var (
	errGitRefEmpty           = errors.New("git reference must not be empty")
	errGitRefEndsWithDot     = errors.New("git reference must not end with a '.'")
	errGitRefAtLeastOneSlash = errors.New("git reference must contain at least one '/'")
	errGitRefEmptyPart       = errors.New("git reference must not have empty parts")
	errGitRefStartsWithDash  = errors.New("git branch and tag references must not start with '-'")
	errGitRefForbiddenChars  = errors.New("git reference contains forbidden characters")
)

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
				assert.ErrorContains(
					t,
					err,
					"see https://git-scm.com/docs/git-check-ref-format for more information on Git reference naming rules",
				)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringGitRef(b *testing.B) {
	for _, tc := range stringGitRefTestCases {
		rule := StringGitRef()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func prepareFileSystemTests(t testing.TB) (root string) {
	t.Helper()
	root = t.TempDir()
	t.Setenv("HOME", root)
	for _, path := range []struct {
		path  string
		perm  os.FileMode
		isDir bool
	}{
		{"file1", 0o755, false},
		{"dir1", 0o755, true},
		{"dir1/file2", 0o755, false},
		{"dir-no-perm", 0o000, true},
		{"dir1/file-no-perm", 0o000, false},
	} {
		if path.isDir {
			err := os.MkdirAll(filepath.Join(root, path.path), path.perm)
			assert.Require(t, assert.NoError(t, err))
		} else {
			err := os.WriteFile(filepath.Join(root, path.path), []byte{}, path.perm)
			assert.Require(t, assert.NoError(t, err))
		}
	}
	return root
}

type stringFileSystemPathTestCase struct {
	in          string
	expectedErr error
}

func getStringFileSystemPathTestCases(root string) []*stringFileSystemPathTestCase {
	addRoot := func(path string) string {
		// We're not using filepath.Join because it cleans the path.
		return root + string(filepath.Separator) + path
	}
	return []*stringFileSystemPathTestCase{
		{"~/dir1", nil},
		{"~/dir1/", nil},
		{addRoot("dir1"), nil},
		{addRoot("dir1/file2"), nil},
		{"~/dir1/file2", nil},
		{addRoot("dir1/file2/.."), nil},
		{"~/dir1/file2/..", nil},
		{"~/dir1/file2/../../", nil},
		{addRoot("."), nil},
		{addRoot("./"), nil},
		{addRoot("./file1"), nil},
		{addRoot("dir-no-perm"), nil},
		{addRoot("dir1/file-no-perm"), nil},
		{addRoot("dir1/file2/"), syscall.ENOTDIR},
		{"~/dir1/../file1/", syscall.ENOTDIR},
		{addRoot("non-existing-dir"), errFilePathNotExists},
		{"", errFilePathEmpty},
		{"	", errFilePathEmpty},
	}
}

func TestStringFileSystemPath(t *testing.T) {
	root := prepareFileSystemTests(t)
	for _, tc := range getStringFileSystemPathTestCases(root) {
		err := StringFileSystemPath().Validate(tc.in)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error())
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringFileSystemPath))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringFileSystemPath(b *testing.B) {
	root := prepareFileSystemTests(b)
	testCases := getStringFileSystemPathTestCases(root)
	for _, tc := range testCases {
		rule := StringFileSystemPath()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func getStringFilePathTestCases(root string) []*stringFileSystemPathTestCase {
	addRoot := func(path string) string {
		// We're not using filepath.Join because it cleans the path.
		return root + string(filepath.Separator) + path
	}
	return []*stringFileSystemPathTestCase{
		{addRoot("dir1/file2"), nil},
		{"~/dir1/file2", nil},
		{addRoot("./file1"), nil},
		{addRoot("dir1/file-no-perm"), nil},
		{addRoot("dir-no-perm"), errFilePathNotFile},
		{addRoot("dir1"), errFilePathNotFile},
		{addRoot("dir1/file2/.."), errFilePathNotFile},
		{addRoot("."), errFilePathNotFile},
		{addRoot("./"), errFilePathNotFile},
		{"~/dir1/file2/..", errFilePathNotFile},
		{"~/dir1/file2/../../", errFilePathNotFile},
		{"~/dir1", errFilePathNotFile},
		{"~/dir1/", errFilePathNotFile},
		{addRoot("dir1/file2/"), syscall.ENOTDIR},
		{"~/dir1/../file1/", syscall.ENOTDIR},
		{addRoot("non-existing-dir"), errFilePathNotExists},
		{"", errFilePathEmpty},
		{"	", errFilePathEmpty},
	}
}

func TestStringFilePath(t *testing.T) {
	root := prepareFileSystemTests(t)
	for _, tc := range getStringFilePathTestCases(root) {
		err := StringFilePath().Validate(tc.in)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error())
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringFilePath))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringFilePath(b *testing.B) {
	root := prepareFileSystemTests(b)
	testCases := getStringFilePathTestCases(root)
	for _, tc := range testCases {
		rule := StringFilePath()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

func getStringDirPathTestCases(root string) []*stringFileSystemPathTestCase {
	addRoot := func(path string) string {
		// We're not using filepath.Join because it cleans the path.
		return root + string(filepath.Separator) + path
	}
	return []*stringFileSystemPathTestCase{
		{addRoot("dir1"), nil},
		{addRoot("dir1/file2/.."), nil},
		{addRoot("."), nil},
		{addRoot("./"), nil},
		{"~/dir1/file2/..", nil},
		{"~/dir1/file2/../../", nil},
		{"~/dir1", nil},
		{"~/dir1/", nil},
		{addRoot("dir-no-perm"), nil},
		{addRoot("dir1/file-no-perm"), errFilePathNotDir},
		{addRoot("dir1/file2"), errFilePathNotDir},
		{"~/dir1/file2", errFilePathNotDir},
		{addRoot("./file1"), errFilePathNotDir},
		{addRoot("dir1/file2/"), syscall.ENOTDIR},
		{"~/dir1/../file1/", syscall.ENOTDIR},
		{addRoot("non-existing-dir"), errFilePathNotExists},
		{"", errFilePathEmpty},
		{"	", errFilePathEmpty},
	}
}

func TestStringDirPath(t *testing.T) {
	root := prepareFileSystemTests(t)
	for _, tc := range getStringDirPathTestCases(root) {
		err := StringDirPath().Validate(tc.in)
		if tc.expectedErr != nil {
			assert.ErrorContains(t, err, tc.expectedErr.Error())
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDirPath))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringDirPath(b *testing.B) {
	root := prepareFileSystemTests(b)
	testCases := getStringDirPathTestCases(root)
	for _, tc := range testCases {
		rule := StringDirPath()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

// test cases copied from Go's [filepath] standard library.
var stringMatchFileSystemPathTestCases = []*struct {
	pattern, in string
	shouldFail  bool
}{
	{"abc", "abc", false},
	{"*", "abc", false},
	{"*c", "abc", false},
	{"a*", "a", false},
	{"a*", "abc", false},
	{"a*/b", "abc/b", false},
	{"a*b*c*d*e*/f", "axbxcxdxe/f", false},
	{"a*b*c*d*e*/f", "axbxcxdxexxx/f", false},
	{"a*b?c*x", "abxbbxdbxebxczzx", false},
	{"ab[c]", "abc", false},
	{"ab[b-d]", "abc", false},
	{"ab[^e-g]", "abc", false},
	{"a\\*b", "a*b", false},
	{"a?b", "a☺b", false},
	{"a[^a]b", "a☺b", false},
	{"[a-ζ]*", "α", false},
	{"[\\]a]", "]", false},
	{"[\\-]", "-", false},
	{"*x", "xxx", false},
	{"[x\\-]", "x", false},
	{"[x\\-]", "-", false},
	{"[\\-x]", "x", false},
	{"[\\-x]", "-", false},
	{"a*/b", "a/c/b", true},
	{"ab[e-g]", "abc", true},
	{"ab[^c]", "abc", true},
	{"a*", "ab/c", true},
	{"a*b*c*d*e*/f", "axbxcxdxe/xxx/f", true},
	{"a*b*c*d*e*/f", "axbxcxdxexxx/fff", true},
	{"a*b?c*x", "abxbbxdbxebxczzy", true},
	{"ab[^b-d]", "abc", true},
	{"a???b", "a☺b", true},
	{"a\\*b", "ab", true},
	{"a[^a][^a][^a]b", "a☺b", true},
	{"*[a-ζ]", "A", true},
	{"a?b", "a/b", true},
	{"a*b", "a/b", true},
	{"[x\\-]", "z", true},
	{"[\\-x]", "a", true},
	{"[]a]", "]", true},
	{"[-]", "-", true},
	{"[x-]", "x", true},
	{"[x-]", "-", true},
	{"[x-]", "z", true},
	{"[-x]", "x", true},
	{"[-x]", "-", true},
	{"[-x]", "a", true},
	{"\\", "a", true},
	{"[a-b-c]", "a", true},
	{"[", "a", true},
	{"[^", "a", true},
	{"[^bc", "a", true},
	{"a[", "a", true},
	{"a[", "ab", true},
	{"a[", "x", true},
	{"a/b[", "x", true},
}

func TestStringMatchFileSystemPath(t *testing.T) {
	for _, tc := range stringMatchFileSystemPathTestCases {
		err := StringMatchFileSystemPath(tc.pattern).Validate(tc.in)
		if tc.shouldFail {
			if !strings.Contains(err.Error(), "string must match file path pattern") &&
				!strings.Contains(err.Error(), filepath.ErrBadPattern.Error()) {
				assert.Fail(t, "unexpected error: %v", err)
			}
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMatchFileSystemPath))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringMatchFileSystemPath(b *testing.B) {
	for _, tc := range stringMatchFileSystemPathTestCases {
		rule := StringMatchFileSystemPath(tc.pattern)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

// test cases copied from Go's [regexp] standard library.
var stringRegexpTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{``, false},
	{`.`, false},
	{`^.$`, false},
	{`a`, false},
	{`a*`, false},
	{`a+`, false},
	{`a?`, false},
	{`a|b`, false},
	{`a*|b*`, false},
	{`(a*|b)(c*|d)`, false},
	{`[a-z]`, false},
	{`[a-abc-c\-\]\[]`, false},
	{`[a-z]+`, false},
	{`[abc]`, false},
	{`[^1234]`, false},
	{`[^\n]`, false},
	{`\!\\`, false},
	{`*`, true},
	{`+`, true},
	{`?`, true},
	{`(abc`, true},
	{`abc)`, true},
	{`x[a-z`, true},
	{`[z-a]`, true},
	{`abc\`, true},
	{`a**`, true},
	{`a*+`, true},
	{`\x`, true},
	{strings.Repeat(`\pL`, 27000), true},
}

func TestStringRegexp(t *testing.T) {
	for _, tc := range stringRegexpTestCases {
		err := StringRegexp().Validate(tc.in)
		if tc.shouldFail {
			assert.ErrorContains(t, err, "string must be a valid regular expression")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringRegexp))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringRegexp(b *testing.B) {
	for _, tc := range stringRegexpTestCases {
		rule := StringRegexp()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

type stringCrontabTestCase struct {
	in         string
	shouldFail bool
}

func getStringCronTestCases() []*stringCrontabTestCase {
	testCases := []*stringCrontabTestCase{
		{"@annually", false},
		{"@yearly", false},
		{"@monthly", false},
		{"@weekly", false},
		{"@daily", false},
		{"@hourly", false},
		{"@reboot", false},
		{"* * * * *", false},
		{"* * * JAN,MAY,DEC *", false},
		{"* * * JAN-DEC *", false},
		{"* * * FEB-MAY/2 *", false},
		{"* * * fEb-may/10 *", false},
		{"* * * SEP-SEP/2 *", false},
		{"* * * JAN-1 *", false},
		{"* * * JAN-12 *", false},
		{"* * * 1-DEC *", false},
		{"* * * * FRI-7", false},
		{"* * * * 2-WED", false},
		{"* * * * THU-FRI", false},
		{"* * * * TUE-THU/10", false},
		{"* * * * SUN-MON", false},
		{"* * * * WED-3", false},
		{"* * * * THU,FRI,MON", false},
		{"* * * * *", false},
		{"", true},
		{"  @hourly", true},
		{"1h @every", true},
		{"@every 1Y", true},
		{"wrong", true},
		{"@minutely", true},
		{"@every 1h", true},
		{"@every 1h30m10ts", true},
		{"a * * * *", true},
		{"1 b * * *", true},
		{"1 1 c * *", true},
		{"1 1 1 d *", true},
		{"1 1 1 1 e", true},
		{"* * * MAZ *", true},
		{"* * * MAY-FEB/2 *", true},
		{"* * * MAY-2 *", true},
		{"* * * 2-JAN *", true},
		{"* * * FEB-JUN/-10 *", true},
		{"* * * JAN,MAY,DEZ *", true},
		{"* * * * MOZ", true},
		{"* * * * MON-SUN", true},
		{"* * * * 7-FRI", true},
		{"* * * * WED-2", true},
		{"* * * * MON-FRI/-10", true},
		{"* * * * THU,FRI,MOZ", true},
	}
	createCron := func(n int, format string, a ...any) string {
		fields := strings.Fields("* * * * *")
		fields[n] = fmt.Sprintf(format, a...)
		return strings.Join(fields, " ")
	}
	for _, field := range []struct {
		n, lower, upper int
	}{
		{0, 0, 59},
		{1, 0, 23},
		{2, 1, 31},
		{3, 1, 12},
		{4, 0, 7},
	} {
		getRandom := func() int {
			return field.lower + rand.Intn(field.upper-field.lower)
		}
		testCases = append(testCases,
			&stringCrontabTestCase{createCron(field.n, "%d", getRandom()), false},
			&stringCrontabTestCase{createCron(field.n, "%d", field.lower), false},
			&stringCrontabTestCase{createCron(field.n, "%d", field.upper), false},
			&stringCrontabTestCase{createCron(field.n, "%d,%d", field.lower, field.upper), false},
			&stringCrontabTestCase{createCron(field.n, "%d,%d", field.upper, field.lower), false},
			&stringCrontabTestCase{createCron(field.n, "%d-%d", field.lower, field.upper), false},
			&stringCrontabTestCase{createCron(field.n, "%d-%d/10", field.lower, field.upper), false},
			&stringCrontabTestCase{createCron(field.n, "*/10"), false},
			&stringCrontabTestCase{createCron(field.n, "%d", field.lower-1), true},
			&stringCrontabTestCase{createCron(field.n, "%d", field.upper+1), true},
			&stringCrontabTestCase{createCron(field.n, "%d,", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "%d,%d", field.lower, field.upper+1), true},
			&stringCrontabTestCase{createCron(field.n, "%d,%d", field.lower-1, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d/10", getRandom()), true},
			&stringCrontabTestCase{createCron(field.n, "%d,%d/10", field.lower, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "a"), true},
			&stringCrontabTestCase{createCron(field.n, "%d,a", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "a,%d", field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d-", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "%d-/", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "-/"), true},
			&stringCrontabTestCase{createCron(field.n, "%d-%d/", field.lower, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d-%d/a", field.lower, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d-%d/-10", field.lower, field.upper), true},
			&stringCrontabTestCase{createCron(field.n, "%d-*/10", field.lower), true},
			&stringCrontabTestCase{createCron(field.n, "*-*/10"), true},
			&stringCrontabTestCase{createCron(field.n, "*-%d/10", field.upper), true},
		)
	}
	for month := range crontabMonthsMap {
		testCases = append(testCases, &stringCrontabTestCase{createCron(3, "%s", month), false})
	}
	for day := range crontabDaysMap {
		// Skip special cases for Sunday.
		if strings.Contains(day, "-") {
			continue
		}
		testCases = append(testCases, &stringCrontabTestCase{createCron(4, "%s", day), false})
	}
	return testCases
}

func TestStringCrontab(t *testing.T) {
	for _, tc := range getStringCronTestCases() {
		t.Run(tc.in, func(t *testing.T) {
			err := StringCrontab().Validate(tc.in)
			if tc.shouldFail {
				assert.ErrorContains(t, err, "string must be a valid cron schedule expression")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCrontab))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringCrontab(b *testing.B) {
	for range b.N {
		testCases := getStringCronTestCases()
		for _, tc := range testCases {
			_ = StringCrontab().Validate(tc.in)
		}
	}
}

var stringDateTimeTestCases = []*struct {
	layout string
	in     string
	errMsg string
}{
	{time.RFC3339, "2024-01-01T15:00:00Z", ""},
	{time.RFC3339, "2024-01-01T15:00:00+01:00", ""},
	{time.DateTime, "2024-01-01 15:00:00", ""},
	{time.DateOnly, "2024-01-01", ""},
	{time.TimeOnly, "15:00:00", ""},
	{
		"invalid-layout",
		"2024-01-01T15:00:00Z",
		"string must be a valid date and time in 'invalid-layout' format",
	},
	{
		time.RFC3339,
		"2024-01-01 15:00:00Z",
		"string must be a valid date and time in '2006-01-02T15:04:05Z07:00' format",
	},
	{
		"15:04",
		"15:00:00",
		"string must be a valid date and time in '15:04'",
	},
}

func TestStringDateTime(t *testing.T) {
	for _, tc := range stringDateTimeTestCases {
		t.Run(tc.layout+tc.in, func(t *testing.T) {
			err := StringDateTime(tc.layout).Validate(tc.in)
			if tc.errMsg != "" {
				assert.ErrorContains(t, err, tc.errMsg)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDateTime))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringDateTime(b *testing.B) {
	for _, tc := range stringDateTimeTestCases {
		rule := StringDateTime(tc.layout)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringTimeZoneTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"UTC", false},
	{"America/New_York", false},
	{"Europe/Warsaw", false},
	{"", true},
	{"Local", true},
	{"America/New_Yorker", true},
	{"x/x", true},
	{"America/Warsaw", true},
}

func TestStringTimeZone(t *testing.T) {
	for _, tc := range stringTimeZoneTestCases {
		t.Run(tc.in, func(t *testing.T) {
			err := StringTimeZone().Validate(tc.in)
			if tc.shouldFail {
				assert.ErrorContains(t, err, "string must be a valid IANA Time Zone Database code")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringTimeZone))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringTimeZone(b *testing.B) {
	for _, tc := range stringDateTimeTestCases {
		rule := StringTimeZone()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringAlphaTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"test", false},
	{"tEsT", false},
	{"s", false},
	{"LOL", false},
	{"test-this", true},
	{" test", true},
	{"  ", true},
	{" ", true},
	{"test1", true},
	{"tęst", true},
}

func TestStringAlpha(t *testing.T) {
	for _, tc := range stringAlphaTestCases {
		err := StringAlpha().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAlpha))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringAlpha(b *testing.B) {
	for _, tc := range stringAlphaTestCases {
		rule := StringAlpha()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringAlphanumericTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"test", false},
	{"tEsT", false},
	{"s", false},
	{"4", false},
	{"LOL", false},
	{"test1", false},
	{"-921", true},
	{"test-this", true},
	{" test", true},
	{" 1", true},
	{"  ", true},
	{" ", true},
	{"tęst", true},
	{"tęst1", true},
}

func TestStringAlphanumeric(t *testing.T) {
	for _, tc := range stringAlphanumericTestCases {
		err := StringAlphanumeric().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAlphanumeric))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringAlphanumeric(b *testing.B) {
	for _, tc := range stringAlphanumericTestCases {
		rule := StringAlphanumeric()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringAlphaUnicodeTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"test", false},
	{"tEsT", false},
	{"s", false},
	{"LOL", false},
	{"tęst", false},
	{"汉语", false},
	{"一二三", false},
	{"test-this", true},
	{" test", true},
	{"  ", true},
	{" ", true},
	{"test1", true},
	{"汉语!", true},
	{"1汉语", true},
}

func TestStringAlphaUnicode(t *testing.T) {
	for _, tc := range stringAlphaUnicodeTestCases {
		err := StringAlphaUnicode().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAlphaUnicode))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringAlphaUnicode(b *testing.B) {
	for _, tc := range stringAlphaUnicodeTestCases {
		rule := StringAlphaUnicode()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringAlphanumericUnicodeTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"test", false},
	{"tEsT", false},
	{"s", false},
	{"5", false},
	{"LOL", false},
	{"tęst", false},
	{"汉语", false},
	{"1汉语", false},
	{"test1", false},
	{"tęst1", false},
	{"一二三", false},
	{"-550", true},
	{"test-this", true},
	{" test", true},
	{"  ", true},
	{" ", true},
	{"汉语!", true},
	{"-921", true},
	{" 1", true},
}

func TestStringAlphanumericUnicode(t *testing.T) {
	for _, tc := range stringAlphanumericUnicodeTestCases {
		err := StringAlphanumericUnicode().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringAlphanumericUnicode))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringAlphanumericUnicode(b *testing.B) {
	for _, tc := range stringAlphanumericUnicodeTestCases {
		rule := StringAlphanumericUnicode()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var stringFQDNTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"test.example.com", false},
	{"example.com", false},
	{"example24.com", false},
	{"test.example24.com", false},
	{"test24.example24.com", false},
	{"test.example.com.", false},
	{"example.com.", false},
	{"example24.com.", false},
	{"test.example24.com.", false},
	{"test24.example24.com.", false},
	{"24.example24.com", false},
	{"test.24.example.com", false},
	{"test24.example24.com..", true},
	{"example", true},
	{"192.168.0.1", true},
	{"email@example.com", true},
	{"2001:cdba:0000:0000:0000:0000:3257:9652", true},
	{"2001:cdba:0:0:0:0:3257:9652", true},
	{"2001:cdba::3257:9652", true},
	{"", true},
}

func TestStringFQDN(t *testing.T) {
	for _, tc := range stringFQDNTestCases {
		err := StringFQDN().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringFQDN))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringFQDN(b *testing.B) {
	for _, tc := range stringFQDNTestCases {
		rule := StringFQDN()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var (
	errK8sQualifiedNameEmptyPrefixPart = errors.New("prefix part must not be empty")
	errK8sQualifiedNamePrefixLength    = errors.New("prefix part length must be less than or equal to 253")
	errK8sQualifiedNamePrefixRegexp    = errors.New(
		`prefix part string must match regular expression: '^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$'`,
	)
	errK8sQualifiedNameTooManyParts   = errors.New("qualified name must have at most 2 parts separated by a '/'")
	errK8sQualifiedNameEmptyNamePart  = errors.New("name part must not be empty")
	errK8sQualifiedNameNamePartLength = errors.New("name part length must be less than or equal to 63")
	errK8sQualifiedNameNamePartRegexp = errors.New(
		"name part string must match regular expression: '^([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$'",
	)
)

var stringK8sQualifiedNameTestCases = []*struct {
	in          string
	expectedErr error
}{
	{"simple", nil},
	{"now-with-dashes", nil},
	{"1-starts-with-num", nil},
	{"1234", nil},
	{"simple/simple", nil},
	{"now-with-dashes/simple", nil},
	{"now-with-dashes/now-with-dashes", nil},
	{"now.with.dots/simple", nil},
	{"now-with.dashes-and.dots/simple", nil},
	{"1-num.2-num/3-num", nil},
	{"1234/5678", nil},
	{"1.2.3.4/5678", nil},
	{"Uppercase_Is_OK_123", nil},
	{"example.com/Uppercase_Is_OK_123", nil},
	{"requests.storage-foo", nil},
	{strings.Repeat("a", 63), nil},
	{strings.Repeat("a", 253) + "/" + strings.Repeat("b", 63), nil},
	// BAD
	{"/", errK8sQualifiedNameEmptyPrefixPart},
	{"nospecialchars%^=@", errK8sQualifiedNameNamePartRegexp},
	{"cantendwithadash-", errK8sQualifiedNameNamePartRegexp},
	{"-cantstartwithadash-", errK8sQualifiedNameNamePartRegexp},
	{"example.com/abc$", errK8sQualifiedNameNamePartRegexp},
	{"only/one/slash", errK8sQualifiedNameTooManyParts},
	{"Example.com/abc", errK8sQualifiedNamePrefixRegexp},
	{"example_com/abc", errK8sQualifiedNamePrefixRegexp},
	{"example.com/", errK8sQualifiedNameEmptyNamePart},
	{"/simple", errK8sQualifiedNameEmptyPrefixPart},
	{"not.Valid/simple", errK8sQualifiedNamePrefixRegexp},
	{strings.Repeat("a", 64), errK8sQualifiedNameNamePartLength},
	{strings.Repeat("a", 254) + "/abc", errK8sQualifiedNamePrefixLength},
	{strings.Repeat("a", 253) + "/" + strings.Repeat("b", 64), errors.New("length must be between 1 and 317")},
}

func TestStringKubernetesQualifiedName(t *testing.T) {
	for _, tc := range stringK8sQualifiedNameTestCases {
		t.Run(tc.in, func(t *testing.T) {
			err := StringKubernetesQualifiedName().Validate(tc.in)
			if tc.expectedErr != nil {
				assert.ErrorContains(t, err, tc.expectedErr.Error())
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringKubernetesQualifiedName))
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkStringKubernetesQualifiedName(b *testing.B) {
	for _, tc := range stringK8sQualifiedNameTestCases {
		rule := StringKubernetesQualifiedName()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var validISBNTestCases = map[string]string{
	"isbn 10 hyphenated":                  "0-306-40615-2",
	"isbn 10 plain":                       "0306406152",
	"isbn 10 x check":                     "0-9752298-0-X",
	"isbn 10 spaced":                      "0 9752298 0 x",
	"isbn 10 library converter numeric":   "0394170660",
	"isbn 10 library converter alternate": "0717941728",
	"isbn 10 library converter x check":   "087779443X",
	"isbn 10 MARC hyphenated":             "0-87068-693-3",
	"isbn 13 hyphenated":                  "978-0-306-40615-7",
	"isbn 13 plain":                       "9780306406157",
	"isbn 13 grouped":                     "978-3-16-148410-0",
	"isbn 13 agency manual hyphenated":    "978-92-95055-12-4",
	"isbn 13 agency manual spaced":        "978 92 95055 12 4",
	"isbn 13 agency manual compact":       "9789295055124",
	"isbn 13 agency manual hardback":      "978-951-45-9693-3",
	"isbn 13 agency manual paperback":     "978-951-45-9694-0",
	"isbn 13 agency manual PDF":           "978-951-45-9695-7",
	"isbn 13 agency manual EPUB":          "978-951-45-9696-4",
	"isbn 13 library converter first":     "9780060723804",
	"isbn 13 library converter second":    "9780060799748",
	"isbn 13 979 prefix":                  "979-10-90636-07-1",
}

var invalidISBNTestCases = map[string]string{
	"empty":                               "",
	"isbn 10 failed check":                "0-306-40615-3",
	"isbn 10 x check mutation":            "0877794430",
	"isbn 10 x in body":                   "08777X443X",
	"isbn 10 x in fourth position":        "087X79443X",
	"isbn 10 short":                       "087779443",
	"isbn 10 trailing space":              "087779443X ",
	"isbn 13 failed check":                "978-0-306-40615-8",
	"isbn 13 manual check mutation":       "978-92-95055-12-5",
	"isbn 13 checksum valid wrong prefix": "9779295055125",
	"isbn 13 x in body":                   "978-92-X5055-12-4",
	"isbn 13 x check":                     "978-92-95055-12-X",
	"isbn 13 en dash separators":          "978–92–95055–12–4",
	"isbn 13 full width digits":           "９７８９２９５０５５１２４",
	"isbn 13 display prefix":              "ISBN 978-92-95055-12-4",
	"repeated separator":                  "978--0-306-40615-7",
	"letters":                             "abc",
}

func TestStringISBN(t *testing.T) {
	rule := StringISBN()
	t.Run("valid inputs", func(t *testing.T) {
		for name, input := range validISBNTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for name, input := range invalidISBNTestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(
					t,
					err,
					"string must be a valid International Standard Book Number (ISBN) in ISBN-10 or ISBN-13 format",
				)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN))
			})
		}
	})
}

func TestStringISBN_VeryLargeInvalid(t *testing.T) {
	input := strings.Repeat("0", 1<<20)
	tests := map[string]struct {
		rule    govy.Rule[string]
		message string
		code    govy.ErrorCode
	}{
		"isbn": {
			rule:    StringISBN(),
			message: "string must be a valid International Standard Book Number (ISBN) in ISBN-10 or ISBN-13 format",
			code:    ErrorCodeStringISBN,
		},
		"isbn-10": {
			rule:    StringISBN10(),
			message: "string must be a valid International Standard Book Number (ISBN) in ISBN-10 format",
			code:    ErrorCodeStringISBN10,
		},
		"isbn-13": {
			rule:    StringISBN13(),
			message: "string must be a valid International Standard Book Number (ISBN) in ISBN-13 format",
			code:    ErrorCodeStringISBN13,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := test.rule.Validate(input)
			assert.EqualError(t, err, test.message)
			assert.True(t, govy.HasErrorCode(err, test.code))
		})
	}
}

func BenchmarkStringISBN(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISBN(),
		validISBNTestCases,
		invalidISBNTestCases,
	)
}

func BenchmarkStringISBNVeryLargeInvalid(b *testing.B) {
	rule := StringISBN()
	input := strings.Repeat("0", 1<<20)

	for b.Loop() {
		_ = rule.Validate(input)
	}
	b.ReportMetric(1, "validations/op")
}

var validISBN10TestCases = map[string]string{
	"hyphenated":                  "0-306-40615-2",
	"plain":                       "0306406152",
	"x check":                     "0-9752298-0-X",
	"spaced":                      "0 9752298 0 x",
	"library converter numeric":   "0394170660",
	"library converter alternate": "0717941728",
	"library converter x check":   "087779443X",
	"MARC hyphenated":             "0-87068-693-3",
}

var invalidISBN10TestCases = map[string]string{
	"empty":                "",
	"failed check":         "0-306-40615-3",
	"x check mutation":     "0877794430",
	"x in body":            "08777X443X",
	"x in fourth position": "087X79443X",
	"short":                "087779443",
	"trailing space":       "087779443X ",
	"isbn 13":              "978-0-306-40615-7",
	"isbn 13 plain":        "9780306406157",
	"repeated separator":   "0-306--40615-2",
}

func TestStringISBN10(t *testing.T) {
	rule := StringISBN10()
	t.Run("valid inputs", func(t *testing.T) {
		for name, input := range validISBN10TestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for name, input := range invalidISBN10TestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(
					t,
					err,
					"string must be a valid International Standard Book Number (ISBN) in ISBN-10 format",
				)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN10))
			})
		}
	})
}

func BenchmarkStringISBN10(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISBN10(),
		validISBN10TestCases,
		invalidISBN10TestCases,
	)
}

var validISBN13TestCases = map[string]string{
	"hyphenated":               "978-0-306-40615-7",
	"plain":                    "9780306406157",
	"grouped":                  "978-3-16-148410-0",
	"agency manual hyphenated": "978-92-95055-12-4",
	"agency manual spaced":     "978 92 95055 12 4",
	"agency manual compact":    "9789295055124",
	"agency manual hardback":   "978-951-45-9693-3",
	"agency manual paperback":  "978-951-45-9694-0",
	"agency manual PDF":        "978-951-45-9695-7",
	"agency manual EPUB":       "978-951-45-9696-4",
	"library converter first":  "9780060723804",
	"library converter second": "9780060799748",
	"979 prefix":               "979-10-90636-07-1",
}

var invalidISBN13TestCases = map[string]string{
	"empty":                       "",
	"isbn 10":                     "0-306-40615-2",
	"failed check":                "978-0-306-40615-8",
	"manual check mutation":       "978-92-95055-12-5",
	"checksum valid wrong prefix": "9779295055125",
	"x in body":                   "978-92-X5055-12-4",
	"x check":                     "978-92-95055-12-X",
	"en dash separators":          "978–92–95055–12–4",
	"full width digits":           "９７８９２９５０５５１２４",
	"display prefix":              "ISBN 978-92-95055-12-4",
	"invalid prefix":              "9770306406157",
	"trailing space":              "978 0 306 40615 7 ",
}

func TestStringISBN13(t *testing.T) {
	rule := StringISBN13()
	t.Run("valid inputs", func(t *testing.T) {
		for name, input := range validISBN13TestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for name, input := range invalidISBN13TestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(
					t,
					err,
					"string must be a valid International Standard Book Number (ISBN) in ISBN-13 format",
				)
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISBN13))
			})
		}
	})
}

func TestISBNPredicatesMatchReference(t *testing.T) {
	tests := map[string]struct {
		predicate func(string) bool
		reference func(string) bool
		inputs    []map[string]string
	}{
		"isbn": {
			predicate: isISBN,
			reference: referenceISBN,
			inputs: []map[string]string{
				validISBNTestCases,
				invalidISBNTestCases,
			},
		},
		"isbn-10": {
			predicate: isISBN10,
			reference: referenceISBN10,
			inputs: []map[string]string{
				validISBN10TestCases,
				invalidISBN10TestCases,
			},
		},
		"isbn-13": {
			predicate: isISBN13,
			reference: referenceISBN13,
			inputs: []map[string]string{
				validISBN13TestCases,
				invalidISBN13TestCases,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			testStringPredicateMatchesReference(t, test.predicate, test.reference, test.inputs...)
		})
	}
}

func BenchmarkStringISBN13(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISBN13(),
		validISBN13TestCases,
		invalidISBN13TestCases,
	)
}

var validISSNTestCases = map[string]string{
	"numeric check":       "2049-3630",
	"numeric example":     "0378-5955",
	"uppercase x":         "2434-561X",
	"lowercase x":         "2434-561x",
	"manual numeric":      "1106-1111",
	"manual uppercase x":  "1092-003X",
	"library check digit": "0317-8471",
	"numeric 2162":        "2162-3546",
	"numeric 1548":        "1548-7180",
	"uppercase x 1204":    "1204-539X",
}

// invalidISSNTestCases includes exact compact construction examples from the
// [ISSN Manual, May 2025] because StringISSN requires the ASCII-hyphenated form.
// Their derived hyphenated forms are accepted in validISSNTestCases.
//
// [ISSN Manual, May 2025]: https://www.issn.org/wp-content/uploads/2025/05/Manual-ISSN_ENG-marc21_May2025.pdf
var invalidISSNTestCases = map[string]string{
	"empty":                      "",
	"missing hyphen":             "20493630",
	"manual compact 2162":        "21623546",
	"manual compact 1548":        "15487180",
	"failed check":               "2049-3631",
	"numeric check mutation":     "1106-1112",
	"uppercase x check mutation": "1092-0030",
	"wrong grouping":             "204-93630",
	"x before check":             "2049-36X0",
	"hyphen as check":            "2049-363-",
	"unicode hyphen":             "1106–1111",
	"U+2010 hyphen":              "1092‐003X",
	"space separator":            "1106 1111",
	"display prefix":             "ISSN 1106-1111",
	"trailing newline":           "1106-1111\n",
	"full width digits":          "１１０６-１１１１",
}

func TestStringISSN(t *testing.T) {
	rule := StringISSN()
	t.Run("valid inputs", func(t *testing.T) {
		for name, input := range validISSNTestCases {
			t.Run(name, func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for name, input := range invalidISSNTestCases {
			t.Run(name, func(t *testing.T) {
				err := rule.Validate(input)
				assert.EqualError(t, err, "string must be a valid International Standard Serial Number (ISSN)")
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISSN))
			})
		}
	})
}

func TestISSNPredicateMatchesReference(t *testing.T) {
	format := regexp.MustCompile(`^\d{4}-\d{3}[0-9Xx]$`)
	testStringPredicateMatchesReference(
		t,
		isISSN,
		func(s string) bool {
			return referenceISSN(format, s)
		},
		validISSNTestCases,
		invalidISSNTestCases,
	)
}

func BenchmarkStringISSN(b *testing.B) {
	benchmarkStringPublicationRule(
		b,
		StringISSN(),
		validISSNTestCases,
		invalidISSNTestCases,
	)
}

func Benchmark_isISSN(b *testing.B) {
	for b.Loop() {
		for name, input := range validISSNTestCases {
			if !isISSN(input) {
				b.Fatalf("%s: expected valid ISSN", name)
			}
		}
		for name, input := range invalidISSNTestCases {
			if isISSN(input) {
				b.Fatalf("%s: expected invalid ISSN", name)
			}
		}
	}
	b.ReportMetric(float64(len(validISSNTestCases)+len(invalidISSNTestCases)), "validations/op")
}

func benchmarkStringPublicationRule(
	b *testing.B,
	rule govy.Rule[string],
	validInputs map[string]string,
	invalidInputs map[string]string,
) {
	b.Helper()
	for b.Loop() {
		for _, in := range validInputs {
			_ = rule.Validate(in)
		}
		for _, in := range invalidInputs {
			_ = rule.Validate(in)
		}
	}
	b.ReportMetric(float64(len(validInputs)+len(invalidInputs)), "validations/op")
}

type stringPredicate func(string) bool

func testStringPredicateMatchesReference(
	t *testing.T,
	predicate stringPredicate,
	reference stringPredicate,
	inputs ...map[string]string,
) {
	t.Helper()
	for _, inputSet := range inputs {
		for name, input := range inputSet {
			assertStringPredicateMatchesReference(t, predicate, reference, name, input)
			testStringPredicateByteEdits(t, predicate, reference, input)
		}
	}
}

func testStringPredicateByteEdits(
	t *testing.T,
	predicate stringPredicate,
	reference stringPredicate,
	input string,
) {
	t.Helper()
	for position := range len(input) {
		candidate := input[:position] + input[position+1:]
		assertStringPredicateMatchesReference(t, predicate, reference, "byte deletion", candidate)
	}

	replacement := []byte(input)
	for position := range len(replacement) {
		original := replacement[position]
		for value := range 256 {
			replacement[position] = byte(value)
			assertStringPredicateMatchesReference(
				t,
				predicate,
				reference,
				"byte replacement",
				string(replacement),
			)
		}
		replacement[position] = original
	}

	for position := range len(input) + 1 {
		insertion := make([]byte, len(input)+1)
		copy(insertion, input[:position])
		copy(insertion[position+1:], input[position:])
		for value := range 256 {
			insertion[position] = byte(value)
			assertStringPredicateMatchesReference(
				t,
				predicate,
				reference,
				"byte insertion",
				string(insertion),
			)
		}
	}
}

func assertStringPredicateMatchesReference(
	t *testing.T,
	predicate stringPredicate,
	reference stringPredicate,
	name string,
	input string,
) {
	t.Helper()
	got := predicate(input)
	expected := reference(input)
	if got != expected {
		t.Fatalf("%s for %q: got %t, expected %t", name, input, got, expected)
	}
}

func referenceISBN(s string) bool {
	return referenceISBN10(s) || referenceISBN13(s)
}

func referenceISBN10(s string) bool {
	isbn, ok := referenceNormalizeISBN(s)
	if !ok || len(isbn) != isbn10Length {
		return false
	}

	sum := 0
	for i := range isbn10Length - 1 {
		if !isASCIIDigit(isbn[i]) {
			return false
		}
		sum += int(isbn[i]-'0') * (isbn10Length - i)
	}

	switch checkDigit := isbn[isbn10Length-1]; {
	case isASCIIDigit(checkDigit):
		sum += int(checkDigit - '0')
	case checkDigit == 'X' || checkDigit == 'x':
		sum += 10
	default:
		return false
	}
	return sum%11 == 0
}

func referenceISBN13(s string) bool {
	isbn, ok := referenceNormalizeISBN(s)
	if !ok || len(isbn) != isbn13Length ||
		(!strings.HasPrefix(isbn, "978") && !strings.HasPrefix(isbn, "979")) {
		return false
	}

	sum := 0
	for i := range isbn13Length - 1 {
		if !isASCIIDigit(isbn[i]) {
			return false
		}
		weight := 1
		if i%2 != 0 {
			weight = 3
		}
		sum += int(isbn[i]-'0') * weight
	}
	if !isASCIIDigit(isbn[isbn13Length-1]) {
		return false
	}
	return (10-sum%10)%10 == int(isbn[isbn13Length-1]-'0')
}

func referenceNormalizeISBN(s string) (string, bool) {
	if s == "" {
		return "", false
	}

	var builder strings.Builder
	previousWasSeparator := false
	for i := range len(s) {
		switch c := s[i]; {
		case isASCIIDigit(c) || c == 'X' || c == 'x':
			builder.WriteByte(c)
			previousWasSeparator = false
		case c == '-' || c == ' ':
			if i == 0 || i == len(s)-1 || previousWasSeparator {
				return "", false
			}
			previousWasSeparator = true
		default:
			return "", false
		}
	}
	return builder.String(), true
}

func referenceISSN(format *regexp.Regexp, s string) bool {
	if !format.MatchString(s) {
		return false
	}

	issn := strings.ReplaceAll(s, "-", "")
	sum := 0
	for i := range 7 {
		if !isASCIIDigit(issn[i]) {
			return false
		}
		sum += int(issn[i]-'0') * (8 - i)
	}

	switch checkDigit := issn[7]; {
	case isASCIIDigit(checkDigit):
		sum += int(checkDigit - '0')
	case checkDigit == 'X' || checkDigit == 'x':
		sum += 10
	default:
		return false
	}
	return sum%11 == 0
}
