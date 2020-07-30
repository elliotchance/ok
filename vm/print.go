package vm

import (
	"fmt"
)

// Print will output a string to stdout.
type Print struct {
	Arguments Registers
}

// Execute implements the Instruction interface for the VM.
func (ins *Print) Execute(_ *int, vm *VM) error {
	for i, register := range ins.Arguments {
		if i > 0 {
			fmt.Fprint(vm.Stdout, " ")
		}

		fmt.Fprint(vm.Stdout, renderLiteral(vm.Get(register), false))
	}

	fmt.Fprint(vm.Stdout, "\n")

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Print) String() string {
	return fmt.Sprintf("print%s", ins.Arguments)
}
