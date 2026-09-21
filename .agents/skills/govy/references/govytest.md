# Govytest

Assertion helpers for validating govy errors in tests.

## Topics

- [Assert validation results](#assert-validation-results)
  - [Assert that validation succeeded.](#assert-that-validation-succeeded)
  - [Assert exact structured rule errors.](#assert-exact-structured-rule-errors)
  - [Assert nested ValidatorErrors produced by slice validation.](#assert-nested-validatorerrors-produced-by-slice-validation)
  - [Assert that one expected error is present among other errors.](#assert-that-one-expected-error-is-present-among-other-errors)

## Assert validation results

Use govytest to match structured validation failures.
The examples below intentionally fail assertions to show their diagnostics.
They use test-only types and a mock testing object from govy's example suite.
In application tests, pass `*testing.T` and expect the actual rule codes.
For example, an empty required address has code `required`, not `greater_than`.

### Assert that validation succeeded

[//]: # (embed: ExampleAssertNoError?comments=false)

```go
func ExampleAssertNoError() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Jake", "George"),
			),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(govy.New(
				govy.For(func(u University) string { return u.Address }).
					WithName("address").
					Required(),
			)),
	)

	teacher := Teacher{
		Name: "John",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	mt := new(mockTestingT)

	err := teacherValidator.WithName("John").Validate(teacher)
	govytest.AssertNoError(mt, err)

	fmt.Println(mt.recordedError)
}
```

### Assert exact structured rule errors

[//]: # (embed: ExampleAssertError?comments=false)

```go
func ExampleAssertError() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Jake", "George"),
			),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(govy.New(
				govy.For(func(u University) string { return u.Address }).
					WithName("address").
					Required(),
			)),
	)

	teacher := Teacher{
		Name: "John",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	mt := new(mockTestingT)

	err := teacherValidator.WithName("John").Validate(teacher)
	govytest.AssertError(mt, err,
		govytest.ExpectedRuleError{
			PropertyPath:    "name",
			ContainsMessage: "one of",
		},
		govytest.ExpectedRuleError{
			PropertyPath: "university.address",
			Code:         "greater_than",
		},
	)

	fmt.Println(mt.recordedError)
}
```

### Assert nested ValidatorErrors produced by slice validation

[//]: # (embed: ExampleAssertError_validatorErrors?comments=false)

```go
func ExampleAssertError_validatorErrors() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Eve", "George"),
			),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(govy.New(
				govy.For(func(u University) string { return u.Address }).
					WithName("address").
					Required(),
			)),
	)

	teacherEve := Teacher{
		Name: "Eve",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}
	teacherJohn := Teacher{
		Name: "John",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "Some address",
		},
	}
	teachers := []Teacher{teacherEve, teacherJohn}

	mt := new(mockTestingT)

	err := teacherValidator.WithNameFunc(func(s Teacher) string { return s.Name }).ValidateSlice(teachers)
	govytest.AssertError(mt, err,
		govytest.ExpectedRuleError{
			PropertyPath:   "university.address",
			Code:           "greater_than",
			ValidatorName:  "Eve",
			ValidatorIndex: ptr(0),
		},
		govytest.ExpectedRuleError{
			PropertyPath:    "name",
			ContainsMessage: "one of",
			ValidatorName:   "John",
			ValidatorIndex:  ptr(1),
		},
	)

	fmt.Println(mt.recordedError)
}
```

### Assert that one expected error is present among other errors

[//]: # (embed: ExampleAssertErrorContains?comments=false)

```go
func ExampleAssertErrorContains() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Jake", "George"),
			),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(govy.New(
				govy.For(func(u University) string { return u.Address }).
					WithName("address").
					Required(),
			)),
	)

	teacher := Teacher{
		Name: "John",
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	mt := new(mockTestingT)

	err := teacherValidator.WithName("John").Validate(teacher)
	govytest.AssertErrorContains(mt, err, govytest.ExpectedRuleError{
		PropertyPath: "name",
		Code:         "one_of",
	})

	err = teacherValidator.WithName("John").Validate(teacher)
	govytest.AssertErrorContains(mt, err, govytest.ExpectedRuleError{
		PropertyPath: "university.address",
		Code:         "greater_than",
	})

	fmt.Println(mt.recordedError)
}
```
