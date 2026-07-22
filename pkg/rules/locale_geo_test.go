package rules

// cspell:ignore biske mingo phonebk Qaaa rozaj USAA

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
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
	{name: "repeated private use subtag", input: "x-abcde-abcde"},
	{name: "grandfathered regular tag", input: "i-default"},
	{name: "grandfathered irregular tag", input: "en-GB-oed"},
	{name: "grandfathered tag with multiple subtags", input: "zh-min-nan"},
	{name: "redundant tag", input: "zh-cmn-Hans-CN"},
	{name: "private use language script and region", input: "qaa-Qaaa-QM-x-southern"},
	{name: "current language unknown to x/text", input: "cls-Latn-US"},
	{name: "case-insensitive current language unknown to x/text", input: "CLS-latn-us"},
}

var invalidStringBCP47LanguageTagCases = []stringRuleCase{
	{name: "underscore separator", input: "en_GB"},
	{name: "non-tag word", input: "English"},
	{name: "empty", input: ""},
	{name: "trailing separator", input: "en-"},
	{name: "leading whitespace", input: " en"},
	{name: "trailing whitespace", input: "en "},
	{name: "duplicate extension singleton", input: "en-a-bbb-a-ccc"},
	{name: "duplicate variant", input: "de-1901-1901"},
	{name: "extension without value", input: "en-a-foo-b"},
	{name: "extension without value on current language unknown to x/text", input: "cls-a"},
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
	benchmarkStringRuleCases(
		b,
		StringBCP47LanguageTag(),
		validStringBCP47LanguageTagCases,
		invalidStringBCP47LanguageTagCases,
	)
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
	{name: "repeated private use subtag", input: "x-abcde-abcde"},
	{name: "extended language with script and region", input: "cmn-Hans-CN"},
	{name: "private use language script and region", input: "qaa-Qaaa-QM-x-southern"},
	{name: "registered suppress script for English", input: "en-Latn"},
	{name: "registered suppress script for Bosnian", input: "bs-Latn"},
	{name: "grandfathered Celtic Gaulish", input: "cel-gaulish"},
	{name: "grandfathered default language", input: "i-default"},
	{name: "grandfathered Enochian", input: "i-enochian"},
	{name: "grandfathered Mingo", input: "i-mingo"},
	{name: "grandfathered Min Chinese", input: "zh-min"},
}

var invalidStringBCP47StrictLanguageTagCases = []stringRuleCase{
	{name: "deprecated language", input: "iw"},
	{name: "non-canonical case", input: "EN-us"},
	{name: "underscore separator", input: "en_GB"},
	{name: "non-tag word", input: "English"},
	{name: "empty", input: ""},
	{name: "deprecated grandfathered tag", input: "en-GB-oed"},
	{name: "deprecated grandfathered tag with multiple subtags", input: "zh-min-nan"},
	{name: "redundant extended language", input: "zh-cmn-Hans-CN"},
	{name: "extension subtag too short", input: "tlh-a-b-foo"},
	{name: "duplicate extension singleton", input: "en-a-bbb-a-ccc"},
	{name: "duplicate variant", input: "de-1901-1901"},
	{name: "non-canonical extension order", input: "en-b-foo-a-bar"},
	{name: "non-canonical extension order on current language unknown to x/text", input: "cls-b-foo-a-bar"},
	{name: "extension without value", input: "en-a-foo-b"},
	{name: "extension without value on current language unknown to x/text", input: "cls-a"},
	{name: "non-canonical case on current language unknown to x/text", input: "CLS-Latn-US"},
	{name: "trailing separator", input: "en-"},
	{name: "leading whitespace", input: " en"},
	{name: "trailing whitespace", input: "en "},
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
	benchmarkStringRuleCases(
		b,
		StringBCP47StrictLanguageTag(),
		validStringBCP47StrictLanguageTagCases,
		invalidStringBCP47StrictLanguageTagCases,
	)
}

// The fixture contains every independently complete language-tag input in the
// IANA Language Subtag Registry dated 2026-06-14:
// https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry
// The source SHA-256 is be1fad86a99e3a932d07b80c9b3c271ec2381a5909ce22420144e5077ab0a43a.
// The normalized fixture SHA-256 is
// 939fc7568793edaa505b8f2fc62d17c6de8a9955e46b096d4778b5349084786a.
// It has 8,275 literal language subtags, all 520 expansions of qaa..qtz,
// 26 grandfathered Tags, and 67 redundant Tags: 8,888 unique inputs total.
// RFC 5646 section 4.5 makes 8,732 canonical and rejects the 156 records
// with a Preferred-Value. The 258 extlang records duplicate language inputs.
// The 225 script, 305 region, and 139 variant records are excluded because
// none is independently a complete language tag.
func TestStringBCP47IANACompleteInputCorpus(t *testing.T) {
	records := loadIANACompleteInputCorpus(t)
	languageTagRule := StringBCP47LanguageTag()
	strictLanguageTagRule := StringBCP47StrictLanguageTag()

	for line, record := range records {
		if err := languageTagRule.Validate(record.input); err != nil {
			t.Errorf("IANA fixture line %d: non-strict rule rejected %q: %v", line+1, record.input, err)
		}

		err := strictLanguageTagRule.Validate(record.input)
		switch record.strictOutcome {
		case "accept":
			if err != nil {
				t.Errorf("IANA fixture line %d: strict rule rejected %q: %v", line+1, record.input, err)
			}
		case "reject":
			if err == nil {
				t.Errorf("IANA fixture line %d: strict rule accepted non-canonical %q", line+1, record.input)
			} else if !govy.HasErrorCode(err, ErrorCodeStringBCP47StrictLanguageTag) {
				t.Errorf("IANA fixture line %d: strict error for %q lacks code %q: %v",
					line+1, record.input, ErrorCodeStringBCP47StrictLanguageTag, err)
			}
		default:
			t.Fatalf("IANA fixture line %d: unknown strict outcome %q", line+1, record.strictOutcome)
		}
	}
}

func TestStringBCP47CurrentIANAParserOverrides(t *testing.T) {
	bases := []string{
		"bih", "cls", "dyl", "hnm", "isv", "lfb", "luh", "oak", "olb",
		"osd", "rrm", "scz", "sjc", "tvg", "vsn", "ynb", "zhk",
	}
	suffixes := []string{"", "-Latn-US", "-u-ca-gregory", "-a-foo-b-bar"}
	languageTagRule := StringBCP47LanguageTag()
	strictLanguageTagRule := StringBCP47StrictLanguageTag()

	for _, base := range bases {
		for _, suffix := range suffixes {
			input := base + suffix
			t.Run(input, func(t *testing.T) {
				assert.NoError(t, languageTagRule.Validate(input))
				assert.NoError(t, strictLanguageTagRule.Validate(input))
			})
		}
	}
}

func TestStringBCP47IANAPreferredValues(t *testing.T) {
	records := loadIANACompleteInputCorpus(t)
	languageTagRule := StringBCP47LanguageTag()
	strictLanguageTagRule := StringBCP47StrictLanguageTag()
	languageCompositeCount := 0
	wholeTagCount := 0
	redundantSignTagCount := 0

	for _, record := range records {
		if record.preferredValue == "" {
			continue
		}
		switch record.sourceKind {
		case "language":
			input := record.input + "-US"
			if err := languageTagRule.Validate(input); err != nil {
				t.Errorf("non-strict rule rejected Preferred-Value language composite %q: %v", input, err)
			}
			if err := strictLanguageTagRule.Validate(input); err == nil {
				t.Errorf("strict rule accepted Preferred-Value language composite %q", input)
			}
			languageCompositeCount++
		case "grandfathered", "redundant":
			if err := languageTagRule.Validate(record.input); err != nil {
				t.Errorf("non-strict rule rejected Preferred-Value Tag %q: %v", record.input, err)
			}
			if err := strictLanguageTagRule.Validate(record.input); err == nil {
				t.Errorf("strict rule accepted Preferred-Value Tag %q", record.input)
			}
			wholeTagCount++
			if record.sourceKind == "redundant" && strings.HasPrefix(record.input, "sgn-") {
				redundantSignTagCount++
			}
		default:
			t.Fatalf("unexpected source kind %q", record.sourceKind)
		}
	}

	assert.Equal(t, 110, languageCompositeCount)
	assert.Equal(t, 46, wholeTagCount)
	assert.Equal(t, 19, redundantSignTagCount)
}

func TestStringBCP47IANAPreferredSignTagExtensions(t *testing.T) {
	records := loadIANACompleteInputCorpus(t)
	languageTagRule := StringBCP47LanguageTag()
	strictLanguageTagRule := StringBCP47StrictLanguageTag()
	suffixes := []string{"-x-foo", "-a-foo"}
	baseCount := 0
	inputCount := 0

	for _, record := range records {
		if record.sourceKind != "redundant" || record.preferredValue == "" ||
			!strings.HasPrefix(record.input, "sgn-") {
			continue
		}
		baseCount++
		for _, suffix := range suffixes {
			input := record.input + suffix
			inputCount++
			t.Run(input, func(t *testing.T) {
				assert.NoError(t, languageTagRule.Validate(input))
				err := strictLanguageTagRule.Validate(input)
				if err == nil {
					t.Errorf("strict rule accepted Preferred-Value Tag extension %q", input)
					return
				}
				assert.True(t, govy.HasErrorCode(err, ErrorCodeStringBCP47StrictLanguageTag))
			})
		}
	}

	assert.Equal(t, 19, baseCount)
	assert.Equal(t, 38, inputCount)
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
	benchmarkStringRuleCases(
		b,
		StringISO3166Alpha2(),
		validStringISO3166Alpha2Cases,
		invalidStringISO3166Alpha2Cases,
	)
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
	benchmarkStringRuleCases(
		b,
		StringISO3166Alpha3(),
		validStringISO3166Alpha3Cases,
		invalidStringISO3166Alpha3Cases,
	)
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
	benchmarkStringRuleCases(
		b,
		StringISO3166Numeric(),
		validStringISO3166NumericCases,
		invalidStringISO3166NumericCases,
	)
}

// The fixture contains all 249 current ISO 3166-1 alpha-2, alpha-3, and
// numeric triplets. It normalizes the 248 triplets in the English table at
// https://unstats.un.org/unsd/methodology/m49/overview/ (source SHA-256
// 748f6ff7380c8a50ea9448f068b79e3a1ee31be63207249e8cc89bf1eb969d11)
// and adds TW/TWN/158 from https://www.iso.org/obp/ui/#iso:code:3166:TW.
// The normalized fixture SHA-256 is
// d7679f496d423a8071e22852dddcc2e261c0ba12da94a7418ad95582e701e033.
// UNSD aggregate and geographic-group rows are excluded because they have no
// ISO alpha code elements; localized duplicate tables are also excluded.
func TestStringISO31661CurrentCodeCorpus(t *testing.T) {
	triplets := loadISO31661Triplets(t)
	alpha2Rule := StringISO3166Alpha2()
	alpha3Rule := StringISO3166Alpha3()
	numericRule := StringISO3166Numeric()

	for line, triplet := range triplets {
		if err := alpha2Rule.Validate(triplet.alpha2); err != nil {
			t.Fatalf("ISO 3166-1 fixture line %d: alpha-2 %q rejected: %v", line+1, triplet.alpha2, err)
		}
		if err := alpha3Rule.Validate(triplet.alpha3); err != nil {
			t.Fatalf("ISO 3166-1 fixture line %d: alpha-3 %q rejected: %v", line+1, triplet.alpha3, err)
		}
		if err := numericRule.Validate(triplet.numeric); err != nil {
			t.Fatalf("ISO 3166-1 fixture line %d: numeric %q rejected: %v", line+1, triplet.numeric, err)
		}
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
	{name: "Metropole de Lyon", input: "FR-69M"},
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
	benchmarkStringRuleCases(
		b,
		StringISO31662(),
		validStringISO31662Cases,
		invalidStringISO31662Cases,
	)
}

// The fixture projects the country and subdivision fields from every record
// in UNECE's ISO 3166-2-derived UN/LOCODE 2024-2 SubdivisionCodes.csv at tag
// commit 03032eb5f0eed1db15dc5255f7dcc0942d7b6238:
// https://opensource.unicc.org/un/unece/uncefact/vocab-locode/-/raw/03032eb5f0eed1db15dc5255f7dcc0942d7b6238/iso-3166/SubdivisionCodes.csv
// The source SHA-256 is b6fcf7c4a554598db32a372fd935a6e12291490ca76a8a0dd712ad849ff7cac6;
// the normalized fixture SHA-256 is
// 020b4ad2c330f2d9cd7e42bd53e84e6b9546ce5462475e6fc348fcd1c73620b9.
// All 4,676 source records are preserved, including the duplicate identifiers
// IN-JK, MA-CHT, MA-KES, and MK-205; there are 4,672 distinct identifiers.
// Names and subdivision types are excluded because this rule accepts only the
// complete country-subdivision identifier.
func TestStringISO31662UNECE2024_2Corpus(t *testing.T) {
	rule := StringISO31662()
	for line, code := range loadISO31662Codes(t) {
		if err := rule.Validate(code); err != nil {
			t.Fatalf("UNECE ISO 3166-2 fixture line %d: code %q rejected: %v", line+1, code, err)
		}
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
	{name: "Mauritanian ouguiya", input: "MRU"},
	{name: "Sierra Leonean leone", input: "SLE"},
	{name: "Salvadoran colon", input: "SVC"},
	{name: "Uruguay nominal wage index unit", input: "UYW"},
	{name: "Venezuelan digital bolivar", input: "VED"},
	{name: "Venezuelan bolivar", input: "VES"},
	{name: "Arab accounting dinar", input: "XAD"},
	{name: "Caribbean guilder", input: "XCG"},
	{name: "Zimbabwe gold", input: "ZWG"},
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
	{name: "withdrawn Netherlands Antillean guilder", input: "ANG"},
	{name: "withdrawn Bulgarian lev", input: "BGN"},
	{name: "unassigned offshore renminbi", input: "CNH"},
	{name: "withdrawn Cuban convertible peso", input: "CUC"},
	{name: "withdrawn Croatian kuna", input: "HRK"},
	{name: "withdrawn Mauritanian ouguiya", input: "MRO"},
	{name: "withdrawn Sierra Leonean leone", input: "SLL"},
	{name: "withdrawn Venezuelan bolivar", input: "VEF"},
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
	benchmarkStringRuleCases(
		b,
		StringISO4217(),
		validStringISO4217Cases,
		invalidStringISO4217Cases,
	)
}

// The fixture contains every distinct populated Ccy value from SIX ISO 4217
// List One published 2026-01-01:
// https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml
// The XML source SHA-256 is
// 838dfb991648cf36df939edd5fe3811737962b75a32252847d239cedd1e291c9;
// the 178-code fixture SHA-256 is
// 6bd74ecf83510ca425c07b1cef1117f9f5a7d76c3e5079365feb841ababa3466.
// Repeated country uses of a currency are collapsed because this rule validates
// the code element. Names, numeric codes, and minor units are excluded.
func TestStringISO4217SIXListOneCorpus(t *testing.T) {
	rule := StringISO4217()
	for line, code := range loadISO4217Codes(t) {
		if err := rule.Validate(code); err != nil {
			t.Fatalf("SIX ISO 4217 fixture line %d: code %q rejected: %v", line+1, code, err)
		}
	}
}

var validStringLatitudeCases = []stringRuleCase{
	{name: "south boundary", input: "-90"},
	{name: "north boundary", input: "90"},
	{name: "southern interior", input: "-89.999999"},
	{name: "northern interior", input: "89.999999"},
	{name: "equator", input: "0"},
	{name: "negative zero", input: "-0"},
	{name: "decimal degrees", input: "12.345678"},
	{name: "high-precision north boundary", input: "90.0000000000000000000"},
	{name: "high-precision south boundary", input: "-90.0000000000000000000"},
	{name: "high-precision northern interior", input: "89.9999999999999999999"},
	{name: "high-precision southern interior", input: "-89.9999999999999999999"},
}

var invalidStringLatitudeCases = []stringRuleCase{
	{name: "above north boundary", input: "90.1"},
	{name: "below south boundary", input: "-90.1"},
	{name: "high-precision above north boundary", input: "90.0000000000000000001"},
	{name: "high-precision below south boundary", input: "-90.0000000000000000001"},
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
	benchmarkStringRuleCases(
		b,
		StringLatitude(),
		validStringLatitudeCases,
		invalidStringLatitudeCases,
	)
}

var validStringLongitudeCases = []stringRuleCase{
	{name: "west boundary", input: "-180"},
	{name: "east boundary", input: "180"},
	{name: "western interior", input: "-179.999999"},
	{name: "eastern interior", input: "179.999999"},
	{name: "prime meridian", input: "0"},
	{name: "negative zero", input: "-0"},
	{name: "decimal degrees", input: "122.4194"},
	{name: "high-precision east boundary", input: "180.0000000000000000000"},
	{name: "high-precision west boundary", input: "-180.0000000000000000000"},
	{name: "high-precision eastern interior", input: "179.9999999999999999999"},
	{name: "high-precision western interior", input: "-179.9999999999999999999"},
}

var invalidStringLongitudeCases = []stringRuleCase{
	{name: "past east boundary", input: "180.1"},
	{name: "past west boundary", input: "-180.1"},
	{name: "high-precision past east boundary", input: "180.0000000000000000001"},
	{name: "high-precision past west boundary", input: "-180.0000000000000000001"},
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
	benchmarkStringRuleCases(
		b,
		StringLongitude(),
		validStringLongitudeCases,
		invalidStringLongitudeCases,
	)
}

// These tables contain every distinct latitude and longitude scalar literal
// used by RFC 5870's geo URI examples and invalid-location example:
// https://www.rfc-editor.org/rfc/rfc5870.txt
// The source SHA-256 is a3bf93084c422508bbf64976e9adee1eae88d688a2e7f7fc3a531aa44ceb3bb1.
// Formatting variants such as 22.300/22.3, -118.44/-118.4400, and 66/66.0
// remain distinct. Repeated literals, altitude, uncertainty, parameters, and
// template placeholders are excluded. Latitude 94 is the RFC's invalid value.
func TestStringCoordinatesRFC5870ScalarCorpus(t *testing.T) {
	type coordinateCase struct {
		input string
		valid bool
	}
	tests := map[string]struct {
		rule  govy.Rule[string]
		cases []coordinateCase
	}{
		"latitude": {
			rule: StringLatitude(),
			cases: []coordinateCase{
				{input: "13.4125", valid: true},
				{input: "48.2010", valid: true},
				{input: "48.198634", valid: true},
				{input: "90", valid: true},
				{input: "22.300", valid: true},
				{input: "22.3", valid: true},
				{input: "66", valid: true},
				{input: "66.0", valid: true},
				{input: "70", valid: true},
				{input: "47", valid: true},
				{input: "22", valid: true},
				{input: "94", valid: false},
			},
		},
		"longitude": {
			rule: StringLongitude(),
			cases: []coordinateCase{
				{input: "103.8667", valid: true},
				{input: "16.3695", valid: true},
				{input: "16.371648", valid: true},
				{input: "-22.43", valid: true},
				{input: "46", valid: true},
				{input: "-118.44", valid: true},
				{input: "-118.4400", valid: true},
				{input: "30", valid: true},
				{input: "20", valid: true},
				{input: "11", valid: true},
				{input: "0", valid: true},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, testCase := range test.cases {
				t.Run(testCase.input, func(t *testing.T) {
					err := test.rule.Validate(testCase.input)
					if testCase.valid {
						assert.NoError(t, err)
					} else {
						assert.Error(t, err)
					}
				})
			}
		})
	}
}

type ianaCompleteInputRecord struct {
	sourceKind     string
	input          string
	preferredValue string
	strictOutcome  string
	strictReason   string
}

type iso31661Triplet struct {
	alpha2  string
	alpha3  string
	numeric string
}

func loadISO31661Triplets(t *testing.T) []iso31661Triplet {
	t.Helper()

	const fixture = "iso_3166_1_triplets_2026-07-21.txt"
	path, lines := loadFixtureLines(
		t,
		fixture,
		249,
		"d7679f496d423a8071e22852dddcc2e261c0ba12da94a7418ad95582e701e033",
	)

	triplets := make([]iso31661Triplet, 0, len(lines))
	seenAlpha2 := make(map[string]int, len(lines))
	seenAlpha3 := make(map[string]int, len(lines))
	seenNumeric := make(map[string]int, len(lines))
	hasTaiwan := false
	for line, value := range lines {
		fields := strings.Split(value, ",")
		if len(fields) != 3 {
			t.Fatalf("%s:%d: got %d fields; want 3", path, line+1, len(fields))
		}
		triplet := iso31661Triplet{alpha2: fields[0], alpha3: fields[1], numeric: fields[2]}
		if !isASCIIUpperTestValue(triplet.alpha2, 2) {
			t.Fatalf("%s:%d: invalid alpha-2 field %q", path, line+1, triplet.alpha2)
		}
		if !isASCIIUpperTestValue(triplet.alpha3, 3) {
			t.Fatalf("%s:%d: invalid alpha-3 field %q", path, line+1, triplet.alpha3)
		}
		if !isASCIIDigitTestValue(triplet.numeric, 3) {
			t.Fatalf("%s:%d: invalid numeric field %q", path, line+1, triplet.numeric)
		}
		checkUniqueFixtureValue(t, path, line+1, "alpha-2", triplet.alpha2, seenAlpha2)
		checkUniqueFixtureValue(t, path, line+1, "alpha-3", triplet.alpha3, seenAlpha3)
		checkUniqueFixtureValue(t, path, line+1, "numeric", triplet.numeric, seenNumeric)
		if triplet == (iso31661Triplet{alpha2: "TW", alpha3: "TWN", numeric: "158"}) {
			hasTaiwan = true
		}
		triplets = append(triplets, triplet)
	}
	if !hasTaiwan {
		t.Fatalf("%s must contain TW,TWN,158", path)
	}
	return triplets
}

func loadISO4217Codes(t *testing.T) []string {
	t.Helper()

	const fixture = "six_iso4217_list_one_2026-01-01.txt"
	path, codes := loadFixtureLines(
		t,
		fixture,
		178,
		"6bd74ecf83510ca425c07b1cef1117f9f5a7d76c3e5079365feb841ababa3466",
	)
	seen := make(map[string]int, len(codes))
	for line, code := range codes {
		if !isASCIIUpperTestValue(code, 3) {
			t.Fatalf("%s:%d: invalid currency code %q", path, line+1, code)
		}
		checkUniqueFixtureValue(t, path, line+1, "currency code", code, seen)
	}
	return codes
}

func loadISO31662Codes(t *testing.T) []string {
	t.Helper()

	const fixture = "unece_iso3166_2_2024-2.txt"
	path, codes := loadFixtureLines(
		t,
		fixture,
		4676,
		"020b4ad2c330f2d9cd7e42bd53e84e6b9546ce5462475e6fc348fcd1c73620b9",
	)

	counts := make(map[string]int, 4672)
	for line, code := range codes {
		country, subdivision, ok := strings.Cut(code, "-")
		if !ok || !isASCIIUpperTestValue(country, 2) || !isASCIIUpperDigitTestValue(subdivision, 1, 3) {
			t.Fatalf("%s:%d: invalid subdivision identifier %q", path, line+1, code)
		}
		counts[code]++
	}
	if len(counts) != 4672 {
		t.Fatalf("%s contains %d distinct identifiers; want 4672", path, len(counts))
	}

	expectedDuplicates := map[string]struct{}{
		"IN-JK":  {},
		"MA-CHT": {},
		"MA-KES": {},
		"MK-205": {},
	}
	for code, count := range counts {
		_, expectedDuplicate := expectedDuplicates[code]
		switch {
		case count == 1 && !expectedDuplicate:
		case count == 2 && expectedDuplicate:
			delete(expectedDuplicates, code)
		default:
			t.Fatalf("%s: identifier %q occurs %d times", path, code, count)
		}
	}
	if len(expectedDuplicates) != 0 {
		t.Fatalf("%s is missing expected duplicate identifiers: %v", path, expectedDuplicates)
	}
	if counts["FR-69M"] != 1 {
		t.Fatalf("%s must contain FR-69M exactly once", path)
	}
	return codes
}

func loadIANACompleteInputCorpus(t *testing.T) []ianaCompleteInputRecord {
	t.Helper()

	const fixture = "iana_bcp47_complete_inputs_2026-06-14.txt"
	path, lines := loadFixtureLines(
		t,
		fixture,
		8888,
		"939fc7568793edaa505b8f2fc62d17c6de8a9955e46b096d4778b5349084786a",
	)

	records := make([]ianaCompleteInputRecord, 0, len(lines))
	seenInputs := make(map[string]int, len(lines))
	rangeInputs := make(map[string]struct{}, 520)
	kindCounts := make(map[string]int, 4)
	languageOutcomeCounts := make(map[string]int, 2)
	tagOutcomeCounts := make(map[string]int, 2)
	reasonCounts := make(map[string]int, 1)
	for line, value := range lines {
		record := parseIANACompleteInputRecord(t, path, line+1, value)
		if record.sourceKind == "language-range" {
			rangeInputs[record.input] = struct{}{}
		}

		if firstLine, ok := seenInputs[record.input]; ok {
			t.Fatalf("%s:%d: duplicate input %q first seen on line %d", path, line+1, record.input, firstLine)
		}
		seenInputs[record.input] = line + 1

		kindCounts[record.sourceKind]++
		if record.sourceKind == "language" || record.sourceKind == "language-range" {
			languageOutcomeCounts[record.strictOutcome]++
		} else {
			tagOutcomeCounts[record.strictOutcome]++
		}
		reasonCounts[record.strictReason]++
		records = append(records, record)
	}

	if kindCounts["language"] != 8275 || kindCounts["language-range"] != 520 ||
		kindCounts["grandfathered"] != 26 || kindCounts["redundant"] != 67 {
		t.Fatalf("%s: source-kind counts are language=%d language-range=%d grandfathered=%d redundant=%d",
			path,
			kindCounts["language"],
			kindCounts["language-range"],
			kindCounts["grandfathered"],
			kindCounts["redundant"],
		)
	}
	if len(rangeInputs) != 520 {
		t.Fatalf("%s expands %d qaa..qtz inputs; want 520", path, len(rangeInputs))
	}
	for second := 'a'; second <= 't'; second++ {
		for third := 'a'; third <= 'z'; third++ {
			input := string([]rune{'q', second, third})
			if _, ok := rangeInputs[input]; !ok {
				t.Fatalf("%s is missing range expansion %q", path, input)
			}
		}
	}
	if languageOutcomeCounts["accept"] != 8685 || languageOutcomeCounts["reject"] != 110 {
		t.Fatalf("%s: language strict outcomes are accept=%d reject=%d; want 8685 and 110",
			path, languageOutcomeCounts["accept"], languageOutcomeCounts["reject"])
	}
	if tagOutcomeCounts["accept"] != 47 || tagOutcomeCounts["reject"] != 46 {
		t.Fatalf("%s: Tag strict outcomes are accept=%d reject=%d; want 47 and 46",
			path, tagOutcomeCounts["accept"], tagOutcomeCounts["reject"])
	}
	if reasonCounts["preferred-value"] != 156 || reasonCounts[""] != 8732 {
		t.Fatalf("%s: reasons are Preferred-Value=%d empty=%d; want 156 and 8732",
			path, reasonCounts["preferred-value"], reasonCounts[""])
	}
	return records
}

func parseIANACompleteInputRecord(
	t *testing.T,
	path string,
	line int,
	value string,
) ianaCompleteInputRecord {
	t.Helper()

	fields := strings.Split(value, "|")
	if len(fields) != 5 {
		t.Fatalf("%s:%d: got %d fields; want 5", path, line, len(fields))
	}
	record := ianaCompleteInputRecord{
		sourceKind:     fields[0],
		input:          fields[1],
		preferredValue: fields[2],
		strictOutcome:  fields[3],
		strictReason:   fields[4],
	}

	switch record.sourceKind {
	case "language":
		if !isASCIILowerTestValue(record.input, 2, 8) {
			t.Fatalf("%s:%d: invalid language input %q", path, line, record.input)
		}
	case "language-range":
		if len(record.input) != 3 || record.input[0] != 'q' ||
			record.input[1] < 'a' || record.input[1] > 't' ||
			record.input[2] < 'a' || record.input[2] > 'z' {
			t.Fatalf("%s:%d: input %q is outside qaa..qtz", path, line, record.input)
		}
	case "grandfathered", "redundant":
		if record.input == "" {
			t.Fatalf("%s:%d: empty Tag field", path, line)
		}
	default:
		t.Fatalf("%s:%d: unknown source kind %q", path, line, record.sourceKind)
	}

	switch record.strictOutcome {
	case "accept":
		if record.preferredValue != "" || record.strictReason != "" {
			t.Fatalf("%s:%d: accepted input has Preferred-Value %q and reason %q",
				path, line, record.preferredValue, record.strictReason)
		}
	case "reject":
		if record.preferredValue == "" || record.strictReason != "preferred-value" {
			t.Fatalf("%s:%d: rejection has Preferred-Value %q and reason %q",
				path, line, record.preferredValue, record.strictReason)
		}
	default:
		t.Fatalf("%s:%d: unknown strict outcome %q", path, line, record.strictOutcome)
	}
	return record
}

func loadFixtureLines(
	t *testing.T,
	fixture string,
	expectedCount int,
	expectedSHA256 string,
) (path string, lines []string) {
	t.Helper()

	path = filepath.Join("testdata", fixture)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if len(data) == 0 || data[len(data)-1] != '\n' {
		t.Fatalf("%s must be non-empty and end with a newline", path)
	}
	if strings.ContainsRune(string(data), '\r') {
		t.Fatalf("%s must use LF line endings", path)
	}
	sum := sha256.Sum256(data)
	if actualSHA256 := hex.EncodeToString(sum[:]); actualSHA256 != expectedSHA256 {
		t.Fatalf("%s SHA-256 is %s; want %s", path, actualSHA256, expectedSHA256)
	}

	lines = strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	if len(lines) != expectedCount {
		t.Fatalf("%s contains %d records; want %d", path, len(lines), expectedCount)
	}
	return path, lines
}

func isASCIIUpperTestValue(value string, expectedLength int) bool {
	if len(value) != expectedLength {
		return false
	}
	for _, char := range value {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}

func isASCIILowerTestValue(value string, minLength, maxLength int) bool {
	if len(value) < minLength || len(value) > maxLength {
		return false
	}
	for _, char := range value {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}

func isASCIIDigitTestValue(value string, expectedLength int) bool {
	if len(value) != expectedLength {
		return false
	}
	for _, char := range value {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func isASCIIUpperDigitTestValue(value string, minLength, maxLength int) bool {
	if len(value) < minLength || len(value) > maxLength {
		return false
	}
	for _, char := range value {
		if (char < 'A' || char > 'Z') && (char < '0' || char > '9') {
			return false
		}
	}
	return true
}

func checkUniqueFixtureValue(
	t *testing.T,
	path string,
	line int,
	field string,
	value string,
	seen map[string]int,
) {
	t.Helper()
	if firstLine, ok := seen[value]; ok {
		t.Fatalf("%s:%d: duplicate %s %q first seen on line %d", path, line, field, value, firstLine)
	}
	seen[value] = line
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

func benchmarkStringRuleCases(
	b *testing.B,
	rule govy.Rule[string],
	validCases []stringRuleCase,
	invalidCases []stringRuleCase,
) {
	b.Helper()
	for b.Loop() {
		for _, testCase := range validCases {
			_ = rule.Validate(testCase.input)
		}
		for _, testCase := range invalidCases {
			_ = rule.Validate(testCase.input)
		}
	}
}
