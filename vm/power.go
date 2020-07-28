package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Power is Left to the power of Right.
type Power struct {
	Base, Power, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Power) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralNumber(
		number.Pow(
			number.NewNumber(registers[ins.Base].Value),
			number.NewNumber(registers[ins.Power].Value),
		).String(),
	)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Power) String() string {
	return fmt.Sprintf("%s = power(%s, %s)", ins.Result, ins.Base, ins.Power)
}
