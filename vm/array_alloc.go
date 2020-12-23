package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// ArrayAlloc allocates an array of fixed size.
type ArrayAlloc struct {
	Size, Result Register
	Kind         TypeRegister
}

// Execute implements the Instruction interface for the VM.
func (ins *ArrayAlloc) Execute(_ *int, vm *VM) error {
	size := number.Int64(number.NewNumber(vm.Get(ins.Size).Value))
	kind := vm.Types[ins.Kind]

	vm.Set(ins.Result, &ast.Literal{
		Kind:  kind,
		Array: make([]*ast.Literal, size),
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *ArrayAlloc) String() string {
	return fmt.Sprintf("%s = %s with %s elements", ins.Result, ins.Kind, ins.Size)
}
