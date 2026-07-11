package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

type stringFormatColorTestCase struct {
	in            string
	expectedError string
}

const (
	errStringHexColor = "string must be a valid CSS hex color"
	errStringRGB      = "string must be a valid legacy comma-separated rgb(...) color"
	errStringRGBA     = "string must be a valid legacy comma-separated rgba(...) color"
	errStringHSL      = "string must be a valid legacy comma-separated hsl(...) color"
	errStringHSLA     = "string must be a valid legacy comma-separated hsla(...) color"
	errStringCMYK     = "string must be a valid comma-separated cmyk(...) color"
)

var stringHexColorTestCases = []*stringFormatColorTestCase{
	{in: "#000"},
	{in: "#fff"},
	{in: "#ABC"},
	{in: "#0000"},
	{in: "#abcd"},
	{in: "#112233"},
	{in: "#A1B2C3"},
	{in: "#11223344"},
	{in: "#AABBCCDD"},
	{in: "", expectedError: errStringHexColor},
	{in: "112233", expectedError: errStringHexColor},
	{in: "#12", expectedError: errStringHexColor},
	{in: "#12345", expectedError: errStringHexColor},
	{in: "#1234567", expectedError: errStringHexColor},
	{in: "#123456789", expectedError: errStringHexColor},
	{in: "#ggg", expectedError: errStringHexColor},
	{in: "#12x", expectedError: errStringHexColor},
	{in: "#123456\n", expectedError: errStringHexColor},
	{in: " #123", expectedError: errStringHexColor},
}

var stringRGBTestCases = []*stringFormatColorTestCase{
	{in: "rgb(0,0,0)"},
	{in: "RGB(0,0,0)"},
	{in: "rgb(255, 255, 255)"},
	{in: "rgb(001, 002, 003)"},
	{in: "rgb(100%, 0%, 50%)"},
	{in: "rGb(33.3%, 100.0%, .5%)"},
	{in: "rgb( 12 , 34 , 56 )"},
	{in: "", expectedError: errStringRGB},
	{in: "rgba(0,0,0)", expectedError: errStringRGB},
	{in: "rgb()", expectedError: errStringRGB},
	{in: "rgb(256,0,0)", expectedError: errStringRGB},
	{in: "rgb(-1,0,0)", expectedError: errStringRGB},
	{in: "rgb(101%,0%,0%)", expectedError: errStringRGB},
	{in: "rgb(100.1%,0%,0%)", expectedError: errStringRGB},
	{in: "rgb(0.%,0%,0%)", expectedError: errStringRGB},
	{in: "rgb(0,0)", expectedError: errStringRGB},
	{in: "rgb(0,0,0,0)", expectedError: errStringRGB},
	{in: "rgb(0.5,0,0)", expectedError: errStringRGB},
	{in: "rgb(0, 0%, 0)", expectedError: errStringRGB},
	{in: "rgb(0 0 0)", expectedError: errStringRGB},
}

var stringRGBATestCases = []*stringFormatColorTestCase{
	{in: "rgba(0,0,0,0)"},
	{in: "RGBA(0,0,0,1.0)"},
	{in: "rgba(255, 255, 255, 1)"},
	{in: "rgba(100%, 0%, 50%, 0.5)"},
	{in: "rgba(100%, 0%, 33.3%, 50%)"},
	{in: "rgba(0%, 0%, .5%, 100.0%)"},
	{in: "rgba(12, 34, 56, 0.25)"},
	{in: "rGbA(12, 34, 56, .5)"},
	{in: "rgba(12, 34, 56, 0.)"},
	{in: "", expectedError: errStringRGBA},
	{in: "rgba(0,0,0)", expectedError: errStringRGBA},
	{in: "rgb(0,0,0)", expectedError: errStringRGBA},
	{in: "rgba(256,0,0,1)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,2)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,1.1)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,-0.1)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,-1%)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,100.1%)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,100.%)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,.)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,1..0)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,50%%)", expectedError: errStringRGBA},
	{in: "rgba(0,0,0,50px)", expectedError: errStringRGBA},
	{in: "rgba(101%,0%,0%,0.5)", expectedError: errStringRGBA},
	{in: "rgba(0, 0%, 0, 0.5)", expectedError: errStringRGBA},
}

var stringHSLTestCases = []*stringFormatColorTestCase{
	{in: "hsl(0,0%,0%)"},
	{in: "HsL(0,0%,0%)"},
	{in: "hsl(360, 100%, 100%)"},
	{in: "hsl(120, 50%, 25%)"},
	{in: "hsl(120, 33.3%, 100.0%)"},
	{in: "hsl(120, .5%, 0%)"},
	{in: "hsl( 42 , 7% , 99% )"},
	{in: "", expectedError: errStringHSL},
	{in: "hsla(0,0%,0%)", expectedError: errStringHSL},
	{in: "hsl(361,0%,0%)", expectedError: errStringHSL},
	{in: "hsl(-1,0%,0%)", expectedError: errStringHSL},
	{in: "hsl(120,101%,0%)", expectedError: errStringHSL},
	{in: "hsl(120,100.1%,0%)", expectedError: errStringHSL},
	{in: "hsl(120,0.%,0%)", expectedError: errStringHSL},
	{in: "hsl(120,50,0%)", expectedError: errStringHSL},
	{in: "hsl(120,50%,0)", expectedError: errStringHSL},
	{in: "hsl(120,50%,0%,1)", expectedError: errStringHSL},
	{in: "hsl(120 50% 0%)", expectedError: errStringHSL},
}

var stringHSLATestCases = []*stringFormatColorTestCase{
	{in: "hsla(0,0%,0%,0)"},
	{in: "HsLa(0,0%,0%,1.0)"},
	{in: "hsla(360, 100%, 100%, 1)"},
	{in: "hsla(120, 50%, 25%, 0.5)"},
	{in: "hsla(120, 33.3%, 100.0%, 50%)"},
	{in: "hsla(120, .5%, 0%, 100.0%)"},
	{in: "hsla( 42 , 7% , 99% , 0.25 )"},
	{in: "hsla( 42 , 7% , 99% , .5 )"},
	{in: "hsla( 42 , 7% , 99% , 0. )"},
	{in: "", expectedError: errStringHSLA},
	{in: "hsla(0,0%,0%)", expectedError: errStringHSLA},
	{in: "hsl(0,0%,0%)", expectedError: errStringHSLA},
	{in: "hsla(361,0%,0%,0)", expectedError: errStringHSLA},
	{in: "hsla(120,101%,0%,0)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,2)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,1.1)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,-0.1)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,-1%)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,100.1%)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,100.%)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,.)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,1..0)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,50%%)", expectedError: errStringHSLA},
	{in: "hsla(120,50%,0%,50px)", expectedError: errStringHSLA},
}

var stringCMYKTestCases = []*stringFormatColorTestCase{
	{in: "cmyk(0%,0%,0%,0%)"},
	{in: "CMYK(0%,0%,0%,0%)"},
	{in: "cmyk(100%, 100%, 100%, 100%)"},
	{in: "cmyk(1%, 2%, 3%, 4%)"},
	{in: "cMyK(33.3%, 100.0%, .5%, 0%)"},
	{in: "cmyk( 0% , 50% , 75% , 100% )"},
	{in: "", expectedError: errStringCMYK},
	{in: "cmyk(0,0,0,0)", expectedError: errStringCMYK},
	{in: "cmyk(101%,0%,0%,0%)", expectedError: errStringCMYK},
	{in: "cmyk(100.1%,0%,0%,0%)", expectedError: errStringCMYK},
	{in: "cmyk(-1%,0%,0%,0%)", expectedError: errStringCMYK},
	{in: "cmyk(0.%,0%,0%,0%)", expectedError: errStringCMYK},
	{in: "cmyk(0%,0%,0%)", expectedError: errStringCMYK},
	{in: "cmyk(0%,0%,0%,0%,0%)", expectedError: errStringCMYK},
}

func TestStringHexColor(t *testing.T) {
	rule := StringHexColor()
	for _, tc := range stringHexColorTestCases {
		assertStringFormatColorRule(t, rule, tc.in, tc.expectedError, ErrorCodeStringHexColor)
	}
}

func TestStringRGB(t *testing.T) {
	rule := StringRGB()
	for _, tc := range stringRGBTestCases {
		assertStringFormatColorRule(t, rule, tc.in, tc.expectedError, ErrorCodeStringRGB)
	}
}

func TestStringRGBA(t *testing.T) {
	rule := StringRGBA()
	for _, tc := range stringRGBATestCases {
		assertStringFormatColorRule(t, rule, tc.in, tc.expectedError, ErrorCodeStringRGBA)
	}
}

func TestStringHSL(t *testing.T) {
	rule := StringHSL()
	for _, tc := range stringHSLTestCases {
		assertStringFormatColorRule(t, rule, tc.in, tc.expectedError, ErrorCodeStringHSL)
	}
}

func TestStringHSLA(t *testing.T) {
	rule := StringHSLA()
	for _, tc := range stringHSLATestCases {
		assertStringFormatColorRule(t, rule, tc.in, tc.expectedError, ErrorCodeStringHSLA)
	}
}

func TestStringCMYK(t *testing.T) {
	rule := StringCMYK()
	for _, tc := range stringCMYKTestCases {
		assertStringFormatColorRule(t, rule, tc.in, tc.expectedError, ErrorCodeStringCMYK)
	}
}

func BenchmarkStringHexColor(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringHexColor(), "#112233")
}

func BenchmarkStringRGB(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringRGB(), "rgb(12, 34, 56)")
}

func BenchmarkStringRGBA(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringRGBA(), "rgba(12, 34, 56, 0.25)")
}

func BenchmarkStringHSL(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringHSL(), "hsl(120, 50%, 25%)")
}

func BenchmarkStringHSLA(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringHSLA(), "hsla(120, 50%, 25%, 0.5)")
}

func BenchmarkStringCMYK(b *testing.B) {
	benchmarkStringFormatColorRule(b, StringCMYK(), "cmyk(1%, 2%, 3%, 4%)")
}

func assertStringFormatColorRule(
	t *testing.T,
	rule govy.Rule[string],
	in string,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	err := rule.Validate(in)
	if expectedError != "" {
		assert.EqualError(t, err, expectedError)
		assert.True(t, govy.HasErrorCode(err, errorCode))
		return
	}
	assert.NoError(t, err)
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
