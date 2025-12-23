package infername

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"slices"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/nobl9/govy/internal/logging"
)

// FunctionsWithGetter is a list of govy package functions that accept a property
// getter as their first argument and thus qualify for property name inference.
// These functions follow the pattern: func(getter func(S) T) where S is the
// struct type being validated and T is the property value type.
var FunctionsWithGetter = []string{
	"For",
	"ForPointer",
	"Transform",
	"ForSlice",
	"ForMap",
}

// InferName infers the name of a field from a file and line number.
func InferName(file string, line int) string {
	parseModuleASTOnce()

	pkg, astFile := modAST.FindFile(file)
	if astFile == nil {
		logging.Logger().Error(
			"AST file not found for name inference",
			"file", file,
			"line", line,
		)
		return ""
	}
	return InferNameFromFile(modAST.FileSet, pkg, astFile, line)
}

// InferNameFromFile infers the property name from a getter function at the given line.
// It traverses the AST to find a govy function call (For, ForPointer, etc.) and
// extracts the property name from the getter's return statement by analyzing
// struct field access patterns like `s.Field` or `s.Nested.Field`.
// Returns an empty string if inference fails.
func InferNameFromFile(fileSet *token.FileSet, pkg *packages.Package, f *ast.File, line int) string {
	var (
		getterNode                   ast.Node
		previousNodeIsFuncWithGetter bool
		// done is used to stop processing nodes after we find the getter function.
		done bool
	)
	importName := GetGovyImportName(f)
	ast.Inspect(f, func(n ast.Node) bool {
		if done || n == nil {
			return false
		}
		nodeLine := fileSet.Position(n.Pos()).Line
		if nodeLine > line {
			return false
		}
		if nodeLine != line {
			return true
		}
		// What follows must be the getter function.
		if previousNodeIsFuncWithGetter {
			getterNode = n
			done = true // Stop processing more nodes
			return false
		}
		switch v := n.(type) {
		case *ast.SelectorExpr:
			// Check if field selector is a govy function with getter.
			if !slices.Contains(FunctionsWithGetter, v.Sel.Name) {
				return true
			}
			// Check if expression is the govy package.
			exprIdent, ok := v.X.(*ast.Ident)
			if ok && exprIdent.Name == importName {
				previousNodeIsFuncWithGetter = true
			}
			return false
		case *ast.Ident:
			// This case is ONLY for dot imports: `import . "govy"`.
			// When using dot import, For() is called without package prefix.
			if slices.Contains(FunctionsWithGetter, v.Name) {
				previousNodeIsFuncWithGetter = true
				return false
			}
		}
		return true
	})

	finder := nameFinder{pkg: pkg}
	return finder.FindName(getterNode, nil)
}

// InferNameDefaultFunc is the default function for inferring field names from struct tags.
// It looks for json and yaml tags, preferring json if both are set.
func InferNameDefaultFunc(fieldName, tagValue string) string {
	for _, tagKey := range []string{"json", "yaml"} {
		tagValues := strings.Split(
			reflect.StructTag(strings.Trim(tagValue, "`")).Get(tagKey),
			",",
		)
		if len(tagValues) > 0 && tagValues[0] != "" {
			fieldName = tagValues[0]
			break
		}
	}
	return fieldName
}

// GetGovyImportName returns the import alias used for the govy package in the file.
// It handles aliased imports like `import v "github.com/nobl9/govy/pkg/govy"`,
// dot imports, and defaults to "govy" if no alias is specified.
func GetGovyImportName(f *ast.File) string {
	importName := "govy"
	for _, imp := range f.Imports {
		if imp.Path.Value == `"github.com/nobl9/govy/pkg/govy"` && imp.Name != nil {
			importName = imp.Name.Name
			break
		}
	}
	return importName
}

// nameFinder is a helper struct for finding the name of an inferred field.
type nameFinder struct {
	pkg *packages.Package
}

// FindName recursively traverses AST nodes to find and construct the property name.
// It dispatches to type-specific handlers based on the AST node type.
// The structType parameter carries the current struct context for nested field lookups.
func (n nameFinder) FindName(a any, structType *types.Struct) string {
	switch v := a.(type) {
	case *ast.SelectorExpr:
		name, _ := n.findNameInSelectorExpr(v, structType)
		return name
	case *ast.Ident:
		return n.findNameInIdent(v, structType)
	case *ast.AssignStmt:
		return n.findNameInAssignStmt(v, structType)
	case *ast.FuncLit:
		return n.findNameInFuncLit(v)
	case *ast.ReturnStmt:
		return n.findNameInReturnStmt(v, structType)
	case *ast.IfStmt:
		return n.findNameInIfStmt(v, structType)
	case *ast.BlockStmt:
		return n.findNameInBlockStmt(v, structType)
	default:
		logging.Logger().Debug(fmt.Sprintf("unexpected type: %T", v))
	}
	return ""
}

func (n nameFinder) findNameInBlockStmt(blockStmt *ast.BlockStmt, structType *types.Struct) string {
	if blockStmt == nil {
		logging.Logger().Debug("*ast.BlockStmt is nil, failed to locate the getter function parent")
		return ""
	}
	for _, stmt := range blockStmt.List {
		if name := n.FindName(stmt, structType); name != "" {
			return name
		}
	}
	return ""
}

func (n nameFinder) findNameInIfStmt(ifStmt *ast.IfStmt, structType *types.Struct) string {
	if ifStmt == nil {
		logging.Logger().Debug("*ast.IfStmt is nil, failed to locate the getter function parent")
		return ""
	}
	return n.FindName(ifStmt.Body, structType)
}

// findNameInFuncLit returns the name of the property that the getter function is supposed to return.
// It attempts to find the return statement which lets us infer the name,
// until it succeeds or there are no more return statements to inspect.
func (n nameFinder) findNameInFuncLit(fl *ast.FuncLit) string {
	if fl == nil {
		logging.Logger().Debug("*ast.FuncLit is nil, failed to locate the getter function parent")
		return ""
	}
	paramsList := fl.Type.Params.List
	if len(paramsList) != 1 {
		logging.Logger().Debug("*ast.FuncLit must have exactly one parameter")
		return ""
	}
	paramIdent, ok := paramsList[0].Type.(*ast.Ident)
	if !ok {
		logging.Logger().Debug("parameter must be an identifier")
		return ""
	}
	object := n.pkg.TypesInfo.ObjectOf(paramIdent)
	if object == nil {
		logging.Logger().Debug("failed to locate the object for the parameter identifier")
		return ""
	}
	var structType *types.Struct
	switch ot := object.Type().(type) {
	case *types.Named:
		switch ut := ot.Underlying().(type) {
		case *types.Struct:
			structType = ut
		default:
			logging.Logger().Debug(fmt.Sprintf("unexpected type: %T", ut))
			return ""
		}
	default:
		logging.Logger().Debug(fmt.Sprintf("unexpected type: %T", ot))
		return ""
	}
	for _, stmt := range fl.Body.List {
		if name := n.FindName(stmt, structType); name != "" {
			return name
		}
	}
	return ""
}

func (n nameFinder) findNameInReturnStmt(returnStmt *ast.ReturnStmt, structType *types.Struct) string {
	if returnStmt == nil {
		logging.Logger().Debug("no return statement found in getter function")
		return ""
	}
	if len(returnStmt.Results) != 1 {
		logging.Logger().Debug("return statement must have exactly one result")
		return ""
	}
	return n.FindName(returnStmt.Results[0], structType)
}

func (n nameFinder) findNameInIdent(ident *ast.Ident, structType *types.Struct) string {
	if ident.Obj == nil {
		logging.Logger().Debug("identifier object is nil")
		return ""
	}
	return n.FindName(ident.Obj.Decl, structType)
}

func (n nameFinder) findNameInAssignStmt(assignment *ast.AssignStmt, structType *types.Struct) string {
	if len(assignment.Rhs) != 1 {
		logging.Logger().Debug("assignment statement must have exactly one right-hand side")
		return ""
	}
	return n.FindName(assignment.Rhs[0], structType)
}

// findNameInSelectorExpr extracts the property name from a selector expression like
// `s.Field` or `s.Nested.Field`. For nested access, it recursively processes the chain
// and constructs a dot-separated path (e.g., "nested.field"). It uses struct tags
// (json/yaml) to determine the final field name via [InferNameDefaultFunc].
func (n nameFinder) findNameInSelectorExpr(
	se *ast.SelectorExpr,
	structType *types.Struct,
) (string, *types.Struct) {
	var name string
	switch v := se.X.(type) {
	case *ast.Ident:
		break
	case *ast.SelectorExpr:
		name, structType = n.findNameInSelectorExpr(v, structType)
	case *ast.IndexExpr:
		name, structType = n.findNameInIndexExpr(v, structType)
	default:
		logging.Logger().Debug(fmt.Sprintf("unexpected type: %T", v))
		return "", nil
	}
	if structType == nil {
		return "", nil
	}
	for i := range structType.NumFields() {
		field := structType.Field(i)
		fieldName := field.Name()
		if fieldName != se.Sel.Name {
			continue
		}
		tagValue := structType.Tag(i)
		if childStructType, isStruct := n.findStructTypeInStructField(field); isStruct {
			structType = childStructType
		}
		fieldName = InferNameDefaultFunc(fieldName, tagValue)
		if name == "" {
			return fieldName, structType
		}
		return name + "." + fieldName, structType
	}
	logging.Logger().Debug(fmt.Sprintf("field matching '%s' name not found in struct type", se.Sel.Name))
	return "", nil
}

// findStructTypeInStructField returns the underlying [*types.Struct] of a field if it's a struct.
// For collection types (slice, array, map), it extracts the element type first.
func (n nameFinder) findStructTypeInStructField(field *types.Var) (*types.Struct, bool) {
	return n.getStructFromType(field.Type())
}

// getStructFromType extracts a struct type from a potentially wrapped type.
// It handles pointers, named types, and collection types (slices, arrays, maps).
func (n nameFinder) getStructFromType(t types.Type) (*types.Struct, bool) {
	// Unwrap pointer.
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}
	// Handle named types.
	if named, ok := t.(*types.Named); ok {
		t = named.Underlying()
	}
	switch ut := t.(type) {
	case *types.Struct:
		return ut, true
	case *types.Slice:
		return n.getStructFromType(ut.Elem())
	case *types.Array:
		return n.getStructFromType(ut.Elem())
	case *types.Map:
		return n.getStructFromType(ut.Elem())
	}
	return nil, false
}

// formatIndexExpr formats an index expression for the property path.
// Returns [0] for integer literals, ["key"] for string literals,
// and [] for variables or other expressions.
func (n nameFinder) formatIndexExpr(index ast.Expr) string {
	if lit, ok := index.(*ast.BasicLit); ok {
		switch lit.Kind {
		case token.INT:
			return "[" + lit.Value + "]"
		case token.STRING:
			// String literals already include quotes.
			return "[" + lit.Value + "]"
		}
	}
	// For any other expression (variables, function calls, etc.)
	return "[]"
}

// findNameInIndexExpr handles index expressions like p.Students[0] or p.Items[i].
// It extracts the indexed property name and appends the index notation.
func (n nameFinder) findNameInIndexExpr(
	ie *ast.IndexExpr,
	structType *types.Struct,
) (string, *types.Struct) {
	var name string
	// Process the base expression (ie.X).
	switch v := ie.X.(type) {
	case *ast.SelectorExpr:
		name, structType = n.findNameInSelectorExpr(v, structType)
	case *ast.IndexExpr:
		// Handle nested indices like matrix[0][1].
		name, structType = n.findNameInIndexExpr(v, structType)
	case *ast.Ident:
		// Base case - just an identifier being indexed.
		break
	default:
		logging.Logger().Debug(fmt.Sprintf("unexpected type in IndexExpr.X: %T", v))
		return "", nil
	}
	if name == "" && structType == nil {
		return "", nil
	}
	// Format the index notation.
	indexStr := n.formatIndexExpr(ie.Index)
	return name + indexStr, structType
}
