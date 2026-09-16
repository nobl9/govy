# Testing

When callers validate whole objects, test through that public entrypoint.
Reuse the consumer's test helpers before introducing direct govytest assertions.
Start with a valid object and change one property per case.
Include a case with independent failures to check error aggregation.

Assert complete property paths, stable error codes, and the expected error count.
For assertion APIs and examples, read [govytest](govytest.md).
Use `govytest.AssertError` for an exact set of failures.
Use `AssertErrorContains` when the test contract allows additional failures.
Assert message text when the message itself is under test.

Test pointer fields with `nil`, a pointer to zero, and a pointer to a valid value.
For ranges and collection limits, test both boundaries and values just outside them.
When callers consume plans, inspect their paths, descriptions, and conditions.
Also test runtime errors.

The examples below inspect structured errors without the govytest package.

## Topics

- [Inspect validation errors](#inspect-validation-errors)
  - [Check whether any nested rule error carries an error code.](#check-whether-any-nested-rule-error-carries-an-error-code)
  - [Match validator errors returned from slice validation.](#match-validator-errors-returned-from-slice-validation)

## Inspect validation errors

Use these helpers to inspect structured errors without matching full messages.

### Check whether any nested rule error carries an error code

[//]: # (embed: ExampleHasErrorCode?comments=false)

```go
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
}
```

### Match validator errors returned from slice validation

[//]: # (embed: ExampleValidatorErrors?comments=false)

```go
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
}
```
