package instruction

import (
	"ok/ast"
	"ok/number"
)

// Divide will multiply two numbers.
type Divide struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Divide) Execute() error {
	divide, err := number.Divide(
		number.NewNumber(ins.Left.Value),
		number.NewNumber(ins.Right.Value),
	)
	if err != nil {
		return err
	}

	ins.Result = ast.NewLiteralNumber(divide.String())
	return nil
}
