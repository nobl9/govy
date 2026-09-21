# Creating and Modifying Rules

## Topics

- [Create custom rules](#create-custom-rules)
  - [Create a custom rule from a validation function.](#create-a-custom-rule-from-a-validation-function)
- [Add rule context](#add-rule-context)
  - [Attach static details to a rule error.](#attach-static-details-to-a-rule-error)
  - [Attach formatted details when the context is computed.](#attach-formatted-details-when-the-context-is-computed)
  - [Attach valid input examples for complex rules.](#attach-valid-input-examples-for-complex-rules)
  - [Attach stable error codes for tests and integrations.](#attach-stable-error-codes-for-tests-and-integrations)
- [Customize rule messages](#customize-rule-messages)
  - [Override a rule message with a fixed string.](#override-a-rule-message-with-a-fixed-string)
  - [Format a fixed rule message from arguments.](#format-a-fixed-rule-message-from-arguments)
- [Use message templates](#use-message-templates)
  - [Attach a template from a string.](#attach-a-template-from-a-string)
  - [Attach a pre-parsed template.](#attach-a-pre-parsed-template)
- [Add template helper functions](#add-template-helper-functions)
  - [Register the builtin helper functions on a template.](#register-the-builtin-helper-functions-on-a-template)
  - [Format example slices the same way default messages do.](#format-example-slices-the-same-way-default-messages-do)
  - [Join slice values in custom template output.](#join-slice-values-in-custom-template-output)
  - [Indent multiline template output.](#indent-multiline-template-output)
- [Describe rules for generated plans](#describe-rules-for-generated-plans)
  - [Attach a plan-only rule description.](#attach-a-plan-only-rule-description)
- [Reuse rule sets](#reuse-rule-sets)
  - [Convert a rule so it can validate pointer values.](#convert-a-rule-so-it-can-validate-pointer-values)
  - [Group multiple rules into a reusable set.](#group-multiple-rules-into-a-reusable-set)
  - [Convert a whole rule set for pointer validation.](#convert-a-whole-rule-set-for-pointer-validation)
  - [Stop evaluating later rules in a rule set after failure.](#stop-evaluating-later-rules-in-a-rule-set-after-failure)

## Create custom rules

Use custom rules for validation that `pkg/rules` does not cover.
Keep domain-specific rules in the consumer.
Share them when several validators need the same constraint.
Keep property getters focused on extracting values.
Put validation decisions in rules and attach metadata to the rule that owns it.

For dependent checks, stop the rule chain after the prerequisite fails.
For example, check a format before comparing a parsed value.
Use `GetSelf` for constraints that need several fields.
Return structured property errors when those failures must identify individual fields.

### Create a custom rule from a validation function

[//]: # (embed: ExampleRule?comments=false)

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
}
```

## Add rule context

Use details to explain failures and examples to show valid inputs.
Use stable error codes for tests and integrations.

### Attach static details to a rule error

[//]: # (embed: ExampleRule_WithDetails?comments=false)

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
}
```

### Attach formatted details when the context is computed

[//]: # (embed: ExampleRule_WithDetailsf?comments=false)

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
}
```

### Attach valid input examples for complex rules

[//]: # (embed: ExampleRule_WithExamples?comments=false)

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
}
```

### Attach stable error codes for tests and integrations

[//]: # (embed: ExampleRule_WithErrorCode?comments=false)

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
}
```

## Customize rule messages

Use `WithMessage` for a fixed message.
`WithMessagef` formats its arguments when the rule is constructed.
Use `.PropertyValue` in a message template to include the failed value.
Message overrides retain details and examples.

### Override a rule message with a fixed string

[//]: # (embed: ExampleRule_WithMessage?comments=false)

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
}
```

### Format a fixed rule message from arguments

[//]: # (embed: ExampleRule_WithMessagef?comments=false)

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
}
```

## Use message templates

Use message templates to format rule details, examples, or property values.
Use a string template for simple cases.
Parse a template first when it needs reuse or validation.

### Attach a template from a string

[//]: # (embed: ExampleRule_WithMessageTemplateString?comments=false)

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
}
```

### Attach a pre-parsed template

[//]: # (embed: ExampleRule_WithMessageTemplate?comments=false)

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
}
```

## Add template helper functions

Template helpers are registered through AddTemplateFunctions.
Use the built-in helpers to format slices and multiline text.

### Register the builtin helper functions on a template

[//]: # (embed: ExampleAddTemplateFunctions?comments=false)

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
}
```

### Format example slices the same way default messages do

[//]: # (embed: ExampleAddTemplateFunctions_formatExamples?comments=false)

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
}
```

### Join slice values in custom template output

[//]: # (embed: ExampleAddTemplateFunctions_joinSlice?comments=false)

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
}
```

### Indent multiline template output

[//]: # (embed: ExampleAddTemplateFunctions_indent?comments=false)

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
}
```

## Describe rules for generated plans

Describe each custom rule's constraint without referring to one failing input.
Descriptions affect validation plans, not runtime decisions.
When plans include conditions, add `govy.WhenDescription` to each relevant `When`.
Use `govy.PlanStrictMode()` to detect missing condition descriptions.
Check rule descriptions separately because strict mode does not require them.

### Attach a plan-only rule description

Use `WithDescription` to set plan text.
For a tested example, read
[Generate a plan from a validator](validation-plan.md#generate-a-plan-from-a-validator).

## Reuse rule sets

Use rule sets to package related rules and reuse them across properties.
Convert rules or rule sets to pointer variants when validating pointed-to values.

### Convert a rule so it can validate pointer values

[//]: # (embed: ExampleRuleToPointer?comments=false)

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
}
```

### Group multiple rules into a reusable set

[//]: # (embed: ExampleRuleSet?comments=false)

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
}
```

### Convert a whole rule set for pointer validation

[//]: # (embed: ExampleRuleSetToPointer?comments=false)

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
}
```

### Stop evaluating later rules in a rule set after failure

[//]: # (embed: ExampleRuleSet_Cascade?comments=false)

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
}
```
