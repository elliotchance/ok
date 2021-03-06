package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// NextArray is used to tick an array iterator forward.
type NextArray struct {
	Array       Register // In (array): Containing the iterating array.
	Cursor      Register // In (number): Containing the current position.
	KeyResult   Register // Out (any): Load the key into this register.
	ValueResult Register // Out (any): Load the value into this register.
	Result      Register // Out (bool): Still more items?
}

// Execute implements the Instruction interface for the VM.
func (ins *NextArray) Execute(_ *int, vm *VM) error {
	array := vm.Get(ins.Array).Array
	pos := number.Int(number.NewNumber(vm.Get(ins.Cursor).Value))
	hasMore := pos < len(array)
	vm.Set(ins.Result, asttest.NewLiteralBool(hasMore))
	if hasMore {
		// Key may be optionally set.
		if ins.KeyResult != "" {
			vm.Set(ins.KeyResult, asttest.NewLiteralNumber(fmt.Sprintf("%d", pos)))
		}

		vm.Set(ins.ValueResult, array[pos])
		vm.Set(ins.Cursor, asttest.NewLiteralNumber(fmt.Sprintf("%d", pos+1)))
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *NextArray) String() string {
	return fmt.Sprintf("%s, %s = next from %s; increment %s; has more %s",
		ins.ValueResult, ins.KeyResult, ins.Array, ins.Cursor, ins.Result)
}
