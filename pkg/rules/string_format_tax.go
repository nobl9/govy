package rules

import (
	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringEIN ensures the property's value is a US EIN in NN-NNNNNNN format
// with a recognized prefix.
func StringEIN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringEINTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidEIN(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringEIN).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

// StringSSN ensures the property's value is a US SSN in NNN-NN-NNNN format.
// It rejects area 000, 666, and 900-999, group 00, and serial 0000.
func StringSSN() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringSSNTemplate)

	return govy.NewRule(func(s string) error {
		if !isValidSSN(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue: s,
			})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringSSN).
		WithMessageTemplate(tpl).
		WithDescription(mustExecuteTemplate(tpl, govy.TemplateVars{}))
}

func isValidEIN(s string) bool {
	if !einRegexp().MatchString(s) {
		return false
	}
	return isValidEINPrefix(s[:2])
}

func isValidEINPrefix(prefix string) bool {
	switch prefix {
	case "01", "02", "03", "04", "05", "06", "10", "11", "12", "13", "14", "15", "16",
		"20", "21", "22", "23", "24", "25", "26", "27",
		"30", "31", "32", "33", "34", "35", "36", "37", "38", "39",
		"40", "41", "42", "43", "44", "45", "46", "47", "48",
		"50", "51", "52", "53", "54", "55", "56", "57", "58", "59",
		"60", "61", "62", "63", "64", "65", "66", "67", "68",
		"71", "72", "73", "74", "75", "76", "77",
		"80", "81", "82", "83", "84", "85", "86", "87", "88",
		"90", "91", "92", "93", "94", "95", "98", "99":
		return true
	default:
		return false
	}
}

func isValidSSN(s string) bool {
	if !ssnRegexp().MatchString(s) {
		return false
	}
	area := s[:3]
	group := s[4:6]
	serial := s[7:]
	return area != "000" &&
		area != "666" &&
		area[0] != '9' &&
		group != "00" &&
		serial != "0000"
}
