# Errors

Validator errors, runtime error naming, structured JSON errors, and manually constructed property errors.

## Topics

- [Work with validator errors](#work-with-validator-errors)
  - [Set or overwrite a validator name on returned errors.](#set-or-overwrite-a-validator-name-on-returned-errors)
  - [Inspect and serialize validator error output.](#inspect-and-serialize-validator-error-output)
- [Create property-scoped errors](#create-property-scoped-errors)
  - [Construct a property error with one or more rule errors.](#construct-a-property-error-with-one-or-more-rule-errors)

## Work with validator errors

Use validator errors when validation failed at the validator level. Names can be attached after validation, and structured fields are available for serialization or inspection.

<a id="set-or-overwrite-a-validator-name-on-returned-errors"></a>

**Set or overwrite a validator name on returned errors.**

[//]: # (embed: ExampleValidatorError_WithName)

```go
// You can also add [govy.Validator] name during runtime,
// by calling [govy.ValidatorError.WithName] function on the returned error.
//
// Note: We left the previous "Teacher" name assignment, to demonstrate that
// the [govy.ValidatorError.WithName] function call will overwrite it.
//
// Note: This would also work:
//
//	err := v.WithName("Jake").Validate(Teacher{})
//
// govy, excluding error handling, tries to follow immutability principle.
// Calling any method on [govy.Validator] will not change its declared instance,
// but rather create a copy of it.
func ExampleValidatorError_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Jake"))
	}

	// Output:
	// Validation for Jake has failed:
	//   - always fails
}
```

<a id="inspect-and-serialize-validator-error-output"></a>

**Inspect and serialize validator error output.**

[//]: # (embed: ExampleValidatorError)

```go
// All errors returned by [govy.Validator] are of type [govy.ValidatorError].
// Type casting directly to [govy.ValidatorError] should be safe once an error
// was asserted to be non-nil.
// However, you shouldn't trust any API with such promises, and always type check in your
// type assignments.
//
// All error types return by govy are JSON serializable.
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

	// Output:
	// {
	//   "errors": [
	//     {
	//       "propertyPath": "name",
	//       "propertyValue": "John",
	//       "errors": [
	//         {
	//           "error": "always fails"
	//         }
	//       ]
	//     }
	//   ],
	//   "name": "Teacher"
	// }
}
```

## Create property-scoped errors

Return property errors from custom rules when top-level validation logic needs to point the failure at a nested property path.

<a id="construct-a-property-error-with-one-or-more-rule-errors"></a>

**Construct a property error with one or more rule errors.**

[//]: # (embed: ExampleNewPropertyError)

```go
// Sometimes you need top level context,
// but you want to scope the error to a specific, nested property.
// One of the ways to do that is to use [govy.NewPropertyError]
// and return [govy.PropertyError] from your validation rule.
// Note that you can still use [govy.ErrorCode] and pass [govy.RuleError] to the constructor.
// You can pass any number of [govy.RuleError].
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

	// Output:
	// Error code: error_code_jake
	//
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - name cannot be Jake
	//     - you can pass me too!
}
```
