package vm

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast/asttest"
)

// Stack returns the current stack trace.
type Stack struct {
	Stack Register // Out
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

// Execute implements the Instruction interface for the VM.
func (ins *Stack) Execute(_ *int, vm *VM) error {
	elements := vm.captureCallStack("")

	// Exclude some elements to avoid including the internal __stack call.
	elements = elements[2 : len(elements)-1]

	vm.Set(ins.Stack, asttest.NewLiteralString(strings.Join(elements, "\n")))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Stack) String() string {
	return fmt.Sprintf("%s = runtime.Stack()", ins.Stack)
}
