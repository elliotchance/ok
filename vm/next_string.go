package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// NextString is used to tick a string iterator forward.
type NextString struct {
	Str         Register // In (string): Containing the iterating string.
	Cursor      Register // In (number): Containing the current position.
	KeyResult   Register // Out (any): Load the key into this register.
	ValueResult Register // Out (any): Load the value into this register.
	Result      Register // Out (bool): Still more items?
}

// Execute implements the Instruction interface for the VM.
func (ins *NextString) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	str := []rune(registers[ins.Str].Value)
	pos := number.Int(number.NewNumber(registers[ins.Cursor].Value))
	hasMore := pos < len(str)
	registers[ins.Result] = asttest.NewLiteralBool(hasMore)
	if hasMore {
		registers[ins.KeyResult] = asttest.NewLiteralNumber(fmt.Sprintf("%d", pos))
		registers[ins.ValueResult] = asttest.NewLiteralChar(str[pos])
		registers[ins.Cursor] = asttest.NewLiteralNumber(fmt.Sprintf("%d", pos+1))
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *NextString) String() string {
	return fmt.Sprintf("%s, %s = next from %s; increment %s; has more %s",
		ins.ValueResult, ins.KeyResult, ins.Str, ins.Cursor, ins.Result)
}
