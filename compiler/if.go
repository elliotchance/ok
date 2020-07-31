package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

func compileIf(compiledFunc *vm.CompiledFunc, n *ast.If, breakIns, continueIns vm.Instruction, file *Compiled) error {
	conditionResults, conditionKinds, err := compileExpr(compiledFunc, n.Condition, file)
	if err != nil {
		return err
	}

	if conditionKinds[0] != "bool" {
		return fmt.Errorf(
			"expression in if condition must be a bool, got %s",
			conditionKinds[0])
	}

	ins := &vm.JumpUnless{
		Condition: conditionResults[0],
		To:        -1,
	}
	compiledFunc.Append(ins)

	err = compileBlock(compiledFunc, n.True, breakIns, continueIns, file)
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

		err = compileBlock(compiledFunc, n.False, breakIns, continueIns, file)
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
