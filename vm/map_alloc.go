package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// MapAllocNumber allocates a "{}number" of fixed size.
type MapAllocNumber struct {
	Size, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *MapAllocNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	size := number.Int64(number.NewNumber(registers[ins.Size].Value))

	registers[ins.Result] = &ast.Literal{
		Kind: "{}number",

		// Array must also be allocated because it will contain the keys for the
		// map.
		Array: make([]*ast.Literal, 0, size),

		Map: make(map[string]*ast.Literal, size),
	}

	return nil
}
