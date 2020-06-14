package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func compileCase(compiledFunc *CompiledFunc, n *ast.Case, afterMatch, breakIns, continueIns instruction.Instruction) error {
	// Condition must always be of type bool.
	conditionResult, conditionKind, err := compileExpr(compiledFunc, n.Condition)
	if err != nil {
		return err
	}

	if conditionKind != "bool" {
		return fmt.Errorf(
			"expression in case condition must be a bool, got %s",
			conditionKind)
	}

	ins := &instruction.JumpUnless{
		Condition: conditionResult,
		To:        -1, // This is corrected at the end.
	}
	compiledFunc.append(ins)

	err = compileBlock(compiledFunc, n.Statements, breakIns, continueIns)
	if err != nil {
		return err
	}

	compiledFunc.append(afterMatch)

	// Correct case jump. This is the jump to the next case statement. Or, if
	// it's the last case it will jump to outside the switch.
	ins.To = len(compiledFunc.Instructions) - 1

	return nil
}

func compileSwitch(compiledFunc *CompiledFunc, n *ast.Switch, breakIns, continueIns instruction.Instruction) error {
	afterMatch := &instruction.Jump{
		To: -1, // Corrected later.
	}

	for _, caseStmt := range n.Cases {
		err := compileCase(compiledFunc, caseStmt, afterMatch, breakIns, continueIns)
		if err != nil {
			return err
		}
	}

	err := compileBlock(compiledFunc, n.Else, breakIns, continueIns)
	if err != nil {
		return err
	}

	afterMatch.To = len(compiledFunc.Instructions) - 1

	return nil
}
