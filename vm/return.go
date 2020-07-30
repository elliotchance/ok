package vm

import (
	"fmt"
)

// Return tells the VM to jump out of this function.
type Return struct {
	Results Registers
}

// Execute implements the Instruction interface for the VM.
func (ins *Return) Execute(_ *int, vm *VM) error {
	vm.Return = ins.Results

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Return) String() string {
	return fmt.Sprintf("return %s", ins.Results)
}
