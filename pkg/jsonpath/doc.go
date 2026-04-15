// Package jsonpath provides utilities for working with [JSONPath] path fragments.
// In govy these paths are usually relative to a validator or parent property,
// for example `name` or `students[0].index`, rather than absolute `$.`-prefixed expressions.
// It implements only the subset of [JSONPath] that govy needs for addressing properties,
// not the full query language from the specification.
//
// [JSONPath]: https://www.rfc-editor.org/rfc/rfc9535.html
package jsonpath
