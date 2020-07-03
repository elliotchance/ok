package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// Concat will create a new string by joining two other strings.
type Concat struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Concat) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralString(
		registers[ins.Left].Value + registers[ins.Right].Value)

	return nil
}
