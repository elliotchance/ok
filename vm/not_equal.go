package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// NotEqualNumber will compare two numbers for equality.
type NotEqualNumber struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *NotEqualNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(!numbersAreEqual(
		registers[ins.Left],
		registers[ins.Right],
	))

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
	registers[ins.Result] = asttest.NewLiteralBool(!compareValue(
		registers[ins.Left], registers[ins.Right],
	))

	return nil
}
