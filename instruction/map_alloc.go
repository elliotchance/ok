package instruction

import (
	"ok/ast"
	"ok/number"
)

// MapAllocNumber allocates a "{}number" of fixed size.
type MapAllocNumber struct {
	Size, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *MapAllocNumber) Execute(registers map[string]*ast.Literal, _ *int) error {
	size := number.NewNumber(registers[ins.Size].Value).Int64()

	registers[ins.Result] = &ast.Literal{
		Kind: "{}number",
		Map:  make(map[string]*ast.Literal, size),
	}

	return nil
}
