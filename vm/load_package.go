package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// LoadPackage loads a package at runtime into a register.
type LoadPackage struct {
	Result      Register
	PackageName string
}

// Execute implements the Instruction interface for the VM.
func (ins *LoadPackage) Execute(_ *int, vm *VM) error {
	vm.fns["__load_package"] = vm.packageFunctions[ins.PackageName]
	call := &Call{
		FunctionName: "__load_package",
		Results:      Registers{"__pkg"},
		Type:         vm.packageFunctions[ins.PackageName].Type,
	}
	err := call.Execute(nil, vm)
	if err != nil {
		return err
	}

	vm.Set(ins.Result, &ast.Literal{
		Map: vm.Stack[len(vm.Stack)-1][StateRegister].Map["__pkg"].Map,
	})

	return nil
}

func (ins *LoadPackage) String() string {
	return fmt.Sprintf("%s = import \"%s\"", ins.Result, ins.PackageName)
}
