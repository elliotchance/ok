package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// DynamicCall tells the VM to jump to another function.
type DynamicCall struct {
	Variable  Register // func literal
	Arguments Register // []any for arguments
	Results   Register // []any for return values
}

// Execute implements the Instruction interface for the VM.
func (ins *DynamicCall) Execute(_ *int, vm *VM) error {
	var args Registers
	i := 0
	for _, arg := range vm.Get(ins.Arguments).Array {
		// TODO(elliot): This won't work with nested dynamic calls. Also, it's
		//  really nasty.
		register := Register(fmt.Sprintf("__%d", i))
		vm.Set(register, arg)
		args = append(args, register)
		i++
	}

	var results Registers
	funcLit := vm.Get(ins.Variable)
	for range ast.NewFuncFromPrototype(funcLit.Kind).Returns {
		register := Register(fmt.Sprintf("__%d", i))
		results = append(results, register)
		i++
	}

	realCall := &Call{
		FunctionName: "*" + string(ins.Variable),
		Arguments:    args,
		Results:      results,
	}

	err := realCall.Execute(nil, vm)
	if err != nil {
		return err
	}

	resultsAsArray := &ast.Literal{
		Kind:  "[]any",
		Array: make([]*ast.Literal, len(realCall.Results)),
	}

	for i, result := range realCall.Results {
		// TODO(elliot): It might be unsafe to pass them by reference this way.
		//  We might need to copy them.
		resultsAsArray.Array[i] = vm.Get(result)
	}

	vm.Set(ins.Results, resultsAsArray)

	return nil
}

// String is the human-readable description of the instruction.
func (ins *DynamicCall) String() string {
	return fmt.Sprintf("%s = reflect.Call(%s, %s)", ins.Results, ins.Variable, ins.Arguments)
}
