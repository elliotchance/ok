package instruction

import "ok/ast"

// An Instruction can be executed by the VM.
type Instruction interface {
	Execute() error

	// Answer will be removed shortly. Right now it's used to evaluate literals
	// because there are no variables.
	Answer() *ast.Literal
}
