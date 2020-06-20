package compiler

import (
	"fmt"
	"ok/ast"
	"ok/vm"
)

func compileCase(compiledFunc *vm.CompiledFunc, n *ast.Case, valueRegister, expectedConditionKind string, afterMatch, breakIns, continueIns vm.Instruction, fns map[string]*ast.Func) error {
	// TODO(elliot): This is a poor solution. It simply expands the conditions
	//  out as if they were individual case statements. This duplicates
	//  statements and uses more memory.

	for _, condition := range n.Conditions {
		conditionResults, conditionKind, err := compileExpr(compiledFunc, condition, fns)
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
			result := compiledFunc.NextRegister()

			op := fmt.Sprintf("%s == %s", conditionKind, conditionKind)
			bop, _ := getBinaryInstruction(op, valueRegister, conditionResults[0], result)
			compiledFunc.Append(bop)

			conditionResults = []string{result}
		}

		ins := &vm.JumpUnless{
			Condition: conditionResults[0],
			To:        -1, // This is corrected at the end.
		}
		compiledFunc.Append(ins)

		err = compileBlock(compiledFunc, n.Statements, breakIns, continueIns, fns)
		if err != nil {
			return err
		}

		compiledFunc.Append(afterMatch)

		// Correct case jump. This is the jump to the next case statement. Or,
		// if it's the last case it will jump to outside the switch.
		ins.To = len(compiledFunc.Instructions) - 1
	}

	return nil
}

func compileSwitch(compiledFunc *vm.CompiledFunc, n *ast.Switch, breakIns, continueIns vm.Instruction, fns map[string]*ast.Func) error {
	afterMatch := &vm.Jump{
		To: -1, // Corrected later.
	}

	// Condition must be a bool if no value has been provided, otherwise all
	// conditions must be the same type as the value.
	expectedConditionKind := "bool"
	valueRegisters := []string{""}
	if n.Expr != nil {
		var err error
		valueRegisters, expectedConditionKind, err = compileExpr(compiledFunc, n.Expr, fns)
		if err != nil {
			return err
		}
	}

	for _, caseStmt := range n.Cases {
		err := compileCase(compiledFunc, caseStmt, valueRegisters[0], expectedConditionKind, afterMatch, breakIns, continueIns, fns)
		if err != nil {
			return err
		}
	}

	err := compileBlock(compiledFunc, n.Else, breakIns, continueIns, fns)
	if err != nil {
		return err
	}

	afterMatch.To = len(compiledFunc.Instructions) - 1

	return nil
}
