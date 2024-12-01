package ast

import "testing"

func Test_FindFuncDeclByName_FuncExists_ShouldReturn(t *testing.T) {
	// arrange
	funcName := "exampleFunction"
	code := `
	package test

    func exampleFunction(a int, b int) int {
        if a > b {
            return a
        } else if a < b {
            return b
        } else {
            return a + b
        }
    }`

	file, err := ParseFile(code)
	if err != nil {
		t.Fatalf("Error should be nil")
	}

	// act
	result, err := FindFuncDeclByName(file, funcName)

	if err != nil {
		t.Fatalf("Error should be nil")
	}

	if result.Name.Name != funcName {
		t.Fatalf("Invalid function returned")
	}
}

func Test_FindFuncDeclByName_FuncDoesNotExist_ShouldReturnError(t *testing.T) {
	// arrange
	funcName := "someName"
	code := `
	package test

    func exampleFunction(a int, b int) int {
        if a > b {
            return a
        } else if a < b {
            return b
        } else {
            return a + b
        }
    }`

	file, err := ParseFile(code)
	if err != nil {
		t.Fatalf("Error should be nil")
	}

	// act
	_, err = FindFuncDeclByName(file, funcName)

	if err == nil {
		t.Errorf("Error should not be nil")
	}

	if err.Error() != "function with name 'someName' not found" {
		t.Errorf("Invalid error message")
	}
}
