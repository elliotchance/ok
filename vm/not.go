package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
)

// Not is a logical NOT of a bool.
type Not struct {
	Left, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Not) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(
		!(registers[ins.Left].Value == "true"),
	)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Not) String() string {
	return fmt.Sprintf("%s = not %s", ins.Result, ins.Left)
}
