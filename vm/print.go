package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// Print will output a string to stdout.
type Print struct {
	Arguments Registers
}

// Execute implements the Instruction interface for the VM.
func (ins *Print) Execute(registers map[Register]*ast.Literal, _ *int, vm *VM) error {
	for i, register := range ins.Arguments {
		if i > 0 {
			fmt.Fprint(vm.Stdout, " ")
		}

		fmt.Fprint(vm.Stdout, renderLiteral(registers[register], false))
	}

	fmt.Fprint(vm.Stdout, "\n")

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Print) String() string {
	return fmt.Sprintf("print%s", ins.Arguments)
}
