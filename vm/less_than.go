package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// LessThanNumber will compare two numbers.
type LessThanNumber struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanNumber) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(number.Cmp(
		number.NewNumber(vm.Get(ins.Left).Value),
		number.NewNumber(vm.Get(ins.Right).Value),
	) < 0))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *LessThanNumber) String() string {
	return fmt.Sprintf("%s = %s < %s", ins.Result, ins.Left, ins.Right)
}

// LessThanString will compare two strings.
type LessThanString struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *LessThanString) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(
		vm.Get(ins.Left).Value < vm.Get(ins.Right).Value,
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *LessThanString) String() string {
	return fmt.Sprintf("%s = %s < %s", ins.Result, ins.Left, ins.Right)
}
