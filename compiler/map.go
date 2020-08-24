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
	file *Compiled,
) (vm.Register, string, error) {
	// TODO(elliot): Check type is valid for the map.
	// TODO(elliot): Maps with duplicate keys should be an error.

	sizeRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.Assign{
		VariableName: sizeRegister,
		Value:        asttest.NewLiteralNumber(fmt.Sprintf("%d", len(n.Elements))),
	})

	mapRegister := compiledFunc.NextRegister()
	mapAlloc := &vm.MapAlloc{
		Kind:   n.Kind,
		Size:   sizeRegister,
		Result: mapRegister,
	}
	compiledFunc.Append(mapAlloc)

	for _, element := range n.Elements {
		// TODO(elliot): Check keyKind is string.
		keyRegisters, _, err := compileExpr(compiledFunc, element.Key, file)
		if err != nil {
			return "", "", err
		}

		// TODO(elliot): Check value is the right type for map.
		valueRegisters, valueKind, _ := compileExpr(compiledFunc, element.Value, file)
		if mapAlloc.Kind == "" {
			mapAlloc.Kind = "{}" + valueKind[0]
		}

		compiledFunc.Append(&vm.MapSet{
			Map:   mapRegister,
			Key:   keyRegisters[0],
			Value: valueRegisters[0],
		})
	}

	return mapRegister, mapAlloc.Kind, nil
}
