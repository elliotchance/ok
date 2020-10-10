package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileStatement(compiledFunc *vm.CompiledFunc, statement ast.Node, breakIns, continueIns vm.Instruction, file *vm.File) error {
	switch n := statement.(type) {
	case *ast.Break:
		compiledFunc.Append(breakIns)
		return nil

	case *ast.Continue:
		compiledFunc.Append(continueIns)
		return nil

	case *ast.Return:
		return compileReturn(compiledFunc, n, file)

	case *ast.Assert:
		return compileAssert(compiledFunc, n, file)

	case *ast.AssertRaise:
		return compileAssertRaise(compiledFunc, n, file)

	case *ast.For:
		return compileFor(compiledFunc, n, file)

	case *ast.If:
		return compileIf(compiledFunc, n, breakIns, continueIns, file)

	case *ast.Switch:
		return compileSwitch(compiledFunc, n, breakIns, continueIns, file)

	case *ast.ErrorScope:
		return compileErrorScope(compiledFunc, n, file)

	case *ast.Raise:
		return compileRaise(compiledFunc, n, file)
	}

	_, _, err := compileExpr(compiledFunc, statement, file)

	return err
}
