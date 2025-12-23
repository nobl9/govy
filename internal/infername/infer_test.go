package infername

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"

	"github.com/nobl9/govy/internal/assert"
)

func TestInferNameDefaultFunc(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		tagValue  string
		expected  string
	}{
		{
			name:      "no tags returns field name",
			fieldName: "MyField",
			tagValue:  "",
			expected:  "MyField",
		},
		{
			name:      "json tag overrides field name",
			fieldName: "MyField",
			tagValue:  "`json:\"my_field\"`",
			expected:  "my_field",
		},
		{
			name:      "yaml tag overrides field name when json absent",
			fieldName: "MyField",
			tagValue:  "`yaml:\"my_yaml_field\"`",
			expected:  "my_yaml_field",
		},
		{
			name:      "json tag takes precedence over yaml",
			fieldName: "MyField",
			tagValue:  "`json:\"json_name\" yaml:\"yaml_name\"`",
			expected:  "json_name",
		},
		{
			name:      "json tag with options",
			fieldName: "MyField",
			tagValue:  "`json:\"name,omitempty\"`",
			expected:  "name",
		},
		{
			name:      "json tag with only options uses field name",
			fieldName: "MyField",
			tagValue:  "`json:\",omitempty\"`",
			expected:  "MyField",
		},
		{
			name:      "yaml tag with options",
			fieldName: "MyField",
			tagValue:  "`yaml:\"name,flow\"`",
			expected:  "name",
		},
		{
			name:      "empty json tag uses field name",
			fieldName: "MyField",
			tagValue:  "`json:\"\"`",
			expected:  "MyField",
		},
		{
			name:      "unrelated tags use field name",
			fieldName: "MyField",
			tagValue:  "`db:\"column_name\"`",
			expected:  "MyField",
		},
		{
			name:      "hyphen json tag returns hyphen",
			fieldName: "MyField",
			tagValue:  "`json:\"-\"`",
			expected:  "-",
		},
		{
			name:      "json tag with multiple commas",
			fieldName: "MyField",
			tagValue:  "`json:\"field_name,omitempty,string\"`",
			expected:  "field_name",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := InferNameDefaultFunc(tc.fieldName, tc.tagValue)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetGovyImportName(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		expected string
	}{
		{
			name: "default import name",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"
`,
			expected: "govy",
		},
		{
			name: "aliased import",
			src: `package test
import v "github.com/nobl9/govy/pkg/govy"
`,
			expected: "v",
		},
		{
			name: "custom alias import",
			src: `package test
import validation "github.com/nobl9/govy/pkg/govy"
`,
			expected: "validation",
		},
		{
			name: "dot import",
			src: `package test
import . "github.com/nobl9/govy/pkg/govy"
`,
			expected: ".",
		},
		{
			name: "underscore import",
			src: `package test
import _ "github.com/nobl9/govy/pkg/govy"
`,
			expected: "_",
		},
		{
			name: "no govy import returns default",
			src: `package test
import "fmt"
`,
			expected: "govy",
		},
		{
			name: "multiple imports with govy",
			src: `package test
import (
	"fmt"
	mygovy "github.com/nobl9/govy/pkg/govy"
	"strings"
)
`,
			expected: "mygovy",
		},
		{
			name: "govy import without alias in import block",
			src: `package test
import (
	"fmt"
	"github.com/nobl9/govy/pkg/govy"
)
`,
			expected: "govy",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "test.go", tc.src, parser.ImportsOnly)
			assert.Require(t, assert.NoError(t, err))

			result := GetGovyImportName(f)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFunctionsWithGetter(t *testing.T) {
	expected := []string{"For", "ForPointer", "Transform", "ForSlice", "ForMap"}
	assert.ElementsMatch(t, expected, FunctionsWithGetter)
}

func TestInferNameFromFile(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		line     int
		expected string
	}{
		{
			name: "simple struct field access",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string ` + "`json:\"name\"`" + `
}

var _ = govy.For(func(p Person) string { return p.Name })
`,
			line:     8,
			expected: "name",
		},
		{
			name: "nested struct field access",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Address struct {
	Street string ` + "`json:\"street\"`" + `
}

type Person struct {
	Addr Address ` + "`json:\"address\"`" + `
}

var _ = govy.For(func(p Person) string { return p.Addr.Street })
`,
			line:     12,
			expected: "address.street",
		},
		{
			name: "field without json tag uses field name",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string
}

var _ = govy.For(func(p Person) string { return p.Name })
`,
			line:     8,
			expected: "Name",
		},
		{
			name: "ForPointer function",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name *string ` + "`json:\"name\"`" + `
}

var _ = govy.ForPointer(func(p Person) *string { return p.Name })
`,
			line:     8,
			expected: "name",
		},
		{
			name: "Transform function",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Age string ` + "`json:\"age\"`" + `
}

var _ = govy.Transform(func(p Person) string { return p.Age }, func(v string) (int, error) { return 0, nil })
`,
			line:     8,
			expected: "age",
		},
		{
			name: "ForSlice function",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Tags []string ` + "`json:\"tags\"`" + `
}

var _ = govy.ForSlice(func(p Person) []string { return p.Tags })
`,
			line:     8,
			expected: "tags",
		},
		{
			name: "ForMap function",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Meta map[string]string ` + "`json:\"meta\"`" + `
}

var _ = govy.ForMap(func(p Person) map[string]string { return p.Meta })
`,
			line:     8,
			expected: "meta",
		},
		{
			name: "aliased import",
			src: `package test
import v "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string ` + "`json:\"name\"`" + `
}

var _ = v.For(func(p Person) string { return p.Name })
`,
			line:     8,
			expected: "name",
		},
		{
			name: "yaml tag when no json tag",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string ` + "`yaml:\"yaml_name\"`" + `
}

var _ = govy.For(func(p Person) string { return p.Name })
`,
			line:     8,
			expected: "yaml_name",
		},
		{
			name: "no matching line returns empty",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string
}

var _ = govy.For(func(p Person) string { return p.Name })
`,
			line:     100,
			expected: "",
		},
		{
			name: "getter with if statement",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string ` + "`json:\"name\"`" + `
}

var _ = govy.For(func(p Person) string {
	if p.Name != "" {
		return p.Name
	}
	return ""
})
`,
			line:     8,
			expected: "name",
		},
		{
			name: "getter with variable assignment",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string ` + "`json:\"name\"`" + `
}

var _ = govy.For(func(p Person) string {
	name := p.Name
	return name
})
`,
			line:     8,
			expected: "name",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := createTestPackage(t, tc.src)

			result := InferNameFromFile(res.fset, res.pkg, res.f, tc.line)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestInferNameFromFile_nonGovySelector(t *testing.T) {
	src := `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string
}

type other struct{}

func (o other) For(f func(p Person) string) {}

var o other
var _ = o.For(func(p Person) string { return p.Name })
`
	res := createTestPackage(t, src)

	result := InferNameFromFile(res.fset, res.pkg, res.f, 13)
	assert.Equal(t, "", result)
}

func TestInferNameFromFile_multipleReturns(t *testing.T) {
	src := `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string ` + "`json:\"name\"`" + `
	Age  int    ` + "`json:\"age\"`" + `
}

var _ = govy.For(func(p Person) string {
	if p.Age > 18 {
		return p.Name
	}
	return ""
})
`
	res := createTestPackage(t, src)

	result := InferNameFromFile(res.fset, res.pkg, res.f, 9)
	assert.Equal(t, "name", result)
}

func TestNameFinder_findNameInBlockStmt_nil(t *testing.T) {
	nf := nameFinder{}
	result := nf.findNameInBlockStmt(nil, nil)
	assert.Equal(t, "", result)
}

func TestNameFinder_findNameInIfStmt_nil(t *testing.T) {
	nf := nameFinder{}
	result := nf.findNameInIfStmt(nil, nil)
	assert.Equal(t, "", result)
}

func TestNameFinder_findNameInFuncLit_nil(t *testing.T) {
	nf := nameFinder{}
	result := nf.findNameInFuncLit(nil)
	assert.Equal(t, "", result)
}

func TestNameFinder_findNameInReturnStmt_nil(t *testing.T) {
	nf := nameFinder{}
	result := nf.findNameInReturnStmt(nil, nil)
	assert.Equal(t, "", result)
}

func TestNameFinder_findNameInReturnStmt_multipleResults(t *testing.T) {
	nf := nameFinder{}
	returnStmt := &ast.ReturnStmt{
		Results: []ast.Expr{
			&ast.Ident{Name: "a"},
			&ast.Ident{Name: "b"},
		},
	}
	result := nf.findNameInReturnStmt(returnStmt, nil)
	assert.Equal(t, "", result)
}

func TestNameFinder_findNameInIdent_nilObj(t *testing.T) {
	nf := nameFinder{}
	ident := &ast.Ident{Name: "test", Obj: nil}
	result := nf.findNameInIdent(ident, nil)
	assert.Equal(t, "", result)
}

func TestNameFinder_findNameInAssignStmt_multipleRhs(t *testing.T) {
	nf := nameFinder{}
	assignStmt := &ast.AssignStmt{
		Rhs: []ast.Expr{
			&ast.Ident{Name: "a"},
			&ast.Ident{Name: "b"},
		},
	}
	result := nf.findNameInAssignStmt(assignStmt, nil)
	assert.Equal(t, "", result)
}

func TestNameFinder_FindName_unexpectedType(t *testing.T) {
	nf := nameFinder{}
	result := nf.FindName("unexpected string type", nil)
	assert.Equal(t, "", result)
}

func TestNameFinder_findNameInFuncLit_multipleParams(t *testing.T) {
	funcLit := &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{{Name: "a"}}},
					{Names: []*ast.Ident{{Name: "b"}}},
				},
			},
		},
		Body: &ast.BlockStmt{},
	}
	nf := nameFinder{}
	result := nf.findNameInFuncLit(funcLit)
	assert.Equal(t, "", result)
}

func TestNameFinder_findNameInFuncLit_nonIdentParam(t *testing.T) {
	funcLit := &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{Type: &ast.StarExpr{X: &ast.Ident{Name: "Person"}}},
				},
			},
		},
		Body: &ast.BlockStmt{},
	}
	nf := nameFinder{}
	result := nf.findNameInFuncLit(funcLit)
	assert.Equal(t, "", result)
}

func TestModuleAST_FindFile_notFound(t *testing.T) {
	modAST := ModuleAST{
		FileSet:  token.NewFileSet(),
		Packages: make(map[string]*packages.Package),
	}

	pkg, file := modAST.FindFile("nonexistent.go")
	assert.Equal(t, (*packages.Package)(nil), pkg)
	assert.Equal(t, (*ast.File)(nil), file)
}

func TestModuleAST_FindFile_emptyPackages(t *testing.T) {
	modAST := ModuleAST{
		FileSet:  token.NewFileSet(),
		Packages: nil,
	}

	pkg, file := modAST.FindFile("test.go")
	assert.Equal(t, (*packages.Package)(nil), pkg)
	assert.Equal(t, (*ast.File)(nil), file)
}

type testPackageResult struct {
	pkg  *packages.Package
	fset *token.FileSet
	f    *ast.File
}

func createTestPackage(t *testing.T, src string) testPackageResult {
	t.Helper()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "test.go", src, parser.AllErrors)
	if err != nil {
		t.Fatalf("failed to parse source: %v", err)
	}

	conf := types.Config{
		Importer: nil,
		Error:    func(err error) {},
	}

	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}

	_, _ = conf.Check("test", fset, []*ast.File{f}, info)

	pkg := &packages.Package{
		Fset:      fset,
		Syntax:    []*ast.File{f},
		TypesInfo: info,
	}

	return testPackageResult{pkg: pkg, fset: fset, f: f}
}

func TestInferNameFromFile_deeplyNestedStruct(t *testing.T) {
	src := `package test
import "github.com/nobl9/govy/pkg/govy"

type Street struct {
	Name string ` + "`json:\"streetName\"`" + `
}

type Address struct {
	Street Street ` + "`json:\"street\"`" + `
}

type Person struct {
	Address Address ` + "`json:\"address\"`" + `
}

var _ = govy.For(func(p Person) string { return p.Address.Street.Name })
`
	res := createTestPackage(t, src)

	result := InferNameFromFile(res.fset, res.pkg, res.f, 16)
	assert.Equal(t, "address.street.streetName", result)
}

func TestInferNameFromFile_embeddedStruct(t *testing.T) {
	src := `package test
import "github.com/nobl9/govy/pkg/govy"

type Base struct {
	ID string ` + "`json:\"id\"`" + `
}

type Person struct {
	Base
	Name string ` + "`json:\"name\"`" + `
}

var _ = govy.For(func(p Person) string { return p.Name })
`
	res := createTestPackage(t, src)

	result := InferNameFromFile(res.fset, res.pkg, res.f, 13)
	assert.Equal(t, "name", result)
}

func TestNameFinder_findNameInSelectorExpr_unexpectedXType(t *testing.T) {
	src := `package test
import "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string
}

func getName() Person {
	return Person{}
}

var _ = govy.For(func(p Person) string { return getName().Name })
`
	res := createTestPackage(t, src)

	result := InferNameFromFile(res.fset, res.pkg, res.f, 12)
	assert.Equal(t, "", result)
}

func TestInferNameFromFile_dotImport(t *testing.T) {
	src := `package test
import . "github.com/nobl9/govy/pkg/govy"

type Person struct {
	Name string ` + "`json:\"name\"`" + `
}

var _ = For(func(p Person) string { return p.Name })
`
	res := createTestPackage(t, src)

	result := InferNameFromFile(res.fset, res.pkg, res.f, 8)
	assert.Equal(t, "name", result)
}

func TestInferNameFromFile_indexExpression(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		line     int
		expected string
	}{
		{
			name: "simple slice index with literal",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Student struct {
	Name string ` + "`json:\"name\"`" + `
}

type Teacher struct {
	Students []Student ` + "`json:\"students\"`" + `
}

var _ = govy.For(func(t Teacher) string { return t.Students[0].Name })
`,
			line:     12,
			expected: "students[0].name",
		},
		{
			name: "map with string key literal",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Value struct {
	Data string ` + "`json:\"data\"`" + `
}

type Container struct {
	Metadata map[string]Value ` + "`json:\"metadata\"`" + `
}

var _ = govy.For(func(c Container) string { return c.Metadata["key"].Data })
`,
			line:     12,
			expected: `metadata["key"].data`,
		},
		{
			name: "slice index with variable (empty brackets)",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Item struct {
	Name string ` + "`json:\"name\"`" + `
}

type List struct {
	Items []Item ` + "`json:\"items\"`" + `
}

var i = 0
var _ = govy.For(func(l List) string { return l.Items[i].Name })
`,
			line:     13,
			expected: "items[].name",
		},
		{
			name: "nested indices",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Cell struct {
	Value string ` + "`json:\"value\"`" + `
}

type Container struct {
	Matrix [][]Cell ` + "`json:\"matrix\"`" + `
}

var _ = govy.For(func(c Container) string { return c.Matrix[0][1].Value })
`,
			line:     12,
			expected: "matrix[0][1].value",
		},
		{
			name: "deep chain with multiple indices",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Employee struct {
	Name string ` + "`json:\"name\"`" + `
}

type Department struct {
	Employees []Employee ` + "`json:\"employees\"`" + `
}

type Company struct {
	Departments []Department ` + "`json:\"departments\"`" + `
}

var _ = govy.For(func(c Company) string { return c.Departments[0].Employees[0].Name })
`,
			line:     16,
			expected: "departments[0].employees[0].name",
		},
		{
			name: "slice of pointers",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Student struct {
	Name string ` + "`json:\"name\"`" + `
}

type Teacher struct {
	Students []*Student ` + "`json:\"students\"`" + `
}

var _ = govy.For(func(t Teacher) string { return t.Students[0].Name })
`,
			line:     12,
			expected: "students[0].name",
		},
		{
			name: "array instead of slice",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Point struct {
	X int ` + "`json:\"x\"`" + `
}

type Shape struct {
	Vertices [3]Point ` + "`json:\"vertices\"`" + `
}

var _ = govy.For(func(s Shape) int { return s.Vertices[0].X })
`,
			line:     12,
			expected: "vertices[0].x",
		},
		{
			name: "map with integer key",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Value struct {
	Data string ` + "`json:\"data\"`" + `
}

type Container struct {
	Items map[int]Value ` + "`json:\"items\"`" + `
}

var _ = govy.For(func(c Container) string { return c.Items[42].Data })
`,
			line:     12,
			expected: "items[42].data",
		},
		{
			name: "index with binary expression falls back to empty brackets",
			src: `package test
import "github.com/nobl9/govy/pkg/govy"

type Item struct {
	Name string ` + "`json:\"name\"`" + `
}

type List struct {
	Items []Item ` + "`json:\"items\"`" + `
}

var offset = 1
var _ = govy.For(func(l List) string { return l.Items[offset+1].Name })
`,
			line:     13,
			expected: "items[].name",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := createTestPackage(t, tc.src)
			result := InferNameFromFile(res.fset, res.pkg, res.f, tc.line)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestInferNameFromFile_indexExpressionWithFunctionCall(t *testing.T) {
	src := `package test
import "github.com/nobl9/govy/pkg/govy"

type Student struct {
	Name string
}

func getStudents() []Student {
	return nil
}

var _ = govy.For(func(_ struct{}) string { return getStudents()[0].Name })
`
	res := createTestPackage(t, src)

	result := InferNameFromFile(res.fset, res.pkg, res.f, 12)
	assert.Equal(t, "", result)
}

func TestNameFinder_getStructFromType_unhandledType(t *testing.T) {
	nf := nameFinder{}
	// Test with a basic type that is not a struct, slice, array, or map.
	basicType := types.Typ[types.Int]
	result, ok := nf.getStructFromType(basicType)
	assert.Equal(t, false, ok)
	assert.Equal(t, (*types.Struct)(nil), result)
}
