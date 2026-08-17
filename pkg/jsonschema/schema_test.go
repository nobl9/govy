package jsonschema_test

import (
	"encoding/json"
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
		},
		Required: []string{"value"},
		AllOf: []*jsonschema.Schema{
			{Not: &jsonschema.Schema{Const: &emptyString}},
		},
	}

	actualJSON, err := json.Marshal(document)
	assert.Require(t, assert.NoError(t, err))

	expectedJSON := []byte(`{
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"title": "Example",
		"type": "object",
		"properties": {
			"value": {
				"type": "string",
				"minLength": 0
			},
			"nothing": {"const": null}
		},
		"required": ["value"],
		"allOf": [{"not": {"const": ""}}]
	}`)

	var expected, actual any
	assert.Require(t, assert.NoError(t, json.Unmarshal(expectedJSON, &expected)))
	assert.Require(t, assert.NoError(t, json.Unmarshal(actualJSON, &actual)))
	assert.Equal(t, expected, actual)
}
