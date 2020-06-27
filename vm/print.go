package vm

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/elliotchance/ok/ast"
)

// Print will output a string to stdout.
type Print struct {
	Arguments []string
	Stdout    io.Writer
}

// Execute implements the Instruction interface for the VM.
func (ins *Print) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	for i, register := range ins.Arguments {
		if i > 0 {
			fmt.Fprint(ins.Stdout, " ")
		}

		r := registers[register]
		switch {
		case strings.HasPrefix(r.Kind, "[]"):
			fmt.Fprint(ins.Stdout, "[")
			for j, element := range r.Array {
				if j > 0 {
					fmt.Fprint(ins.Stdout, ", ")
				}

				printLiteral(ins.Stdout, element, true)
			}
			fmt.Fprint(ins.Stdout, "]")

		case strings.HasPrefix(r.Kind, "{}"):
			// Keys are always sorted.
			var keys []string
			for key := range r.Map {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			fmt.Fprint(ins.Stdout, "{")
			for j, key := range keys {
				element := r.Map[key]
				if j > 0 {
					fmt.Fprint(ins.Stdout, ", ")
				}

				fmt.Fprintf(ins.Stdout, `"%s": `, key)
				printLiteral(ins.Stdout, element, true)
			}
			fmt.Fprint(ins.Stdout, "}")

		default:
			printLiteral(ins.Stdout, r, false)
		}
	}

	fmt.Fprint(ins.Stdout, "\n")

	return nil
}

func printLiteral(out io.Writer, literal *ast.Literal, asJSON bool) {
	switch literal.Kind {
	case "char", "string", "data":
		if asJSON {
			// TODO(elliot): This is not escaped correctly.
			fmt.Fprintf(out, `"%s"`, literal.Value)

			return
		}
	}

	fmt.Fprint(out, literal.Value)
}
