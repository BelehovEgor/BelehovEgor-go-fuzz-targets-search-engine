package metrics

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func Test_CalculateCyclomaticComplexity_Func_ShouldCalculate(t *testing.T) {
	// arrage
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{"Simple Function", "package pkg\nfunc simple() {}", 1},
		{"If Statement", "package pkg\nfunc withIf() { if true {} }", 2},
		{"For Loop", "package pkg\nfunc withFor() { for i := 0; i < 10; i++ {} }", 2},
		{"Range Loop", "package pkg\nfunc withRange() { arr := []int{1, 2, 3}; for _, v := range arr {} }", 2},
		{"Switch Statement", "package pkg\nfunc withSwitch() { switch x { case 1: default: } }", 3},
		{"Type Switch Statement", "package pkg\nfunc withTypeSwitch() { switch x := y.(type) { case int: default: } }", 3},
		{"Binary Expression", "package pkg\nfunc withBinaryExpr() { if a && b {} }", 3},
		{"Multiple Constructs", "package pkg\nfunc complex() { if true {}; for i := 0; i < 10; i++ {}; switch x { case 1: default: }; if a && b {} }", 7},
	}

	// act
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", tc.input, 0)
			if err != nil {
				t.Fatalf("Failed to parse input: %v", err)
			}

			funcDecl := f.Decls[0].(*ast.FuncDecl)
			actual := calculateCyclomaticComplexity(funcDecl)

			// assert
			if actual != tc.expected {
				t.Errorf("Test case: %s. Expected cyclomatic complexity of %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func Test_CountCycles_Func_ShouldCount(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{"No Loops", "package pkg\nfunc noLoops() {}", 0},
		{"Single For Loop", "package pkg\nfunc singleFor() { for i := 0; i < 10; i++ {} }", 1},
		{"Single Range Loop", "package pkg\nfunc singleRange() { arr := []int{1, 2, 3}; for _, v := range arr {} }", 1},
		{"Multiple Loops", "package pkg\nfunc multipleLoops() { for i := 0; i < 10; i++ {}; arr := []int{1, 2, 3}; for _, v := range arr {} }", 2},
		{"Nested Loops", "package pkg\nfunc nestedLoops() { for i := 0; i < 10; i++ { arr := []int{1, 2, 3}; for _, v := range arr {} } }", 2},
		{"Complex Function", "package pkg\nfunc complex() { for i := 0; i < 10; i++ {}; arr := []int{1, 2, 3}; for _ = range arr {}; for j := 0; j < 5; j++ {} }", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", tc.input, 0)
			if err != nil {
				t.Fatalf("Test case: %s. Failed to parse input: %v", tc.name, err)
			}

			funcDecl := f.Decls[0].(*ast.FuncDecl)
			actual := countCycles(funcDecl)
			if actual != tc.expected {
				t.Errorf("Test case: %s. Expected cycle count of %d, got %d", tc.name, tc.expected, actual)
			}
		})
	}
}

func Test_CountNestedLoops_Func_ShouldCount(t *testing.T) {
	var testCases = []struct {
		name           string
		input          string
		expectedCount  int
		expectedMaxDep int
	}{
		{"No Loops", "package pkg\nfunc noLoops() {}", 0, 0},
		{"Single For Loop", "package pkg\nfunc singleFor() { for i := 0; i < 10; i++ {} }", 0, 1},
		{"Single Range Loop", "package pkg\nfunc singleRange() { arr := []int{1, 2, 3}; for _ = range arr {} }", 0, 1},
		{"Nested Loops", "package pkg\nfunc nestedLoops() { for i := 0; i < 10; i++ { arr := []int{1, 2, 3}; for _ = range arr { } } }", 1, 2},
		{"Double Nested Loops", "package pkg\nfunc doubleNestedLoops() { for i := 0; i < 10; i++ { for j := 0; j < 5; j++ { arr := []int{1, 2, 3}; for _, v := range arr {} } } }", 3, 3},
		{"Mixed Nested Loops", "package pkg\nfunc mixedNestedLoops() { for i := 0; i < 10; i++ { arr := []int{1, 2, 3}; for _ = range arr {}; for j := 0; j < 5; j++ {} } }", 2, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", tc.input, 0)
			if err != nil {
				t.Fatalf("Test case: %s. Failed to parse input: %v", tc.name, err)
			}

			funcDecl := f.Decls[0].(*ast.FuncDecl)
			actualCount, actualMaxDep := countNestedLoops(funcDecl)
			if actualCount != tc.expectedCount {
				t.Errorf("Test case: %s. Expected nested loop count of %d, got %d", tc.name, tc.expectedCount, actualCount)
			}
			if actualMaxDep != tc.expectedMaxDep {
				t.Errorf("Test case: %s. Expected maximum depth of %d, got %d", tc.name, tc.expectedMaxDep, actualMaxDep)
			}
		})
	}
}
