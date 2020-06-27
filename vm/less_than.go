package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// LessThanNumber will compare two numbers.
type LessThanNumber struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(registers[ins.Left].Value),
		number.NewNumber(registers[ins.Right].Value),
	) < 0)

	return nil
}

// LessThanString will compare two strings.
type LessThanString struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanString) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralBool(
		registers[ins.Left].Value < registers[ins.Right].Value,
	)

	return nil
}
