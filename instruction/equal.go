package instruction

import (
	"ok/ast"
	"ok/number"
)

// EqualNumber will compare two numbers for equality.
type EqualNumber struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *EqualNumber) Execute() error {
	ins.Result = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(ins.Left.Value),
		number.NewNumber(ins.Right.Value),
	) == 0)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *EqualNumber) Answer() *ast.Literal {
	return ins.Result
}

// Equal will compare two non-numbers for equality. This works for every other
// type because every other type is stored as a string. When optimizations are
// made in the future this will need to be expanded to one instruction per type.
type Equal struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Equal) Execute() error {
	ins.Result = ast.NewLiteralBool(ins.Left.Value == ins.Right.Value)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Equal) Answer() *ast.Literal {
	return ins.Result
}
