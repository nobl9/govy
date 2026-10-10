package jsonschema_test

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/pkg/jsonschema"
)

func TestDocument_MarshalJSON(t *testing.T) {
	t.Parallel()

	zero := uint64(0)
	null := any(nil)
	emptyString := any("")
	document := jsonschema.Document{
		Title: "Example",
		Type:  jsonschema.TypeObject,
		Properties: map[string]*jsonschema.Schema{
			"value": {
				Type:      jsonschema.TypeString,
				MinLength: &zero,
			},
			"nothing": {Const: &null},
			"pair": {
				Type: jsonschema.TypeArray,
				PrefixItems: []*jsonschema.Schema{
					{Type: jsonschema.TypeString},
					{Type: jsonschema.TypeInteger},
				},
			},
		},
		Required: []string{"value"},
		AllOf: []*jsonschema.Schema{
			{Not: &jsonschema.Schema{Const: &emptyString}},
		},
	}

	expected, err := os.ReadFile("testdata/expected_document.json")
	assert.Require(t, assert.NoError(t, err))

	var actual bytes.Buffer
	encoder := json.NewEncoder(&actual)
	encoder.SetIndent("", "  ")
	assert.Require(t, assert.NoError(t, encoder.Encode(document)))
	assert.Equal(t, string(expected), actual.String())
}
