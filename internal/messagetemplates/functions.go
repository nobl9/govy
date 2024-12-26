package messagetemplates

import (
	"strings"
	"text/template"

	"github.com/nobl9/govy/internal"
)

// AddFunctions adds a set of custom functions to the provided template.
// These functions are used by builtin templates.
func AddFunctions(t *template.Template) *template.Template {
	return t.Funcs(functions)
}

var functions = template.FuncMap{
	"formatExamples":    formatExamplesTplFunc,
	"formatStringSlice": formatStringSliceTplFunc,
}

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

func formatStringSliceTplFunc(values []string, surroundingStr string) string {
	if len(values) == 0 {
		return ""
	}
	b := strings.Builder{}
	internal.PrettyStringListBuilder(&b, values, surroundingStr)
	return b.String()
}
