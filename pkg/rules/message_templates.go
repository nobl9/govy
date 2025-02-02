package rules

import (
	"bytes"
	"log/slog"
	"text/template"

	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/pkg/govy"
)

func mustExecuteTemplate(tpl *template.Template, vars govy.TemplateVars) string {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, vars); err != nil {
		logging.Logger().Error("failed to execute message template",
			slog.String("template", tpl.Name()),
			slog.String("error", err.Error()))
	}
	return buf.String()
}
