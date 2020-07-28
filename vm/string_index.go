package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// StringIndex returns a character from an index of a string.
type StringIndex struct {
	Str, Index, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *StringIndex) Execute(registers map[Register]*ast.Literal, _ *int, _ *VM) error {
	index := number.Int64(number.NewNumber(registers[ins.Index].Value))

	// TODO(elliot): This won't work with multibyte characters.
	registers[ins.Result] = asttest.NewLiteralChar([]rune(registers[ins.Str].Value)[index])

	return nil
}

// String is the human-readable description of the instruction.
func (ins *StringIndex) String() string {
	return fmt.Sprintf("%s = %s[%s]", ins.Result, ins.Str, ins.Index)
}
