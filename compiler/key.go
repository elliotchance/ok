package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler/kind"
	"github.com/elliotchance/ok/vm"
)

func compileKey(compiledFunc *vm.CompiledFunc, n *ast.Key, file *Compiled) (vm.Register, string, error) {
	// It could be an imported constant.
	// TODO(elliot): This is hack for now, because the package name is not yet a
	//  variable.
	if node, ok := n.Expr.(*ast.Identifier); ok {
		if key, ok := n.Key.(*ast.Literal); ok {
			if c, ok := vm.Constants[fmt.Sprintf("%s.%s", node.Name, key.Value)]; ok {
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

		if iface, ok := file.Interfaces[arrayOrMapKind[0]]; ok {
			ty := iface[n.Key.(*ast.Literal).Value]

			return resultRegister, ty, nil
		}

		if iface, ok := vm.Interfaces[arrayOrMapKind[0]]; ok {
			ty := iface[n.Key.(*ast.Literal).Value]

			return resultRegister, ty, nil
		}

		return "", "", fmt.Errorf("unknown type: %s", arrayOrMapKind[0])

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
