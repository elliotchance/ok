package vm

import (
	"fmt"
	"strconv"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
	"github.com/elliotchance/ok/types"
)

// CastString returns a string value of a value.
type CastString struct {
	X, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *CastString) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, &ast.Literal{
		Kind:  types.String,
		Value: renderLiteral(vm.Get(ins.X), false),
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *CastString) String() string {
	return fmt.Sprintf("%s = string %s", ins.Result, ins.X)
}

// CastNumber returns a number value of a value.
type CastNumber struct {
	X, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *CastNumber) Execute(_ *int, vm *VM) error {
	x := vm.Get(ins.X)
	var value string

	switch x.Kind.Kind {
	case types.KindChar:
		value = fmt.Sprintf("%d", int([]rune(x.Value)[0]))

	case types.KindString:
		_, err := strconv.ParseFloat(x.Value, 64)
		if err != nil {
			vm.Raise(fmt.Sprintf("not a number: %s", x.Value))

			return nil
		}

		value = x.Value
	}

	vm.Set(ins.Result, &ast.Literal{
		Kind:  types.Number,
		Value: value,
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *CastNumber) String() string {
	return fmt.Sprintf("%s = number %s", ins.Result, ins.X)
}

// CastChar returns a char value of a value.
type CastChar struct {
	X, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *CastChar) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, &ast.Literal{
		Kind:  types.Char,
		Value: fmt.Sprintf("%c", number.Int(number.NewNumber(vm.Get(ins.X).Value))),
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *CastChar) String() string {
	return fmt.Sprintf("%s = char %s", ins.Result, ins.X)
}

// CastData returns a data value of a value.
type CastData struct {
	X, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *CastData) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, &ast.Literal{
		Kind:  types.Data,
		Value: renderLiteral(vm.Get(ins.X), false),
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *CastData) String() string {
	return fmt.Sprintf("%s = data %s", ins.Result, ins.X)
}
