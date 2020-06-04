package instruction

import (
	"fmt"
	"io"
	"ok/ast"
)

// Print will output a string to stdout.
type Print struct {
	Value *ast.Literal
}

// Execute implements the Instruction interface for the VM.
func (ins *Print) Execute(stdout io.Writer) {
	if ins.Value != nil {
		fmt.Fprintln(stdout, ins.Value.Value)
	}
}
