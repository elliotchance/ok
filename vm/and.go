package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// And is a logical AND between two bools.
type And struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *And) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(
		(registers[ins.Left].Value == "true") &&
			(registers[ins.Right].Value == "true"),
	)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *And) String() string {
	return fmt.Sprintf("%s = %s and %s", ins.Result, ins.Left, ins.Right)
}
