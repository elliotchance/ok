package vm

import (
	"fmt"
	"time"

	"github.com/cockroachdb/apd/v2"
	"github.com/elliotchance/ok/number"
)

// FromUnix returns a time components from a unix timestamp.
type FromUnix struct {
	Seconds                                Register // in
	Year, Month, Day, Hour, Minute, Second Register // out
}

// Execute implements the Instruction interface for the VM.
func (ins *FromUnix) Execute(_ *int, vm *VM) error {
	t := vm.Get(ins.Seconds)

	rawSeconds := number.NewNumber(t.Value)
	seconds := apd.New(1, 0)
	nanoseconds := apd.New(1, 0)
	rawSeconds.Modf(seconds, nanoseconds)
	nanoseconds = number.Multiply(nanoseconds, apd.New(1, 9))

	tm := time.Unix(number.Int64(seconds), number.Int64(nanoseconds)).UTC()
	setTimeLiterals(vm, tm, ins.Year, ins.Month, ins.Day, ins.Hour, ins.Minute, ins.Second)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *FromUnix) String() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s = time.FromUnix(%s)",
		ins.Year, ins.Month, ins.Day, ins.Hour, ins.Minute, ins.Second,
		ins.Seconds)
}
