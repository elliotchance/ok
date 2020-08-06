package vm

import (
	"fmt"
)

// ParentScope sets the parent scope of a function literal.
type ParentScope struct {
	X Register
}

// Execute implements the Instruction interface for the VM.
func (ins *ParentScope) Execute(_ *int, vm *VM) error {
	vm.Stack[len(vm.Stack)-1][ins.X].Map = vm.Stack[len(vm.Stack)-1][StateRegister].Map

	return nil
}

func (ins *ParentScope) String() string {
	return fmt.Sprintf("%s", ins.X)
}
