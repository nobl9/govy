package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/fs"
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
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "embed":
			embedExamples(root, os.Args[2:])
			return
		default:
			logFatal(nil, "Unknown docextractor command %q", os.Args[1])
		}
	}

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

func embedExamples(root string, paths []string) {
	if len(paths) == 0 {
		logFatal(nil, "No Markdown files provided for example embedding")
	}

	for _, path := range paths {
		info, err := os.Stat(path) // #nosec G703
		if err != nil {
			logFatal(err, "Failed to stat path %q", path)
		}
		if !info.IsDir() {
			embedExamplesInMarkdown(root, path)
			continue
		}
		if err = filepath.WalkDir(path, func(path string, entry fs.DirEntry, err error) error { // #nosec G703
			if err != nil {
				return err
			}
			if entry.IsDir() || filepath.Ext(path) != ".md" {
				return nil
			}
			embedExamplesInMarkdown(root, path)
			return nil
		}); err != nil {
			logFatal(err, "Failed to walk Markdown directory %q", path)
		}
	}
}

func embedExamplesInMarkdown(root, path string) {
	contents, err := os.ReadFile(path) // #nosec G304,G703
	if err != nil {
		logFatal(err, "Failed to read Markdown file %q", path)
	}

	updated := replaceEmbeddedExamples(root, string(contents), path)
	if updated == string(contents) {
		return
	}
	if err = os.WriteFile(path, []byte(updated), 0o600); err != nil { // #nosec G703
		logFatal(err, "Failed to write Markdown file %q", path)
	}
}

func replaceEmbeddedExamples(root, markdown, markdownPath string) string {
	const (
		embedPrefix = "[//]: # (embed: "
		embedSuffix = ")"
		goFence     = "```go\n"
	)

	var builder strings.Builder
	cursor := 0
	for {
		directiveStart := strings.Index(markdown[cursor:], embedPrefix)
		if directiveStart < 0 {
			builder.WriteString(markdown[cursor:])
			return builder.String()
		}
		directiveStart += cursor
		lineEnd := strings.IndexByte(markdown[directiveStart:], '\n')
		if lineEnd < 0 {
			logFatal(nil, "Embed directive in %q is not followed by a code block", markdownPath)
		}
		lineEnd += directiveStart
		directive := strings.TrimSpace(markdown[directiveStart:lineEnd])
		if !strings.HasPrefix(directive, embedPrefix) || !strings.HasSuffix(directive, embedSuffix) {
			logFatal(nil, "Malformed embed directive %q in %q", directive, markdownPath)
		}
		exampleRef := strings.TrimSuffix(strings.TrimPrefix(directive, embedPrefix), embedSuffix)
		example := readEmbeddedExample(root, exampleRef)

		fenceStart := strings.Index(markdown[lineEnd:], goFence)
		if fenceStart < 0 {
			logFatal(nil, "Embed directive %q in %q is missing a Go code block", exampleRef, markdownPath)
		}
		fenceStart += lineEnd
		codeStart := fenceStart + len(goFence)
		fenceEnd := strings.Index(markdown[codeStart:], "\n```")
		if fenceEnd < 0 {
			logFatal(nil, "Embed directive %q in %q has an unterminated Go code block", exampleRef, markdownPath)
		}
		fenceEnd += codeStart

		builder.WriteString(markdown[cursor:codeStart])
		builder.WriteString(strings.TrimRight(example, "\n"))
		cursor = fenceEnd
	}
}

func readEmbeddedExample(root, exampleRef string) string {
	sourcePath, functionName, hasFunctionName := strings.Cut(exampleRef, "#")
	if hasFunctionName {
		return readEmbeddedFunction(root, sourcePath, functionName)
	}
	if strings.HasPrefix(exampleRef, "Example") {
		return readEmbeddedFunctionByName(root, exampleRef)
	}

	contents, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(sourcePath))) // #nosec G304,G703
	if err != nil {
		logFatal(err, "Failed to read embedded example %q", exampleRef)
	}
	return string(contents)
}

func readEmbeddedFunctionByName(root, functionName string) string {
	var matches []string
	if err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error { // #nosec G703
		if err != nil {
			return err
		}
		if entry.IsDir() || filepath.Base(path) != "example_test.go" {
			return nil
		}
		if hasEmbeddedFunction(path, functionName) {
			matches = append(matches, path)
		}
		return nil
	}); err != nil {
		logFatal(err, "Failed to find embedded example %q", functionName)
	}

	switch len(matches) {
	case 0:
		logFatal(nil, "Function %q was not found in example_test.go files", functionName)
	case 1:
		relPath, err := filepath.Rel(root, matches[0])
		if err != nil {
			logFatal(err, "Failed to resolve embedded example source %q", matches[0])
		}
		return readEmbeddedFunction(root, filepath.ToSlash(relPath), functionName)
	default:
		logFatal(
			nil,
			"Function %q is ambiguous across example_test.go files: %s",
			functionName,
			strings.Join(matches, ", "),
		)
	}
	return ""
}

func hasEmbeddedFunction(path, functionName string) bool {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
	if err != nil {
		logFatal(err, "Failed to parse embedded example source %q", path)
	}
	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if ok && funcDecl.Name.Name == functionName {
			return true
		}
	}
	return false
}

func readEmbeddedFunction(root, sourcePath, functionName string) string {
	path := filepath.Join(root, filepath.FromSlash(sourcePath))
	source, err := os.ReadFile(path) // #nosec G304,G703
	if err != nil {
		logFatal(err, "Failed to read embedded example source %q", path)
	}

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, path, source, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		logFatal(err, "Failed to parse embedded example source %q", path)
	}

	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok || funcDecl.Name.Name != functionName {
			continue
		}
		startPos := funcDecl.Pos()
		if funcDecl.Doc != nil {
			startPos = funcDecl.Doc.Pos()
		}
		start := fset.PositionFor(startPos, false)
		end := fset.PositionFor(funcDecl.End(), false)
		return strings.TrimSpace(string(source[start.Offset:end.Offset]))
	}
	logFatal(nil, "Function %q was not found in embedded example source %q", functionName, sourcePath)
	return ""
}

func logFatal(err error, msg string, a ...any) {
	var attrs []slog.Attr
	if err != nil {
		attrs = append(attrs, slog.String("error", err.Error()))
	}
	logging.Logger().LogAttrs(context.Background(), slog.LevelError, fmt.Sprintf(msg, a...), attrs...)
	os.Exit(1)
}
