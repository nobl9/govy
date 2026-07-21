package rules

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

const (
	wptColorCorpusRevision = "073a34f030ad68af980d7a627e829900b003961b"
	wptColorCorpusSource   = "https://github.com/web-platform-tests/wpt/tree/" +
		wptColorCorpusRevision + "/css/css-color"
)

// These tables copy every input from the six WPT color corpora at
// wptColorCorpusRevision. RGB and HSL source-valid inputs are partitioned by
// Govy's contract; each partition preserves the inputs' relative source order.
// Keep the complete finite corpora: reducing them to representative subsets
// defeats their regression value.
//
// Source files:
//   - parsing/color-computed-hex-color.html (6 computed-value inputs)
//   - parsing/color-invalid-hex-color.html (10 invalid parsing inputs)
//   - parsing/color-valid-rgb.html (70 valid parsing inputs)
//   - parsing/color-invalid-rgb.html (30 invalid parsing inputs)
//   - parsing/color-valid-hsl.html (59 valid parsing inputs)
//   - parsing/color-invalid-hsl.html (23 invalid parsing inputs)
//
// The two additional HSL inputs come from
// css/css-color/hsl-clamp-negative-saturation.html at the same revision.
var (
	wptHexSourceValidInputs = enumerateWPTColorInputs(1,
		"#fff",
		"#ffff",
		"#ffffff",
		"#ffffffff",
		"#FFCc99",
		"#369",
	)
	wptHexSourceInvalidInputs = enumerateWPTColorInputs(1,
		"#",
		"#f",
		"#ff",
		// cspell:disable
		"#ffg",
		"#fffg",
		"#fffff",
		"#fffffg",
		"#fffffff",
		"#fffffffg",
		"#fffffffff",
		// cspell:enable
	)
	wptRGBSourceValidInContractInputs = indexWPTColorInputs([]int{30},
		"rgb(255 20% 102)",
	)
	wptRGBSourceValidOutOfContractInputs = []wptColorContractExclusion{
		{
			classification: "none keyword excluded by concrete-literal contract",
			inputs: enumerateWPTColorInputs(1,
				"rgb(none none none)",
				"rgb(none none none / none)",
				"rgb(128 none none)",
				"rgb(128 none none / none)",
				"rgb(none none none / .5)",
				"rgb(20% none none)",
				"rgb(20% none none / none)",
				"rgb(none none none / 50%)",
				"rgba(none none none)",
				"rgba(none none none / none)",
				"rgba(128 none none)",
				"rgba(128 none none / none)",
				"rgba(none none none / .5)",
				"rgba(20% none none)",
				"rgba(20% none none / none)",
				"rgba(none none none / 50%)",
			),
		},
		{
			classification: "CSS clamping excluded by in-range contract",
			inputs: indexWPTColorInputs(
				[]int{17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 31, 32},
				"rgb(-2 3 4)",
				"rgb(-20% 20% 40%)",
				"rgb(257 30 40)",
				"rgb(250% 20% 40%)",
				"rgba(-2 3 4)",
				"rgba(-20% 20% 40%)",
				"rgba(257 30 40)",
				"rgba(250% 20% 40%)",
				"rgba(-2 3 4 / .5)",
				"rgba(-20% 20% 40% / 50%)",
				"rgba(257 30 40 / 50%)",
				"rgba(250% 20% 40% / .5)",
				"rgb(250% 51 40%)",
				"rgb(500, 0, 0)",
				"rgb(-500, 64, 128)",
			),
		},
		{
			classification: "calculation excluded by concrete-literal contract",
			inputs: enumerateWPTColorInputs(33,
				"rgb(calc(infinity), 0, 0)",
				"rgb(0, calc(infinity), 0)",
				"rgb(0, 0, calc(infinity))",
				"rgba(0, 0, 0, calc(infinity))",
				"rgb(calc(-infinity), 0, 0)",
				"rgb(0, calc(-infinity), 0)",
				"rgb(0, 0, calc(-infinity))",
				"rgba(0, 0, 0, calc(-infinity))",
				"rgb(calc(NaN), 0, 0)",
				"rgb(0, calc(NaN), 0)",
				"rgb(0, 0, calc(NaN))",
				"rgba(0, 0, 0, calc(NaN))",
				"rgb(calc(0 / 0), 0, 0)",
				"rgb(0, calc(0 / 0), 0)",
				"rgb(0, 0, calc(0 / 0))",
				"rgba(0, 0, 0, calc(0 / 0))",
				"rgb(calc(50% + (sign(1em - 10px) * 10%)), 0%, 0%, 50%)",
				"rgba(calc(50% + (sign(1em - 10px) * 10%)), 0%, 0%, 50%)",
				"rgb(calc(50 + (sign(1em - 10px) * 10)), 0, 0, 0.5)",
				"rgba(calc(50 + (sign(1em - 10px) * 10)), 0, 0, 0.5)",
				"rgb(0%, 0%, 0%, calc(50% + (sign(1em - 10px) * 10%)))",
				"rgba(0%, 0%, 0%, calc(50% + (sign(1em - 10px) * 10%)))",
				"rgb(0, 0, 0, calc(0.75 + (sign(1em - 10px) * 0.1)))",
				"rgba(0, 0, 0, calc(0.75 + (sign(1em - 10px) * 0.1)))",
				"rgb(calc(50% + (sign(1em - 10px) * 10%)) 0% 0% / 50%)",
				"rgba(calc(50% + (sign(1em - 10px) * 10%)) 0% 0% / 50%)",
				"rgb(calc(50 + (sign(1em - 10px) * 10)) 0 0 / 0.5)",
				"rgba(calc(50 + (sign(1em - 10px) * 10)) 0 0 / 0.5)",
				"rgb(0% 0% 0% / calc(50% + (sign(1em - 10px) * 10%)))",
				"rgba(0% 0% 0% / calc(50% + (sign(1em - 10px) * 10%)))",
				"rgb(0 0 0 / calc(0.75 + (sign(1em - 10px) * 0.1)))",
				"rgba(0 0 0 / calc(0.75 + (sign(1em - 10px) * 0.1)))",
				"rgba(calc(50% + (sign(1em - 10px) * 10%)) 0 0% / 0.5)",
				"rgba(0% 0 0% / calc(0.75 + (sign(1em - 10px) * 0.1)))",
				"rgba(calc(50 + (sign(1em - 10px) * 10)) 400 -400 / 0.5)",
				"rgba(calc(50% + (sign(1em - 10px) * 10%)) 400% -400% / 0.5)",
				"rgba(calc(50 + (sign(1em - 10px) * 10)), 400, -400, 0.5)",
				"rgba(calc(50% + (sign(1em - 10px) * 10%)), 400%, -400%, 0.5)",
			),
		},
	}
	wptRGBSourceInvalidInputs = enumerateWPTColorInputs(1,
		"rgb(none, none, none)",
		"rgba(none, none, none, none)",
		"rgb(128, 0, none)",
		"rgb(255, 255, 255, none)",
		"rgb(10%, 50%, 0)",
		"rgb(255, 50%, 0%)",
		"rgb(0, 0 0)",
		"rgb(0 0, 0)",
		"rgb(,0, 0, 0)",
		"rgb(0, 0, 0,)",
		"rgb(0, 0,, 0)",
		"rgb(0, 0, 0deg)",
		"rgb(0, 0, light)",
		"rgb()",
		"rgb(0)",
		"rgb(0, 0)",
		"rgb(0%)",
		"rgb(0%, 0%)",
		"rgba(10%, 50%, 0, 1)",
		"rgba(255, 50%, 0%, 1)",
		"rgba(0, 0, 0 0)",
		"rgba(0, 0, 0, 0deg)",
		"rgba(0, 0, 0, light)",
		"rgba()",
		"rgba(0)",
		"rgba(0, 0, 0, 0, 0)",
		"rgba(0%)",
		"rgba(0%, 0%)",
		"rgba(0%, 0%, 0%, 0%, 0%)",
		"rgb(257, 0, 5 / 0)",
	)
	wptHSLSourceValidInContractInputs = indexWPTColorInputs(
		[]int{1, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 21, 22, 23, 24, 25, 26},
		"hsl(120 30% 50%)",
		"hsl(120 30% 50% / 0.5)",
		"hsl(0 0% 0%)",
		"hsl(0 0% 0% / 0)",
		"hsla(0 0% 0%)",
		"hsla(0 0% 0% / 0)",
		"hsl(120 0% 0%)",
		"hsl(120 80% 0%)",
		"hsl(120 0% 50%)",
		"hsl(120 100% 50% / 0)",
		"hsl(0 100% 50%)",
		"hsl(120 30 50)",
		"hsl(120 30 50 / 0.5)",
		"hsl(120 30% 50)",
		"hsl(120 30% 50 / 0.5)",
		"hsl(120 30 50%)",
		"hsl(120 30 50% / 0.5)",
	)
	wptHSLSourceValidOutOfContractInputs = []wptColorContractExclusion{
		{
			classification: "none keyword excluded by concrete-literal contract",
			inputs: indexWPTColorInputs(
				[]int{3, 5, 7, 9, 11, 13, 15, 17, 19, 27, 28, 29, 30, 31},
				"hsl(none none none)",
				"hsl(none none none / none)",
				"hsla(none none none)",
				"hsla(none none none / none)",
				"hsl(120 none none)",
				"hsl(120 80% none)",
				"hsl(120 none 50%)",
				"hsl(120 100% 50% / none)",
				"hsl(none 100% 50%)",
				"hsl(120 none 50)",
				"hsl(120 none 50 / 0.5)",
				"hsl(120 30 none)",
				"hsl(120 30 none / 0.5)",
				"hsl(120 30 50 / none)",
			),
		},
		{
			classification: "CSS clamping excluded by in-range contract",
			inputs: enumerateWPTColorInputs(32,
				"hsl(0 -50% 40%)",
				"hsl(30 -50% 60)",
				"hsl(0 -50 40%)",
				"hsl(30 -50 60)",
			),
		},
		{
			classification: "calculation excluded by concrete-literal contract",
			inputs: enumerateWPTColorInputs(36,
				"hsl(calc(infinity) 100% 50%)",
				"hsl(calc(-infinity) 100% 50%)",
				"hsl(calc(0 / 0) 100% 50%)",
				"hsl(90 50% 50% / calc(infinity))",
				"hsl(90 50% 50% / calc(-infinity))",
				"hsl(90 50% 50% / calc(0 / 0))",
				"hsl(calc(50deg + (sign(1em - 10px) * 10deg)), 0%, 0%, 50%)",
				"hsla(calc(50deg + (sign(1em - 10px) * 10deg)), 0%, 0%, 50%)",
				"hsl(calc(50 + (sign(1em - 10px) * 10)), 0%, 0%, 50%)",
				"hsla(calc(50 + (sign(1em - 10px) * 10)), 0%, 0%, 50%)",
				"hsl(0deg, 0%, 0%, calc(50% + (sign(1em - 10px) * 10%)))",
				"hsla(0deg, 0%, 0%, calc(50% + (sign(1em - 10px) * 10%)))",
				"hsl(0, 0%, 0%, calc(50% + (sign(1em - 10px) * 10%)))",
				"hsla(0, 0%, 0%, calc(50% + (sign(1em - 10px) * 10%)))",
				"hsl(calc(50deg + (sign(1em - 10px) * 10deg)) 0% 0% / 50%)",
				"hsla(calc(50deg + (sign(1em - 10px) * 10deg)) 0% 0% / 50%)",
				"hsl(calc(50 + (sign(1em - 10px) * 10)) 0 0 / 0.5)",
				"hsla(calc(50 + (sign(1em - 10px) * 10)) 0 0 / 0.5)",
				"hsl(0deg 0% 0% / calc(50% + (sign(1em - 10px) * 10%)))",
				"hsla(0deg 0% 0% / calc(50% + (sign(1em - 10px) * 10%)))",
				"hsl(0 0 0 / calc(0.75 + (sign(1em - 10px) * 0.1)))",
				"hsla(0 0 0 / calc(0.75 + (sign(1em - 10px) * 0.1)))",
				"hsla(calc(50deg + (sign(1em - 10px) * 10deg)) -100 300 / 0.5)",
				"hsla(calc(50deg + (sign(1em - 10px) * 10deg)) -100% 300% / 0.5)",
			),
		},
	}
	wptHSLSourceInvalidInputs = enumerateWPTColorInputs(1,
		"hsl(none, none, none)",
		"hsla(none, none, none, none)",
		"hsl(none, 100%, 50%)",
		"hsla(120, 100%, 50%, none)",
		"hsl(10, 50%, 0)",
		"hsl(50%, 50%, 0%)",
		"hsl(0, 0% 0%)",
		"hsl(0, 0%, light)",
		"hsl()",
		"hsl(0)",
		"hsl(0, 0%)",
		"hsl(0, 50, 30%)",
		"hsl(0, 50%, 30)",
		"hsla(10, 50%, 0, 1)",
		"hsla(50%, 50%, 0%, 1)",
		"hsla(0, 0% 0%, 1)",
		"hsla(0, 0%, light, 1)",
		"hsla()",
		"hsla(0)",
		"hsla(0, 0%)",
		"hsla(0, 0%, 0%, 1, 0%)",
		"hsl(0, 50, 30%, 1)",
		"hsl(0, 50%, 30, 1)",
	)
	wptHSLNegativeSaturationReftestInputs = enumerateWPTColorInputs(1,
		"hsl(0, -50%, 40%)",
		"hsl(30, -50%, 60%)",
	)
)

type wptColorContractExclusion struct {
	classification string
	inputs         []wptColorInput
}

type wptColorInput struct {
	sourceOrdinal int
	value         string
}

func TestStringColorWPTCorpora(t *testing.T) {
	t.Parallel()
	t.Logf("WPT color corpus source: %s", wptColorCorpusSource)
	t.Run("hex/source-valid", func(t *testing.T) {
		t.Parallel()
		assertWPTColorInputs(t, StringHexColor(), wptHexSourceValidInputs, 6, "", ErrorCodeStringHexColor)
	})
	t.Run("hex/source-invalid", func(t *testing.T) {
		t.Parallel()
		assertWPTColorInputs(
			t,
			StringHexColor(),
			wptHexSourceInvalidInputs,
			10,
			errStringHexColor,
			ErrorCodeStringHexColor,
		)
	})
	t.Run("RGB/source-valid/in-contract", func(t *testing.T) {
		t.Parallel()
		assertWPTColorInputs(t, StringRGB(), wptRGBSourceValidInContractInputs, 1, "", ErrorCodeStringRGB)
	})
	t.Run("RGB/source-valid/out-of-contract", func(t *testing.T) {
		t.Parallel()
		assertWPTColorContractExclusions(
			t,
			StringRGB(),
			wptRGBSourceValidOutOfContractInputs,
			69,
			errStringRGB,
			ErrorCodeStringRGB,
		)
	})
	t.Run("RGB/source-invalid", func(t *testing.T) {
		t.Parallel()
		assertWPTColorInputs(t, StringRGB(), wptRGBSourceInvalidInputs, 30, errStringRGB, ErrorCodeStringRGB)
	})
	t.Run("HSL/source-valid/in-contract", func(t *testing.T) {
		t.Parallel()
		assertWPTColorInputs(t, StringHSL(), wptHSLSourceValidInContractInputs, 17, "", ErrorCodeStringHSL)
	})
	t.Run("HSL/source-valid/out-of-contract", func(t *testing.T) {
		t.Parallel()
		assertWPTColorContractExclusions(
			t,
			StringHSL(),
			wptHSLSourceValidOutOfContractInputs,
			42,
			errStringHSL,
			ErrorCodeStringHSL,
		)
	})
	t.Run("HSL/source-invalid", func(t *testing.T) {
		t.Parallel()
		assertWPTColorInputs(t, StringHSL(), wptHSLSourceInvalidInputs, 23, errStringHSL, ErrorCodeStringHSL)
	})
	t.Run("HSL/negative-saturation-reftest/out-of-contract", func(t *testing.T) {
		t.Parallel()
		assertWPTColorInputs(
			t,
			StringHSL(),
			wptHSLNegativeSaturationReftestInputs,
			2,
			errStringHSL,
			ErrorCodeStringHSL,
		)
	})
}

func assertWPTColorContractExclusions(
	t *testing.T,
	rule govy.Rule[string],
	exclusions []wptColorContractExclusion,
	expectedCount int,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	count := 0
	for _, exclusion := range exclusions {
		count += len(exclusion.inputs)
	}
	if count != expectedCount {
		t.Fatalf("WPT contract-exclusion input count: got %d, want %d", count, expectedCount)
	}
	for _, exclusion := range exclusions {
		t.Run(exclusion.classification, func(t *testing.T) {
			t.Parallel()
			assertWPTColorInputs(t, rule, exclusion.inputs, len(exclusion.inputs), expectedError, errorCode)
		})
	}
}

func assertWPTColorInputs(
	t *testing.T,
	rule govy.Rule[string],
	inputs []wptColorInput,
	expectedCount int,
	expectedError string,
	errorCode govy.ErrorCode,
) {
	t.Helper()
	if len(inputs) != expectedCount {
		t.Fatalf("WPT input count: got %d, want %d", len(inputs), expectedCount)
	}
	for _, input := range inputs {
		t.Run(wptColorSubtestName(input), func(t *testing.T) {
			t.Parallel()
			t.Logf("WPT source input %03d: %q", input.sourceOrdinal, input.value)
			err := rule.Validate(input.value)
			if expectedError != "" {
				assert.EqualError(t, err, expectedError)
				assert.True(t, govy.HasErrorCode(err, errorCode))
				return
			}
			assert.NoError(t, err)
		})
	}
}

func enumerateWPTColorInputs(firstSourceOrdinal int, values ...string) []wptColorInput {
	inputs := make([]wptColorInput, 0, len(values))
	for index, value := range values {
		inputs = append(inputs, wptColorInput{
			sourceOrdinal: firstSourceOrdinal + index,
			value:         value,
		})
	}
	return inputs
}

func indexWPTColorInputs(sourceOrdinals []int, values ...string) []wptColorInput {
	if len(sourceOrdinals) != len(values) {
		panic("WPT source ordinal count must match input count")
	}
	inputs := make([]wptColorInput, 0, len(values))
	for index, value := range values {
		inputs = append(inputs, wptColorInput{
			sourceOrdinal: sourceOrdinals[index],
			value:         value,
		})
	}
	return inputs
}

func wptColorSubtestName(input wptColorInput) string {
	return fmt.Sprintf("%03d_%s", input.sourceOrdinal, url.QueryEscape(input.value))
}
