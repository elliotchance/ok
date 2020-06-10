package instruction

import (
	"ok/ast"
	"ok/number"
)

// Subtract will subtract two numbers.
type Subtract struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Subtract) Execute(registers map[string]*ast.Literal) error {
	registers[ins.Result] = ast.NewLiteralNumber(
		number.Subtract(
			number.NewNumber(registers[ins.Left].Value),
			number.NewNumber(registers[ins.Right].Value),
		).String(),
	)

	return nil
}
