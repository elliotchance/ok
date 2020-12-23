package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileMap(
	compiledFunc *vm.CompiledFunc,
	n *ast.Map,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) (vm.Register, *types.Type, error) {
	// TODO(elliot): Check type is valid for the map.
	// TODO(elliot): Maps with duplicate keys should be an error.

	sizeRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.Assign{
		VariableName: sizeRegister,
		Value:        asttest.NewLiteralNumber(fmt.Sprintf("%d", len(n.Elements))),
	})

	mapRegister := compiledFunc.NextRegister()
	mapAlloc := &vm.MapAlloc{
		Size:   sizeRegister,
		Result: mapRegister,
	}
	compiledFunc.Append(mapAlloc)

	for _, element := range n.Elements {
		// TODO(elliot): Check keyKind is string.
		keyRegisters, _, err := compileExpr(compiledFunc, element.Key, file,
			scopeOverrides)
		if err != nil {
			return "", nil, err
		}

		// TODO(elliot): Check value is the right type for map.
		// TODO(elliot): Check all elements are the same kind.
		valueRegisters, valueKind, _ := compileExpr(compiledFunc, element.Value,
			file, scopeOverrides)
		if n.Kind == nil {
			n.Kind = valueKind[0].ToMap()
		}

		compiledFunc.Append(&vm.MapSet{
			Map:   mapRegister,
			Key:   keyRegisters[0],
			Value: valueRegisters[0],
		})
	}

	typeRegister := file.AddType(n.Kind)
	mapAlloc.Kind = typeRegister

	return mapRegister, n.Kind, nil
}
