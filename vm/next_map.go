package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// NextMap is used to tick a map iterator forward.
type NextMap struct {
	Map         Register // In (map): Containing the iterating map.
	Cursor      Register // In (number): Containing the current position.
	KeyResult   Register // Out (any): Load the key into this register.
	ValueResult Register // Out (any): Load the value into this register.
	Result      Register // Out (bool): Still more items?
}

// Execute implements the Instruction interface for the VM.
func (ins *NextMap) Execute(_ *int, vm *VM) error {
	array := vm.Get(ins.Map).Array
	pos := number.Int(number.NewNumber(vm.Get(ins.Cursor).Value))

	hasMore := pos < len(array)
	vm.Set(ins.Result, asttest.NewLiteralBool(hasMore))
	if hasMore {
		m := vm.Get(ins.Map).Map
		vm.Set(ins.KeyResult, array[pos])
		vm.Set(ins.ValueResult, m[array[pos].Value])
		vm.Set(ins.Cursor, asttest.NewLiteralNumber(fmt.Sprintf("%d", pos+1)))
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *NextMap) String() string {
	return fmt.Sprintf("%s, %s = next from %s; increment %s; has more %s",
		ins.ValueResult, ins.KeyResult, ins.Map, ins.Cursor, ins.Result)
}
