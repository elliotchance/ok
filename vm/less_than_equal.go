package vm

import (
	"ok/ast"
	"ok/number"
)

// LessThanEqualNumber will compare two numbers.
type LessThanEqualNumber struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanEqualNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(registers[ins.Left].Value),
		number.NewNumber(registers[ins.Right].Value),
	) <= 0)

	return nil
}

// LessThanEqualString will compare two strings.
type LessThanEqualString struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanEqualString) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(
		registers[ins.Left].Value <= registers[ins.Right].Value,
	)

	return nil
}
