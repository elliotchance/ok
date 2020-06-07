package vm

import (
	"ok/instruction"
)

// VM is an instance of a virtual machine to run ok instructions.
type VM struct {
	instructions []instruction.Instruction
}

// NewVM will create a new VM ready to run the provided instructions.
func NewVM(instructions []instruction.Instruction) *VM {
	return &VM{
		instructions: instructions,
	}
}

// Run will run the program.
func (vm *VM) Run() {
	for _, i := range vm.instructions {
		i.Execute()
	}
}
