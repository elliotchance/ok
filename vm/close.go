package vm

import (
	"fmt"
	"os"
)

// Close closes a file handle.
type Close struct {
	Fd Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Close) Execute(_ *int, vm *VM) error {
	f := vm.Get(ins.Fd).File.(*os.File)
	err := f.Close()
	if err != nil {
		vm.Raise(err.Error())
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Close) String() string {
	return fmt.Sprintf("os.Close(%s)", ins.Fd)
}
