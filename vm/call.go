package vm

import (
	"github.com/elliotchance/ok/ast"
)

// Call tells the VM to jump to another function.
type Call struct {
	FunctionName string
	Arguments    []string
	Results      []string
}

// Execute implements the Instruction interface for the VM.
func (ins *Call) Execute(registers map[string]*ast.Literal, _ *int, vm *VM) error {
	results, err := vm.call(ins.FunctionName, ins.Arguments)
	if err != nil {
		return err
	}

	for i, result := range results {
		vm.stack[len(vm.stack)-2][ins.Results[i]] = vm.stack[len(vm.stack)-1][result]
	}

	vm.stack = vm.stack[:len(vm.stack)-1]
	vm.Return = nil
	return nil
}
