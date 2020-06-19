package instruction

import (
	"ok/ast"
)

// MapGet gets a value from the map by its key.
type MapGet struct {
	Map, Key, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *MapGet) Execute(registers map[string]*ast.Literal, _ *int) error {
	key := registers[ins.Key].Value
	registers[ins.Result] = registers[ins.Map].Map[key]

	return nil
}
