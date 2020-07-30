package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Remainder will return the remainder when dividing two numbers. This is not
// the same as a modulo. A remainder may be negative.
type Remainder struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Remainder) Execute(_ *int, vm *VM) error {
	divide, err := number.Remainder(
		number.NewNumber(vm.Get(ins.Left).Value),
		number.NewNumber(vm.Get(ins.Right).Value),
	)
	if err != nil {
		// TODO(elliot): This needs to be the same precision of zero.
		vm.Set(ins.Result, asttest.NewLiteralNumber("0"))
		return err
	}

	vm.Set(ins.Result, asttest.NewLiteralNumber(divide.String()))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Remainder) String() string {
	return fmt.Sprintf("%s = %s %% %s", ins.Result, ins.Left, ins.Right)
}
