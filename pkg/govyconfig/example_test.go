package govyconfig_test

import (
	"fmt"

	"github.com/nobl9/govy/pkg/govy"
	"github.com/nobl9/govy/pkg/govyconfig"
	"github.com/nobl9/govy/pkg/rules"
)

type Teacher struct {
	Name string `json:"name"`
}

// In the interactive tutorial for govy, we've been using
// [govy.PropertyRules.WithName] to provide the name for our properties.
//
// Ideally, we'd want to make sure the names govy assigns to each property,
// match the name of the real-world struct representation that the user interacts with.
// Go uses struct tags to achieve just that,
// and libraries like [encoding/json] use their values to encode/decode structs.
// Unfortunately, there's no easy way to tell what exact property we're returning from [govy.PropertyGetter].
//
// To solve this problem, govy provides a way to infer the name of the property (with a catch).
// The catch being that the name inference mechanism needs to parse the whole modules' AST.
// This can be a performance hit, especially for large projects if not done properly.
//
// By default govy WILL NOT attempt to infer ANY property names.
//
// So, how do we do that properly?
// It depends on the [govyconfig.NameInferMode] used:
//   - [govyconfig.NameInferModeDisable], name inference is disabled (default), nothing to do here
//   - [govyconfig.NameInferModeRuntime], the name is inferred during runtime, whenever [govy.For] is called.
//     This is the most flexible option, but also the slowest.
//     However, If you make sure that [govy.PropertyRules] are created only once and don't mind
//     the initial/startup performance hit, this should be enough for you.
//   - [govyconfig.NameInferModeGenerate], the name is inferred during code generation.
//     This mode requires you to run the cmd/nameinfer BEFORE you run your code.
//     It will generate a file with inferred names for your structs which automatically
//     registers these names using [govyconfig.SetInferredName].
//
// Since this tutorial is run as a test,
// we need to explicitly instruct govy to infer names from test files.
// In order to do that, we use [govyconfig.SetNameInferIncludeTestFiles].
func ExampleSetNameInferMode() {
	govyconfig.SetNameInferIncludeTestFiles(true)
	govyconfig.SetNameInferMode(govyconfig.NameInferModeRuntime)
	defer govyconfig.SetNameInferIncludeTestFiles(false)
	defer govyconfig.SetNameInferMode(govyconfig.NameInferModeDisable)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).WithName("Teacher")

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

// In the previous example we've seen [govyconfig.NameInferModeRuntime] in action.
// An alternative for the aforementioned mode which offers better runtime performance
// is [govyconfig.NameInferModeGenerate].
//
// It comes at a cost of having to run the code generation utility before running your code.
// The utility generates code which uses [govyconfig.SetInferredName].
// We'll use this very function in this example to simulate the code generation step.
// The first validator, 'v1', is created with [govyconfig.NameInferModeDisable],
// the second validator, 'v2' is created with [govyconfig.NameInferModeGenerate].
// As you can see in the output, only the second validator, 'v2' has the inferred name.
func ExampleNameInferModeGenerate() {
	govyconfig.SetNameInferIncludeTestFiles(true)
	defer govyconfig.SetNameInferIncludeTestFiles(false)
	defer govyconfig.SetNameInferMode(govyconfig.NameInferModeDisable)

	v1 := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).WithName("Teacher")

	govyconfig.SetNameInferMode(govyconfig.NameInferModeGenerate)
	govyconfig.SetInferredName(govyconfig.InferredName{
		Name: "name",
		File: "pkg/govyconfig/example_test.go",
		Line: 96,
	})

	v2 := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Thomas")),
	).WithName("NotTeacher")

	teacher := Teacher{Name: "Tom"}
	if err := v1.Validate(teacher); err != nil {
		fmt.Println(err)
	}
	if err := v2.Validate(teacher); err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - should be equal to 'Jerry'
	// Validation for NotTeacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - should be equal to 'Thomas'
}

// Beware where you call [govyconfig.SetNameInferMode].
// If you call it after the [govy.For] has been called, it won't do anything.
// This is because the name inference is done during the creation of [govy.PropertyRules].
func ExampleSetNameInferMode_invalidUsage() {
	govyconfig.SetNameInferIncludeTestFiles(true)
	defer govyconfig.SetNameInferIncludeTestFiles(false)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).WithName("Teacher")

	govyconfig.SetNameInferMode(govyconfig.NameInferModeRuntime)
	defer govyconfig.SetNameInferMode(govyconfig.NameInferModeDisable)

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - should be equal to 'Jerry'
}
