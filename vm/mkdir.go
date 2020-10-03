package vm

import (
	"fmt"
	"os"
)

// Mkdir creates a directory for Path and any directories needed along the way.
type Mkdir struct {
	Path Register // In
}

// Execute implements the Instruction interface for the VM.
func (ins *Mkdir) Execute(_ *int, vm *VM) error {
	path := vm.Get(ins.Path).Value
	err := os.MkdirAll(path, 0666)
	if err != nil {
		vm.Raise(err.Error())
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Mkdir) String() string {
	return fmt.Sprintf("os.CreateDirectory(%s)", ins.Path)
}
