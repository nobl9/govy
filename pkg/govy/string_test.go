package validation

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringNotEmpty(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringNotEmpty().Validate("                s")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringNotEmpty().Validate("     ")
		assert.Error(t, err)
		assert.True(t, HasErrorCode(err, ErrorCodeStringNotEmpty))
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
		assert.True(t, HasErrorCode(err, ErrorCodeStringMatchRegexp))
	})
	t.Run("examples output", func(t *testing.T) {
		err := StringMatchRegexp(re, "ab", "a", "b").Validate("cd")
		assert.EqualError(t, err, "string must match regular expression: '[ab]+' (e.g. 'ab', 'a', 'b')")
		assert.True(t, HasErrorCode(err, ErrorCodeStringMatchRegexp))
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
		assert.True(t, HasErrorCode(err, ErrorCodeStringDenyRegexp))
	})
	t.Run("examples output", func(t *testing.T) {
		err := StringDenyRegexp(re, "ab", "a", "b").Validate("ab")
		assert.EqualError(t, err, "string must not match regular expression: '[ab]+' (e.g. 'ab', 'a', 'b')")
		assert.True(t, HasErrorCode(err, ErrorCodeStringDenyRegexp))
	})
}

func TestStringIsDNSSubdomain(t *testing.T) {
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
			err := StringIsDNSSubdomain().Validate(input)
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
			err := StringIsDNSSubdomain().Validate(input)
			assert.Error(t, err)
			for _, e := range err.(ruleSetError) {
				assert.True(t, HasErrorCode(e, ErrorCodeStringIsDNSSubdomain))
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
			assert.True(t, HasErrorCode(err, ErrorCodeStringASCII))
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
			assert.True(t, HasErrorCode(err, ErrorCodeStringUUID))
		}
	})
}

func TestStringDescription(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := StringDescription().Validate(strings.Repeat("l", 1050))
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := StringDescription().Validate(strings.Repeat("l", 1051))
		assert.Error(t, err)
		assert.True(t, HasErrorCode(err, ErrorCodeStringDescription))
	})
}

func TestStringIsURL(t *testing.T) {
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
			assert.True(t, HasErrorCode(err, ErrorCodeStringURL))
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
		assert.True(t, HasErrorCode(err, ErrorCodeStringJSON))
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
		assert.True(t, HasErrorCode(err, ErrorCodeStringContains))
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
		assert.True(t, HasErrorCode(err, ErrorCodeStringStartsWith))
	})
	t.Run("fails with multiple prefixes", func(t *testing.T) {
		err := StringStartsWith("th", "ht").Validate("one")
		assert.Error(t, err)
		assert.EqualError(t, err, "string must start with one of the following prefixes: 'th', 'ht'")
		assert.True(t, HasErrorCode(err, ErrorCodeStringStartsWith))
	})
}
