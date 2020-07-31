package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileCase(compiledFunc *vm.CompiledFunc, n *ast.Case, valueRegister vm.Register, expectedConditionKind string, afterMatch, breakIns, continueIns vm.Instruction, file *Compiled) error {
	// TODO(elliot): This is a poor solution. It simply expands the conditions
	//  out as if they were individual case statements. This duplicates
	//  statements and uses more memory.

	for _, condition := range n.Conditions {
		conditionResults, conditionKinds, err := compileExpr(compiledFunc, condition, file)
		if err != nil {
			return err
		}

		if conditionKinds[0] != expectedConditionKind {
			return fmt.Errorf(
				"expression in case condition must be %s, got %s",
				expectedConditionKind, conditionKinds[0])
		}

		// If we are comparing against a value we need to add the equality step.
		if valueRegister != "" {
			result := compiledFunc.NextRegister()

			op := fmt.Sprintf("%s == %s", conditionKinds[0], conditionKinds[0])
			bop, _ := getBinaryInstruction(op, valueRegister, conditionResults[0], result)
			compiledFunc.Append(bop)

			conditionResults = []vm.Register{result}
		}

		ins := &vm.JumpUnless{
			Condition: conditionResults[0],
			To:        -1, // This is corrected at the end.
		}
		compiledFunc.Append(ins)

		err = compileBlock(compiledFunc, n.Statements, breakIns, continueIns, file)
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

func compileSwitch(compiledFunc *vm.CompiledFunc, n *ast.Switch, breakIns, continueIns vm.Instruction, file *Compiled) error {
	afterMatch := &vm.Jump{
		To: -1, // Corrected later.
	}

	// Condition must be a bool if no value has been provided, otherwise all
	// conditions must be the same type as the value.
	expectedConditionKinds := []string{"bool"}
	valueRegisters := []vm.Register{""}
	if n.Expr != nil {
		var err error
		valueRegisters, expectedConditionKinds, err = compileExpr(compiledFunc, n.Expr, file)
		if err != nil {
			return err
		}
	}

	for _, caseStmt := range n.Cases {
		err := compileCase(compiledFunc, caseStmt, valueRegisters[0], expectedConditionKinds[0], afterMatch, breakIns, continueIns, file)
		if err != nil {
			return err
		}
	}

	err := compileBlock(compiledFunc, n.Else, breakIns, continueIns, file)
	if err != nil {
		return err
	}

	afterMatch.To = len(compiledFunc.Instructions) - 1

	return nil
}
