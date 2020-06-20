package compiler

import (
	"ok/ast"
	"ok/vm"
)

func compileUnary(compiledFunc *vm.CompiledFunc, e *ast.Unary, fns map[string]*ast.Func) (string, string, error) {
	returns1, kind, err := compileExpr(compiledFunc, e.Expr, fns)
	if err != nil {
		return "", "", err
	}

	var ins vm.Instruction
	switch e.Op {
	case "not":
		returns2 := compiledFunc.NextRegister()
		ins = &vm.Not{
			Left:   returns1[0],
			Result: returns2,
		}
		compiledFunc.Append(ins)

		return returns2, kind, nil

	case "-":
		zeroAt := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: zeroAt,
			Value:        ast.NewLiteralNumber("0"),
		})

		returns2 := compiledFunc.NextRegister()
		ins = &vm.Subtract{
			Left:   zeroAt,
			Right:  returns1[0],
			Result: returns2,
		}
		compiledFunc.Append(ins)

		return returns2, kind, nil

	case "++":
		oneAt := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: oneAt,
			Value:        ast.NewLiteralNumber("1"),
		})

		ins = &vm.Add{
			Left:   returns1[0],
			Right:  oneAt,
			Result: returns1[0],
		}
		compiledFunc.Append(ins)

		return returns1[0], kind, nil

	case "--":
		oneAt := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: oneAt,
			Value:        ast.NewLiteralNumber("1"),
		})

		ins = &vm.Subtract{
			Left:   returns1[0],
			Right:  oneAt,
			Result: returns1[0],
		}
		compiledFunc.Append(ins)

		return returns1[0], kind, nil
	}

	panic(e.Op)
}
