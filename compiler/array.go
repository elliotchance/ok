package compiler

import (
	"errors"
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"
)

func compileArray(compiledFunc *vm.CompiledFunc, n *ast.Array, fns map[string]*ast.Func) (string, error) {
	if len(n.Elements) == 0 && n.Kind == "" {
		return "", errors.New("empty array needs to specify a type")
	}

	sizeRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.Assign{
		VariableName: sizeRegister,
		Value:        asttest.NewLiteralNumber(fmt.Sprintf("%d", len(n.Elements))),
	})

	arrayRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.ArrayAllocNumber{
		Size:   sizeRegister,
		Result: arrayRegister,
	})

	for index, value := range n.Elements {
		indexRegister := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: indexRegister,
			Value:        asttest.NewLiteralNumber(fmt.Sprintf("%d", index)),
		})

		valueRegisters, _, _ := compileExpr(compiledFunc, value, fns)

		compiledFunc.Append(&vm.ArraySet{
			Array: arrayRegister,
			Index: indexRegister,
			Value: valueRegisters[0],
		})
	}

	return arrayRegister, nil
}
