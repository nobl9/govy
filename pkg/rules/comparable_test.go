package rules

import (
	"cmp"
	"encoding/json"
	"fmt"
	"math"
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
)

type comparisonTestCase[T comparable] struct {
	value                T
	input                T
	expectedError        string
	jsonSchemaDifference string
	schemaName           string
}

var eqTestCases = []*comparisonTestCase[any]{
	{value: 1, input: 1},
	{value: 1, input: 0, expectedError: "must be equal to '1'"},
	{value: 1.1, input: 1.3, expectedError: "must be equal to '1.1'"},
	{value: 1.1, input: 1.1},
	{value: 0, input: 0},
	{value: 0, input: -1, expectedError: "must be equal to '0'"},
	{value: false, input: false},
	{value: false, input: true, expectedError: "must be equal to 'false'"},
	{value: "", input: "", schemaName: "empty_string"},
	{value: "", input: "other", expectedError: "must be equal to ''", schemaName: "empty_string"},
	{value: nil, input: nil, schemaName: "null"},
	{value: nil, input: 0, expectedError: "must be equal to '<no value>'", schemaName: "null"},
	{
		value: 1, input: float64(1), expectedError: "must be equal to '1'",
		jsonSchemaDifference: "JSON numbers do not preserve the distinction between Go int and float64 values.",
	},
	{value: [2]int{1, 2}, input: [2]int{1, 2}, schemaName: "array"},
	{
		value: [2]int{1, 2}, input: [2]int{2, 1}, expectedError: "must be equal to '[1 2]'",
		schemaName: "array",
	},
}

func TestEQ(t *testing.T) {
	for _, tc := range eqTestCases {
		err := EQ(tc.value).Validate(tc.input)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeEqualTo))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestEQ_JSONSchema(t *testing.T) {
	t.Parallel()
	assertComparisonJSONSchema(t, "eq", EQ[any], eqTestCases)
}

func BenchmarkEQ(b *testing.B) {
	for _, tc := range eqTestCases {
		rule := EQ(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var neqTestCases = []*comparisonTestCase[any]{
	{value: 1.1, input: 1.3},
	{value: 1.1, input: 1.1, expectedError: "must not be equal to '1.1'"},
	{value: 0, input: -1},
	{value: 0, input: 0, expectedError: "must not be equal to '0'"},
	{value: false, input: true},
	{value: false, input: false, expectedError: "must not be equal to 'false'"},
	{value: "", input: "other", schemaName: "empty_string"},
	{value: "", input: "", expectedError: "must not be equal to ''", schemaName: "empty_string"},
	{value: nil, input: 0, schemaName: "null"},
	{value: nil, input: nil, expectedError: "must not be equal to '<no value>'", schemaName: "null"},
	{
		value: 0, input: float64(0),
		jsonSchemaDifference: "JSON numbers do not preserve the distinction between Go int and float64 values.",
	},
	{value: [2]int{1, 2}, input: [2]int{2, 1}, schemaName: "array"},
	{
		value: [2]int{1, 2}, input: [2]int{1, 2}, expectedError: "must not be equal to '[1 2]'",
		schemaName: "array",
	},
}

func TestNEQ(t *testing.T) {
	for _, tc := range neqTestCases {
		err := NEQ(tc.value).Validate(tc.input)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeNotEqualTo))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestNEQ_JSONSchema(t *testing.T) {
	t.Parallel()
	assertComparisonJSONSchema(t, "neq", NEQ[any], neqTestCases)
}

func BenchmarkNEQ(b *testing.B) {
	for _, tc := range neqTestCases {
		rule := NEQ(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

func TestComparableRule_DescriptionTemplateFailure(t *testing.T) {
	compared, other := make(chan int), make(chan int)
	for _, tc := range []struct {
		name        string
		newRule     func(chan int) govy.Rule[chan int]
		valid       chan int
		invalid     chan int
		errorCode   govy.ErrorCode
		description string
	}{
		{
			name:        "EQ",
			newRule:     EQ[chan int],
			valid:       compared,
			invalid:     other,
			errorCode:   ErrorCodeEqualTo,
			description: "must be equal to '",
		},
		{
			name:        "NEQ",
			newRule:     NEQ[chan int],
			valid:       other,
			invalid:     compared,
			errorCode:   ErrorCodeNotEqualTo,
			description: "must not be equal to '",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("plan before validation", func(t *testing.T) {
				validator := govy.New(
					govy.For(govy.GetSelf[chan int]()).WithName("channel").Rules(tc.newRule(compared)),
				)
				plan, err := govy.Plan(validator)
				assert.Require(t, assert.NoError(t, err))
				assert.Require(t, assert.Len(t, plan.Properties, 1))
				assert.Require(t, assert.Len(t, plan.Properties[0].Rules, 1))
				assert.Equal(t, tc.description, plan.Properties[0].Rules[0].Description)
				assert.Equal(t, tc.errorCode, plan.Properties[0].Rules[0].ErrorCode)
			})

			for _, consumer := range []struct {
				name     string
				override func(govy.Rule[chan int]) govy.Rule[chan int]
			}{
				{
					name: "message",
					override: func(rule govy.Rule[chan int]) govy.Rule[chan int] {
						return rule.WithMessage("invalid channel")
					},
				},
				{
					name: "message template",
					override: func(rule govy.Rule[chan int]) govy.Rule[chan int] {
						return rule.WithMessageTemplateString("invalid channel")
					},
				},
			} {
				t.Run(consumer.name, func(t *testing.T) {
					rule := consumer.override(tc.newRule(compared))
					assert.NoError(t, rule.Validate(tc.valid))
					assert.Equal(t, &govy.RuleError{
						Message:     "invalid channel",
						Code:        tc.errorCode,
						Description: tc.description,
					}, rule.Validate(tc.invalid))
				})
			}
		})
	}
}

var gtTestCases = []*comparisonTestCase[int]{
	{value: 1, input: 2},
	{value: 1, input: 1, expectedError: "must be greater than '1'"},
	{value: 4, input: 2, expectedError: "must be greater than '4'"},
	{value: 0, input: 1},
	{value: 0, input: 0, expectedError: "must be greater than '0'"},
	{value: 0, input: -1, expectedError: "must be greater than '0'"},
}

func TestGT(t *testing.T) {
	for _, tc := range gtTestCases {
		err := GT(tc.value).Validate(tc.input)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGreaterThan))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestGT_JSONSchema(t *testing.T) {
	t.Parallel()
	assertComparisonJSONSchema(t, "gt", GT[int], gtTestCases)
}

func TestGT_JSONSchema_LargeIntegers(t *testing.T) {
	t.Parallel()

	const boundary uint64 = 1 << 53
	rule := GT(boundary)
	cases := []jsonschematest.Case[uint64]{
		{Name: "below", Input: boundary - 1},
		{Name: "equal", Input: boundary},
		{
			Name: "above loses precision in Ajv", Input: boundary + 1, Valid: true,
			JSONSchemaDifference: "Ajv uses JavaScript numbers, so 2^53 + 1 rounds down to 2^53.",
		},
		{Name: "above exactly representable", Input: boundary + 2, Valid: true},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, rule.Validate(tc.Input) == nil)
	}
	schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[uint64]()).Rules(rule)))
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_gt_large_integer.json", cases)
}

func BenchmarkGT(b *testing.B) {
	for _, tc := range gtTestCases {
		rule := GT(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var gteTestCases = []*comparisonTestCase[int]{
	{value: 1, input: 1},
	{value: 2, input: 4},
	{value: 4, input: 2, expectedError: "must be greater than or equal to '4'"},
	{value: 0, input: 1},
	{value: 0, input: 0},
	{value: 0, input: -1, expectedError: "must be greater than or equal to '0'"},
}

func TestGTE(t *testing.T) {
	for _, tc := range gteTestCases {
		err := GTE(tc.value).Validate(tc.input)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGreaterThanOrEqualTo))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestGTE_JSONSchema(t *testing.T) {
	t.Parallel()
	assertComparisonJSONSchema(t, "gte", GTE[int], gteTestCases)
}

func BenchmarkGTE(b *testing.B) {
	for _, tc := range gteTestCases {
		rule := GTE(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var ltTestCases = []*comparisonTestCase[int]{
	{value: 4, input: 2},
	{value: 1, input: 1, expectedError: "must be less than '1'"},
	{value: 2, input: 4, expectedError: "must be less than '2'"},
	{value: 0, input: -1},
	{value: 0, input: 0, expectedError: "must be less than '0'"},
	{value: 0, input: 1, expectedError: "must be less than '0'"},
}

func TestLT(t *testing.T) {
	for _, tc := range ltTestCases {
		err := LT(tc.value).Validate(tc.input)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLessThan))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestLT_JSONSchema(t *testing.T) {
	t.Parallel()
	assertComparisonJSONSchema(t, "lt", LT[int], ltTestCases)
}

func BenchmarkLT(b *testing.B) {
	for _, tc := range ltTestCases {
		rule := LT(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var lteTestCases = []*comparisonTestCase[int]{
	{value: 1, input: 1},
	{value: 4, input: 2},
	{value: 2, input: 4, expectedError: "must be less than or equal to '2'"},
	{value: 0, input: -1},
	{value: 0, input: 0},
	{value: 0, input: 1, expectedError: "must be less than or equal to '0'"},
}

func TestLTE(t *testing.T) {
	for _, tc := range lteTestCases {
		err := LTE(tc.value).Validate(tc.input)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLessThanOrEqualTo))
		} else {
			assert.NoError(t, err)
		}
	}
}

func TestLTE_JSONSchema(t *testing.T) {
	t.Parallel()
	assertComparisonJSONSchema(t, "lte", LTE[int], lteTestCases)
}

func BenchmarkLTE(b *testing.B) {
	for _, tc := range lteTestCases {
		rule := LTE(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var equalPropertiesTestCases = []*struct {
	run           func() error
	expectedError string
}{
	{
		run: func() error {
			return EqualProperties(CompareDeepEqualFunc, paymentMethodGetters).Validate(paymentMethod{
				Cash:     ptr("2$"),
				Card:     ptr("2$"),
				Transfer: ptr("2$"),
			})
		},
	},
	{
		run: func() error {
			return EqualProperties(CompareFunc, paymentMethodGetters).Validate(paymentMethod{
				Cash:     nil,
				Card:     ptr("2$"),
				Transfer: ptr("2$"),
			})
		},
		expectedError: "all of [Card, Cash, Transfer] properties must be equal, but 'Card' is not equal to 'Cash'",
	},
	{
		run: func() error {
			return EqualProperties(CompareFunc, paymentMethodGetters).Validate(paymentMethod{
				Cash:     nil,
				Card:     nil,
				Transfer: nil,
			})
		},
	},
	{
		run: func() error {
			return EqualProperties(CompareDeepEqualFunc, paymentMethodGetters).Validate(paymentMethod{
				Cash:     ptr("2$"),
				Card:     ptr("2$"),
				Transfer: ptr("3$"),
			})
		},
		expectedError: "all of [Card, Cash, Transfer] properties must be equal, but 'Cash' is not equal to 'Transfer'",
	},
	{
		run: func() error {
			return EqualProperties(CompareDeepEqualFunc, paymentMethodGetters).Validate(paymentMethod{
				Cash:     ptr("1$"),
				Card:     ptr("2$"),
				Transfer: ptr("3$"),
			})
		},
		expectedError: "all of [Card, Cash, Transfer] properties must be equal, but 'Card' is not equal to 'Cash'",
	},
}

func TestEqualProperties(t *testing.T) {
	for _, tc := range equalPropertiesTestCases {
		err := tc.run()
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeEqualProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkEqualProperties(b *testing.B) {
	for _, tc := range equalPropertiesTestCases {
		rule := tc.run()
		for range b.N {
			_ = rule
		}
	}
}

type intRange struct {
	Min int
	Max int
}

var ltPropertiesTestCases = []*struct {
	value         intRange
	expectedError string
}{
	{value: intRange{Min: 1, Max: 10}},
	{value: intRange{Min: 10, Max: 1}, expectedError: "'min' must be less than 'max'"},
	{value: intRange{Min: 5, Max: 5}, expectedError: "'min' must be less than 'max'"},
}

func TestLTProperties(t *testing.T) {
	rule := LTProperties(
		"min", func(r intRange) int { return r.Min },
		"max", func(r intRange) int { return r.Max },
	)
	for _, tc := range ltPropertiesTestCases {
		err := rule.Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLTProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkLTProperties(b *testing.B) {
	rule := LTProperties(
		"min", func(r intRange) int { return r.Min },
		"max", func(r intRange) int { return r.Max },
	)
	for _, tc := range ltPropertiesTestCases {
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

type timeRange struct {
	StartTime time.Time
	EndTime   time.Time
}

var ltComparablePropertiesTestCases = []*struct {
	value         timeRange
	expectedError string
}{
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		expectedError: "'startTime' must be before 'endTime'",
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		expectedError: "'startTime' must be before 'endTime'",
	},
}

func TestLTComparableProperties(t *testing.T) {
	rule := LTComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)
	for _, tc := range ltComparablePropertiesTestCases {
		err := rule.Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLTComparableProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkLTComparableProperties(b *testing.B) {
	rule := LTComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)
	for _, tc := range ltComparablePropertiesTestCases {
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var gtPropertiesTestCases = []*struct {
	value         intRange
	expectedError string
}{
	{value: intRange{Min: 10, Max: 1}},
	{value: intRange{Min: 1, Max: 10}, expectedError: "'min' must be greater than 'max'"},
	{value: intRange{Min: 5, Max: 5}, expectedError: "'min' must be greater than 'max'"},
}

func TestGTProperties(t *testing.T) {
	rule := GTProperties(
		"min", func(r intRange) int { return r.Min },
		"max", func(r intRange) int { return r.Max },
	)
	for _, tc := range gtPropertiesTestCases {
		err := rule.Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGTProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkGTProperties(b *testing.B) {
	rule := GTProperties(
		"min", func(r intRange) int { return r.Min },
		"max", func(r intRange) int { return r.Max },
	)
	for _, tc := range gtPropertiesTestCases {
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var gtComparablePropertiesTestCases = []*struct {
	value         timeRange
	expectedError string
}{
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		expectedError: "'startTime' must be after 'endTime'",
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		expectedError: "'startTime' must be after 'endTime'",
	},
}

func TestGTComparableProperties(t *testing.T) {
	rule := GTComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)
	for _, tc := range gtComparablePropertiesTestCases {
		err := rule.Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGTComparableProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkGTComparableProperties(b *testing.B) {
	rule := GTComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)
	for _, tc := range gtComparablePropertiesTestCases {
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var ltePropertiesTestCases = []*struct {
	value         intRange
	expectedError string
}{
	{value: intRange{Min: 1, Max: 10}},
	{value: intRange{Min: 5, Max: 5}},
	{value: intRange{Min: 10, Max: 1}, expectedError: "'min' must be less than or equal to 'max'"},
}

func TestLTEProperties(t *testing.T) {
	rule := LTEProperties(
		"min", func(r intRange) int { return r.Min },
		"max", func(r intRange) int { return r.Max },
	)
	for _, tc := range ltePropertiesTestCases {
		err := rule.Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLTEProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkLTEProperties(b *testing.B) {
	rule := LTEProperties(
		"min", func(r intRange) int { return r.Min },
		"max", func(r intRange) int { return r.Max },
	)
	for _, tc := range ltePropertiesTestCases {
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var lteComparablePropertiesTestCases = []*struct {
	value         timeRange
	expectedError string
}{
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		expectedError: "'startTime' must be before or equal to 'endTime'",
	},
}

func TestLTEComparableProperties(t *testing.T) {
	rule := LTEComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)
	for _, tc := range lteComparablePropertiesTestCases {
		err := rule.Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLTEComparableProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkLTEComparableProperties(b *testing.B) {
	rule := LTEComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)
	for _, tc := range lteComparablePropertiesTestCases {
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var gtePropertiesTestCases = []*struct {
	value         intRange
	expectedError string
}{
	{value: intRange{Min: 10, Max: 1}},
	{value: intRange{Min: 5, Max: 5}},
	{value: intRange{Min: 1, Max: 10}, expectedError: "'min' must be greater than or equal to 'max'"},
}

func TestGTEProperties(t *testing.T) {
	rule := GTEProperties(
		"min", func(r intRange) int { return r.Min },
		"max", func(r intRange) int { return r.Max },
	)
	for _, tc := range gtePropertiesTestCases {
		err := rule.Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGTEProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkGTEProperties(b *testing.B) {
	rule := GTEProperties(
		"min", func(r intRange) int { return r.Min },
		"max", func(r intRange) int { return r.Max },
	)
	for _, tc := range gtePropertiesTestCases {
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

var gteComparablePropertiesTestCases = []*struct {
	value         timeRange
	expectedError string
}{
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		expectedError: "'startTime' must be after or equal to 'endTime'",
	},
}

func TestGTEComparableProperties(t *testing.T) {
	rule := GTEComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)
	for _, tc := range gteComparablePropertiesTestCases {
		err := rule.Validate(tc.value)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGTEComparableProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkGTEComparableProperties(b *testing.B) {
	rule := GTEComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)
	for _, tc := range gteComparablePropertiesTestCases {
		for range b.N {
			_ = rule.Validate(tc.value)
		}
	}
}

type customComparable struct {
	value int
}

func (c customComparable) Compare(other customComparable) int {
	return cmp.Compare(c.value, other.value)
}

type customComparableRange struct {
	First  customComparable
	Second customComparable
}

func TestIsTemporal(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name     string
		value    any
		expected bool
	}{
		{name: "time.Time", value: now, expected: true},
		{name: "*time.Time", value: &now, expected: true},
		{name: "int", value: 42, expected: false},
		{name: "string", value: "test", expected: false},
		{name: "customComparable", value: customComparable{value: 10}, expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isTemporal(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLTComparablePropertiesUsesBeforeForTime(t *testing.T) {
	rule := LTComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)

	err := rule.Validate(timeRange{
		StartTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})

	assert.Require(t, assert.Error(t, err))
	assert.EqualError(t, err, "'startTime' must be before 'endTime'")
}

func TestLTComparablePropertiesUsesLessThanForCustomType(t *testing.T) {
	rule := LTComparableProperties(
		"first", func(r customComparableRange) customComparable { return r.First },
		"second", func(r customComparableRange) customComparable { return r.Second },
	)

	err := rule.Validate(customComparableRange{
		First:  customComparable{value: 10},
		Second: customComparable{value: 5},
	})

	assert.Require(t, assert.Error(t, err))
	assert.EqualError(t, err, "'first' must be less than 'second'")
}

func TestGTComparablePropertiesUsesAfterForTime(t *testing.T) {
	rule := GTComparableProperties(
		"startTime", func(tr timeRange) time.Time { return tr.StartTime },
		"endTime", func(tr timeRange) time.Time { return tr.EndTime },
	)

	err := rule.Validate(timeRange{
		StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
	})

	assert.Require(t, assert.Error(t, err))
	assert.EqualError(t, err, "'startTime' must be after 'endTime'")
}

func TestGTComparablePropertiesUsesGreaterThanForCustomType(t *testing.T) {
	rule := GTComparableProperties(
		"first", func(r customComparableRange) customComparable { return r.First },
		"second", func(r customComparableRange) customComparable { return r.Second },
	)

	err := rule.Validate(customComparableRange{
		First:  customComparable{value: 5},
		Second: customComparable{value: 10},
	})

	assert.Require(t, assert.Error(t, err))
	assert.EqualError(t, err, "'first' must be greater than 'second'")
}

func TestEqualityRules_JSONSchema_TypedValues(t *testing.T) {
	t.Parallel()

	t.Run("integer", func(t *testing.T) {
		t.Parallel()
		assertEqualityJSONSchemas(t, "integer", 1, []jsonschematest.Case[int]{
			{Name: "equal", Input: 1, Valid: true},
			{Name: "different", Input: 2},
		})
	})
	t.Run("number", func(t *testing.T) {
		t.Parallel()
		assertEqualityJSONSchemas(t, "number", 1.1, []jsonschematest.Case[float64]{
			{Name: "equal", Input: 1.1, Valid: true},
			{Name: "different", Input: 1.3},
		})
	})
	t.Run("array", func(t *testing.T) {
		t.Parallel()
		assertEqualityJSONSchemas(t, "array", [2]int{1, 2}, []jsonschematest.Case[[2]int]{
			{Name: "equal", Input: [2]int{1, 2}, Valid: true},
			{Name: "different", Input: [2]int{2, 1}},
		})
	})
	t.Run("object", func(t *testing.T) {
		t.Parallel()
		type value struct {
			Name     string `json:"name"`
			Revision int    `json:"-"`
		}
		assertEqualityJSONSchemas(t, "object", value{Name: "first", Revision: 1}, []jsonschematest.Case[value]{
			{Name: "equal", Input: value{Name: "first", Revision: 1}, Valid: true},
			{Name: "different", Input: value{Name: "second", Revision: 1}},
			{
				Name:                 "ignored field differs",
				Input:                value{Name: "first", Revision: 2},
				JSONSchemaDifference: "Go compares the ignored Revision field; JSON Schema compares only the serialized object.",
			},
		})
	})
	t.Run("pointer", func(t *testing.T) {
		t.Parallel()
		compared := ptr(1)
		assertEqualityJSONSchemas(t, "pointer", compared, []jsonschematest.Case[*int]{
			{Name: "same pointer", Input: compared, Valid: true},
			{
				Name: "same value at another address", Input: ptr(1),
				JSONSchemaDifference: "Go compares pointer identity; JSON Schema compares the pointed-to value.",
			},
			{Name: "different value", Input: ptr(2)},
		})
	})
}

func assertEqualityJSONSchemas[T comparable](
	t *testing.T,
	name string,
	compared T,
	cases []jsonschematest.Case[T],
) {
	t.Helper()

	for _, tc := range []struct {
		name string
		rule govy.Rule[T]
		code govy.ErrorCode
	}{
		{name: "eq", rule: EQ(compared), code: ErrorCodeEqualTo},
		{name: "neq", rule: NEQ(compared), code: ErrorCodeNotEqualTo},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			inputs := make([]jsonschematest.Case[T], len(cases))
			for i, input := range cases {
				if tc.name == "neq" {
					input.Valid = !input.Valid
				}
				inputs[i] = input
				err := tc.rule.Validate(input.Input)
				if input.Valid {
					assert.NoError(t, err)
				} else {
					assert.True(t, govy.HasErrorCode(err, tc.code))
				}
			}
			schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[T]()).Rules(tc.rule)))
			assert.Require(t, assert.NoError(t, err))
			fixture := fmt.Sprintf("testdata/jsonschema/expected_%s_typed_%s.json", tc.name, name)
			jsonschematest.Assert(t, schema, fixture, inputs)
		})
	}
}

func assertComparisonJSONSchema[T comparable](
	t *testing.T,
	name string,
	newRule func(T) govy.Rule[T],
	tests []*comparisonTestCase[T],
) {
	t.Helper()

	groups := make(map[string][]jsonschematest.Case[T])
	values := make(map[string]T)
	for i, tc := range tests {
		schemaName := tc.schemaName
		if schemaName == "" {
			schemaName = fmt.Sprint(tc.value)
		}
		values[schemaName] = tc.value
		groups[schemaName] = append(groups[schemaName], jsonschematest.Case[T]{
			Name:                 fmt.Sprintf("%d/%v", i, tc.input),
			Input:                tc.input,
			Valid:                tc.expectedError == "",
			JSONSchemaDifference: tc.jsonSchemaDifference,
		})
	}
	for schemaName, cases := range groups {
		t.Run(schemaName, func(t *testing.T) {
			t.Parallel()
			rule := newRule(values[schemaName])
			schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[T]()).Rules(rule)))
			assert.Require(t, assert.NoError(t, err))
			fixture := fmt.Sprintf("testdata/jsonschema/expected_%s_%s.json", name, schemaName)
			jsonschematest.Assert(t, schema, fixture, cases)
		})
	}
}

func TestOrderedStringRules_JSONSchema(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rule  govy.Rule[string]
		cases []jsonschematest.Case[string]
	}{
		{
			name: "GT",
			rule: GT("middle"),
			cases: []jsonschematest.Case[string]{
				{Name: "after", Input: "zulu", Valid: true},
				{
					Name:                 "equal",
					Input:                "middle",
					JSONSchemaDifference: "JSON Schema has no string ordering constraint.",
				},
				{
					Name:                 "before",
					Input:                "alpha",
					JSONSchemaDifference: "JSON Schema has no string ordering constraint.",
				},
			},
		},
		{
			name: "GTE",
			rule: GTE("middle"),
			cases: []jsonschematest.Case[string]{
				{Name: "after", Input: "zulu", Valid: true},
				{Name: "equal", Input: "middle", Valid: true},
				{
					Name:                 "before",
					Input:                "alpha",
					JSONSchemaDifference: "JSON Schema has no string ordering constraint.",
				},
			},
		},
		{
			name: "LT",
			rule: LT("middle"),
			cases: []jsonschematest.Case[string]{
				{Name: "before", Input: "alpha", Valid: true},
				{
					Name:                 "equal",
					Input:                "middle",
					JSONSchemaDifference: "JSON Schema has no string ordering constraint.",
				},
				{Name: "after", Input: "zulu", JSONSchemaDifference: "JSON Schema has no string ordering constraint."},
			},
		},
		{
			name: "LTE",
			rule: LTE("middle"),
			cases: []jsonschematest.Case[string]{
				{Name: "before", Input: "alpha", Valid: true},
				{Name: "equal", Input: "middle", Valid: true},
				{Name: "after", Input: "zulu", JSONSchemaDifference: "JSON Schema has no string ordering constraint."},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			for _, input := range tc.cases {
				assert.Equal(t, input.Valid, tc.rule.Validate(input.Input) == nil)
			}
			schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(tc.rule)))
			assert.Require(t, assert.NoError(t, err))
			jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_ordered_string.json", tc.cases)
		})
	}
}

func TestOrderedNumberRules_JSONSchema_FractionalBounds(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		rule  govy.Rule[float64]
		cases []jsonschematest.Case[float64]
	}{
		{
			name: "gt", rule: GT(1.5),
			cases: []jsonschematest.Case[float64]{
				{Name: "below", Input: 1.25},
				{Name: "equal", Input: 1.5},
				{Name: "above", Input: 1.75, Valid: true},
			},
		},
		{
			name: "gte", rule: GTE(1.5),
			cases: []jsonschematest.Case[float64]{
				{Name: "below", Input: 1.25},
				{Name: "equal", Input: 1.5, Valid: true},
				{Name: "above", Input: 1.75, Valid: true},
			},
		},
		{
			name: "lt", rule: LT(1.5),
			cases: []jsonschematest.Case[float64]{
				{Name: "below", Input: 1.25, Valid: true},
				{Name: "equal", Input: 1.5},
				{Name: "above", Input: 1.75},
			},
		},
		{
			name: "lte", rule: LTE(1.5),
			cases: []jsonschematest.Case[float64]{
				{Name: "below", Input: 1.25, Valid: true},
				{Name: "equal", Input: 1.5, Valid: true},
				{Name: "above", Input: 1.75},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			for _, input := range tc.cases {
				assert.Equal(t, input.Valid, tc.rule.Validate(input.Input) == nil)
			}
			schema, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[float64]()).Rules(tc.rule)))
			assert.Require(t, assert.NoError(t, err))
			fixture := fmt.Sprintf("testdata/jsonschema/expected_%s_fractional.json", tc.name)
			jsonschematest.Assert(t, schema, fixture, tc.cases)
		})
	}
}

func TestComparisonRules_JSONSchema_NonJSONValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		value         any
		expectedError string
	}{
		{name: "channel", value: make(chan int), expectedError: "json: unsupported type: chan int"},
		{name: "complex", value: complex(1, 2), expectedError: "json: unsupported type: complex128"},
		{name: "NaN", value: math.NaN(), expectedError: "json: unsupported value: NaN"},
		{name: "positive infinity", value: math.Inf(1), expectedError: "json: unsupported value: +Inf"},
		{name: "negative infinity", value: math.Inf(-1), expectedError: "json: unsupported value: -Inf"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			for _, rule := range []struct {
				name string
				code govy.ErrorCode
				rule govy.Rule[any]
			}{
				{name: "EQ", code: ErrorCodeEqualTo, rule: EQ(tc.value)},
				{name: "NEQ", code: ErrorCodeNotEqualTo, rule: NEQ(tc.value)},
			} {
				t.Run(rule.name, func(t *testing.T) {
					t.Parallel()
					_, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[any]()).Rules(rule.rule)))
					assert.EqualError(t, err, fmt.Sprintf(
						"failed to build JSON Schema for \"$\" property and %q rule: marshal value: %s",
						rule.code, tc.expectedError,
					))
				})
			}
		})
	}
}

func TestOrderedNumberRules_JSONSchema_NonFiniteBounds(t *testing.T) {
	t.Parallel()

	for name, value := range map[string]float64{
		"NaN":  math.NaN(),
		"+Inf": math.Inf(1),
		"-Inf": math.Inf(-1),
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			for _, rule := range []struct {
				name string
				code govy.ErrorCode
				rule govy.Rule[float64]
			}{
				{name: "GT", code: ErrorCodeGreaterThan, rule: GT(value)},
				{name: "GTE", code: ErrorCodeGreaterThanOrEqualTo, rule: GTE(value)},
				{name: "LT", code: ErrorCodeLessThan, rule: LT(value)},
				{name: "LTE", code: ErrorCodeLessThanOrEqualTo, rule: LTE(value)},
			} {
				t.Run(rule.name, func(t *testing.T) {
					t.Parallel()
					_, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[float64]()).Rules(rule.rule)))
					assert.EqualError(t, err, fmt.Sprintf(
						"failed to build JSON Schema for \"$\" property and %q rule: marshal value: json: unsupported value: %s",
						rule.code,
						name,
					))
				})
			}
		})
	}
}

type stringEncodedComparisonNumber float64

// MarshalJSON encodes numeric test values as JSON strings.
func (n stringEncodedComparisonNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(float64(n)))
}

func TestOrderedNumberRules_JSONSchema_NonNumericEncoding(t *testing.T) {
	t.Parallel()

	const compared = stringEncodedComparisonNumber(1.5)
	for _, tc := range []struct {
		name string
		code govy.ErrorCode
		rule govy.Rule[stringEncodedComparisonNumber]
	}{
		{name: "GT", code: ErrorCodeGreaterThan, rule: GT(compared)},
		{name: "GTE", code: ErrorCodeGreaterThanOrEqualTo, rule: GTE(compared)},
		{name: "LT", code: ErrorCodeLessThan, rule: LT(compared)},
		{name: "LTE", code: ErrorCodeLessThanOrEqualTo, rule: LTE(compared)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[stringEncodedComparisonNumber]()).Rules(tc.rule)))
			assert.EqualError(t, err, fmt.Sprintf(
				"failed to build JSON Schema for \"$\" property and %q rule: value marshaled as string instead of a JSON number",
				tc.code,
			))
		})
	}
}
