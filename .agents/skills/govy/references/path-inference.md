# Path Inference

Runtime and generated path inference, govyconfig registration, and inference caching behavior.

## Topics

- [Choose a path inference mode](#choose-a-path-inference-mode)
  - [Configure inference modes and include test files when needed.](#configure-inference-modes-and-include-test-files-when-needed)
  - [Use generated inference by registering generated paths.](#use-generated-inference-by-registering-generated-paths)
- [Avoid late inference changes](#avoid-late-inference-changes)
  - [Changing inference mode after first validation does not recompute paths.](#changing-inference-mode-after-first-validation-does-not-recompute-paths)

## Choose a path inference mode

Prefer explicit paths unless inferred paths are part of the design. Runtime inference trades startup simplicity for a one-time AST lookup, while generated inference moves that work into code generation.

<a id="configure-inference-modes-and-include-test-files-when-needed"></a>

**Configure inference modes and include test files when needed.**

[//]: # (embed: ExampleInferPathMode)

```go
// In the interactive tutorial for govy, we've been using
// [govy.PropertyRules.WithName] to provide explicit path segments for our properties.
//
// Ideally, we'd want govy to derive those paths directly from the getter expressions,
// matching the struct fields selected by the user.
// Go uses struct tags to achieve that,
// and libraries like [encoding/json] use these tags to encode/decode structs.
// Unfortunately, there's no easy way to tell what exact property we're returning from [govy.PropertyGetter].
//
// To solve this problem, govy provides a way to infer the property path (with a catch).
// The catch being that the path inference mechanism needs to parse the whole modules' AST.
// This can be a performance hit, especially for large projects if not done properly.
//
// By default govy will not attempt to infer any property paths.
//
// So, how do we do that properly?
// Both [govy.Validator] and [govy.PropertyRules] (including variants) have a dedicated method
// to configure how property paths are inferred.
//
// It depends on the [govy.InferPathMode] used:
//   - [govy.InferPathModeDisable], path inference is disabled (default), nothing to do here
//   - [govy.InferPathModeRuntime], the path is inferred during runtime from the getter expression.
//     This is the most flexible option, but also the slowest, although the slowdown
//     is incurred only once, whenever [govy.PropertyRules.Validate] is first called.
//     If you make sure that [govy.PropertyRules] is created only once and don't mind
//     the one-time performance hit, this should be enough for you.
//   - [govy.InferPathModeGenerate], the path is inferred during a separate code generation phase.
//     This mode requires you to run `govy inferpath` before you run your code.
//     It generates a file with inferred relative paths for your getter call sites,
//     which automatically registers them using [govyconfig.SetInferredPath].
//
// Since this tutorial is run as a test,
// we need to explicitly instruct govy to infer paths from test files.
// By default, test files are not parsed to improve performance.
// In order to do that, we use [govyconfig.SetInferPathIncludeTestFiles].
func ExampleInferPathMode() {
	govyconfig.SetInferPathIncludeTestFiles(true)
	defer govyconfig.SetInferPathIncludeTestFiles(false)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		InferPath(govy.InferPathModeRuntime).
		WithName("Teacher")

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - must be equal to 'Jerry'
}
```

<a id="use-generated-inference-by-registering-generated-paths"></a>

**Use generated inference by registering generated paths.**

[//]: # (embed: ExampleInferPathModeGenerate)

```go
// In the previous example we've seen [govy.InferPathModeRuntime] in action.
// An alternative for the aforementioned mode which offers better runtime performance
// is [govy.InferPathModeGenerate].
//
// It comes at a cost of having to run the code generation utility before running your code.
// The utility generates code which uses [govyconfig.SetInferredPath].
// We'll use this very function in this example to simulate the code generation step.
// The first validator, 'v1', is created with [govy.InferPathModeDisable],
// the second validator, 'v2' is created with [govy.InferPathModeGenerate].
// As you can see in the output, only the second validator, 'v2' has the inferred path.
func ExampleInferPathModeGenerate() {
	govyconfig.SetInferPathIncludeTestFiles(true)
	defer govyconfig.SetInferPathIncludeTestFiles(false)

	v1 := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		InferPath(govy.InferPathModeDisable).
		WithName("Teacher")

	govyconfig.SetInferredPath(govyconfig.InferredPath{
		Path: jsonpath.New().Name("name"),
		File: "pkg/govy/example_test.go",
		Line: 2042,
	})

	v2 := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Thomas")),
	).
		InferPath(govy.InferPathModeGenerate).
		WithName("NotTeacher")

	teacher := Teacher{Name: "Tom"}
	if err := v1.Validate(teacher); err != nil {
		fmt.Println(err)
	}
	if err := v2.Validate(teacher); err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed:
	//   - must be equal to 'Jerry'
	// Validation for NotTeacher has failed for the following properties:
	//   - 'name' with value 'Tom':
	//     - must be equal to 'Thomas'
}
```

## Avoid late inference changes

Path inference is cached after first validation. Configure the mode before validation, not after an empty path has already been inferred.

<a id="changing-inference-mode-after-first-validation-does-not-recompute-paths"></a>

**Changing inference mode after first validation does not recompute paths.**

[//]: # (embed: ExampleValidator_InferPath_changeModeInRuntime)

```go
// Knowing when to call [govy.Validator.InferPath] is important.
// The path inference runs only once per [govy.PropertyRules] instance, on the first validation.
// Once this happens, the result is cached, even if that result is an empty path.
//
// This example demonstrates that changing the mode after the first validation has no effect.
// The first validation runs with [govy.InferPathModeDisable], which produces an empty path.
// This empty result is then cached. Even after switching to [govy.InferPathModeRuntime],
// the cached empty result persists, so no property path appears in the output.
func ExampleValidator_InferPath_changeModeInRuntime() {
	govyconfig.SetInferPathIncludeTestFiles(true)
	defer govyconfig.SetInferPathIncludeTestFiles(false)

	v := govy.New(
		govy.For(func(t Teacher) string { return t.Name }).
			Rules(rules.EQ("Jerry")),
	).
		InferPath(govy.InferPathModeDisable).
		WithName("Teacher")

	teacher := Teacher{Name: "Tom"}
	err := v.Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---\nAfter setting Runtime infer mode.\n---")
	err = v.InferPath(govy.InferPathModeRuntime).Validate(teacher)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// Validation for Teacher has failed:
	//   - must be equal to 'Jerry'
	// ---
	// After setting Runtime infer mode.
	// ---
	// Validation for Teacher has failed:
	//   - must be equal to 'Jerry'
}
```
