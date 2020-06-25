package compiler

import (
	"ok/ast"
	"ok/vm"
)

// CompileTest will compile a test.
func CompileTest(fn *ast.Test, fns map[string]*ast.Func) (*vm.CompiledTest, error) {
	// Tests can be compiled as if they were functions, then wrapped in a
	// CompiledTest.
	compiledFunc, err := CompileFunc(&ast.Func{
		Statements: fn.Statements,
	}, fns)
	if err != nil {
		return nil, err
	}

	return &vm.CompiledTest{
		CompiledFunc: compiledFunc,
		TestName:     fn.Name,
	}, nil
}
