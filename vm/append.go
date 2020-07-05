package vm

import (
	"github.com/elliotchance/ok/ast"
)

// Append returns an array by combining two other arrays.
type Append struct {
	A, B, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Append) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = &ast.Literal{
		Kind:  registers[ins.A].Kind,
		Array: append(registers[ins.A].Array, registers[ins.B].Array...),
	}

	return nil
}
