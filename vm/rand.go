package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast/asttest"
)

// Rand returns a random number between 0 and 1.
type Rand struct {
	Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Rand) Execute(_ *int, vm *VM) error {
	val := vm.rand.Float64()
	vm.Set(ins.Result, asttest.NewLiteralNumber(
		fmt.Sprintf("%f", val),
	))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Rand) String() string {
	return fmt.Sprintf("%s = math.Rand()", ins.Result)
}
