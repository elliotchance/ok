package vm

import (
	"github.com/elliotchance/ok/ast"
)

// Finally will activate or deactivate a finally block.
type Finally struct {
	Index int
	Run   bool
}

// Execute implements the Instruction interface for the VM.
func (ins *Finally) Execute(registers map[string]*ast.Literal, _ *int, vm *VM) error {
	vm.FinallyBlocks[len(vm.FinallyBlocks)-1][ins.Index].Run = ins.Run

	return nil
}
