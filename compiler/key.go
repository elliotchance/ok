package compiler

import (
	"ok/ast"
	"ok/instruction"
)

func compileKey(compiledFunc *CompiledFunc, n *ast.Key) (string, string, error) {
	arrayRegister, _, err := compileExpr(compiledFunc, n.Expr)
	if err != nil {
		return "", "", err
	}

	// TODO(elliot): Check key is the correct type.
	keyRegister, _, err := compileExpr(compiledFunc, n.Key)
	if err != nil {
		return "", "", err
	}

	resultRegister := compiledFunc.nextRegister()
	compiledFunc.append(&instruction.ArrayGet{
		Array:  arrayRegister,
		Index:  keyRegister,
		Result: resultRegister,
	})

	// TODO(elliot): Does not return element type.
	return resultRegister, "", nil
}
