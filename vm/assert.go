package vm

import (
	"github.com/elliotchance/ok/ast"
)

// Assert is used in tests.
type Assert struct {
	Left, Op, Right, Final string
	Pos                    string
}

// Execute implements the Instruction interface for the VM.
func (ins *Assert) Execute(registers map[string]*ast.Literal, _ *int, vm *VM) error {
	pass := registers[ins.Final].Value == "true"
	left := renderLiteral(registers[ins.Left], true)
	right := renderLiteral(registers[ins.Right], true)
	vm.assert(pass, left, ins.Op, right, ins.Pos)

	return nil
}
