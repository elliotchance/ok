package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// Append returns an array by combining two other arrays.
type Append struct {
	A, B, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Append) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, &ast.Literal{
		Kind:  vm.Get(ins.A).Kind,
		Array: append(vm.Get(ins.A).Array, vm.Get(ins.B).Array...),
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Append) String() string {
	return fmt.Sprintf("%s = append(%s, %s)", ins.Result, ins.A, ins.B)
}
