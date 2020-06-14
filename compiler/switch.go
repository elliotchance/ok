package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
)

func compileCase(compiledFunc *CompiledFunc, n *ast.Case, valueRegister, expectedConditionKind string, afterMatch, breakIns, continueIns instruction.Instruction) error {
	// TODO(elliot): This is a poor solution. It simply expands the conditions
	//  out as if they were individual case statements. This duplicates
	//  statements and uses more memory.

	for _, condition := range n.Conditions {
		conditionResult, conditionKind, err := compileExpr(compiledFunc, condition)
		if err != nil {
			return err
		}

		if conditionKind != expectedConditionKind {
			return fmt.Errorf(
				"expression in case condition must be %s, got %s",
				expectedConditionKind, conditionKind)
		}

		// If we are comparing against a value we need to add the equality step.
		if valueRegister != "" {
			result := compiledFunc.nextRegister()

			op := fmt.Sprintf("%s == %s", conditionKind, conditionKind)
			bop, _ := getBinaryInstruction(op, valueRegister, conditionResult, result)
			compiledFunc.append(bop)

			conditionResult = result
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

		// Correct case jump. This is the jump to the next case statement. Or,
		// if it's the last case it will jump to outside the switch.
		ins.To = len(compiledFunc.Instructions) - 1
	}

	return nil
}

func compileSwitch(compiledFunc *CompiledFunc, n *ast.Switch, breakIns, continueIns instruction.Instruction) error {
	afterMatch := &instruction.Jump{
		To: -1, // Corrected later.
	}

	// Condition must be a bool if no value has been provided, otherwise all
	// conditions must be the same type as the value.
	expectedConditionKind := "bool"
	var valueRegister string
	if n.Expr != nil {
		var err error
		valueRegister, expectedConditionKind, err = compileExpr(compiledFunc, n.Expr)
		if err != nil {
			return err
		}
	}

	for _, caseStmt := range n.Cases {
		err := compileCase(compiledFunc, caseStmt, valueRegister, expectedConditionKind, afterMatch, breakIns, continueIns)
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
