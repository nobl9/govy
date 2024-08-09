# govy

Validation library for Go that uses a functional interface for building
strongly-typed validation rules, powered by generics and
[reflection free](#reflection).

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
to access it visit [pkg.go.dev](TODO) or locally at [example_test.go](./example_test.go).

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
