package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileAssert(
	compiledFunc *vm.CompiledFunc,
	n *ast.Assert,
	fns map[string]*ast.Func,
) error {
	left, right, returns, returnKind, err := compileComparison(compiledFunc, n.Expr, fns)
	if err != nil {
		return err
	}

	if returnKind != "bool" {
		return fmt.Errorf("assert condition must be a bool but is %s", returnKind)
	}

	compiledFunc.Append(&vm.Assert{
		Left:  left,
		Op:    n.Expr.Op,
		Right: right,
		Final: returns,
		Pos:   n.Position(),
	})

	return nil
}
