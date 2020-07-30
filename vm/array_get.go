package vm

import (
	"fmt"

	"github.com/elliotchance/ok/number"
)

// ArrayGet gets a value from the array by its index.
type ArrayGet struct {
	Array, Index, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *ArrayGet) Execute(_ *int, vm *VM) error {
	index := number.Int64(number.NewNumber(vm.Get(ins.Index).Value))
	vm.Set(ins.Result, vm.Get(ins.Array).Array[index])

	return nil
}

// String is the human-readable description of the instruction.
func (ins *ArrayGet) String() string {
	return fmt.Sprintf("%s = %s[%s]", ins.Result, ins.Array, ins.Index)
}
