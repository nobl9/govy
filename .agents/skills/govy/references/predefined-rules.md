# Predefined Rules

Examples from the rules package for collection uniqueness, property relationships, and comparable values.

Source examples:

- [pkg/rules/example_test.go](../../../../pkg/rules/example_test.go)

## Examples

- [ExampleSliceUnique](#examplesliceunique)
- [ExampleMutuallyExclusive](#examplemutuallyexclusive)
- [ExampleOneOfProperties](#exampleoneofproperties)
- [ExampleEqualProperties](#exampleequalproperties)
- [ExampleLTProperties](#exampleltproperties)
- [ExampleLTComparableProperties](#exampleltcomparableproperties)

## ExampleSliceUnique

Source: [pkg/rules/example_test.go:21](../../../../pkg/rules/example_test.go#L21)

[//]: # (embed: pkg/rules/example_test.go#ExampleSliceUnique)

```go
func ExampleSliceUnique() {
	v := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(rules.SliceUnique(func(v Student) string { return v.Index },
				"each student must have unique index")),
	)
	teacher := Teacher{
		Students: []Student{
			{Index: "foo"},
			{Index: "bar"}, // 2nd element
			{Index: "baz"},
			{Index: "bar"}, // 4th element
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students' with value '[{"index":"foo"},{"index":"bar"},{"index":"baz"},{"index":"bar"}]':
	//     - elements are not unique, 2nd and 4th elements collide based on constraints: each student must have unique index
}
```

## ExampleMutuallyExclusive

Source: [pkg/rules/example_test.go:47](../../../../pkg/rules/example_test.go#L47)

[//]: # (embed: pkg/rules/example_test.go#ExampleMutuallyExclusive)

```go
func ExampleMutuallyExclusive() {
	v := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			RulesForEach(rules.MutuallyExclusive(true, map[string]func(Student) any{
				"index": func(s Student) any { return s.Index },
				"name":  func(s Student) any { return s.Name },
			})),
	)
	teacher := Teacher{
		Students: []Student{
			{Index: "foo"},
			{Index: "bar", Name: "John"},
			{Name: "Eve"},
			{},
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students[1]' with value '{"index":"bar","name":"John"}':
	//     - [index, name] properties are mutually exclusive, provide only one of them
	//   - 'students[3]':
	//     - one of [index, name] properties must be set, none was provided
}
```

## ExampleOneOfProperties

Source: [pkg/rules/example_test.go:77](../../../../pkg/rules/example_test.go#L77)

[//]: # (embed: pkg/rules/example_test.go#ExampleOneOfProperties)

```go
func ExampleOneOfProperties() {
	v := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			RulesForEach(rules.OneOfProperties(map[string]func(Student) any{
				"index": func(s Student) any { return s.Index },
				"name":  func(s Student) any { return s.Name },
			})),
	)
	teacher := Teacher{
		Students: []Student{
			{Index: "foo"},
			{},
			{Name: "John"},
			{Index: "bar", Name: "Eve"},
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students[1]':
	//     - one of [index, name] properties must be set, none was provided
}
```

## ExampleEqualProperties

Source: [pkg/rules/example_test.go:105](../../../../pkg/rules/example_test.go#L105)

[//]: # (embed: pkg/rules/example_test.go#ExampleEqualProperties)

```go
func ExampleEqualProperties() {
	v := govy.New(
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			RulesForEach(rules.EqualProperties(rules.CompareFunc, map[string]func(Student) any{
				"index":     func(s Student) any { return s.Index },
				"indexCopy": func(s Student) any { return s.IndexCopy },
			})),
	)
	teacher := Teacher{
		Students: []Student{
			{Index: "foo", IndexCopy: "foo"},
			{Index: "bar"},
			{IndexCopy: "foo"},
			{}, // Both index and indexCopy are empty strings, and thus equal.
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'students[1]' with value '{"index":"bar"}':
	//     - all of [index, indexCopy] properties must be equal, but 'index' is not equal to 'indexCopy'
	//   - 'students[2]' with value '{"indexCopy":"foo"}':
	//     - all of [index, indexCopy] properties must be equal, but 'index' is not equal to 'indexCopy'
}
```

## ExampleLTProperties

Source: [pkg/rules/example_test.go:135](../../../../pkg/rules/example_test.go#L135)

[//]: # (embed: pkg/rules/example_test.go#ExampleLTProperties)

```go
func ExampleLTProperties() {
	type IntRange struct {
		Min int `json:"min"`
		Max int `json:"max"`
	}

	v := govy.New(
		govy.For(govy.GetSelf[IntRange]()).
			Rules(
				rules.LTProperties(
					"min", func(r IntRange) int { return r.Min },
					"max", func(r IntRange) int { return r.Max },
				),
			),
	)

	// Valid case: min < max
	err := v.Validate(IntRange{Min: 1, Max: 10})
	fmt.Println("Valid:", err == nil)

	// Invalid case: min >= max
	err = v.Validate(IntRange{Min: 10, Max: 1})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Valid: true
	// Validation has failed:
	//   - 'min' must be less than 'max'
}
```

## ExampleLTComparableProperties

Source: [pkg/rules/example_test.go:169](../../../../pkg/rules/example_test.go#L169)

LTComparableProperties and other *ComparableProperties functions work with types
that implement [rules.Comparable] interface, such as [time.Time].

[//]: # (embed: pkg/rules/example_test.go#ExampleLTComparableProperties)

```go
func ExampleLTComparableProperties() {
	type TimeRange struct {
		StartTime time.Time `json:"startTime"`
		EndTime   time.Time `json:"endTime"`
	}

	v := govy.New(
		govy.For(govy.GetSelf[TimeRange]()).
			Rules(
				rules.LTComparableProperties(
					"startTime", func(tr TimeRange) time.Time { return tr.StartTime },
					"endTime", func(tr TimeRange) time.Time { return tr.EndTime },
				),
			),
	)

	// Valid case: start is before end
	err := v.Validate(TimeRange{
		StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
	})
	fmt.Println("Valid:", err == nil)

	// Invalid case: start is after end
	err = v.Validate(TimeRange{
		StartTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Valid: true
	// Validation has failed:
	//   - 'startTime' must be before 'endTime'
}
```

