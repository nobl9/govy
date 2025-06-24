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
	{"those", []string{"this", "that"}, "must be one of: this, that"},
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

func BenchmarkOneOf(b *testing.B) {
	for _, tc := range oneOfTestCases {
		rule := OneOf(tc.options...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

var notOneOfTestCases = []*struct {
	in            string
	options       []string
	expectedError string
}{
	{"that", []string{"this", "that"}, "must not be one of: this, that"},
	{"those", []string{"this", "that"}, ""},
}

func TestNotOneOf(t *testing.T) {
	for _, tc := range notOneOfTestCases {
		err := NotOneOf(tc.options...).Validate(tc.in)
		if tc.expectedError != "" {
			assert.Require(t, assert.Error(t, err))
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeNotOneOf))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkNotOneOf(b *testing.B) {
	for _, tc := range notOneOfTestCases {
		rule := NotOneOf(tc.options...)
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}

type paymentMethod struct {
	Cash     *string
	Card     *string
	Transfer *string
}

var paymentMethodGetters = map[string]func(p paymentMethod) any{
	"Cash":     func(p paymentMethod) any { return p.Cash },
	"Card":     func(p paymentMethod) any { return p.Card },
	"Transfer": func(p paymentMethod) any { return p.Transfer },
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
	for _, tc := range oneOfPropertiesTestCases {
		rule := OneOfProperties(paymentMethodGetters)
		for range b.N {
			_ = rule.Validate(tc.paymentMethod)
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
	for _, tc := range mutuallyExclusiveTestCases {
		rule := MutuallyExclusive(tc.required, paymentMethodGetters)
		for range b.N {
			_ = rule.Validate(tc.paymentMethod)
		}
	}
}

var mutuallyDependentTestCases = []*struct {
	paymentMethod paymentMethod
	expectedError string
}{
	{
		paymentMethod: paymentMethod{
			Cash:     ptr("2$"),
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		},
	},
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: nil,
		},
		expectedError: "[Card, Cash, Transfer] properties are mutually dependent, since [Card] is provided, [Cash, Transfer] properties must also be set",
	},
	{
		paymentMethod: paymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: ptr("2$"),
		},
		expectedError: "[Card, Cash, Transfer] properties are mutually dependent, since [Card, Transfer] are provided, [Cash] property must also be set",
	},
}

func TestMutuallyDependent(t *testing.T) {
	for _, tc := range mutuallyDependentTestCases {
		err := MutuallyDependent(paymentMethodGetters).Validate(tc.paymentMethod)
		if tc.expectedError != "" {
			assert.EqualError(t, err, tc.expectedError)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMutuallyDependent))
		} else {
			assert.NoError(t, err)
		}
	}
}

func BenchmarkMutuallyDependent(b *testing.B) {
	for _, tc := range mutuallyDependentTestCases {
		rule := MutuallyDependent(paymentMethodGetters)
		for range b.N {
			_ = rule.Validate(tc.paymentMethod)
		}
	}
}

func ptr[T any](v T) *T { return &v }
