package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileStatement(compiledFunc *vm.CompiledFunc, statement ast.Node, breakIns, continueIns vm.Instruction, fns map[string]*ast.Func) error {
	switch n := statement.(type) {
	case *ast.Break:
		compiledFunc.Append(breakIns)
		return nil

	case *ast.Continue:
		compiledFunc.Append(continueIns)
		return nil

	case *ast.Return:
		return compileReturn(compiledFunc, n, fns)

	case *ast.Assert:
		return compileAssert(compiledFunc, n, fns)

	case *ast.For:
		return compileFor(compiledFunc, n, fns)

	case *ast.If:
		return compileIf(compiledFunc, n, breakIns, continueIns, fns)

	case *ast.Switch:
		return compileSwitch(compiledFunc, n, breakIns, continueIns, fns)

	case *ast.ErrorScope:
		return compileErrorScope(compiledFunc, n, fns)

	case *ast.Raise:
		return compileRaise(compiledFunc, n, fns)
	}

	_, _, err := compileExpr(compiledFunc, statement, fns)

	return err
}
