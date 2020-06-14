package instruction

import (
	"ok/ast"
)

// Combine will create a new data by joining two other datas.
type Combine struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Combine) Execute(registers map[string]*ast.Literal, _ *int) error {
	registers[ins.Result] = ast.NewLiteralData(
		[]byte(registers[ins.Left].Value + registers[ins.Right].Value))

	return nil
}
