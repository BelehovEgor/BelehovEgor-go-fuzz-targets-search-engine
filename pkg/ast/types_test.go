package ast

import (
	"fmt"
	"go/ast"
	"testing"
)

func Test_getVariableType(t *testing.T) {
	pkg, err := GetPackage("D:/itmo/3-sem-2024-2025/go-fuzz-targets-search-engine/pkg/ast", "ast")
	if err != nil {
		t.Errorf("should be nil")
	}

	ast.Inspect(pkg.Syntax[0], func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			if t, ok := pkg.TypesInfo.Types[x.Lhs[0]]; ok {
				fmt.Print(t)
			}
		}
		return true
	})

}
