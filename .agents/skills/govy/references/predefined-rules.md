# Predefined Rules

Examples from the rules package for collection uniqueness, property relationships, and comparable values.

## Topics

- [Collection constraints](#collection-constraints)
  - [Require values in a slice to be unique.](#require-values-in-a-slice-to-be-unique)
- [Property relationship rules](#property-relationship-rules)
  - [Require exactly one value across mutually exclusive properties.](#require-exactly-one-value-across-mutually-exclusive-properties)
  - [Require one populated value from a set of properties.](#require-one-populated-value-from-a-set-of-properties)
- [Property comparison rules](#property-comparison-rules)
  - [Compare two properties for equality.](#compare-two-properties-for-equality)
  - [Compare ordered primitive properties.](#compare-ordered-primitive-properties)
  - [Compare custom comparable properties.](#compare-custom-comparable-properties)

## Collection constraints

Use collection rules when the validation decision depends on all values, not a single property value.

<a id="require-values-in-a-slice-to-be-unique"></a>

**Require values in a slice to be unique.**

[//]: # (embed: ExampleSliceUnique)

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

## Property relationship rules

Use property relationship rules when one field controls whether another field may or must be set.

<a id="require-exactly-one-value-across-mutually-exclusive-properties"></a>

**Require exactly one value across mutually exclusive properties.**

[//]: # (embed: ExampleMutuallyExclusive)

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

<a id="require-one-populated-value-from-a-set-of-properties"></a>

**Require one populated value from a set of properties.**

[//]: # (embed: ExampleOneOfProperties)

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

## Property comparison rules

Use comparison rules when multiple properties must preserve ordering or equality. Pick comparable variants for types such as time.Time that define their own ordering contract.

<a id="compare-two-properties-for-equality"></a>

**Compare two properties for equality.**

[//]: # (embed: ExampleEqualProperties)

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

<a id="compare-ordered-primitive-properties"></a>

**Compare ordered primitive properties.**

[//]: # (embed: ExampleLTProperties)

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

<a id="compare-custom-comparable-properties"></a>

**Compare custom comparable properties.**

[//]: # (embed: ExampleLTComparableProperties)

```go
// LTComparableProperties and other *ComparableProperties functions work with types
// that implement [rules.Comparable] interface, such as [time.Time].
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
