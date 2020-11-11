package vm

import (
	"fmt"
	"os"
)

// Remove removes (unlinks) a file.
type Remove struct {
	Path Register // In
}

// Execute implements the Instruction interface for the VM.
func (ins *Remove) Execute(_ *int, vm *VM) error {
	err := os.Remove(vm.Get(ins.Path).Value)
	if err != nil {
		vm.Raise(err.Error())
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Remove) String() string {
	return fmt.Sprintf("os.Remove(%s)", ins.Path)
}
