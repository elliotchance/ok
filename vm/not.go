package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
)

// Not is a logical NOT of a bool.
type Not struct {
	Left, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Not) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(
		!(vm.Get(ins.Left).Value == "true"),
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Not) String() string {
	return fmt.Sprintf("%s = not %s", ins.Result, ins.Left)
}
