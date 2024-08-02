# Validation

Package validation implements a functional API for consistent,
type safe validation.
It puts heavy focus on end user errors readability,
providing means of crafting clear and information-rich error messages.

Validation pipeline is immutable and lazily loaded.

- Immutable, as changing the pipeline through chained functions,
  will return a new pipeline.
  It allows extended reusability of validation components.
- Lazily loaded, as properties are extracted through getter functions,
  which are only called when you call the `Validate` method.
  Functional approach allows validation components to only be called when
  needed.
  You should define your pipeline once and call it
  whenever you validate instances of your entity.

All that has been made possible by the introduction of generics in Go.
Prior to that, there wasn't really any viable way to create type safe
validation API.
Although the current state of type inference is somewhat clunky,
the API can only improve in time when generics support in Go is further
extended.

## NOTE: Work in progress

Although already battle tested through SLO hellfire,
this library is still a work in progress.
The principles and the API at its core won't change,
but the details and capabilities might hopefully will.
Contributions and suggestions are most welcome!

## Usage

**This README goes through an abstract overview of the library. \
Refer to [example_test.go](./example_test.go)
for a hands-on tutorial with runnable examples.**

### Legend

- [Validator](#validator)
- [Property rules](#property-rules)
  - [PropertyRules](#propertyrules) _(single property)_
  - [PropertyRulesForEach](#propertyrulesforeach) _(slice of properties)_
- [Rule](#rule)
  - [SingleRule](#singlerule)
  - [RuleSet](#ruleset)
- [Errors](#errors)
  - [ValidatorError](#validatorerror)
  - [PropertyError](#propertyerror)
  - [RuleError](#ruleerror)
- [FAQ](#faq)

### Validator

Validator aggregates [property rules](#property-rules) into a single validation scenario,
most commonly associated with an entity, like `struct`.

If any property rules fail [ValidatorError](#validatorerror) is returned.

### Property rules

When validating structured data, namely `struct`,
each structure consists of multiple properties.
For `struct`, these will be its fields.

Most commonly, property has its name and value.
Property name should be derived from the struct
representation visible by the errors consumer,
this will most likely be JSON format.

Nested properties are represented by paths,
where each property is delimited by `.`.
Arrays are represented by `[<index>]`.
Let's examine a simple teacher/student example:

```go
package university

type Teacher struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

type Student struct {
	Index string
}
```

We can distinguish the following property paths:

- `name`
- `students`
- `students[0].Index` _(let's assume there's only a single student)_

If any property rule fails [PropertyError](#propertyerror) is returned.

#### PropertyRules

`PropertyRules` aggregates [rules](#rule) for a single property.

#### PropertyRulesForEach

`PropertyRulesForEach` is an extension of [PropertyRules](#propertyrules),
it provides means of defining rules for each property in a slice.

Currently, it only works with slices, maps are not supported.

### Rule

Rules validate a concrete value.
If a rule is not met it returns [RuleError](#ruleerror).

#### SingleRule

This is the most basic validation building block.
Its error code can be set using `WithErrorCode` function and its error message can
also be enhanced using `WithDetails` function.
Details are delimited by `;` character.

#### RuleSet

Rule sets are used to aggregate multiple [SingleRule](#singlerule)
into a single validation rule.
It wraps any and all errors returned from single rules in a container which is later
on unpacked. If you use either `WithErrorCode` or `WithDetails` functions, each error
will be extended with the provided details and error code.

### Errors

Each validation level defines an error which adds further details of what went wrong.

#### ValidatorError

Adds top level entity name, following our teacher example,
it would be simply `teacher`.
Although that once again depends on how your end use perceives this entity.
It wraps multiple [PropertyError](#propertyerror).

#### PropertyError

Adds both property name and value. Property value is converted to a string
representation. It wraps multiple [RuleError](#ruleerror).

#### RuleError

The most basic building block for validation errors, associated with a single
failing [SingleRule](#singlerule).
It conveys an error message and [ErrorCode](#error-codes).

#### Error codes

To aid the process of testing, `ErrorCode` has been introduced along
with a helper functions `WithErrorCode` to associate [Rule](#rule) with an error
code and `AddCode` to associate multiple error codes with a single [Rule](#rule).
Multiple error codes are delimited by `:`,
similar to how wrapped errors are represented in Go.

To check if `ErrorCode` is part if a given validation error, use `HasErrorCode`.

## FAQ

### Why not use existing validation library?

Existing, established solutions are mostly based on struct tags and heavily
utilize reflection.
This leaves type safety an issue to be solved and handled by developers.
For simple use cases, covered by predefined validation functions,
this solutions works well enough.
However when adding custom validation rules,
type casting has to be heavily utilized,
and it becomes increasingly harder to track what exactly is being validated.
Another issue is the readability of the errors,
it's often hard or even impossible to shape the error to the developer liking.

### Acknowledgements

Heavily inspired by [C# FluentValidation](https://docs.fluentvalidation.net/).
