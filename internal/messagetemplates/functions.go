package messagetemplates

import (
	"strings"
	"text/template"

	"github.com/nobl9/govy/internal"
)

// AddFunctions adds a set of custom functions to the provided template.
// These functions are used by builtin templates.
func AddFunctions(tpl *template.Template) *template.Template {
	return tpl.Funcs(templateFunctions)
}

var templateFunctions = template.FuncMap{
	"formatExamples":   formatExamplesTplFunc,
	"joinStringsSlice": joinStringsSliceTplFunc,
}

// formatExamplesTplFunc formats a list of strings which are example valid values
// as a single string representation.
//
// Example: `{{ formatExamples ["foo", "bar"] }}` -> "(e.g. 'foo', 'bar')"
func formatExamplesTplFunc(examples []string) string {
	if len(examples) == 0 {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("(e.g. ")
	internal.PrettyStringListBuilder(&b, examples, "'")
	b.WriteString(")")
	return b.String()
}

// joinStringsSliceTplFunc joins a list of strings into a comma separated list of values.
// Its second argument determines a surrounding string for each value.
//
// Example: `{{ joinStringsSlice ["foo", "bar"] "'" }}` -> "'foo', 'bar'"
func joinStringsSliceTplFunc(values []string, surroundingStr string) string {
	if len(values) == 0 {
		return ""
	}
	b := strings.Builder{}
	internal.PrettyStringListBuilder(&b, values, surroundingStr)
	return b.String()
}
