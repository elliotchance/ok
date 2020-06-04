package instruction

import "fmt"

// Print will output a string to stdout.
type Print struct {
	Value string
}

// Execute implements the Instruction interface for the VM.
func (ins *Print) Execute() {
	fmt.Println(ins.Value)
}
