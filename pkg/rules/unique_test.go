package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var sliceUniqueTestCases = []*struct {
	slice         []string
	hashFunc      HashFunction[string, string]
	constraints   []string
	expectedError string
	shouldFail    bool
}{
	{
		slice:    []string{"a"},
		hashFunc: HashFuncSelf[string](),
	},
	{
		slice:    []string{"a", "b", "c"},
		hashFunc: HashFuncSelf[string](),
	},
	{
		slice:    []string{"a", "b", "c"},
		hashFunc: func(v string) string { return v },
	},
	{
		slice:    []string{"a", "b", "c", "d", "e", "f"},
		hashFunc: HashFuncSelf[string](),
	},
	{
		slice:         []string{"a", "a"},
		hashFunc:      HashFuncSelf[string](),
		expectedError: "elements are not unique, 1st and 2nd elements collide",
		shouldFail:    true,
	},
	{
		slice:         []string{"a", "b", "a"},
		hashFunc:      HashFuncSelf[string](),
		expectedError: "elements are not unique, 1st and 3rd elements collide",
		shouldFail:    true,
	},
	{
		slice:         []string{"a", "b", "c", "b"},
		hashFunc:      HashFuncSelf[string](),
		expectedError: "elements are not unique, 2nd and 4th elements collide",
		shouldFail:    true,
	},
	{
		slice:         []string{"a", "b", "c", "b"},
		hashFunc:      HashFuncSelf[string](),
		constraints:   []string{"values must be unique"},
		expectedError: "elements are not unique, 2nd and 4th elements collide based on constraints: values must be unique", // nolint: lll
		shouldFail:    true,
	},
	{
		slice:         []string{"a", "b", "c", "b"},
		hashFunc:      HashFuncSelf[string](),
		constraints:   []string{"constraint 1", "constraint 2"},
		expectedError: "elements are not unique, 2nd and 4th elements collide based on constraints: constraint 1, constraint 2", // nolint: lll
		shouldFail:    true,
	},
	{
		slice:         []string{"a", "b", "c", "b"},
		hashFunc:      func(v string) string { return v },
		expectedError: "elements are not unique, 2nd and 4th elements collide",
		shouldFail:    true,
	},
}

func TestSliceUnique(t *testing.T) {
	for _, tc := range sliceUniqueTestCases {
		err := SliceUnique(tc.hashFunc, tc.constraints...).Validate(tc.slice)
		if tc.shouldFail {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceUnique))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkSliceUnique(b *testing.B) {
	for _, tc := range sliceUniqueTestCases {
		rule := SliceUnique(tc.hashFunc, tc.constraints...)
		for range b.N {
			_ = rule.Validate(tc.slice)
		}
	}
}
