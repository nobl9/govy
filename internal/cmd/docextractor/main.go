package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/internal/collections"
	"github.com/nobl9/govy/internal/logging"
)

const (
	variableName = "templateFunctions"
	firstComment = "// The following functions are made available for use in the templates:"
)

// docextractor is a tool that extracts documentation from builtin template functions
// and adds them to the AddTemplateFunctions function in the message_templates.go file.
func main() {
	fmt.Println("Running docextractor...")

	root := internal.FindModuleRoot()
	docs := findTemplateFunctionsDocs(root)

	path := filepath.Join(root, "pkg", "govy", "message_templates.go")
	fileContents := readFile(path)
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, path, fileContents, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		logFatal(err, "Failed to parse file AST %q", path)
	}

	ast.Inspect(astFile, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.FuncDecl:
			if v.Name.Name != "AddTemplateFunctions" {
				return true
			}
			addTemplateFunctionComments(docs, v)
			return false
		}
		return true
	})

	file, err := os.Create(path) // #nosec G304
	if err != nil {
		logFatal(err, "Failed to open file %q", path)
	}
	defer func() { _ = file.Close() }()
	if err = format.Node(file, fset, astFile); err != nil {
		logFatal(err, "Failed to format and write file %q", path)
	}
}

func addTemplateFunctionComments(docs [][]string, funcDecl *ast.FuncDecl) {
	comments := funcDecl.Doc.List

	appendComments := func(texts ...string) {
		for _, text := range texts {
			comments = append(comments, &ast.Comment{
				Slash: funcDecl.Pos() - 1,
				Text:  text,
			})
		}
	}

	appendComments(firstComment)
	for _, templateFuncDocsLines := range docs {
		appendComments("//")
		for i, line := range templateFuncDocsLines {
			if i == 0 {
				line = "// - " + line
			} else {
				line = "//   " + line
			}
			appendComments(line)
		}
	}
	appendComments(
		"//",
		"// Refer to the testable examples of [AddTemplateFunctions] for more details",
		"// on each builtin function.",
	)
	funcDecl.Doc.List = comments
}

func readFile(path string) string {
	fileContents, err := os.ReadFile(path) // #nosec G304
	if err != nil {
		logFatal(err, "Failed to read file %q contents", path)
	}
	fileContentsStr := string(fileContents)
	if strings.Contains(fileContentsStr, firstComment) {
		firstCommentIdx := strings.Index(fileContentsStr, firstComment)
		funcIdx := strings.Index(fileContentsStr, "func AddTemplateFunctions(")
		fileContentsStr = fileContentsStr[:firstCommentIdx] + fileContentsStr[funcIdx:]
	}
	return fileContentsStr
}

func findTemplateFunctionsDocs(root string) [][]string {
	path := filepath.Join(root, "internal", "messagetemplates", "functions.go")

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		logFatal(err, "Failed to parse file %q", path)
	}

	var templateFunctionsExpr ast.Expr
	ast.Inspect(astFile, func(n ast.Node) bool {
		switch v := n.(type) {
		case *ast.ValueSpec:
			if len(v.Names) == 1 && v.Names[0].Name == variableName &&
				len(v.Values) == 1 {
				logging.Logger().Debug(fmt.Sprintf("found variable: %s", v.Names[0].Name))
				templateFunctionsExpr = v.Values[0]
				return false
			}
		}
		return true
	})

	if templateFunctionsExpr == nil {
		logFatal(nil, "Template functions %q variable was not found", variableName)
	}

	compositeLiteral, ok := templateFunctionsExpr.(*ast.CompositeLit)
	if !ok {
		logFatal(nil, "Template functions %q variable is not a %T", variableName, &ast.CompositeLit{})
	}

	templateFunctions := make(map[string]string, len(compositeLiteral.Elts))
	for _, el := range compositeLiteral.Elts {
		kv := el.(*ast.KeyValueExpr)

		key := kv.Key.(*ast.BasicLit)
		value := kv.Value.(*ast.Ident)

		funcName, err := strconv.Unquote(key.Value)
		if err != nil {
			logFatal(nil, "Failed to unquote template function name %q", key.Value)
		}
		templateFunctions[funcName] = value.Name
	}

	docsList := make([][]string, 0, len(templateFunctions))
	for _, templateFuncName := range collections.SortedKeys(templateFunctions) {
		goFuncName := templateFunctions[templateFuncName]
		var funcDecl *ast.FuncDecl
		ast.Inspect(astFile, func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncDecl:
				if v.Name.Name == goFuncName {
					funcDecl = v
					return false
				}
			}
			return true
		})
		if funcDecl == nil {
			logFatal(nil, "Function %q is not defined in the file", goFuncName)
		}
		if funcDecl.Doc == nil {
			logFatal(nil, "Function %q is missing documentation", goFuncName)
		}
		docLines := make([]string, 0, len(funcDecl.Doc.List))
		for _, comment := range funcDecl.Doc.List {
			text := strings.TrimPrefix(comment.Text, "//")
			text = strings.TrimSpace(text)
			text = strings.ReplaceAll(text, goFuncName, "'"+templateFuncName+"'")
			docLines = append(docLines, text)
		}
		docsList = append(docsList, docLines)
	}
	return docsList
}

func logFatal(err error, msg string, a ...any) {
	var attrs []slog.Attr
	if err != nil {
		attrs = append(attrs, slog.String("error", err.Error()))
	}
	logging.Logger().LogAttrs(context.Background(), slog.LevelError, fmt.Sprintf(msg, a...), attrs...)
}
