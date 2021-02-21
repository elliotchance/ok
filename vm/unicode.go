package vm

import (
	"fmt"
	"unicode"

	"github.com/elliotchance/ok/ast/asttest"
)

// UnicodeIs is a collection of unicode instructions that provides OK with basic
// unicode functionality without having to build all the tables into the source
// code, yet.
//
// It was put into a single instruction as to not bloat the VM instructions
// since this is intended to be removed in the future.
type UnicodeIs struct {
	Op     Register // in: string
	Char   Register // in: char
	Result Register // out: bool
}

var unicodeIsFns = map[string]func(rune) bool{
	"Control": unicode.IsControl,
	"Digit":   unicode.IsDigit,
	"Graphic": unicode.IsGraphic,
	"Letter":  unicode.IsLetter,
	"Lower":   unicode.IsLower,
	"Mark":    unicode.IsMark,
	"Number":  unicode.IsNumber,
	"Print":   unicode.IsPrint,
	"Punct":   unicode.IsPunct,
	"Space":   unicode.IsSpace,
	"Symbol":  unicode.IsSymbol,
	"Title":   unicode.IsTitle,
	"Upper":   unicode.IsUpper,
}

// Execute implements the Instruction interface for the VM.
func (ins *UnicodeIs) Execute(_ *int, vm *VM) error {
	c := []rune(vm.Get(ins.Char).Value)[0]
	op := vm.Get(ins.Op).Value
	vm.Set(ins.Result, asttest.NewLiteralBool(unicodeIsFns[op](c)))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *UnicodeIs) String() string {
	return fmt.Sprintf("%s = __unicode_is(%s, %s)", ins.Result, ins.Op, ins.Char)
}

// UnicodeTo works the same way as UnicodeIs, but we need a separate instruction
// to handle returning a character rather than a bool.
type UnicodeTo struct {
	Op     Register // in: string
	Char   Register // in: char
	Result Register // out: char
}

var unicodeToFns = map[string]func(rune) rune{
	"Lower": unicode.ToLower,
	"Title": unicode.ToTitle,
	"Upper": unicode.ToUpper,
}

// Execute implements the Instruction interface for the VM.
func (ins *UnicodeTo) Execute(_ *int, vm *VM) error {
	c := []rune(vm.Get(ins.Char).Value)[0]
	op := vm.Get(ins.Op).Value
	vm.Set(ins.Result, asttest.NewLiteralChar(unicodeToFns[op](c)))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *UnicodeTo) String() string {
	return fmt.Sprintf("%s = __unicode_to(%s, %s)", ins.Result, ins.Op, ins.Char)
}
