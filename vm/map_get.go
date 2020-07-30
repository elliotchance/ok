package vm

import (
	"fmt"
)

// MapGet gets a value from the map by its key.
type MapGet struct {
	Map, Key, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *MapGet) Execute(_ *int, vm *VM) error {
	key := vm.Get(ins.Key).Value
	vm.Set(ins.Result, vm.Get(ins.Map).Map[key])

	return nil
}

// String is the human-readable description of the instruction.
func (ins *MapGet) String() string {
	return fmt.Sprintf("%s = %s[%s]", ins.Result, ins.Map, ins.Key)
}
