---
name: govy
description: >
  Use this skill when writing, reviewing, or explaining validation code with
  github.com/nobl9/govy. Trigger for govy validators, property rules, custom
  rules, predefined rules, govytest assertions, validation plans, error
  messages, message templates, path inference, or any task that mentions the
  govy library.
---

# Govy

Govy is a validation library for Go that uses a functional interface for building
strongly-typed validation rules, powered by generics and reflection free.
It puts heavy focus on end user errors readability,
providing means of crafting clear and information-rich error messages.
It also allows writing self-documenting validation rules through a
validation plan.

The references in this skill are self-contained.
Load only the reference that matches the task.

## Reference selection

| Task | Load |
| :--- | :--- |
| Build a validator, name it, add conditions, validate slices, or compose validators | [core-validation.md](references/core-validation.md) |
| Work with property getters, paths, pointers, transforms, required values, optional values, hidden values, or property-level conditions | [properties-and-paths.md](references/properties-and-paths.md) |
| Create or modify rules, rule sets, error details, examples, error codes, custom messages, or message templates | [rules-and-messages.md](references/rules-and-messages.md) |
| Validate nested objects, slices, maps, each element, or compose validators over collections | [collections-and-composition.md](references/collections-and-composition.md) |
| Inspect, serialize, rename, or construct govy errors | [errors.md](references/errors.md) |
| Test govy validation behavior without the govytest package | [testing.md](references/testing.md) |
| Generate or validate a validation plan | [validation-plan.md](references/validation-plan.md) |
| Configure runtime or generated path inference | [path-inference.md](references/path-inference.md) |
| Choose from existing predefined rules in `pkg/rules` | [existing-rules.md](references/existing-rules.md) |
| See examples of predefined rules in use | [predefined-rules.md](references/predefined-rules.md) |
| Use `pkg/govytest` assertion helpers | [govytest.md](references/govytest.md) |

## Usage guidance

- Define rules and validators once and reuse them across validation calls.
  They are relatively costly to construct and may lazily initialize internal
  state during first use, so avoid rebuilding them every time validation runs.
- Keep validator declarations immutable-style: chained methods return copies,
  so derive runtime variants from existing validators instead of mutating them
  in place.
- Prefer explicit `WithName` or `WithPath` unless the user explicitly
  mentions path inference or if the inference mechanism is already used.
- Use `WithName` for one path segment and `WithPath` for multi-segment paths.
- Use `ForPointer` for optional pointer fields so rules operate on the
  pointed-to value and nil values are skipped unless `Required` is added.
- Use `Required` and `OmitEmpty` on property rules to short-circuit empty values
  before running normal rules.
- Attach plan and error metadata to the receiver that owns it:
  use `Rule.WithDescription`, `Rule.WithDetails`, `Rule.WithExamples`,
  and `Rule.WithErrorCode` for rules; use `PropertyRules.WithExamples`
  for property examples; use `govy.WhenDescription` for conditional rules.
- Use `govytest` helpers for tests that need to match structured validation
  errors rather than brittle full error strings.

## General example

[//]: # (embed: internal/examples/readme_intro_example_test.go)

```go
package examples

import (
	"fmt"
	"regexp"
	"time"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func Example_basicUsage() {
	type University struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	type Student struct {
		Index string `json:"index"`
	}
	type Teacher struct {
		Name       string        `json:"name"`
		Age        time.Duration `json:"age"`
		Students   []Student     `json:"students"`
		MiddleName *string       `json:"middleName,omitempty"`
		University University    `json:"university"`
	}

	universityValidation := govy.New(
		govy.For(func(u University) string { return u.Name }).
			WithName("name").
			Required(),
		govy.For(func(u University) string { return u.Address }).
			WithName("address").
			Required().
			Rules(rules.StringMatchRegexp(
				regexp.MustCompile(`[\w\s.]+, \d{2}-\d{3} \w+`),
			).
				WithDetails("Polish address format must consist of the main address and zip code").
				WithExamples("5 M. Skłodowska-Curie Square, 60-965 Poznan")),
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
		govy.ForPointer(func(t Teacher) *string { return t.MiddleName }).
			WithName("middleName").
			Rules(rules.StringTitle()),
		govy.ForSlice(func(t Teacher) []Student { return t.Students }).
			WithName("students").
			Rules(
				rules.SliceMaxLength[[]Student](2),
				rules.SliceUnique(func(v Student) string { return v.Index })).
			IncludeForEach(studentValidator),
		govy.For(func(t Teacher) University { return t.University }).
			WithName("university").
			Include(universityValidation),
	).
		When(func(t Teacher) bool { return t.Age < 50 })

	teacher := Teacher{
		Name:       "John",
		MiddleName: nil, // Validation for nil pointers by default is skipped.
		Age:        48,
		Students: []Student{
			{Index: "918230014"},
			{Index: "9182300123"},
			{Index: "918230014"},
		},
		University: University{
			Name:    "",
			Address: "10th University St.",
		},
	}

	if err := teacherValidator.WithName("John").Validate(teacher); err != nil {
		fmt.Println(err)
	}
	// When condition is not met, no validation errors.
	johnFromTheFuture := teacher
	johnFromTheFuture.Age = 51
	if err := teacherValidator.WithName("John From The Future").Validate(johnFromTheFuture); err != nil {
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
	//   - 'university.name':
	//     - property is required but was empty
	//   - 'university.address' with value '10th University St.':
	//     - string must match regular expression: '[\w\s.]+, \d{2}-\d{3} \w+' (e.g. '5 M. Skłodowska-Curie Square, 60-965 Poznan'); Polish address format must consist of the main address and zip code
}
```
