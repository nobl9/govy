package rules

import (
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/govy"
)

func TestOneOf(t *testing.T) {
	t.Run("passes", func(t *testing.T) {
		err := OneOf("this", "that").Validate("that")
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := OneOf("this", "that").Validate("those")
		assert.Require(t, assert.Error(t, err))
		assert.EqualError(t, err, "must be one of [this, that]")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeOneOf))
	})
}

func TestMutuallyExclusive(t *testing.T) {
	type PaymentMethod struct {
		Cash     *string
		Card     *string
		Transfer *string
	}
	getters := map[string]func(p PaymentMethod) any{
		"Cash":     func(p PaymentMethod) any { return p.Cash },
		"Card":     func(p PaymentMethod) any { return p.Card },
		"Transfer": func(p PaymentMethod) any { return p.Transfer },
	}

	t.Run("passes with required", func(t *testing.T) {
		err := MutuallyExclusive(true, getters).Validate(PaymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: nil,
		})
		assert.NoError(t, err)
	})
	t.Run("passes with non-required", func(t *testing.T) {
		err := MutuallyExclusive(false, getters).Validate(PaymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		for _, required := range []bool{true, false} {
			err := MutuallyExclusive(required, getters).Validate(PaymentMethod{
				Cash:     nil,
				Card:     ptr("2$"),
				Transfer: ptr("2$"),
			})
			assert.EqualError(t, err, "[Card, Transfer] properties are mutually exclusive, provide only one of them")
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMutuallyExclusive))
		}
	})
	t.Run("fails, multiple conflicts", func(t *testing.T) {
		for _, required := range []bool{true, false} {
			err := MutuallyExclusive(required, getters).Validate(PaymentMethod{
				Cash:     ptr("2$"),
				Card:     ptr("2$"),
				Transfer: ptr("2$"),
			})
			assert.EqualError(
				t,
				err,
				"[Card, Cash, Transfer] properties are mutually exclusive, provide only one of them",
			)
			assert.True(t, govy.HasErrorCode(err, ErrorCodeMutuallyExclusive))
		}
	})
	t.Run("required fails", func(t *testing.T) {
		err := MutuallyExclusive(true, getters).Validate(PaymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		})
		assert.EqualError(t, err, "one of [Card, Cash, Transfer] properties must be set, none was provided")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeMutuallyExclusive))
	})
}

func TestOneOfProperties(t *testing.T) {
	type PaymentMethod struct {
		Cash     *string
		Card     *string
		Transfer *string
	}
	getters := map[string]func(p PaymentMethod) any{
		"Cash":     func(p PaymentMethod) any { return p.Cash },
		"Card":     func(p PaymentMethod) any { return p.Card },
		"Transfer": func(p PaymentMethod) any { return p.Transfer },
	}
	t.Run("passes", func(t *testing.T) {
		err := OneOfProperties(getters).Validate(PaymentMethod{
			Cash:     nil,
			Card:     ptr("2$"),
			Transfer: nil,
		})
		assert.NoError(t, err)
		err = OneOfProperties(getters).Validate(PaymentMethod{
			Cash:     ptr("1$"),
			Card:     ptr("2$"),
			Transfer: nil,
		})
		assert.NoError(t, err)
		err = OneOfProperties(getters).Validate(PaymentMethod{
			Cash:     ptr("1$"),
			Card:     ptr("2$"),
			Transfer: ptr("3$"),
		})
		assert.NoError(t, err)
	})
	t.Run("fails", func(t *testing.T) {
		err := OneOfProperties(getters).Validate(PaymentMethod{
			Cash:     nil,
			Card:     nil,
			Transfer: nil,
		})
		assert.EqualError(t, err, "one of [Card, Cash, Transfer] properties must be set, none was provided")
		assert.True(t, govy.HasErrorCode(err, ErrorCodeOneOfProperties))
	})
}

func ptr[T any](v T) *T { return &v }
