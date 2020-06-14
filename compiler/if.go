package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func compileIf(compiledFunc *CompiledFunc, n *ast.If, breakIns, continueIns instruction.Instruction) error {
	conditionResult, conditionKind, err := compileExpr(compiledFunc, n.Condition)
	if err != nil {
		return err
	}

	if conditionKind != "bool" {
		return fmt.Errorf(
			"expression in if condition must be a bool, got %s",
			conditionKind)
	}

	ins := &instruction.JumpUnless{
		Condition: conditionResult,
		To:        -1,
	}
	compiledFunc.append(ins)

	err = compileBlock(compiledFunc, n.True, breakIns, continueIns)
	if err != nil {
		return err
	}

	// Else
	if len(n.False) > 0 {
		elseIns := &instruction.Jump{
			To: -1,
		}
		compiledFunc.append(elseIns)

		ins.To = len(compiledFunc.Instructions) - 1

		err = compileBlock(compiledFunc, n.False, breakIns, continueIns)
		if err != nil {
			return err
		}

		// Correct else jump. The "-1" is to correct for the "+1" that would
		// happen after every instruction.
		elseIns.To = len(compiledFunc.Instructions) - 1
	} else {
		// Correct condition jump. The "-1" is to correct for the "+1" that would
		// happen after every instruction.
		ins.To = len(compiledFunc.Instructions) - 1
	}

	return nil
}
