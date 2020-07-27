package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Multiply will multiply two numbers.
type Multiply struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Multiply) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralNumber(
		number.Multiply(
			number.NewNumber(registers[ins.Left].Value),
			number.NewNumber(registers[ins.Right].Value),
		).String(),
	)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Multiply) String() string {
	return fmt.Sprintf("%s = %s * %s", ins.Result, ins.Left, ins.Right)
}
