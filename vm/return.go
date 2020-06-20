package vm

import (
	"ok/ast"
)

// Return tells the VM to jump out of this function.
type Return struct {
	Results []string
}

// Execute implements the Instruction interface for the VM.
func (ins *Return) Execute(registers map[string]*ast.Literal, _ *int, vm *VM) error {
	vm.Return = ins.Results

	return nil
}
