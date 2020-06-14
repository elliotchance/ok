package instruction

import (
	"ok/ast"
)

// Assign sets a variable to the result of an expression.
type Assign struct {
	VariableName string
	Value        *ast.Literal
	Register     string
}

// Execute implements the Instruction interface for the VM.
func (ins *Assign) Execute(registers map[string]*ast.Literal, _ *int) error {
	if ins.Value != nil {
		registers[ins.VariableName] = ins.Value
	} else {
		registers[ins.VariableName] = registers[ins.Register]
	}

	return nil
}
