package instruction

import (
	"ok/ast"
	"ok/number"
)

// GreaterThanEqualNumber will compare two numbers.
type GreaterThanEqualNumber struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanEqualNumber) Execute(registers map[string]*ast.Literal, _ *int) error {
	registers[ins.Result] = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(registers[ins.Left].Value),
		number.NewNumber(registers[ins.Right].Value),
	) >= 0)

	return nil
}

// GreaterThanEqualString will compare two strings.
type GreaterThanEqualString struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanEqualString) Execute(registers map[string]*ast.Literal, _ *int) error {
	registers[ins.Result] = ast.NewLiteralBool(
		registers[ins.Left].Value >= registers[ins.Right].Value,
	)

	return nil
}
