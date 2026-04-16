package inferpath

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/nobl9/govy/internal/logging"
	"github.com/nobl9/govy/pkg/jsonpath"
)

// FunctionsWithGetter is a list of govy package functions that accept a property
// getter as their first argument and thus qualify for property path inference.
// These functions follow the pattern: func(getter func(S) T) where S is the
// struct type being validated and T is the property value type.
var FunctionsWithGetter = []string{
	"For",
	"ForPointer",
	"Transform",
	"ForSlice",
	"ForMap",
}

// InferPath infers the relative property path for the govy getter constructor call
// recorded at the given file and line.
// The returned [jsonpath.Path] does not include a leading `$` segment.
func InferPath(file string, line int) jsonpath.Path {
	parseModuleASTOnce()

	pkg, astFile := modAST.FindFile(file)
	if astFile == nil {
		logging.Logger().Error(
			"AST file not found for path inference",
			"file", file,
			"line", line,
		)
		return jsonpath.Path{}
	}
	return InferPathFromFile(modAST.FileSet, pkg, astFile, line)
}

// InferPathFromFile infers the relative property path from a getter function at the given line.
// It traverses the AST to find a govy function call (For, ForPointer, etc.) and
// extracts the property path from the getter's return statement by analyzing
// struct field access patterns like `s.Field` or `s.Nested.Field`.
// The returned path is rooted at the getter parameter, not at `$`.
func InferPathFromFile(
	fileSet *token.FileSet,
	pkg *packages.Package,
	f *ast.File,
	line int,
) jsonpath.Path {
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

	finder := pathFinder{pkg: pkg}
	return finder.findPath(getterNode, nil)
}

func resolveFieldPathSegment(fieldName, tagValue string) string {
	tag := reflect.StructTag(strings.Trim(tagValue, "`"))
	for _, value := range []string{tag.Get("json"), tag.Get("yaml")} {
		name, _, _ := strings.Cut(value, ",")
		if name != "" {
			return name
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

type pathFinder struct {
	pkg *packages.Package
}

func (n pathFinder) findPath(a any, structType *types.Struct) jsonpath.Path {
	switch v := a.(type) {
	case *ast.SelectorExpr:
		path, _ := n.findPathInSelectorExpr(v, structType)
		return path
	case *ast.Ident:
		return n.findPathInIdent(v, structType)
	case *ast.AssignStmt:
		return n.findPathInAssignStmt(v, structType)
	case *ast.FuncLit:
		return n.findPathInFuncLit(v)
	case *ast.ReturnStmt:
		return n.findPathInReturnStmt(v, structType)
	case *ast.IfStmt:
		return n.findPathInIfStmt(v, structType)
	case *ast.BlockStmt:
		return n.findPathInBlockStmt(v, structType)
	default:
		logging.Logger().Debug(fmt.Sprintf("unexpected type: %T", v))
	}
	return jsonpath.Path{}
}

func (n pathFinder) findPathInBlockStmt(
	blockStmt *ast.BlockStmt,
	structType *types.Struct,
) jsonpath.Path {
	if blockStmt == nil {
		logging.Logger().Debug("*ast.BlockStmt is nil, failed to locate the getter function parent")
		return jsonpath.Path{}
	}
	for _, stmt := range blockStmt.List {
		if path := n.findPath(stmt, structType); !path.IsEmpty() {
			return path
		}
	}
	return jsonpath.Path{}
}

func (n pathFinder) findPathInIfStmt(
	ifStmt *ast.IfStmt,
	structType *types.Struct,
) jsonpath.Path {
	if ifStmt == nil {
		logging.Logger().Debug("*ast.IfStmt is nil, failed to locate the getter function parent")
		return jsonpath.Path{}
	}
	return n.findPath(ifStmt.Body, structType)
}

func (n pathFinder) findPathInFuncLit(fl *ast.FuncLit) jsonpath.Path {
	if fl == nil {
		logging.Logger().Debug("*ast.FuncLit is nil, failed to locate the getter function parent")
		return jsonpath.Path{}
	}
	paramsList := fl.Type.Params.List
	if len(paramsList) != 1 {
		logging.Logger().Debug("*ast.FuncLit must have exactly one parameter")
		return jsonpath.Path{}
	}
	paramIdent, ok := paramsList[0].Type.(*ast.Ident)
	if !ok {
		logging.Logger().Debug("parameter must be an identifier")
		return jsonpath.Path{}
	}
	object := n.pkg.TypesInfo.ObjectOf(paramIdent)
	if object == nil {
		logging.Logger().Debug("failed to locate the object for the parameter identifier")
		return jsonpath.Path{}
	}
	var structType *types.Struct
	switch ot := object.Type().(type) {
	case *types.Named:
		switch ut := ot.Underlying().(type) {
		case *types.Struct:
			structType = ut
		default:
			logging.Logger().Debug(fmt.Sprintf("unexpected type: %T", ut))
			return jsonpath.Path{}
		}
	default:
		logging.Logger().Debug(fmt.Sprintf("unexpected type: %T", ot))
		return jsonpath.Path{}
	}
	for _, stmt := range fl.Body.List {
		if path := n.findPath(stmt, structType); !path.IsEmpty() {
			return path
		}
	}
	return jsonpath.Path{}
}

func (n pathFinder) findPathInReturnStmt(
	returnStmt *ast.ReturnStmt,
	structType *types.Struct,
) jsonpath.Path {
	if returnStmt == nil {
		logging.Logger().Debug("no return statement found in getter function")
		return jsonpath.Path{}
	}
	if len(returnStmt.Results) != 1 {
		logging.Logger().Debug("return statement must have exactly one result")
		return jsonpath.Path{}
	}
	return n.findPath(returnStmt.Results[0], structType)
}

func (n pathFinder) findPathInIdent(
	ident *ast.Ident,
	structType *types.Struct,
) jsonpath.Path {
	if ident.Obj == nil {
		logging.Logger().Debug("identifier object is nil")
		return jsonpath.Path{}
	}
	return n.findPath(ident.Obj.Decl, structType)
}

func (n pathFinder) findPathInAssignStmt(
	assignment *ast.AssignStmt,
	structType *types.Struct,
) jsonpath.Path {
	if len(assignment.Rhs) != 1 {
		logging.Logger().Debug("assignment statement must have exactly one right-hand side")
		return jsonpath.Path{}
	}
	return n.findPath(assignment.Rhs[0], structType)
}

// findPathInSelectorExpr extracts a relative property path from a selector expression like
// `s.Field` or `s.Nested.Field`.
// For nested access, it recursively processes the chain and constructs a multi-segment path
// such as `nested.field`.
// It uses struct tags (json/yaml) to determine each named segment.
func (n pathFinder) findPathInSelectorExpr(
	se *ast.SelectorExpr,
	structType *types.Struct,
) (jsonpath.Path, *types.Struct) {
	var path jsonpath.Path
	switch v := se.X.(type) {
	case *ast.Ident:
		break
	case *ast.SelectorExpr:
		path, structType = n.findPathInSelectorExpr(v, structType)
	case *ast.IndexExpr:
		path, structType = n.findPathInIndexExpr(v, structType)
	default:
		logging.Logger().Debug(fmt.Sprintf("unexpected type: %T", v))
		return jsonpath.Path{}, nil
	}
	if structType == nil {
		return jsonpath.Path{}, nil
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
		fieldName = resolveFieldPathSegment(fieldName, tagValue)
		return path.Name(fieldName), structType
	}
	logging.Logger().Debug(fmt.Sprintf("field matching '%s' name not found in struct type", se.Sel.Name))
	return jsonpath.Path{}, nil
}

// findStructTypeInStructField returns the underlying [*types.Struct] of a field if it's a struct.
// For collection types (slice, array, map), it extracts the element type first.
func (n pathFinder) findStructTypeInStructField(field *types.Var) (*types.Struct, bool) {
	return n.getStructFromType(field.Type())
}

// getStructFromType extracts a struct type from a potentially wrapped type.
// It handles pointers, named types, and collection types (slices, arrays, maps).
func (n pathFinder) getStructFromType(t types.Type) (*types.Struct, bool) {
	if ptr, ok := t.(*types.Pointer); ok {
		t = ptr.Elem()
	}
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
	default:
		return nil, false
	}
}

// findPathInIndexExpr handles index expressions like p.Students[0] or p.Items[i].
// It extracts the indexed relative property path and appends the index or key segment.
func (n pathFinder) findPathInIndexExpr(
	ie *ast.IndexExpr,
	structType *types.Struct,
) (jsonpath.Path, *types.Struct) {
	var path jsonpath.Path
	switch v := ie.X.(type) {
	case *ast.SelectorExpr:
		path, structType = n.findPathInSelectorExpr(v, structType)
	case *ast.IndexExpr:
		path, structType = n.findPathInIndexExpr(v, structType)
	case *ast.Ident:
		logging.Logger().Debug(fmt.Sprintf("index expression base is identifier: %s", v.Name))
	default:
		logging.Logger().Debug(fmt.Sprintf("unexpected type in IndexExpr.X: %T", v))
		return jsonpath.Path{}, nil
	}
	if path.IsEmpty() && structType == nil {
		return jsonpath.Path{}, nil
	}
	return n.appendIndexExpr(path, ie.Index), structType
}

// appendIndexExpr appends an index expression to the path.
// For integer literals it uses [Path.Index], for string literals [Path.Key],
// and for non-literal expressions it falls back to a raw "[]" segment.
func (n pathFinder) appendIndexExpr(path jsonpath.Path, index ast.Expr) jsonpath.Path {
	lit, ok := index.(*ast.BasicLit)
	if !ok {
		logging.Logger().Debug(fmt.Sprintf("non-literal index expression: %T", index))
		return path.UnknownIndex()
	}
	switch lit.Kind {
	case token.INT:
		v, err := strconv.ParseUint(lit.Value, 10, 64)
		if err != nil {
			logging.Logger().Debug(fmt.Sprintf("failed to parse integer literal: %v", err))
			return path.UnknownIndex()
		}
		return path.Index(uint(v))
	case token.STRING:
		key, err := strconv.Unquote(lit.Value)
		if err != nil {
			logging.Logger().Debug(fmt.Sprintf("failed to unquote string literal: %v", err))
			return path.UnknownIndex()
		}
		return path.Key(key)
	default:
		logging.Logger().Debug(fmt.Sprintf("unhandled BasicLit kind in index expression: %v", lit.Kind))
	}
	return path.UnknownIndex()
}
