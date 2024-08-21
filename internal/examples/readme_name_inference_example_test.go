package examples

import (
	"fmt"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govyconfig"
	"github.com/nobl9/govy/pkg/rules"
)

func Example_nameInference() {
	govyconfig.SetNameInferIncludeTestFiles(true) // Required for the example to run.
	govyconfig.SetNameInferMode(govyconfig.NameInferModeRuntime)
	defer govyconfig.SetNameInferIncludeTestFiles(false)
	defer govyconfig.SetNameInferMode(govyconfig.NameInferModeDisable)

	type Teacher struct {
		Name string `json:"name"`
	}

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).InferName()

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - should be equal to 'Jerry'
}
