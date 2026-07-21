# Properties and Paths

Property getters, names, explicit JSON paths, pointers, transforms, required and optional values, hidden values, and property-level conditions.

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

Prefer WithName for a single path segment. Use WithPath when the rule logically targets a nested path or a specific index.

<a id="name-a-property-rule-with-one-path-segment"></a>

**Name a property rule with one path segment.**

[//]: # (embed: ExamplePropertyRules_WithName)

```go
// So far we've been using a very simple [govy.PropertyRules] instance:
//
//	validation.For(func(t Teacher) string { return t.Name }).
//		Rules(validation.NewRule(func(name string) error { return fmt.Errorf("always fails") }))
//
// The error message returned by this property rule does not tell us
// which property is failing.
// Let's change that by adding an explicit path segment using [govy.PropertyRules.WithName].
//
// We can also change the [govy.Rule] to be something more real.
// govy comes with a number of predefined [govy.Rule], we'll use
// [rules.EQ] which accepts a single argument, value to compare with.
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

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - must be equal to 'Tom'
}
```

<a id="do-not-pass-dotted-paths-to-withname"></a>

**Do not pass dotted paths to WithName.**

[//]: # (embed: ExamplePropertyRules_WithName_wrongUsage)

```go
// Beware that anything passed into [govy.PropertyRules.WithName] is treated as a single path segment.
// If you pass a dot-separated path-like string into this method, govy renders
// the dots as escaped characters inside one bracket-quoted segment.
// For multi-segment paths, use [govy.PropertyRules.WithPath] instead.
//
// Note: Prior to v0.25.0, [govy.PropertyRules.WithName] treated every string
// as a path, so this usage was valid then.
func ExamplePropertyRules_WithName_wrongUsage() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.University.Name }).
			WithName("university.name"). // WRONG USAGE!
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

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - '['university.name']' with value 'Poznan University of Technology':
	//     - yikes, looks like you used WithName instead of WithPath!
}
```

<a id="set-multi-segment-paths-with-jsonpathpath"></a>

**Set multi-segment paths with jsonpath.Path.**

[//]: # (embed: ExamplePropertyRules_WithPath)

```go
// While [govy.PropertyRules.WithName] is convenient and we recommend using it,
// sometimes you might want to define rules that access nested fields directly.
// That's what [govy.PropertyRules.WithPath] is for.
//
// Unlike [govy.PropertyRules.WithName], [govy.PropertyRules.WithPath] accepts a
// [jsonpath.Path] with one or more segments.
// [govy.PropertyRules.WithName] is just shorthand for `jsonpath.New().Name(...)`.
//
// You can either:
//   - pass a string representation of path directly with [jsonpath.Parse]
//   - construct the path with a builder API, starting with [jsonpath.New]
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

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'university.name' with value 'Poznan University of Technology':
	//     - must be equal to 'Tom'
	//   - 'students[0].index' with value '1':
	//     - must be equal to '2'
}
```

## Validate optional and transformed values

Use ForPointer for optional pointer fields and Transform when rule input should differ from the stored field type.

<a id="validate-the-pointed-to-value-while-skipping-nil-by-default"></a>

**Validate the pointed-to value while skipping nil by default.**

[//]: # (embed: ExampleForPointer)

```go
// [govy.For] constructor creates new [govy.PropertyRules] instance.
// It's only argument, [govy.PropertyGetter] is used to extract the property value.
// It works fine for direct values, but falls short when working with pointers.
// Often times we use pointers to indicate that a property is optional,
// or we want to discern between nil and zero values.
// In either case we want our validation rules to work on direct values,
// not the pointer, otherwise we'd have to always check if pointer != nil.
//
// [govy.ForPointer] constructor can be used to solve this problem and allow
// us to work with the underlying value in our rules.
// Under the hood it wraps [govy.PropertyGetter] and safely extracts the underlying value.
// If the value was nil, it will not attempt to evaluate any rules for this property.
// The rationale for that is it doesn't make sense to evaluate any rules for properties
// which are essentially empty. The only rule that makes sense in this context is to
// ensure the property is required.
// We'll learn about a way to achieve that in the next example: [ExamplePropertyRules_Required].
//
// Let's define a rule for [Teacher.MiddleName] property.
// Not everyone has to have a middle name, that's why we've defined this field
// as a pointer to string, rather than a string itself.
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

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'middleName' with value 'Thaddeus':
	//     - length must be less than or equal to 5
}
```

<a id="transform-a-property-value-before-applying-rules"></a>

**Transform a property value before applying rules.**

[//]: # (embed: ExampleTransform)

```go
// [govy.Transform] constructor can be used to transform the property's value
// before it's passed to the rules' evaluation.
// It's useful when you want to use rules that operate on a different type than the property's.
//
// Along with the standard [govy.PropertyGetter] it accepts a [govy.Transformer] function
// which takes the property value and returns the transformed value along with an error.
// If the error is not nil, the validation will fail with the error message returned by [govy.Transformer] error.
//
// In this example we'll use [time.ParseDuration] to transform the string value of [Clock.Duration] to [time.Duration].
// The first value we'll validate will force [govy.Transformer] to return an error,
// the second will succeed transformation, but it will fail the validation for [rules.DurationPrecision].
//
// Notice how the [govy.Transformer] shape adheres to a lot of standard library conversion/parsing functions.
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

	// Output:
	// Validation for MyClock has failed for the following properties:
	//   - 'duration' with value 'bad duration!':
	//     - time: invalid duration "bad duration!"
	// Validation for MyClock has failed for the following properties:
	//   - 'duration' with value '4m16s':
	//     - duration must be defined with 1m0s precision
}
```

## Handle empty or sensitive values

Use Required to fail fast on empty values, OmitEmpty to skip optional direct values, and HideValue when errors must not expose the original input.

<a id="require-a-non-empty-value-before-normal-rules-run"></a>

**Require a non-empty value before normal rules run.**

[//]: # (embed: ExamplePropertyRules_Required)

```go
// By default, when [govy.PropertyRules] is constructed using [govy.ForPointer]
// it will skip validation of the property if the pointer is nil.
// To enforce a value is set for pointer use [govy.PropertyRules.Required].
//
// You may ask yourself why not just use [rules.Required] rule instead?
// If we were to do that, we'd be forced to operate on pointer in all of our rules.
// Other than checking if the pointer is nil, there aren't any rules which would
// benefit from working on the pointer instead of the underlying value.
//
// If you want to also make sure the underlying value is filled,
// i.e. it's not a zero value, you can also use [rules.Required] rule
// on top of [govy.PropertyRules.Required].
//
// [govy.PropertyRules.Required] when used with [govy.For] constructor, will ensure
// the property does not contain a zero value.
//
// Note: [govy.PropertyRules.Required] is introducing a short circuit.
// If the assertion fails, validation will stop and return [rules.ErrorCodeRequired].
// None of the rules you've defined would be evaluated.
//
// Note: Placement of [govy.PropertyRules.Required] does not matter,
// it's not evaluated in a sequential loop, unlike standard [govy.Rule].
// However, we recommend you always place it below [govy.PropertyRules.WithName]
// to make your rules more readable.
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

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'middleName':
	//     - property is required but was empty
	//   - 'name':
	//     - property is required but was empty
}
```

<a id="skip-rule-evaluation-for-empty-optional-values"></a>

**Skip rule evaluation for empty optional values.**

[//]: # (embed: ExamplePropertyRules_OmitEmpty)

```go
// While [govy.ForPointer] will by default omit validation for nil pointers,
// it might be useful to have a similar behavior for optional properties
// which are direct values.
// [govy.PropertyRules.OmitEmpty] will do the trick.
//
// Note: [govy.PropertyRules.OmitEmpty] will have no effect on pointers handled
// by [govy.ForPointer], as they already behave in the same way.
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

	// Output:
	// no error! we skipped 'name' validation and 'middleName' is implicitly skipped
}
```

<a id="hide-sensitive-property-values-in-errors"></a>

**Hide sensitive property values in errors.**

[//]: # (embed: ExamplePropertyRules_HideValue)

```go
// Sometimes you want to hide the value of the property in the error message.
// It can contain sensitive information, like a secret access key.
// You can use [govy.PropertyRules.HideValue] to achieve that.
//
// You can see that the error message now contains "[hidden]" instead of the actual value,
// and the property value is not included in the property bullet point (- 'name').
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

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name':
	//     - that [hidden] is secret
}
```

## Use whole-object context

Use GetSelf when the rule compares multiple fields in the same object. Use property-level When to keep branch-specific rules isolated.

<a id="validate-using-the-whole-current-object"></a>

**Validate using the whole current object.**

[//]: # (embed: ExampleGetSelf)

```go
// If you want to access the value of the entity you're writing the [govy.Validator] for,
// you can use [govy.GetSelf] function which is a convenience [govy.PropertyGetter] that returns self.
// Note that we don't call [govy.PropertyRules.WithName] here,
// as we're comparing two properties in our top level, [Teacher] scope.
//
// You can provide your own rules using [govy.NewRule] constructor.
// It returns new [govy.Rule] instance which wraps your validation function.
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

	// Output:
	// Validation for Teacher has failed:
	//   - now I have access to the whole teacher
}
```

<a id="run-property-rules-only-when-predicates-match"></a>

**Run property rules only when predicates match.**

[//]: # (embed: ExamplePropertyRules_When)

```go
// To only run property validation on condition, use [govy.PropertyRules.When].
// Predicates set through [govy.PropertyRules.When] are evaluated in the order they are provided.
// If any predicate is not met, validation rules are not evaluated for the whole [govy.PropertyRules].
//
// It's recommended to define [govy.PropertyRules.When] before [govy.PropertyRules.Rules] declaration.
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

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jerry':
	//     - must not be equal to 'Jerry'
}
```

## Control property-level aggregation

Property cascade overrides validator cascade for that property and is useful when later rules add noise after an earlier failure.

<a id="stop-evaluating-later-rules-on-one-property"></a>

**Stop evaluating later rules on one property.**

[//]: # (embed: ExamplePropertyRules_Cascade)

```go
// To customize how [govy.Rule] are evaluated use [govy.PropertyRules.Cascade].
// Use [govy.CascadeModeStop] to stop validation after the first error.
// If you wish to revert to the default behavior, use [govy.CascadeModeContinue].
//
// Note: the cascade mode change only applies to the given [govy.PropertyRules] instance
// and not the parent [govy.Validator] or neighboring [govy.PropertyRules].
// It does however override the [govy.CascadeMode] set for [govy.Validator].
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

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - always fails
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jerry':
	//     - must not be equal to 'Jerry'
}
```
