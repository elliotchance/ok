package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func compileAssign(compiledFunc *CompiledFunc, assign *ast.Assign) error {
	returns, kind, err := compileExpr(compiledFunc, assign.Expr)
	if err != nil {
		return err
	}

	// Make sure we do not assign the wrong type to an existing variable.
	if v, ok := compiledFunc.variables[assign.VariableName]; ok && kind != v {
		return fmt.Errorf("cannot assign %s to variable %s (expecting %s)",
			kind, assign.VariableName, v)
	}

	ins := &instruction.Assign{
		VariableName: assign.VariableName,
		Register:     returns,
	}
	compiledFunc.append(ins)
	compiledFunc.variables[assign.VariableName] = kind

	return nil
}
