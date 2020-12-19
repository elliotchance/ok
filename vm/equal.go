package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/types"
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
	// This is also used when comparing any values so we have to test for type
	// here to be sure. Ideally, all of the types will have custom equality -
	// like EqualNumber - or, at least notify if an any is being used to avoid
	// this runtime check.
	if a.Kind.Kind != b.Kind.Kind {
		return false
	}

	switch a.Kind.Kind {
	case types.KindArray:
		if len(a.Array) == len(b.Array) {
			for i, v := range a.Array {
				if !compareValue(v, b.Array[i]) {
					return false
				}
			}

			return true
		}

		return false

	case types.KindMap:
		if len(a.Map) != len(b.Map) {
			return false
		}

		keys := make(map[string]struct{})
		for key := range a.Map {
			keys[key] = struct{}{}
		}
		for key := range b.Map {
			keys[key] = struct{}{}
		}

		if len(keys) != len(a.Map) {
			return false
		}

		for k, v := range a.Map {
			if !compareValue(v, b.Map[k]) {
				return false
			}
		}

		return true

	case types.KindNumber:
		return numbersAreEqual(a, b)

	default:
		return a.Value == b.Value
	}
}
