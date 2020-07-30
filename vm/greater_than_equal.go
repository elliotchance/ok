package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// GreaterThanEqualNumber will compare two numbers.
type GreaterThanEqualNumber struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanEqualNumber) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(number.Cmp(
		number.NewNumber(vm.Get(ins.Left).Value),
		number.NewNumber(vm.Get(ins.Right).Value),
	) >= 0))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *GreaterThanEqualNumber) String() string {
	return fmt.Sprintf("%s = %s >= %s", ins.Result, ins.Left, ins.Right)
}

// GreaterThanEqualString will compare two strings.
type GreaterThanEqualString struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *GreaterThanEqualString) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(
		vm.Get(ins.Left).Value >= vm.Get(ins.Right).Value,
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *GreaterThanEqualString) String() string {
	return fmt.Sprintf("%s = %s >= %s", ins.Result, ins.Left, ins.Right)
}
