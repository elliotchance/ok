package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
)

// Is checks a type at runtime.
type Is struct {
	Value  Register // In any
	Type   Register // In string
	Result Register // Out bool
}

// Execute implements the Instruction interface for the VM.
func (ins *Is) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralBool(
		vm.Get(ins.Value).Kind.String() == vm.Get(ins.Type).Value,
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Is) String() string {
	return fmt.Sprintf("%s = %s is %s", ins.Result, ins.Value, ins.Type)
}
