# govy

Validation library for Go that uses a functional interface for building
strongly-typed validation rules, powered by generics and
[reflection free](#reflection).
It puts heavy focus on end user errors readability,
providing means of crafting clear and information-rich error messages.
It also allows writing self-documenting validation rules through validation plan generator.

**DISCLAIMER** govy is in active development, while the core API is unlikely
to change, breaking changes may be introduced with new versions until v1
is released.

## Getting started

In order to add the library to your project, run:

```shell
go get github.com/nobl9/govy
```

There's an interactive tutorial available,
powered by Go's [testable examples](https://go.dev/blog/examples),
to access it visit [pkg.go.dev](https://pkg.go.dev/github.com/nobl9/govy)
or locally at [example_test.go](./pkg/govy/example_test.go).
Note that `godoc` does not render testable examples' code comments well,
it's advised to go through the interactive tutorial inside an IDE of choice.

Here's a quick example of how to use `govy`:

[//]: # (embed: internal/examples/readme_intro_example_test.go)

```go
```

## Features

### Type safety

### Immutability

govy components are largely immutable and lazily loaded:

- Immutable, as changing the pipeline through chained functions,
  will return a new pipeline.
  It allows extended reusability of validation components.
- Lazily loaded, as properties are extracted through getter functions,
  which are only called when you call the `Validate` method.
  Functional approach allows validation components to only be called when
  needed.
  You should define your pipeline once and call it
  whenever you validate instances of your entity.

### Verbose error messages

TODO

### Validation plan

TODO

### Properties name inference

TODO

## Rationale

Why was this library created?

Most of the established Go libraries for validation were created
in a pre-generics era. They often use reflection in order to provide
generic validation API, it also allows them to utilize struct tags, which
further minimize the amount of code users need to write.

Unfortunately, the ease of use compromises Go's core language feature,
type safety.

Enter, generics.
With generics on board, it's finally possible to write a robust and type safe
API for validation, thus `govy` was born.

### Reflection

Is `govy` truly reflection free?
The short answer is yes, the long answer is `govy` does not utilize
reflection other than to serve better error messages or devise a validation
plan. Some builtin rules might also use `reflect`, but the core functionality
does not rely on it.

## Acknowledgments

The API along with the accompanying nomenclature was heavily inspired by the
awesome [Fluent Validation](https://github.com/FluentValidation/FluentValidation)
library for C#.
