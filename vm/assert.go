package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// Assert is used in tests.
type Assert struct {
	Left, Right, Final Register
	Op                 string
	Pos                string
}

// Execute implements the Instruction interface for the VM.
func (ins *Assert) Execute(registers map[Register]*ast.Literal, _ *int, vm *VM) error {
	pass := registers[ins.Final].Value == "true"
	left := renderLiteral(registers[ins.Left], true)
	right := renderLiteral(registers[ins.Right], true)
	vm.assert(pass, left, ins.Op, right, ins.Pos)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Assert) String() string {
	return fmt.Sprintf("assert(%s %s %s)", ins.Left, ins.Op, ins.Right)
}
