package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// Or is a logical OR between two bools.
type Or struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Or) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(
		(registers[ins.Left].Value == "true") ||
			(registers[ins.Right].Value == "true"),
	)

	return nil
}
