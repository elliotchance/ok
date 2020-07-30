package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Divide will multiply two numbers.
type Divide struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Divide) Execute(_ *int, vm *VM) error {
	divide, err := number.Divide(
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
func (ins *Divide) String() string {
	return fmt.Sprintf("%s = %s / %s", ins.Result, ins.Left, ins.Right)
}
