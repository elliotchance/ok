package vm

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/number"
)

// StringIndex returns a character from an index of a string.
type StringIndex struct {
	String, Index, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *StringIndex) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	index := number.Int64(number.NewNumber(registers[ins.Index].Value))

	// TODO(elliot): This won't work with multibyte characters.
	registers[ins.Result] = asttest.NewLiteralChar([]rune(registers[ins.String].Value)[index])

	return nil
}
