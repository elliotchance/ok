package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func compileMap(compiledFunc *CompiledFunc, n *ast.Map) (string, error) {
	// TODO(elliot): Check type is valid for the map.
	// TODO(elliot): Maps with duplicate keys should be an error.

	sizeRegister := compiledFunc.nextRegister()
	compiledFunc.append(&instruction.Assign{
		VariableName: sizeRegister,
		Value:        ast.NewLiteralNumber(fmt.Sprintf("%d", len(n.Elements))),
	})

	mapRegister := compiledFunc.nextRegister()
	compiledFunc.append(&instruction.MapAllocNumber{
		Size:   sizeRegister,
		Result: mapRegister,
	})

	for _, element := range n.Elements {
		// TODO(elliot): Check keyKind is string.
		keyRegister, _, err := compileExpr(compiledFunc, element.Key)
		if err != nil {
			return "", err
		}

		// TODO(elliot): Check value is the right type for map.
		valueRegister, _, _ := compileExpr(compiledFunc, element.Value)

		compiledFunc.append(&instruction.MapSet{
			Map:   mapRegister,
			Key:   keyRegister,
			Value: valueRegister,
		})
	}

	return mapRegister, nil
}
