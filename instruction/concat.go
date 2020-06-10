package instruction

import (
	"ok/ast"
)

// Concat will create a new string by joining two other strings.
type Concat struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Concat) Execute() error {
	ins.Result = ast.NewLiteralString(ins.Left.Value + ins.Right.Value)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Concat) Answer() *ast.Literal {
	return ins.Result
}
