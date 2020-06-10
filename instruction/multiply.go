package instruction

import (
	"ok/ast"
	"ok/number"
)

// Multiply will multiply two numbers.
type Multiply struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Multiply) Execute() error {
	ins.Result = ast.NewLiteralNumber(
		number.Multiply(
			number.NewNumber(ins.Left.Value),
			number.NewNumber(ins.Right.Value),
		).String(),
	)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Multiply) Answer() *ast.Literal {
	return ins.Result
}
