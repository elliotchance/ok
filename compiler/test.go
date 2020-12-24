package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

// CompileTest will compile a test.
func CompileTest(
	fn *ast.Test,
	file *vm.File,
	constants map[string]*ast.Literal,
) (*vm.CompiledTest, error) {
	// Tests can be compiled as if they were functions, then wrapped in a
	// CompiledTest.
	//
	// TODO(elliot): When tests can be nested the third argument for parentFunc
	//  should not be nil.
	compiledFunc, err := CompileFunc(&ast.Func{
		Statements: fn.Statements,
		Pos:        fn.Pos,
	}, file, nil, constants)
	if err != nil {
		return nil, err
	}

	return &vm.CompiledTest{
		CompiledFunc: compiledFunc,
		TestName:     fn.Name,
	}, nil
}
