package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// Assign sets a variable to the result of an expression.
type Assign struct {
	VariableName Register
	Value        *ast.Literal
	Register     Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Assign) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	if ins.Value != nil {
		registers[ins.VariableName] = ins.Value
	} else {
		registers[ins.VariableName] = registers[ins.Register]
	}

	return nil
}

func (ins *Assign) String() string {
	if ins.Value != nil {
		return fmt.Sprintf("%s = %s", ins.VariableName, ins.Value)
	}

	return fmt.Sprintf("%s = %s", ins.VariableName, ins.Register)
}
