package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// ArrayAlloc allocates an array of fixed size.
type ArrayAlloc struct {
	Size, Result Register
	Kind         string
}

// Execute implements the Instruction interface for the VM.
func (ins *ArrayAlloc) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	size := number.Int64(number.NewNumber(registers[ins.Size].Value))

	registers[ins.Result] = &ast.Literal{
		Kind:  ins.Kind,
		Array: make([]*ast.Literal, size),
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *ArrayAlloc) String() string {
	return fmt.Sprintf("%s = %s with %s elements", ins.Result, ins.Kind, ins.Size)
}
