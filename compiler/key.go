package compiler

import (
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileKey(compiledFunc *vm.CompiledFunc, n *ast.Key, fns map[string]*ast.Func) (string, string, error) {
	arrayOrMapRegisters, arrayOrMapKind, err := compileExpr(compiledFunc, n.Expr, fns)
	if err != nil {
		return "", "", err
	}

	// TODO(elliot): Check key is the correct type.
	keyRegisters, _, err := compileExpr(compiledFunc, n.Key, fns)
	if err != nil {
		return "", "", err
	}

	resultRegister := compiledFunc.NextRegister()
	switch {
	case strings.HasPrefix(arrayOrMapKind[0], "[]"):
		compiledFunc.Append(&vm.ArrayGet{
			Array:  arrayOrMapRegisters[0],
			Index:  keyRegisters[0],
			Result: resultRegister,
		})

	case arrayOrMapKind[0] == "string":
		compiledFunc.Append(&vm.StringIndex{
			String: arrayOrMapRegisters[0],
			Index:  keyRegisters[0],
			Result: resultRegister,
		})

		return resultRegister, "char", nil

	default:
		// This applies for both maps and objects.
		compiledFunc.Append(&vm.MapGet{
			Map:    arrayOrMapRegisters[0],
			Key:    keyRegisters[0],
			Result: resultRegister,
		})
	}

	// TODO(elliot): Does not return element type.
	return resultRegister, "", nil
}
