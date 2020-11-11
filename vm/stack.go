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
	var elements []string

	// We must exclude the last stack item because that's the code that called
	// the internal __stack.
	for _, stack := range vm.Stack[:len(vm.Stack)-1] {
		elements = append(elements, stack[StackRegister].Value)
	}

	reverse(elements)

	vm.Set(ins.Stack, asttest.NewLiteralString(strings.Join(elements, "\n")))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Stack) String() string {
	return fmt.Sprintf("%s = runtime.Stack()", ins.Stack)
}
