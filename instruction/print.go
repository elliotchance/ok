package instruction

import (
	"fmt"
	"io"
	"ok/ast"
)

// Print will output a string to stdout.
type Print struct {
	Values []*ast.Literal
	Stdout io.Writer
}

// Execute implements the Instruction interface for the VM.
func (ins *Print) Execute() error {
	for _, value := range ins.Values {
		fmt.Fprintln(ins.Stdout, value.Value)
	}

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Print) Answer() *ast.Literal {
	return nil
}
