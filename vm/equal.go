package vm

import (
	"ok/ast"
	"ok/number"
)

// EqualNumber will compare two numbers for equality.
type EqualNumber struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *EqualNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(registers[ins.Left].Value),
		number.NewNumber(registers[ins.Right].Value),
	) == 0)

	return nil
}

// Equal will compare two non-numbers for equality. This works for every other
// type because every other type is stored as a string. When optimizations are
// made in the future this will need to be expanded to one instruction per type.
type Equal struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Equal) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(
		registers[ins.Left].Value == registers[ins.Right].Value,
	)

	return nil
}
