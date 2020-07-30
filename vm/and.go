package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
)

// And is a logical AND between two bools.
type And struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *And) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(
		(vm.Get(ins.Left).Value == "true") &&
			(vm.Get(ins.Right).Value == "true"),
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *And) String() string {
	return fmt.Sprintf("%s = %s and %s", ins.Result, ins.Left, ins.Right)
}
