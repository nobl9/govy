package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

var oneOfTestCases = []*struct {
	in            string
	options       []string
	expectedError string
}{
	{"that", []string{"this", "that"}, ""},
	{"those", []string{"this", "that"}, "must be one of [this, that]"},
}

func TestOneOf(t *testing.T) {
	for _, tc := range oneOfTestCases {
		err := OneOf(tc.options...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeOneOf))
		} else {
			assert.NoError(t, err)
		}
	}
}

type paymentMethod struct {
	Cash     *string
	Card     *string
	Transfer *string
}

func BenchmarkOneOf(b *testing.B) {
	for range b.N {
		for _, tc := range oneOfTestCases {
			_ = OneOf(tc.options...).Validate(tc.in)
		}
	}
}

var mutuallyExclusiveTestCases = []*struct {
	required      bool
	paymentMethod paymentMethod
	expectedError string
}{
	{
		required: true,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: nil,
		},
	},
	{
		required: false,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		},
	},
	{
		required: true,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Transfer] properties are mutually exclusive, provide only one of them",
	},
	{
		required: false,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Transfer] properties are mutually exclusive, provide only one of them",
	},
	{
		required: true,
		paymentMethod: paymentMethod{
			Cash:     ptr("2$"),
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Cash, Transfer] properties are mutually exclusive, provide only one of them",
	},
	{
		required: false,
		paymentMethod: paymentMethod{
			Cash:     ptr("2$"),
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Cash, Transfer] properties are mutually exclusive, provide only one of them",
	},
	{
		required: true,
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		},
		expectedError: "one of [Card, Cash, Transfer] properties must be set, none was provided",
	},
}

var paymentMethodGetters = map[string]func(p paymentMethod) any{
	"Cash":     func(p paymentMethod) any { return p.Cash },
	"Card":     func(p paymentMethod) any { return p.Card },
	"Transfer": func(p paymentMethod) any { return p.Transfer },
}

func TestMutuallyExclusive(t *testing.T) {
	for _, tc := range mutuallyExclusiveTestCases {
		err := MutuallyExclusive(tc.required, paymentMethodGetters).Validate(tc.paymentMethod)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMutuallyExclusive))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkMutuallyExclusive(b *testing.B) {
	for range b.N {
		for _, tc := range mutuallyExclusiveTestCases {
			_ = MutuallyExclusive(tc.required, paymentMethodGetters).Validate(tc.paymentMethod)
		}
	}
}

var oneOfPropertiesTestCases = []*struct {
	paymentMethod paymentMethod
	expectedError string
}{
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: nil,
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     ptr("1$"),
			Card:     ptr("2$"),
			Transfer: nil,
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     ptr("1$"),
			Card:     ptr("2$"),
			Transfer: ptr("3$"),
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		},
		expectedError: "one of [Card, Cash, Transfer] properties must be set, none was provided",
	},
}

func TestOneOfProperties(t *testing.T) {
	for _, tc := range oneOfPropertiesTestCases {
		err := OneOfProperties(paymentMethodGetters).Validate(tc.paymentMethod)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeOneOfProperties))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkOneOfProperties(b *testing.B) {
	for range b.N {
		for _, tc := range oneOfPropertiesTestCases {
			_ = OneOfProperties(paymentMethodGetters).Validate(tc.paymentMethod)
		}
	}
}

func ptr[T any](v T) *T { return &v }
