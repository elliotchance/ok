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
		get := vm.Get(ins.Register)
		vm.Set(ins.VariableName, &ast.Literal{
			Kind:  get.Kind,
			Value: get.Value,
			Array: get.Array,
			Map:   get.Map,
			Pos:   get.Pos,
		})
	}

	return nil
}

func (ins *Assign) String() string {
	if ins.Value != nil {
		return fmt.Sprintf("%s = %s (%s)", ins.VariableName, ins.Value,
			ins.Value.Kind.String())
	}

	return fmt.Sprintf("%s = %s", ins.VariableName, ins.Register)
}
