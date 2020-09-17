package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/util"
)

// Set is used to set by array indexes, map keys and object properties at
// runtime.
type Set struct {
	Object, Prop, Value Register

	// Result will always be true. This is just a way to make the instruction
	// set easier, but it doesn't need to be here. Or at least return something
	// more useful.
	Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Set) Execute(_ *int, vm *VM) error {
	object := vm.Get(ins.Object)
	prop := vm.Get(ins.Prop)
	value := vm.Get(ins.Value)

	switch object.Kind.Kind {
	case types.KindArray:
		if prop.Kind.Kind != types.KindNumber {
			vm.Raise("prop must be a number for an array")

			return nil
		}

		index := number.Int(number.NewNumber(prop.Value))
		if index < 0 || index >= len(object.Array) {
			vm.Raise(fmt.Sprintf("index out of bounds [%d] with length %d",
				index, len(object.Array)))

			return nil
		}

		object.Array[index] = value
		vm.Set(ins.Result, asttest.NewLiteralBool(true))

		return nil

	case types.KindMap:
		if prop.Kind.Kind != types.KindString {
			vm.Raise("prop must be a string for a map")

			return nil
		}

		object.Map[prop.Value] = value
		vm.Set(ins.Result, asttest.NewLiteralBool(true))

		return nil

	case types.KindResolvedInterface, types.KindUnresolvedInterface:
		// TODO(elliot): This should not allow KindUnresolvedInterface.

		if prop.Kind.Kind != types.KindString {
			vm.Raise("prop must be a string for an object")

			return nil
		}

		if !util.IsPublic(prop.Value) {
			vm.Raise("cannot mutate private property: " + prop.Value)

			return nil
		}

		object.Map[prop.Value] = value
		vm.Set(ins.Result, asttest.NewLiteralBool(true))

		return nil
	}

	vm.Raise("cannot use Set on " + object.Kind.String())

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Set) String() string {
	return fmt.Sprintf("reflect.Set(%s, %s, %s)", ins.Object, ins.Prop, ins.Value)
}
