package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

// compileExpr return the (result register, result type, error)
func compileExpr(compiledFunc *vm.CompiledFunc, expr ast.Node, fns map[string]*ast.Func) ([]string, []string, error) {
	switch e := expr.(type) {
	case *ast.Assign:
		err := compileAssign(compiledFunc, e, fns)

		return nil, nil, err

	case *ast.Literal:
		returns := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: returns,
			Value:        e,
		})

		return []string{returns}, []string{e.Kind}, nil

	case *ast.Array:
		returns, err := compileArray(compiledFunc, e, fns)
		if err != nil {
			return nil, nil, err
		}

		// TODO(elliot): Doesn't return type.
		return []string{returns}, []string{"[]"}, nil

	case *ast.Map:
		returns, err := compileMap(compiledFunc, e, fns)
		if err != nil {
			return nil, nil, err
		}

		// TODO(elliot): Doesn't return type.
		return []string{returns}, []string{"{}"}, nil

	case *ast.Call:
		results, resultKinds, err := compileCall(compiledFunc, e, fns)
		if err != nil {
			return nil, nil, err
		}

		// TODO(elliot): Doesn't return kind.
		return results, resultKinds, nil

	case *ast.Identifier:
		if v, ok := compiledFunc.Variables[e.Name]; ok {
			return []string{e.Name}, []string{v}, nil
		}

		return nil, nil, fmt.Errorf("undefined variable: %s", e.Name)

	case *ast.Binary:
		result, ty, err := compileBinary(compiledFunc, e, fns)
		if err != nil {
			return nil, nil, err
		}

		return []string{result}, []string{ty}, nil

	case *ast.Group:
		return compileExpr(compiledFunc, e.Expr, fns)

	case *ast.Unary:
		result, ty, err := compileUnary(compiledFunc, e, fns)
		if err != nil {
			return nil, nil, err
		}

		return []string{result}, []string{ty}, nil

	case *ast.Key:
		result, ty, err := compileKey(compiledFunc, e, fns)
		if err != nil {
			return nil, nil, err
		}

		return []string{result}, []string{ty}, nil

	case *ast.Interpolate:
		result, err := compileInterpolate(compiledFunc, e, fns)
		if err != nil {
			return nil, nil, err
		}

		return []string{result}, []string{"string"}, nil
	}

	panic(expr)
}
