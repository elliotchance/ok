package vm

import (
	"fmt"
)

// Jump will jump to the instruction.
type Jump struct {
	To int
}

// Execute implements the Instruction interface for the VM.
func (ins *Jump) Execute(i *int, vm *VM) error {
	*i = ins.To

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Jump) String() string {
	return fmt.Sprintf("jump to #%d", ins.To)
}
