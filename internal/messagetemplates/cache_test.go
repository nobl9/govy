package messagetemplates

import (
	"bytes"
	"testing"
	"text/template"

	"github.com/nobl9/govy/internal/assert"
)

func TestGet(t *testing.T) {
	expected := "length must be between 1 and 2"
	vars := templateVariables{
		MinLength: 1,
		MaxLength: 2,
	}

	assert.Len(t, messageTemplatesCache.tmpl, 0)
	tpl := Get(StringLengthTemplate)
	actual := mustExecuteTemplate(t, tpl, vars)
	assert.Equal(t, expected, actual)

	assert.Len(t, messageTemplatesCache.tmpl, 1)
	tpl = Get(StringLengthTemplate)
	actual = mustExecuteTemplate(t, tpl, vars)
	assert.Equal(t, expected, actual)
}

func mustExecuteTemplate(t *testing.T, tpl *template.Template, vars any) string {
	t.Helper()
	var buf bytes.Buffer
	err := tpl.Execute(&buf, vars)
	assert.Require(t, assert.NoError(t, err))
	return buf.String()
}
