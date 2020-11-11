package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

// compileExpr return the (result register, result type, error)
func compileExpr(
	compiledFunc *vm.CompiledFunc,
	expr ast.Node,
	file *vm.File,
) ([]vm.Register, []*types.Type, error) {
	switch e := expr.(type) {
	case *ast.Assign:
		err := compileAssign(compiledFunc, e, file)

		return nil, nil, err

	case *ast.Literal:
		returns, kind := compileLiteral(compiledFunc, e)

		return []vm.Register{returns}, []*types.Type{kind}, nil

	case *ast.Func:
		return compileFunc(compiledFunc, e)

	case *ast.Array:
		returns, kind, err := compileArray(compiledFunc, e, file)
		if err != nil {
			return nil, nil, err
		}

		return []vm.Register{returns}, []*types.Type{kind}, nil

	case *ast.Map:
		returns, kind, err := compileMap(compiledFunc, e, file)
		if err != nil {
			return nil, nil, err
		}

		return []vm.Register{returns}, []*types.Type{kind}, nil

	case *ast.Call:
		results, resultKinds, err := compileCall(compiledFunc, e, file)
		if err != nil {
			return nil, nil, err
		}

		// TODO(elliot): Doesn't return kind.
		return results, resultKinds, nil

	case *ast.Identifier:
		return compileIdentifier(compiledFunc, e, file)

	case *ast.Binary:
		result, ty, err := compileBinary(compiledFunc, e, file)
		if err != nil {
			return nil, nil, err
		}

		return []vm.Register{result}, []*types.Type{ty}, nil

	case *ast.Group:
		return compileExpr(compiledFunc, e.Expr, file)

	case *ast.Unary:
		result, ty, err := compileUnary(compiledFunc, e, file)
		if err != nil {
			return nil, nil, err
		}

		return []vm.Register{result}, []*types.Type{ty}, nil

	case *ast.Key:
		result, ty, err := compileKey(compiledFunc, e, file)
		if err != nil {
			return nil, nil, err
		}

		return []vm.Register{result}, []*types.Type{ty}, nil

	case *ast.Interpolate:
		result, err := compileInterpolate(compiledFunc, e, file)
		if err != nil {
			return nil, nil, err
		}

		return []vm.Register{result}, []*types.Type{types.String}, nil
	}

	panic(expr)
}
