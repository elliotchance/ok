package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileReturn(
	compiledFunc *vm.CompiledFunc,
	n *ast.Return,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) error {
	var results []vm.Register
	for _, expr := range n.Exprs {
		// TODO(elliot): Check return types are valid.
		result, _, err := compileExpr(compiledFunc, expr, file, scopeOverrides)
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
