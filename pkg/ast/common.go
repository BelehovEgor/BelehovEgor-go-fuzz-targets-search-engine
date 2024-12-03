package ast

import (
	"fmt"
	ast "go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(code string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", string(code), 0)
}

func FindFuncDecls(file *ast.File) []*ast.FuncDecl {
	foundFuncs := make([]*ast.FuncDecl, 0)

	ast.Inspect(file, func(node ast.Node) bool {
		decl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		foundFuncs = append(foundFuncs, decl)

		return true
	})

	return foundFuncs
}

func FindFilesFuncDecls(files []*ast.File) []*ast.FuncDecl {
	foundFuncs := make([]*ast.FuncDecl, 0)

	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) bool {
			decl, ok := node.(*ast.FuncDecl)
			if !ok {
				return true
			}

			foundFuncs = append(foundFuncs, decl)

			return true
		})
	}

	return foundFuncs
}

func FindFuncDeclByName(file *ast.File, funcName string) (*ast.FuncDecl, error) {
	var foundFunc *ast.FuncDecl

	ast.Inspect(file, func(node ast.Node) bool {
		decl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if decl.Name.Name == funcName {
			foundFunc = decl
			return false
		}

		return true
	})

	if foundFunc == nil {
		return nil, fmt.Errorf("function with name '%s' not found", funcName)
	}

	return foundFunc, nil
}
