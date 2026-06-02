# Testing

Core govy error helpers and structured error flows useful when writing tests without the govytest package.

## Topics

- [Inspect validation errors](#inspect-validation-errors)
  - [Check whether any nested rule error carries an error code.](#check-whether-any-nested-rule-error-carries-an-error-code)
  - [Match validator errors returned from slice validation.](#match-validator-errors-returned-from-slice-validation)

## Inspect validation errors

Use these helpers when tests need to match structured govy errors without depending on entire formatted error strings.

<a id="check-whether-any-nested-rule-error-carries-an-error-code"></a>

**Check whether any nested rule error carries an error code.**

[//]: # (embed: ExampleHasErrorCode)

```go
// To inspect if an error contains a given [govy.ErrorCode], use [govy.HasErrorCode] function.
// This function will also return true if the expected [govy.ErrorCode]
// is part of a chain of wrapped error codes.
// In this example we're dealing with two error code chains:
//   - 'teacher_name:string_length'
//   - 'teacher_name:string_match_regexp'
func ExampleHasErrorCode() {
	teacherNameRule := govy.NewRuleSet(
		rules.StringLength(1, 5),
		rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")),
	).
		WithErrorCode("teacher_name")

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(teacherNameRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jonathan",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		for _, code := range []govy.ErrorCode{
			"teacher_name",
			"string_length",
			"string_match_regexp",
		} {
			if govy.HasErrorCode(err, code) {
				fmt.Println("Has error code:", code)
			}
		}
	}

	// Output:
	// Has error code: teacher_name
	// Has error code: string_length
	// Has error code: string_match_regexp
}
```

<a id="match-validator-errors-returned-from-slice-validation"></a>

**Match validator errors returned from slice validation.**

[//]: # (embed: ExampleValidatorErrors)

```go
// [govy.Validator.ValidateSlice] outputs [govy.ValidatorErrors] which is a slice of [govy.ValidatorError].
// Each [govy.ValidatorError] has an additional property set: SliceIndex, which is a 0-based slice element index.
func ExampleValidatorErrors() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(govy.NewRule(func(name string) error {
				if name == "John" || name == "Jake" {
					return fmt.Errorf("fails for John and Jake")
				}
				return nil
			})),
	).WithName("Teacher")

	err := v.ValidateSlice([]Teacher{
		{Name: "John"},
		{Name: "George"},
		{Name: "Jake"},
	})
	if err != nil {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err = enc.Encode(err); err != nil {
			fmt.Printf("error encoding: %v\n", err)
		}
	}

	// Output:
	// [
	//   {
	//     "errors": [
	//       {
	//         "propertyPath": "name",
	//         "propertyValue": "John",
	//         "errors": [
	//           {
	//             "error": "fails for John and Jake"
	//           }
	//         ]
	//       }
	//     ],
	//     "name": "Teacher",
	//     "sliceIndex": 0
	//   },
	//   {
	//     "errors": [
	//       {
	//         "propertyPath": "name",
	//         "propertyValue": "Jake",
	//         "errors": [
	//           {
	//             "error": "fails for John and Jake"
	//           }
	//         ]
	//       }
	//     ],
	//     "name": "Teacher",
	//     "sliceIndex": 2
	//   }
	// ]
}
```
