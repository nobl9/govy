package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
)

var forbiddenTestCases = []*struct {
	in         string
	shouldFail bool
}{
	{"", false},
	{"test", true},
}

func TestForbidden(t *testing.T) {
	for _, tc := range forbiddenTestCases {
		err := Forbidden[string]().Validate(tc.in)
		if tc.shouldFail {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, "property is forbidden")
			assert.Equal(t, true, govy.HasErrorCode(err, ErrorCodeForbidden))
		} else {
			assert.Equal(t, nil, err)
		}
	}
}

func BenchmarkForbidden(b *testing.B) {
	for range b.N {
		for _, tc := range forbiddenTestCases {
			_ = Forbidden[string]().Validate(tc.in)
		}
	}
}
