package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Add will sum two numbers.
type Add struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Add) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralNumber(
		number.Add(
			number.NewNumber(vm.Get(ins.Left).Value),
			number.NewNumber(vm.Get(ins.Right).Value),
		).String(),
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Add) String() string {
	return fmt.Sprintf("%s = %s + %s", ins.Result, ins.Left, ins.Right)
}
