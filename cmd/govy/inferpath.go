package main

import (
	"bytes"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/inferpath"
	"github.com/nobl9/govy/pkg/govyconfig"
)

type inferPathTemplateData struct {
	ProgramInvocation string
	Package           string
	Names             map[string]govyconfig.InferredPath
}

//go:embed inferred_names.go.tmpl
var inferredNamesTemplateStr string

var inferredNamesTemplate = template.Must(
	template.New("inferred_names").Parse(inferredNamesTemplateStr))

func newInferPathCommand() *inferPathCommand {
	fset := flag.NewFlagSet(inferPathCmdName, flag.ExitOnError)
	cmd := &inferPathCommand{fset: fset}
	fset.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s %s:\n", govyCmdName, inferPathCmdName)
		fmt.Fprintf(os.Stderr, "  %s %s [flags]\n", govyCmdName, inferPathCmdName)
		fmt.Fprintf(os.Stderr, "Flags:\n")
		fset.PrintDefaults()
	}
	fset.StringVar(&cmd.outputDir, "dir", "", "directory path to save the generated file")
	fset.StringVar(&cmd.pkg, "pkg", "", "package name of the generated file")
	fset.StringVar(&cmd.fileName, "filename", "govy_inferred_names.go", "generated file name")
	return cmd
}

type inferPathCommand struct {
	fset      *flag.FlagSet
	outputDir string
	pkg       string
	fileName  string
}

func (n *inferPathCommand) Run() error {
	_ = n.fset.Parse(os.Args[2:])
	if n.outputDir == "" {
		errFatalWithUsage(n.fset, "'-dir' flag is required")
	}
	if n.pkg == "" {
		errFatalWithUsage(n.fset, "'-pkg' flag is required")
	}

	root := internal.FindModuleRoot()
	if root == "" {
		return errors.New("failed to find module root")
	}

	modAST := inferpath.NewModuleAST(root)

	names := make(map[string]govyconfig.InferredPath)
	for _, pkg := range modAST.Packages {
		for i, f := range pkg.Syntax {
			importName := inferpath.GetGovyImportName(f)
			ast.Inspect(f, func(n ast.Node) bool {
				selectorExpr, ok := n.(*ast.SelectorExpr)
				if !ok {
					return true
				}
				exprIdent, ok := selectorExpr.X.(*ast.Ident)
				if !ok {
					return true
				}
				if exprIdent.Name != importName ||
					!slices.Contains(inferpath.FunctionsWithGetter, selectorExpr.Sel.Name) {
					return true
				}
				line := modAST.FileSet.Position(selectorExpr.Pos()).Line
				inferredPath := inferpath.InferPathFromFile(modAST.FileSet, pkg, f, line)
				name := govyconfig.InferredPath{
					Path: inferredPath.String(),
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
	if err := inferredNamesTemplate.Execute(buf, inferPathTemplateData{
		ProgramInvocation: fmt.Sprintf("%s %s", govyCmdName, strings.Join(os.Args[1:], " ")),
		Package:           n.pkg,
		Names:             names,
	}); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format produced template: %w", err)
	}
	outputName := filepath.Join(n.outputDir, n.fileName)
	if err = os.WriteFile(outputName, formatted, 0o600); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
