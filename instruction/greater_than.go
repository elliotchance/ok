package instruction

import (
	"ok/ast"
	"ok/number"
)

// GreaterThanNumber will compare two numbers.
type GreaterThanNumber struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanNumber) Execute() error {
	ins.Result = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(ins.Left.Value),
		number.NewNumber(ins.Right.Value),
	) > 0)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *GreaterThanNumber) Answer() *ast.Literal {
	return ins.Result
}

// GreaterThanString will compare two strings.
type GreaterThanString struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanString) Execute() error {
	ins.Result = ast.NewLiteralBool(ins.Left.Value > ins.Right.Value)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *GreaterThanString) Answer() *ast.Literal {
	return ins.Result
}
