# Validation Plan

Validation plan generation and strict validation of plan metadata.

## Examples

- [Generate a validation plan](#generate-a-validation-plan)
- [Validate generated plan metadata](#validate-generated-plan-metadata)

## Generate a validation plan

[//]: # (embed: ExamplePlan)

```go
// When documenting an API it's often a struggle to keep consistency
// between the code and documentation we write for it.
// What If your code could be self-descriptive?
// Specifically, what If we could generate documentation out of our validation rules?
// We can achieve that by using [govy.Plan] function!
//
// There are multiple ways to improve the generated documentation:
//   - Use [govy.PropertyRules.WithExamples] to provide a list of example values for the property.
//   - Use [govy.Rule.WithDescription] to provide a plan-only description for your rule.
//     For builtin rules, the description is already provided.
//   - Use [govy.WhenDescription] to provide a plan-only description for your when conditions.
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

	// Output:
	// {
	//   "name": "Teacher",
	//   "properties": [
	//     {
	//       "path": "$.name",
	//       "typeInfo": {
	//         "name": "string",
	//         "kind": "string"
	//       },
	//       "examples": [
	//         "Jake",
	//         "John"
	//       ],
	//       "rules": [
	//         {
	//           "description": "must not be equal to 'Jerry'",
	//           "details": "Jerry is just a name!",
	//           "errorCode": "not_equal_to",
	//           "conditions": [
	//             "name is Jerry"
	//           ]
	//         },
	//         {
	//           "description": "this is a custom error!",
	//           "conditions": [
	//             "name is Jerry"
	//           ]
	//         }
	//       ]
	//     }
	//   ]
	// }
}
```

## Validate generated plan metadata

[//]: # (embed: ExamplePlan_validation)

```go
// You can enforce certain rules upon [govy.Plan].
// For instance, If you'd want to make sure every [govy.Predicate]
// has a description attached to it, provide [govy.Plan] with [govy.PlanRequirePredicateDescription] option.
//
// If you want to follow our best recommendations, use [govy.PlanStrictMode].
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

	// Output:
	// predicates without description found at: validator level, $.name
}
```

