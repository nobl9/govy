package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"go/ast"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/assert"
)

const fatalHelperEnv = "GOVY_DOCEXTRACTOR_FATAL_HELPER"

type fatalLog struct {
	Time   string `json:"time"`
	Level  string `json:"level"`
	Source struct {
		Function string `json:"function"`
		File     string `json:"file"`
		Line     int    `json:"line"`
	} `json:"source"`
	Message string `json:"msg"`
	Error   string `json:"error"`
}

func TestParseGeneratedDocsConfig(t *testing.T) {
	tests := []struct {
		name     string
		docsRef  string
		expected generatedDocsConfig
	}{
		{
			name:    "defaults to functions without a return filter",
			docsRef: "pkg/rules",
			expected: generatedDocsConfig{
				path: "pkg/rules",
				kind: "func",
			},
		},
		{
			name:    "parses kind and multiple return filters",
			docsRef: "pkg/rules?kind=func&returns=govy.Rule,govy.RuleSet",
			expected: generatedDocsConfig{
				path:    "pkg/rules",
				kind:    "func",
				returns: []string{"govy.Rule", "govy.RuleSet"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, parseGeneratedDocsConfig(test.docsRef))
		})
	}
}

func TestFirstDocSentence(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		expected string
	}{
		{
			name:     "ordinary sentence",
			lines:    []string{"Validates a value. Additional detail."},
			expected: "Validates a value.",
		},
		{
			name:     "doc link",
			lines:    []string{"Uses [jsonpath.Path]. Additional detail."},
			expected: "Uses [jsonpath.Path].",
		},
		{
			name:     "known abbreviation",
			lines:    []string{"Handles e.g. URLs. Additional detail."},
			expected: "Handles e.g. URLs.",
		},
		{
			name:     "decimal number",
			lines:    []string{"Version 1.2 is supported. Additional detail."},
			expected: "Version 1.2 is supported.",
		},
		{
			name:     "no terminal period",
			lines:    []string{"No terminal period"},
			expected: "No terminal period",
		},
		{
			name:     "URL",
			lines:    []string{"See https://example.com/docs. Additional detail."},
			expected: "See https://example.com/docs.",
		},
		{
			name:     "Unicode",
			lines:    []string{"Value 🧪 validates. Additional detail."},
			expected: "Value 🧪 validates.",
		},
		{
			name:     "blank second paragraph",
			lines:    []string{"First paragraph without a period", "", "Second paragraph."},
			expected: "First paragraph without a period",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			comments := make([]*ast.Comment, 0, len(test.lines))
			for _, line := range test.lines {
				comments = append(comments, &ast.Comment{Text: "// " + line})
			}
			doc := &ast.CommentGroup{List: comments}
			assert.Equal(t, test.expected, firstDocSentence(doc))
		})
	}
}

func TestRenderGeneratedDocs(t *testing.T) {
	root := testdataRoot(t)
	expected := "- `Alpha` - Alpha uses [jsonpath.Path].\n" +
		"- `Middle` - Middle handles e.g. values.\n" +
		"- `Zulu` - Zulu validates a value.\n"

	actual := renderGeneratedDocs(
		root,
		"generateddocs?kind=func&returns=govy.Rule,govy.RuleSet",
	)
	assert.Equal(t, expected, actual)
	assert.Equal(t, actual, renderGeneratedDocs(
		root,
		"generateddocs?kind=func&returns=govy.Rule,govy.RuleSet",
	))
}

func TestReplaceGeneratedDocs(t *testing.T) {
	root := testdataRoot(t)
	input := readTestFile(t, filepath.Join(root, "markdown", "generated-docs-input.md"))
	expected := readTestFile(t, filepath.Join(root, "golden", "generated-docs.md"))

	actual := replaceGeneratedDocs(root, input, "catalog.md")
	assert.Equal(t, expected, actual)
	assert.Equal(t, actual, replaceGeneratedDocs(root, actual, "catalog.md"))

	t.Run("region ending at EOF", func(t *testing.T) {
		input := "[//]: # (docs: generateddocs?kind=func&returns=govy.RuleSet)\n" +
			"stale\n" +
			"[//]: # (end-docs)"
		expected := "[//]: # (docs: generateddocs?kind=func&returns=govy.RuleSet)\n\n" +
			"- `Alpha` - Alpha uses [jsonpath.Path].\n\n" +
			"[//]: # (end-docs)"
		assert.Equal(t, expected, replaceGeneratedDocs(root, input, "catalog.md"))
	})

	t.Run("zero matches", func(t *testing.T) {
		root := t.TempDir()
		writeTestFile(t, filepath.Join(root, "empty", "ignored_test.go"), "package empty\n", 0o644)
		input := "[//]: # (docs: empty?kind=func&returns=govy.Rule)\n" +
			"stale\n" +
			"[//]: # (end-docs)\n"
		expected := "[//]: # (docs: empty?kind=func&returns=govy.Rule)\n\n\n\n" +
			"[//]: # (end-docs)\n"
		assert.Equal(t, expected, replaceGeneratedDocs(root, input, "empty.md"))
	})
}

func TestReadEmbeddedExample(t *testing.T) {
	root := testdataRoot(t)
	tests := []struct {
		name       string
		exampleRef string
		expected   string
	}{
		{
			name:       "full file",
			exampleRef: "examples/full.go",
			expected:   readTestFile(t, filepath.Join(root, "examples", "full.go")),
		},
		{
			name:       "named function",
			exampleRef: "examples/named.go#ExampleNamed",
			expected: "// ExampleNamed demonstrates named extraction.\n" +
				"func ExampleNamed() {\n" +
				"\tprintln(\"named\")\n" +
				"}",
		},
		{
			name:       "unique bare function",
			exampleRef: "ExampleUnique",
			expected: "// ExampleUnique demonstrates bare lookup.\n" +
				"func ExampleUnique() {\n" +
				"\tprintln(\"unique\")\n" +
				"}",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, readEmbeddedExample(root, test.exampleRef))
		})
	}
}

func TestReplaceEmbeddedExamples(t *testing.T) {
	root := testdataRoot(t)
	input := readTestFile(t, filepath.Join(root, "markdown", "embedded-input.md"))
	expected := readTestFile(t, filepath.Join(root, "golden", "embedded.md"))

	actual := replaceEmbeddedExamples(root, input, "examples.md")
	assert.Equal(t, expected, actual)
	assert.Equal(t, actual, replaceEmbeddedExamples(root, actual, "examples.md"))
}

func TestEmbedExamples(t *testing.T) {
	root := testdataRoot(t)
	workspace := t.TempDir()
	input := readTestFile(t, filepath.Join(root, "markdown", "embedded-input.md"))
	expected := readTestFile(t, filepath.Join(root, "golden", "embedded.md"))

	explicitPath := filepath.Join(workspace, "single document.md")
	docsDir := filepath.Join(workspace, "docs with spaces")
	nestedPath := filepath.Join(docsDir, "nested", "nested.md")
	executablePath := filepath.Join(docsDir, "executable.md")
	ignoredPath := filepath.Join(docsDir, "ignored.txt")
	writeTestFile(t, explicitPath, input, 0o600)
	writeTestFile(t, nestedPath, input, 0o644)
	writeTestFile(t, executablePath, input, 0o755)
	writeTestFile(t, ignoredPath, input, 0o644)

	embedExamples(root, []string{explicitPath, docsDir})

	for path, mode := range map[string]os.FileMode{
		explicitPath:   0o600,
		nestedPath:     0o644,
		executablePath: 0o755,
	} {
		assert.Equal(t, expected, readTestFile(t, path))
		info, err := os.Stat(path)
		assert.Require(t, assert.NoError(t, err))
		assert.Equal(t, mode, info.Mode().Perm())
	}
	assert.Equal(t, input, readTestFile(t, ignoredPath))

	embedExamples(root, []string{explicitPath, docsDir})
	for _, path := range []string{explicitPath, nestedPath, executablePath} {
		assert.Equal(t, expected, readTestFile(t, path))
	}
}

func TestDocextractorFatalErrors(t *testing.T) {
	t.Run("missing constructor documentation", func(t *testing.T) {
		root := t.TempDir()
		path := filepath.Join(root, "rules", "missing.go")
		writeTestFile(t, path, "package rules\n\nfunc Missing() govy.Rule[string] { panic(\"fixture\") }\n", 0o644)

		entry := runFatalHelper(
			t,
			"render-docs",
			root,
			"rules?kind=func&returns=govy.Rule",
		)
		assert.Equal(
			t,
			"Exported function \"Missing\" in \""+path+"\" is missing documentation",
			entry.Message,
		)
	})

	t.Run("missing bare example", func(t *testing.T) {
		entry := runFatalHelper(t, "read-example", t.TempDir(), "ExampleMissing")
		assert.Equal(
			t,
			"Function \"ExampleMissing\" was not found in example_test.go files",
			entry.Message,
		)
	})

	t.Run("ambiguous bare example", func(t *testing.T) {
		root := t.TempDir()
		for _, dir := range []string{"alpha", "zulu"} {
			writeTestFile(
				t,
				filepath.Join(root, dir, "example_test.go"),
				"package "+dir+"\n\nfunc ExampleDuplicate() {}\n",
				0o644,
			)
		}

		entry := runFatalHelper(t, "read-example", root, "ExampleDuplicate")
		assert.Equal(
			t,
			"Function \"ExampleDuplicate\" is ambiguous across example_test.go files: "+
				filepath.Join(root, "alpha", "example_test.go")+", "+
				filepath.Join(root, "zulu", "example_test.go"),
			entry.Message,
		)
	})

	t.Run("malformed embed directive", func(t *testing.T) {
		root := t.TempDir()
		writeTestFile(t, filepath.Join(root, "source.go"), "package fixture\n", 0o644)
		markdownPath := filepath.Join(root, "malformed.md")
		writeTestFile(
			t,
			markdownPath,
			"[//]: # (embed: source.go\n\n```go\nstale\n```\n",
			0o644,
		)

		entry := runFatalHelper(t, "embed-file", root, markdownPath)
		assert.Equal(
			t,
			"Malformed embed directive \"[//]: # (embed: source.go\" in \""+markdownPath+"\"",
			entry.Message,
		)
	})

	t.Run("missing Go fence", func(t *testing.T) {
		root := t.TempDir()
		writeTestFile(t, filepath.Join(root, "source.go"), "package fixture\n", 0o644)
		markdownPath := filepath.Join(root, "missing-fence.md")
		writeTestFile(t, markdownPath, "[//]: # (embed: source.go)\n\nNo code block.\n", 0o644)

		entry := runFatalHelper(t, "embed-file", root, markdownPath)
		assert.Equal(
			t,
			"Embed directive \"source.go\" in \""+markdownPath+"\" is missing a Go code block",
			entry.Message,
		)
	})

	t.Run("unterminated Go fence", func(t *testing.T) {
		root := t.TempDir()
		writeTestFile(t, filepath.Join(root, "source.go"), "package fixture\n", 0o644)
		markdownPath := filepath.Join(root, "unterminated.md")
		writeTestFile(t, markdownPath, "[//]: # (embed: source.go)\n\n```go\nstale\n", 0o644)

		entry := runFatalHelper(t, "embed-file", root, markdownPath)
		assert.Equal(
			t,
			"Embed directive \"source.go\" in \""+markdownPath+"\" has an unterminated Go code block",
			entry.Message,
		)
	})

	t.Run("missing generated docs end marker", func(t *testing.T) {
		root := testdataRoot(t)
		markdownPath := filepath.Join(t.TempDir(), "missing-end.md")
		writeTestFile(
			t,
			markdownPath,
			"[//]: # (docs: generateddocs?kind=func&returns=govy.RuleSet)\nstale\n",
			0o644,
		)

		entry := runFatalHelper(t, "replace-docs-file", root, markdownPath)
		assert.Equal(
			t,
			"Docs directive \"generateddocs?kind=func&returns=govy.RuleSet\" in \""+
				markdownPath+"\" is missing \"[//]: # (end-docs)\"",
			entry.Message,
		)
	})

	t.Run("unsupported generated docs kind", func(t *testing.T) {
		root := testdataRoot(t)
		markdownPath := filepath.Join(t.TempDir(), "unsupported-kind.md")
		writeTestFile(
			t,
			markdownPath,
			"[//]: # (docs: generateddocs?kind=type)\nstale\n[//]: # (end-docs)\n",
			0o644,
		)

		entry := runFatalHelper(t, "replace-docs-file", root, markdownPath)
		assert.Equal(t, "Unsupported docs kind \"type\"", entry.Message)
	})
}

func TestDocextractorCLI(t *testing.T) {
	repoRoot := internal.FindModuleRoot()
	binaryPath := filepath.Join(t.TempDir(), "docextractor")
	// #nosec G204 -- the output path is created by the test.
	build := exec.CommandContext(t.Context(), "go", "build", "-o", binaryPath, "./internal/cmd/docextractor")
	build.Dir = repoRoot
	buildOutput, err := build.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to build docextractor: %v\n%s", err, buildOutput)
	}

	t.Run("default mode remains idempotent", func(t *testing.T) {
		root := t.TempDir()
		writeTestFile(t, filepath.Join(root, "go.mod"), "module example.com/fixture\n\ngo 1.25.5\n", 0o644)
		writeTestFile(
			t,
			filepath.Join(root, "internal", "messagetemplates", "functions.go"),
			"package messagetemplates\n\n"+
				"var templateFunctions = map[string]any{\n\t\"upper\": upper,\n}\n\n"+
				"// upper transforms input.\nfunc upper(value string) string { return value }\n",
			0o644,
		)
		messageTemplatesPath := filepath.Join(root, "pkg", "govy", "message_templates.go")
		writeTestFile(
			t,
			messageTemplatesPath,
			"package govy\n\n// AddTemplateFunctions adds template functions.\n"+
				"func AddTemplateFunctions() {}\n",
			0o644,
		)
		expected := "package govy\n\n" +
			"// AddTemplateFunctions adds template functions.\n" +
			"// The following functions are made available for use in the templates:\n" +
			"//\n" +
			"// - 'upper' transforms input.\n" +
			"//\n" +
			"// Refer to the testable examples of [AddTemplateFunctions] for more details\n" +
			"// on each builtin function.\n" +
			"func AddTemplateFunctions() {}\n"

		stdout, stderr, err := runCommand(t, binaryPath, root)
		assert.Require(t, assert.NoError(t, err))
		assert.Equal(t, "Running docextractor...\n", stdout)
		assert.Equal(t, "", stderr)
		assert.Equal(t, expected, readTestFile(t, messageTemplatesPath))

		stdout, stderr, err = runCommand(t, binaryPath, root)
		assert.Require(t, assert.NoError(t, err))
		assert.Equal(t, "Running docextractor...\n", stdout)
		assert.Equal(t, "", stderr)
		assert.Equal(t, expected, readTestFile(t, messageTemplatesPath))
	})

	t.Run("embed requires paths", func(t *testing.T) {
		stdout, stderr, err := runCommand(t, binaryPath, repoRoot, "embed")
		assertExitCode(t, err)
		assert.Equal(t, "Running docextractor...\n", stdout)
		entry := decodeFatalLog(t, stderr)
		assert.Equal(t, "No Markdown files provided for example embedding", entry.Message)
	})

	t.Run("unknown subcommand", func(t *testing.T) {
		stdout, stderr, err := runCommand(t, binaryPath, repoRoot, "unknown")
		assertExitCode(t, err)
		assert.Equal(t, "Running docextractor...\n", stdout)
		entry := decodeFatalLog(t, stderr)
		assert.Equal(t, "Unknown docextractor command \"unknown\"", entry.Message)
	})
}

func TestEmbedExamplesScript(t *testing.T) {
	repoRoot := internal.FindModuleRoot()
	scriptPath := filepath.Join(repoRoot, "scripts", "embed-examples-in-markdown.bash")

	t.Run("no arguments", func(t *testing.T) {
		stdout, stderr, err := runCommand(t, scriptPath, repoRoot)
		assertExitCode(t, err)
		assert.Equal(t, "Usage: "+scriptPath+" <MARKDOWN_PATH>...\n", stdout)
		assert.Equal(t, "", stderr)
	})

	t.Run("multiple paths including spaces", func(t *testing.T) {
		workspace := t.TempDir()
		refs := []string{
			"pkg/govy/example_test.go#ExampleNew",
			"pkg/govy/example_test.go#ExampleValidator_WithName",
		}
		paths := []string{
			filepath.Join(workspace, "first document.md"),
			filepath.Join(workspace, "second document.md"),
		}
		for i, path := range paths {
			writeTestFile(t, path, markdownWithEmbeddedBody(refs[i], "stale"), 0o644)
		}

		stdout, stderr, err := runCommand(t, scriptPath, repoRoot, paths...)
		assert.Require(t, assert.NoError(t, err))
		assert.Equal(t, "Running docextractor...\n", stdout)
		assert.Equal(t, "", stderr)
		for i, path := range paths {
			example := readEmbeddedExample(repoRoot, refs[i])
			assert.Equal(t, markdownWithEmbeddedBody(refs[i], example), readTestFile(t, path))
		}
	})

	t.Run("generator failure is propagated", func(t *testing.T) {
		missingPath := filepath.Join(t.TempDir(), "missing.md")
		stdout, stderr, err := runCommand(t, scriptPath, repoRoot, missingPath)
		assertExitCode(t, err)
		assert.Equal(t, "Running docextractor...\n", stdout)
		entry := decodeGoRunFatalLog(t, stderr)
		assert.Equal(t, "Failed to stat path \""+missingPath+"\"", entry.Message)
		assert.Equal(t, "stat "+missingPath+": no such file or directory", entry.Error)
	})
}

func TestDocextractorFatalHelper(t *testing.T) {
	if os.Getenv(fatalHelperEnv) != "1" {
		return
	}

	root := os.Getenv("GOVY_DOCEXTRACTOR_ROOT")
	target := os.Getenv("GOVY_DOCEXTRACTOR_TARGET")
	switch os.Getenv("GOVY_DOCEXTRACTOR_MODE") {
	case "parse-config":
		parseGeneratedDocsConfig(target)
	case "render-docs":
		renderGeneratedDocs(root, target)
	case "read-example":
		readEmbeddedExample(root, target)
	case "embed-file":
		embedExamplesInMarkdown(root, target)
	case "replace-docs-file":
		contents, err := os.ReadFile(target) // #nosec G703 -- the parent test supplies this fixture path.
		if err != nil {
			t.Fatal(err)
		}
		replaceGeneratedDocs(root, string(contents), target)
	default:
		t.Fatalf("unknown helper mode %q", os.Getenv("GOVY_DOCEXTRACTOR_MODE"))
	}
	t.Fatal("fatal helper returned without exiting")
}

func runFatalHelper(t *testing.T, mode, root, target string) fatalLog {
	t.Helper()
	// #nosec G204,G702 -- os.Args[0] is the current test binary.
	cmd := exec.CommandContext(t.Context(), os.Args[0], "-test.run=^TestDocextractorFatalHelper$")
	cmd.Env = append(
		os.Environ(),
		fatalHelperEnv+"=1",
		"GOVY_DOCEXTRACTOR_MODE="+mode,
		"GOVY_DOCEXTRACTOR_ROOT="+root,
		"GOVY_DOCEXTRACTOR_TARGET="+target,
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	assertExitCode(t, err)
	assert.Equal(t, "", stdout.String())
	return decodeFatalLog(t, stderr.String())
}

func decodeFatalLog(t *testing.T, output string) fatalLog {
	t.Helper()
	decoder := json.NewDecoder(strings.NewReader(output))
	decoder.DisallowUnknownFields()
	var entry fatalLog
	assert.Require(t, assert.NoError(t, decoder.Decode(&entry)))
	assert.True(t, entry.Time != "")
	assert.Equal(t, "ERROR", entry.Level)
	assert.True(t, entry.Source.Function != "")
	assert.True(t, entry.Source.File != "")
	assert.True(t, entry.Source.Line > 0)
	assert.True(t, entry.Message != "")
	var extra any
	assert.Require(t, assert.True(t, errors.Is(decoder.Decode(&extra), io.EOF)))
	return entry
}

func decodeGoRunFatalLog(t *testing.T, output string) fatalLog {
	t.Helper()
	const trailer = "exit status 1\n"
	assert.Require(t, assert.True(t, strings.HasSuffix(output, trailer)))
	return decodeFatalLog(t, strings.TrimSuffix(output, trailer))
}

func assertExitCode(t *testing.T, err error) {
	t.Helper()
	var exitError *exec.ExitError
	if !errors.As(err, &exitError) {
		t.Fatalf("expected exit code 1, got error %v", err)
	}
	assert.Equal(t, 1, exitError.ExitCode())
}

func runCommand(t *testing.T, name, dir string, args ...string) (stdoutText, stderrText string, err error) {
	t.Helper()
	cmd := exec.CommandContext(t.Context(), name, args...) // #nosec G204 -- arguments are controlled by tests.
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	return stdout.String(), stderr.String(), err
}

func testdataRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs("testdata")
	assert.Require(t, assert.NoError(t, err))
	return root
}

func readTestFile(t *testing.T, path string) string {
	t.Helper()
	contents, err := os.ReadFile(path)
	assert.Require(t, assert.NoError(t, err))
	return string(contents)
}

func writeTestFile(t *testing.T, path, contents string, mode os.FileMode) {
	t.Helper()
	assert.Require(t, assert.NoError(t, os.MkdirAll(filepath.Dir(path), 0o750)))
	assert.Require(t, assert.NoError(t, os.WriteFile(path, []byte(contents), mode)))
	assert.Require(t, assert.NoError(t, os.Chmod(path, mode)))
}

func markdownWithEmbeddedBody(ref, body string) string {
	return "[//]: # (embed: " + ref + ")\n\n```go\n" + body + "\n```\n"
}
