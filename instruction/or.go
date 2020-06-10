package instruction

import (
	"ok/ast"
)

// Or is a logical OR between two bools.
type Or struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Or) Execute() error {
	ins.Result = ast.NewLiteralBool(
		(ins.Left.Value == "true") || (ins.Right.Value == "true"),
	)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Or) Answer() *ast.Literal {
	return ins.Result
}
