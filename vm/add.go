package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Add will sum two numbers.
type Add struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Add) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralNumber(
		number.Add(
			number.NewNumber(registers[ins.Left].Value),
			number.NewNumber(registers[ins.Right].Value),
		).String(),
	)

	return nil
}
