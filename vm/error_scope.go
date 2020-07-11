package vm

import (
	"github.com/elliotchance/ok/ast"
)

// On is a pragma for the vm to handle errors. It can also be used to indicate
// the end of the on's so that the VM can send the error up to the caller.
type On struct {
	// Type will be the name of the error, or it can be empty to signal there
	// are no more errors to check. If the VM hits an empty Type it will return
	// and pass the error up to the caller.
	Type string
}

// Execute implements the Instruction interface for the VM.
func (ins *On) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	// Nothing happens here because On is just a pragma for the VM to look
	// forward to find the error handler.
	return nil
}
