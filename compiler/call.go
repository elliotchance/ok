package compiler

import (
	"ok/ast"
	"ok/instruction"
	"os"
)

func compileCall(compiledFunc *CompiledFunc, call *ast.Call) (string, error) {
	if call.FunctionName == "len" {
		argumentResult, _, err := compileExpr(compiledFunc, call.Arguments[0])
		if err != nil {
			return "", err
		}

		result := compiledFunc.nextRegister()
		ins := &instruction.Len{
			Argument: argumentResult,
			Result:   result,
		}

		compiledFunc.append(ins)

		return result, nil
	}

	ins := &instruction.Print{
		Stdout: os.Stdout,
	}
	for _, arg := range call.Arguments {
		returns, _, err := compileExpr(compiledFunc, arg)
		if err != nil {
			return "", err
		}

		ins.Arguments = append(ins.Arguments, returns)
	}

	compiledFunc.append(ins)

	return "", nil
}
