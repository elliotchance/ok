package vm

import (
	"fmt"
	"os"

	"github.com/elliotchance/ok/number"
)

// Exit is the process exit with status code.
type Exit struct {
	Status Register // In
}

// Execute implements the Instruction interface for the VM.
func (ins *Exit) Execute(_ *int, vm *VM) error {
	os.Exit(number.Int(number.NewNumber(vm.Get(ins.Status).Value)))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Exit) String() string {
	return fmt.Sprintf("runtime.Exit(%s)", ins.Status)
}
