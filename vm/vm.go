package vm

import (
	"ok/ast"
)

// VM is an instance of a virtual machine to run ok instructions.
type VM struct {
	fns    map[string]*CompiledFunc
	Return []string
	stack  []map[string]*ast.Literal
}

// NewVM will create a new VM ready to run the provided instructions.
func NewVM(fns map[string]*CompiledFunc) *VM {
	return &VM{
		fns: fns,
	}
}

// Run will run the program.
func (vm *VM) Run() error {
	_, err := vm.call("main", nil)

	return err
}

// Run will run the program.
func (vm *VM) call(name string, arguments []string) ([]string, error) {
	// TODO(elliot): Check function exists, especially main.

	registers := map[string]*ast.Literal{}

	// Copy the registers of this context into the new call context.
	fn := vm.fns[name]
	for i, arg := range arguments {
		registers[fn.Arguments[i]] = vm.stack[len(vm.stack)-1][arg]
	}
	vm.stack = append(vm.stack, registers)
	//fmt.Println("add")
	//defer func() {
	//	//if len(vm.stack) > 1 {
	//	//}
	//	//fmt.Println("delete")
	//}()

	totalInstructions := len(fn.Instructions)
	for i := 0; i < totalInstructions; i++ {
		ins := fn.Instructions[i]
		err := ins.Execute(registers, &i, vm)
		if err != nil {
			return nil, err
		}

		if vm.Return != nil {
			return vm.Return, nil
		}
	}

	return nil, nil
}
