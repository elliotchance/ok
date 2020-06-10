package instruction

import (
	"ok/ast"
	"ok/number"
)

// Subtract will subtract two numbers.
type Subtract struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Subtract) Execute() error {
	ins.Result = ast.NewLiteralNumber(
		number.Subtract(
			number.NewNumber(ins.Left.Value),
			number.NewNumber(ins.Right.Value),
		).String(),
	)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Subtract) Answer() *ast.Literal {
	return ins.Result
}
