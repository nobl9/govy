package rules

import (
	"strconv"
	"strings"
	"sync"

	textcurrency "golang.org/x/text/currency"
	"golang.org/x/text/language"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

var (
	iso31662Regexp          = lazyRegexCompile(`^[A-Z]{2}-[A-Z0-9]{1,3}$`)
	decimalCoordinateRegexp = lazyRegexCompile(`^[+-]?(?:\d+(?:\.\d+)?|\.\d+)$`)
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

// StringISO3166Numeric ensures the property's value is a valid ISO 3166-1 numeric country code.
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

// StringISO31662 ensures the property's value has ISO 3166-2 subdivision code syntax
// and a valid ISO 3166-1 alpha-2 country prefix.
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
		WithDescription("string must have ISO 3166-2 subdivision code syntax with a valid ISO 3166-1 alpha-2 country prefix")
}

// StringISO4217 ensures the property's value is a valid ISO 4217 currency code.
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
		if !isCoordinate(s, -90, 90) {
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
		if !isCoordinate(s, -180, 180) {
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
	if strings.Contains(s, "_") {
		return false
	}
	_, err := language.Parse(s)
	return err == nil
}

func isCanonicalBCP47LanguageTag(s string) bool {
	if strings.Contains(s, "_") {
		return false
	}
	tag, err := language.BCP47.Parse(s)
	return err == nil && tag.String() == s
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
	country, _, _ := strings.Cut(s, "-")
	return isISO3166Alpha2(country)
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

func isISO3166CountryRegion(region language.Region) bool {
	if !region.IsCountry() {
		return false
	}
	_, ok := iso3166Alpha2Codes()[region.String()]
	return ok
}

func isCoordinate(s string, minValue, maxValue float64) bool {
	if !decimalCoordinateRegexp().MatchString(s) {
		return false
	}
	value, err := strconv.ParseFloat(s, 64)
	return err == nil && value >= minValue && value <= maxValue
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
