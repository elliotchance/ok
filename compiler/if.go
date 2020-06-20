package compiler

import (
	"fmt"
	"ok/ast"
	"ok/vm"
)

func compileIf(compiledFunc *vm.CompiledFunc, n *ast.If, breakIns, continueIns vm.Instruction, fns map[string]*ast.Func) error {
	conditionResults, conditionKind, err := compileExpr(compiledFunc, n.Condition, fns)
	if err != nil {
		return err
	}

	if conditionKind != "bool" {
		return fmt.Errorf(
			"expression in if condition must be a bool, got %s",
			conditionKind)
	}

	ins := &vm.JumpUnless{
		Condition: conditionResults[0],
		To:        -1,
	}
	compiledFunc.Append(ins)

	err = compileBlock(compiledFunc, n.True, breakIns, continueIns, fns)
	if err != nil {
		return err
	}

	// Else
	if len(n.False) > 0 {
		elseIns := &vm.Jump{
			To: -1,
		}
		compiledFunc.Append(elseIns)

		ins.To = len(compiledFunc.Instructions) - 1

		err = compileBlock(compiledFunc, n.False, breakIns, continueIns, fns)
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
