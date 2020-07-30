package vm

import (
	"fmt"
)

// Assert is used in tests.
type Assert struct {
	Left, Right, Final Register
	Op                 string
	Pos                string
}

// Execute implements the Instruction interface for the VM.
func (ins *Assert) Execute(_ *int, vm *VM) error {
	pass := vm.Get(ins.Final).Value == "true"
	left := renderLiteral(vm.Get(ins.Left), true)
	right := renderLiteral(vm.Get(ins.Right), true)
	vm.assert(pass, left, ins.Op, right, ins.Pos)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Assert) String() string {
	return fmt.Sprintf("assert(%s %s %s)", ins.Left, ins.Op, ins.Right)
}
