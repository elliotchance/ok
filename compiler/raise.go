package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileRaise(
	compiledFunc *vm.CompiledFunc,
	n *ast.Raise,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) error {
	// TODO(elliot): Check this call returns a type that satisfies an error
	//  interface.
	//
	// TODO(elliot): It may not be a Call.
	result, resultKind, err := compileCall(compiledFunc, n.Err.(*ast.Call),
		file, scopeOverrides)
	if err != nil {
		return err
	}

	compiledFunc.Append(&vm.Raise{
		Err:  result[0],
		Type: resultKind[0],
	})

	return nil
}
