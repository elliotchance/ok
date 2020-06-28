package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// NextMap is used to tick a map iterator forward.
type NextMap struct {
	Map         string // Register In (map): Containing the iterating map.
	Cursor      string // Register In (number): Containing the current position.
	KeyResult   string // Register Out (any): Load the key into this register.
	ValueResult string // Register Out (any): Load the value into this register.
	Result      string // Register Out (bool): Still more items?
}

// Execute implements the Instruction interface for the VM.
func (ins *NextMap) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	array := registers[ins.Map].Array
	pos := number.Int(number.NewNumber(registers[ins.Cursor].Value))

	hasMore := pos < len(array)
	registers[ins.Result] = ast.NewLiteralBool(hasMore)
	if hasMore {
		m := registers[ins.Map].Map
		registers[ins.KeyResult] = array[pos]
		registers[ins.ValueResult] = m[array[pos].Value]
		registers[ins.Cursor].Value = fmt.Sprintf("%d", pos+1)
	}

	return nil
}
