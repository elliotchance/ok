package vm

import (
	"fmt"
	"sort"
	"strings"

	"github.com/elliotchance/ok/ast"
)

// Print will output a string to stdout.
type Print struct {
	Arguments []string
}

// Execute implements the Instruction interface for the VM.
func (ins *Print) Execute(registers map[string]*ast.Literal, _ *int, vm *VM) error {
	for i, register := range ins.Arguments {
		if i > 0 {
			fmt.Fprint(vm.Stdout, " ")
		}

		r := registers[register]
		switch {
		case strings.HasPrefix(r.Kind, "[]"):
			fmt.Fprint(vm.Stdout, "[")
			for j, element := range r.Array {
				if j > 0 {
					fmt.Fprint(vm.Stdout, ", ")
				}

				fmt.Fprint(vm.Stdout, renderLiteral(element, true))
			}
			fmt.Fprint(vm.Stdout, "]")

		case strings.HasPrefix(r.Kind, "{}"):
			// Keys are always sorted.
			var keys []string
			for key := range r.Map {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			fmt.Fprint(vm.Stdout, "{")
			for j, key := range keys {
				element := r.Map[key]
				if j > 0 {
					fmt.Fprint(vm.Stdout, ", ")
				}

				fmt.Fprintf(vm.Stdout, `"%s": `, key)
				fmt.Fprint(vm.Stdout, renderLiteral(element, true))
			}
			fmt.Fprint(vm.Stdout, "}")

		default:
			fmt.Fprint(vm.Stdout, renderLiteral(r, false))
		}
	}

	fmt.Fprint(vm.Stdout, "\n")

	return nil
}
