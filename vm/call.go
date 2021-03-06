package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// Call tells the VM to jump to another function.
type Call struct {
	FunctionName string
	Arguments    Registers
	Results      Registers

	// Type is the resolved interface when calling constructors. In any other
	// case this should be types.Any.
	Type TypeRegister

	// Pos is used to append to the call stack.
	Pos string
}

// Execute implements the Instruction interface for the VM.
func (ins *Call) Execute(_ *int, vm *VM) error {
	// TODO(elliot): This is a hack that matches up with compiler/call.go.
	parentScope := map[string]*ast.Literal{}
	funcName := ins.FunctionName
	if ins.FunctionName[0] == '*' {
		funcLit := vm.Get(Register(ins.FunctionName[1:]))
		funcName = funcLit.Value
		parentScope = funcLit.Map
	}

	ty := vm.Types[ins.Type]
	results, err := vm.call(funcName, ins.Arguments, parentScope, ty, ins.Pos)
	if err != nil {
		return err
	}

	for i, result := range results {
		vm.set(ins.Results[i], vm.Get(result), 2)
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
