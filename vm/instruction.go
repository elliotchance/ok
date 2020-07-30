package vm

import (
	"fmt"
)

// An Instruction can be executed by the VM.
type Instruction interface {
	// Stringer provides human-readable descriptions of instructions. It's
	// helpful for debugging and used directly by "ok asm".
	fmt.Stringer

	Execute(i *int, vm *VM) error
}
