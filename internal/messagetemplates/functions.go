package messagetemplates

import (
	"reflect"
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
	"formatExamples": formatExamplesTplFunc,
	"joinSlice":      joinSliceTplFunc,
	"indent":         indentTplFunc,
}

// formatExamplesTplFunc formats a list of strings which are example valid values
// as a single string representation.
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

// joinSliceTplFunc joins a list of values into a comma separated list of strings.
// Its second argument determines the surrounding string for each value.
// Example: `{{ joinSlice ["foo", "bar"] "'" }}` -> "'foo', 'bar'"
func joinSliceTplFunc(input any, surroundingStr string) string {
	rv := reflect.ValueOf(input)
	if rv.Kind() != reflect.Slice {
		panic("first argument must be a slice")
	}
	if rv.Len() == 0 {
		return ""
	}

	values := make([]any, 0, rv.Len())
	for i := range rv.Len() {
		values = append(values, rv.Index(i).Interface())
	}
	b := strings.Builder{}
	internal.PrettyStringListBuilder(&b, values, surroundingStr)
	return b.String()
}

// indentTplFunc indents every line in a given string to the specified indent width.
// Example: `{{ indent 2 "foo\nbar" }}` -> "  foo\n  bar"
func indentTplFunc(spaces int, v string) string {
	pad := strings.Repeat(" ", spaces)
	return pad + strings.ReplaceAll(v, "\n", "\n"+pad)
}
