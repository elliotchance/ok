package instruction

// An Instruction can be executed by the VM.
type Instruction interface {
	Execute()
}
