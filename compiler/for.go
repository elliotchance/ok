package compiler

import (
	"fmt"
	"ok/ast"
	"ok/instruction"
	"strings"
)

func compileFor(compiledFunc *CompiledFunc, n *ast.For) error {
	// There's nothing special about Init here. It just executes once before the
	// loop.
	if n.Init != nil {
		_, _, err := compileExpr(compiledFunc, n.Init)
		if err != nil {
			return err
		}
	}

	var conditionResult string
	switch cond := n.Condition.(type) {
	case nil:
		// Error here should not be possible.
		conditionResult, _, _ = compileExpr(compiledFunc, ast.NewLiteralBool(true))

	case *ast.In:
		// TODO(elliot): Check arrayOrMapResult is actually an array or map.
		arrayOrMapResult, arrayOrMapKind, err := compileExpr(compiledFunc, cond.Expr)
		if err != nil {
			return err
		}

		// TODO(elliot): Not always a number.
		compiledFunc.newVariable(cond.Key, "number")

		if cond.Value != "" {
			// TODO(elliot): Not always a number.
			compiledFunc.newVariable(cond.Value, "number")
		}

		cursorRegister := compiledFunc.nextRegister()
		compiledFunc.append(&instruction.Assign{
			VariableName: cursorRegister,
			Value:        ast.NewLiteralNumber("0"),
		})

		conditionResult = compiledFunc.nextRegister()
		if strings.HasPrefix(arrayOrMapKind, "[]") {
			compiledFunc.append(&instruction.NextArray{
				Array:       arrayOrMapResult,
				Cursor:      cursorRegister,
				KeyResult:   cond.Key,
				ValueResult: cond.Value,
				Result:      conditionResult,
			})
		} else {
			compiledFunc.append(&instruction.NextMap{
				Map:         arrayOrMapResult,
				Cursor:      cursorRegister,
				KeyResult:   cond.Key,
				ValueResult: cond.Value,
				Result:      conditionResult,
			})
		}

	default:
		var conditionKind string
		var err error
		conditionResult, conditionKind, err = compileExpr(compiledFunc, n.Condition)
		if err != nil {
			return err
		}

		if conditionKind != "bool" {
			return fmt.Errorf(
				"expression in for condition must be a bool, got %s",
				conditionKind)
		}
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
	err := compileBlock(compiledFunc, n.Statements, breakIns, continueIns)
	if err != nil {
		return err
	}

	// Next is optional and runs at the end of each iteration even if continue
	// is called. However, it does not run is break is called.
	if n.Next != nil {
		// Since there is a Next we need the continue to jump to the end of the
		// iteration where the Next is before it jumps back to the start of the
		// next iteration.
		continueIns.To = len(compiledFunc.Instructions) - 1

		_, _, err := compileExpr(compiledFunc, n.Next)
		if err != nil {
			return err
		}
	} else {
		// Since there is no Next, the continue can jump right to the start of
		// the iteration.
		continueIns.To = conditionPosition
	}

	// Jump back to the condition.
	compiledFunc.append(&instruction.Jump{
		To: conditionPosition,
	})

	// Correct condition jump. The "-1" is to correct for the "+1" that would
	// happen after every instruction.
	ins.To = len(compiledFunc.Instructions) - 1

	// Correct the break instruction.
	breakIns.To = len(compiledFunc.Instructions) - 1

	return nil
}
