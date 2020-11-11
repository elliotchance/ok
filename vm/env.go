package vm

import (
	"fmt"
	"os"

	"github.com/elliotchance/ok/ast/asttest"
)

// EnvGet gets an environment variable.
type EnvGet struct {
	Name   Register // In string
	Value  Register // Out string
	Exists Register // Out bool
}

// Execute implements the Instruction interface for the VM.
func (ins *EnvGet) Execute(_ *int, vm *VM) error {
	r := vm.Get(ins.Name).Value
	value, exists := os.LookupEnv(r)
	vm.Set(ins.Value, asttest.NewLiteralString(value))
	vm.Set(ins.Exists, asttest.NewLiteralBool(exists))

	return nil
}

// String is the human-readable description of the instruction.
func (ins *EnvGet) String() string {
	return fmt.Sprintf("%s, %s = runtime.LookupEnv(%s)", ins.Value, ins.Exists, ins.Name)
}

// EnvSet sets an environment variable.
type EnvSet struct {
	Name  Register // In string
	Value Register // In string
}

// Execute implements the Instruction interface for the VM.
func (ins *EnvSet) Execute(_ *int, vm *VM) error {
	name := vm.Get(ins.Name).Value
	value := vm.Get(ins.Value).Value
	err := os.Setenv(name, value)
	if err != nil {
		vm.Raise(err.Error())
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *EnvSet) String() string {
	return fmt.Sprintf("runtime.SetEnv(%s, %s)", ins.Name, ins.Value)
}

// EnvUnset unsets an environment variable.
type EnvUnset struct {
	Name Register // In string
}

// Execute implements the Instruction interface for the VM.
func (ins *EnvUnset) Execute(_ *int, vm *VM) error {
	name := vm.Get(ins.Name).Value
	err := os.Unsetenv(name)
	if err != nil {
		vm.Raise(err.Error())
	}

	return nil
}

// String is the human-readable description of the instruction.
func (ins *EnvUnset) String() string {
	return fmt.Sprintf("runtime.UnsetENv(%s)", ins.Name)
}
