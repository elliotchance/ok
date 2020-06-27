package vm

import (
	"github.com/elliotchance/ok/ast"
)

// And is a logical AND between two bools.
type And struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *And) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(
		(registers[ins.Left].Value == "true") &&
			(registers[ins.Right].Value == "true"),
	)

	return nil
}
