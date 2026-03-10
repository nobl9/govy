package rules

import (
	"testing"
	"time"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var eqTestCases = []*struct {
	value         any
	input         any
	expectedError string
}{
	{value: 1, input: 1},
	{value: 1.1, input: 1.3, expectedError: "should be equal to '1.1'"},
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

func BenchmarkEQ(b *testing.B) {
	for _, tc := range eqTestCases {
		rule := EQ(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var neqTestCases = []*struct {
	value         any
	input         any
	expectedError string
}{
	{value: 1.1, input: 1.3},
	{value: 1.1, input: 1.1, expectedError: "should be not equal to '1.1'"},
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

func BenchmarkNEQ(b *testing.B) {
	for _, tc := range neqTestCases {
		rule := NEQ(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var gtTestCases = []*struct {
	value         int
	input         int
	expectedError string
}{
	{value: 1, input: 2},
	{value: 1, input: 1, expectedError: "should be greater than '1'"},
	{value: 4, input: 2, expectedError: "should be greater than '4'"},
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

func BenchmarkGT(b *testing.B) {
	for _, tc := range gtTestCases {
		rule := GT(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var gteTestCases = []*struct {
	value         int
	input         int
	expectedError string
}{
	{value: 1, input: 1},
	{value: 2, input: 4},
	{value: 4, input: 2, expectedError: "should be greater than or equal to '4'"},
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

func BenchmarkGTE(b *testing.B) {
	for _, tc := range gteTestCases {
		rule := GTE(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var ltTestCases = []*struct {
	value         int
	input         int
	expectedError string
}{
	{value: 4, input: 2},
	{value: 1, input: 1, expectedError: "should be less than '1'"},
	{value: 2, input: 4, expectedError: "should be less than '2'"},
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

func BenchmarkLT(b *testing.B) {
	for _, tc := range ltTestCases {
		rule := LT(tc.value)
		for range b.N {
			_ = rule.Validate(tc.input)
		}
	}
}

var lteTestCases = []*struct {
	value         int
	input         int
	expectedError string
}{
	{value: 1, input: 1},
	{value: 4, input: 2},
	{value: 2, input: 4, expectedError: "should be less than or equal to '2'"},
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
		expectedError: "'startTime' must be less than 'endTime'",
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		expectedError: "'startTime' must be less than 'endTime'",
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
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLTProperties))
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
		expectedError: "'startTime' must be greater than 'endTime'",
	},
	{
		value: timeRange{
			StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		expectedError: "'startTime' must be greater than 'endTime'",
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
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGTProperties))
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
		expectedError: "'startTime' must be less than or equal to 'endTime'",
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
			assert.True(t, govy.HasErrorCode(err, ErrorCodeLTEProperties))
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
		expectedError: "'startTime' must be greater than or equal to 'endTime'",
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
			assert.True(t, govy.HasErrorCode(err, ErrorCodeGTEProperties))
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
