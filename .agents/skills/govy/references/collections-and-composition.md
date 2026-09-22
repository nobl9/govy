# Collections and Composition

## Topics

- [Compose nested validators](#compose-nested-validators)
  - [Include a validator for a nested object.](#include-a-validator-for-a-nested-object)
- [Validate collection elements](#validate-collection-elements)
  - [Validate a slice and each of its elements.](#validate-a-slice-and-each-of-its-elements)
  - [Handle slices whose elements are pointers.](#handle-slices-whose-elements-are-pointers)
  - [Validate map keys, values, and key-value items.](#validate-map-keys-values-and-key-value-items)
- [Derive validator variants](#derive-validator-variants)
  - [Remove selected properties by path.](#remove-selected-properties-by-path)

## Compose nested validators

Use Include when a property has its own validator.
Govy appends nested paths automatically so errors still point to the leaf property.

### Include a validator for a nested object

[//]: # (embed: ExamplePropertyRules_Include?comments=false)

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
}
```

## Validate collection elements

Use collection builders to validate a collection and its elements.
They preserve map keys or slice indexes in paths.

### Validate a slice and each of its elements

[//]: # (embed: ExampleForSlice?comments=false)

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
				rules.SliceUnique(func(v Student) string { return v.Index }),
			).
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
}
```

### Handle slices whose elements are pointers

[//]: # (embed: ExampleForSlice_sliceOfPointers?comments=false)

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
}
```

### Validate map keys, values, and key-value items

[//]: # (embed: ExampleForMap?comments=false)

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
}
```

## Derive validator variants

Derive variants with `RemovePropertiesByPath` or `RemovePropertiesByID`.
The original validator remains unchanged.

### Remove selected properties by ID

Set `WithID` on scalar, slice, or map property rules.
`RemovePropertiesByID` removes every property with a matching nonempty ID.
It traverses included `Validator` values and pointers, including recursive references.
Other `ValidatorInterface` implementations, such as wrappers with custom `Validate`
methods, remain unchanged.
Tag the containing property to remove an entire wrapper.
Empty and unknown IDs have no effect.

[//]: # (embed: ExampleValidator_RemovePropertiesByID?comments=false)

```go
func ExampleValidator_RemovePropertiesByID() {
	ageProperty := govy.For(func(t Teacher) time.Duration { return t.Age }).
		WithName("age").
		WithID("age").
		Rules(rules.GT(time.Duration(0)))
	baseValidator := govy.New(ageProperty)
	modifiedValidator := baseValidator.RemovePropertiesByID("age")
	teacher := Teacher{Age: -1}

	fmt.Println(baseValidator.Validate(teacher) != nil)
	fmt.Println(modifiedValidator.Validate(teacher) != nil)
}
```

### Remove selected properties by path

[//]: # (embed: ExampleValidator_RemovePropertiesByPath?comments=false)

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

	err := baseValidator.Validate(teacher)
	if err != nil {
		fmt.Println("Base validator failed")
	}

	modifiedValidator := baseValidator.RemovePropertiesByPath(jsonpath.New().Name("age"))
	err = modifiedValidator.Validate(teacher)
	if err == nil {
		fmt.Println("Modified validator passed")
	}
}
```
