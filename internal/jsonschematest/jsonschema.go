// Package jsonschematest provides fixture and Ajv assertions for generated schemas.
package jsonschematest

import (
	"bytes"
	"embed"
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/nobl9/govy/internal/assert"
	"github.com/nobl9/govy/pkg/jsonschema"
)

// Embed the runner and lockfile so changes to either invalidate Go's test cache.
//
//go:embed testdata/validate.cjs testdata/package-lock.json
var jsonSchemaRunnerFiles embed.FS

// Case holds an input and its expected Govy validation result.
type Case[T any] struct {
	Name  string
	Input T
	Valid bool
	// A non-empty reason means JSON Schema must return the opposite of Valid.
	JSONSchemaDifference string
}

// Assert compares schema with the JSON fixture and validates cases with Ajv.
// The fixture path is relative to the calling test's working directory.
func Assert[T any](t *testing.T, schema *jsonschema.Document, fixture string, cases []Case[T]) {
	t.Helper()

	expected, err := os.ReadFile(fixture)
	assert.Require(t, assert.NoError(t, err))
	actual, err := json.MarshalIndent(schema, "", "  ")
	assert.Require(t, assert.NoError(t, err))
	assert.Equal(t, string(expected), string(actual)+"\n")

	inputs := make([]T, len(cases))
	for i, tc := range cases {
		inputs[i] = tc.Input
	}
	request, err := json.Marshal(struct {
		Schema *jsonschema.Document `json:"schema"`
		Inputs []T                  `json:"inputs"`
	}{Schema: schema, Inputs: inputs})
	assert.Require(t, assert.NoError(t, err))

	runner, err := jsonSchemaRunnerFiles.ReadFile("testdata/validate.cjs")
	assert.Require(t, assert.NoError(t, err))
	// Resolve Ajv relative to this package, not the package under test.
	_, source, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate JSON Schema test dependencies: source path unavailable")
	}
	//nolint:gosec // Only embedded test code is executed. Test inputs go through stdin.
	cmd := exec.CommandContext(t.Context(), "node", "-e", string(runner))
	cmd.Dir = filepath.Join(filepath.Dir(source), "testdata")
	cmd.Stdin = bytes.NewReader(request)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("run Ajv (use devbox run -- make test): %v\n%s", err, stderr.String())
	}
	var results []struct {
		Valid  bool            `json:"valid"`
		Errors json.RawMessage `json:"errors"`
	}
	assert.Require(t, assert.NoError(t, json.Unmarshal(output, &results)))
	assert.Require(t, assert.Len(t, results, len(cases)))
	for i, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			wantValid := tc.Valid
			if tc.JSONSchemaDifference != "" {
				wantValid = !wantValid
				t.Logf("Difference from Govy: %s", tc.JSONSchemaDifference)
			}
			if !assert.Equal(t, wantValid, results[i].Valid) {
				t.Logf("Input: %v\nAjv errors: %s", tc.Input, results[i].Errors)
			}
		})
	}
}
