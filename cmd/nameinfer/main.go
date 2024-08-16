package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/exp/slices"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/nameinfer"
	"github.com/nobl9/govy/pkg/govyconfig"
)

type templateData struct {
	ProgramInvocation string
	Package           string
	Names             map[string]govyconfig.InferredName
}

//go:embed inferred_names.go.tmpl
var inferredNamesTemplateStr string

var inferredNamesTemplate = template.Must(
	template.New("inferred_names").Parse(inferredNamesTemplateStr))

func main() {
	govyconfig.SetLogLevel(slog.LevelDebug)

	dirFlag := flag.String("dir", "", "directory path to save the generated file")
	pkgFlag := flag.String("pkg", "", "package name")
	fileNameFlag := flag.String("filename", "govy_inferred_names.go", "file name")
	flag.Parse()

	outputDir := *dirFlag
	if outputDir == "" {
		errFatal("'-dir' flag is required")
	}
	outputPkg := *pkgFlag
	if outputPkg == "" {
		errFatal("'-outputPkg' flag is required")
	}
	fileName := *fileNameFlag

	programName := filepath.Base(os.Args[0])
	root := internal.FindModuleRoot()
	if root == "" {
		errFatal("failed to find module root")
	}

	modAST := nameinfer.NewModuleAST(root)

	names := make(map[string]govyconfig.InferredName)
	for _, pkg := range modAST.Packages {
		for i, f := range pkg.Syntax {
			importName := nameinfer.GetGovyImportName(f)
			ast.Inspect(f, func(n ast.Node) bool {
				selectorExpr, ok := n.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				exprIdent, ok := selectorExpr.X.(*ast.Ident)
				if !ok {
					return true
				}
				if exprIdent.Name != importName || !slices.Contains(nameinfer.FunctionsWithGetter, selectorExpr.Sel.Name) {
					return true
				}
				line := modAST.FileSet.Position(selectorExpr.Pos()).Line
				inferredName := nameinfer.InferNameFromFile(modAST.FileSet, pkg, f, line)
				name := govyconfig.InferredName{
					Name: inferredName,
					File: strings.TrimPrefix(pkg.GoFiles[i], root+"/"),
					Line: line,
				}
				fmt.Printf("Found 'govy.%s' function at: %s:%d\n", selectorExpr.Sel.Name, name.File, name.Line)
				key := fmt.Sprintf("%s:%d", name.File, name.Line)
				names[key] = name
				return false
			})
		}
	}

	if len(names) == 0 {
		errFatal("no names inferred")
	}

	buf := new(bytes.Buffer)
	if err := inferredNamesTemplate.Execute(buf, templateData{
		ProgramInvocation: fmt.Sprintf("%s %s", programName, strings.Join(os.Args[1:], " ")),
		Package:           outputPkg,
		Names:             names,
	}); err != nil {
		errFatal("failed to execute template: %v", err)
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		errFatal("failed to format produced template: %v", err)
	}
	outputName := filepath.Join(outputDir, fileName)
	if err = os.WriteFile(outputName, formatted, 0o600); err != nil {
		errFatal("failed to write file: %v", err)
	}
}

func errFatal(f string, a ...any) {
	if len(a) == 0 {
		fmt.Fprintln(os.Stderr, f)
	} else {
		fmt.Fprintf(os.Stderr, f+"\n", a...)
	}
	os.Exit(1)
}
