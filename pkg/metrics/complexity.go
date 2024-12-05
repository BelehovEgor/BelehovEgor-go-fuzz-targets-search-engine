package metrics

import (
	ast "go/ast"
	"go/token"

	myAst "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/ast"
	"golang.org/x/tools/go/packages"
)

type Complexity struct {
	// Func info
	Package string
	Name    string

	// Functions
	cyclomatic int

	// Loops
	number_of_loops                int
	number_of_nested_loops         int
	maximum_nesting_level_of_loops int
}

func (comp Complexity) GetPackage() string {
	return comp.Package
}

func (comp Complexity) GetName() string {
	return comp.Name
}

func (comp Complexity) GetScore() int {
	return comp.cyclomatic + comp.number_of_loops + comp.number_of_nested_loops + comp.maximum_nesting_level_of_loops
}

func CalculateComplexitiesFromPackage(pkg *packages.Package) ([]Rankable, error) {
	var complexity = make([]Rankable, 0)
	for _, target := range myAst.GetFuncs(pkg) {
		dimension, err := calculateComplexity(target, pkg)
		if err != nil {
			return nil, err
		}

		complexity = append(complexity, dimension)
	}

	return complexity, nil
}

func CalculateComplexityFromPackage(pkg *packages.Package, funcName string) (*Complexity, error) {
	targetFunc, err := myAst.GetFunc(pkg, funcName)
	if err != nil {
		return nil, err
	}

	return calculateComplexity(targetFunc, pkg)
}

func calculateComplexity(targetFunc *ast.FuncDecl, pkg *packages.Package) (*Complexity, error) {
	var complexity *Complexity = &Complexity{
		Package: pkg.PkgPath,
		Name:    targetFunc.Name.Name,
	}

	complexity.cyclomatic = calculateCyclomaticComplexity(targetFunc)
	complexity.number_of_loops = countCycles(targetFunc)

	countNestedLoops, maxDepth := countNestedLoops(targetFunc)
	complexity.number_of_nested_loops = countNestedLoops
	complexity.maximum_nesting_level_of_loops = maxDepth

	return complexity, nil
}

func calculateCyclomaticComplexity(f *ast.FuncDecl) int {
	var complexity int = 1

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.IfStmt:
			complexity++
		case *ast.ForStmt:
			complexity++
		case *ast.RangeStmt:
			complexity++
		case *ast.SwitchStmt:
			complexity += len(x.Body.List)
		case *ast.TypeSwitchStmt:
			complexity += len(x.Body.List)
		case *ast.BinaryExpr:
			if x.Op == token.LAND || x.Op == token.LOR {
				complexity++
			}
		}
		return true
	})

	return complexity
}

func countCycles(f *ast.FuncDecl) int {
	var cycleCount int

	ast.Inspect(f, func(node ast.Node) bool {
		switch node.(type) {
		case *ast.ForStmt, *ast.RangeStmt:
			cycleCount++
		}
		return true
	})

	return cycleCount
}

func countNestedLoops(f *ast.FuncDecl) (int, int) {
	nestedLoopCount := 0
	maxDepth := 0
	currentDepth := 0

	myAst.Inspect(f,
		func(node ast.Node) {
			switch node.(type) {
			case *ast.RangeStmt, *ast.ForStmt:
				nestedLoopCount += currentDepth

				currentDepth++
			}
		},
		func(node ast.Node) bool {
			return true
		},
		func(node ast.Node) {
			switch node.(type) {
			case *ast.RangeStmt, *ast.ForStmt:
				maxDepth = max(maxDepth, currentDepth)
				currentDepth--
			}
		})

	return nestedLoopCount, maxDepth
}
