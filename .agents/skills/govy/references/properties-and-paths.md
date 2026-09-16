# Properties and Paths

## Topics

- [Name and path properties](#name-and-path-properties)
  - [Name a property rule with one path segment.](#name-a-property-rule-with-one-path-segment)
  - [Do not pass dotted paths to WithName.](#do-not-pass-dotted-paths-to-withname)
  - [Set multi-segment paths with jsonpath.Path.](#set-multi-segment-paths-with-jsonpathpath)
- [Validate optional and transformed values](#validate-optional-and-transformed-values)
  - [Validate the pointed-to value while skipping nil by default.](#validate-the-pointed-to-value-while-skipping-nil-by-default)
  - [Transform a property value before applying rules.](#transform-a-property-value-before-applying-rules)
- [Handle empty or sensitive values](#handle-empty-or-sensitive-values)
  - [Require a non-empty value before normal rules run.](#require-a-non-empty-value-before-normal-rules-run)
  - [Skip rule evaluation for empty optional values.](#skip-rule-evaluation-for-empty-optional-values)
  - [Hide sensitive property values in errors.](#hide-sensitive-property-values-in-errors)
- [Use whole-object context](#use-whole-object-context)
  - [Validate using the whole current object.](#validate-using-the-whole-current-object)
  - [Run property rules only when predicates match.](#run-property-rules-only-when-predicates-match)
- [Control property-level aggregation](#control-property-level-aggregation)
  - [Stop evaluating later rules on one property.](#stop-evaluating-later-rules-on-one-property)

## Name and path properties

Prefer WithName for a single path segment.
Use WithPath when the rule logically targets a nested path or a specific index.

### Name a property rule with one path segment

[//]: # (embed: ExamplePropertyRules_WithName?comments=false)

```go
func ExamplePropertyRules_WithName() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.EQ("Tom")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

### Do not pass dotted paths to WithName

[//]: # (embed: ExamplePropertyRules_WithName_wrongUsage?comments=false)

```go
func ExamplePropertyRules_WithName_wrongUsage() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.University.Name }).
			WithName("university.name").
			Rules(rules.EQ("Tom").WithMessage("yikes, looks like you used WithName instead of WithPath!")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		University: University{
			Name: "Poznan University of Technology",
		},
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

### Set multi-segment paths with jsonpath.Path

[//]: # (embed: ExamplePropertyRules_WithPath?comments=false)

```go
func ExamplePropertyRules_WithPath() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.University.Name }).
			WithPath(jsonpath.Parse("university.name")).
			Rules(rules.EQ("Tom")),
		govy.For(func(t Teacher) string { return t.Students[0].Index }).
			WithPath(jsonpath.New().Name("students").Index(0).Name("index")).
			Rules(rules.EQ("2")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		University: University{
			Name: "Poznan University of Technology",
		},
		Students: []Student{
			{Index: "1"},
		},
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

## Validate optional and transformed values

Use `ForPointer` to validate a pointed-to value.
It skips `nil` unless the property is required.
`ForPointer(...).Required()` checks pointer presence.
Add a content rule to reject an empty string or zero number.

Use `Transform` when rules need a different type from the stored value.
Check zero-value behavior before combining it with `OmitEmpty`.
A nonempty duration string can transform to a zero duration.
`OmitEmpty` then skips that value.

### Validate the pointed-to value while skipping nil by default

[//]: # (embed: ExampleForPointer?comments=false)

```go
func ExampleForPointer() {
	v := govy.New(
		govy.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Rules(rules.StringMaxLength(5)),
	).WithName("Teacher")

	middleName := "Thaddeus"
	teacher := Teacher{
		Name:       "Jake",
		Age:        51 * year,
		MiddleName: &middleName,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

### Transform a property value before applying rules

[//]: # (embed: ExampleTransform?comments=false)

```go
func ExampleTransform() {
	type Clock struct {
		Duration string `json:"duration"`
	}
	v := govy.New(
		govy.Transform(func(c Clock) string { return c.Duration }, time.ParseDuration).
			WithName("duration").
			Rules(rules.DurationPrecision(time.Minute)),
	).WithName("MyClock")

	err := v.Validate(Clock{Duration: "bad duration!"})
	if err != nil {
		fmt.Println(err)
	}

	err = v.Validate(Clock{Duration: (256 * time.Second).String()})
	if err != nil {
		fmt.Println(err)
	}
}
```

## Handle empty or sensitive values

Use `Required` to stop on empty values.
Use `OmitEmpty` to skip optional direct values.
For value hiding and its limits, read [Errors](errors.md#hide-property-values).

### Require a non-empty value before normal rules run

[//]: # (embed: ExamplePropertyRules_Required?comments=false)

```go
func ExamplePropertyRules_Required() {
	alwaysFailingRule := govy.NewRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := govy.New(
		govy.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Required().
			Rules(alwaysFailingRule),
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(alwaysFailingRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name:       "",
		Age:        51 * year,
		MiddleName: nil,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

### Skip rule evaluation for empty optional values

[//]: # (embed: ExamplePropertyRules_OmitEmpty?comments=false)

```go
func ExamplePropertyRules_OmitEmpty() {
	alwaysFailingRule := govy.NewRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			OmitEmpty().
			Rules(alwaysFailingRule),
		govy.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Rules(alwaysFailingRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name:       "",
		Age:        51 * year,
		MiddleName: nil,
	}

	err := v.Validate(teacher)
	if err == nil {
		fmt.Println("no error! we skipped 'name' validation and 'middleName' is implicitly skipped")
	}
}
```

### Hide sensitive property values in errors

[//]: # (embed: ExamplePropertyRules_HideValue?comments=false)

```go
func ExamplePropertyRules_HideValue() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			HideValue().
			Rules(govy.NewRule(func(name string) error { return fmt.Errorf("that Jake is secret") })),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

## Use whole-object context

Use GetSelf when the rule compares multiple fields in the same object.
Use property-level When to keep branch-specific rules isolated.

### Validate using the whole current object

[//]: # (embed: ExampleGetSelf?comments=false)

```go
func ExampleGetSelf() {
	customRule := govy.NewRule(func(v Teacher) error {
		return fmt.Errorf("now I have access to the whole teacher")
	})

	v := govy.New(
		govy.For(govy.GetSelf[Teacher]()).
			Rules(customRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}
}
```

### Run property rules only when predicates match

[//]: # (embed: ExamplePropertyRules_When?comments=false)

```go
func ExamplePropertyRules_When() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			When(func(t Teacher) bool { return t.Name == "Jerry" }).
			Rules(rules.NEQ("Jerry")),
	).WithName("Teacher")

	for _, name := range []string{"Tom", "Jerry", "Mickey"} {
		teacher := Teacher{Name: name}
		err := v.Validate(teacher)
		if err != nil {
			fmt.Println(err)
		}
	}
}
```

## Control property-level aggregation

Property cascade overrides the inherited mode for rules within that property.
Use a stop when later rules depend on an earlier result.
A property override does not change a validator-level stop between properties.

### Stop evaluating later rules on one property

[//]: # (embed: ExamplePropertyRules_Cascade?comments=false)

```go
func ExamplePropertyRules_Cascade() {
	alwaysFailingRule := govy.NewRule(func(string) error {
		return fmt.Errorf("always fails")
	})

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Cascade(govy.CascadeModeStop).
			Rules(rules.NEQ("Jerry")).
			Rules(alwaysFailingRule),
	).WithName("Teacher")

	for _, name := range []string{"Tom", "Jerry"} {
		teacher := Teacher{Name: name}
		err := v.Validate(teacher)
		if err != nil {
			fmt.Println(err)
		}
	}
}
```
