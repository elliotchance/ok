package compiler

import (
	"fmt"
	"ok/ast"
	"ok/vm"
)

// compileExpr return the (result register, result type, error)
func compileExpr(compiledFunc *vm.CompiledFunc, expr ast.Node, fns map[string]*ast.Func) ([]string, string, error) {
	switch e := expr.(type) {
	case *ast.Assign:
		err := compileAssign(compiledFunc, e, fns)

		return nil, "", err

	case *ast.Literal:
		returns := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: returns,
			Value:        e,
		})

		return []string{returns}, e.Kind, nil

	case *ast.Array:
		returns, err := compileArray(compiledFunc, e, fns)
		if err != nil {
			return nil, "", err
		}

		// TODO(elliot): Doesn't return type.
		return []string{returns}, "[]", nil

	case *ast.Map:
		returns, err := compileMap(compiledFunc, e, fns)
		if err != nil {
			return nil, "", err
		}

		// TODO(elliot): Doesn't return type.
		return []string{returns}, "{}", nil

	case *ast.Call:
		results, err := compileCall(compiledFunc, e, fns)
		if err != nil {
			return nil, "", err
		}

		// TODO(elliot): Doesn't return kind.
		return results, "", nil

	case *ast.Identifier:
		if v, ok := compiledFunc.Variables[e.Name]; ok {
			return []string{e.Name}, v, nil
		}

		return nil, "", fmt.Errorf("undefined variable: %s", e.Name)

	case *ast.Binary:
		result, ty, err := compileBinary(compiledFunc, e, fns)
		if err != nil {
			return nil, "", err
		}

		return []string{result}, ty, nil

	case *ast.Group:
		return compileExpr(compiledFunc, e.Expr, fns)

	case *ast.Unary:
		result, ty, err := compileUnary(compiledFunc, e, fns)
		if err != nil {
			return nil, "", err
		}

		return []string{result}, ty, nil

	case *ast.Key:
		result, ty, err := compileKey(compiledFunc, e, fns)
		if err != nil {
			return nil, "", err
		}

		return []string{result}, ty, nil
	}

	panic(expr)
}
