package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// Or is a logical OR between two bools.
type Or struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Or) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(
		(registers[ins.Left].Value == "true") ||
			(registers[ins.Right].Value == "true"),
	)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Or) String() string {
	return fmt.Sprintf("%s = %s or %s", ins.Result, ins.Left, ins.Right)
}
