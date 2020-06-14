package instruction

import (
	"ok/ast"
	"ok/number"
)

// GreaterThanNumber will compare two numbers.
type GreaterThanNumber struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanNumber) Execute(registers map[string]*ast.Literal, _ *int) error {
	registers[ins.Result] = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(registers[ins.Left].Value),
		number.NewNumber(registers[ins.Right].Value),
	) > 0)

	return nil
}

// GreaterThanString will compare two strings.
type GreaterThanString struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanString) Execute(registers map[string]*ast.Literal, _ *int) error {
	registers[ins.Result] = ast.NewLiteralBool(
		registers[ins.Left].Value > registers[ins.Right].Value,
	)

	return nil
}
