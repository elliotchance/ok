package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Add will sum two numbers.
type Add struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Add) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralNumber(
		number.Add(
			number.NewNumber(registers[ins.Left].Value),
			number.NewNumber(registers[ins.Right].Value),
		).String(),
	)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Add) String() string {
	return fmt.Sprintf("%s = %s + %s", ins.Result, ins.Left, ins.Right)
}
