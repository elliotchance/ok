package instruction

import (
	"ok/ast"
	"ok/number"
)

// Remainder will return the remainder when dividing two numbers. This is not
// the same as a modulo. A remainder may be negative.
type Remainder struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Remainder) Execute() error {
	divide, err := number.Remainder(
		number.NewNumber(ins.Left.Value),
		number.NewNumber(ins.Right.Value),
	)
	if err != nil {
		return err
	}

	ins.Result = ast.NewLiteralNumber(divide.String())
	return nil
}
