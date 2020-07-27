package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// ArrayGet gets a value from the array by its index.
type ArrayGet struct {
	Array, Index, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *ArrayGet) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	index := number.Int64(number.NewNumber(registers[ins.Index].Value))
	registers[ins.Result] = registers[ins.Array].Array[index]

	return nil
}

// String is the human-readable description of the instruction.
func (ins *ArrayGet) String() string {
	return fmt.Sprintf("%s = %s[%s]", ins.Result, ins.Array, ins.Index)
}
