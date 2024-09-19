package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"

	"github.com/nobl9/govy/internal"
)

// nolint: lll
const expectedGeneratedFile = `// Code generated by "govy nameinfer -dir %s -pkg validation -filename govy_inferred_names.go"; DO NOT EDIT.

package validation

import (
	"github.com/nobl9/govy/pkg/govyconfig"
)

func init() {
	inferredNames := map[string]govyconfig.InferredName{
		"student/nested.go:13": {
			Name: "name",
			File: "student/nested.go",
			Line: 13,
		},
		"university/university.go:13": {
			Name: "name",
			File: "university/university.go",
			Line: 13,
		},
		"validation.go:11": {
			Name: "Name",
			File: "validation.go",
			Line: 11,
		},
		"validation.go:13": {
			Name: "university.name",
			File: "validation.go",
			Line: 13,
		},
		"validation.go:15": {
			Name: "university",
			File: "validation.go",
			Line: 15,
		},
	}
	for _, name := range inferredNames {
		govyconfig.SetInferredName(name)
	}
}
`

var expectedCommandOutputLines = []string{
	"Found 'govy.For' function at: university/university.go:13",
	"Found 'govy.For' function at: validation.go:11",
	"Found 'govy.For' function at: validation.go:13",
	"Found 'govy.For' function at: validation.go:15",
	"Found 'govy.For' function at: student/nested.go:13",
}

var moduleRoot = internal.FindModuleRoot()

func TestCmd_NameInfer(t *testing.T) {
	tmpDir := t.TempDir()
	fileName := "govy_inferred_names.go"
	out := execCmd(t,
		"go", "run", "../../cmd/govy",
		nameInferCmdName,
		"-dir", tmpDir,
		"-pkg", "validation",
		"-filename", fileName,
	)
	assert.ElementsMatch(t, expectedCommandOutputLines, strings.Split(strings.TrimSpace(out.String()), "\n"))
	generatedFilePath := filepath.Join(tmpDir, fileName)
	data, err := os.ReadFile(generatedFilePath)
	assert.Require(t, assert.NoError(t, err))
	assert.Equal(t, fmt.Sprintf(expectedGeneratedFile, tmpDir), string(data))
}

func execCmd(t *testing.T, name string, arg ...string) *bytes.Buffer {
	t.Helper()
	// #nosec G204
	cmd := exec.Command(name, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	cmd.Dir = filepath.Join(moduleRoot, "tests/examplemodule")
	if err := cmd.Run(); err != nil {
		out := stderr.String()
		if out == "" {
			out = stdout.String()
		}
		t.Errorf("failed to execute '%s' command; err: %s, output: %s", cmd, err, out)
		t.FailNow()
	}
	return &stdout
}
