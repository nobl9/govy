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

	messageTemplatesCache = newMessageTemplatesMap()
	assert.Len(t, messageTemplatesCache.tmpl, 0)

	tpl := Get(LengthTemplate)
	actual := mustExecuteTemplate(t, tpl, vars)
	assert.Equal(t, expected, actual)
	assert.Len(t, messageTemplatesCache.tmpl, 1)

	tpl = Get(LengthTemplate)
	actual = mustExecuteTemplate(t, tpl, vars)
	assert.Equal(t, expected, actual)
	assert.Len(t, messageTemplatesCache.tmpl, 1)
}

func TestGet_TemplateWithDependencies(t *testing.T) {
	messageTemplatesCache = newMessageTemplatesMap()
	assert.Len(t, messageTemplatesCache.tmpl, 0)

	tpl := Get(StringKubernetesQualifiedNameTemplate)
	actual := mustExecuteTemplate(t, tpl, map[string]any{"Custom": map[string]any{"EmptyPrefixPart": true}})
	assert.Equal(t, "prefix part must not be empty", actual)
	assert.Len(t, messageTemplatesCache.tmpl, 1)

	tpl = Get(StringKubernetesQualifiedNameTemplate)
	actual = mustExecuteTemplate(t, tpl, map[string]any{
		"ComparisonValue": "^foo$",
		"Custom":          map[string]any{"PrefixRegexp": true},
	})
	assert.Equal(t, "prefix part string must match regular expression: '^foo$'", actual)
	assert.Len(t, messageTemplatesCache.tmpl, 1)

	tpl = Get(StringMatchRegexpTemplate)
	actual = mustExecuteTemplate(t, tpl, map[string]any{"ComparisonValue": "^foo$"})
	assert.Equal(t, "string must match regular expression: '^foo$'", actual)
	assert.Len(t, messageTemplatesCache.tmpl, 2)
}

func mustExecuteTemplate(t *testing.T, tpl *template.Template, vars any) string {
	t.Helper()
	var buf bytes.Buffer
	err := tpl.Execute(&buf, vars)
	assert.Require(t, assert.NoError(t, err))
	return buf.String()
}
