# Core Validation

## Topics

- [Construct reusable validators](#construct-reusable-validators)
  - [Create a validator from property rules.](#create-a-validator-from-property-rules)
  - [Attach a static validator name.](#attach-a-static-validator-name)
  - [Compute the validator name from the validated value.](#compute-the-validator-name-from-the-validated-value)
- [Pass validation callbacks](#pass-validation-callbacks)
- [Control when validators run](#control-when-validators-run)
  - [Run a validator only when a predicate matches.](#run-a-validator-only-when-a-predicate-matches)
  - [Branch validation by including different validators under property conditions.](#branch-validation-by-including-different-validators-under-property-conditions)
- [Validate slices](#validate-slices)
  - [Validate a slice while preserving indexed property paths.](#validate-a-slice-while-preserving-indexed-property-paths)
  - [Validate each element directly with ValidateSlice.](#validate-each-element-directly-with-validateslice)
- [Control error aggregation](#control-error-aggregation)
  - [Stop evaluating later properties after a validator failure.](#stop-evaluating-later-properties-after-a-validator-failure)
- [Compose a complete validator](#compose-a-complete-validator)
  - [Build a full Teacher validator.](#build-a-full-teacher-validator)

## Construct reusable validators

Define validators once and reuse them.
Name validators when the resulting error should identify the validated entity.

### Create a validator from property rules

[//]: # (embed: ExampleNew?comments=false)

```go
func ExampleNew() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	)

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err)
	}
}
```

### Attach a static validator name

[//]: # (embed: ExampleValidator_WithName?comments=false)

```go
func ExampleValidator_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.Validate(Teacher{})
	if err != nil {
		fmt.Println(err)
	}
}
```

### Compute the validator name from the validated value

[//]: # (embed: ExampleValidator_WithNameFunc?comments=false)

```go
func ExampleValidator_WithNameFunc() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithNameFunc(func(t Teacher) string { return "Teacher " + t.Name })

	err := v.Validate(Teacher{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}
}
```

## Pass validation callbacks

The `Validate` methods take `...govy.ValidationOption`.
Direct `Validate(value)` calls still work.
To pass validation as `func(T) error`, wrap the call in a function.
Custom validation interfaces must include the variadic option parameter.
Use builders to configure validation.
There are no exported validation option constructors.

## Control when validators run

Use validator-level conditions for whole-object gates.
Use property-level conditions when only one property branch should be skipped.

### Run a validator only when a predicate matches

[//]: # (embed: ExampleValidator_When?comments=false)

```go
func ExampleValidator_When() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).
		When(func(t Teacher) bool { return t.Age < (50 * year) })

	teacherTom := Teacher{
		Name: "Tom",
		Age:  51 * year,
	}
	teacherJerry := Teacher{
		Name: "Jerry",
		Age:  30 * year,
	}

	err := v.Validate(teacherTom)
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Tom"))
	}
	err = v.Validate(teacherJerry)
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Jerry"))
	}
}
```

### Branch validation by including different validators under property conditions

[//]: # (embed: ExampleValidator_branchingPattern?comments=false)

```go
func ExampleValidator_branchingPattern() {
	type (
		CSV struct {
			Separator string `json:"separator"`
		}
		JSON struct {
			Indent string `json:"indent"`
		}
		File struct {
			Format string `json:"format"`
			CSV    *CSV   `json:"csv,omitempty"`
			JSON   *JSON  `json:"json,omitempty"`
		}
	)

	csvValidation := govy.New(
		govy.For(func(c CSV) string { return c.Separator }).
			WithName("separator").
			Required().
			Rules(rules.OneOf(",", ";")),
	)

	jsonValidation := govy.New(
		govy.For(func(j JSON) string { return j.Indent }).
			WithName("indent").
			Required().
			Rules(rules.StringMatchRegexp(regexp.MustCompile(`^\s*$`))),
	)

	fileValidation := govy.New(
		govy.ForPointer(func(f File) *CSV { return f.CSV }).
			When(func(f File) bool { return f.Format == "csv" }).
			Include(csvValidation),
		govy.ForPointer(func(f File) *JSON { return f.JSON }).
			When(func(f File) bool { return f.Format == "json" }).
			Include(jsonValidation),
		govy.For(func(f File) string { return f.Format }).
			WithName("format").
			Required().
			Rules(rules.OneOf("csv", "json")),
	).WithName("File")

	file := File{
		Format: "json",
		CSV:    nil,
		JSON: &JSON{
			Indent: "invalid",
		},
	}

	err := fileValidation.Validate(file)
	if err != nil {
		fmt.Println(err)
	}
}
```

## Validate slices

Use ForSlice when the slice itself has rules or the path should include indexes.
Use ValidateSlice when each value can be validated independently.

### Validate a slice while preserving indexed property paths

[//]: # (embed: ExampleValidator_Validate_slice?comments=false)

```go
func ExampleValidator_Validate_slice() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	)
	v := govy.New(
		govy.ForSlice(govy.GetSelf[[]Teacher]()).
			IncludeForEach(teacherValidator),
	)

	err := v.Validate([]Teacher{
		{Name: "John"},
		{Name: "Jake"},
	})
	if err != nil {
		fmt.Println(err)
	}
}
```

### Validate each element directly with ValidateSlice

[//]: # (embed: ExampleValidator_ValidateSlice?comments=false)

```go
func ExampleValidator_ValidateSlice() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).WithName("Teacher")

	err := v.ValidateSlice([]Teacher{
		{Name: "John"},
		{Name: "Jake"},
	})
	if err != nil {
		fmt.Println(err)
	}
}
```

## Control error aggregation

Validator cascade controls whether validation continues to later properties.
Property cascade controls the rules within that property.
A property override does not make later properties run after a validator-level stop.
Included validators retain their own cascade modes.

### Stop evaluating later properties after a validator failure

[//]: # (embed: ExampleValidator_Cascade?comments=false)

```go
func ExampleValidator_Cascade() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Cascade(govy.CascadeModeContinue).
			Rules(rules.NEQ("Jerry")).
			Rules(rules.EQ("Tom")),
		govy.For(func(t Teacher) time.Duration { return t.Age }).
			WithName("age").
			Rules(
				rules.GT(18*year),
				govy.NewRule(func(time.Duration) error {
					return fmt.Errorf("always fails")
				}),
			),
	).
		Cascade(govy.CascadeModeStop)

	for _, name := range []string{"Tom", "Jerry"} {
		teacher := Teacher{
			Name: name,
			Age:  17 * year,
		}
		err := v.WithName(name).Validate(teacher)
		if err != nil {
			fmt.Println(err)
		}
	}
}
```

## Compose a complete validator

This example combines named properties, predefined rules, and nested validators.

### Build a full Teacher validator

[//]: # (embed: ExampleValidator?comments=false)

```go
func ExampleValidator() {
	universityValidation := govy.New(
		govy.For(func(u University) string { return u.Address }).
			WithName("address").
			Required(),
	)
	studentValidator := govy.New(
		govy.For(func(s Student) string { return s.Index }).
			WithName("index").
			Rules(rules.StringLength(9, 9)),
	)
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				rules.StringNotEmpty(),
				rules.OneOf("Jake", "George"),
			),
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				rules.SliceMaxLength[[]Student](2),
				rules.SliceUnique(func(v Student) string { return v.Index }),
			).
			IncludeForEach(studentValidator),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(universityValidation),
	).When(func(t Teacher) bool { return t.Age < 50 })

	teacher := Teacher{
		Name: "John",
		Students: []Student{
			{Index: "918230014"},
			{Index: "9182300123"},
			{Index: "918230014"},
		},
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	err := teacherValidator.WithName("John").Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```
