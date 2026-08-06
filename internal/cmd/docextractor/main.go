package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
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

	resolver := newEmbeddedExampleResolver(root)
	for _, path := range paths {
		info, err := os.Stat(path) // #nosec G703
		if err != nil {
			logFatal(err, "Failed to stat path %q", path)
		}
		if !info.IsDir() {
			embedExamplesInMarkdown(resolver, path)
			continue
		}
		if err = filepath.WalkDir(path, func(path string, entry fs.DirEntry, err error) error { // #nosec G703
			if err != nil {
				return err
			}
			if entry.IsDir() || filepath.Ext(path) != ".md" {
				return nil
			}
			embedExamplesInMarkdown(resolver, path)
			return nil
		}); err != nil {
			logFatal(err, "Failed to walk Markdown directory %q", path)
		}
	}
}

func embedExamplesInMarkdown(resolver *embeddedExampleResolver, path string) {
	contents, err := os.ReadFile(path) // #nosec G304,G703
	if err != nil {
		logFatal(err, "Failed to read Markdown file %q", path)
	}

	original := string(contents)
	updated := replaceEmbeddedExamples(resolver, original, path)
	updated = replaceGeneratedDocs(resolver.root, updated, path)
	if updated == original {
		return
	}
	if err = os.WriteFile(path, []byte(updated), 0o600); err != nil { // #nosec G703
		logFatal(err, "Failed to write Markdown file %q", path)
	}
}

func replaceGeneratedDocs(root, markdown, markdownPath string) string {
	const (
		docsPrefix = "[//]: # (docs: "
		docsSuffix = ")"
		docsEnd    = "[//]: # (end-docs)"
	)

	if !strings.Contains(markdown, docsPrefix) {
		return markdown
	}

	var builder strings.Builder
	cursor := 0
	for {
		directiveStart := strings.Index(markdown[cursor:], docsPrefix)
		if directiveStart < 0 {
			builder.WriteString(markdown[cursor:])
			return builder.String()
		}
		directiveStart += cursor
		lineEnd := strings.IndexByte(markdown[directiveStart:], '\n')
		if lineEnd < 0 {
			logFatal(nil, "Docs directive in %q is not followed by an end directive", markdownPath)
		}
		lineEnd += directiveStart
		directive := strings.TrimSpace(markdown[directiveStart:lineEnd])
		if !strings.HasPrefix(directive, docsPrefix) || !strings.HasSuffix(directive, docsSuffix) {
			logFatal(nil, "Malformed docs directive %q in %q", directive, markdownPath)
		}
		docsRef := strings.TrimSuffix(strings.TrimPrefix(directive, docsPrefix), docsSuffix)
		docs := renderGeneratedDocs(root, docsRef)

		endStart := strings.Index(markdown[lineEnd:], docsEnd)
		if endStart < 0 {
			logFatal(nil, "Docs directive %q in %q is missing %q", docsRef, markdownPath, docsEnd)
		}
		endStart += lineEnd
		endLineEnd := strings.IndexByte(markdown[endStart:], '\n')
		if endLineEnd < 0 {
			endLineEnd = len(markdown)
		} else {
			endLineEnd += endStart
		}

		builder.WriteString(markdown[cursor : lineEnd+1])
		builder.WriteByte('\n')
		builder.WriteString(strings.TrimRight(docs, "\n"))
		builder.WriteString("\n\n")
		builder.WriteString(markdown[endStart:endLineEnd])
		cursor = endLineEnd
	}
}

type generatedDocsConfig struct {
	path    string
	kind    string
	returns []string
}

type documentedFunction struct {
	name        string
	description string
}

func renderGeneratedDocs(root, docsRef string) string {
	config := parseGeneratedDocsConfig(docsRef)
	if config.kind != "func" {
		logFatal(nil, "Unsupported docs kind %q", config.kind)
	}

	path := filepath.Join(root, filepath.FromSlash(config.path))
	var files []string
	info, err := os.Stat(path) // #nosec G703
	if err != nil {
		logFatal(err, "Failed to stat generated docs path %q", path)
	}
	if info.IsDir() {
		if err = filepath.WalkDir(path, func(path string, entry fs.DirEntry, err error) error { // #nosec G703
			if err != nil {
				return err
			}
			if entry.IsDir() || filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
				return nil
			}
			files = append(files, path)
			return nil
		}); err != nil {
			logFatal(err, "Failed to walk generated docs path %q", path)
		}
	} else {
		files = append(files, path)
	}
	slices.Sort(files)

	functions := make([]documentedFunction, 0, len(files))
	for _, file := range files {
		functions = append(functions, readDocumentedFunctions(file, config)...)
	}
	slices.SortFunc(functions, func(a, b documentedFunction) int {
		return strings.Compare(a.name, b.name)
	})

	var builder strings.Builder
	for _, function := range functions {
		builder.WriteString("- `")
		builder.WriteString(function.name)
		builder.WriteString("` - ")
		builder.WriteString(function.description)
		builder.WriteByte('\n')
	}
	return builder.String()
}

func parseGeneratedDocsConfig(docsRef string) generatedDocsConfig {
	path, query, _ := strings.Cut(docsRef, "?")
	config := generatedDocsConfig{path: path, kind: "func"}
	for part := range strings.SplitSeq(query, "&") {
		key, value, ok := strings.Cut(part, "=")
		if !ok {
			continue
		}
		switch key {
		case "kind":
			config.kind = value
		case "returns":
			config.returns = strings.Split(value, ",")
		default:
			logFatal(nil, "Unsupported docs query parameter %q", key)
		}
	}
	return config
}

func readDocumentedFunctions(path string, config generatedDocsConfig) []documentedFunction {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, path, nil, parser.ParseComments|parser.SkipObjectResolution)
	if err != nil {
		logFatal(err, "Failed to parse generated docs source %q", path)
	}

	var functions []documentedFunction
	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok || !funcDecl.Name.IsExported() || !matchesReturnFilter(fset, funcDecl, config.returns) {
			continue
		}
		if funcDecl.Doc == nil {
			logFatal(nil, "Exported function %q in %q is missing documentation", funcDecl.Name.Name, path)
		}
		functions = append(functions, documentedFunction{
			name:        funcDecl.Name.Name,
			description: firstDocSentence(funcDecl.Doc),
		})
	}
	return functions
}

func firstDocSentence(doc *ast.CommentGroup) string {
	lines := make([]string, 0, len(doc.List))
	for _, comment := range doc.List {
		line := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	return firstSentence(strings.Join(lines, " "))
}

func firstSentence(text string) string {
	bracketDepth := 0
	for i, r := range text {
		switch r {
		case '[':
			bracketDepth++
		case ']':
			if bracketDepth > 0 {
				bracketDepth--
			}
		case '.':
			if bracketDepth > 0 || isKnownAbbreviation(text[:i+1]) {
				continue
			}
			if i == len(text)-1 || text[i+1] == ' ' || text[i+1] == '\n' || text[i+1] == '\t' {
				return strings.TrimSpace(text[:i+1])
			}
		}
	}
	return strings.TrimSpace(text)
}

func isKnownAbbreviation(text string) bool {
	lower := strings.ToLower(text)
	return strings.HasSuffix(lower, "i.e.") || strings.HasSuffix(lower, "e.g.")
}

func matchesReturnFilter(fset *token.FileSet, funcDecl *ast.FuncDecl, returns []string) bool {
	if len(returns) == 0 {
		return true
	}
	if funcDecl.Type.Results == nil {
		return false
	}
	for _, result := range funcDecl.Type.Results.List {
		var builder strings.Builder
		if err := printer.Fprint(&builder, fset, result.Type); err != nil {
			logFatal(err, "Failed to print return type for function %q", funcDecl.Name.Name)
		}
		returnType := builder.String()
		for _, expected := range returns {
			if strings.HasPrefix(returnType, expected+"[") || returnType == expected {
				return true
			}
		}
	}
	return false
}

func replaceEmbeddedExamples(
	resolver *embeddedExampleResolver,
	markdown,
	markdownPath string,
) string {
	const (
		embedPrefix = "[//]: # (embed: "
		embedSuffix = ")"
		goFence     = "```go\n"
	)

	if !strings.Contains(markdown, embedPrefix) {
		return markdown
	}

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
		example := resolver.read(exampleRef)

		fenceStart := strings.Index(markdown[lineEnd:], goFence)
		if fenceStart < 0 {
			logFatal(nil, "Embed directive %q in %q is missing a Go code block", exampleRef, markdownPath)
		}
		fenceStart += lineEnd
		codeStart := fenceStart + len(goFence)
		fenceEnd := findCodeFenceEnd(markdown[codeStart:])
		if fenceEnd < 0 {
			logFatal(nil, "Embed directive %q in %q has an unterminated Go code block", exampleRef, markdownPath)
		}
		fenceEnd += codeStart

		builder.WriteString(markdown[cursor:codeStart])
		builder.WriteString(strings.TrimRight(example, "\n"))
		if markdown[fenceEnd] != '\n' {
			builder.WriteByte('\n')
		}
		cursor = fenceEnd
	}
}

func findCodeFenceEnd(markdown string) int {
	if strings.HasPrefix(markdown, "```") {
		return 0
	}
	if fenceEnd := strings.Index(markdown, "\n```"); fenceEnd >= 0 {
		return fenceEnd
	}
	return strings.Index(markdown, "```")
}

type embeddedExampleResolver struct {
	root                string
	files               map[string]*embeddedExampleFile
	functions           map[string][]*embeddedExampleFile
	exampleFilesIndexed bool
}

type embeddedExampleFile struct {
	path      string
	contents  string
	functions map[string]embeddedFunctionSpan
	parsed    bool
	indexed   bool
}

type embeddedFunctionSpan struct {
	start int
	end   int
}

func newEmbeddedExampleResolver(root string) *embeddedExampleResolver {
	return &embeddedExampleResolver{
		root:      filepath.Clean(root),
		files:     make(map[string]*embeddedExampleFile),
		functions: make(map[string][]*embeddedExampleFile),
	}
}

func readEmbeddedExample(root, exampleRef string) string {
	return newEmbeddedExampleResolver(root).read(exampleRef)
}

func (r *embeddedExampleResolver) read(exampleRef string) string {
	sourcePath, functionName, hasFunctionName := strings.Cut(exampleRef, "#")
	if hasFunctionName {
		return r.readFunction(sourcePath, functionName)
	}
	if strings.HasPrefix(exampleRef, "Example") {
		return r.readFunctionByName(exampleRef)
	}
	return r.readFile(sourcePath, exampleRef)
}

func (r *embeddedExampleResolver) readFunctionByName(functionName string) string {
	r.indexExampleFiles(functionName)
	matches := r.functions[functionName]
	switch len(matches) {
	case 0:
		logFatal(nil, "Function %q was not found in example_test.go files", functionName)
	case 1:
		return matches[0].function(functionName)
	default:
		paths := make([]string, 0, len(matches))
		for _, match := range matches {
			paths = append(paths, match.path)
		}
		logFatal(
			nil,
			"Function %q is ambiguous across example_test.go files: %s",
			functionName,
			strings.Join(paths, ", "),
		)
	}
	return ""
}

func (r *embeddedExampleResolver) readFunction(sourcePath, functionName string) string {
	path := filepath.Join(r.root, filepath.FromSlash(sourcePath))
	file := r.loadFile(path)
	r.parseFile(file)
	if _, ok := file.functions[functionName]; !ok {
		logFatal(nil, "Function %q was not found in embedded example source %q", functionName, sourcePath)
	}
	return file.function(functionName)
}

func (r *embeddedExampleResolver) readFile(sourcePath, exampleRef string) string {
	path := filepath.Join(r.root, filepath.FromSlash(sourcePath))
	if file, ok := r.files[path]; ok {
		return file.contents
	}
	contents, err := os.ReadFile(path) // #nosec G304,G703
	if err != nil {
		logFatal(err, "Failed to read embedded example %q", exampleRef)
	}
	file := &embeddedExampleFile{path: path, contents: string(contents)}
	r.files[path] = file
	return file.contents
}

func (r *embeddedExampleResolver) indexExampleFiles(functionName string) {
	if r.exampleFilesIndexed {
		return
	}
	if err := filepath.WalkDir(r.root, func(path string, entry fs.DirEntry, err error) error { // #nosec G703
		if err != nil {
			return err
		}
		if entry.IsDir() {
			if path != r.root && skipExampleSearchDir(r.root, path) {
				return fs.SkipDir
			}
			return nil
		}
		if filepath.Base(path) != "example_test.go" {
			return nil
		}
		file := r.loadFile(path)
		r.parseFile(file)
		r.indexFile(file)
		return nil
	}); err != nil {
		logFatal(err, "Failed to find embedded example %q", functionName)
	}
	for _, matches := range r.functions {
		slices.SortFunc(matches, func(a, b *embeddedExampleFile) int {
			return strings.Compare(a.path, b.path)
		})
	}
	r.exampleFilesIndexed = true
}

func (r *embeddedExampleResolver) loadFile(path string) *embeddedExampleFile {
	path = filepath.Clean(path)
	if file, ok := r.files[path]; ok {
		return file
	}
	contents, err := os.ReadFile(path) // #nosec G304,G703
	if err != nil {
		logFatal(err, "Failed to read embedded example source %q", path)
	}
	file := &embeddedExampleFile{path: path, contents: string(contents)}
	r.files[path] = file
	return file
}

func (r *embeddedExampleResolver) parseFile(file *embeddedExampleFile) {
	if file.parsed {
		return
	}
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(
		fset,
		file.path,
		file.contents,
		parser.ParseComments|parser.SkipObjectResolution,
	)
	if err != nil {
		logFatal(err, "Failed to parse embedded example source %q", file.path)
	}
	file.functions = make(map[string]embeddedFunctionSpan)
	for _, decl := range astFile.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if _, exists := file.functions[funcDecl.Name.Name]; exists {
			continue
		}
		startPos := funcDecl.Pos()
		if funcDecl.Doc != nil {
			startPos = funcDecl.Doc.Pos()
		}
		start := fset.PositionFor(startPos, false)
		end := fset.PositionFor(funcDecl.End(), false)
		file.functions[funcDecl.Name.Name] = embeddedFunctionSpan{
			start: start.Offset,
			end:   end.Offset,
		}
	}
	file.parsed = true
}

func (r *embeddedExampleResolver) indexFile(file *embeddedExampleFile) {
	if file.indexed {
		return
	}
	for functionName := range file.functions {
		r.functions[functionName] = append(r.functions[functionName], file)
	}
	file.indexed = true
}

func (f *embeddedExampleFile) function(functionName string) string {
	span := f.functions[functionName]
	return strings.TrimSpace(f.contents[span.start:span.end])
}

func skipExampleSearchDir(root, path string) bool {
	switch filepath.Base(path) {
	case ".git", ".hg", ".svn", ".worktrees",
		"vendor", "node_modules":
		return true
	case "bin", "build", "dist", "out", "target":
		return filepath.Dir(path) == root
	default:
		return false
	}
}

func logFatal(err error, msg string, a ...any) {
	var attrs []slog.Attr
	if err != nil {
		attrs = append(attrs, slog.String("error", err.Error()))
	}
	logging.Logger().LogAttrs(context.Background(), slog.LevelError, fmt.Sprintf(msg, a...), attrs...)
	os.Exit(1)
}
