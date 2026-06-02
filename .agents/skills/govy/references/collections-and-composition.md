# Collections and Composition

Nested validators, slice and map property rules, slice pointers, and validator variants derived from path removal.

Source examples:

- [pkg/govy/example_test.go](../../../../pkg/govy/example_test.go)

## Examples

- [ExamplePropertyRules_Include](#examplepropertyrules_include)
- [ExampleForSlice](#exampleforslice)
- [ExampleForSlice_sliceOfPointers](#exampleforslice_sliceofpointers)
- [ExampleForMap](#exampleformap)
- [ExampleValidator_RemovePropertiesByPath](#examplevalidator_removepropertiesbypath)

## ExamplePropertyRules_Include

Source: [pkg/govy/example_test.go:1240](../../../../pkg/govy/example_test.go#L1240)

So far we've defined validation rules for simple, top-level properties.
What If we want to define validation rules for nested properties?
We can use [govy.PropertyRules.Include] to include another [govy.Validator] in our [govy.PropertyRules].

Let's extend our [Teacher] struct to include a nested [University] property.
[University] in of itself is another struct with its own validation rules.

Notice how the nested property path is automatically built for you,
each segment separated by a dot.

[//]: # (embed: pkg/govy/example_test.go#ExamplePropertyRules_Include)

```go
func ExamplePropertyRules_Include() {
	universityValidation := govy.New(
		govy.For(func(u University) string { return u.Address }).
			WithName("address").
			Required(),
	)
	teacherValidation := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.EQ("Tom")),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(universityValidation),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jerry",
		Age:  51 * year,
		University: University{
			Name:    "Poznan University of Technology",
			Address: "",
		},
	}

	err := teacherValidation.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jerry':
	//     - must be equal to 'Tom'
	//   - 'university.address':
	//     - property is required but was empty
}
```

## ExampleForSlice

Source: [pkg/govy/example_test.go:1299](../../../../pkg/govy/example_test.go#L1299)

When dealing with slices we often want to both validate the whole slice
and each of its elements.
You can use [govy.ForSlice] function to do just that.
It returns a new struct [govy.PropertyRulesForSlice] which behaves exactly
the same as [govy.PropertyRules], but extends its API slightly.

To define rules for each element use:
  - [govy.PropertyRulesForSlice.RulesForEach]
  - [govy.PropertyRulesForSlice.IncludeForEach]

These work exactly the same way as [govy.PropertyRules.Rules] and [govy.PropertyRules.Include]
verifying each slice element.

[govy.PropertyRulesForSlice.Rules] is in turn used to define rules for the whole slice.

Note: [govy.PropertyRulesForSlice] does not implement Include function for the whole slice.

In the below example, we're defining that students slice must have at most 2 elements
and that each element's index must be unique.
For each element we're also including [Student] [govy.Validator].
Notice that property path for slices has the following format:
<slice_name>[<index>].<slice_property_name>

[//]: # (embed: pkg/govy/example_test.go#ExampleForSlice)

```go
func ExampleForSlice() {
	studentValidator := govy.New(
		govy.For(func(s Student) string { return s.Index }).
			WithName("index").
			Rules(rules.StringLength(9, 9)),
	)
	teacherValidator := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				rules.SliceMaxLength[[]Student](2),
				rules.SliceUnique(func(v Student) string { return v.Index })).
			IncludeForEach(studentValidator),
	).When(func(t Teacher) bool { return t.Age < 50 })

	teacher := Teacher{
		Name: "John",
		Students: []Student{
			{Index: "918230014"},
			{Index: "9182300123"},
			{Index: "918230014"},
		},
	}

	err := teacherValidator.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students' with value '[{"index":"918230014"},{"index":"9182300123"},{"index":"918230014"}]':
	//     - length must be less than or equal to 2
	//     - elements are not unique, 1st and 3rd elements collide
	//   - 'students[1].index' with value '9182300123':
	//     - length must be between 9 and 9
}
```

## ExampleForSlice_sliceOfPointers

Source: [pkg/govy/example_test.go:1351](../../../../pkg/govy/example_test.go#L1351)

When dealing with slices of pointers you may find it problematic to add [govy.Rule]
with [govy.PropertyRulesForSlice.RulesForEach].
The builtin rules, and most likely your custom rules as well, all operate on non-pointer values.
This means you cannot use them on your slice's pointer elements.

To solve this problem you can use [govy.ForPointer] constructor and convert any [govy.Rule]
to work on pointers.

In the below example we're defining two [govy.Validator] instances:
  - 'faultyValidator' which will not fail for 'nil' value
  - 'goodValidator' which will fail for 'nil' value by using [rules.Required] rule

This behavior is consistent with [govy.ForPointer] constructor, which will skip the validation
unless you add [govy.PropertyRules.Required] to enforce the value to be a non-nil pointer.

[//]: # (embed: pkg/govy/example_test.go#ExampleForSlice_sliceOfPointers)

```go
func ExampleForSlice_sliceOfPointers() {
	type Pointers struct {
		Pointers []*string `json:"pointers"`
	}
	pointersRules := govy.ForSlice(func(p Pointers) []*string { return p.Pointers }).
		WithName("pointers").
		Rules(rules.SliceMaxLength[[]*string](2)).
		RulesForEach(
			govy.RuleToPointer(rules.StringLength(9, 9)),
		)
	faultyValidator := govy.New(
		pointersRules,
	)
	goodValidator := govy.New(
		pointersRules.RulesForEach(rules.Required[*string]()),
	)

	pointers := Pointers{
		Pointers: []*string{ptr("918230014"), ptr("9182300123"), ptr("918230014"), nil},
	}

	err := faultyValidator.Validate(pointers)
	if err != nil {
		fmt.Println(err)
	}
	err = goodValidator.Validate(pointers)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'pointers' with value '["918230014","9182300123","918230014",null]':
	//     - length must be less than or equal to 2
	//   - 'pointers[1]' with value '9182300123':
	//     - length must be between 9 and 9
	// Validation has failed for the following properties:
	//   - 'pointers' with value '["918230014","9182300123","918230014",null]':
	//     - length must be less than or equal to 2
	//   - 'pointers[1]' with value '9182300123':
	//     - length must be between 9 and 9
	//   - 'pointers[3]':
	//     - property is required but was empty
}
```

## ExampleForMap

Source: [pkg/govy/example_test.go:1428](../../../../pkg/govy/example_test.go#L1428)

When dealing with maps there are three forms of iteration:
  - keys
  - values
  - key-value pairs (items)

You can use [govy.ForMap] function to define rules for all the aforementioned iterators.
It returns a new struct [govy.PropertyRulesForMap] which behaves similar to
[govy.PropertyRulesForSlice]..

To define rules for keys use:
  - [govy.PropertyRulesForMap.RulesForKeys]
  - [govy.PropertyRulesForMap.IncludeForKeys]
  - [govy.PropertyRulesForMap.RulesForValues]
  - [govy.PropertyRulesForMap.IncludeForValues]
  - [govy.PropertyRulesForMap.RulesForItems]
  - [govy.PropertyRulesForMap.IncludeForItems]

These work exactly the same way as [govy.PropertyRules.Rules] and [govy.PropertyRules.Include]
verifying each map's key, value or [govy.MapItem].

[govy.PropertyRulesForMap.Rules] is in turn used to define rules for the whole map.

Note: [govy.PropertyRulesForMap] does not implement Include function for the whole map.

In the below example, we're defining that student index to [Teacher] map:
  - Must have at most 2 elements (map).
  - Keys must have a length of 9 (keys).
  - Eve cannot be a teacher for any student (values).
  - Joan cannot be a teacher for student with index 918230013 (items).

Notice that property path for maps has the following format:
<map_name>.<key>.<map_property_name>

[//]: # (embed: pkg/govy/example_test.go#ExampleForMap)

```go
func ExampleForMap() {
	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.NEQ("Eve")),
	)
	tutoringValidator := govy.New(
		govy.ForMap(func(t Tutoring) map[string]Teacher { return t.StudentIndexToTeacher }).
			WithName("students").
			Rules(
				rules.MapMaxLength[map[string]Teacher](2),
			).
			RulesForKeys(
				rules.StringLength(9, 9),
			).
			IncludeForValues(teacherValidator).
			RulesForItems(govy.NewRule(func(v govy.MapItem[string, Teacher]) error {
				if v.Key == "918230013" && v.Value.Name == "Joan" {
					return govy.NewRuleError(
						"Joan cannot be a teacher for student with index 918230013",
						"joan_teacher",
					)
				}
				return nil
			})),
	)

	tutoring := Tutoring{
		StudentIndexToTeacher: map[string]Teacher{
			"918230013":  {Name: "Joan"},
			"9182300123": {Name: "Eve"},
			"918230014":  {Name: "Joan"},
		},
	}

	err := tutoringValidator.Validate(tutoring)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students' with value '{"9182300123":{"name":"Eve","age":0,"students":null,"university":{"name":"","address":""}},"91823001...':
	//     - length must be less than or equal to 2
	//   - 'students['9182300123']' with key '9182300123':
	//     - length must be between 9 and 9
	//   - 'students['9182300123'].name' with value 'Eve':
	//     - must not be equal to 'Eve'
	//   - 'students['918230013']' with value '{"name":"Joan","age":0,"students":null,"university":{"name":"","address":""}}':
	//     - Joan cannot be a teacher for student with index 918230013
}
```

## ExampleValidator_RemovePropertiesByPath

Source: [pkg/govy/example_test.go:1926](../../../../pkg/govy/example_test.go#L1926)

This example demonstrates how to remove specific properties from a [govy.Validator] by their paths.
This is useful when you want to create a modified validator without certain rules.

[//]: # (embed: pkg/govy/example_test.go#ExampleValidator_RemovePropertiesByPath)

```go
func ExampleValidator_RemovePropertiesByPath() {
	baseValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringNotEmpty()),
		govy.For(func(t Teacher) time.Duration { return t.Age }).
			WithName("age").
			Rules(rules.GT(time.Duration(0))),
	)

	teacher := Teacher{Name: "John", Age: -1}

	// Base validator fails because age is negative
	err := baseValidator.Validate(teacher)
	if err != nil {
		fmt.Println("Base validator failed")
	}

	// Modified validator passes because age validation is removed
	modifiedValidator := baseValidator.RemovePropertiesByPath(jsonpath.New().Name("age"))
	err = modifiedValidator.Validate(teacher)
	if err == nil {
		fmt.Println("Modified validator passed")
	}

	// Output:
	// Base validator failed
	// Modified validator passed
}
```

