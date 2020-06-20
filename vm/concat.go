package vm

import (
	"ok/ast"
)

// Concat will create a new string by joining two other strings.
type Concat struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Concat) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralString(
		registers[ins.Left].Value + registers[ins.Right].Value)

	return nil
}
