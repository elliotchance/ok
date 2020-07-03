package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// Combine will create a new data by joining two other datas.
type Combine struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Combine) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralData(
		[]byte(registers[ins.Left].Value + registers[ins.Right].Value))

	return nil
}
