package instruction

import (
	"ok/ast"
	"ok/number"
)

// LessThanNumber will compare two numbers.
type LessThanNumber struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanNumber) Execute() error {
	ins.Result = ast.NewLiteralBool(number.Cmp(
		number.NewNumber(ins.Left.Value),
		number.NewNumber(ins.Right.Value),
	) < 0)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *LessThanNumber) Answer() *ast.Literal {
	return ins.Result
}

// LessThanString will compare two strings.
type LessThanString struct {
	Left, Right, Result *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanString) Execute() error {
	ins.Result = ast.NewLiteralBool(ins.Left.Value < ins.Right.Value)

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *LessThanString) Answer() *ast.Literal {
	return ins.Result
}
