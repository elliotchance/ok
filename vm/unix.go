package vm

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cockroachdb/apd/v2"
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/types"
)

// Unix returns a unix timestamp from a time.Time.
type Unix struct {
	Time, Result Register
}

func i(t *ast.Literal, prop string) int {
	n, _ := strconv.Atoi(t.Map[prop].Value)
	return n
}

// Execute implements the Instruction interface for the VM.
func (ins *Unix) Execute(_ *int, vm *VM) error {
	t := vm.Get(ins.Time)
	rawSeconds := number.NewNumber(t.Map["Second"].Value)
	seconds := apd.New(1, 0)
	nanoseconds := apd.New(1, 0)
	rawSeconds.Modf(seconds, nanoseconds)

	nanoseconds = number.Multiply(nanoseconds, apd.New(1, 9))

	tm := time.Date(
		i(t, "Year"), time.Month(i(t, "Month")), i(t, "Day"),
		i(t, "Hour"), i(t, "Minute"), number.Int(seconds),
		number.Int(nanoseconds),
		time.UTC,
	)

	// An error is impossible because the denominator can never be 0.
	d, _ := number.Divide(
		number.NewNumber(fmt.Sprintf("%d", tm.UnixNano())),
		number.NewNumber("1000000000"),
	)

	vm.Set(ins.Result, &ast.Literal{
		Kind:  types.Number,
		Value: number.Format(d, -1),
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Unix) String() string {
	return fmt.Sprintf("%s = time.Unix(%s)", ins.Result, ins.Time)
}
