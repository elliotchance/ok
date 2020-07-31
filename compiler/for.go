package compiler

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler/kind"
	"github.com/elliotchance/ok/vm"
)

func compileFor(compiledFunc *vm.CompiledFunc, n *ast.For, file *Compiled) error {
	// There's nothing special about Init here. It just executes once before the
	// loop.
	if n.Init != nil {
		_, _, err := compileExpr(compiledFunc, n.Init, file)
		if err != nil {
			return err
		}
	}

	var conditionResults []vm.Register
	switch cond := n.Condition.(type) {
	case nil:
		// Error here should not be possible.
		conditionResults, _, _ = compileExpr(compiledFunc, asttest.NewLiteralBool(true), file)

	case *ast.In:
		arrayOrMapResults, arrayOrMapKind, err := compileExpr(compiledFunc, cond.Expr, file)
		if err != nil {
			return err
		}

		switch {
		case kind.IsArray(arrayOrMapKind[0]),
			kind.IsMap(arrayOrMapKind[0]),
			arrayOrMapKind[0] == "string":
			// Allowed

		default:
			return fmt.Errorf("%s is not iterable", arrayOrMapKind[0])
		}

		compiledFunc.NewVariable(cond.Value, kind.ElementType(arrayOrMapKind[0]))

		if cond.Key != "" {
			switch {
			case kind.IsArray(arrayOrMapKind[0]):
				compiledFunc.NewVariable(cond.Key, "number")

			case kind.IsMap(arrayOrMapKind[0]):
				compiledFunc.NewVariable(cond.Key, "string")

			case arrayOrMapKind[0] == "string":
				compiledFunc.NewVariable(cond.Key, "char")
			}
		}

		cursorRegister := compiledFunc.NextRegister()
		compiledFunc.Append(&vm.Assign{
			VariableName: cursorRegister,
			Value:        asttest.NewLiteralNumber("0"),
		})

		conditionResults = []vm.Register{compiledFunc.NextRegister()}
		switch {
		case kind.IsArray(arrayOrMapKind[0]):
			compiledFunc.Append(&vm.NextArray{
				Array:       arrayOrMapResults[0],
				Cursor:      cursorRegister,
				KeyResult:   vm.Register(cond.Key),
				ValueResult: vm.Register(cond.Value),
				Result:      conditionResults[0],
			})

		case kind.IsMap(arrayOrMapKind[0]):
			compiledFunc.Append(&vm.NextMap{
				Map:         arrayOrMapResults[0],
				Cursor:      cursorRegister,
				KeyResult:   vm.Register(cond.Key),
				ValueResult: vm.Register(cond.Value),
				Result:      conditionResults[0],
			})

		case arrayOrMapKind[0] == "string":
			compiledFunc.Append(&vm.NextString{
				Str:         arrayOrMapResults[0],
				Cursor:      cursorRegister,
				KeyResult:   vm.Register(cond.Key),
				ValueResult: vm.Register(cond.Value),
				Result:      conditionResults[0],
			})
		}

	default:
		var conditionKinds []string
		var err error
		conditionResults, conditionKinds, err = compileExpr(compiledFunc, n.Condition, file)
		if err != nil {
			return err
		}

		if conditionKinds[0] != "bool" {
			return fmt.Errorf(
				"expression in for condition must be a bool, got %s",
				conditionKinds[0])
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
	err := compileBlock(compiledFunc, n.Statements, breakIns, continueIns, file)
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

		_, _, err := compileExpr(compiledFunc, n.Next, file)
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
