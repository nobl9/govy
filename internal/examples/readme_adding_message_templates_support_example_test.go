package examples

import (
	"fmt"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func Example_addingMessageTemplatesSupportToCustomRules() {
	type Teacher struct {
		Name string `json:"name"`
	}

	template := `{{ .PropertyValue }} must be {{ .ComparisonValue }}; {{ .Custom.Foo }} and {{ .Custom.Baz }}`

	customRule := govy.NewRule(func(name string) error {
		if name != "John" {
			return govy.NewRuleErrorTemplate(govy.TemplateVars{
				PropertyValue:   name,
				ComparisonValue: "John",
				Custom: map[string]any{
					"Foo": "Bar",
					"Baz": 42,
				},
			})
		}
		return nil
	}).
		WithErrorCode("custom_rule").
		WithMessageTemplateString(template).
		WithDetails("we just don't like anyone but Johns...").
		WithDescription("must be John")

	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				customRule,
				rules.StringStartsWith("J"),
			),
	).WithNameFunc(govy.NameFuncFromTypeName[Teacher]())

	teacher := Teacher{Name: "George"}

	if err := teacherValidator.Validate(teacher); err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'George':
	//     - George must be John; Bar and 42
	//     - string must start with 'J' prefix
}
