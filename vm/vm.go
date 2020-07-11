package vm

import (
	"fmt"
	"io"
	"os"

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
	Stdout io.Writer

	// Stats when running tests.
	TestsPass, TestsFailed int
	TotalAssertions        int
	CurrentTestName        string
	CurrentTestPassed      bool

	// Err will be non-empty once an error is raised. It contains the type to
	// match for a handler.
	Err string
}

// NewVM will create a new VM ready to run the provided instructions.
func NewVM(fns map[string]*CompiledFunc, tests []*CompiledTest, pkg string) *VM {
	return &VM{
		fns:    fns,
		tests:  tests,
		pkg:    pkg,
		Stdout: os.Stdout,
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

		// If we are in an error state, we keep moving forward until we find an
		// appropriate error handler.
		//
		// TODO(elliot): This can not differentiate a handler in a lower scope
		//  that should be ignored. It would be best for raise to provide a jump
		//  to the first (or each) of the handlers, ideally.
		if vm.Err != "" {
			if on, ok := ins.(*On); ok {
				switch on.Type {
				case vm.Err:
					// We found the handler. Remove the error state and continue
					// as normal. There is a jump at the end of the handler that
					// will launch us out when it's done.
					vm.Err = ""

				case "":
					// An empty type signals the end of the error handlers.
					// Making it to here means none of the error handlers worked
					// for us. We need to return now and let the parent scope
					// try to handle it.
					return nil, nil
				}

			}

			continue
		}

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

func (vm *VM) assert(pass bool, left, op, right, pos string) {
	if !pass {
		fmt.Printf("%s: %s: %s: assert(%s %s %s) failed\n",
			vm.pkg, pos, vm.CurrentTestName, left, op, right)
		vm.CurrentTestPassed = false
	}
	vm.TotalAssertions++
}
