package nameinfer

import (
	"go/ast"
	"go/token"
	"log"
	"sync"

	"golang.org/x/tools/go/packages"

	"github.com/nobl9/govy/internal"
	"github.com/nobl9/govy/pkg/govyconfig"
)

const packagesMode = packages.NeedName |
	packages.NeedFiles |
	packages.NeedTypes |
	packages.NeedSyntax |
	packages.NeedTypesInfo |
	packages.NeedImports

func NewModuleAST(root string) ModuleAST {
	fileSet := token.NewFileSet()
	cfg := &packages.Config{
		Fset:  fileSet,
		Mode:  packagesMode,
		Dir:   root,
		Tests: govyconfig.GetNameInferIncludeTestFiles(),
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		log.Fatal(err)
	}
	pkgMap := make(map[string]*packages.Package, len(pkgs))
	for _, pkg := range pkgs {
		pkgMap[pkg.PkgPath] = pkg
	}
	return ModuleAST{
		FileSet:  fileSet,
		Packages: pkgMap,
	}
}

type ModuleAST struct {
	FileSet  *token.FileSet
	Packages map[string]*packages.Package
}

func (a ModuleAST) FindFile(file string) (*packages.Package, *ast.File) {
	for _, pkg := range a.Packages {
		for i, filePath := range pkg.GoFiles {
			if filePath == file {
				return pkg, pkg.Syntax[i]
			}
		}
	}
	return nil, nil
}

var (
	modAST       ModuleAST
	parseASTOnce sync.Once
)

func parseModuleASTOnce() {
	parseASTOnce.Do(func() { modAST = NewModuleAST(internal.FindModuleRoot()) })
}
