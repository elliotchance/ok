package instruction

import (
	"ok/ast"
)

// Combine will create a new data by joining two other datas.
type Combine struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Combine) Execute() error {
	ins.Result = ast.NewLiteralData([]byte(ins.Left.Value + ins.Right.Value))

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Combine) Answer() *ast.Literal {
	return ins.Result
}
