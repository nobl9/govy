// Package jsonpath provides utilities for working with [JSONPath] path fragments.
// In govy these paths are usually relative to a validator or parent property,
// for example `name` or `students[0].index`, rather than absolute `$.`-prefixed expressions.
// It implements only a subset of [JSONPath] that govy needs for addressing properties,
// not the full query language from the specification.
//
// It also uses the non-standard `*~` selector for plan paths that refer to map keys.
// That selector is not part of [JSONPath] as defined by RFC 9535; govy uses it with the
// same meaning as the key-selector extension in [JSONPath-Plus].
//
// [JSONPath]: https://www.rfc-editor.org/rfc/rfc9535.html
// [JSONPath-Plus]: https://github.com/JSONPath-Plus/JSONPath
package jsonpath
