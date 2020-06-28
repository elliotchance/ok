package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// ArrayAllocNumber allocates a "[]number" of fixed size.
type ArrayAllocNumber struct {
	Size, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *ArrayAllocNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	size := number.Int64(number.NewNumber(registers[ins.Size].Value))

	registers[ins.Result] = &ast.Literal{
		Kind:  "[]number",
		Array: make([]*ast.Literal, size),
	}

	return nil
}
