package govy

import (
	"text/template"

	"github.com/nobl9/govy/internal/messagetemplates"
)

//go:generate go run ../../internal/cmd/docextractor/main.go

// AddTemplateFunctions adds a set of utility functions to the provided template.
// The following functions are made available for use in the templates:
//
//   - 'formatExamples' formats a list of strings which are example valid values
//     as a single string representation.
//     Example: `{{ formatExamples ["foo", "bar"] }}` -> "(e.g. 'foo', 'bar')"
//
//   - 'joinSlice' joins a list of strings into a comma separated list of values.
//     Its second argument determines a surrounding string for each value.
//     Example: `{{ joinSlice ["foo", "bar"] "'" }}` -> "'foo', 'bar'"
//
// Refer to the testable examples of [AddTemplateFunctions] for more details
// on each builtin function.
func AddTemplateFunctions(tpl *template.Template) *template.Template {
	return messagetemplates.AddFunctions(tpl)
}
