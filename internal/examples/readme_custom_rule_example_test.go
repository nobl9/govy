package examples

import (
	"fmt"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/rules"
)

func Example_customRules() {
	type Teacher struct {
		Name string `json:"name"`
	}

	customRule := govy.NewRule(func(name string) error {
		if name != "John" {
			return fmt.Errorf("must be John")
		}
		return nil
	}).
		WithErrorCode("custom_rule").
		WithDetails("we just don't like anyone but Johns...").
		WithDescription("must be John")

	teacherValidator := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			WithName("name").
			Required().
			Rules(
				customRule,
				rules.StringStartsWith("J")),
	).InferName()

	teacher := Teacher{Name: "George"}

	if err := teacherValidator.Validate(teacher); err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'George':
	//     - must be John; we just don't like anyone but Johns...
	//     - string must start with 'J' prefix
}
