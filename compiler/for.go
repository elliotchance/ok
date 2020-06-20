package compiler

import (
	"fmt"
	"ok/ast"
	"ok/vm"
	"strings"
)

func compileFor(compiledFunc *vm.CompiledFunc, n *ast.For, fns map[string]*ast.Func) error {
	// There's nothing special about Init here. It just executes once before the
	// loop.
	if n.Init != nil {
		_, _, err := compileExpr(compiledFunc, n.Init, fns)
		if err != nil {
			return err
		}
	}

	var conditionResults []string
	switch cond := n.Condition.(type) {
	case nil:
		// Error here should not be possible.
		conditionResults, _, _ = compileExpr(compiledFunc, ast.NewLiteralBool(true), fns)

	case *ast.In:
		// TODO(elliot): Check arrayOrMapResult is actually an array or map.
		arrayOrMapResults, arrayOrMapKind, err := compileExpr(compiledFunc, cond.Expr, fns)
		if err != nil {
			return err
		}

		// TODO(elliot): Not always a number.
		compiledFunc.NewVariable(cond.Key, "number")

		if cond.Value != "" {
			// TODO(elliot): Not always a number.
			compiledFunc.NewVariable(cond.Value, "number")
		}

		cursorRegister := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: cursorRegister,
			Value:        ast.NewLiteralNumber("0"),
		})

		conditionResults = []string{compiledFunc.NextRegister()}
		if strings.HasPrefix(arrayOrMapKind, "[]") {
			compiledFunc.Append(&vm.NextArray{
				Array:       arrayOrMapResults[0],
				Cursor:      cursorRegister,
				KeyResult:   cond.Key,
				ValueResult: cond.Value,
				Result:      conditionResults[0],
			})
		} else {
			compiledFunc.Append(&vm.NextMap{
				Map:         arrayOrMapResults[0],
				Cursor:      cursorRegister,
				KeyResult:   cond.Key,
				ValueResult: cond.Value,
				Result:      conditionResults[0],
			})
		}

	default:
		var conditionKind string
		var err error
		conditionResults, conditionKind, err = compileExpr(compiledFunc, n.Condition, fns)
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

	ins := &vm.JumpUnless{
		Condition: conditionResults[0],
		To:        -1,
	}
	compiledFunc.Append(ins)

	breakIns := &vm.Jump{
		To: 0, // This is corrected later on.
	}
	continueIns := &vm.Jump{
		To: 0, // This is corrected later on.
	}
	err := compileBlock(compiledFunc, n.Statements, breakIns, continueIns, fns)
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

		_, _, err := compileExpr(compiledFunc, n.Next, fns)
		if err != nil {
			return err
		}
	} else {
		// Since there is no Next, the continue can jump right to the start of
		// the iteration.
		continueIns.To = conditionPosition
	}

	// Jump back to the condition.
	compiledFunc.Append(&vm.Jump{
		To: conditionPosition,
	})

	// Correct condition jump. The "-1" is to correct for the "+1" that would
	// happen after every instruction.
	ins.To = len(compiledFunc.Instructions) - 1

	// Correct the break instruction.
	breakIns.To = len(compiledFunc.Instructions) - 1

	return nil
}
