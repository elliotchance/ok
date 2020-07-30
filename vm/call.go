package vm

import (
	"fmt"
)

// Call tells the VM to jump to another function.
type Call struct {
	FunctionName string
	Arguments    Registers
	Results      Registers
}

// Execute implements the Instruction interface for the VM.
func (ins *Call) Execute(_ *int, vm *VM) error {
	results, err := vm.call(ins.FunctionName, ins.Arguments)
	if err != nil {
		return err
	}

	for i, result := range results {
		vm.Stack[len(vm.Stack)-2][ins.Results[i]] = vm.Stack[len(vm.Stack)-1][result]
	}

	vm.Stack = vm.Stack[:len(vm.Stack)-1]

	// We cannot rollback the FinallyBlocks stack here because we may need to
	// run some of them as part of the return. They will be removed in the
	// parent caller when it's finished with them

	vm.Return = nil

	return nil
}

// String is the human-readable description of the instruction.
func (ins *Call) String() string {
	return fmt.Sprintf("%s = %s%s", ins.Results, ins.FunctionName, ins.Arguments)
}
