package rules

// cspell:ignore biske phonebk Qaaa rozaj USAA

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

type stringRuleCase struct {
	name  string
	input string
}

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

var validStringBCP47LanguageTagCases = []stringRuleCase{
	{name: "language", input: "en"},
	{name: "language and region", input: "en-US"},
	{name: "language script and region", input: "zh-Hant-TW"},
	{name: "deprecated language", input: "iw"},
	{name: "variant", input: "de-CH-1901"},
	{name: "multiple variants", input: "sl-rozaj-biske-1994"},
	{name: "Unicode extension", input: "en-US-u-ca-gregory"},
	{name: "multiple extensions", input: "en-t-ja-u-ca-gregory"},
	{name: "private use suffix", input: "de-CH-x-phonebk"},
	{name: "private use only", input: "x-whatever"},
	{name: "grandfathered regular tag", input: "i-default"},
	{name: "grandfathered irregular tag", input: "en-GB-oed"},
	{name: "grandfathered tag with multiple subtags", input: "zh-min-nan"},
	{name: "redundant tag", input: "zh-cmn-Hans-CN"},
	{name: "private use language script and region", input: "qaa-Qaaa-QM-x-southern"},
}

var invalidStringBCP47LanguageTagCases = []stringRuleCase{
	{name: "underscore separator", input: "en_GB"},
	{name: "non-tag word", input: "English"},
	{name: "empty", input: ""},
	{name: "trailing separator", input: "en-"},
	{name: "leading whitespace", input: " en"},
	{name: "Unicode hyphen", input: "en‐US"},
}

func TestStringBCP47LanguageTag(t *testing.T) {
	rule := StringBCP47LanguageTag()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringBCP47LanguageTagCases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringBCP47LanguageTagCases)
		err := rule.Validate(invalidStringBCP47LanguageTagCases[0].input)
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

var validStringBCP47StrictLanguageTagCases = []stringRuleCase{
	{name: "language", input: "en"},
	{name: "language and region", input: "en-US"},
	{name: "language script and region", input: "zh-Hant-TW"},
	{name: "variant", input: "de-CH-1901"},
	{name: "multiple variants", input: "sl-rozaj-biske-1994"},
	{name: "Unicode extension", input: "en-US-u-ca-gregory"},
	{name: "multiple extensions", input: "en-t-ja-u-ca-gregory"},
	{name: "private use suffix", input: "de-CH-x-phonebk"},
	{name: "private use only", input: "x-whatever"},
	{name: "extended language with script and region", input: "cmn-Hans-CN"},
	{name: "private use language script and region", input: "qaa-Qaaa-QM-x-southern"},
}

var invalidStringBCP47StrictLanguageTagCases = []stringRuleCase{
	{name: "deprecated language", input: "iw"},
	{name: "non-canonical case", input: "EN-us"},
	{name: "suppressed script", input: "en-Latn"},
	{name: "underscore separator", input: "en_GB"},
	{name: "non-tag word", input: "English"},
	{name: "empty", input: ""},
	{name: "deprecated grandfathered tag", input: "en-GB-oed"},
	{name: "deprecated grandfathered tag with multiple subtags", input: "zh-min-nan"},
	{name: "redundant extended language", input: "zh-cmn-Hans-CN"},
	{name: "extension subtag too short", input: "tlh-a-b-foo"},
	{name: "duplicate variant", input: "sl-rozaj-rozaj"},
	{name: "trailing separator", input: "en-"},
	{name: "leading whitespace", input: " en"},
	{name: "Unicode hyphen", input: "en‐US"},
}

func TestStringBCP47StrictLanguageTag(t *testing.T) {
	rule := StringBCP47StrictLanguageTag()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringBCP47StrictLanguageTagCases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringBCP47StrictLanguageTagCases)
		err := rule.Validate(invalidStringBCP47StrictLanguageTagCases[0].input)
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

var validStringISO3166Alpha2Cases = []stringRuleCase{
	{name: "AD AND 020", input: "AD"},
	{name: "AX ALA 248", input: "AX"},
	{name: "BQ BES 535", input: "BQ"},
	{name: "CW CUW 531", input: "CW"},
	{name: "SS SSD 728", input: "SS"},
	{name: "UM UMI 581", input: "UM"},
	{name: "US USA 840", input: "US"},
	{name: "PL POL 616", input: "PL"},
	{name: "JP JPN 392", input: "JP"},
	{name: "TW TWN 158", input: "TW"},
	{name: "ZW ZWE 716", input: "ZW"},
}

var invalidStringISO3166Alpha2Cases = []stringRuleCase{
	{name: "numeric code", input: "001"},
	{name: "lowercase", input: "us"},
	{name: "unassigned", input: "ZZ"},
	{name: "exceptionally reserved", input: "UK"},
	{name: "deleted Soviet Union code", input: "SU"},
	{name: "deleted Netherlands Antilles code", input: "AN"},
	{name: "user-assigned Kosovo code", input: "XK"},
	{name: "exceptionally reserved territory", input: "AC"},
	{name: "too short", input: "U"},
	{name: "too long", input: "USA"},
	{name: "non-letter", input: "U1"},
	{name: "leading whitespace", input: " US"},
	{name: "trailing whitespace", input: "US "},
	{name: "full-width letters", input: "ＵＳ"},
}

func TestStringISO3166Alpha2(t *testing.T) {
	rule := StringISO3166Alpha2()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringISO3166Alpha2Cases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringISO3166Alpha2Cases)
		err := rule.Validate(invalidStringISO3166Alpha2Cases[0].input)
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

var validStringISO3166Alpha3Cases = []stringRuleCase{
	{name: "AD AND 020", input: "AND"},
	{name: "AX ALA 248", input: "ALA"},
	{name: "BQ BES 535", input: "BES"},
	{name: "CW CUW 531", input: "CUW"},
	{name: "SS SSD 728", input: "SSD"},
	{name: "UM UMI 581", input: "UMI"},
	{name: "US USA 840", input: "USA"},
	{name: "PL POL 616", input: "POL"},
	{name: "JP JPN 392", input: "JPN"},
	{name: "TW TWN 158", input: "TWN"},
	{name: "ZW ZWE 716", input: "ZWE"},
}

var invalidStringISO3166Alpha3Cases = []stringRuleCase{
	{name: "numeric code", input: "001"},
	{name: "lowercase", input: "usa"},
	{name: "unassigned", input: "ZZZ"},
	{name: "deleted Soviet Union code", input: "SUN"},
	{name: "deleted Netherlands Antilles code", input: "ANT"},
	{name: "deleted East Germany code", input: "DDR"},
	{name: "user-assigned Kosovo code", input: "XKK"},
	{name: "exceptionally reserved territory", input: "ASC"},
	{name: "too short", input: "US"},
	{name: "too long", input: "USAA"},
	{name: "non-letter", input: "US1"},
	{name: "leading whitespace", input: " USA"},
	{name: "trailing whitespace", input: "USA "},
	{name: "full-width letters", input: "ＵＳＡ"},
}

func TestStringISO3166Alpha3(t *testing.T) {
	rule := StringISO3166Alpha3()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringISO3166Alpha3Cases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringISO3166Alpha3Cases)
		err := rule.Validate(invalidStringISO3166Alpha3Cases[0].input)
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

var validStringISO3166NumericCases = []stringRuleCase{
	{name: "AD AND 020", input: "020"},
	{name: "AX ALA 248", input: "248"},
	{name: "BQ BES 535", input: "535"},
	{name: "CW CUW 531", input: "531"},
	{name: "SS SSD 728", input: "728"},
	{name: "UM UMI 581", input: "581"},
	{name: "US USA 840", input: "840"},
	{name: "PL POL 616", input: "616"},
	{name: "JP JPN 392", input: "392"},
	{name: "TW TWN 158", input: "158"},
	{name: "ZW ZWE 716", input: "716"},
}

var invalidStringISO3166NumericCases = []stringRuleCase{
	{name: "world aggregate", input: "001"},
	{name: "too short", input: "84"},
	{name: "too long", input: "0840"},
	{name: "alphabetic", input: "USA"},
	{name: "non-digit", input: "84O"},
	{name: "deleted Soviet Union code", input: "810"},
	{name: "deleted Netherlands Antilles code", input: "530"},
	{name: "deleted East Germany code", input: "278"},
	{name: "user-assigned Kosovo code", input: "983"},
	{name: "leading whitespace", input: " 840"},
	{name: "trailing whitespace", input: "840 "},
	{name: "full-width digits", input: "８４０"},
}

func TestStringISO3166Numeric(t *testing.T) {
	rule := StringISO3166Numeric()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringISO3166NumericCases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringISO3166NumericCases)
		err := rule.Validate(invalidStringISO3166NumericCases[0].input)
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

var validStringISO31662Cases = []stringRuleCase{
	{name: "US state", input: "US-CA"},
	{name: "UK country", input: "GB-ENG"},
	{name: "Polish region", input: "PL-14"},
	{name: "two-digit subdivision", input: "AD-02"},
	{name: "mixed subdivision", input: "CZ-20A"},
	{name: "three-digit subdivision", input: "PW-002"},
	{name: "one-letter subdivision", input: "BD-H"},
	{name: "two-letter subdivision", input: "ID-PD"},
	{name: "Indian subdivision", input: "IN-LA"},
	{name: "Korean subdivision", input: "KR-50"},
	{name: "letter and digit subdivision", input: "NP-P7"},
}

var invalidStringISO31662Cases = []stringRuleCase{
	{name: "unknown US subdivision", input: "US-XXX"},
	{name: "unknown French subdivision", input: "FR-999"},
	{name: "unknown country", input: "ZZ-CA"},
	{name: "lowercase country", input: "us-CA"},
	{name: "lowercase subdivision", input: "US-ca"},
	{name: "lowercase and oversized subdivision", input: "US-cal"},
	{name: "underscore separator", input: "US_CA"},
	{name: "whitespace before separator", input: "US -CA"},
	{name: "trailing whitespace", input: "US-CA "},
	{name: "country code too long", input: "USA-CA"},
	{name: "subdivision code too long", input: "US-0000"},
	{name: "exceptionally reserved country", input: "UK-ENG"},
	{name: "deleted Soviet Union country", input: "SU-MOW"},
	{name: "deleted Netherlands Antilles country", input: "AN-CW"},
	{name: "user-assigned Kosovo country", input: "XK-01"},
	{name: "exceptionally reserved territory", input: "AC-SH"},
}

func TestStringISO31662(t *testing.T) {
	rule := StringISO31662()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringISO31662Cases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringISO31662Cases)
		err := rule.Validate(invalidStringISO31662Cases[0].input)
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

var validStringISO4217Cases = []stringRuleCase{
	{name: "US dollar", input: "USD"},
	{name: "euro", input: "EUR"},
	{name: "Japanese yen", input: "JPY"},
	{name: "Bolivian settlement fund", input: "BOV"},
	{name: "Chilean unit of account", input: "CLF"},
	{name: "Colombian real value unit", input: "COU"},
	{name: "Mexican investment unit", input: "MXV"},
	{name: "WIR euro", input: "CHE"},
	{name: "WIR franc", input: "CHW"},
	{name: "US dollar next day", input: "USN"},
	{name: "Uruguay peso indexed unit", input: "UYI"},
	{name: "gold", input: "XAU"},
	{name: "silver", input: "XAG"},
	{name: "palladium", input: "XPD"},
	{name: "platinum", input: "XPT"},
	{name: "testing code", input: "XTS"},
	{name: "no currency", input: "XXX"},
}

var invalidStringISO4217Cases = []stringRuleCase{
	{name: "lowercase", input: "usd"},
	{name: "unassigned", input: "ZZZ"},
	{name: "too short", input: "US"},
	{name: "too long", input: "USDD"},
	{name: "non-letter", input: "US1"},
	{name: "withdrawn Andorran peseta", input: "ADP"},
	{name: "withdrawn Estonian kroon", input: "EEK"},
	{name: "withdrawn Zimbabwe dollar", input: "ZWD"},
	{name: "withdrawn Russian ruble", input: "RUR"},
	{name: "withdrawn Belarusian ruble", input: "BYR"},
	{name: "leading whitespace", input: " USD"},
	{name: "trailing whitespace", input: "USD "},
	{name: "full-width letters", input: "ＵＳＤ"},
}

func TestStringISO4217(t *testing.T) {
	rule := StringISO4217()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringISO4217Cases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringISO4217Cases)
		err := rule.Validate(invalidStringISO4217Cases[0].input)
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

var validStringLatitudeCases = []stringRuleCase{
	{name: "south boundary", input: "-90"},
	{name: "north boundary", input: "90"},
	{name: "southern interior", input: "-89.999999"},
	{name: "northern interior", input: "89.999999"},
	{name: "equator", input: "0"},
	{name: "decimal degrees", input: "12.345678"},
}

var invalidStringLatitudeCases = []stringRuleCase{
	{name: "above north boundary", input: "90.1"},
	{name: "below south boundary", input: "-90.1"},
	{name: "cardinal direction", input: "north"},
	{name: "exponent notation", input: "1e1"},
	{name: "not a number", input: "NaN"},
	{name: "positive infinity", input: "+Inf"},
	{name: "negative infinity", input: "-Inf"},
	{name: "leading whitespace", input: " 45"},
	{name: "trailing whitespace", input: "45 "},
	{name: "decimal comma", input: "45,5"},
	{name: "full-width digits", input: "４５"},
	{name: "empty", input: ""},
}

func TestStringLatitude(t *testing.T) {
	rule := StringLatitude()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringLatitudeCases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringLatitudeCases)
		err := rule.Validate(invalidStringLatitudeCases[0].input)
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

var validStringLongitudeCases = []stringRuleCase{
	{name: "west boundary", input: "-180"},
	{name: "east boundary", input: "180"},
	{name: "western interior", input: "-179.999999"},
	{name: "eastern interior", input: "179.999999"},
	{name: "prime meridian", input: "0"},
	{name: "decimal degrees", input: "122.4194"},
}

var invalidStringLongitudeCases = []stringRuleCase{
	{name: "past east boundary", input: "180.1"},
	{name: "past west boundary", input: "-180.1"},
	{name: "cardinal direction", input: "east"},
	{name: "exponent notation", input: "1e2"},
	{name: "not a number", input: "NaN"},
	{name: "positive infinity", input: "+Inf"},
	{name: "negative infinity", input: "-Inf"},
	{name: "leading whitespace", input: " 45"},
	{name: "trailing whitespace", input: "45 "},
	{name: "decimal comma", input: "45,5"},
	{name: "full-width digits", input: "４５"},
	{name: "empty", input: ""},
}

func TestStringLongitude(t *testing.T) {
	rule := StringLongitude()
	t.Run("valid inputs", func(t *testing.T) {
		testStringRuleAccepts(t, rule, validStringLongitudeCases)
	})
	t.Run("invalid inputs", func(t *testing.T) {
		testStringRuleRejects(t, rule, invalidStringLongitudeCases)
		err := rule.Validate(invalidStringLongitudeCases[0].input)
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

func testStringRuleAccepts(t *testing.T, rule govy.Rule[string], cases []stringRuleCase) {
	t.Helper()
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.NoError(t, rule.Validate(testCase.input))
		})
	}
}

func testStringRuleRejects(t *testing.T, rule govy.Rule[string], cases []stringRuleCase) {
	t.Helper()
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Error(t, rule.Validate(testCase.input))
		})
	}
}
