package rules

import (
	"strconv"
	"strings"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringHexColor ensures the property's value is a CSS hexadecimal color.
// It accepts #RGB, #RGBA, #RRGGBB, and #RRGGBBAA forms.
func StringHexColor() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringHexColorTemplate)
	return govy.NewRule(func(s string) error {
		if !hexColorRegexp().MatchString(s) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{PropertyValue: s})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringHexColor).
		WithMessageTemplate(tpl).
		WithDescription("string must be a CSS hex color in #RGB, #RGBA, #RRGGBB, or #RRGGBBAA format")
}

// StringRGB ensures the property's value is a legacy comma-separated rgb(...) color.
// It accepts either three numeric 0-255 components or three percentage 0-100% components.
func StringRGB() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringRGBTemplate)
	return govy.NewRule(func(s string) error {
		args, ok := splitCSSColorFunctionArgs(s, "rgb", 3)
		if !ok || !validRGBComponents(args) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{PropertyValue: s})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringRGB).
		WithMessageTemplate(tpl).
		WithDescription("string must be a legacy comma-separated rgb(...) color with numeric 0-255 or percentage 0-100% components")
}

// StringRGBA ensures the property's value is a legacy comma-separated rgba(...) color.
// It accepts RGB components and an alpha component of 0, 1, or a fractional 0.x value.
func StringRGBA() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringRGBATemplate)
	return govy.NewRule(func(s string) error {
		args, ok := splitCSSColorFunctionArgs(s, "rgba", 4)
		if !ok || !validRGBComponents(args[:3]) || !validAlphaComponent(args[3]) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{PropertyValue: s})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringRGBA).
		WithMessageTemplate(tpl).
		WithDescription("string must be a legacy comma-separated rgba(...) color with numeric 0-255 or percentage 0-100% RGB components and alpha 0, 1, or 0.x")
}

// StringHSL ensures the property's value is a legacy comma-separated hsl(...) color.
// It accepts a hue from 0 to 360 and saturation/lightness percentages from 0% to 100%.
func StringHSL() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringHSLTemplate)
	return govy.NewRule(func(s string) error {
		args, ok := splitCSSColorFunctionArgs(s, "hsl", 3)
		if !ok || !validHSLComponents(args) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{PropertyValue: s})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringHSL).
		WithMessageTemplate(tpl).
		WithDescription("string must be a legacy comma-separated hsl(...) color with hue 0-360 and saturation/lightness 0-100%")
}

// StringHSLA ensures the property's value is a legacy comma-separated hsla(...) color.
// It accepts HSL components and an alpha component of 0, 1, or a fractional 0.x value.
func StringHSLA() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringHSLATemplate)
	return govy.NewRule(func(s string) error {
		args, ok := splitCSSColorFunctionArgs(s, "hsla", 4)
		if !ok || !validHSLComponents(args[:3]) || !validAlphaComponent(args[3]) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{PropertyValue: s})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringHSLA).
		WithMessageTemplate(tpl).
		WithDescription("string must be a legacy comma-separated hsla(...) color with hue 0-360, saturation/lightness 0-100%, and alpha 0, 1, or 0.x")
}

// StringCMYK ensures the property's value is a cmyk(...) color.
// It accepts four percentage components from 0% to 100%.
func StringCMYK() govy.Rule[string] {
	tpl := messagetemplates.Get(messagetemplates.StringCMYKTemplate)
	return govy.NewRule(func(s string) error {
		args, ok := splitCSSColorFunctionArgs(s, "cmyk", 4)
		if !ok || !allCSSColorComponents(args, validPercentComponent) {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{PropertyValue: s})
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringCMYK).
		WithMessageTemplate(tpl).
		WithDescription("string must be a cmyk(...) color with four percentage components from 0% to 100%")
}

func splitCSSColorFunctionArgs(s, name string, count int) ([]string, bool) {
	prefix := name + "("
	if !strings.HasPrefix(s, prefix) || !strings.HasSuffix(s, ")") {
		return nil, false
	}
	body := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(s, prefix), ")"))
	parts := strings.Split(body, ",")
	if len(parts) != count {
		return nil, false
	}
	args := make([]string, 0, count)
	for _, part := range parts {
		arg := strings.TrimSpace(part)
		if arg == "" {
			return nil, false
		}
		args = append(args, arg)
	}
	return args, true
}

func validRGBComponents(args []string) bool {
	return allCSSColorComponents(args, validRGBNumericComponent) ||
		allCSSColorComponents(args, validPercentComponent)
}

func validHSLComponents(args []string) bool {
	return validIntComponent(args[0], 0, 360) &&
		validPercentComponent(args[1]) &&
		validPercentComponent(args[2])
}

func allCSSColorComponents(args []string, valid func(string) bool) bool {
	for _, arg := range args {
		if !valid(arg) {
			return false
		}
	}
	return true
}

func validRGBNumericComponent(s string) bool {
	return validIntComponent(s, 0, 255)
}

func validPercentComponent(s string) bool {
	value, ok := strings.CutSuffix(s, "%")
	return ok && validIntComponent(value, 0, 100)
}

func validAlphaComponent(s string) bool {
	if s == "0" || s == "1" {
		return true
	}
	value, ok := strings.CutPrefix(s, "0.")
	return ok && isASCIIDigits(value)
}

func validIntComponent(s string, minValue, maxValue int) bool {
	if !isASCIIDigits(s) {
		return false
	}
	n, err := strconv.Atoi(s)
	return err == nil && n >= minValue && n <= maxValue
}

func isASCIIDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
