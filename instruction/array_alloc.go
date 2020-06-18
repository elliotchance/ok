package instruction

import (
	"ok/ast"
	"ok/number"
)

// ArrayAllocNumber allocates a "[]number" of fixed size.
type ArrayAllocNumber struct {
	Size, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *ArrayAllocNumber) Execute(registers map[string]*ast.Literal, _ *int) error {
	size := number.NewNumber(registers[ins.Size].Value).Int64()

	registers[ins.Result] = &ast.Literal{
		Kind:  "[]number",
		Array: make([]*ast.Literal, size),
	}

	return nil
}
