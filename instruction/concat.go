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
