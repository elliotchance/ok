package vm

import (
	"ok/ast"
	"ok/number"
)

// NotEqualNumber will compare two numbers for equality.
type NotEqualNumber struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *NotEqualNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(registers[ins.Left].Value),
		number.NewNumber(registers[ins.Right].Value),
	) != 0)

	return nil
}

// NotEqual will compare two non-numbers for non-equality. This works for every
// other type because every other type is stored as a string. When optimizations
// are made in the future this will need to be expanded to one instruction per
// type.
type NotEqual struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *NotEqual) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(
		registers[ins.Left].Value != registers[ins.Right].Value,
	)

	return nil
}
