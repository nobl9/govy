package messagetemplates

import (
	"testing"
	"text/template"

	"github.com/nobl9/govy/internal/assert"
)

type templateVariables struct {
	Examples  []string
	Details   string
	MinLength int
	MaxLength int
}

func TestAddFunctions(t *testing.T) {
	tests := map[string][]struct {
		Text     string
		Vars     templateVariables
		Expected string
	}{
		"formatExamples": {
			{
				Text:     "{{ formatExamples .Examples }}",
				Vars:     templateVariables{Examples: nil},
				Expected: "",
			},
			{
				Text:     "{{ formatExamples .Examples }}",
				Vars:     templateVariables{Examples: []string{}},
				Expected: "",
			},
			{
				Text:     "{{ formatExamples .Examples }}",
				Vars:     templateVariables{Examples: []string{"foo"}},
				Expected: "(e.g. 'foo')",
			},
			{
				Text:     "{{ formatExamples .Examples }}",
				Vars:     templateVariables{Examples: []string{"foo", "bar", "baz"}},
				Expected: "(e.g. 'foo', 'bar', 'baz')",
			},
		},
		"formatStringSlice": {
			{
				Text:     `{{ formatStringSlice .Examples "'" }}`,
				Vars:     templateVariables{Examples: nil},
				Expected: "",
			},
			{
				Text:     `{{ formatStringSlice .Examples "'" }}`,
				Vars:     templateVariables{Examples: []string{}},
				Expected: "",
			},
			{
				Text:     `{{ formatStringSlice .Examples "'" }}`,
				Vars:     templateVariables{Examples: []string{"foo"}},
				Expected: "'foo'",
			},
			{
				Text:     `{{ formatStringSlice .Examples "'" }}`,
				Vars:     templateVariables{Examples: []string{"foo", "bar", "baz"}},
				Expected: "'foo', 'bar', 'baz'",
			},
		},
	}

	for funcName, testCases := range tests {
		t.Run(funcName, func(t *testing.T) {
			for _, tc := range testCases {
				tpl := AddFunctions(template.New(""))
				tpl, err := tpl.Parse(tc.Text)
				assert.Require(t, assert.NoError(t, err))

				actual := mustExecuteTemplate(t, tpl, tc.Vars)
				assert.Equal(t, tc.Expected, actual)
			}
		})
	}
}
