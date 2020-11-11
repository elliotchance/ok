package vm

import (
	"fmt"
	"os"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
)

// Open opens a file.
type Open struct {
	Path, Result Register
}

// Execute implements the Instruction interface for the VM.
func (ins *Open) Execute(_ *int, vm *VM) error {
	f, err := os.OpenFile(vm.Get(ins.Path).Value, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		vm.Raise(err.Error())
	}

	vm.Set(ins.Result, &ast.Literal{
		Kind: types.Data,
		File: f,
	})

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Open) String() string {
	return fmt.Sprintf("%s = os.Open(%s)", ins.Result, ins.Path)
}
