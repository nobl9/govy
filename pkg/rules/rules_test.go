package rules

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal"
)

func TestRules_EnsureTestsAndBenchmarksAreWritten(t *testing.T) {
	// Functions that should be excluded from this test.
	// It's easier to list them here rather than complicate the AST traversal.
	excludeFuncs := map[string]bool{
		"HashFuncSelf":     true,
		"Compare":          true,
		"CompareDeepEqual": true,
	}
	rulesDir := filepath.Join(internal.FindModuleRoot(), "pkg/rules")
	fset := token.NewFileSet()

	// Parse the directory
	pkgs, err := parser.ParseDir(fset, rulesDir, nil, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse directory: %v", err)
	}

	exportedFuncs := make(map[string]struct{})
	testFuncs := make(map[string]bool)
	benchmarkFuncs := make(map[string]bool)

	// Collect exported functions
	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			if strings.HasSuffix(fileName, "_test.go") {
				continue
			}
			ast.Inspect(file, func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)
				if !ok {
					return true
				}
				if excludeFuncs[fn.Name.Name] || !fn.Name.IsExported() || fn.Recv != nil {
					return false
				}
				exportedFuncs[fn.Name.Name] = struct{}{}
				return false
			})
		}
	}

	// Collect test and benchmark functions
	for _, pkg := range pkgs {
		for fileName, file := range pkg.Files {
			if !strings.HasSuffix(fileName, "_test.go") {
				continue
			}
			ast.Inspect(file, func(n ast.Node) bool {
				if fn, ok := n.(*ast.FuncDecl); ok {
					if strings.HasPrefix(fn.Name.Name, "Test") {
						testFuncs[fn.Name.Name] = true
					} else if strings.HasPrefix(fn.Name.Name, "Benchmark") {
						benchmarkFuncs[fn.Name.Name] = true
					}
				}
				return true
			})
		}
	}

	// Check for corresponding test and benchmark functions
	for funcName := range exportedFuncs {
		testName := "Test" + funcName
		benchmarkName := "Benchmark" + funcName

		if !mapContainsString(testFuncs, testName) {
			t.Errorf("Missing test function for %s", funcName)
		}
		if !mapContainsString(benchmarkFuncs, benchmarkName) {
			t.Errorf("Missing benchmark function for %s", funcName)
		}
	}
}

func mapContainsString(m map[string]bool, s string) bool {
	if _, ok := m[s]; ok {
		return true
	}
	for k := range m {
		if strings.Contains(k, s) {
			return true
		}
	}
	return false
}
