package govy

import (
	"text/template"

	"github.com/nobl9/govy/internal/messagetemplates"
)

// AddTemplateFunctions adds a set of utility functions to the provided template.
func AddTemplateFunctions(tpl *template.Template) *template.Template {
	return messagetemplates.AddFunctions(tpl)
}
