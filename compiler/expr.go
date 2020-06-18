package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

// compileExpr return the (result register, result type, error)
func compileExpr(compiledFunc *CompiledFunc, expr ast.Node) (string, string, error) {
	switch e := expr.(type) {
	case *ast.Literal:
		returns := compiledFunc.nextRegister()
		compiledFunc.append(&instruction.Assign{
			VariableName: returns,
			Value:        e,
		})

		return returns, e.Kind, nil

	case *ast.Array:
		returns, err := compileArray(compiledFunc, e)
		if err != nil {
			return "", "", err
		}

		// FIXME: doesn't return type
		return returns, "", nil

	case *ast.Call:
		err := compileCall(compiledFunc, e)
		if err != nil {
			return "", "", err
		}

		// FIXME: doesn't return register
		return "", "", nil

	case *ast.Identifier:
		if v, ok := compiledFunc.variables[e.Name]; ok {
			return e.Name, v, nil
		}

		return "", "", fmt.Errorf("undefined variable: %s", e.Name)

	case *ast.Binary:
		return compileBinary(compiledFunc, e)

	case *ast.Group:
		return compileExpr(compiledFunc, e.Expr)

	case *ast.Unary:
		return compileUnary(compiledFunc, e)

	case *ast.Key:
		return compileKey(compiledFunc, e)
	}

	panic(expr)
}
