# Predefined Rules

Examples that show selected predefined rules in use.
For available constructors, read [Existing Rules](existing-rules.md).

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

Use collection rules for constraints that depend on the collection's values.

### Require values in a slice to be unique

[//]: # (embed: ExampleSliceUnique?comments=false)

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
			{Index: "bar"},
			{Index: "baz"},
			{Index: "bar"},
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

## Property relationship rules

Use relationship rules when one field controls another field's presence.

### Require exactly one value across mutually exclusive properties

[//]: # (embed: ExampleMutuallyExclusive?comments=false)

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
}
```

### Require one populated value from a set of properties

[//]: # (embed: ExampleOneOfProperties?comments=false)

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
}
```

## Property comparison rules

Use comparison rules when multiple properties must preserve ordering or equality.
Use comparable variants for types with custom ordering, such as `time.Time`.

### Compare two properties for equality

[//]: # (embed: ExampleEqualProperties?comments=false)

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
			{},
		},
	}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

### Compare ordered primitive properties

[//]: # (embed: ExampleLTProperties?comments=false)

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

	err := v.Validate(IntRange{Min: 1, Max: 10})
	fmt.Println("Valid:", err == nil)

	err = v.Validate(IntRange{Min: 10, Max: 1})
	if err != nil {
		fmt.Println(err)
	}
}
```

### Compare custom comparable properties

[//]: # (embed: ExampleLTComparableProperties?comments=false)

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

	err := v.Validate(TimeRange{
		StartTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
	})
	fmt.Println("Valid:", err == nil)

	err = v.Validate(TimeRange{
		StartTime: time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if err != nil {
		fmt.Println(err)
	}
}
```
