package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// CastString returns a string value of a value.
type CastString struct {
	X, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *CastString) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, &ast.Literal{
		Kind:  "string",
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
	vm.Set(ins.Result, &ast.Literal{
		Kind:  "number",
		Value: fmt.Sprintf("%d", int([]rune(vm.Get(ins.X).Value)[0])),
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
		Kind:  "char",
		Value: fmt.Sprintf("%c", number.Int(number.NewNumber(vm.Get(ins.X).Value))),
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *CastChar) String() string {
	return fmt.Sprintf("%s = char %s", ins.Result, ins.X)
}
