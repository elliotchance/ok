package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// Assign sets a variable to the result of an expression.
type Assign struct {
	Result   Register // Out (any)
	Register Register // In (any)
}

// Execute implements the Instruction interface for the VM.
func (ins *Assign) Execute(_ *int, vm *VM) error {
	get := vm.Get(ins.Register)
	vm.Set(ins.Result, &ast.Literal{
		Kind:  get.Kind,
		Value: get.Value,
		Array: get.Array,
		Map:   get.Map,
		Pos:   get.Pos,
	})

	return nil
}

func (ins *Assign) String() string {
	return fmt.Sprintf("%s = %s", ins.Result, ins.Register)
}

// AssignSymbol sets a variable to the symbol.
type AssignSymbol struct {
	Result Register       // Out (any)
	Symbol SymbolRegister // In (any)
}

// Execute implements the Instruction interface for the VM.
func (ins *AssignSymbol) Execute(_ *int, vm *VM) error {
	symbol := vm.Symbols[ins.Symbol]
	vm.Set(ins.Result, symbol)

	return nil
}

func (ins *AssignSymbol) String() string {
	return fmt.Sprintf("%s = %s", ins.Result, ins.Symbol)
}

// AssignFunc sets a function to a register.
type AssignFunc struct {
	Result     Register     // Out (any)
	Type       TypeRegister // In
	UniqueName string       // In
}

// Execute implements the Instruction interface for the VM.
func (ins *AssignFunc) Execute(_ *int, vm *VM) error {
	vm.Set(ins.Result, &ast.Literal{
		Kind:  vm.Types[ins.Type],
		Value: ins.UniqueName,
	})

	return nil
}

func (ins *AssignFunc) String() string {
	return fmt.Sprintf("%s = %s (%s)", ins.Result, ins.UniqueName, ins.Type)
}
