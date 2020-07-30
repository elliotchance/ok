package vm

import (
	"fmt"
)

// JumpUnless will jump to the instruction if the expression is false.
type JumpUnless struct {
	Condition Register
	To        int
}

// Execute implements the Instruction interface for the VM.
func (ins *JumpUnless) Execute(i *int, vm *VM) error {
	if vm.Get(ins.Condition).Value != "true" {
		*i = ins.To
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *JumpUnless) String() string {
	return fmt.Sprintf("if %s is false then jump to #%d", ins.Condition, ins.To)
}
