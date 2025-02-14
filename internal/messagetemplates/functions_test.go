package messagetemplates

import (
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"testing"
	"text/template"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/assert"
)

type templateVariables struct {
	Examples  []string
	Values    any
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
		"joinSlice": {
			{
				Text:     `{{ joinSlice .Examples "'" }}`,
				Vars:     templateVariables{Examples: nil},
				Expected: "",
			},
			{
				Text:     `{{ joinSlice .Examples "'" }}`,
				Vars:     templateVariables{Examples: []string{}},
				Expected: "",
			},
			{
				Text:     `{{ joinSlice .Examples "'" }}`,
				Vars:     templateVariables{Examples: []string{"foo"}},
				Expected: "'foo'",
			},
			{
				Text:     `{{ joinSlice .Examples "'" }}`,
				Vars:     templateVariables{Examples: []string{"foo", "bar", "baz"}},
				Expected: "'foo', 'bar', 'baz'",
			},
			{
				Text:     `{{ joinSlice .Values "'" }}`,
				Vars:     templateVariables{Values: []int{1, 2, 3}},
				Expected: "'1', '2', '3'",
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

func TestFunctions_EnsureExamplesAreDefined(t *testing.T) {
	root := internal.FindModuleRoot()
	path := filepath.Join(root, "pkg", "govy", "example_test.go")

	data, err := os.ReadFile(path)
	assert.Require(t, assert.NoError(t, err))

	re := regexp.MustCompile(`(?m)^func ExampleAddTemplateFunctions_(\w+)\(\)`)
	matches := re.FindAllStringSubmatch(string(data), -1)

	for funcName := range templateFunctions {
		if !slices.ContainsFunc(matches, func(match []string) bool { return match[1] == funcName }) {
			assert.Fail(t,
				"Example for template function %[1]q is missing"+
					", expected the following signature: 'func ExampleAddTemplateFunctions_%[1]s()'",
				funcName)
		}
	}
}
