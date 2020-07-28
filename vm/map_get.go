package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// MapGet gets a value from the map by its key.
type MapGet struct {
	Map, Key, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *MapGet) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	key := registers[ins.Key].Value
	registers[ins.Result] = registers[ins.Map].Map[key]

	return nil
}

// String is the human-readable description of the instruction.
func (ins *MapGet) String() string {
	return fmt.Sprintf("%s = %s[%s]", ins.Result, ins.Map, ins.Key)
}
