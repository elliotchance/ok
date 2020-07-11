package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/number"
)

// CastString returns a string value of a value.
type CastString struct {
	X, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *CastString) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = &ast.Literal{
		Kind:  "string",
		Value: renderLiteral(registers[ins.X], false),
	}

	return nil
}

// CastNumber returns a number value of a value.
type CastNumber struct {
	X, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *CastNumber) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = &ast.Literal{
		Kind:  "number",
		Value: fmt.Sprintf("%d", int([]rune(registers[ins.X].Value)[0])),
	}

	return nil
}

// CastChar returns a char value of a value.
type CastChar struct {
	X, Result string
}

// Execute implements the Instruction interface for the VM.
func (ins *CastChar) Execute(registers map[string]*ast.Literal, _ *int, _ *VM) error {
	registers[ins.Result] = &ast.Literal{
		Kind:  "char",
		Value: fmt.Sprintf("%c", number.Int(number.NewNumber(registers[ins.X].Value))),
	}

	return nil
}
