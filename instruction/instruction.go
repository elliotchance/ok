package instruction

import "io"

// An Instruction can be executed by the VM.
type Instruction interface {
	Execute(stdout io.Writer)
}
