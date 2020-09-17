package compiler

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
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
		if key, ok := n.Key.(*ast.Literal); ok {
			if c, ok := vm.Packages[node.Name].Constants[key.Value]; ok {
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

	// TODO(elliot): Check key is the correct type.
	keyRegisters, _, err := compileExpr(compiledFunc, n.Key, file)
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

	case types.KindResolvedInterface, types.KindUnresolvedInterface:
		// TODO(elliot): This should not allow KindUnresolvedInterface. It only
		//  exists here for now as a hack to permit errors.Error to work.

		compiledFunc.Append(&vm.MapGet{
			Map:    arrayOrMapRegisters[0],
			Key:    keyRegisters[0],
			Result: resultRegister,
		})

		if iface, ok := file.Interfaces[arrayOrMapKind[0].Name]; ok {
			ty := iface[n.Key.(*ast.Literal).Value]

			return resultRegister, ty, nil
		}

		parts := strings.Split(arrayOrMapKind[0].Name, ".")
		if len(parts) == 2 {
			if iface, ok := vm.Packages[parts[0]].Interfaces[parts[1]]; ok {
				ty := iface[n.Key.(*ast.Literal).Value]

				return resultRegister, ty, nil
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
