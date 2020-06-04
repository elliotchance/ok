package vm

import (
	"io"
	"ok/instruction"
)

// VM is an instance of a virtual machine to run ok instructions.
type VM struct {
	instructions []instruction.Instruction
	stdout       io.Writer
}

// NewVM will create a new VM ready to run the provided instructions.
func NewVM(instructions []instruction.Instruction, stdout io.Writer) *VM {
	return &VM{
		instructions: instructions,
		stdout:       stdout,
	}
}

// Run will run the program.
func (vm *VM) Run() {
	for _, i := range vm.instructions {
		i.Execute(vm.stdout)
	}
}
