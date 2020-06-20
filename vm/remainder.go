package vm

import (
	"ok/ast"
	"ok/number"
)

// Remainder will return the remainder when dividing two numbers. This is not
// the same as a modulo. A remainder may be negative.
type Remainder struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Remainder) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	divide, err := number.Remainder(
		number.NewNumber(registers[ins.Left].Value),
		number.NewNumber(registers[ins.Right].Value),
	)
	if err != nil {
		// TODO(elliot): This needs to be the same precision of zero.
		registers[ins.Result] = ast.NewLiteralNumber("0")
		return err
	}

	registers[ins.Result] = ast.NewLiteralNumber(divide.String())
	return nil
}
