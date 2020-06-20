package compiler

import (
	"ok/ast"
	"ok/vm"
)

// compileBlock translates list of statements into a set of instructions. The
// number of instructions returned may be zero.
func compileBlock(compiledFunc *vm.CompiledFunc, stmts []ast.Node, breakIns, continueIns vm.Instruction, fns map[string]*ast.Func) error {
	for _, statement := range stmts {
		err := compileStatement(compiledFunc, statement, breakIns, continueIns, fns)
		if err != nil {
			return err
		}
	}

	return nil
}
