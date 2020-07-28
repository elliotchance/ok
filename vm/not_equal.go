package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// NotEqualNumber will compare two numbers for equality.
type NotEqualNumber struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *NotEqualNumber) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(!numbersAreEqual(
		registers[ins.Left],
		registers[ins.Right],
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *NotEqualNumber) String() string {
	return fmt.Sprintf("%s = %s != %s", ins.Result, ins.Left, ins.Right)
}

// NotEqual will compare two non-numbers for non-equality. This works for every
// other type because every other type is stored as a string. When optimizations
// are made in the future this will need to be expanded to one instruction per
// type.
type NotEqual struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *NotEqual) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(!compareValue(
		registers[ins.Left], registers[ins.Right],
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *NotEqual) String() string {
	return fmt.Sprintf("%s = %s != %s", ins.Result, ins.Left, ins.Right)
}
