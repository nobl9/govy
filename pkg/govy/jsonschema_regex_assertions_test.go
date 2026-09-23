package govy_test

import (
	"regexp"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_QuantifiedAssertions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		pattern string
		valid   string
		invalid string
	}{
		{name: "start", pattern: "^*a", valid: "ba", invalid: "b"},
		{name: "boundary", pattern: `\b+word`, valid: "word", invalid: "sword"},
		{name: "end", pattern: "a$+", valid: "a", invalid: "ab"},
		{name: "counted", pattern: "^{1,2}a", valid: "a", invalid: "ba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			validator := govy.New(govy.For(govy.GetSelf[string]()).
				Rules(rules.StringMatchRegexp(regexp.MustCompile("(?-m)" + tt.pattern))))
			schema, err := govy.JSONSchema(validator)
			assert.Require(t, assert.NoError(t, err))
			cases := []jsonschematest.Case[string]{
				{Name: "matches", Input: tt.valid, Valid: true},
				{Name: "does not match", Input: tt.invalid},
			}
			for _, tc := range cases {
				assert.Equal(t, tc.Valid, validator.Validate(tc.Input) == nil)
			}
			jsonschematest.Assert(t, schema, "test_data/expected_quantified_"+tt.name+"_json_schema.json", cases)
		})
	}
}
