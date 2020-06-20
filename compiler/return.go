package compiler

import (
	"ok/ast"
	"ok/vm"
)

func compileReturn(compiledFunc *vm.CompiledFunc, n *ast.Return, fns map[string]*ast.Func) error {
	var results []string
	for _, expr := range n.Exprs {
		// TODO(elliot): Check return types are valid.
		result, _, err := compileExpr(compiledFunc, expr, fns)
		if err != nil {
			return err
		}

		results = append(results, result...)
	}

	compiledFunc.Append(&vm.Return{
		Results: results,
	})

	return nil
}
