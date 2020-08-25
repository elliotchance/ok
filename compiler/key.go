package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler/kind"
	"github.com/elliotchance/ok/vm"
)

func compileKey(compiledFunc *vm.CompiledFunc, n *ast.Key, file *Compiled) (vm.Register, string, error) {
	arrayOrMapRegisters, arrayOrMapKind, err := compileExpr(compiledFunc, n.Expr, file)
	if err != nil {
		return "", "", err
	}

	// TODO(elliot): Check key is the correct type.
	keyRegisters, _, err := compileExpr(compiledFunc, n.Key, file)
	if err != nil {
		return "", "", err
	}

	resultRegister := compiledFunc.NextRegister()
	switch {
	case kind.IsArray(arrayOrMapKind[0]):
		compiledFunc.Append(&vm.ArrayGet{
			Array:  arrayOrMapRegisters[0],
			Index:  keyRegisters[0],
			Result: resultRegister,
		})

		return resultRegister, kind.ElementType(arrayOrMapKind[0]), nil

	case kind.IsMap(arrayOrMapKind[0]):
		compiledFunc.Append(&vm.MapGet{
			Map:    arrayOrMapRegisters[0],
			Key:    keyRegisters[0],
			Result: resultRegister,
		})

		return resultRegister, kind.ElementType(arrayOrMapKind[0]), nil

	case kind.IsObject(arrayOrMapKind[0]):
		compiledFunc.Append(&vm.MapGet{
			Map:    arrayOrMapRegisters[0],
			Key:    keyRegisters[0],
			Result: resultRegister,
		})

		// TODO(elliot): Check exists.
		ty := file.Interfaces[arrayOrMapKind[0]][n.Key.(*ast.Literal).Value]

		return resultRegister, ty, nil

	case arrayOrMapKind[0] == "string":
		compiledFunc.Append(&vm.StringIndex{
			Str:    arrayOrMapRegisters[0],
			Index:  keyRegisters[0],
			Result: resultRegister,
		})

		return resultRegister, "char", nil
	}

	panic(arrayOrMapKind[0])
}
