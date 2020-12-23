package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compileIf(
	compiledFunc *vm.CompiledFunc,
	n *ast.If,
	breakIns,
	continueIns vm.Instruction,
	file *vm.File,
	scopeOverrides map[string]*types.Type,
) error {
	conditionResults, conditionKinds, err := compileExpr(compiledFunc,
		n.Condition, file, scopeOverrides)
	if err != nil {
		return err
	}

	if conditionKinds[0].Kind != types.KindBool {
		return fmt.Errorf(
			"expression in if condition must be a bool, got %s",
			conditionKinds[0])
	}

	ins := &vm.JumpUnless{
		Condition: conditionResults[0],
		To:        -1,
	}
	compiledFunc.Append(ins)

	// The true scope may contain an override, but the false scope must not.
	trueScope := scopeOverrides
	if condition, ok := n.Condition.(*ast.Binary); ok && condition.Op == lexer.TokenIs {
		trueScope = appendScope(trueScope,
			condition.Left.(*ast.Identifier).Name,
			types.TypeFromString(condition.Right.(*ast.Identifier).Name))
	}

	err = compileBlock(compiledFunc, n.True, breakIns, continueIns, file,
		trueScope)
	if err != nil {
		return err
	}

	// Else
	if len(n.False) > 0 {
		elseIns := &vm.Jump{
			To: -1,
		}
		compiledFunc.Append(elseIns)

		ins.To = len(compiledFunc.Instructions.Instructions) - 1

		err = compileBlock(compiledFunc, n.False, breakIns, continueIns, file,
			scopeOverrides)
		if err != nil {
			return err
		}

		// Correct else jump. The "-1" is to correct for the "+1" that would
		// happen after every instruction.
		elseIns.To = len(compiledFunc.Instructions.Instructions) - 1
	} else {
		// Correct condition jump. The "-1" is to correct for the "+1" that would
		// happen after every instruction.
		ins.To = len(compiledFunc.Instructions.Instructions) - 1
	}

	return nil
}

func appendScope(
	scope map[string]*types.Type,
	name string,
	ty *types.Type,
) map[string]*types.Type {
	newScope := map[string]*types.Type{
		name: ty,
	}
	for k, v := range scope {
		newScope[k] = v
	}

	return newScope
}
