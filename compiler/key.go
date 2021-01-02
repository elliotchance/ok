package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileKey(
	compiledFunc *vm.CompiledFunc,
	n *ast.Key,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) (vm.Register, *types.Type, error) {
	arrayOrMapRegisters, arrayOrMapKind, err := compileExpr(compiledFunc,
		n.Expr, file, scopeOverrides)
	if err != nil {
		return "", nil, err
	}

	// TODO(elliot): This can be removed once the compiler can understand Key
	//  expressions better.
	key := n.Key
	switch arrayOrMapKind[0].Kind {
	case types.KindResolvedInterface:
		if k, ok := key.(*ast.Identifier); ok {
			key = asttest.NewLiteralString(k.Name)
		}
	}

	// TODO(elliot): Check key is the correct type.
	keyRegisters, _, err := compileExpr(compiledFunc, key, file, scopeOverrides)
	if err != nil {
		return "", nil, err
	}

	resultRegister := compiledFunc.NextRegister()
	switch arrayOrMapKind[0].Kind {
	case types.KindArray:
		compiledFunc.Append(&vm.ArrayGet{
			Array:  arrayOrMapRegisters[0],
			Index:  keyRegisters[0],
			Result: resultRegister,
		})

		return resultRegister, arrayOrMapKind[0].Element, nil

	case types.KindMap:
		compiledFunc.Append(&vm.MapGet{
			Map:    arrayOrMapRegisters[0],
			Key:    keyRegisters[0],
			Result: resultRegister,
		})

		return resultRegister, arrayOrMapKind[0].Element, nil

	case types.KindResolvedInterface:
		compiledFunc.Append(&vm.MapGet{
			Map:    arrayOrMapRegisters[0],
			Key:    keyRegisters[0],
			Result: resultRegister,
		})

		// TODO(elliot): Might not be an identifier?
		return resultRegister,
			arrayOrMapKind[0].Properties[n.Key.(*ast.Identifier).Name], nil

	case types.KindString:
		compiledFunc.Append(&vm.StringIndex{
			Str:    arrayOrMapRegisters[0],
			Index:  keyRegisters[0],
			Result: resultRegister,
		})

		return resultRegister, types.Char, nil
	}

	panic(arrayOrMapKind[0])
}
