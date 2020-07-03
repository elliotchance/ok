package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Power is Left to the power of Right.
type Power struct {
	Base, Power, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Power) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralNumber(
		number.Pow(
			number.NewNumber(registers[ins.Base].Value),
			number.NewNumber(registers[ins.Power].Value),
		).String(),
	)

	return nil
}
