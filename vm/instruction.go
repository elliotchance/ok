package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// An Instruction can be executed by the VM.
type Instruction interface {
	// Stringer provides human-readable descriptions of instructions. It's
	// helpful for debugging and used directly by "ok asm".
	fmt.Stringer

	Execute(registers map[Register]*ast.Literal, currentInstruction *int, vm *VM) error
}
