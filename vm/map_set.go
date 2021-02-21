package vm

import (
	"fmt"
)

// MapSet sets a number value to an index.
type MapSet struct {
	Map, Key, Value Register
}

// Execute implements the Instruction interface for the VM.
func (ins *MapSet) Execute(_ *int, vm *VM) error {
	key := vm.Get(ins.Key).Value
	vm.Get(ins.Map).Map[key] = vm.Get(ins.Value).Copy()
	vm.Get(ins.Map).Array = append(vm.Get(ins.Map).Array, vm.Get(ins.Key).Copy())

	return nil
}

// String is the human-readable description of the instruction.
func (ins *MapSet) String() string {
	return fmt.Sprintf("%s[%s] = %s", ins.Map, ins.Key, ins.Value)
}
