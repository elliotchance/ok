package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func compileFor(compiledFunc *CompiledFunc, n *ast.For) error {
	condition := n.Condition
	if condition == nil {
		condition = ast.NewLiteralBool(true)
	}

	conditionResult, conditionKind, err := compileExpr(compiledFunc, condition)
	if err != nil {
		return err
	}

	if conditionKind != "bool" {
		return fmt.Errorf(
			"expression in for condition must be a bool, got %s",
			conditionKind)
	}

	// "-2" because we need to jump before the previous instruction + we
	// need to offset the +1 that will always occur after an instruction.
	conditionPosition := len(compiledFunc.Instructions) - 2

	ins := &instruction.JumpUnless{
		Condition: conditionResult,
		To:        -1,
	}
	compiledFunc.append(ins)

	breakIns := &instruction.Jump{
		To: 0, // This is corrected later on.
	}
	continueIns := &instruction.Jump{
		To: 0, // This is corrected later on.
	}
	err = compileBlock(compiledFunc, n.Statements, breakIns, continueIns)
	if err != nil {
		return err
	}

	// Jump back to the condition.
	compiledFunc.append(&instruction.Jump{
		To: conditionPosition,
	})

	// Correct condition jump. The "-1" is to correct for the "+1" that would
	// happen after every instruction.
	ins.To = len(compiledFunc.Instructions) - 1

	// Correct the continue instruction.
	continueIns.To = conditionPosition

	// Correct the break instruction.
	breakIns.To = len(compiledFunc.Instructions) - 1

	return nil
}
