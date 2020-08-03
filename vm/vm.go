package vm

import (
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/elliotchance/ok/ast"
)

// StateRegister is a reserved register for holding the map of the state.
// Effectively the instance or "this" context.
const StateRegister = "0"

// InternalDefinition is used by Lib to hold the definitions and compiled code
// for internal functions.
type InternalDefinition struct {
	CompiledFunc *CompiledFunc
	FuncDef      *ast.Func
}

// These are populated with the generated lib.go file. See Makefile.
var Lib map[string]*InternalDefinition
var Packages map[string]bool

// CompiledTest is a runnable test.
type CompiledTest struct {
	*CompiledFunc
	TestName string
}

// VM is an instance of a virtual machine to run ok instructions.
type VM struct {
	fns    map[string]*CompiledFunc
	Return []Register
	Stack  []map[Register]*ast.Literal
	tests  []*CompiledTest
	pkg    string
	Stdout io.Writer

	// Stats when running tests.
	TestsPass, TestsFailed int
	TotalAssertions        int
	CurrentTestName        string
	CurrentTestPassed      bool

	// ErrType will be non-empty once an error is raised. It contains the type
	// to match for a handler. ErrValue contains the actual error.
	ErrType  string
	ErrValue *ast.Literal

	// FinallyBlocks are stacked with stack.
	FinallyBlocks [][]*FinallyBlock
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

func (vm *VM) appendStack() {
	registers := map[Register]*ast.Literal{}
	vm.Stack = append(vm.Stack, registers)

	// Setup a new state, even we don't use it.
	vm.Set(StateRegister, &ast.Literal{
		// TODO(elliot): This should probably be the actual type returned.
		Kind: "any",
		Map:  make(map[string]*ast.Literal),
	})
}

func (vm *VM) call(name string, arguments []Register) ([]Register, error) {
	// TODO(elliot): Check function exists, especially main.

	vm.appendStack()

	// Copy the registers of this context into the new call context.
	fn := vm.fns[name]

	// TODO(elliot): It's probably not a good idea to fallback onto the standard
	//  lib. Perhaps the compiler can append these functions so this isn't
	//  necessary?
	if fn == nil {
		fn = Lib[name].CompiledFunc
	}

	// Setup the finally blocks. Copy so they all start disabled.
	var finallyBlocks []*FinallyBlock
	for _, ins := range fn.Finally {
		finallyBlocks = append(finallyBlocks, &FinallyBlock{
			Run:          false,
			Instructions: ins,
		})
	}
	vm.FinallyBlocks = append(vm.FinallyBlocks, finallyBlocks)

	// Copy the arguments in.
	for i, arg := range arguments {
		vm.Set(Register(fn.Arguments[i]), vm.get(arg, 2))
	}

	returns, err := vm.runInstructions(name, fn.Instructions, false)
	if err != nil {
		return nil, err
	}

	for _, fb := range finallyBlocks {
		if fb.Run {
			// Ignore returns here.
			// TODO(elliot): The compiler must disallow return
			//  statements within a finally block.
			_, err := vm.runInstructions(name, fb.Instructions, true)
			if err != nil {
				return nil, err
			}
		}
	}

	vm.FinallyBlocks = vm.FinallyBlocks[:len(vm.FinallyBlocks)-1]

	return returns, nil
}

func (vm *VM) dumpMemory() {
	// TODO(elliot): It would be nice to have both of these sorted.

	fmt.Printf("Registers:\n")

	for n, v := range vm.Stack[len(vm.Stack)-1] {
		if n == StateRegister {
			fmt.Printf("  %s: <State>\n", n)
		} else {
			fmt.Printf("  %s: %v\n", n, v)
		}
	}

	fmt.Printf("\nState:\n")

	for n, v := range vm.Stack[len(vm.Stack)-1][StateRegister].Map {
		fmt.Printf("  %s: %v\n", n, v)
	}
}

func (vm *VM) runInstructions(funcName string, ins []Instruction, inFinally bool) ([]Register, error) {
	i := 0

	defer func() {
		if r := recover(); r != nil {
			// i+1 because the first instruction shown in "ok asm" is #1.
			fmt.Printf("VM panicked in function %s at instruction #%d: %s\n\n",
				funcName, i+1, ins[i].String())
			vm.dumpMemory()

			fmt.Println()
			fmt.Println(r)
			fmt.Println(string(debug.Stack()))
			os.Exit(1)
		}
	}()

	totalInstructions := len(ins)
	for ; i < totalInstructions; i++ {
		ins := ins[i]

		// If we are in an error state, we keep moving forward until we find an
		// appropriate error handler.
		//
		// TODO(elliot): This can not differentiate a handler in a lower scope
		//  that should be ignored. It would be best for raise to provide a jump
		//  to the first (or each) of the handlers, ideally.
		if vm.ErrType != "" && !inFinally {
			if on, ok := ins.(*On); ok {
				switch on.Type {
				case vm.ErrType,
					// TODO(elliot): This is a stupid hack for now. This was
					//  created before interfaces could determine this properly.
					"Error":

					// TODO(elliot): The err register might be better as a
					//  fixed position register rather than a variable?
					vm.Set(Register("err"), vm.ErrValue)

					// We found the handler. Remove the error state and continue
					// as normal. There is a jump at the end of the handler that
					// will launch us out when it's done.
					vm.ErrType = ""

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

		err := ins.Execute(&i, vm)
		if err != nil {
			return nil, err
		}

		if vm.Return != nil && !inFinally {
			return vm.Return, nil
		}
	}

	return nil, nil
}

func (vm *VM) runTest(testName string, instructions []Instruction) error {
	vm.CurrentTestName = testName

	vm.appendStack()

	totalInstructions := len(instructions)
	for i := 0; i < totalInstructions; i++ {
		ins := instructions[i]
		err := ins.Execute(&i, vm)
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

func isRegister(register Register) bool {
	r := register[0]

	return r >= '0' && r <= '9'
}

// Set will set a register.
func (vm *VM) Set(register Register, val *ast.Literal) {
	vm.set(register, val, 1)
}

func (vm *VM) set(register Register, val *ast.Literal, offset int) {
	if isRegister(register) {
		vm.Stack[len(vm.Stack)-offset][register] = val
	} else {
		vm.Stack[len(vm.Stack)-offset][StateRegister].Map[string(register)] = val
	}
}

// Get will get a register.
func (vm *VM) get(register Register, offset int) *ast.Literal {
	if isRegister(register) {
		return vm.Stack[len(vm.Stack)-offset][register]
	}

	return vm.Stack[len(vm.Stack)-offset][StateRegister].Map[string(register)]
}

// Get will get a register.
func (vm *VM) Get(register Register) *ast.Literal {
	return vm.get(register, 1)
}
