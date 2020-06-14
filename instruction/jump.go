package instruction

import (
	"ok/ast"
)

// Jump will jump to the instruction.
type Jump struct {
	To int
}

// Execute implements the Instruction interface for the VM.
func (ins *Jump) Execute(registers map[string]*ast.Literal, i *int) error {
	*i = ins.To

	return nil
}
