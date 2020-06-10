package instruction

import (
	"ok/ast"
	"ok/number"
)

// GreaterThanEqualNumber will compare two numbers.
type GreaterThanEqualNumber struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanEqualNumber) Execute() error {
	ins.Result = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(ins.Left.Value),
		number.NewNumber(ins.Right.Value),
	) >= 0)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *GreaterThanEqualNumber) Answer() *ast.Literal {
	return ins.Result
}

// GreaterThanEqualString will compare two strings.
type GreaterThanEqualString struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanEqualString) Execute() error {
	ins.Result = ast.NewLiteralBool(ins.Left.Value >= ins.Right.Value)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *GreaterThanEqualString) Answer() *ast.Literal {
	return ins.Result
}
