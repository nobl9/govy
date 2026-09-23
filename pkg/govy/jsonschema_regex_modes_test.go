package govy_test

import (
	"regexp"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/internal/jsonschematest"
	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func TestJSONSchema_RegexpCompilationModes(t *testing.T) {
	t.Parallel()

	for _, pattern := range []string{`^foo$`, `[^a]`} {
		for mode, compile := range map[string]func(string) *regexp.Regexp{
			"Perl": regexp.MustCompile, "POSIX": regexp.MustCompilePOSIX,
		} {
			for name, constructor := range map[string]func(*regexp.Regexp) govy.Rule[string]{
				"match": rules.StringMatchRegexp, "deny": rules.StringDenyRegexp,
			} {
				t.Run(mode+"/"+name+"/"+pattern, func(t *testing.T) {
					t.Parallel()
					re := compile(pattern)
					rule := constructor(re)
					for _, input := range []string{"foo", "foo\n", "\n", "a", "b"} {
						matches := re.MatchString(input)
						if name == "deny" {
							matches = !matches
						}
						assert.Equal(t, matches, rule.Validate(input) == nil)
					}
					_, err := govy.JSONSchema(govy.New(govy.For(govy.GetSelf[string]()).Rules(rule)))
					assert.ErrorContains(
						t,
						err,
						"has different Perl and POSIX meanings; provide an explicit WithJSONSchema mapping",
					)
				})
			}
		}
	}
}

func TestJSONSchema_RegexpExplicitMode(t *testing.T) {
	t.Parallel()
	v := govy.New(govy.For(govy.GetSelf[string]()).Rules(rules.StringMatchRegexp(regexp.MustCompile(`(?-m)^foo$`))))
	schema, err := govy.JSONSchema(v)
	assert.Require(t, assert.NoError(t, err))
	jsonschematest.Assert(t, schema, "test_data/expected_explicit_regexp_mode.json", []jsonschematest.Case[string]{
		{Name: "whole input", Input: "foo", Valid: true},
		{Name: "final newline", Input: "foo\n"},
		{Name: "interior line", Input: "x\nfoo\ny"},
	})
}

func TestJSONSchema_RegexpExplicitMapping(t *testing.T) {
	t.Parallel()
	rule := rules.StringMatchRegexp(regexp.MustCompilePOSIX(`[^a]`)).
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{Pattern: `[^a\n]`}, nil
		})
	v := govy.New(govy.For(govy.GetSelf[string]()).Rules(rule))
	schema, err := govy.JSONSchema(v)
	assert.Require(t, assert.NoError(t, err))
	cases := []jsonschematest.Case[string]{
		{Name: "matching character", Input: "b", Valid: true},
		{Name: "excluded character", Input: "a"},
		{Name: "newline", Input: "\n"},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.Valid, v.Validate(tc.Input) == nil)
	}
	jsonschematest.Assert(t, schema, "test_data/expected_explicit_regexp_mapping.json", cases)
}
