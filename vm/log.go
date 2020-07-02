package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// Log is a natural logarithm (base e).
type Log struct {
	X, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Log) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = ast.NewLiteralNumber(
		number.Log(
			number.NewNumber(registers[ins.X].Value),
		).String(),
	)

	return nil
}
