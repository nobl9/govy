# Path Inference

## Topics

- [Choose a path inference mode](#choose-a-path-inference-mode)
  - [Configure inference modes and include test files when needed.](#configure-inference-modes-and-include-test-files-when-needed)
  - [Use generated inference by registering generated paths.](#use-generated-inference-by-registering-generated-paths)
- [Avoid late inference changes](#avoid-late-inference-changes)
  - [Changing inference mode after first validation does not recompute paths.](#changing-inference-mode-after-first-validation-does-not-recompute-paths)

## Choose a path inference mode

Prefer explicit paths unless inferred paths are part of the design.
Runtime inference resolves paths from source on first use.
Generated inference resolves paths before the application runs.

### Configure inference modes and include test files when needed

[//]: # (embed: ExampleInferPathMode?comments=false)

```go
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
}
```

### Use generated inference by registering generated paths

Run `govy inferpath -dir <directory> -pkg <package-name>` for the consumer package.
Generate and load the registrations before validation.
Use the package name, not its import path, for `-pkg`.
The example below manually simulates generated registrations.
Its file and line identify a getter in govy's test source.
For application code, generate entries from the actual getters.

[//]: # (embed: ExampleInferPathModeGenerate?comments=false)

```go
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
		Line: 2059,
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
}
```

## Avoid late inference changes

Path inference is cached after first validation.
Configure the mode before validation, not after an empty path has already been inferred.

### Changing inference mode after first validation does not recompute paths

[//]: # (embed: ExampleValidator_InferPath_changeModeInRuntime?comments=false)

```go
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
}
```
