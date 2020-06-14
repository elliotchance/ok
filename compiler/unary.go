package compiler

import (
	"ok/ast"
	"ok/instruction"
)

func compileUnary(compiledFunc *CompiledFunc, e *ast.Unary) (string, string, error) {
	returns1, kind, err := compileExpr(compiledFunc, e.Expr)
	if err != nil {
		return "", "", err
	}

	var ins instruction.Instruction
	switch e.Op {
	case "not":
		returns2 := compiledFunc.nextRegister()
		ins = &instruction.Not{
			Left:   returns1,
			Result: returns2,
		}
		compiledFunc.append(ins)

		return returns2, kind, nil

	case "-":
		zeroAt := compiledFunc.nextRegister()
		compiledFunc.append(&instruction.Assign{
			VariableName: zeroAt,
			Value:        ast.NewLiteralNumber("0"),
		})

		returns2 := compiledFunc.nextRegister()
		ins = &instruction.Subtract{
			Left:   zeroAt,
			Right:  returns1,
			Result: returns2,
		}
		compiledFunc.append(ins)

		return returns2, kind, nil

	case "++":
		oneAt := compiledFunc.nextRegister()
		compiledFunc.append(&instruction.Assign{
			VariableName: oneAt,
			Value:        ast.NewLiteralNumber("1"),
		})

		ins = &instruction.Add{
			Left:   returns1,
			Right:  oneAt,
			Result: returns1,
		}
		compiledFunc.append(ins)

		return returns1, kind, nil

	case "--":
		oneAt := compiledFunc.nextRegister()
		compiledFunc.append(&instruction.Assign{
			VariableName: oneAt,
			Value:        ast.NewLiteralNumber("1"),
		})

		ins = &instruction.Subtract{
			Left:   returns1,
			Right:  oneAt,
			Result: returns1,
		}
		compiledFunc.append(ins)

		return returns1, kind, nil
	}

	panic(e.Op)
}
