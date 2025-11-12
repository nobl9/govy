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

var FunctionsWithGetter = []string{
	"For",
	"ForPointer",
	"Transform",
	"ForSlice",
	"ForMap",
}

func InferName(file string, line int) string {
	parseModuleASTOnce()

	pkg, astFile := modAST.FindFile(file)
	if astFile == nil {
		return ""
	}
	return InferNameFromFile(modAST.FileSet, pkg, astFile, line)
}

func InferNameFromFile(fileSet *token.FileSet, pkg *packages.Package, f *ast.File, line int) string {
	var (
		getterNode                   ast.Node
		previousNodeIsFuncWithGetter bool
	)
	importName := GetGovyImportName(f)
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
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
			return false
		}
		switch v := n.(type) {
		case *ast.SelectorExpr:
			if se, isSelectorExpr := n.(*ast.SelectorExpr); isSelectorExpr {
				exprIdent, ok := se.X.(*ast.Ident)
				// FIXME: It's not safe to assume package name like that.
				if ok && exprIdent.Name == importName && slices.Contains(FunctionsWithGetter, se.Sel.Name) {
					previousNodeIsFuncWithGetter = true
					return false
				}
			}
		case *ast.Ident:
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

type nameFinder struct {
	pkg *packages.Package
}

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

// findStructTypeInStructField returns the underlying [*types.Struct] of [*ast.Field] if it's a struct.
func (n nameFinder) findStructTypeInStructField(field *types.Var) (*types.Struct, bool) {
	switch ut := field.Type().Underlying().(type) {
	case *types.Struct:
		return ut, true
	default:
		return nil, false
	}
}
