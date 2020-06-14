package instruction

import (
	"ok/ast"
	"ok/number"
)

// Multiply will multiply two numbers.
type Multiply struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Multiply) Execute(registers map[string]*ast.Literal, _ *int) error {
	registers[ins.Result] = ast.NewLiteralNumber(
		number.Multiply(
			number.NewNumber(registers[ins.Left].Value),
			number.NewNumber(registers[ins.Right].Value),
		).String(),
	)

	return nil
}
