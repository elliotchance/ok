package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// Combine will create a new data by joining two other datas.
type Combine struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Combine) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralData(
		[]byte(registers[ins.Left].Value + registers[ins.Right].Value))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Combine) String() string {
	return fmt.Sprintf("%s = combine(%s, %s)", ins.Result, ins.Left, ins.Right)
}
