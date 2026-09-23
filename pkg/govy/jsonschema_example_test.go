package govy_test

import (
	"encoding/json"
	"fmt"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/jsonschema"
	"github.com/nobl9/govy/pkg/rules"
)

func ExampleJSONSchema() {
	type Service struct {
		Name     string `json:"name"`
		Replicas int    `json:"replicas"`
	}
	v := govy.New(
		govy.For(func(s Service) string { return s.Name }).
			WithName("name").
			Required().
			Rules(rules.StringMinLength(1)),
		govy.For(func(s Service) int { return s.Replicas }).
			WithName("replicas").
			Rules(rules.GTE(1)),
	).WithName("Service")

	schema, err := govy.JSONSchema(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	// Output:
	// {
	//   "$schema": "https://json-schema.org/draft/2020-12/schema",
	//   "title": "Service",
	//   "type": "object",
	//   "properties": {
	//     "name": {
	//       "type": "string",
	//       "minLength": 1
	//     },
	//     "replicas": {
	//       "type": "integer",
	//       "minimum": 1
	//     }
	//   },
	//   "required": [
	//     "name"
	//   ]
	// }
}

func ExampleRule_WithJSONSchema() {
	even := govy.NewRule(func(value int) error {
		if value%2 != 0 {
			return fmt.Errorf("%d is not even", value)
		}
		return nil
	}).
		WithErrorCode("even").
		WithDescription("value must be even").
		WithJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
			return &jsonschema.Schema{MultipleOf: "2"}, nil
		})
	v := govy.New(
		govy.For(govy.GetSelf[int]()).
			Rules(
				rules.GTE(2),
				even,
			),
	)

	schema, err := govy.JSONSchema(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	// Output:
	// {
	//   "$schema": "https://json-schema.org/draft/2020-12/schema",
	//   "type": "integer",
	//   "multipleOf": 2,
	//   "minimum": 2
	// }
}

func ExampleWhenJSONSchema() {
	type Endpoint struct {
		Mode    string `json:"mode"`
		Address string `json:"address,omitempty"`
	}
	remote := any("remote")
	v := govy.New(
		govy.For(func(e Endpoint) string { return e.Mode }).
			WithName("mode").
			Required().
			Rules(rules.OneOf("local", "remote")),
		govy.For(func(e Endpoint) string { return e.Address }).
			WithName("address").
			Required().
			When(
				func(e Endpoint) bool { return e.Mode == "remote" },
				govy.WhenDescription("mode is remote"),
				govy.WhenJSONSchema(func(govy.JSONSchemaBuilderContext) (*jsonschema.Schema, error) {
					return &jsonschema.Schema{
						Properties: map[string]*jsonschema.Schema{
							"mode": {Const: &remote},
						},
						Required: []string{"mode"},
					}, nil
				}),
			),
	)

	schema, err := govy.JSONSchema(v)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	// Output:
	// {
	//   "$schema": "https://json-schema.org/draft/2020-12/schema",
	//   "type": "object",
	//   "allOf": [
	//     {
	//       "if": {
	//         "properties": {
	//           "mode": {
	//             "const": "remote"
	//           }
	//         },
	//         "required": [
	//           "mode"
	//         ]
	//       },
	//       "then": {
	//         "properties": {
	//           "address": {}
	//         },
	//         "required": [
	//           "address"
	//         ]
	//       }
	//     }
	//   ],
	//   "properties": {
	//     "address": {
	//       "type": "string"
	//     },
	//     "mode": {
	//       "type": "string",
	//       "enum": [
	//         "local",
	//         "remote"
	//       ]
	//     }
	//   },
	//   "required": [
	//     "mode"
	//   ]
	// }
}

func ExampleJSONSchemaIncludeOmittedRules() {
	type Settings struct {
		TimeZone string `json:"timeZone"`
	}
	v := govy.New(
		govy.For(func(s Settings) string { return s.TimeZone }).
			WithName("timeZone").
			Rules(rules.StringTimeZone()),
	)

	schema, err := govy.JSONSchema(v, govy.JSONSchemaIncludeOmittedRules())
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	// Output:
	// {
	//   "$schema": "https://json-schema.org/draft/2020-12/schema",
	//   "type": "object",
	//   "properties": {
	//     "timeZone": {
	//       "type": "string"
	//     }
	//   },
	//   "x-govy-omittedRules": [
	//     {
	//       "path": "$.timeZone",
	//       "rule": "string_time_zone",
	//       "reason": "missing JSON Schema builder"
	//     }
	//   ]
	// }
}
