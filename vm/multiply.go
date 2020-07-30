package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Multiply will multiply two numbers.
type Multiply struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Multiply) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralNumber(
		number.Multiply(
			number.NewNumber(vm.Get(ins.Left).Value),
			number.NewNumber(vm.Get(ins.Right).Value),
		).String(),
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Multiply) String() string {
	return fmt.Sprintf("%s = %s * %s", ins.Result, ins.Left, ins.Right)
}
