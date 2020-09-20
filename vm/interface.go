package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/util"
)

// Interface assigns the runtime interface of a value to a string destination.
type Interface struct {
	Value, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Interface) Execute(_ *int, vm *VM) error {
	i := util.Interface(vm.Get(ins.Value).Kind.Properties)
	vm.Set(ins.Result, asttest.NewLiteralString(i))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Interface) String() string {
	return fmt.Sprintf("%s = reflect.Interface(%s)", ins.Result, ins.Value)
}
