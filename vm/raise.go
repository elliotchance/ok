package vm

import (
	"fmt"
)

// Raise put the VM into an error mode. The VM will look for an error handler.
type Raise struct {
	// Err is the register containing the error.
	Err Register

	// Type is used to match the handler.
	Type TypeRegister
}

// Execute implements the Instruction interface for the VM.
func (ins *Raise) Execute(_ *int, vm *VM) error {
	vm.ErrType = vm.Types[ins.Type]
	vm.ErrValue = vm.Get(ins.Err)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Raise) String() string {
	return fmt.Sprintf("raise %s (%s)", ins.Err, ins.Type)
}
