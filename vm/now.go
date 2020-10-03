package vm

import (
	"fmt"
	"strconv"
	"time"

	"github.com/elliotchance/ok/ast/asttest"
)

// Now sets the Result to the current time as a time.Time.
type Now struct {
	Year, Month, Day, Hour, Minute, Second Register // out
}

// Execute implements the Instruction interface for the VM.
func (ins *Now) Execute(_ *int, vm *VM) error {
	setTimeLiterals(vm, time.Now(),
		ins.Year, ins.Month, ins.Day, ins.Hour, ins.Minute, ins.Second)

	return nil
}

func setTimeLiterals(vm *VM, t time.Time, year, month, day, hour, minute, second Register) {
	vm.Set(year, asttest.NewLiteralNumber(strconv.Itoa(t.Year())))
	vm.Set(month, asttest.NewLiteralNumber(strconv.Itoa(int(t.Month()))))
	vm.Set(day, asttest.NewLiteralNumber(strconv.Itoa(t.Day())))
	vm.Set(hour, asttest.NewLiteralNumber(strconv.Itoa(t.Hour())))
	vm.Set(minute, asttest.NewLiteralNumber(strconv.Itoa(t.Minute())))
	vm.Set(second, asttest.NewLiteralNumber(fmt.Sprintf("%.9f",
		float64(t.Second())+float64(t.Nanosecond())/float64(time.Second))))
}

// String is the human-readable description of the instruction.
func (ins *Now) String() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s = time.Now()",
		ins.Year, ins.Month, ins.Day, ins.Hour, ins.Minute, ins.Second)
}
