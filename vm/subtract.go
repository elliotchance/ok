package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Subtract will subtract two numbers.
type Subtract struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Subtract) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralNumber(
		number.Subtract(
			number.NewNumber(vm.Get(ins.Left).Value),
			number.NewNumber(vm.Get(ins.Right).Value),
		).String(),
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Subtract) String() string {
	return fmt.Sprintf("%s = %s - %s", ins.Result, ins.Left, ins.Right)
}
