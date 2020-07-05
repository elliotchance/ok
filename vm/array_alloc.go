package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// ArrayAlloc allocates a "[]number" of fixed size.
type ArrayAlloc struct {
	Size, Result string
	Kind         string
}

// Execute implements the Instruction interface for the VM.
func (ins *ArrayAlloc) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	size := number.Int64(number.NewNumber(registers[ins.Size].Value))

	registers[ins.Result] = &ast.Literal{
		Kind:  ins.Kind,
		Array: make([]*ast.Literal, size),
	}

	return nil
}
