package compiler

import (
	"errors"
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func compileArray(compiledFunc *CompiledFunc, n *ast.Array) (string, error) {
	if len(n.Elements) == 0 && n.Kind == "" {
		return "", errors.New("empty array needs to specify a type")
	}

	sizeRegister := compiledFunc.nextRegister()
	compiledFunc.append(&instruction.Assign{
		VariableName: sizeRegister,
		Value:        ast.NewLiteralNumber(fmt.Sprintf("%d", len(n.Elements))),
	})

	arrayRegister := compiledFunc.nextRegister()
	compiledFunc.append(&instruction.ArrayAllocNumber{
		Size:   sizeRegister,
		Result: arrayRegister,
	})

	for index, value := range n.Elements {
		indexRegister := compiledFunc.nextRegister()
		compiledFunc.append(&instruction.Assign{
			VariableName: indexRegister,
			Value:        ast.NewLiteralNumber(fmt.Sprintf("%d", index)),
		})

		valueRegister, _, _ := compileExpr(compiledFunc, value)

		compiledFunc.append(&instruction.ArraySet{
			Array: arrayRegister,
			Index: indexRegister,
			Value: valueRegister,
		})
	}

	return arrayRegister, nil
}
