package compiler

import (
	"ok/ast"
	"ok/instruction"
)

func compileStatement(compiledFunc *CompiledFunc, statement ast.Node, breakIns, continueIns instruction.Instruction) error {
	switch n := statement.(type) {
	case *ast.Break:
		compiledFunc.append(breakIns)
		return nil

	case *ast.Continue:
		compiledFunc.append(continueIns)
		return nil

	case *ast.For:
		return compileFor(compiledFunc, n)

	case *ast.If:
		return compileIf(compiledFunc, n, breakIns, continueIns)

	case *ast.Switch:
		return compileSwitch(compiledFunc, n, breakIns, continueIns)
	}

	_, _, err := compileExpr(compiledFunc, statement)

	return err
}
