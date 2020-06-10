package instruction

import (
	"ok/ast"
	"ok/number"
)

// LessThanEqualNumber will compare two numbers.
type LessThanEqualNumber struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanEqualNumber) Execute() error {
	ins.Result = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(ins.Left.Value),
		number.NewNumber(ins.Right.Value),
	) <= 0)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *LessThanEqualNumber) Answer() *ast.Literal {
	return ins.Result
}

// LessThanEqualString will compare two strings.
type LessThanEqualString struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanEqualString) Execute() error {
	ins.Result = ast.NewLiteralBool(ins.Left.Value <= ins.Right.Value)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *LessThanEqualString) Answer() *ast.Literal {
	return ins.Result
}
