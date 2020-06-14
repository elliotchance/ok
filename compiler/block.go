package compiler

import (
	"ok/ast"
	"ok/instruction"
)

// compileBlock translates list of statements into a set of instructions. The
// number of instructions returned may be zero.
func compileBlock(compiledFunc *CompiledFunc, stmts []ast.Node, breakIns, continueIns instruction.Instruction) error {
	for _, statement := range stmts {
		err := compileStatement(compiledFunc, statement, breakIns, continueIns)
		if err != nil {
			return err
		}
	}

	return nil
}
