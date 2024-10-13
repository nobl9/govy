package rules

import (
	"testing"

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
	for range b.N {
		for _, tc := range eqTestCases {
			_ = EQ(tc.value).Validate(tc.input)
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
	for range b.N {
		for _, tc := range neqTestCases {
			_ = NEQ(tc.value).Validate(tc.input)
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
	for range b.N {
		for _, tc := range gtTestCases {
			_ = GT(tc.value).Validate(tc.input)
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
	for range b.N {
		for _, tc := range gteTestCases {
			_ = GTE(tc.value).Validate(tc.input)
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
	for range b.N {
		for _, tc := range ltTestCases {
			_ = LT(tc.value).Validate(tc.input)
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
	for range b.N {
		for _, tc := range lteTestCases {
			_ = LTE(tc.value).Validate(tc.input)
		}
	}
}
