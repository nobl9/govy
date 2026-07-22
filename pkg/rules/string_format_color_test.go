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
		"three mixed digits":     {in: "#369"},
		"four digits":            {in: "#abcd"},
		"four digit white":       {in: "#ffff"},
		"six digits":             {in: "#A1B2C3"},
		"six digit white":        {in: "#ffffff"},
		"six mixed case digits":  {in: "#FFCc99"},
		"eight digits":           {in: "#AABBCCDD"},
		"eight digit white":      {in: "#ffffffff"},
	}
	stringHexColorInvalidTestCases = map[string]stringFormatColorTestCase{
		"empty":        {in: ""},
		"hash only":    {in: "#"},
		"missing hash": {in: "112233"},
		"one digit":    {in: "#f"},
		"two digits":   {in: "#ff"},
		// cspell:disable
		"three digits ending non-hex": {in: "#ffg"},
		"four digits ending non-hex":  {in: "#fffg"},
		"five digits":                 {in: "#fffff"},
		"six digits ending non-hex":   {in: "#fffffg"},
		"seven digits":                {in: "#fffffff"},
		"eight digits ending non-hex": {in: "#fffffffg"},
		"nine digits":                 {in: "#fffffffff"},
		// cspell:enable
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
	benchmarkStringFormatColorRule(
		b,
		func(...StringColorOption) govy.Rule[string] { return StringHexColor() },
		stringHexColorValidTestCases,
		stringHexColorInvalidTestCases,
	)
}

var (
	stringRGBValidTestCases = map[string]stringFormatColorTestCase{
		"legacy numeric":                  {in: "rgb(0, 127.5, 255)"},
		"legacy leading-dot number":       {in: "rgb(.5,0,0)"},
		"legacy explicit plus sign":       {in: "rgb(+1,2,3)"},
		"legacy exponent components":      {in: "rgb(1e2,2,3)"},
		"legacy negative zero":            {in: "rgb(-0,0,0)"},
		"legacy signed exponent numbers":  {in: "rgb(-0, +5e1, 2.55e2)"},
		"legacy percentage":               {in: "RGB(-0%, +5e1%, 1e2%)"},
		"legacy signed percentage":        {in: "rgb(+50%,0%,0%)"},
		"legacy exponent percentage":      {in: "rgb(5e1%,0%,0%)"},
		"legacy rgba fractional alpha":    {in: "rgba(12, 34, 56, .25)"},
		"legacy rgba exponent alpha":      {in: "rgba(12, 34, 56, 5e-1)"},
		"legacy rgba exponent components": {in: "rgba(1e2,2,3,5e-1)"},
		"legacy rgba signed alpha":        {in: "rgba(.5,0,255,+.5)"},
		"legacy rgba negative zero alpha": {in: "rgba(-0,0,0,-0)"},
		"legacy rgb optional alpha":       {in: "rgb(0,0,0,.5)"},
		"legacy rgba optional alpha":      {in: "rgba(12, 34, 56)"},
		"legacy rgba alias without alpha": {in: "rgba(0,0,0)"},
		"legacy CSS whitespace":           {in: "rgb(0,\t0,\n0)"},
		"legacy whitespace before commas": {in: "rgb(0\t, 0\n, 0)"},
		"modern numeric":                  {in: "rgb(0 127.5 255)"},
		"modern mixed components":         {in: "rgb(0 50% 2.55e2)"},
		"modern alpha":                    {in: "rgb(0 0 0 / 50%)"},
		"modern rgba alias":               {in: "RGBA(100% 0% 33.3% / .5)"},
		"modern CSS whitespace":           {in: "rgb(\t0\n50%\f255\r/\t1 )"},
		"nil option":                      {in: "rgb(0 0 0)", options: []StringColorOption{nil}},
		"legacy-only numeric":             {in: "rgb(+0, 1.27e2, 255)", options: legacyColorOptions()},
		"legacy-only rgba alpha":          {in: "rgba(0%, 50%, 100%, 1)", options: legacyColorOptions()},
	}
	stringRGBInvalidLegacyTestCases = map[string]stringFormatColorTestCase{
		"legacy-only rejects modern": {
			in:      "rgb(0 0 0)",
			options: legacyColorOptions(),
		},
	}
	stringRGBInvalidTestCases = map[string]stringFormatColorTestCase{
		"empty":                          {in: ""},
		"wrong function":                 {in: "hsl(0 0 0)"},
		"same-length wrong function":     {in: "rxb(0,0,0)"},
		"component above number range":   {in: "rgb(256 0 0)"},
		"legacy component above range":   {in: "rgb(256,0,0)"},
		"component below number range":   {in: "rgb(-1 0 0)"},
		"legacy component below range":   {in: "rgb(-1,0,0)"},
		"percentage above range":         {in: "rgb(100.1% 0% 0%)"},
		"legacy percentage above range":  {in: "rgb(101%,0%,0%)"},
		"alpha above range":              {in: "rgb(0 0 0 / 1.1)"},
		"legacy alpha above range":       {in: "rgba(0,0,0,1.1)"},
		"alpha below range":              {in: "rgb(0 0 0 / -1%)"},
		"legacy mixed components":        {in: "rgb(0, 0%, 0)"},
		"legacy percentages then number": {in: "rgb(10%, 50%, 0)"},
		"legacy number then percentages": {in: "rgb(255, 50%, 0%)"},
		"legacy none component":          {in: "rgb(none, 0, 0)"},
		"legacy all none rgb":            {in: "rgb(none, none, none)"},
		"legacy leading empty component": {in: "rgb(,0, 0, 0)"},
		"legacy middle empty component":  {in: "rgb(0, 0,, 0)"},
		"legacy all none rgba":           {in: "rgba(none, none, none, none)"},
		"mixed comma after first":        {in: "rgb(0, 0 0)"},
		"mixed comma after second":       {in: "rgb(0 0, 0)"},
		"legacy slash alpha":             {in: "rgb(0, 0, 0 / .5)"},
		"modern comma alpha":             {in: "rgb(0 0 0, .5)"},
		"missing alpha after slash":      {in: "rgb(0 0 0 /)"},
		"sign-only number":               {in: "rgb(+ 0 0)"},
		"angle component":                {in: "rgb(0, 0, 0deg)"},
		"keyword component":              {in: "rgb(0, 0, light)"},
		"empty function":                 {in: "rgb()"},
		"one component":                  {in: "rgb(0)"},
		"legacy trailing comma":          {in: "rgb(0, 0, 0,)"},
		"missing modern separator":       {in: "rgb(0 0 0.5.5)"},
		"missing component":              {in: "rgb(0 0)"},
		"extra component":                {in: "rgb(0 0 0 0)"},
		"legacy extra component":         {in: "rgba(0, 0, 0, 0, 0)"},
		"trailing decimal point":         {in: "rgb(0. 0 0)"},
		"trailing alpha decimal point":   {in: "rgba(0, 0, 0, 0.)"},
		"compact trailing alpha point":   {in: "rgba(0,0,0,0.)"},
		"malformed exponent":             {in: "rgb(1e 0 0)"},
		"non-finite exponent":            {in: "rgb(1e999 0 0)"},
		"empty percentage":               {in: "rgb(%, 0%, 0%)"},
		"compact empty percentage":       {in: "rgb(%,0%,0%)"},
		"alpha dimension":                {in: "rgba(0, 0, 0, 50px)"},
		"legacy angle alpha":             {in: "rgba(0, 0, 0, 0deg)"},
		"legacy keyword alpha":           {in: "rgba(0, 0, 0, light)"},
		"none component":                 {in: "rgb(none 0 0)"},
		"calculated component":           {in: "rgb(calc(1) 0 0)"},
		"variable component":             {in: "rgb(var(--red) 0 0)"},
		"CSS comment":                    {in: "rgb(0/**/ 0 0)"},
		"escaped function name":          {in: "r\\67 b(0 0 0)"},
		"leading external whitespace":    {in: " rgb(0 0 0)"},
		"trailing external whitespace":   {in: "rgb(0 0 0) "},
		"non-CSS whitespace":             {in: "rgb(0\u00a00\u00a00)"},
		"vertical tab whitespace":        {in: "rgb(0\v0\v0)"},
		"legacy non-CSS whitespace":      {in: "rgb(0,\u00a00,0)"},
		"legacy vertical tab whitespace": {in: "rgb(0,\v0,0)"},
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
	benchmarkStringFormatColorRule(
		b,
		StringRGB,
		stringRGBValidTestCases,
		stringRGBInvalidLegacyTestCases,
		stringRGBInvalidTestCases,
	)
}

var (
	stringHSLValidTestCases = map[string]stringFormatColorTestCase{
		"legacy percentage":                {in: "hsl(120, 50%, 25%)"},
		"legacy exact fractional hue":      {in: "hsl(120.5,50%,50%)"},
		"legacy explicit plus hue":         {in: "hsl(+120,50%,50%)"},
		"legacy exact exponent hue":        {in: "hsl(1.2e2,50%,50%)"},
		"legacy exact signed zero hue":     {in: "hsl(-0,50%,50%)"},
		"legacy signed zero hue":           {in: "hsl(-0, 50%, 25%)"},
		"legacy negative cyclic hue":       {in: "hsl(-1, 0%, 0%)"},
		"legacy exact negative cyclic hue": {in: "hsl(-1,0%,0%)"},
		"legacy positive cyclic hue":       {in: "hsl(361,0%,0%)"},
		"legacy fractional hue":            {in: "hsl(.5, 50%, 25%)"},
		"legacy exponent hue":              {in: "hsl(1.2e2, 50%, 25%)"},
		"legacy angle hue":                 {in: "HSLA(-.5turn, 50%, 25%, .5)"},
		"legacy degree hue":                {in: "hsl(120deg,50%,50%)"},
		"legacy turn hue":                  {in: "hsl(.5turn,50%,50%)"},
		"legacy radian hue":                {in: "hsl(2.0944rad,50%,50%)"},
		"legacy gradian hue":               {in: "hsl(133.333grad,50%,50%)"},
		"legacy fractional alpha":          {in: "hsla(120.5,50%,50%,5e-1)"},
		"legacy hsl optional alpha":        {in: "hsl(120,50%,50%,.5)"},
		"legacy hsla optional alpha":       {in: "hsla(120,50%,50%)"},
		"legacy cyclic hue":                {in: "hsl(720, 50%, 25%)"},
		"modern percentage":                {in: "hsl(120 50% 25%)"},
		"modern numeric":                   {in: "hsl(1.2e2 50 25)"},
		"modern mixed components":          {in: "hsl(120deg 50 25%)"},
		"modern alpha":                     {in: "hsla(120 50% 25% / 50%)"},
		"modern hsla optional alpha":       {in: "hsla(120 50% 25%)"},
		"modern CSS whitespace":            {in: "hsl(\t+120\n50%\f25%\r/\t5e-1 )"},
		"grad hue":                         {in: "hsl(100grad 50% 25%)"},
		"radian hue":                       {in: "hsl(3.14rad 50% 25%)"},
		"turn hue":                         {in: "hsl(.5TURN 50% 25%)"},
		"legacy-only percentage":           {in: "hsl(-120deg, 50%, 25%)", options: legacyColorOptions()},
		"legacy-only alpha":                {in: "hsla(120, 50%, 25%, 25%)", options: legacyColorOptions()},
	}
	stringHSLInvalidLegacyTestCases = map[string]stringFormatColorTestCase{
		"legacy-only rejects modern": {
			in:      "hsl(120 50% 25%)",
			options: legacyColorOptions(),
		},
	}
	stringHSLInvalidTestCases = map[string]stringFormatColorTestCase{
		"empty":                               {in: ""},
		"wrong function":                      {in: "rgb(0 0 0)"},
		"legacy numeric saturation":           {in: "hsl(120, 50, 25%)"},
		"legacy none hue":                     {in: "hsl(none, 100%, 50%)"},
		"legacy all none hsl":                 {in: "hsl(none, none, none)"},
		"legacy all none hsla":                {in: "hsla(none, none, none, none)"},
		"percentage hue":                      {in: "hsl(50%, 50%, 0%)"},
		"legacy numeric lightness":            {in: "hsl(0, 50%, 30)"},
		"legacy percentage hue alpha":         {in: "hsla(50%, 50%, 0%, 1)"},
		"saturation above range":              {in: "hsl(120 101% 25%)"},
		"legacy saturation above range":       {in: "hsl(120,101%,0%)"},
		"lightness below range":               {in: "hsl(120 50% -1)"},
		"alpha above range":                   {in: "hsl(120 50% 25% / 101%)"},
		"legacy alpha percentage above range": {in: "hsla(120,50%,0%,100.1%)"},
		"unknown angle unit":                  {in: "hsl(120foo 50% 25%)"},
		"space before angle unit":             {in: "hsl(120 deg 50% 25%)"},
		"trailing hue decimal point":          {in: "hsl(120. 50% 25%)"},
		"trailing alpha decimal point":        {in: "hsla(0, 50%, 25%, 0.)"},
		"compact trailing alpha point":        {in: "hsla(0,0%,0%,0.)"},
		"alpha dimension":                     {in: "hsla(0, 50%, 25%, 50px)"},
		"empty percentage":                    {in: "hsl(0, %, 25%)"},
		"none hue":                            {in: "hsl(none 50% 25%)"},
		"none saturation":                     {in: "hsl(120 none 25%)"},
		"calculated hue":                      {in: "hsl(calc(120) 50% 25%)"},
		"empty function":                      {in: "hsl()"},
		"one component":                       {in: "hsl(0)"},
		"legacy trailing comma":               {in: "hsl(0, 50%, 25%,)"},
		"legacy extra alpha":                  {in: "hsla(0, 50%, 25%, 1, 0%)"},
		"legacy mixed delimiters":             {in: "hsl(120, 50%, 25% / .5)"},
		"legacy mixed comma and space":        {in: "hsl(0, 0% 0%)"},
		"modern comma delimiter":              {in: "hsl(120 50%, 25%)"},
		"keyword lightness":                   {in: "hsl(0, 0%, light)"},
		"leading external whitespace":         {in: " hsl(120 50% 25%)"},
		"trailing external whitespace":        {in: "hsl(120 50% 25%) "},
		"non-CSS whitespace":                  {in: "hsl(0\u00a050%\u00a025%)"},
		"vertical tab whitespace":             {in: "hsl(0\v50%\v25%)"},
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
	benchmarkStringFormatColorRule(
		b,
		StringHSL,
		stringHSLValidTestCases,
		stringHSLInvalidLegacyTestCases,
		stringHSLInvalidTestCases,
	)
}

var (
	stringDeviceCMYKValidTestCases = map[string]stringFormatColorTestCase{
		"legacy numeric":                        {in: "device-cmyk(-0, +.25, 5e-1, 1)"},
		"legacy zero":                           {in: "device-cmyk(0, 0, 0, 0)"},
		"legacy fractional components":          {in: "device-cmyk(.1, .2, .3, .4)"},
		"legacy uppercase":                      {in: "DEVICE-CMYK(0, 0, 0, 0)"},
		"modern numeric":                        {in: "device-cmyk(0 .25 .5 1)"},
		"modern percentage":                     {in: "device-cmyk(0% 25% 50% 100%)"},
		"modern exact percentage":               {in: "device-cmyk(0% 81% 81% 30%)"},
		"modern percentage signs and exponents": {in: "device-cmyk(-0% +2.5e1% 5e1% 1e2%)"},
		"modern mixed components":               {in: "device-cmyk(0 25% .5 100%)"},
		"modern exact mixed components":         {in: "device-cmyk(0 81% 81% 30%)"},
		"modern percentage alpha":               {in: "device-cmyk(0 0 0 1 / 50%)"},
		"modern exponent alpha":                 {in: "device-cmyk(0 0 0 1 / 5e-1)"},
		"modern CSS whitespace":                 {in: "device-cmyk(\t0\n25%\f.5\r1 /\t.5 )"},
		"legacy-only numeric":                   {in: "device-cmyk(0, .25, .5, 1)", options: legacyColorOptions()},
	}
	stringDeviceCMYKInvalidLegacyTestCases = map[string]stringFormatColorTestCase{
		"legacy-only rejects modern": {
			in:      "device-cmyk(0 0 0 1)",
			options: legacyColorOptions(),
		},
	}
	stringDeviceCMYKInvalidTestCases = map[string]stringFormatColorTestCase{
		"empty":                            {in: ""},
		"non-standard function":            {in: "cmyk(0 0 0 1)"},
		"non-standard percentage function": {in: "cmyk(0%,0%,0%,0%)"},
		"legacy percentage":                {in: "device-cmyk(0%, 25%, 50%, 100%)"},
		"legacy all-zero percentage":       {in: "device-cmyk(0%, 0%, 0%, 0%)"},
		"legacy alpha":                     {in: "device-cmyk(0, 0, 0, 1, .5)"},
		"legacy fifth zero component":      {in: "device-cmyk(0,0,0,0,0)"},
		"legacy fifth one component":       {in: "device-cmyk(0,0,0,0,1)"},
		"legacy components with slash":     {in: "device-cmyk(0,0,0,0 / .5)"},
		"mixed comma and space":            {in: "device-cmyk(0, 0 0 1)"},
		"modern comma alpha":               {in: "device-cmyk(0 0 0 1, .5)"},
		"legacy trailing comma":            {in: "device-cmyk(0, 0, 0, 1,)"},
		"number above range":               {in: "device-cmyk(0 0 0 1.1)"},
		"number below range":               {in: "device-cmyk(-.1 0 0 1)"},
		"percentage above range":           {in: "device-cmyk(0% 0% 0% 101%)"},
		"alpha above range":                {in: "device-cmyk(0 0 0 1 / 1.1)"},
		"trailing alpha decimal point":     {in: "device-cmyk(0 0 0 1 / 0.)"},
		"alpha dimension":                  {in: "device-cmyk(0 0 0 1 / 50px)"},
		"empty function":                   {in: "device-cmyk()"},
		"legacy missing component":         {in: "device-cmyk(0,0,0)"},
		"missing component":                {in: "device-cmyk(0 0 1)"},
		"extra component":                  {in: "device-cmyk(0 0 0 1 0)"},
		"none component":                   {in: "device-cmyk(none 0 0 1)"},
		"mixed none component":             {in: "device-cmyk(none 0% .5 1 / 50%)"},
		"calculated component":             {in: "device-cmyk(calc(0) 0 0 1)"},
		"fallback color":                   {in: "device-cmyk(0 0 0 1, black)"},
		"fallback after alpha":             {in: "device-cmyk(0 0 0 0 / 50%, red)"},
		"leading external whitespace":      {in: " device-cmyk(0 0 0 1)"},
		"trailing external whitespace":     {in: "device-cmyk(0 0 0 1) "},
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
	benchmarkStringFormatColorRule(
		b,
		StringDeviceCMYK,
		stringDeviceCMYKValidTestCases,
		stringDeviceCMYKInvalidLegacyTestCases,
		stringDeviceCMYKInvalidTestCases,
	)
}

func assertStringColorRuleTestCases(
	t *testing.T,
	ruleFactory func(...StringColorOption) govy.Rule[string],
	testCases map[string]stringFormatColorTestCase,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := ruleFactory(tc.options...).Validate(tc.in)
			if expectedError != "" {
				assert.EqualError(t, err, expectedError)
				assert.True(t, govy.HasErrorCode(err, errorCode))
				return
			}
			assert.NoError(t, err)
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
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			err := rule.Validate(tc.in)
			if expectedError != "" {
				assert.EqualError(t, err, expectedError)
				assert.True(t, govy.HasErrorCode(err, errorCode))
				return
			}
			assert.NoError(t, err)
		})
	}
}

func benchmarkStringFormatColorRule(
	b *testing.B,
	ruleFactory func(...StringColorOption) govy.Rule[string],
	testCaseGroups ...map[string]stringFormatColorTestCase,
) {
	b.Helper()
	type benchmarkCase struct {
		rule govy.Rule[string]
		in   string
	}
	testCaseCount := 0
	for _, testCases := range testCaseGroups {
		testCaseCount += len(testCases)
	}
	testCases := make([]benchmarkCase, 0, testCaseCount)
	for _, testCaseGroup := range testCaseGroups {
		for _, tc := range testCaseGroup {
			testCases = append(testCases, benchmarkCase{
				rule: ruleFactory(tc.options...),
				in:   tc.in,
			})
		}
	}
	for b.Loop() {
		for _, tc := range testCases {
			_ = tc.rule.Validate(tc.in)
		}
	}
}

func legacyColorOptions() []StringColorOption {
	return []StringColorOption{
		StringColorLegacySyntaxOnly(),
		nil,
		StringColorLegacySyntaxOnly(),
	}
}
