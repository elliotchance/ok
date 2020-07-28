package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// ArraySet sets a number value to an index.
type ArraySet struct {
	Array, Index, Value Register
}

// Execute implements the Instruction interface for the VM.
func (ins *ArraySet) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	index := number.Int64(number.NewNumber(registers[ins.Index].Value))
	registers[ins.Array].Array[index] = registers[ins.Value]

	return nil
}

// String is the human-readable description of the instruction.
func (ins *ArraySet) String() string {
	return fmt.Sprintf("%s[%s] = %s", ins.Array, ins.Index, ins.Value)
}
