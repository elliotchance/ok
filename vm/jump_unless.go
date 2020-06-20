package vm

import (
	"ok/ast"
)

// JumpUnless will jump to the instruction if the expression is false.
type JumpUnless struct {
	Condition string
	To        int
}

// Execute implements the Instruction interface for the VM.
func (ins *JumpUnless) Execute(registers map[string]*ast.Literal, i *int, _ *VM) error {
	if registers[ins.Condition].Value != "true" {
		*i = ins.To
	}

	return nil
}
