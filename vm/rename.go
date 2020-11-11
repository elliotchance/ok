package vm

import (
	"fmt"
	"os"
)

// Rename renames (moves) a file.
type Rename struct {
	OldPath Register // In
	NewPath Register // In
}

// Execute implements the Instruction interface for the VM.
func (ins *Rename) Execute(_ *int, vm *VM) error {
	err := os.Rename(vm.Get(ins.OldPath).Value, vm.Get(ins.NewPath).Value)
	if err != nil {
		vm.Raise(err.Error())
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Rename) String() string {
	return fmt.Sprintf("os.Rename(%s, %s)", ins.OldPath, ins.NewPath)
}
