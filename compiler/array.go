package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/vm"
)

func compileArray(
	compiledFunc *vm.CompiledFunc,
	n *ast.Array,
	file *vm.File,
) (vm.Register, string, error) {
	if len(n.Elements) == 0 && n.Kind == "" {
		err := fmt.Errorf("%s empty array needs to specify a type",
			n.Position())

		return "", "", err
	}

	sizeRegister := compiledFunc.NextRegister()
	compiledFunc.Append(&vm.Assign{
		VariableName: sizeRegister,
		Value:        asttest.NewLiteralNumber(fmt.Sprintf("%d", len(n.Elements))),
	})

	arrayRegister := compiledFunc.NextRegister()
	arrayAlloc := &vm.ArrayAlloc{
		Size:   sizeRegister,
		Result: arrayRegister,
		Kind:   n.Kind,
	}
	compiledFunc.Append(arrayAlloc)

	for index, value := range n.Elements {
		indexRegister := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: indexRegister,
			Value:        asttest.NewLiteralNumber(fmt.Sprintf("%d", index)),
		})

		// TODO(elliot): Check all elements are the same kind.
		valueRegisters, valueKind, _ := compileExpr(compiledFunc, value, file)
		if arrayAlloc.Kind == "" {
			arrayAlloc.Kind = "[]" + valueKind[0]
		}

		compiledFunc.Append(&vm.ArraySet{
			Array: arrayRegister,
			Index: indexRegister,
			Value: valueRegisters[0],
		})
	}

	return arrayRegister, arrayAlloc.Kind, nil
}
