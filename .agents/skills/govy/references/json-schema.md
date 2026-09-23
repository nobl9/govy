# JSON Schema generation

Check the Govy README for feature status and compatibility guidance.
Check that the consumer's pinned Govy version provides the APIs below.
Keep Govy validation as the authoritative check for Go values.

## Generate a document

Call `govy.JSONSchema(validator)` and marshal the returned `*jsonschema.Document`.
The document declares JSON Schema Draft 2020-12.
Generation needs no Node.js or Ajv dependency.

Use the consumer's JSON property paths through `WithName`, `WithPath`,
or existing path inference.
Generation uses the validation plan, not struct-tag discovery.
It does not include unvalidated fields or guarantee types for intermediate containers.
Unsupported Go kinds and builder errors stop generation.

[//]: # (embed: pkg/govy/jsonschema_example_test.go#ExampleJSONSchema?comments=false)

```go
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
}
```

## Map a custom rule

Use `Rule.WithJSONSchema` to return an independent `*jsonschema.Schema`
containing the rule's constraints.
Descriptions, details, and error codes are optional for schema builders.
Generation calls the builder and propagates its errors even without this metadata.
The builder context exposes the selected value's absolute path and JSON type.
It does not expose a mutable baseline schema.
Return only the constraints owned by the rule.
Govy merges non-conflicting keywords and uses `allOf` when keywords overlap.
Custom `additionalProperties` applies only to unnamed properties.
Custom `items` applies only after `prefixItems`.
Govy normalizes its own wildcard paths without broadening these custom constraints.
A builder that returns `nil, nil` contributes no constraint.

[//]: # (embed: pkg/govy/jsonschema_example_test.go#ExampleRule_WithJSONSchema?comments=false)

```go
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
}
```

## Map a condition

Pass `govy.WhenJSONSchema(builder)` as an option to `When`.
The builder returns the condition, not an `if`/`then` wrapper.
Govy applies it at the same scope as the Go predicate.
A property-level predicate receives the parent object
and can check sibling properties.
A rule-level predicate receives the selected value.
Use `required` inside a condition when an absent property must not match.
Govy supplies `if`/`then` and combines multiple conditions with AND semantics.
Repeated conditions receive the same known type for their scope.
Temporary conditional branches do not change the builder context type.

Govy cannot infer a condition from a Go function or `WhenDescription`.
Without a schema builder for every guarding predicate,
Govy omits the guarded rule's constraints.
A condition builder that returns `nil, nil` also omits those constraints.

[//]: # (embed: pkg/govy/jsonschema_example_test.go#ExampleWhenJSONSchema?comments=false)

```go
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
}
```

## Inspect omissions

Pass `govy.JSONSchemaIncludeOmittedRules()`
to include the root annotation `x-govy-omittedRules`.
Each record has a `path`, an optional rule error code in `rule`, and a `reason`.
The annotation reports missing rule or condition builders
and rules on transformed values.
It does not affect validation or classify approximate mappings.
An executed builder that returns `nil, nil` produces no omission record.
The annotation is absent when no omissions are recorded.
Consumers with strict unknown-keyword checks must register the annotation.
For Ajv, use `ajv.addKeyword("x-govy-omittedRules")` before compiling.

[//]: # (embed: pkg/govy/jsonschema_example_test.go#ExampleJSONSchemaIncludeOmittedRules?comments=false)

```go
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
}
```

## Explain the limits

Do not promise equivalent Govy and JSON Schema validation:

- `Required` checks property presence in the schema, not Go zero values.
  Generation recognizes the `required` component in the rule's error-code chain.
  Replacing that component removes this constraint.
  RuleSet error-code prefixes preserve it.
- Optional pointers can be absent, but their schemas do not accept explicit `null`.
  Decoding either form into a zero-valued Go struct produces a nil pointer.
- `Transform` preserves the input type and directly attached required constraints.
  It omits constraints and included structure for the transformed value.
- Some rules have only approximate mappings, annotations, or no mapping.
  No omission record does not prove exact validation.
- Format assertions depend on the JSON Schema consumer's configuration.

When working in the Govy repository, inspect `JSONSchemaDifference` cases
in `pkg/rules` and `pkg/govy` for documented validation differences.
Test generated schemas through `govy.JSONSchema`, not private helpers.
Keep expected schemas in JSON fixtures and validate accepted and rejected inputs
with a real schema validator.
