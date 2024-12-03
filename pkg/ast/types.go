package ast

import (
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
)

func GetPackage(folder, packageName string) (*packages.Package, error) {
	cfg := &packages.Config{
		Dir:  folder,
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("load packages error: %v", err)
	}

	var searchedPackage *packages.Package
	for _, pkg := range pkgs {
		if pkg.Name == packageName {
			searchedPackage = pkg
			break
		}
	}

	if searchedPackage == nil {
		return nil, fmt.Errorf("package not found")
	}

	return searchedPackage, nil
}

func GetFuncs(pkg *packages.Package) []*ast.FuncDecl {
	fdecls := make([]*ast.FuncDecl, 0)
	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.FuncDecl:
				fdecls = append(fdecls, x)
			}
			return true
		})
	}

	return fdecls
}

func GetFunc(pkg *packages.Package, funcName string) (*ast.FuncDecl, error) {
	var fdecl *ast.FuncDecl
	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			switch x := node.(type) {
			case *ast.FuncDecl:
				if x.Name.Name == funcName {
					fdecl = x
					return false
				}
			}
			return fdecl == nil
		})
	}

	if fdecl == nil {
		return nil, fmt.Errorf("func not found")
	}

	return fdecl, nil
}

func GetType(pkg *packages.Package, target ast.Expr) (*types.TypeAndValue, error) {
	targetType, ok := pkg.TypesInfo.Types[target]
	if !ok {
		return nil, fmt.Errorf("type undefined")
	}

	return &targetType, nil
}
