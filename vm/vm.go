package vm

import (
	"fmt"

	"github.com/elliotchance/ok/ast"
)

// InternalDefinition is used by Lib to hold the definitions and compiled code
// for internal functions.
type InternalDefinition struct {
	CompiledFunc *CompiledFunc
	FuncDef      *ast.Func
}

// Lib is populated with the generated lib.go file. See Makefile.
var Lib map[string]*InternalDefinition

// CompiledTest is a runnable test.
type CompiledTest struct {
	*CompiledFunc
	TestName string
}

// VM is an instance of a virtual machine to run ok instructions.
type VM struct {
	fns    map[string]*CompiledFunc
	Return []string
	stack  []map[string]*ast.Literal
	tests  []*CompiledTest
	pkg    string

	// Stats when running tests.
	TestsPass, TestsFailed int
	TotalAssertions        int
	CurrentTestName        string
	CurrentTestPassed      bool
}

// NewVM will create a new VM ready to run the provided instructions.
func NewVM(fns map[string]*CompiledFunc, tests []*CompiledTest, pkg string) *VM {
	return &VM{
		fns:   fns,
		tests: tests,
		pkg:   pkg,
	}
}

// Run will run the program.
func (vm *VM) Run() error {
	_, err := vm.call("main", nil)

	return err
}

// Run will run the tests only.
func (vm *VM) RunTests() error {
	for _, t := range vm.tests {
		vm.CurrentTestPassed = true
		err := vm.runTest(t.TestName, t.Instructions)
		if err != nil {
			return err
		}

		if vm.CurrentTestPassed {
			vm.TestsPass++
		} else {
			vm.TestsFailed++
		}
	}

	return nil
}

func (vm *VM) call(name string, arguments []string) ([]string, error) {
	// TODO(elliot): Check function exists, especially main.

	registers := map[string]*ast.Literal{}

	// Copy the registers of this context into the new call context.
	fn := vm.fns[name]

	// TODO(elliot): It's probably not a good idea to fallback onto the standard
	//  lib. Perhaps the compiler can append these functions so this isn't
	//  necessary?
	if fn == nil {
		fn = Lib[name].CompiledFunc
	}

	for i, arg := range arguments {
		registers[fn.Arguments[i]] = vm.stack[len(vm.stack)-1][arg]
	}
	vm.stack = append(vm.stack, registers)

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

func (vm *VM) runTest(testName string, instructions []Instruction) error {
	vm.CurrentTestName = testName

	registers := map[string]*ast.Literal{}
	vm.stack = append(vm.stack, registers)

	totalInstructions := len(instructions)
	for i := 0; i < totalInstructions; i++ {
		ins := instructions[i]
		err := ins.Execute(registers, &i, vm)
		if err != nil {
			return err
		}
	}

	return nil
}

func (vm *VM) assert(pass bool, left, op, right string) {
	if !pass {
		fmt.Printf("%s: %s: assert(%s %s %s) failed\n",
			vm.pkg, vm.CurrentTestName, left, op, right)
		vm.CurrentTestPassed = false
	}
	vm.TotalAssertions++
}
