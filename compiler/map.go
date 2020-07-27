package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"
)

func compileMap(
	compiledFunc *vm.CompiledFunc,
	n *ast.Map,
	fns map[string]*ast.Func,
) (vm.Register, error) {
	// TODO(elliot): Check type is valid for the map.
	// TODO(elliot): Maps with duplicate keys should be an error.

	sizeRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.Assign{
		VariableName: sizeRegister,
		Value:        asttest.NewLiteralNumber(fmt.Sprintf("%d", len(n.Elements))),
	})

	mapRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.MapAllocNumber{
		Size:   sizeRegister,
		Result: mapRegister,
	})

	for _, element := range n.Elements {
		// TODO(elliot): Check keyKind is string.
		keyRegisters, _, err := compileExpr(compiledFunc, element.Key, fns)
		if err != nil {
			return "", err
		}

		// TODO(elliot): Check value is the right type for map.
		valueRegisters, _, _ := compileExpr(compiledFunc, element.Value, fns)

		compiledFunc.Append(&vm.MapSet{
			Map:   mapRegister,
			Key:   keyRegisters[0],
			Value: valueRegisters[0],
		})
	}

	return mapRegister, nil
}
