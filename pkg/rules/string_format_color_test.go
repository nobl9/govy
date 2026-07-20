package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

type stringFormatColorTestCase struct {
	in      string
	options []StringColorOption
}

const (
	errStringHexColor         = "string must be a valid CSS hexadecimal color"
	errStringRGB              = "string must be a valid CSS rgb(...) or rgba(...) color"
	errStringRGBLegacy        = "string must be a valid legacy comma-separated CSS rgb(...) or rgba(...) color"
	errStringHSL              = "string must be a valid CSS hsl(...) or hsla(...) color"
	errStringHSLLegacy        = "string must be a valid legacy comma-separated CSS hsl(...) or hsla(...) color"
	errStringDeviceCMYK       = "string must be a valid CSS device-cmyk(...) color"
	errStringDeviceCMYKLegacy = "string must be a valid legacy comma-separated CSS device-cmyk(...) color"
)

func TestStringColorLegacySyntaxOnly(t *testing.T) {
	t.Parallel()
	config := newStringColorConfig([]StringColorOption{StringColorLegacySyntaxOnly()})
	assert.True(t, config.legacySyntaxOnly)
}

func BenchmarkStringColorLegacySyntaxOnly(b *testing.B) {
	for b.Loop() {
		_ = StringColorLegacySyntaxOnly()
	}
}

var (
	stringHexColorValidTestCases = map[string]stringFormatColorTestCase{
		"three digits":           {in: "#000"},
		"three uppercase digits": {in: "#ABC"},
		"four digits":            {in: "#abcd"},
		"six digits":             {in: "#A1B2C3"},
		"eight digits":           {in: "#AABBCCDD"},
	}
	stringHexColorInvalidTestCases = map[string]stringFormatColorTestCase{
		"empty":                       {in: ""},
		"missing hash":                {in: "112233"},
		"two digits":                  {in: "#12"},
		"five digits":                 {in: "#12345"},
		"seven digits":                {in: "#1234567"},
		"nine digits":                 {in: "#123456789"},
		"non-hex digit":               {in: "#ggg"},
		"trailing newline":            {in: "#123456\n"},
		"leading external whitespace": {in: " #123"},
	}
)

func TestStringHexColor(t *testing.T) {
	t.Parallel()
	rule := StringHexColor()
	assertStringFormatColorRuleTestCases(
		t,
		rule,
		stringHexColorValidTestCases,
		"",
		ErrorCodeStringHexColor,
	)
	assertStringFormatColorRuleTestCases(
		t,
		rule,
		stringHexColorInvalidTestCases,
		errStringHexColor,
		ErrorCodeStringHexColor,
	)
}

func BenchmarkStringHexColor(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringHexColor(), "#112233")
}

var (
	stringRGBValidTestCases = map[string]stringFormatColorTestCase{
		"legacy numeric":             {in: "rgb(0, 127.5, 255)"},
		"legacy percentage":          {in: "RGB(+0%, 5e1%, 100%)"},
		"legacy rgba alpha":          {in: "rgba(12, 34, 56, .25)"},
		"legacy rgba optional alpha": {in: "rgba(12, 34, 56)"},
		"modern numeric":             {in: "rgb(0 127.5 255)"},
		"modern mixed components":    {in: "rgb(0 50% 2.55e2)"},
		"modern alpha":               {in: "rgb(0 0 0 / 50%)"},
		"modern rgba alias":          {in: "RGBA(100% 0% 33.3% / .5)"},
		"modern CSS whitespace":      {in: "rgb(\t0\n50%\f255\r/\t1 )"},
		"nil option":                 {in: "rgb(0 0 0)", options: []StringColorOption{nil}},
		"legacy-only numeric":        {in: "rgb(+0, 1.27e2, 255)", options: legacyColorOptions()},
		"legacy-only rgba alpha":     {in: "rgba(0%, 50%, 100%, 1)", options: legacyColorOptions()},
	}
	stringRGBInvalidLegacyTestCases = map[string]stringFormatColorTestCase{
		"legacy-only rejects modern": {
			in:      "rgb(0 0 0)",
			options: legacyColorOptions(),
		},
	}
	stringRGBInvalidTestCases = map[string]stringFormatColorTestCase{
		"empty":                        {in: ""},
		"wrong function":               {in: "hsl(0 0 0)"},
		"component above number range": {in: "rgb(256 0 0)"},
		"component below number range": {in: "rgb(-1 0 0)"},
		"percentage above range":       {in: "rgb(100.1% 0% 0%)"},
		"alpha above range":            {in: "rgb(0 0 0 / 1.1)"},
		"alpha below range":            {in: "rgb(0 0 0 / -1%)"},
		"legacy mixed components":      {in: "rgb(0, 0%, 0)"},
		"legacy slash alpha":           {in: "rgb(0, 0, 0 / .5)"},
		"modern comma alpha":           {in: "rgb(0 0 0, .5)"},
		"missing modern separator":     {in: "rgb(0 0 0.5.5)"},
		"missing component":            {in: "rgb(0 0)"},
		"extra component":              {in: "rgb(0 0 0 0)"},
		"trailing decimal point":       {in: "rgb(0. 0 0)"},
		"malformed exponent":           {in: "rgb(1e 0 0)"},
		"non-finite exponent":          {in: "rgb(1e999 0 0)"},
		"none component":               {in: "rgb(none 0 0)"},
		"calculated component":         {in: "rgb(calc(1) 0 0)"},
		"variable component":           {in: "rgb(var(--red) 0 0)"},
		"CSS comment":                  {in: "rgb(0/**/ 0 0)"},
		"escaped function name":        {in: "r\\67 b(0 0 0)"},
		"leading external whitespace":  {in: " rgb(0 0 0)"},
		"trailing external whitespace": {in: "rgb(0 0 0) "},
		"non-CSS whitespace":           {in: "rgb(0\u00a00\u00a00)"},
	}
)

func TestStringRGB(t *testing.T) {
	t.Parallel()
	assertStringColorRuleTestCases(t, StringRGB, stringRGBValidTestCases, "", ErrorCodeStringRGB)
	assertStringColorRuleTestCases(
		t,
		StringRGB,
		stringRGBInvalidLegacyTestCases,
		errStringRGBLegacy,
		ErrorCodeStringRGB,
	)
	assertStringColorRuleTestCases(
		t,
		StringRGB,
		stringRGBInvalidTestCases,
		errStringRGB,
		ErrorCodeStringRGB,
	)
}

func BenchmarkStringRGB(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringRGB(), "rgb(12 34 56 / .25)")
}

var (
	stringHSLValidTestCases = map[string]stringFormatColorTestCase{
		"legacy percentage":       {in: "hsl(120, 50%, 25%)"},
		"legacy angle hue":        {in: "HSLA(-.5turn, 50%, 25%, .5)"},
		"legacy cyclic hue":       {in: "hsl(720, 50%, 25%)"},
		"modern percentage":       {in: "hsl(120 50% 25%)"},
		"modern numeric":          {in: "hsl(1.2e2 50 25)"},
		"modern mixed components": {in: "hsl(120deg 50 25%)"},
		"modern alpha":            {in: "hsla(120 50% 25% / 50%)"},
		"grad hue":                {in: "hsl(100grad 50% 25%)"},
		"radian hue":              {in: "hsl(3.14rad 50% 25%)"},
		"turn hue":                {in: "hsl(.5TURN 50% 25%)"},
		"legacy-only percentage":  {in: "hsl(-120deg, 50%, 25%)", options: legacyColorOptions()},
		"legacy-only alpha":       {in: "hsla(120, 50%, 25%, 25%)", options: legacyColorOptions()},
	}
	stringHSLInvalidLegacyTestCases = map[string]stringFormatColorTestCase{
		"legacy-only rejects modern": {
			in:      "hsl(120 50% 25%)",
			options: legacyColorOptions(),
		},
	}
	stringHSLInvalidTestCases = map[string]stringFormatColorTestCase{
		"empty":                        {in: ""},
		"wrong function":               {in: "rgb(0 0 0)"},
		"legacy numeric saturation":    {in: "hsl(120, 50, 25%)"},
		"saturation above range":       {in: "hsl(120 101% 25%)"},
		"lightness below range":        {in: "hsl(120 50% -1)"},
		"alpha above range":            {in: "hsl(120 50% 25% / 101%)"},
		"unknown angle unit":           {in: "hsl(120foo 50% 25%)"},
		"space before angle unit":      {in: "hsl(120 deg 50% 25%)"},
		"trailing hue decimal point":   {in: "hsl(120. 50% 25%)"},
		"none hue":                     {in: "hsl(none 50% 25%)"},
		"none saturation":              {in: "hsl(120 none 25%)"},
		"calculated hue":               {in: "hsl(calc(120) 50% 25%)"},
		"legacy mixed delimiters":      {in: "hsl(120, 50%, 25% / .5)"},
		"modern comma delimiter":       {in: "hsl(120 50%, 25%)"},
		"leading external whitespace":  {in: " hsl(120 50% 25%)"},
		"trailing external whitespace": {in: "hsl(120 50% 25%) "},
	}
)

func TestStringHSL(t *testing.T) {
	t.Parallel()
	assertStringColorRuleTestCases(t, StringHSL, stringHSLValidTestCases, "", ErrorCodeStringHSL)
	assertStringColorRuleTestCases(
		t,
		StringHSL,
		stringHSLInvalidLegacyTestCases,
		errStringHSLLegacy,
		ErrorCodeStringHSL,
	)
	assertStringColorRuleTestCases(
		t,
		StringHSL,
		stringHSLInvalidTestCases,
		errStringHSL,
		ErrorCodeStringHSL,
	)
}

func BenchmarkStringHSL(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringHSL(), "hsl(120 50% 25% / .5)")
}

var (
	stringDeviceCMYKValidTestCases = map[string]stringFormatColorTestCase{
		"legacy numeric":          {in: "device-cmyk(0, .25, 5e-1, 1)"},
		"legacy uppercase":        {in: "DEVICE-CMYK(0, 0, 0, 0)"},
		"modern numeric":          {in: "device-cmyk(0 .25 .5 1)"},
		"modern percentage":       {in: "device-cmyk(0% 25% 50% 100%)"},
		"modern mixed components": {in: "device-cmyk(0 25% .5 100%)"},
		"modern alpha":            {in: "device-cmyk(0 0 0 1 / 50%)"},
		"legacy-only numeric":     {in: "device-cmyk(0, .25, .5, 1)", options: legacyColorOptions()},
	}
	stringDeviceCMYKInvalidLegacyTestCases = map[string]stringFormatColorTestCase{
		"legacy-only rejects modern": {
			in:      "device-cmyk(0 0 0 1)",
			options: legacyColorOptions(),
		},
	}
	stringDeviceCMYKInvalidTestCases = map[string]stringFormatColorTestCase{
		"empty":                        {in: ""},
		"non-standard function":        {in: "cmyk(0 0 0 1)"},
		"legacy percentage":            {in: "device-cmyk(0%, 25%, 50%, 100%)"},
		"legacy alpha":                 {in: "device-cmyk(0, 0, 0, 1, .5)"},
		"number above range":           {in: "device-cmyk(0 0 0 1.1)"},
		"number below range":           {in: "device-cmyk(-.1 0 0 1)"},
		"percentage above range":       {in: "device-cmyk(0% 0% 0% 101%)"},
		"alpha above range":            {in: "device-cmyk(0 0 0 1 / 1.1)"},
		"missing component":            {in: "device-cmyk(0 0 1)"},
		"extra component":              {in: "device-cmyk(0 0 0 1 0)"},
		"none component":               {in: "device-cmyk(none 0 0 1)"},
		"calculated component":         {in: "device-cmyk(calc(0) 0 0 1)"},
		"fallback color":               {in: "device-cmyk(0 0 0 1, black)"},
		"leading external whitespace":  {in: " device-cmyk(0 0 0 1)"},
		"trailing external whitespace": {in: "device-cmyk(0 0 0 1) "},
	}
)

func TestStringDeviceCMYK(t *testing.T) {
	t.Parallel()
	assertStringColorRuleTestCases(
		t,
		StringDeviceCMYK,
		stringDeviceCMYKValidTestCases,
		"",
		ErrorCodeStringDeviceCMYK,
	)
	assertStringColorRuleTestCases(
		t,
		StringDeviceCMYK,
		stringDeviceCMYKInvalidLegacyTestCases,
		errStringDeviceCMYKLegacy,
		ErrorCodeStringDeviceCMYK,
	)
	assertStringColorRuleTestCases(
		t,
		StringDeviceCMYK,
		stringDeviceCMYKInvalidTestCases,
		errStringDeviceCMYK,
		ErrorCodeStringDeviceCMYK,
	)
}

func BenchmarkStringDeviceCMYK(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringDeviceCMYK(), "device-cmyk(0 25% .5 100% / .5)")
}

func assertStringColorRuleTestCases(
	t *testing.T,
	ruleFactory func(...StringColorOption) govy.Rule[string],
	testCases map[string]stringFormatColorTestCase,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	if expectedError != "" {
		for _, tc := range testCases {
			err := ruleFactory(tc.options...).Validate(tc.in)
			assert.EqualError(t, err, expectedError)
			assert.True(t, govy.HasErrorCode(err, errorCode))
			break
		}
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if expectedError != "" {
				assert.Error(t, ruleFactory(tc.options...).Validate(tc.in))
				return
			}
			assert.NoError(t, ruleFactory(tc.options...).Validate(tc.in))
		})
	}
}

func assertStringFormatColorRuleTestCases(
	t *testing.T,
	rule govy.Rule[string],
	testCases map[string]stringFormatColorTestCase,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	if expectedError != "" {
		for _, tc := range testCases {
			err := rule.Validate(tc.in)
			assert.EqualError(t, err, expectedError)
			assert.True(t, govy.HasErrorCode(err, errorCode))
			break
		}
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if expectedError != "" {
				assert.Error(t, rule.Validate(tc.in))
				return
			}
			assert.NoError(t, rule.Validate(tc.in))
		})
	}
}

func benchmarkStringFormatColorRule(
	b *testing.B,
	rule govy.Rule[string],
	in string,
) {
	b.Helper()
	for b.Loop() {
		_ = rule.Validate(in)
	}
}

func legacyColorOptions() []StringColorOption {
	return []StringColorOption{
		StringColorLegacySyntaxOnly(),
		nil,
		StringColorLegacySyntaxOnly(),
	}
}
