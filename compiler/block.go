package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

// compileBlock translates list of statements into a set of instructions. The
// number of instructions returned may be zero.
func compileBlock(compiledFunc *vm.CompiledFunc, stmts []ast.Node, breakIns, continueIns vm.Instruction, file *Compiled) error {
	for _, statement := range stmts {
		err := compileStatement(compiledFunc, statement, breakIns, continueIns, file)
		if err != nil {
			return err
		}
	}

	return nil
}
