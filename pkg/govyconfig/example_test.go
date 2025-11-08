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
// Ideally, we'd want to make sure the names govy assigns to each property
// match the name of the struct fields' names that the user defines the validation for.
// Go uses struct tags to achieve that,
// and libraries like [encoding/json] use these tags to encode/decode structs.
// Unfortunately, there's no easy way to tell what exact property we're returning from [govy.PropertyGetter].
//
// To solve this problem, govy provides a way to infer the name of the property (with a catch).
// The catch being that the name inference mechanism needs to parse the whole modules' AST.
// This can be a performance hit, especially for large projects if not done properly.
//
// By default govy **WILL NOT** attempt to infer **ANY** property names.
//
// So, how do we do that properly?
// Both [govy.Validator] and [govy.PropertyRules] (including variants) have a dedicated function -- TODO
//
// It depends on the [govy.NameInferMode] used:
//   - [govy.NameInferModeDisable], name inference is disabled (default), nothing to do here
//   - [govy.NameInferModeRuntime], the name is inferred during runtime, whenever [govy.For] is called.
//     This is the most flexible option, but also the slowest.
//     However, If you make sure that [govy.PropertyRules] are created only once and don't mind
//     the initial/startup performance hit, this should be enough for you.
//   - [govy.NameInferModeGenerate], the name is inferred during code generation.
//     This mode requires you to run the 'cmd/govy nameinfer' BEFORE you run your code.
//     It will generate a file with inferred names for your structs which automatically
//     registers these names using [govy.SetInferredName].
//
// Since this tutorial is run as a test,
// we need to explicitly instruct govy to infer names from test files.
// By default test files are not parsed to improve performance.
// In order to do that, we use [govyconfig.SetNameInferIncludeTestFiles].
func ExampleNameInferMode() {
	govyconfig.SetNameInferIncludeTestFiles(true)
	defer govyconfig.SetNameInferIncludeTestFiles(false)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		NameInferMode(govy.NameInferModeRuntime).
		WithName("Teacher")

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

// In the previous example we've seen [govy.NameInferModeRuntime] in action.
// An alternative for the aforementioned mode which offers better runtime performance
// is [govy.NameInferModeGenerate].
//
// It comes at a cost of having to run the code generation utility before running your code.
// The utility generates code which uses [govy.SetInferredName].
// We'll use this very function in this example to simulate the code generation step.
// The first validator, 'v1', is created with [govy.NameInferModeDisable],
// the second validator, 'v2' is created with [govy.NameInferModeGenerate].
// As you can see in the output, only the second validator, 'v2' has the inferred name.
func ExampleNameInferModeGenerate() {
	govyconfig.SetNameInferIncludeTestFiles(true)
	defer govyconfig.SetNameInferIncludeTestFiles(false)

	v1 := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		NameInferMode(govy.NameInferModeGenerate).
		WithName("Teacher")

	govy.SetInferredName(govy.InferredName{
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
	// Validation for Teacher has failed:
	//   - should be equal to 'Jerry'
	// Validation for NotTeacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - should be equal to 'Thomas'
}

// Knowing when to call [govy.Validator.NameInferMode] is important.
// Beware that once a [govy.Validator.Validate] is called, it will cache the inferred name.
//
// To demonstrate this we'll set the [govy.NameInferModeDisable] and you will observe
// that the name is still inferred, although to be precise, it's not inferred anymore,
// it was inferred the first time [govy.Validator.Validate] was called and now it's cached.
func ExampleSetNameInferMode_changeModeInRuntime() {
	govyconfig.SetNameInferIncludeTestFiles(true)
	defer govyconfig.SetNameInferIncludeTestFiles(false)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		NameInferMode(govy.NameInferModeDisable).
		WithName("Teacher")

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---\nAfter setting Runtime infer mode.\n---")
	err = v.NameInferMode(govy.NameInferModeRuntime).Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed:
	//   - should be equal to 'Jerry'
	// ---
	// After setting Runtime infer mode.
	// ---
	// Validation for Teacher has failed:
	//   - should be equal to 'Jerry'
}

// By default govy's name inference is first checking for either json or yaml tags.
// If either is set it will use the value of the tag as the property name, see [govy.NameInferDefaultFunc].
//
// This behavior can be customized by providing a custom [govy.NameInferFunc]
// via [govy.SetNameInferFunc].
// Note that the tag value is the raw value of the struct tag,
// it needs to be further parsed with [reflect.StructTag].
//
// In the example below we're setting a custom name inference function which always returns the exact field name.
func ExampleSetNameInferFunc() {
	govyconfig.SetNameInferIncludeTestFiles(true)
	defer govyconfig.SetNameInferIncludeTestFiles(false)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		NameInferMode(govy.NameInferModeRuntime).
		WithName("Teacher")

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'Name' with value 'Tom':
	//     - should be equal to 'Jerry'
}
