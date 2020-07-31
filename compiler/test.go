package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

// CompileTest will compile a test.
func CompileTest(fn *ast.Test, file *Compiled) (*vm.CompiledTest, error) {
	// Tests can be compiled as if they were functions, then wrapped in a
	// CompiledTest.
	compiledFunc, err := CompileFunc(&ast.Func{
		Statements: fn.Statements,
		Pos:        fn.Pos,
	}, file)
	if err != nil {
		return nil, err
	}

	return &vm.CompiledTest{
		CompiledFunc: compiledFunc,
		TestName:     fn.Name,
	}, nil
}
