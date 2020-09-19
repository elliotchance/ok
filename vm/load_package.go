package vm

import (
	"fmt"
	"strings"

	"github.com/elliotchance/ok/ast"
)

// LoadPackage loads a package at runtime into a register.
type LoadPackage struct {
	Result      Register
	PackageName string
}

// Execute implements the Instruction interface for the VM.
func (ins *LoadPackage) Execute(_ *int, vm *VM) error {
	// TODO(elliot): This is a horrible implementation. Needs to be rewritten.

	pkgMap := map[string]*ast.Literal{}

	prefix := ins.PackageName + "."
	for fnName, fn := range vm.fns {
		if strings.HasPrefix(fnName, prefix) {
			pkgMap[fnName[len(prefix):]] = &ast.Literal{
				Kind:  fn.Type,
				Value: fnName,
			}
		}
	}

	vm.Set(ins.Result, &ast.Literal{
		Map: pkgMap,
	})

	return nil
}

func (ins *LoadPackage) String() string {
	return fmt.Sprintf("%s = import \"%s\"", ins.Result, ins.PackageName)
}
