package vm

import (
	"fmt"

	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/util"
)

// Get is used to access array indexes, map keys and object properties at
// runtime.
type Get struct {
	Object, Prop, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Get) Execute(_ *int, vm *VM) error {
	object := vm.Get(ins.Object)
	prop := vm.Get(ins.Prop)
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

		vm.Set(ins.Result, object.Array[index])

		return nil

	case types.KindMap:
		if prop.Kind.Kind != types.KindString {
			vm.Raise("prop must be a string for a map")

			return nil
		}

		if v, ok := object.Map[prop.Value]; ok {
			vm.Set(ins.Result, v)

			return nil
		}

		vm.Raise(fmt.Sprintf("no such key in map: %s", prop.Value))

		return nil

	case types.KindResolvedInterface, types.KindUnresolvedInterface:
		// TODO(elliot): This should not allow KindUnresolvedInterface.

		if prop.Kind.Kind != types.KindString {
			vm.Raise("prop must be a string for an object")

			return nil
		}

		if !util.IsPublic(prop.Value) {
			vm.Raise("cannot access private property: " + prop.Value)

			return nil
		}

		if v, ok := object.Map[prop.Value]; ok {
			vm.Set(ins.Result, v)

			return nil
		}

		vm.Raise(fmt.Sprintf("no such property in object: %s", prop.Value))

		return nil
	}

	vm.Raise("cannot use Get on " + object.Kind.String())

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Get) String() string {
	return fmt.Sprintf("%s = reflect.Get(%s, %s)", ins.Result, ins.Object, ins.Prop)
}
