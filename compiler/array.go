package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileArray(
	compiledFunc *vm.CompiledFunc,
	n *ast.Array,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) (vm.Register, *types.Type, error) {
	if len(n.Elements) == 0 && n.Kind == nil {
		err := fmt.Errorf("%s empty array needs to specify a type",
			n.Position())

		return "", nil, err
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
		// We'll fix the Type later, see below.
	}
	compiledFunc.Append(arrayAlloc)

	for index, value := range n.Elements {
		indexRegister := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: indexRegister,
			Value:        asttest.NewLiteralNumber(fmt.Sprintf("%d", index)),
		})

		// TODO(elliot): Check all elements are the same kind.
		valueRegisters, valueKind, _ := compileExpr(compiledFunc, value, file, scopeOverrides)
		if n.Kind == nil {
			n.Kind = valueKind[0].ToArray()
		}

		compiledFunc.Append(&vm.ArraySet{
			Array: arrayRegister,
			Index: indexRegister,
			Value: valueRegisters[0],
		})
	}

	typeRegister := file.AddType(n.Kind)
	arrayAlloc.Kind = typeRegister

	return arrayRegister, n.Kind, nil
}
