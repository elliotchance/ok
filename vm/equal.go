package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler/kind"
	"github.com/elliotchance/ok/number"
)

// EqualNumber will compare two numbers for equality.
type EqualNumber struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *EqualNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(numbersAreEqual(
		registers[ins.Left],
		registers[ins.Right],
	))

	return nil
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
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Equal) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = asttest.NewLiteralBool(compareValue(
		registers[ins.Left], registers[ins.Right]))

	return nil
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
