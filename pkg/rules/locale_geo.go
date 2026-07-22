package rules

// cspell:ignore guoyu lojban mingo xiang

import (
	"strconv"
	"strings"
	"sync"

	textcurrency "golang.org/x/text/currency"
	"golang.org/x/text/language"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// These BCP 47 compatibility tables are derived from every applicable record
// in the IANA Language Subtag Registry dated 2026-06-14. The source SHA-256 is
// be1fad86a99e3a932d07b80c9b3c271ec2381a5909ce22420144e5077ab0a43a.
const (
	bcp47LanguagePreferredValueBasesData = `
bh in iw ji jw mo aam adp ajp ajt asd aue ayx bgm bic bjd blg ccq cjr cka cmk coy cqu dek dit drh
drr drw gav gfx ggn gli gti guv hrr ibi ilw jeg kgc kgh kgm koj krm ktr kvs kwq kxe kxl kzj kzt
lak lii llo lmm meg mst mwj myd myt nad ncp nns nnx nom nte nts nxu oun pat pcr pmc pmk pmu ppa
ppr prp pry puz sca skk smd snb szd tdu thc thw thx tie tkk tlw tmk tmp tne tnf tpw tsf uok xba
xia xkh xrq xss ybd yma ymt yol yos yuu zir zkb
`
	bcp47TagPreferredValuesData = `
art-lojban en-GB-oed i-ami i-bnn i-hak i-klingon i-lux i-navajo i-pwn i-tao i-tay i-tsu no-bok
no-nyn sgn-BE-FR sgn-BE-NL sgn-CH-DE zh-guoyu zh-hakka zh-min-nan zh-xiang sgn-BR sgn-CO sgn-DE
sgn-DK sgn-ES sgn-FR sgn-GB sgn-GR sgn-IE sgn-IT sgn-JP sgn-MX sgn-NI sgn-NL sgn-NO sgn-PT
sgn-SE sgn-US sgn-ZA zh-cmn zh-cmn-Hans zh-cmn-Hant zh-gan zh-wuu zh-yue
`
	bcp47ParserOverrideBasesData = `
bih cls dyl hnm isv lfb luh oak olb osd rrm scz sjc tvg vsn ynb zhk
`
	bcp47GrandfatheredWithoutPreferredValueData = `
cel-gaulish i-default i-enochian i-mingo zh-min
`
	bcp47NeutralParserBase = "qaa"
)

var (
	bcp47LanguagePreferredValueBases = lazyLookupMap(func() map[string]struct{} {
		return buildStringLookup(bcp47LanguagePreferredValueBasesData)
	})
	bcp47TagPreferredValues = lazyLookupMap(func() map[string]struct{} {
		return buildStringLookup(bcp47TagPreferredValuesData)
	})
	bcp47ParserOverrideBases = lazyLookupMap(func() map[string]struct{} {
		return buildStringLookup(bcp47ParserOverrideBasesData)
	})
	bcp47GrandfatheredWithoutPreferredValue = lazyLookupMap(func() map[string]struct{} {
		return buildStringLookup(bcp47GrandfatheredWithoutPreferredValueData)
	})
)

// iso3166Alpha2Codes returns the accepted ISO 3166-1 alpha-2 code elements.
// x/text also recognizes aliases, reserved codes, and deleted codes.
var iso3166Alpha2Codes = lazyLookupMap(func() map[string]struct{} {
	return map[string]struct{}{
		"AD": {},
		"AE": {},
		"AF": {},
		"AG": {},
		"AI": {},
		"AL": {},
		"AM": {},
		"AO": {},
		"AQ": {},
		"AR": {},
		"AS": {},
		"AT": {},
		"AU": {},
		"AW": {},
		"AX": {},
		"AZ": {},
		"BA": {},
		"BB": {},
		"BD": {},
		"BE": {},
		"BF": {},
		"BG": {},
		"BH": {},
		"BI": {},
		"BJ": {},
		"BL": {},
		"BM": {},
		"BN": {},
		"BO": {},
		"BQ": {},
		"BR": {},
		"BS": {},
		"BT": {},
		"BV": {},
		"BW": {},
		"BY": {},
		"BZ": {},
		"CA": {},
		"CC": {},
		"CD": {},
		"CF": {},
		"CG": {},
		"CH": {},
		"CI": {},
		"CK": {},
		"CL": {},
		"CM": {},
		"CN": {},
		"CO": {},
		"CR": {},
		"CU": {},
		"CV": {},
		"CW": {},
		"CX": {},
		"CY": {},
		"CZ": {},
		"DE": {},
		"DJ": {},
		"DK": {},
		"DM": {},
		"DO": {},
		"DZ": {},
		"EC": {},
		"EE": {},
		"EG": {},
		"EH": {},
		"ER": {},
		"ES": {},
		"ET": {},
		"FI": {},
		"FJ": {},
		"FK": {},
		"FM": {},
		"FO": {},
		"FR": {},
		"GA": {},
		"GB": {},
		"GD": {},
		"GE": {},
		"GF": {},
		"GG": {},
		"GH": {},
		"GI": {},
		"GL": {},
		"GM": {},
		"GN": {},
		"GP": {},
		"GQ": {},
		"GR": {},
		"GS": {},
		"GT": {},
		"GU": {},
		"GW": {},
		"GY": {},
		"HK": {},
		"HM": {},
		"HN": {},
		"HR": {},
		"HT": {},
		"HU": {},
		"ID": {},
		"IE": {},
		"IL": {},
		"IM": {},
		"IN": {},
		"IO": {},
		"IQ": {},
		"IR": {},
		"IS": {},
		"IT": {},
		"JE": {},
		"JM": {},
		"JO": {},
		"JP": {},
		"KE": {},
		"KG": {},
		"KH": {},
		"KI": {},
		"KM": {},
		"KN": {},
		"KP": {},
		"KR": {},
		"KW": {},
		"KY": {},
		"KZ": {},
		"LA": {},
		"LB": {},
		"LC": {},
		"LI": {},
		"LK": {},
		"LR": {},
		"LS": {},
		"LT": {},
		"LU": {},
		"LV": {},
		"LY": {},
		"MA": {},
		"MC": {},
		"MD": {},
		"ME": {},
		"MF": {},
		"MG": {},
		"MH": {},
		"MK": {},
		"ML": {},
		"MM": {},
		"MN": {},
		"MO": {},
		"MP": {},
		"MQ": {},
		"MR": {},
		"MS": {},
		"MT": {},
		"MU": {},
		"MV": {},
		"MW": {},
		"MX": {},
		"MY": {},
		"MZ": {},
		"NA": {},
		"NC": {},
		"NE": {},
		"NF": {},
		"NG": {},
		"NI": {},
		"NL": {},
		"NO": {},
		"NP": {},
		"NR": {},
		"NU": {},
		"NZ": {},
		"OM": {},
		"PA": {},
		"PE": {},
		"PF": {},
		"PG": {},
		"PH": {},
		"PK": {},
		"PL": {},
		"PM": {},
		"PN": {},
		"PR": {},
		"PS": {},
		"PT": {},
		"PW": {},
		"PY": {},
		"QA": {},
		"RE": {},
		"RO": {},
		"RS": {},
		"RU": {},
		"RW": {},
		"SA": {},
		"SB": {},
		"SC": {},
		"SD": {},
		"SE": {},
		"SG": {},
		"SH": {},
		"SI": {},
		"SJ": {},
		"SK": {},
		"SL": {},
		"SM": {},
		"SN": {},
		"SO": {},
		"SR": {},
		"SS": {},
		"ST": {},
		"SV": {},
		"SX": {},
		"SY": {},
		"SZ": {},
		"TC": {},
		"TD": {},
		"TF": {},
		"TG": {},
		"TH": {},
		"TJ": {},
		"TK": {},
		"TL": {},
		"TM": {},
		"TN": {},
		"TO": {},
		"TR": {},
		"TT": {},
		"TV": {},
		"TW": {},
		"TZ": {},
		"UA": {},
		"UG": {},
		"UM": {},
		"US": {},
		"UY": {},
		"UZ": {},
		"VA": {},
		"VC": {},
		"VE": {},
		"VG": {},
		"VI": {},
		"VN": {},
		"VU": {},
		"WF": {},
		"WS": {},
		"YE": {},
		"YT": {},
		"ZA": {},
		"ZM": {},
		"ZW": {},
	}
})

// iso4217Codes returns current tender and non-tender ISO 4217 code elements.
// ParseISO also recognizes withdrawn codes.
var iso4217Codes = lazyLookupMap(buildISO4217Codes)

// StringBCP47LanguageTag ensures the property's value is a valid BCP 47 language tag.
func StringBCP47LanguageTag() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBCP47LanguageTagTemplate)

	return govy.NewRule(func(s string) error {
		if !isBCP47LanguageTag(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBCP47LanguageTag).
		WithMessageTemplate(tpl).
		WithExamples("en", "en-US", "zh-Hant-TW").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringBCP47StrictLanguageTag ensures the property's value is a valid canonical BCP 47 language tag.
func StringBCP47StrictLanguageTag() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringBCP47StrictLanguageTagTemplate)

	return govy.NewRule(func(s string) error {
		if !isCanonicalBCP47LanguageTag(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringBCP47StrictLanguageTag).
		WithMessageTemplate(tpl).
		WithExamples("en", "en-US", "zh-Hant-TW").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringISO3166Alpha2 ensures the property's value is a valid ISO 3166-1 alpha-2 country code.
func StringISO3166Alpha2() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISO3166Alpha2Template)

	return govy.NewRule(func(s string) error {
		if !isISO3166Alpha2(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISO3166Alpha2).
		WithMessageTemplate(tpl).
		WithExamples("US", "PL", "JP").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringISO3166Alpha3 ensures the property's value is a valid ISO 3166-1 alpha-3 country code.
func StringISO3166Alpha3() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISO3166Alpha3Template)

	return govy.NewRule(func(s string) error {
		if !isISO3166Alpha3(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISO3166Alpha3).
		WithMessageTemplate(tpl).
		WithExamples("USA", "POL", "JPN").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringISO3166Numeric ensures the property's value is a valid ISO 3166-1 numeric-3 country code.
func StringISO3166Numeric() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISO3166NumericTemplate)

	return govy.NewRule(func(s string) error {
		if !isISO3166Numeric(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISO3166Numeric).
		WithMessageTemplate(tpl).
		WithExamples("840", "616", "392").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringISO31662 ensures the property's value is a valid ISO 3166-2 country subdivision code.
func StringISO31662() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISO31662Template)

	return govy.NewRule(func(s string) error {
		if !isISO31662(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISO31662).
		WithMessageTemplate(tpl).
		WithExamples("US-CA", "GB-ENG", "PL-14").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringISO4217 ensures the property's value is a valid ISO 4217 three-letter alphabetic currency code.
func StringISO4217() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringISO4217Template)

	return govy.NewRule(func(s string) error {
		if !isISO4217(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringISO4217).
		WithMessageTemplate(tpl).
		WithExamples("USD", "EUR", "JPY").
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringLatitude ensures the property's value is a decimal latitude coordinate between -90 and 90 degrees.
func StringLatitude() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringLatitudeTemplate)

	return govy.NewRule(func(s string) error {
		if !isCoordinate(s, "90") {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringLatitude).
		WithMessageTemplate(tpl).
		WithExamples("0", "-45.25", "90").
		WithDescription("string must be a decimal latitude coordinate between -90 and 90 degrees")
}

// StringLongitude ensures the property's value is a decimal longitude coordinate between -180 and 180 degrees.
func StringLongitude() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringLongitudeTemplate)

	return govy.NewRule(func(s string) error {
		if !isCoordinate(s, "180") {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringLongitude).
		WithMessageTemplate(tpl).
		WithExamples("0", "-122.4194", "180").
		WithDescription("string must be a decimal longitude coordinate between -180 and 180 degrees")
}

func isBCP47LanguageTag(s string) bool {
	if strings.Contains(s, "_") || hasDuplicateBCP47Subtags(s) {
		return false
	}
	base, _, _ := strings.Cut(s, "-")
	if _, ok := bcp47ParserOverrideBases()[strings.ToLower(base)]; ok {
		_, err := language.Default.Parse(withBCP47NeutralParserBase(s))
		return err == nil
	}
	_, err := language.Parse(s)
	return err == nil
}

func isCanonicalBCP47LanguageTag(s string) bool {
	if strings.Contains(s, "_") || hasDuplicateBCP47Subtags(s) {
		return false
	}
	if isCanonicalBCP47GrandfatheredTag(s) {
		return true
	}
	if hasBCP47PreferredTagPrefix(s) {
		return false
	}
	base, _, _ := strings.Cut(s, "-")
	if _, ok := bcp47LanguagePreferredValueBases()[strings.ToLower(base)]; ok {
		return false
	}
	if _, ok := bcp47ParserOverrideBases()[base]; ok {
		neutralTag := withBCP47NeutralParserBase(s)
		tag, err := language.Deprecated.Parse(neutralTag)
		return err == nil && tag.String() == neutralTag
	}
	tag, err := language.Deprecated.Parse(s)
	return err == nil && tag.String() == s
}

func hasBCP47PreferredTagPrefix(s string) bool {
	preferredTags := bcp47TagPreferredValues()
	for {
		if _, ok := preferredTags[s]; ok {
			return true
		}
		separator := strings.LastIndexByte(s, '-')
		if separator < 0 {
			return false
		}
		s = s[:separator]
	}
}

func withBCP47NeutralParserBase(s string) string {
	_, suffix, found := strings.Cut(s, "-")
	if !found {
		return bcp47NeutralParserBase
	}
	return bcp47NeutralParserBase + "-" + suffix
}

// hasDuplicateBCP47Subtags reports duplicates that RFC 5646 prohibits even
// though x/text accepts and normalizes them.
func hasDuplicateBCP47Subtags(s string) bool {
	seenSingletons := make(map[string]struct{})
	seenVariants := make(map[string]struct{})
	inExtension := false
	for index, subtag := range strings.Split(s, "-") {
		subtag = strings.ToLower(subtag)
		if index == 0 {
			if subtag == "x" {
				break
			}
			continue
		}

		if len(subtag) == 1 {
			if subtag == "x" {
				break
			}
			if _, ok := seenSingletons[subtag]; ok {
				return true
			}
			seenSingletons[subtag] = struct{}{}
			inExtension = true
			continue
		}
		if inExtension || !isBCP47VariantSubtag(subtag) {
			continue
		}
		if _, ok := seenVariants[subtag]; ok {
			return true
		}
		seenVariants[subtag] = struct{}{}
	}
	return false
}

func isBCP47VariantSubtag(subtag string) bool {
	return len(subtag) >= 5 && len(subtag) <= 8 ||
		len(subtag) == 4 && subtag[0] >= '0' && subtag[0] <= '9'
}

// isCanonicalBCP47GrandfatheredTag reports registered tags without a
// Preferred-Value that x/text rewrites to private-use representations.
func isCanonicalBCP47GrandfatheredTag(s string) bool {
	_, ok := bcp47GrandfatheredWithoutPreferredValue()[s]
	return ok
}

func isISO3166Alpha2(s string) bool {
	if len(s) != 2 || !isASCIIUpper(s) {
		return false
	}
	region, err := language.ParseRegion(s)
	return err == nil && isISO3166CountryRegion(region) && region.String() == s
}

func isISO3166Alpha3(s string) bool {
	if len(s) != 3 || !isASCIIUpper(s) {
		return false
	}
	region, err := language.ParseRegion(s)
	return err == nil && isISO3166CountryRegion(region) && region.ISO3() == s
}

func isISO3166Numeric(s string) bool {
	if len(s) != 3 || !isASCIIDigit(s) {
		return false
	}
	region, err := language.ParseRegion(s)
	if err != nil || !isISO3166CountryRegion(region) {
		return false
	}
	code, _ := strconv.Atoi(s)
	return region.M49() == code
}

func isISO31662(s string) bool {
	if !iso31662Regexp().MatchString(s) {
		return false
	}
	_, ok := iso31662Codes()[s]
	return ok
}

func isISO4217(s string) bool {
	if len(s) != 3 || !isASCIIUpper(s) {
		return false
	}
	_, ok := iso4217Codes()[s]
	return ok
}

func buildISO4217Codes() map[string]struct{} {
	codes := map[string]struct{}{}
	for iter := textcurrency.Query(textcurrency.NonTender); iter.Next(); {
		codes[iter.Unit().String()] = struct{}{}
	}
	listOneAdditions := [...]string{"MRU", "SLE", "SVC", "UYW", "VED", "VES", "XAD", "XCG", "ZWG"}
	for _, code := range listOneAdditions {
		codes[code] = struct{}{}
	}
	listOneRemovals := [...]string{"ANG", "BGN", "CNH", "CUC", "HRK", "MRO", "SLL", "VEF"}
	for _, code := range listOneRemovals {
		delete(codes, code)
	}
	return codes
}

func lazyLookupMap(build func() map[string]struct{}) func() map[string]struct{} {
	var (
		lookup map[string]struct{}
		once   sync.Once
	)
	return func() map[string]struct{} {
		once.Do(func() {
			lookup = build()
		})
		return lookup
	}
}

func buildStringLookup(data string) map[string]struct{} {
	values := strings.Fields(data)
	lookup := make(map[string]struct{}, len(values))
	for _, value := range values {
		lookup[value] = struct{}{}
	}
	return lookup
}

func isISO3166CountryRegion(region language.Region) bool {
	if !region.IsCountry() {
		return false
	}
	_, ok := iso3166Alpha2Codes()[region.String()]
	return ok
}

func isCoordinate(s, maxMagnitude string) bool {
	if !decimalCoordinateRegexp().MatchString(s) {
		return false
	}
	return decimalMagnitudeAtMost(s, maxMagnitude)
}

func decimalMagnitudeAtMost(value, maximum string) bool {
	value = strings.TrimPrefix(strings.TrimPrefix(value, "-"), "+")
	integer, fraction, _ := strings.Cut(value, ".")
	integer = strings.TrimLeft(integer, "0")
	if integer == "" {
		integer = "0"
	}

	switch {
	case len(integer) < len(maximum):
		return true
	case len(integer) > len(maximum):
		return false
	case integer < maximum:
		return true
	case integer > maximum:
		return false
	default:
		return strings.Trim(fraction, "0") == ""
	}
}

func isASCIIUpper(s string) bool {
	for _, c := range s {
		if c < 'A' || c > 'Z' {
			return false
		}
	}
	return true
}

func isASCIIDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
