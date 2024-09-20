package rules

import (
	"regexp"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestStringNotEmpty(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringNotEmpty().Validate("                s")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringNotEmpty().Validate("     ")
		assert.Error(t, err)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringNotEmpty))
	})
}

func TestStringMatchRegexp(t *testing.T) {
	re := regexp.MustCompile("[ab]+")
	t.Run("passes", func(t *testing.T) {
		err := StringMatchRegexp(re).Validate("ab")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringMatchRegexp(re).Validate("cd")
		assert.EqualError(t, err, "string must match regular expression: '[ab]+'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMatchRegexp))
	})
	t.Run("examples output", func(t *testing.T) {
		err := StringMatchRegexp(re, "ab", "a", "b").Validate("cd")
		assert.EqualError(t, err, "string must match regular expression: '[ab]+' (e.g. 'ab', 'a', 'b')")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMatchRegexp))
	})
}

func TestStringDenyRegexp(t *testing.T) {
	re := regexp.MustCompile("[ab]+")
	t.Run("passes", func(t *testing.T) {
		err := StringDenyRegexp(re).Validate("cd")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringDenyRegexp(re).Validate("ab")
		assert.EqualError(t, err, "string must not match regular expression: '[ab]+'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDenyRegexp))
	})
	t.Run("examples output", func(t *testing.T) {
		err := StringDenyRegexp(re, "ab", "a", "b").Validate("ab")
		assert.EqualError(t, err, "string must not match regular expression: '[ab]+' (e.g. 'ab', 'a', 'b')")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringDenyRegexp))
	})
}

func TestStringDNSLabel(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringDNSLabel().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			for _, e := range err.(govy.RuleSetError) {
				assert.True(t, govy.HasErrorCode(e, ErrorCodeStringDNSLabel))
			}
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringASCII(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringASCII().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringASCII))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringUUID(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringUUID().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringUUID))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringEmail(t *testing.T) {
	tests := []struct {
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
		// Custom domains are allowed by RFC 5322.
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
	for _, tc := range tests {
		err := StringEmail().Validate(tc.in)
		if tc.shouldFail {
			assert.ErrorContains(t, err, "string must be a valid email address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEmail))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringURL(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for _, input := range validURLs {
			err := StringURL().Validate(input)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		for _, input := range invalidURLs {
			err := StringURL().Validate(input)
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringURL))
		}
	})
}

func TestStringMAC(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringMAC().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid MAC address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringMAC))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringIP(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringIP().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IP address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIP))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringIPv4(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringIPv4().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IPv4 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIPv4))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringIPv6(t *testing.T) {
	tests := []struct {
		in         string
		shouldFail bool
	}{
		// cspell:disable
		{"10.0.0.1", true},
		{"172.16.0.1", true},
		{"192.168.0.1", true},
		{"192.168.255.254", true},
		{"192.168.255.256", true},
		{"172.16.255.254", true},
		{"172.16.256.255", true},
		{"2001:cdba:0000:0000:0000:0000:3257:9652", false},
		{"2001:cdba:0:0:0:0:3257:9652", false},
		{"2001:cdba::3257:9652", false},
		// cspell:enable
	}
	for _, tc := range tests {
		err := StringIPv6().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid IPv6 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringIPv6))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringCIDR(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringCIDR().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IP address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDR))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringCIDRv4(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringCIDRv4().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IPv4 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDRv4))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringCIDRv6(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringCIDRv6().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "string must be a valid CIDR notation IPv6 address")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringCIDRv6))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestStringJSON(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringJSON().Validate(`{"foo": "bar"}`)
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringJSON().Validate(`{]}`)
		assert.Error(t, err)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringJSON))
	})
}

func TestStringContains(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringContains("th", "is").Validate("this")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringContains("th", "ht").Validate("one")
		assert.EqualError(t, err, "string must contain the following substrings: 'th', 'ht'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringContains))
	})
}

func TestStringExcludes(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringExcludes("ts", "ih").Validate("this")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringExcludes("oe", "ne").Validate("one")
		assert.EqualError(t, err, "string must not contain any of the following substrings: 'oe', 'ne'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringExcludes))
	})
}

func TestStringStartsWith(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for _, prefixes := range [][]string{
			{"th"},
			{"is", "th"},
		} {
			err := StringStartsWith(prefixes...).Validate("this")
			assert.NoError(t, err)
		}
	})
	t.Run("fails with single prefix", func(t *testing.T) {
		err := StringStartsWith("th").Validate("one")
		assert.EqualError(t, err, "string must start with 'th' prefix")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringStartsWith))
	})
	t.Run("fails with multiple prefixes", func(t *testing.T) {
		err := StringStartsWith("th", "ht").Validate("one")
		assert.EqualError(t, err, "string must start with one of the following prefixes: 'th', 'ht'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringStartsWith))
	})
}

func TestStringEndsWith(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for _, prefixes := range [][]string{
			{"is"},
			{"th", "is"},
		} {
			err := StringEndsWith(prefixes...).Validate("this")
			assert.NoError(t, err)
		}
	})
	t.Run("fails with single suffix", func(t *testing.T) {
		err := StringEndsWith("th").Validate("one")
		assert.EqualError(t, err, "string must end with 'th' suffix")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEndsWith))
	})
	t.Run("fails with multiple suffixes", func(t *testing.T) {
		err := StringEndsWith("th", "ht").Validate("one")
		assert.EqualError(t, err, "string must end with one of the following suffixes: 'th', 'ht'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEndsWith))
	})
}

func TestStringTitle(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringTitle().Validate(tc.in)
		if tc.shouldFail {
			assert.EqualError(t, err, "each word in a string must start with a capital letter")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringTitle))
		} else {
			assert.NoError(t, err)
		}
	}
}
