package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/types"
)

// MapAlloc allocates a map of fixed size.
type MapAlloc struct {
	Kind         *types.Type
	Size, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *MapAlloc) Execute(_ *int, vm *VM) error {
	size := number.Int64(number.NewNumber(vm.Get(ins.Size).Value))

	vm.Set(ins.Result, &ast.Literal{
		Kind: ins.Kind,

		// Array must also be allocated because it will contain the keys for the
		// map.
		Array: make([]*ast.Literal, 0, size),

		Map: make(map[string]*ast.Literal, size),
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *MapAlloc) String() string {
	return fmt.Sprintf("%s = {}any with %s elements", ins.Result, ins.Size)
}
