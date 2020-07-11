package vm

import (
	"github.com/elliotchance/ok/ast"
)

// Raise put the VM into an error mode. The VM will look for an error handler.
type Raise struct {
	// Err is the register containing the error.
	Err string

	// Type is used to match the handler.
	Type string
}

// Execute implements the Instruction interface for the VM.
func (ins *Raise) Execute(registers map[string]*ast.Literal, _ *int, vm *VM) error {
	vm.Err = ins.Type

	return nil
}
