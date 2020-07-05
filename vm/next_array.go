package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// NextArray is used to tick an array iterator forward.
type NextArray struct {
	Array       string // Register In (array): Containing the iterating array.
	Cursor      string // Register In (number): Containing the current position.
	KeyResult   string // Register Out (any): Load the key into this register.
	ValueResult string // Register Out (any): Load the value into this register.
	Result      string // Register Out (bool): Still more items?
}

// Execute implements the Instruction interface for the VM.
func (ins *NextArray) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	array := registers[ins.Array].Array
	pos := number.Int(number.NewNumber(registers[ins.Cursor].Value))
	hasMore := pos < len(array)
	registers[ins.Result] = asttest.NewLiteralBool(hasMore)
	if hasMore {
		registers[ins.KeyResult] = asttest.NewLiteralNumber(fmt.Sprintf("%d", pos))
		registers[ins.ValueResult] = array[pos]
		registers[ins.Cursor] = asttest.NewLiteralNumber(fmt.Sprintf("%d", pos+1))
	}

	return nil
}
