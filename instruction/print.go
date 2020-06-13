package instruction

import (
	"fmt"
	"io"
	"ok/ast"
)

// Print will output a string to stdout.
type Print struct {
	Arguments []string
	Stdout    io.Writer
}

// Execute implements the Instruction interface for the VM.
func (ins *Print) Execute(registers map[string]*ast.Literal) error {
	for i, register := range ins.Arguments {
		if i > 0 {
			fmt.Fprint(ins.Stdout, " ")
		}

		fmt.Fprint(ins.Stdout, registers[register].Value)
	}

	fmt.Fprint(ins.Stdout, "\n")

	return nil
}
