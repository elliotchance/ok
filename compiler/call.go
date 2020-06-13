package compiler

import (
	"ok/ast"
	"ok/instruction"
	"os"
)

func compileCall(compiledFunc *CompiledFunc, call *ast.Call) error {
	ins := &instruction.Print{
		Stdout: os.Stdout,
	}
	for _, arg := range call.Arguments {
		returns, _, err := compileExpr(compiledFunc, arg)
		if err != nil {
			return err
		}

		ins.Arguments = append(ins.Arguments, returns)
	}

	compiledFunc.append(ins)

	return nil
}
