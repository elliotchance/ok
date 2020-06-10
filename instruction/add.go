package instruction

import (
	"ok/ast"
	"ok/number"
)

// Add will sum two numbers.
type Add struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Add) Execute() error {
	ins.Result = ast.NewLiteralNumber(
		number.Add(
			number.NewNumber(ins.Left.Value),
			number.NewNumber(ins.Right.Value),
		).String(),
	)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Add) Answer() *ast.Literal {
	return ins.Result
}
