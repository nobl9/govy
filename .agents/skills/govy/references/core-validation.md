# Core Validation

Validator construction, naming, conditions, slice validation, cascade behavior, and validator composition patterns.

Source examples:

- [pkg/govy/example_test.go](../../../../pkg/govy/example_test.go)

## Examples

- [ExampleNew](#examplenew)
- [ExampleValidator_WithName](#examplevalidator_withname)
- [ExampleValidator_WithNameFunc](#examplevalidator_withnamefunc)
- [ExampleValidator_When](#examplevalidator_when)
- [ExampleValidator_Validate_slice](#examplevalidator_validate_slice)
- [ExampleValidator_ValidateSlice](#examplevalidator_validateslice)
- [ExampleValidator_Cascade](#examplevalidator_cascade)
- [ExampleValidator](#examplevalidator)
- [ExampleValidator_branchingPattern](#examplevalidator_branchingpattern)

## ExampleNew

Source: [pkg/govy/example_test.go:44](../../../../pkg/govy/example_test.go#L44)

In order to create a new [govy.Validator] use [govy.New] constructor.
Let's define simple [govy.PropertyRules] for [Teacher.Name].
For now, it will be always failing.

[//]: # (embed: pkg/govy/example_test.go#ExampleNew)

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

	// Output:
	// Validation has failed:
	//   - always fails
}
```

## ExampleValidator_WithName

Source: [pkg/govy/example_test.go:62](../../../../pkg/govy/example_test.go#L62)

To associate [govy.Validator] with an entity name use [govy.Validator.WithName] function.
When any of the rules fails, the error will contain the entity name you've provided.

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator_WithName)

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

	// Output:
	// Validation for Teacher has failed:
	//   - always fails
}
```

## ExampleValidator_WithNameFunc

Source: [pkg/govy/example_test.go:81](../../../../pkg/govy/example_test.go#L81)

If statically defined name through [govy.Validator.WithName] is not enough,
you can use [govy.Validator.WithNameFunc].
The function receives the entity's instance you're validating and returns a string name.

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator_WithNameFunc)

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

	// Output:
	// Validation for Teacher John has failed:
	//   - always fails
}
```

## ExampleValidator_When

Source: [pkg/govy/example_test.go:130](../../../../pkg/govy/example_test.go#L130)

[govy.Validator] rules can be evaluated on condition, to specify the predicate use [govy.Validator.When] function.

In this example, validation for [Teacher] instance will only be evaluated
if the [Teacher.Age] property is less than 50 years.

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator_When)

```go
func ExampleValidator_When() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("always fails") })),
	).
		When(func(t Teacher) bool { return t.Age < (50 * year) })

	// Prepare teachers.
	teacherTom := Teacher{
		Name: "Tom",
		Age:  51 * year,
	}
	teacherJerry := Teacher{
		Name: "Jerry",
		Age:  30 * year,
	}

	// Run validation.
	err := v.Validate(teacherTom)
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Tom"))
	}
	err = v.Validate(teacherJerry)
	if err != nil {
		fmt.Println(err.(*govy.ValidatorError).WithName("Jerry"))
	}

	// Output:
	// Validation for Jerry has failed:
	//   - always fails
}
```

## ExampleValidator_Validate_slice

Source: [pkg/govy/example_test.go:207](../../../../pkg/govy/example_test.go#L207)

If you want to validate a slice of entities, you can combine [govy.New] with [govy.ForSlice].
The produced errors will contain information about the failing entity's index
in their [govy.PropertyError.PropertyPath].

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator_Validate_slice)

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

	// Output:
	// Validation has failed for the following properties:
	//   - '[0].name' with value 'John':
	//     - always fails
	//   - '[1].name' with value 'Jake':
	//     - always fails
}
```

## ExampleValidator_ValidateSlice

Source: [pkg/govy/example_test.go:1556](../../../../pkg/govy/example_test.go#L1556)

If combining [govy.New] with [govy.ForSlice] is not verbose enough for you,
you can use [govy.Validator.ValidateSlice] function.
It will validate each element according to the rules defined by [govy.Validator].
It returns [govy.ValidatorErrors].

Note: If you need to perform additional validation on the whole slice,
you should rather use [govy.New] with [govy.ForSlice] and [govy.GetSelf].
[govy.Validator.ValidateSlice] is designed to be used for processing independent values.

Note: Since each element is validated in isolation,
the reported property paths will not start with the slice index,
they will instead start at the element's root.

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator_ValidateSlice)

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

	// Output:
	// Validation for Teacher at index 0 has failed for the following properties:
	//   - 'name' with value 'John':
	//     - always fails
	// Validation for Teacher at index 1 has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - always fails
}
```

## ExampleValidator_Cascade

Source: [pkg/govy/example_test.go:1588](../../../../pkg/govy/example_test.go#L1588)

Unlike [govy.PropertyRules.Cascade] which works on [govy.PropertyRules] level,
[govy.Validator.Cascade] propagates to all the properties of [govy.Validator] and
furthermore, will stop evaluating the next property if any preceding property fails.

If [govy.PropertyRules.Cascade] is set, the setting will take precedence over
[govy.Validator] cascade mode.

See [ExamplePropertyRules_Cascade] for more details on [govy.PropertyRules.Cascade].

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator_Cascade)

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

	// Output:
	// Validation for Tom has failed for the following properties:
	//   - 'age' with value '148920h0m0s':
	//     - must be greater than '157680h0m0s'
	// Validation for Jerry has failed for the following properties:
	//   - 'name' with value 'Jerry':
	//     - must not be equal to 'Jerry'
	//     - must be equal to 'Tom'
}
```

## ExampleValidator

Source: [pkg/govy/example_test.go:1690](../../../../pkg/govy/example_test.go#L1690)

Bringing it all (mostly) together, let's create a fully fledged [govy.Validator] for [Teacher].

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator)

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
				rules.OneOf("Jake", "George")),
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				rules.SliceMaxLength[[]Student](2),
				rules.SliceUnique(func(v Student) string { return v.Index })).
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

	// Output:
	// Validation for John has failed for the following properties:
	//   - 'name' with value 'John':
	//     - must be one of: Jake, George
	//   - 'students' with value '[{"index":"918230014"},{"index":"9182300123"},{"index":"918230014"}]':
	//     - length must be less than or equal to 2
	//     - elements are not unique, 1st and 3rd elements collide
	//   - 'students[1].index' with value '9182300123':
	//     - length must be between 9 and 9
	//   - 'university.address':
	//     - property is required but was empty
}
```

## ExampleValidator_branchingPattern

Source: [pkg/govy/example_test.go:1760](../../../../pkg/govy/example_test.go#L1760)

When dealing with properties that should only be validated if a certain other
property has specific value, it's recommended to use [govy.PropertyRules.When] and [govy.PropertyRules.Include]
to separate validation paths into non-overlapping branches.

Notice how in the below example [File.Format] is the common,
shared property between [CSV] and [JSON] files.
We define separate [govy.Validator] for [CSV] and [JSON] and use [govy.PropertyRules.When] to only validate
their included [govy.Validator] if the correct [File.Format] is provided.

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator_branchingPattern)

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

	// Output:
	// Validation for File has failed for the following properties:
	//   - 'indent' with value 'invalid':
	//     - string must match regular expression: '^\s*$'
}
```

