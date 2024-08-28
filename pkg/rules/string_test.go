package rules

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

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
	t.Run("passes", func(t *testing.T) {
		for _, input := range []string{
			"test",
			"s",
			"test-this",
			"test-1-this",
			"test1-this",
			"123",
			strings.Repeat("l", 63),
		} {
			err := StringDNSLabel().Validate(input)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		for _, input := range []string{
			"tesT",
			"",
			strings.Repeat("l", 64),
			"test?",
			"test this",
			"1_2",
			"LOL",
		} {
			err := StringDNSLabel().Validate(input)
			assert.Error(t, err)
			for _, e := range err.(govy.RuleSetError) {
				assert.True(t, govy.HasErrorCode(e, ErrorCodeStringDNSLabel))
			}
		}
	})
}

func TestStringASCII(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for _, input := range []string{
			"foobar",
			"0987654321",
			"test@example.com",
			"1234abcDEF",
			"",
		} {
			err := StringASCII().Validate(input)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		for _, input := range []string{
			// cspell:disable
			"ｆｏｏbar",
			"ｘｙｚ０９８",
			"１２３456",
			"ｶﾀｶﾅ",
			// cspell:enable
		} {
			err := StringASCII().Validate(input)
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringASCII))
		}
	})
}

func TestStringUUID(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		for _, input := range []string{
			"00000000-0000-0000-0000-000000000000",
			"e190c630-8873-11ee-b9d1-0242ac120002",
			"79258D24-01A7-47E5-ACBB-7E762DE52298",
		} {
			err := StringUUID().Validate(input)
			assert.NoError(t, err)
		}
	})
	t.Run("fails", func(t *testing.T) {
		for _, input := range []string{
			// cspell:disable
			"foobar",
			"0987654321",
			"AXAXAXAX-AAAA-AAAA-AAAA-AAAAAAAAAAAA",
			"00000000-0000-0000-0000-0000000000",
			// cspell:enable
		} {
			err := StringUUID().Validate(input)
			assert.Error(t, err)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringUUID))
		}
	})
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
		assert.Error(t, err)
		assert.EqualError(t, err, "string must contain the following substrings: 'th', 'ht'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringContains))
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
		assert.Error(t, err)
		assert.EqualError(t, err, "string must start with 'th' prefix")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringStartsWith))
	})
	t.Run("fails with multiple prefixes", func(t *testing.T) {
		err := StringStartsWith("th", "ht").Validate("one")
		assert.Error(t, err)
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
		assert.Error(t, err)
		assert.EqualError(t, err, "string must end with 'th' suffix")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEndsWith))
	})
	t.Run("fails with multiple suffixes", func(t *testing.T) {
		err := StringEndsWith("th", "ht").Validate("one")
		assert.Error(t, err)
		assert.EqualError(t, err, "string must end with one of the following suffixes: 'th', 'ht'")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringEndsWith))
	})
}

func TestStringTitle(t *testing.T) {
	tests := []struct {
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
	for _, tc := range tests {
		err := StringTitle().Validate(tc.in)
		if tc.shouldFail {
			assert.Error(t, err, "input: %q", tc.in)
			assert.EqualError(t, err, "each word in a string must start with a capital letter")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringTitle))
		} else {
			assert.NoError(t, err, "input: %q", tc.in)
		}
	}
}
