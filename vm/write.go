package vm

import (
	"fmt"
)

// Write writes data to a file.
type Write struct {
	Data, Fd Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Write) Execute(_ *int, vm *VM) error {
	f := vm.Get(ins.Fd).File
	_, err := f.WriteString(vm.Get(ins.Data).Value)
	if err != nil {
		vm.Raise(err.Error())
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Write) String() string {
	return fmt.Sprintf("os.Write(%s, %s)", ins.Fd, ins.Data)
}
