package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// Assign sets a variable to the result of an expression.
type Assign struct {
	// VariableName is the destination.
	VariableName Register

	// Value or Register must be supplied, but not both.
	Value    *ast.Literal
	Register Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Assign) Execute(_ *int, vm *VM) error {
	if ins.Value != nil {
		vm.Set(ins.VariableName, ins.Value)
	} else {
		vm.Set(ins.VariableName, vm.Get(ins.Register))
	}

	return nil
}

func (ins *Assign) String() string {
	if ins.Value != nil {
		return fmt.Sprintf("%s = %s", ins.VariableName, ins.Value)
	}

	return fmt.Sprintf("%s = %s", ins.VariableName, ins.Register)
}
