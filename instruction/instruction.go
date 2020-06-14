package instruction

import "ok/ast"

// An Instruction can be executed by the VM.
type Instruction interface {
	Execute(registers map[string]*ast.Literal, currentInstruction *int) error
}
