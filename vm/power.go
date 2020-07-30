package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Power is Left to the power of Right.
type Power struct {
	Base, Power, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Power) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralNumber(
		number.Pow(
			number.NewNumber(vm.Get(ins.Base).Value),
			number.NewNumber(vm.Get(ins.Power).Value),
		).String(),
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Power) String() string {
	return fmt.Sprintf("%s = power(%s, %s)", ins.Result, ins.Base, ins.Power)
}
