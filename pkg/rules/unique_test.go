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

type currency struct {
	Dollar string
	Euro   string
	Pound  string
}

var uniquePropertiesGetters = map[string]func(c currency) string{
	"Dollar": func(c currency) string { return c.Dollar },
	"Euro":   func(c currency) string { return c.Euro },
	"Pound":  func(c currency) string { return c.Pound },
}

var uniquePropertiesTestCases = []*struct {
	run           func() error
	expectedError string
}{
	{
		run: func() error {
			return UniqueProperties(HashFuncSelf[string](), uniquePropertiesGetters).Validate(currency{
				Dollar: "unique1",
				Euro:   "unique2",
				Pound:  "unique3",
			})
		},
	},
	{
		run: func() error {
			return UniqueProperties(HashFuncSelf[string](), uniquePropertiesGetters).Validate(currency{
				Dollar: "",
				Euro:   "",
				Pound:  "",
			})
		},
		expectedError: "all of [Dollar, Euro, Pound] properties must be unique, but 'Dollar' collides with 'Euro'",
	},
	{
		run: func() error {
			return UniqueProperties(HashFuncSelf[string](), uniquePropertiesGetters).Validate(currency{
				Dollar: "same",
				Euro:   "same",
				Pound:  "unique",
			})
		},
		expectedError: "all of [Dollar, Euro, Pound] properties must be unique, but 'Dollar' collides with 'Euro'",
	},
	{
		run: func() error {
			return UniqueProperties(HashFuncSelf[string](), uniquePropertiesGetters).Validate(currency{
				Dollar: "same",
				Euro:   "unique",
				Pound:  "same",
			})
		},
		expectedError: "all of [Dollar, Euro, Pound] properties must be unique, but 'Dollar' collides with 'Pound'",
	},
	{
		run: func() error {
			return UniqueProperties(HashFuncSelf[string](), uniquePropertiesGetters).Validate(currency{
				Dollar: "unique",
				Euro:   "same",
				Pound:  "same",
			})
		},
		expectedError: "all of [Dollar, Euro, Pound] properties must be unique, but 'Euro' collides with 'Pound'",
	},
	{
		run: func() error {
			return UniqueProperties(
				HashFuncSelf[string](),
				uniquePropertiesGetters,
				"each currency must be unique",
			).Validate(currency{
				Dollar: "same",
				Euro:   "same",
				Pound:  "unique",
			})
		},
		expectedError: "all of [Dollar, Euro, Pound] properties must be unique, but 'Dollar' collides with 'Euro'" +
			", based on constraints: each currency must be unique",
	},
	{
		run: func() error {
			return UniqueProperties(
				HashFuncSelf[string](),
				uniquePropertiesGetters,
				"each currency must be unique", "another constraint",
			).Validate(currency{
				Dollar: "same",
				Euro:   "same",
				Pound:  "unique",
			})
		},
		expectedError: "all of [Dollar, Euro, Pound] properties must be unique, but 'Dollar' collides with 'Euro'" +
			", based on constraints: each currency must be unique, another constraint",
	},
}

func TestUniqueProperties(t *testing.T) {
	for _, tc := range uniquePropertiesTestCases {
		err := tc.run()
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeSliceUnique))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkUniqueProperties(b *testing.B) {
	for _, tc := range uniquePropertiesTestCases {
		for range b.N {
			_ = tc.run()
		}
	}
}
