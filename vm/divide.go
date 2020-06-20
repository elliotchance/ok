package vm

import (
	"ok/ast"
	"ok/number"
)

// Divide will multiply two numbers.
type Divide struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Divide) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	divide, err := number.Divide(
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
