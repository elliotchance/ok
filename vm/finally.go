package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// Finally will activate or deactivate a finally block.
type Finally struct {
	Index int
	Run   bool
}

// Execute implements the Instruction interface for the VM.
func (ins *Finally) Execute(registers map[Register]*ast.Literal, _ *int, vm *VM) error {
	vm.FinallyBlocks[len(vm.FinallyBlocks)-1][ins.Index].Run = ins.Run

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Finally) String() string {
	return fmt.Sprintf("finally %d", ins.Index)
}
