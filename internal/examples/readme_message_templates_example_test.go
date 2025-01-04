package examples

import (
	"fmt"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func Example_messageTemplates() {
	type Teacher struct {
		Name string `json:"name"`
	}

	templateString := "name length should be between {{ .MinLength }} and {{ .MaxLength }} {{ formatExamples .Examples }}"

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Rules(
				rules.StringLength(5, 10).
					WithExamples("Joanna", "Jerry").
					WithMessageTemplateString(templateString),
			),
	).WithName("Teacher")

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - name length should be between 5 and 10 (e.g. 'Joanna', 'Jerry')
}
