package govy_test

import (
	"testing"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func TestValidatorForSlice(t *testing.T) {
	t.Run("no errors", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return "test" }).
				WithName("test").
				Rules(govy.NewRule(func(v string) error { return nil })),
		)
		vs := govy.NewForSlice(v)
		err := vs.Validate([]mockValidatorStruct{{}})
		assert.NoError(t, err)
	})

	t.Run("errors", func(t *testing.T) {
		v := govy.New(
			govy.For(func(m mockValidatorStruct) string { return m.Field }).
				WithName("Field").
				Required(),
		).WithName("mock")
		vs := govy.NewForSlice(v)
		errs := mustValidatorErrors(t, vs.Validate([]mockValidatorStruct{
			{Field: "0"},
			{},
			{Field: "2"},
			{},
		}))
		assert.Require(t, assert.Len(t, errs, 2))
		assert.Require(t, assert.Len(t, errs[0].Errors, 1))
		assert.Equal(t, govy.ValidatorErrors{
			{
				Name:       "mock",
				SliceIndex: ptr(1),
				Errors: govy.PropertyErrors{
					&govy.PropertyError{
						PropertyName: "Field",
						Errors: []*govy.RuleError{{
							Message: internal.RequiredErrorMessage,
							Code:    rules.ErrorCodeRequired,
						}},
					},
				},
			},
			{
				Name:       "mock",
				SliceIndex: ptr(3),
				Errors: govy.PropertyErrors{
					&govy.PropertyError{
						PropertyName: "Field",
						Errors: []*govy.RuleError{{
							Message: internal.RequiredErrorMessage,
							Code:    rules.ErrorCodeRequired,
						}},
					},
				},
			},
		}, errs)
	})
}

func mustValidatorErrors(t *testing.T, err error) govy.ValidatorErrors {
	t.Helper()
	return mustErrorType[govy.ValidatorErrors](t, err)
}
