package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileLiteral(
	compiledFunc *vm.CompiledFunc,
	e *ast.Literal,
	file *vm.File,
) (vm.Register, *types.Type) {
	returns := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.AssignSymbol{
		Result: returns,
		Symbol: file.AddSymbolLiteral(e),
	})

	return returns, e.Kind
}
