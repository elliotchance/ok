package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileLiteral(
	compiledFunc *vm.CompiledFunc,
	e *ast.Literal,
) (vm.Register, *types.Type) {
	returns := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.Assign{
		VariableName: returns,
		Value:        e,
	})

	return returns, e.Kind
}
