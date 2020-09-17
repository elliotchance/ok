package compiler

import (
	"fmt"

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
		cf, err := CompileFunc(e, file)
		if err != nil {
			return nil, nil, err
		}

		file.FuncDefs[e.Name] = e
		file.Funcs[e.Name] = cf
		fnType := e.Type()

		returns := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: returns,
			Value: &ast.Literal{
				Kind:  fnType,
				Value: e.Name,
			},
		})

		compiledFunc.Append(&vm.ParentScope{
			X: returns,
		})

		return []vm.Register{returns}, []*types.Type{fnType}, nil

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
		// TODO(elliot): Doesn't check that the upper scope variable exists or
		//  fetches the correct type.
		if e.Name[0] == '^' {
			return []vm.Register{vm.Register(e.Name)}, []*types.Type{types.Number}, nil
		}

		if v, ok := compiledFunc.Variables[e.Name]; ok || e.Name[0] == '^' {
			return []vm.Register{vm.Register(e.Name)}, []*types.Type{v}, nil
		}

		if c, ok := file.Constants[e.Name]; ok {
			// We copy it locally to make sure it's value isn't changed. The
			// compiler will prevent a constant from being modified directly.
			//
			// TODO(elliot): The compiler needs to raise an error when trying to
			//  modify a constant.
			literalRegister := compiledFunc.NextRegister()
			compiledFunc.Append(&vm.Assign{
				VariableName: literalRegister,
				// TODO(elliot): Does not support non-scalar values.
				Value: &ast.Literal{
					Kind:  c.Kind,
					Value: c.Value,
				},
			})

			return []vm.Register{literalRegister}, []*types.Type{c.Kind}, nil
		}

		// It could also reference a package-level function.
		if fn, ok := file.FuncDefs[e.Name]; ok {
			literalRegister := compiledFunc.NextRegister()
			compiledFunc.Append(&vm.Assign{
				VariableName: literalRegister,
				Value: &ast.Literal{
					Kind:  fn.Type(),
					Value: fn.Name,
				},
			})

			return []vm.Register{literalRegister}, []*types.Type{fn.Type()}, nil
		}

		return nil, nil, fmt.Errorf("%s undefined variable: %s",
			e.Pos, e.Name)

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
