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
	for i, value := range ins.Values {
		if i > 0 {
			fmt.Fprint(ins.Stdout, " ")
		}

		fmt.Fprint(ins.Stdout, value.Value)
	}

	fmt.Fprint(ins.Stdout, "\n")

	return nil
}

// Answer will be removed shortly. Right now it's used to evaluate literals
// because there are no variables.
func (ins *Print) Answer() *ast.Literal {
	return nil
}
