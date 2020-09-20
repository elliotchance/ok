package compiler

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileKey(
	compiledFunc *vm.CompiledFunc,
	n *ast.Key,
	file *vm.File,
) (vm.Register, *types.Type, error) {
	// It could be an imported constant.
	// TODO(elliot): This is hack for now, because the package name is not yet a
	//  variable.
	if node, ok := n.Expr.(*ast.Identifier); ok && vm.Packages[node.Name] != nil {
		if key, ok := n.Key.(*ast.Identifier); ok {
			if c, ok := vm.Packages[node.Name].Constants[key.Name]; ok {
				resultRegister := compiledFunc.NextRegister()

				// Copy the value in case it's modified.
				//
				// TODO(elliot): This does not support non-scalar values.
				compiledFunc.Append(&vm.Assign{
					VariableName: resultRegister,
					Value: &ast.Literal{
						Kind:  c.Kind,
						Value: c.Value,
					},
				})

				return resultRegister, c.Kind, nil
			}
		}
	}

	arrayOrMapRegisters, arrayOrMapKind, err := compileExpr(compiledFunc, n.Expr, file)
	if err != nil {
		return "", nil, err
	}

	// TODO(elliot): This can be removed once the compiler can understand Key
	//  expressions better.
	key := n.Key
	switch arrayOrMapKind[0].Kind {
	case types.KindResolvedInterface, types.KindUnresolvedInterface:
		// TODO(elliot): This should not allow KindUnresolvedInterface. It only
		//  exists here for now as a hack to permit errors.Error to work.

		if k, ok := key.(*ast.Identifier); ok {
			key = asttest.NewLiteralString(k.Name)
		}
	}

	// TODO(elliot): Check key is the correct type.
	keyRegisters, _, err := compileExpr(compiledFunc, key, file)
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

	case types.KindUnresolvedInterface:
		// TODO(elliot): This should not allow KindUnresolvedInterface. It only
		//  exists here for now as a hack to permit errors.Error to work.

		compiledFunc.Append(&vm.MapGet{
			Map:    arrayOrMapRegisters[0],
			Key:    keyRegisters[0],
			Result: resultRegister,
		})

		if fn, ok := file.FuncDefs[arrayOrMapKind[0].Name]; ok {
			if !fn.IsConstructor() {
				return "", nil, fmt.Errorf("%s %s is not an interface",
					n.Position(), fn.Name)
			}

			return resultRegister, fn.Returns[0], nil
		}

		parts := strings.Split(arrayOrMapKind[0].Name, ".")
		if len(parts) == 2 {
			if fn, ok := vm.Packages[parts[0]].FuncDefs[parts[1]]; ok {
				if !fn.IsConstructor() {
					return "", nil, fmt.Errorf("%s %s is not an interface",
						n.Position(), fn.Name)
				}

				return resultRegister, fn.Returns[0], nil
			}
		}

		return "", nil, fmt.Errorf("%s unknown type: %s",
			n.Position(), arrayOrMapKind[0])

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
