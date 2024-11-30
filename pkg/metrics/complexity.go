package metrics

import (
	"fmt"
	ast "go/ast"
	"go/parser"
	"go/token"

	myAst "github.com/BelehovEgor/go-fuzz-targets-search-engine/pkg/ast"
)

type Code string

type Dimension struct {
	// Func info
	Name string

	// Functions
	cyclomatic_complexity int

	// Loops
	number_of_loops                int
	number_of_nested_loops         int
	maximum_nesting_level_of_loops int
}

func (code Code) CalculateDimensions() ([]*Dimension, error) {
	f, err := parseFile(code)
	if err != nil {
		return nil, err
	}

	var dimensions = make([]*Dimension, 0)
	for _, target := range findFuncDecls(f) {
		dimension, err := calculateDimension(target)
		if err != nil {
			return nil, err
		}

		dimensions = append(dimensions, dimension)
	}

	return dimensions, nil
}

func (code Code) CalculateDimension(funcName string) (*Dimension, error) {
	f, err := parseFile(code)
	if err != nil {
		return nil, err
	}

	targetFunc, err := findFuncDeclByName(f, funcName)
	if err != nil {
		return nil, err
	}

	return calculateDimension(targetFunc)
}

func calculateDimension(targetFunc *ast.FuncDecl) (*Dimension, error) {
	var dimension *Dimension = &Dimension{}

	dimension.Name = targetFunc.Name.Name

	dimension.cyclomatic_complexity = calculateCyclomaticComplexity(targetFunc)
	dimension.number_of_loops = countCycles(targetFunc)

	countNestedLoops, maxDepth := countNestedLoops(targetFunc)
	dimension.number_of_nested_loops = countNestedLoops
	dimension.maximum_nesting_level_of_loops = maxDepth

	return dimension, nil
}

func parseFile(code Code) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", string(code), 0)
}

func findFuncDecls(file *ast.File) []*ast.FuncDecl {
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

func findFuncDeclByName(file *ast.File, funcName string) (*ast.FuncDecl, error) {
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
