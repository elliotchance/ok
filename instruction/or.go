package instruction

import (
	"ok/ast"
)

// Or is a logical OR between two bools.
type Or struct {
	Left, Right, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *Or) Execute(registers map[string]*ast.Literal) error {
	registers[ins.Result] = ast.NewLiteralBool(
		(registers[ins.Left].Value == "true") || (registers[ins.Right].Value == "true"),
	)

	return nil
}
