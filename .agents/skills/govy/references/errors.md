# Errors

## Topics

- [Work with validator errors](#work-with-validator-errors)
  - [Set or overwrite a validator name on returned errors.](#set-or-overwrite-a-validator-name-on-returned-errors)
  - [Inspect and serialize validator error output.](#inspect-and-serialize-validator-error-output)
- [Hide property values](#hide-property-values)
- [Create property-scoped errors](#create-property-scoped-errors)
  - [Construct a property error with one or more rule errors.](#construct-a-property-error-with-one-or-more-rule-errors)

## Work with validator errors

Use validator errors when validation failed at the validator level.
Attach names after validation.
Inspect or serialize the structured fields.

### Set or overwrite a validator name on returned errors

[//]: # (embed: ExampleValidatorError_WithName?comments=false)

```go
func ExampleValidatorError_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Jake"))
	}
}
```

### Inspect and serialize validator error output

[//]: # (embed: ExampleValidatorError?comments=false)

```go
func ExampleValidatorError() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })).
			WithName("name"),
	).WithName("Teacher")

	err := v.Validate(Teacher{Name: "John"})
	if err != nil {
		if validatorErr, ok := err.(*govy.ValidatorError); ok {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			if err = enc.Encode(validatorErr); err != nil {
				fmt.Printf("error encoding: %v\n", err)
			}
		}
	}
}
```

## Hide property values

Configure `PropertyRules.HideValue()` before validation.
It applies to nested rules, included validators, and collection checks.
It clears `PropertyError.PropertyValue`
and redacts value text from ordinary rule errors.
Message templates receive `[hidden]` as `.PropertyValue` and a redacted `.Error`.

Template rendering does not redact details, examples, custom fields,
comparison values, or literal text.
Keep secrets out of those inputs.
Property paths stay visible, including map keys.

For migrations, replace error-level `HideValue` calls with property configuration.
`PropertyErrors`, `PropertyError`, and `RuleError` no longer expose those methods.

## Create property-scoped errors

Return property errors from custom rules to identify failures at nested paths.

### Construct a property error with one or more rule errors

[//]: # (embed: ExampleNewPropertyError?comments=false)

```go
func ExampleNewPropertyError() {
	v := govy.New(
		govy.For(govy.GetSelf[Teacher]()).
			Rules(govy.NewRule(func(t Teacher) error {
				if t.Name == "Jake" {
					return govy.NewPropertyError(
						jsonpath.Parse("name"),
						t.Name,
						govy.NewRuleError("name cannot be Jake", "error_code_jake"),
						govy.NewRuleError("you can pass me too!"),
					)
				}
				return nil
			})),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		propertyErrors := err.(*govy.ValidatorError).Errors
		ruleErrors := propertyErrors[0].Errors
		fmt.Printf("Error code: %s\n\n", ruleErrors[0].Code)
		fmt.Println(err)
	}
}
```
