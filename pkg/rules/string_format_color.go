package rules

import (
	"math"
	"strconv"
	"strings"

	"github.com/nobl9/govy/internal/messagetemplates"
	"github.com/nobl9/govy/pkg/govy"
)

// StringColorOption configures [StringRGB], [StringHSL], or [StringDeviceCMYK].
type StringColorOption interface {
	apply(*stringColorConfig)
}

type stringColorOptionFunc func(*stringColorConfig)

func (f stringColorOptionFunc) apply(config *stringColorConfig) {
	f(config)
}

type stringColorConfig struct {
	legacySyntaxOnly bool
}

type stringColorTemplateVars struct {
	LegacySyntaxOnly bool
}

// StringColorLegacySyntaxOnly restricts a string color format rule to the
// comma-separated legacy syntax defined by CSS Color Modules Level 4 and 5.
func StringColorLegacySyntaxOnly() StringColorOption {
	return stringColorOptionFunc(func(config *stringColorConfig) {
		config.legacySyntaxOnly = true
	})
}

// StringHexColor ensures the property's value is a CSS hexadecimal color.
// It accepts the #RGB, #RGBA, #RRGGBB, and #RRGGBBAA forms defined by
// [CSS Color Module Level 4 hexadecimal notation].
//
// [CSS Color Module Level 4 hexadecimal notation]: https://www.w3.org/TR/css-color-4/#hex-notation
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
		WithDescription("string must be a CSS hexadecimal color in #RGB, #RGBA, #RRGGBB, or #RRGGBBAA format")
}

// StringRGB ensures the property's value is an rgb(...) or rgba(...) color
// using the modern or legacy syntax defined by [CSS Color Module Level 4].
// Pass [StringColorLegacySyntaxOnly] to accept only the comma-separated legacy syntax.
//
// StringRGB validates concrete numeric literals rather than arbitrary CSS.
// It rejects none, calculated values, and out-of-range components instead of
// applying CSS clamping during value computation.
//
// [CSS Color Module Level 4]: https://www.w3.org/TR/css-color-4/#rgb-functions
func StringRGB(options ...StringColorOption) govy.Rule[string] {
	config := newStringColorConfig(options)
	tpl := messagetemplates.Get(messagetemplates.StringRGBTemplate)
	return govy.NewRule(func(s string) error {
		if !validRGBColor(s, config.legacySyntaxOnly) {
			return newStringColorRuleError(s, config)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringRGB).
		WithMessageTemplate(tpl).
		WithDescription(stringRGBDescription(config))
}

// StringHSL ensures the property's value is an hsl(...) or hsla(...) color
// using the modern or legacy syntax defined by [CSS Color Module Level 4].
// Pass [StringColorLegacySyntaxOnly] to accept only the comma-separated legacy syntax.
//
// StringHSL validates concrete numeric literals rather than arbitrary CSS.
// It rejects none, calculated values, and out-of-range saturation, lightness,
// and alpha components instead of applying CSS clamping during value computation.
// Hue values are cyclic and therefore are not range-restricted.
//
// [CSS Color Module Level 4]: https://www.w3.org/TR/css-color-4/#the-hsl-notation
func StringHSL(options ...StringColorOption) govy.Rule[string] {
	config := newStringColorConfig(options)
	tpl := messagetemplates.Get(messagetemplates.StringHSLTemplate)
	return govy.NewRule(func(s string) error {
		if !validHSLColor(s, config.legacySyntaxOnly) {
			return newStringColorRuleError(s, config)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringHSL).
		WithMessageTemplate(tpl).
		WithDescription(stringHSLDescription(config))
}

// StringDeviceCMYK ensures the property's value is a device-cmyk(...) color
// using the modern or legacy syntax defined by [CSS Color Module Level 5].
// Pass [StringColorLegacySyntaxOnly] to accept only the comma-separated legacy syntax.
//
// StringDeviceCMYK validates concrete numeric literals rather than arbitrary CSS.
// It rejects none, calculated values, fallback colors, and out-of-range components
// instead of applying CSS clamping during value computation.
//
// [CSS Color Module Level 5]: https://www.w3.org/TR/css-color-5/#device-CMYK
func StringDeviceCMYK(options ...StringColorOption) govy.Rule[string] {
	config := newStringColorConfig(options)
	tpl := messagetemplates.Get(messagetemplates.StringDeviceCMYKTemplate)
	return govy.NewRule(func(s string) error {
		if !validDeviceCMYKColor(s, config.legacySyntaxOnly) {
			return newStringColorRuleError(s, config)
		}
		return nil
	}).
		WithErrorCode(ErrorCodeStringDeviceCMYK).
		WithMessageTemplate(tpl).
		WithDescription(stringDeviceCMYKDescription(config))
}

func newStringColorConfig(options []StringColorOption) stringColorConfig {
	config := stringColorConfig{}
	for _, option := range options {
		if option != nil {
			option.apply(&config)
		}
	}
	return config
}

func newStringColorRuleError(s string, config stringColorConfig) error {
	return govy.NewRuleErrorTemplate(govy.TemplateVars{
		PropertyValue: s,
		Custom: stringColorTemplateVars{
			LegacySyntaxOnly: config.legacySyntaxOnly,
		},
	})
}

func stringRGBDescription(config stringColorConfig) string {
	if config.legacySyntaxOnly {
		return "string must be a CSS rgb(...) or rgba(...) color in legacy comma-separated syntax with in-range literal components"
	}
	return "string must be a CSS rgb(...) or rgba(...) color in modern or legacy syntax with in-range literal components"
}

func stringHSLDescription(config stringColorConfig) string {
	if config.legacySyntaxOnly {
		return "string must be a CSS hsl(...) or hsla(...) color in legacy comma-separated syntax with in-range literal components"
	}
	return "string must be a CSS hsl(...) or hsla(...) color in modern or legacy syntax with in-range literal components"
}

func stringDeviceCMYKDescription(config stringColorConfig) string {
	if config.legacySyntaxOnly {
		return "string must be a CSS device-cmyk(...) color in legacy comma-separated syntax with in-range literal components"
	}
	return "string must be a CSS device-cmyk(...) color in modern or legacy syntax with in-range literal components"
}

func validRGBColor(s string, legacyOnly bool) bool {
	body, ok := splitCSSColorFunction(s, "rgb", "rgba")
	if !ok {
		return false
	}
	if strings.Contains(body, ",") {
		return validLegacyRGB(body)
	}
	return !legacyOnly && validModernRGB(body)
}

func validLegacyRGB(body string) bool {
	components, ok := splitLegacyColorComponents(body, 3, 4)
	if !ok {
		return false
	}
	rgb := components[:3]
	if !allCSSColorComponents(rgb, validRGBNumber) &&
		!allCSSColorComponents(rgb, validPercentage) {
		return false
	}
	return len(components) == 3 || validAlpha(components[3])
}

func validModernRGB(body string) bool {
	components, alpha, ok := splitModernColorComponents(body, 3)
	if !ok || !allCSSColorComponents(components, validRGBComponent) {
		return false
	}
	return alpha == "" || validAlpha(alpha)
}

func validHSLColor(s string, legacyOnly bool) bool {
	body, ok := splitCSSColorFunction(s, "hsl", "hsla")
	if !ok {
		return false
	}
	if strings.Contains(body, ",") {
		return validLegacyHSL(body)
	}
	return !legacyOnly && validModernHSL(body)
}

func validLegacyHSL(body string) bool {
	components, ok := splitLegacyColorComponents(body, 3, 4)
	if !ok || !validHue(components[0]) ||
		!validPercentage(components[1]) ||
		!validPercentage(components[2]) {
		return false
	}
	return len(components) == 3 || validAlpha(components[3])
}

func validModernHSL(body string) bool {
	components, alpha, ok := splitModernColorComponents(body, 3)
	if !ok || !validHue(components[0]) ||
		!validHSLPercentage(components[1]) ||
		!validHSLPercentage(components[2]) {
		return false
	}
	return alpha == "" || validAlpha(alpha)
}

func validDeviceCMYKColor(s string, legacyOnly bool) bool {
	body, ok := splitCSSColorFunction(s, "device-cmyk")
	if !ok {
		return false
	}
	if strings.Contains(body, ",") {
		return validLegacyDeviceCMYK(body)
	}
	return !legacyOnly && validModernDeviceCMYK(body)
}

func validLegacyDeviceCMYK(body string) bool {
	components, ok := splitLegacyColorComponents(body, 4)
	return ok && allCSSColorComponents(components, validUnitIntervalNumber)
}

func validModernDeviceCMYK(body string) bool {
	components, alpha, ok := splitModernColorComponents(body, 4)
	if !ok || !allCSSColorComponents(components, validDeviceCMYKComponent) {
		return false
	}
	return alpha == "" || validAlpha(alpha)
}

func splitCSSColorFunction(s string, names ...string) (string, bool) {
	open := strings.IndexByte(s, '(')
	if open <= 0 || !strings.HasSuffix(s, ")") {
		return "", false
	}
	name := s[:open]
	validName := false
	for _, expected := range names {
		if equalASCIIFold(name, expected) {
			validName = true
			break
		}
	}
	if !validName {
		return "", false
	}
	return s[open+1 : len(s)-1], true
}

func splitLegacyColorComponents(body string, counts ...int) ([]string, bool) {
	parts := strings.Split(body, ",")
	validCount := false
	for _, count := range counts {
		if len(parts) == count {
			validCount = true
			break
		}
	}
	if !validCount {
		return nil, false
	}
	components := make([]string, 0, len(parts))
	for _, part := range parts {
		component := trimCSSWhitespace(part)
		if component == "" {
			return nil, false
		}
		components = append(components, component)
	}
	return components, true
}

func splitModernColorComponents(
	body string,
	count int,
) (components []string, alpha string, ok bool) {
	position := skipCSSWhitespace(body, 0)
	components = make([]string, 0, count)
	for index := range count {
		component, next := scanCSSColorComponent(body, position)
		if component == "" {
			return nil, "", false
		}
		components = append(components, component)
		position = next
		beforeWhitespace := position
		position = skipCSSWhitespace(body, position)
		if index < count-1 && position == beforeWhitespace {
			return nil, "", false
		}
	}
	if position == len(body) {
		return components, "", true
	}
	if body[position] != '/' {
		return nil, "", false
	}
	position = skipCSSWhitespace(body, position+1)
	alpha, position = scanCSSColorComponent(body, position)
	if alpha == "" {
		return nil, "", false
	}
	position = skipCSSWhitespace(body, position)
	return components, alpha, position == len(body)
}

func scanCSSColorComponent(s string, position int) (component string, next int) {
	start := position
	for position < len(s) &&
		!isCSSWhitespace(s[position]) &&
		s[position] != '/' &&
		s[position] != ',' {
		position++
	}
	return s[start:position], position
}

func skipCSSWhitespace(s string, position int) int {
	for position < len(s) && isCSSWhitespace(s[position]) {
		position++
	}
	return position
}

func trimCSSWhitespace(s string) string {
	start := skipCSSWhitespace(s, 0)
	end := len(s)
	for end > start && isCSSWhitespace(s[end-1]) {
		end--
	}
	return s[start:end]
}

func isCSSWhitespace(b byte) bool {
	switch b {
	case '\t', '\n', '\f', '\r', ' ':
		return true
	default:
		return false
	}
}

func allCSSColorComponents(components []string, valid func(string) bool) bool {
	for _, component := range components {
		if !valid(component) {
			return false
		}
	}
	return true
}

func validRGBComponent(s string) bool {
	return validRGBNumber(s) || validPercentage(s)
}

func validRGBNumber(s string) bool {
	return validCSSNumberMax(s, 255)
}

func validHSLPercentage(s string) bool {
	return validCSSNumberMax(s, 100) || validPercentage(s)
}

func validDeviceCMYKComponent(s string) bool {
	return validUnitIntervalNumber(s) || validPercentage(s)
}

func validUnitIntervalNumber(s string) bool {
	return validCSSNumberMax(s, 1)
}

func validPercentage(s string) bool {
	value, ok := strings.CutSuffix(s, "%")
	return ok && validCSSNumberMax(value, 100)
}

func validAlpha(s string) bool {
	if value, ok := strings.CutSuffix(s, "%"); ok {
		return validCSSNumberMax(value, 100)
	}
	return validCSSNumberMax(s, 1)
}

func validHue(s string) bool {
	if _, ok := parseCSSNumber(s); ok {
		return true
	}
	for _, unit := range []string{"grad", "turn", "deg", "rad"} {
		if len(s) > len(unit) && equalASCIIFold(s[len(s)-len(unit):], unit) {
			_, ok := parseCSSNumber(s[:len(s)-len(unit)])
			return ok
		}
	}
	return false
}

func validCSSNumberMax(s string, maxValue float64) bool {
	value, ok := parseCSSNumber(s)
	return ok && value >= 0 && value <= maxValue
}

func parseCSSNumber(s string) (float64, bool) {
	if s == "" {
		return 0, false
	}
	position := 0
	if s[position] == '+' || s[position] == '-' {
		position++
		if position == len(s) {
			return 0, false
		}
	}
	numberStart := position
	digitsBefore := scanASCIIDigits(s, position)
	position = digitsBefore
	if position < len(s) && s[position] == '.' {
		position++
		digitsAfter := scanASCIIDigits(s, position)
		if digitsAfter == position {
			return 0, false
		}
		position = digitsAfter
	} else if digitsBefore == numberStart {
		return 0, false
	}
	if position < len(s) && (s[position] == 'e' || s[position] == 'E') {
		position++
		if position < len(s) && (s[position] == '+' || s[position] == '-') {
			position++
		}
		exponentEnd := scanASCIIDigits(s, position)
		if exponentEnd == position {
			return 0, false
		}
		position = exponentEnd
	}
	if position != len(s) {
		return 0, false
	}
	value, err := strconv.ParseFloat(s, 64)
	return value, err == nil && !math.IsInf(value, 0) && !math.IsNaN(value)
}

func scanASCIIDigits(s string, position int) int {
	for position < len(s) && s[position] >= '0' && s[position] <= '9' {
		position++
	}
	return position
}

func equalASCIIFold(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range len(a) {
		if lowerASCII(a[i]) != lowerASCII(b[i]) {
			return false
		}
	}
	return true
}

func lowerASCII(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}
