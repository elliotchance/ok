package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// Log is a natural logarithm (base e).
type Log struct {
	X, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Log) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, asttest.NewLiteralNumber(
		number.Log(
			number.NewNumber(vm.Get(ins.X).Value),
		).String(),
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Log) String() string {
	return fmt.Sprintf("%s = log(%s)", ins.Result, ins.X)
}
