package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// Not is a logical NOT of a bool.
type Not struct {
	Left, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Not) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(
		!(registers[ins.Left].Value == "true"),
	)

	return nil
}
