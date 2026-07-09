package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

const (
	stringBCP47LanguageTagError       = "string must be a valid BCP 47 language tag (e.g. 'en', 'en-US', 'zh-Hant-TW')"
	stringBCP47StrictLanguageTagError = "string must be a valid canonical BCP 47 language tag (e.g. 'en', 'en-US', 'zh-Hant-TW')"
	stringISO3166Alpha2Error          = "string must be a valid ISO 3166-1 alpha-2 country code (e.g. 'US', 'PL', 'JP')"
	stringISO3166Alpha3Error          = "string must be a valid ISO 3166-1 alpha-3 country code (e.g. 'USA', 'POL', 'JPN')"
	stringISO3166NumericError         = "string must be a valid ISO 3166-1 numeric country code (e.g. '840', '616', '392')"
	stringISO31662Error               = "string must be a valid ISO 3166-2 subdivision code (e.g. 'US-CA', 'GB-ENG', 'PL-14')"
	stringISO4217Error                = "string must be a valid ISO 4217 currency code (e.g. 'USD', 'EUR', 'JPY')"
	stringLatitudeError               = "string must be a valid latitude coordinate (e.g. '0', '-45.25', '90')"
	stringLongitudeError              = "string must be a valid longitude coordinate (e.g. '0', '-122.4194', '180')"
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

var stringBCP47LanguageTagTestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "en"},
	{in: "en-US"},
	{in: "zh-Hant-TW"},
	{in: "iw"},
	{in: "en_GB", expectedError: stringBCP47LanguageTagError},
	{in: "English", expectedError: stringBCP47LanguageTagError},
	{in: "", expectedError: stringBCP47LanguageTagError},
}

func TestStringBCP47LanguageTag(t *testing.T) {
	for _, tc := range stringBCP47LanguageTagTestCases {
		err := StringBCP47LanguageTag().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBCP47LanguageTag))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringBCP47LanguageTag(b *testing.B) {
	rule := StringBCP47LanguageTag()
	for b.Loop() {
		_ = rule.Validate("zh-Hant-TW")
	}
}

var stringBCP47StrictLanguageTagTestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "en"},
	{in: "en-US"},
	{in: "zh-Hant-TW"},
	{in: "iw", expectedError: stringBCP47StrictLanguageTagError},
	{in: "EN-us", expectedError: stringBCP47StrictLanguageTagError},
	{in: "en-Latn", expectedError: stringBCP47StrictLanguageTagError},
	{in: "en_GB", expectedError: stringBCP47StrictLanguageTagError},
	{in: "English", expectedError: stringBCP47StrictLanguageTagError},
}

func TestStringBCP47StrictLanguageTag(t *testing.T) {
	for _, tc := range stringBCP47StrictLanguageTagTestCases {
		err := StringBCP47StrictLanguageTag().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBCP47StrictLanguageTag))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringBCP47StrictLanguageTag(b *testing.B) {
	rule := StringBCP47StrictLanguageTag()
	for b.Loop() {
		_ = rule.Validate("zh-Hant-TW")
	}
}

var stringISO3166Alpha2TestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "US"},
	{in: "PL"},
	{in: "JP"},
	{in: "001", expectedError: stringISO3166Alpha2Error},
	{in: "us", expectedError: stringISO3166Alpha2Error},
	{in: "ZZ", expectedError: stringISO3166Alpha2Error},
	{in: "UK", expectedError: stringISO3166Alpha2Error},
	{in: "SU", expectedError: stringISO3166Alpha2Error},
	{in: "AN", expectedError: stringISO3166Alpha2Error},
	{in: "XK", expectedError: stringISO3166Alpha2Error},
	{in: "AC", expectedError: stringISO3166Alpha2Error},
}

func TestStringISO3166Alpha2(t *testing.T) {
	for _, tc := range stringISO3166Alpha2TestCases {
		err := StringISO3166Alpha2().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO3166Alpha2))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISO3166Alpha2(b *testing.B) {
	rule := StringISO3166Alpha2()
	for b.Loop() {
		_ = rule.Validate("US")
	}
}

var stringISO3166Alpha3TestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "USA"},
	{in: "POL"},
	{in: "JPN"},
	{in: "001", expectedError: stringISO3166Alpha3Error},
	{in: "usa", expectedError: stringISO3166Alpha3Error},
	{in: "ZZZ", expectedError: stringISO3166Alpha3Error},
	{in: "SUN", expectedError: stringISO3166Alpha3Error},
	{in: "ANT", expectedError: stringISO3166Alpha3Error},
	{in: "DDR", expectedError: stringISO3166Alpha3Error},
	{in: "XKK", expectedError: stringISO3166Alpha3Error},
	{in: "ASC", expectedError: stringISO3166Alpha3Error},
}

func TestStringISO3166Alpha3(t *testing.T) {
	for _, tc := range stringISO3166Alpha3TestCases {
		err := StringISO3166Alpha3().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO3166Alpha3))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISO3166Alpha3(b *testing.B) {
	rule := StringISO3166Alpha3()
	for b.Loop() {
		_ = rule.Validate("USA")
	}
}

var stringISO3166NumericTestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "840"},
	{in: "616"},
	{in: "392"},
	{in: "001", expectedError: stringISO3166NumericError},
	{in: "84", expectedError: stringISO3166NumericError},
	{in: "USA", expectedError: stringISO3166NumericError},
	{in: "810", expectedError: stringISO3166NumericError},
	{in: "530", expectedError: stringISO3166NumericError},
	{in: "278", expectedError: stringISO3166NumericError},
	{in: "983", expectedError: stringISO3166NumericError},
}

func TestStringISO3166Numeric(t *testing.T) {
	for _, tc := range stringISO3166NumericTestCases {
		err := StringISO3166Numeric().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO3166Numeric))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISO3166Numeric(b *testing.B) {
	rule := StringISO3166Numeric()
	for b.Loop() {
		_ = rule.Validate("840")
	}
}

var stringISO31662TestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "US-CA"},
	{in: "GB-ENG"},
	{in: "PL-14"},
	{in: "US-XXX", expectedError: stringISO31662Error},
	{in: "FR-999", expectedError: stringISO31662Error},
	{in: "ZZ-CA", expectedError: stringISO31662Error},
	{in: "US-cal", expectedError: stringISO31662Error},
	{in: "USA-CA", expectedError: stringISO31662Error},
	{in: "UK-ENG", expectedError: stringISO31662Error},
	{in: "SU-MOW", expectedError: stringISO31662Error},
	{in: "AN-CW", expectedError: stringISO31662Error},
	{in: "XK-01", expectedError: stringISO31662Error},
	{in: "AC-SH", expectedError: stringISO31662Error},
}

func TestStringISO31662(t *testing.T) {
	for _, tc := range stringISO31662TestCases {
		err := StringISO31662().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO31662))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISO31662(b *testing.B) {
	rule := StringISO31662()
	for b.Loop() {
		_ = rule.Validate("US-CA")
	}
}

var stringISO4217TestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "USD"},
	{in: "EUR"},
	{in: "JPY"},
	{in: "XXX"},
	{in: "XTS"},
	{in: "XAU"},
	{in: "usd", expectedError: stringISO4217Error},
	{in: "ZZZ", expectedError: stringISO4217Error},
	{in: "US", expectedError: stringISO4217Error},
	{in: "ADP", expectedError: stringISO4217Error},
	{in: "EEK", expectedError: stringISO4217Error},
	{in: "ZWD", expectedError: stringISO4217Error},
	{in: "RUR", expectedError: stringISO4217Error},
	{in: "BYR", expectedError: stringISO4217Error},
}

func TestStringISO4217(t *testing.T) {
	for _, tc := range stringISO4217TestCases {
		err := StringISO4217().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringISO4217))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringISO4217(b *testing.B) {
	rule := StringISO4217()
	for b.Loop() {
		_ = rule.Validate("USD")
	}
}

var stringLatitudeTestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "0"},
	{in: "-45.25"},
	{in: "+90"},
	{in: ".5"},
	{in: "90.1", expectedError: stringLatitudeError},
	{in: "north", expectedError: stringLatitudeError},
	{in: "1e2", expectedError: stringLatitudeError},
}

func TestStringLatitude(t *testing.T) {
	for _, tc := range stringLatitudeTestCases {
		err := StringLatitude().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringLatitude))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringLatitude(b *testing.B) {
	rule := StringLatitude()
	for b.Loop() {
		_ = rule.Validate("-45.25")
	}
}

var stringLongitudeTestCases = []*struct {
	in            string
	expectedError string
}{
	{in: "0"},
	{in: "-122.4194"},
	{in: "+180"},
	{in: ".5"},
	{in: "180.1", expectedError: stringLongitudeError},
	{in: "east", expectedError: stringLongitudeError},
	{in: "1e2", expectedError: stringLongitudeError},
}

func TestStringLongitude(t *testing.T) {
	for _, tc := range stringLongitudeTestCases {
		err := StringLongitude().Validate(tc.in)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeStringLongitude))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkStringLongitude(b *testing.B) {
	rule := StringLongitude()
	for b.Loop() {
		_ = rule.Validate("-122.4194")
	}
}
