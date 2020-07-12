package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileRaise(compiledFunc *vm.CompiledFunc, n *ast.Raise, fns map[string]*ast.Func) error {
	// TODO(elliot): Check this call returns a type that satisfies an error
	//  interface.
	result, resultKind, err := compileCall(compiledFunc, n.Err, fns)
	if err != nil {
		return err
	}

	compiledFunc.Append(&vm.Raise{
		Err:  result[0],
		Type: resultKind[0],
	})

	return nil
}