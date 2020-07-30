package vm

import (
	"fmt"

	"github.com/elliotchance/ok/number"
)

// ArraySet sets a number value to an index.
type ArraySet struct {
	Array, Index, Value Register
}

// Execute implements the Instruction interface for the VM.
func (ins *ArraySet) Execute(_ *int, vm *VM) error {
	index := number.Int64(number.NewNumber(vm.Get(ins.Index).Value))
	vm.Get(ins.Array).Array[index] = vm.Get(ins.Value)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *ArraySet) String() string {
	return fmt.Sprintf("%s[%s] = %s", ins.Array, ins.Index, ins.Value)
}
