package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

func TestLazyLookupMap(t *testing.T) {
	calls := 0
	lazyMap := lazyLookupMap(func() map[string]struct{} {
		calls++
		return map[string]struct{}{
			"test": {},
		}
	})

	assert.Equal(t, 0, calls)

	lookup := lazyMap()
	_, ok := lookup["test"]
	assert.True(t, ok)
	assert.Equal(t, 1, calls)

	_ = lazyMap()
	assert.Equal(t, 1, calls)
}

var validStringBCP47LanguageTagInputs = []string{
	"en",
	"en-US",
	"zh-Hant-TW",
	"iw",
}

var invalidStringBCP47LanguageTagInputs = []string{
	"en_GB",
	"English",
	"",
}

func TestStringBCP47LanguageTag(t *testing.T) {
	rule := StringBCP47LanguageTag()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringBCP47LanguageTagInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringBCP47LanguageTagInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringBCP47LanguageTagInputs[0])
		assert.EqualError(
			t,
			err,
			"string must be a valid BCP 47 language tag (e.g. 'en', 'en-US', 'zh-Hant-TW')",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBCP47LanguageTag))
	})
}

func BenchmarkStringBCP47LanguageTag(b *testing.B) {
	rule := StringBCP47LanguageTag()
	for b.Loop() {
		_ = rule.Validate("zh-Hant-TW")
	}
}

var validStringBCP47StrictLanguageTagInputs = []string{
	"en",
	"en-US",
	"zh-Hant-TW",
}

var invalidStringBCP47StrictLanguageTagInputs = []string{
	"iw",
	"EN-us",
	"en-Latn",
	"en_GB",
	"English",
}

func TestStringBCP47StrictLanguageTag(t *testing.T) {
	rule := StringBCP47StrictLanguageTag()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringBCP47StrictLanguageTagInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringBCP47StrictLanguageTagInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringBCP47StrictLanguageTagInputs[0])
		assert.EqualError(
			t,
			err,
			"string must be a valid canonical BCP 47 language tag (e.g. 'en', 'en-US', 'zh-Hant-TW')",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBCP47StrictLanguageTag))
	})
}

func BenchmarkStringBCP47StrictLanguageTag(b *testing.B) {
	rule := StringBCP47StrictLanguageTag()
	for b.Loop() {
		_ = rule.Validate("zh-Hant-TW")
	}
}

var validStringISO3166Alpha2Inputs = []string{
	"US",
	"PL",
	"JP",
}

var invalidStringISO3166Alpha2Inputs = []string{
	"001",
	"us",
	"ZZ",
	"UK",
	"SU",
	"AN",
	"XK",
	"AC",
}

func TestStringISO3166Alpha2(t *testing.T) {
	rule := StringISO3166Alpha2()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringISO3166Alpha2Inputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringISO3166Alpha2Inputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringISO3166Alpha2Inputs[0])
		assert.EqualError(
			t,
			err,
			"string must be a valid ISO 3166-1 alpha-2 country code (e.g. 'US', 'PL', 'JP')",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO3166Alpha2))
	})
}

func BenchmarkStringISO3166Alpha2(b *testing.B) {
	rule := StringISO3166Alpha2()
	for b.Loop() {
		_ = rule.Validate("US")
	}
}

var validStringISO3166Alpha3Inputs = []string{
	"USA",
	"POL",
	"JPN",
}

var invalidStringISO3166Alpha3Inputs = []string{
	"001",
	"usa",
	"ZZZ",
	"SUN",
	"ANT",
	"DDR",
	"XKK",
	"ASC",
}

func TestStringISO3166Alpha3(t *testing.T) {
	rule := StringISO3166Alpha3()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringISO3166Alpha3Inputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringISO3166Alpha3Inputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringISO3166Alpha3Inputs[0])
		assert.EqualError(
			t,
			err,
			"string must be a valid ISO 3166-1 alpha-3 country code (e.g. 'USA', 'POL', 'JPN')",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO3166Alpha3))
	})
}

func BenchmarkStringISO3166Alpha3(b *testing.B) {
	rule := StringISO3166Alpha3()
	for b.Loop() {
		_ = rule.Validate("USA")
	}
}

var validStringISO3166NumericInputs = []string{
	"840",
	"616",
	"392",
}

var invalidStringISO3166NumericInputs = []string{
	"001",
	"84",
	"USA",
	"810",
	"530",
	"278",
	"983",
}

func TestStringISO3166Numeric(t *testing.T) {
	rule := StringISO3166Numeric()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringISO3166NumericInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringISO3166NumericInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringISO3166NumericInputs[0])
		assert.EqualError(
			t,
			err,
			"string must be a valid ISO 3166-1 numeric-3 country code (e.g. '840', '616', '392')",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO3166Numeric))
	})
}

func BenchmarkStringISO3166Numeric(b *testing.B) {
	rule := StringISO3166Numeric()
	for b.Loop() {
		_ = rule.Validate("840")
	}
}

var validStringISO31662Inputs = []string{
	"US-CA",
	"GB-ENG",
	"PL-14",
}

var invalidStringISO31662Inputs = []string{
	"US-XXX",
	"FR-999",
	"ZZ-CA",
	"US-cal",
	"USA-CA",
	"UK-ENG",
	"SU-MOW",
	"AN-CW",
	"XK-01",
	"AC-SH",
}

func TestStringISO31662(t *testing.T) {
	rule := StringISO31662()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringISO31662Inputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringISO31662Inputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringISO31662Inputs[0])
		assert.EqualError(
			t,
			err,
			"string must be a valid ISO 3166-2 country subdivision code (e.g. 'US-CA', 'GB-ENG', 'PL-14')",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO31662))
	})
}

func BenchmarkStringISO31662(b *testing.B) {
	rule := StringISO31662()
	for b.Loop() {
		_ = rule.Validate("US-CA")
	}
}

var validStringISO4217Inputs = []string{
	"USD",
	"EUR",
	"JPY",
	"XXX",
	"XTS",
	"XAU",
}

var invalidStringISO4217Inputs = []string{
	"usd",
	"ZZZ",
	"US",
	"ADP",
	"EEK",
	"ZWD",
	"RUR",
	"BYR",
}

func TestStringISO4217(t *testing.T) {
	rule := StringISO4217()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringISO4217Inputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringISO4217Inputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringISO4217Inputs[0])
		assert.EqualError(
			t,
			err,
			"string must be a valid ISO 4217 three-letter alphabetic currency code (e.g. 'USD', 'EUR', 'JPY')",
		)
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO4217))
	})
}

func BenchmarkStringISO4217(b *testing.B) {
	rule := StringISO4217()
	for b.Loop() {
		_ = rule.Validate("USD")
	}
}

var validStringLatitudeInputs = []string{
	"0",
	"-45.25",
	"+90",
	".5",
}

var invalidStringLatitudeInputs = []string{
	"90.1",
	"north",
	"1e2",
}

func TestStringLatitude(t *testing.T) {
	rule := StringLatitude()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringLatitudeInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringLatitudeInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringLatitudeInputs[0])
		assert.EqualError(t, err, "string must be a valid latitude coordinate (e.g. '0', '-45.25', '90')")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringLatitude))
	})
}

func BenchmarkStringLatitude(b *testing.B) {
	rule := StringLatitude()
	for b.Loop() {
		_ = rule.Validate("-45.25")
	}
}

var validStringLongitudeInputs = []string{
	"0",
	"-122.4194",
	"+180",
	".5",
}

var invalidStringLongitudeInputs = []string{
	"180.1",
	"east",
	"1e2",
}

func TestStringLongitude(t *testing.T) {
	rule := StringLongitude()
	t.Run("valid inputs", func(t *testing.T) {
		for _, input := range validStringLongitudeInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.NoError(t, rule.Validate(input))
			})
		}
	})
	t.Run("invalid inputs", func(t *testing.T) {
		for _, input := range invalidStringLongitudeInputs {
			t.Run(fmt.Sprintf("%q", input), func(t *testing.T) {
				assert.Error(t, rule.Validate(input))
			})
		}
		err := rule.Validate(invalidStringLongitudeInputs[0])
		assert.EqualError(t, err, "string must be a valid longitude coordinate (e.g. '0', '-122.4194', '180')")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeStringLongitude))
	})
}

func BenchmarkStringLongitude(b *testing.B) {
	rule := StringLongitude()
	for b.Loop() {
		_ = rule.Validate("-122.4194")
	}
}
