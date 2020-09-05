package vm

import (
	"fmt"
	"time"

	"github.com/elliotchance/ok/number"
)

// Sleep will sleep for a fractional amount of seconds.
type Sleep struct {
	Seconds Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Sleep) Execute(_ *int, vm *VM) error {
	seconds := number.NewNumber(vm.Get(ins.Seconds).Value)
	duration := number.Multiply(seconds, number.NewNumber("1000000000"))
	time.Sleep(time.Duration(number.Int64(duration)))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Sleep) String() string {
	return fmt.Sprintf("time.Sleep(%s)", ins.Seconds)
}
