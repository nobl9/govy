package rules

import (
	"fmt"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
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

func TestForbidden_JSONSchema(t *testing.T) {
	t.Parallel()
	validator := govy.New(
		govy.For(func(v map[string]string) string { return v["value"] }).
			WithName("value").
			Rules(Forbidden[string]()),
	)
	schema, err := govy.JSONSchema(validator)
	assert.Require(t, assert.NoError(t, err))
	cases := make([]jsonschematest.Case[map[string]string], 0, len(forbiddenTestCases)+1)
	for _, tc := range forbiddenTestCases {
		cases = append(cases, jsonschematest.Case[map[string]string]{
			Name:  fmt.Sprintf("present/%q", tc.in),
			Input: map[string]string{"value": tc.in},
			Valid: !tc.shouldFail,
		})
	}
	cases = append(cases, jsonschematest.Case[map[string]string]{
		Name:  "missing",
		Input: map[string]string{},
		Valid: true,
	})
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "testdata/jsonschema/expected_forbidden.json", cases)
}

func TestForbidden_JSONSchemaUnsupportedValue(t *testing.T) {
	t.Parallel()
	type value struct {
		Channel chan struct{}
	}
	validator := govy.New(govy.For(govy.GetSelf[value]()).Rules(Forbidden[value]()))
	schema, err := govy.JSONSchema(validator)
	assert.True(t, schema == nil)
	assert.EqualError(
		t,
		err,
		`failed to build JSON Schema for "$" property and "forbidden" rule: marshal value: json: unsupported type: chan struct {}`,
	)
}

func BenchmarkForbidden(b *testing.B) {
	for _, tc := range forbiddenTestCases {
		rule := Forbidden[string]()
		for range b.N {
			_ = rule.Validate(tc.in)
		}
	}
}
