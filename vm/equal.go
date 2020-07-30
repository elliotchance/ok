package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler/kind"
	"github.com/elliotchance/ok/number"
)

// EqualNumber will compare two numbers for equality.
type EqualNumber struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *EqualNumber) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(numbersAreEqual(
		vm.Get(ins.Left),
		vm.Get(ins.Right),
	)))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *EqualNumber) String() string {
	return fmt.Sprintf("%s = %s == %s", ins.Result, ins.Left, ins.Right)
}

func numbersAreEqual(a, b *ast.Literal) bool {
	return number.Cmp(
		number.NewNumber(a.Value),
		number.NewNumber(b.Value),
	) == 0
}

// Equal will compare two non-numbers for equality. This works for every other
// type because every other type is stored as a string. When optimizations are
// made in the future this will need to be expanded to one instruction per type.
type Equal struct {
	Left, Right, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Equal) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(compareValue(
		vm.Get(ins.Left), vm.Get(ins.Right))))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Equal) String() string {
	return fmt.Sprintf("%s = %s == %s", ins.Result, ins.Left, ins.Right)
}

func compareValue(a, b *ast.Literal) bool {
	switch {
	case kind.IsArray(a.Kind):
		if len(a.Array) == len(b.Array) {
			for i, v := range a.Array {
				if !compareValue(v, b.Array[i]) {
					return false
				}
			}

			return true
		}

		return false

	case a.Kind == "number":
		return numbersAreEqual(a, b)

	default:
		return a.Value == b.Value
	}
}
