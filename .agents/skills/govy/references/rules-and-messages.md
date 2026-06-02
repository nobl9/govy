# Rules and Messages

Custom rules, details, examples, error codes, messages, message templates, template functions, pointer rules, and rule sets.

## Examples

- [Create a custom rule](#examplerule)
- [Attach static rule details](#examplerule_withdetails)
- [Attach formatted rule details](#examplerule_withdetailsf)
- [Attach rule examples](#examplerule_withexamples)
- [Attach rule error codes](#examplerule_witherrorcode)
- [Override rule messages](#examplerule_withmessage)
- [Override rule messages dynamically](#examplerule_withmessagef)
- [Use string message templates](#examplerule_withmessagetemplatestring)
- [Use parsed message templates](#examplerule_withmessagetemplate)
- [Add custom template functions](#exampleaddtemplatefunctions)
- [Format example lists in templates](#exampleaddtemplatefunctions_formatexamples)
- [Join slices in templates](#exampleaddtemplatefunctions_joinslice)
- [Indent template output](#exampleaddtemplatefunctions_indent)
- [Describe rules for validation plans](#examplerule_withdescription)
- [Convert rules to pointer rules](#exampleruletopointer)
- [Group rules into reusable sets](#exampleruleset)
- [Convert rule sets to pointer rule sets](#examplerulesettopointer)
- [Cascade failures within rule sets](#exampleruleset_cascade)

## ExampleRule

Govy comes with a set of predefined rules,
which you can use out of the box by importing [rules] package.

However, you can also create your own rules by using [govy.NewRule] constructor.
It accepts a simple validation function which takes in a value
and returns an error if the validation failed.

Note: the [govy.Rule] struct has all its fields private,
so you can only create and modify them using exported constructor and methods.

[//]: # (embed: ExampleRule)

```go
func ExampleRule() {
	myRule := govy.NewRule(func(name string) error {
		if name != "Tom" {
			return fmt.Errorf("Teacher can be only Tom")
		}
		return nil
	})
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(myRule),
	)

	teacher := Teacher{Name: "Jake"}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'name' with value 'Jake':
	//     - Teacher can be only Tom
}
```

## ExampleRule_WithDetails

You can use [govy.Rule.WithDetails] to add additional details to the error message.
This allows you to extend existing rules by adding your use case context.
Let's give a regex validation some more clarity.

[//]: # (embed: ExampleRule_WithDetails)

```go
func ExampleRule_WithDetails() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)")),
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
	//     - string must match regular expression: '^(Tom|Jerry)$'; Teacher can be either Tom or Jerry :)
}
```

## ExampleRule_WithDetailsf

You can use [govy.Rule.WithDetailsf] to add formatted details to the returned [govy.RuleError] error message.

[//]: # (embed: ExampleRule_WithDetailsf)

```go
func ExampleRule_WithDetailsf() {
	minLen := 3
	maxLen := 10
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringLength(minLen, maxLen).
				WithDetailsf("Teacher name must be between %d and %d characters", minLen, maxLen)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jo",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jo':
	//     - length must be between 3 and 10; Teacher name must be between 3 and 10 characters
}
```

## ExampleRule_WithExamples

You can use [govy.Rule.WithExamples] to add examples of valid inputs
which pass the [govy.Rule].
This can be useful for more complex rules, especially regex based, where
it might not be immediately obvious how a valid value should look like.

Note: examples are added between the error message and details
(configured with [govy.Rule.WithDetails]).

[//]: # (embed: ExampleRule_WithExamples)

```go
func ExampleRule_WithExamples() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)").
				WithExamples("Tom", "Jerry")),
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
	//     - string must match regular expression: '^(Tom|Jerry)$' (e.g. 'Tom', 'Jerry'); Teacher can be either Tom or Jerry :)
}
```

## ExampleRule_WithErrorCode

When testing, it can be tedious to always rely on error messages as these can change over time.
Enter [govy.ErrorCode], which is a simple string type alias used to ease testing,
but also potentially allow third parties to integrate with your validation results.
Use [govy.Rule.WithErrorCode] to associate [govy.ErrorCode] with a [govy.Rule].
Notice that our modified version of [rules.StringMatchRegexp] will now return a new [govy.ErrorCode].
Predefined rules have [govy.ErrorCode] already associated with them.
To view the list of predefined [govy.ErrorCode] checkout error_codes.go file.

[//]: # (embed: ExampleRule_WithErrorCode)

```go
func ExampleRule_WithErrorCode() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)").
				WithErrorCode("custom_code")),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jake",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		propertyErrors := err.(*govy.ValidatorError).Errors
		ruleErrors := propertyErrors[0].Errors
		fmt.Println(ruleErrors[0].Code)
	}

	// Output:
	// custom_code
}
```

## ExampleRule_WithMessage

If you want to override the default error message, you can use [govy.Rule.WithMessage].

[//]: # (embed: ExampleRule_WithMessage)

```go
func ExampleRule_WithMessage() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)").
				WithMessage("unsupported name")),
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
	//     - unsupported name; Teacher can be either Tom or Jerry :)
}
```

## ExampleRule_WithMessagef

You can use [govy.Rule.WithMessagef] to override the default error message using printf-like formatting.

[//]: # (embed: ExampleRule_WithMessagef)

```go
func ExampleRule_WithMessagef() {
	allowedNames := []string{"Tom", "Jerry"}
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithMessagef("name must be one of: %v", allowedNames)),
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
	//     - name must be one of: [Tom Jerry]
}
```

## ExampleRule_WithMessageTemplateString

If you want to have more control over the resulting error message, but [govy.Rule.WithMessage]
is not enough, you can utilize a template string which is parsed by [govy.Rule] into
[template.Template] to construct a custom error message.

Each builtin rule supports different variables.
For instance, [rules.StringLength] supports 'MinLength' and 'MaxLength' variables.
Refer to the rule's documentation to see which variables are supported.

Note: Builtin functions provided by [govy.AddTemplateFunctions], like 'formatExamples',
are automatically added to the parsed [template.Template].

[//]: # (embed: ExampleRule_WithMessageTemplateString)

```go
func ExampleRule_WithMessageTemplateString() {
	tplString := `Teacher's name must be between {{ .MinLength }} and {{ .MaxLength }} characters {{ formatExamples .Examples }}.`

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringLength(5, 10).
				WithExamples("Joanna", "Angeline").
				WithMessageTemplateString(tplString)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Eve",
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Eve':
	//     - Teacher's name must be between 5 and 10 characters (e.g. 'Joanna', 'Angeline').
}
```

## ExampleRule_WithMessageTemplate

If you want to have more control over the [template.Template] used for error message creation,
for instance, add custom functions, use [govy.Rule.WithMessageTemplate].

In the example below, we're defining a custom template function 'join' which calls [strings.Join]
under the hood to join a slice of strings with a comma.

Note: 'Examples' field is a plain slice of strings, If you wish to format it the same way
as the default message does, use 'formatExamples' function provided by [govy.AddTemplateFunctions].

[//]: # (embed: ExampleRule_WithMessageTemplate)

```go
func ExampleRule_WithMessageTemplate() {
	tplString := `Teacher's name '{{ .PropertyValue }}' is not supported. {{ .Details }} (e.g. {{ join .Examples ", " }}).`
	tpl := template.New("").Funcs(template.FuncMap{"join": strings.Join})
	tpl = template.Must(tpl.Parse(tplString))

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringLength(5, 10).
				WithDetails("Teacher's name must be between 5 and 10 characters").
				WithExamples("Joanna", "Angeline").
				WithMessageTemplate(tpl)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Eve",
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Eve':
	//     - Teacher's name 'Eve' is not supported. Teacher's name must be between 5 and 10 characters (e.g. Joanna, Angeline).
}
```

## ExampleAddTemplateFunctions

Under the hood builtin rules' message templates utilize a set of custom template functions.
If you want to use them in your custom templates, you can add them to your [template.Template]
instance by calling [govy.AddTemplateFunctions].

An example of such function is 'formatExamples' which takes in a slice of strings
and returns a formatted string.

Note: Builtin functions are automatically added to the parsed [template.Template] if you're using
[govy.Rule.WithMessageTemplateString].

Note: [govy.AddTemplateFunctions] calls [template.Template.Funcs], which will not add the functions
to your template If it was already parsed.

[//]: # (embed: ExampleAddTemplateFunctions)

```go
func ExampleAddTemplateFunctions() {
	tplString := `Teacher's name '{{ .PropertyValue }}' is not supported {{ formatExamples .Examples }}.`
	tpl := template.New("")
	tpl = govy.AddTemplateFunctions(tpl)
	tpl = template.Must(tpl.Parse(tplString))

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringLength(5, 10).
				WithExamples("Joanna", "Angeline").
				WithMessageTemplate(tpl)),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Eve",
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Eve':
	//     - Teacher's name 'Eve' is not supported (e.g. 'Joanna', 'Angeline').
}
```

## ExampleAddTemplateFunctions_formatExamples

[//]: # (embed: ExampleAddTemplateFunctions_formatExamples)

```go
func ExampleAddTemplateFunctions_formatExamples() {
	tplString := "{{ formatExamples .Examples }}"
	tpl := template.New("")
	tpl = govy.AddTemplateFunctions(tpl)
	tpl = template.Must(tpl.Parse(tplString))

	err := tpl.Execute(
		os.Stdout,
		map[string]any{"Examples": []string{"Joanna", "Angeline"}},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// (e.g. 'Joanna', 'Angeline')
}
```

## ExampleAddTemplateFunctions_joinSlice

[//]: # (embed: ExampleAddTemplateFunctions_joinSlice)

```go
func ExampleAddTemplateFunctions_joinSlice() {
	tplString := `{{ joinSlice .Slice "'" }}`
	tpl := template.New("")
	tpl = govy.AddTemplateFunctions(tpl)
	tpl = template.Must(tpl.Parse(tplString))

	err := tpl.Execute(
		os.Stdout,
		map[string]any{"Slice": []string{"Joanna", "Angeline"}},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// 'Joanna', 'Angeline'
}
```

## ExampleAddTemplateFunctions_indent

[//]: # (embed: ExampleAddTemplateFunctions_indent)

```go
func ExampleAddTemplateFunctions_indent() {
	tplString := "{{ indent 2 .Details }}"
	tpl := template.New("")
	tpl = govy.AddTemplateFunctions(tpl)
	tpl = template.Must(tpl.Parse(tplString))

	err := tpl.Execute(
		os.Stdout,
		map[string]any{"Details": "foo\nbar"},
	)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	//   foo
	//   bar
}
```

## ExampleRule_WithDescription

[govy.Rule] error might be static, i.e. a single [govy.Rule] always returns
the same exact error message, but they don't have to.
For instance, consider a rule which parses a URL using [net/url] package.

This makes it very hard to infer error message for [govy.RulePlan], if not
impossible, since the exact error might only be known during runtime.

To solve this problem, you can use [govy.Rule.WithDescription] to provide a
verbose and informative rule description.
It will be only included in the [govy.RulePlan] and otherwise not displayed in
the default [govy.RuleError.Error].
However, it is available in the structured [govy.RuleError].

[//]: # (embed: ExampleRule_WithDescription)

```go
func ExampleRule_WithDescription() {
	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
				WithDetails("Teacher can be either Tom or Jerry :)").
				WithMessage("unsupported name")),
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
	//     - unsupported name; Teacher can be either Tom or Jerry :)
}
```

## ExampleRuleToPointer

The builtin rules, and most likely your custom rules as well, all operate on non-pointer values.
This means you cannot use them on pointers to the same type.

If for whatever reason you don't want to use [govy.ForPointer] constructor,
you can use [govy.RuleToPointer] constructor and convert any [govy.Rule] to work on pointers.

Note: [govy.RuleToPointer] will skip validation for nil pointers.
If you want to enforce the value to be non-nil, you can use [rules.Required].
This behavior is consistent with [govy.ForPointer] constructor, which will skip the validation
unless you add [govy.PropertyRules.Required] to enforce the value to be a non-nil pointer.

[//]: # (embed: ExampleRuleToPointer)

```go
func ExampleRuleToPointer() {
	type Pointer struct {
		Pointer *string `json:"pointer"`
	}
	validator := govy.New(
		govy.For(func(p Pointer) *string { return p.Pointer }).
			WithName("pointer").
			Rules(govy.RuleToPointer(rules.EQ("foo"))),
	)

	pointer := Pointer{Pointer: ptr("bar")}

	err := validator.Validate(pointer)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'pointer' with value 'bar':
	//     - must be equal to 'foo'
}
```

## ExampleRuleSet

Sometimes it's useful to aggregate multiple [govy.Rule] into a single, composite rule.
To do that we'll use [govy.RuleSet] and [govy.NewRuleSet] constructor.
RuleSet is a simple container for multiple [govy.Rule].
During validation it is unpacked and each [govy.RuleError] is reported separately.

Note that govy uses similar syntax to wrapped errors in Go;
a ':' delimiter is used to chain error codes together.

[//]: # (embed: ExampleRuleSet)

```go
func ExampleRuleSet() {
	teacherNameRule := govy.NewRuleSet(
		rules.StringLength(1, 5),
		rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")).
			WithDetails("Teacher can be either Tom or Jerry :)"),
	).
		WithErrorCode("teacher_name")

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(teacherNameRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jonathan",
		Age:  51 * year,
	}

	err := v.Validate(teacher)
	if err != nil {
		propertyErrors := err.(*govy.ValidatorError).Errors
		ruleErrors := propertyErrors[0].Errors
		fmt.Printf("Error codes: %s, %s\n\n", ruleErrors[0].Code, ruleErrors[1].Code)
		fmt.Println(err)
	}

	// Output:
	// Error codes: teacher_name:string_length, teacher_name:string_match_regexp
	//
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jonathan':
	//     - length must be between 1 and 5
	//     - string must match regular expression: '^(Tom|Jerry)$'; Teacher can be either Tom or Jerry :)
}
```

## ExampleRuleSetToPointer

Similar to [govy.RuleToPointer], you can use [govy.RuleSetToPointer] to convert
[govy.RuleSet] to work with pointers.

See [ExampleRuleToPointer] for more details.

[//]: # (embed: ExampleRuleSetToPointer)

```go
func ExampleRuleSetToPointer() {
	type Pointer struct {
		Pointer *string `json:"pointer"`
	}
	ruleSet := govy.NewRuleSet(
		rules.StringStartsWith("f"),
		rules.StringEndsWith("o"),
	)
	validator := govy.New(
		govy.For(func(p Pointer) *string { return p.Pointer }).
			WithName("pointer").
			Rules(govy.RuleSetToPointer(ruleSet)),
	)

	pointer := Pointer{Pointer: ptr("bar")}

	err := validator.Validate(pointer)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation has failed for the following properties:
	//   - 'pointer' with value 'bar':
	//     - string must start with 'f' prefix
	//     - string must end with 'o' suffix
}
```

## ExampleRuleSet_Cascade

If you wish to control how rules aggregated by [govy.RuleSet] evaluate
you can use [govy.RuleSet.Cascade] to set a [govy.CascadeMode].

Similar to how the cascade mode works when evaluating [govy.PropertyRules],
the [govy.CascadeModeStop] will stop validation after the first encountered error.

In the example below we can see that although both rules should fail,
only the first one (order of definitions matters here!) returns an error.

[//]: # (embed: ExampleRuleSet_Cascade)

```go
func ExampleRuleSet_Cascade() {
	teacherNameRule := govy.NewRuleSet(
		rules.StringLength(1, 5),
		rules.StringMatchRegexp(regexp.MustCompile("^(Tom|Jerry)$")),
	).
		Cascade(govy.CascadeModeStop)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(teacherNameRule),
	).WithName("Teacher")

	teacher := Teacher{
		Name: "Jonathan",
	}

	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Jonathan':
	//     - length must be between 1 and 5
}
```

