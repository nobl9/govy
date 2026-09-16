# Validation Plan

Validation plan generation and strict validation of plan metadata.

## Topics

- [Generate and enforce validation plans](#generate-and-enforce-validation-plans)
  - [Generate a plan from a validator.](#generate-a-plan-from-a-validator)
  - [Validate plan metadata with strict plan options.](#validate-plan-metadata-with-strict-plan-options)

## Generate and enforce validation plans

Use validation plans when validation rules also need to describe an API contract.
Add examples, descriptions, and predicate descriptions before generating the plan.

`ValidatorPlan.TypeInfo` describes the root Go type with `Name`, `Kind`, and `Package`.
It is separate from the display name set by `WithName`.
Serialized plans include this metadata as the root `typeInfo` field.

### Generate a plan from a validator

[//]: # (embed: ExamplePlan?comments=false)

```go
func ExamplePlan() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			WithExamples("Jake", "John").
			When(
				func(t Teacher) bool { return t.Name == "Jerry" },
				govy.WhenDescription("name is Jerry"),
			).
			Rules(
				rules.NEQ("Jerry").
					WithDetails("Jerry is just a name!"),
				govy.NewRule(func(v string) error {
					return fmt.Errorf("some custom error")
				}).
					WithDescription("this is a custom error!"),
			),
	).WithName("Teacher")

	properties, err := govy.Plan(v)
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(properties)
}
```

### Validate plan metadata with strict plan options

[//]: # (embed: ExamplePlan_validation?comments=false)

```go
func ExamplePlan_validation() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			WithExamples("Jake", "John").
			When(func(t Teacher) bool { return t.Name == "Jerry" }).
			Rules(
				rules.NEQ("Jerry").
					WithDetails("Jerry is just a name!"),
				govy.NewRule(func(v string) error {
					return fmt.Errorf("some custom error")
				}).
					WithDescription("this is a custom error!"),
			),
	).
		When(func(t Teacher) bool { return t.Age > 18 }).
		WithName("Teacher")

	_, err := govy.Plan(v, govy.PlanStrictMode())
	fmt.Println(err)
}
```
