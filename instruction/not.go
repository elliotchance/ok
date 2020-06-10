package instruction

import (
	"ok/ast"
)

// Not is a logical NOT of a bool.
type Not struct {
	Left, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Not) Execute() error {
	ins.Result = ast.NewLiteralBool(
		!(ins.Left.Value == "true"),
	)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Not) Answer() *ast.Literal {
	return ins.Result
}
