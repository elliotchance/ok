package instruction

import (
	"ok/ast"
)

// And is a logical AND between two bools.
type And struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *And) Execute() error {
	ins.Result = ast.NewLiteralBool(
		(ins.Left.Value == "true") && (ins.Right.Value == "true"),
	)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *And) Answer() *ast.Literal {
	return ins.Result
}
