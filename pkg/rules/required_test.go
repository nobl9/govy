package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var requiredTestCases = []*struct {
	in         any
	shouldFail bool
}{
	{1, false},
	{"s", false},
	{0.1, false},
	{[]int{}, false},
	{map[string]int{}, false},
	{nil, true},
	{struct{}{}, true},
	{"", true},
	{false, true},
	{0, true},
	{0.0, true},
}

func TestRequired(t *testing.T) {
	for _, tc := range requiredTestCases {
		err := Required[any]().Validate(tc.in)
		if tc.shouldFail {
			assert.Require(t, assert.Error(t, err))
			assert.True(t, govy.HasErrorCode(err, ErrorCodeRequired))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkRequired(b *testing.B) {
	for _, tc := range requiredTestCases {
		rule := Required[any]()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}
