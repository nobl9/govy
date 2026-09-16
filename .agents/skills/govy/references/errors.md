# Errors

## Topics

- [Work with validator errors](#work-with-validator-errors)
  - [Set or overwrite a validator name on returned errors.](#set-or-overwrite-a-validator-name-on-returned-errors)
  - [Inspect and serialize validator error output.](#inspect-and-serialize-validator-error-output)
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
						govy.NewRuleError("you can pass me too!"))
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
