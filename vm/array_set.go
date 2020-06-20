package vm

import (
	"ok/ast"
	"ok/number"
)

// ArraySet sets a number value to an index.
type ArraySet struct {
	Array, Index, Value string
}

// Execute implements the Instruction interface for the VM.
func (ins *ArraySet) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	index := number.NewNumber(registers[ins.Index].Value).Int64()
	registers[ins.Array].Array[index] = registers[ins.Value]

	return nil
}
