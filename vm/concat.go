package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
)

// Concat will create a new string by joining two other strings.
type Concat struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Concat) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralString(
		vm.Get(ins.Left).Value+vm.Get(ins.Right).Value))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Concat) String() string {
	return fmt.Sprintf("%s = concat(%s, %s)", ins.Result, ins.Left, ins.Right)
}
