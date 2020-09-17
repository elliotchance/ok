package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
)

// Type assigns the runtime type of a value to a string destination.
type Type struct {
	Value, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Type) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralString(vm.Get(ins.Value).Kind.String()))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Type) String() string {
	return fmt.Sprintf("%s = reflect.Type(%s)", ins.Result, ins.Value)
}
