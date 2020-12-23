package vm

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"regexp"
	"runtime/debug"
	"sort"
	"time"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/types"
)

const (
	// StateRegister is a reserved register for holding the map of the state.
	// Effectively the instance or "this" context.
	StateRegister = "0"

	// StackRegister holds the stack information.
	StackRegister = "__stack"
)

// These are populated with the generated lib.go file. See Makefile.
var Packages map[string]*File

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
	ErrType  *types.Type
	ErrValue *ast.Literal

	// FinallyBlocks are stacked with stack.
	FinallyBlocks [][]*FinallyBlock

	packageFunctions map[string]*CompiledFunc

	// Used for generating random numbers within this VM.
	rand *rand.Rand

	// Types can be referenced by instructions.
	Types map[TypeRegister]*types.Type
}

// NewVM will create a new VM ready to run the provided instructions.
func NewVM(pkg string) *VM {
	return &VM{
		fns:              make(map[string]*CompiledFunc),
		pkg:              pkg,
		Stdout:           os.Stdout,
		packageFunctions: make(map[string]*CompiledFunc),
		rand:             rand.New(rand.NewSource(int64(time.Now().Nanosecond()))),
		Types:            map[TypeRegister]*types.Type{},
	}
}

// Run will run the program. That is, execute the "main" function. If there is
// no main() method then nothing will run, but no error will be raised.
//
// TODO(elliot): Change missing main into an error in the future. It was done
//  this way so I could use "ok run" like an "ok compile" (that didn't exist at
//  the time) for compiling the standard libraries.
func (vm *VM) Run() error {
	for uniqueName, fn := range vm.fns {
		if fn.Name == "main" {
			_, err := vm.call(uniqueName, nil, map[string]*ast.Literal{}, types.Any)

			vm.catchUnhandledError()

			return err
		}
	}

	return nil
}

// Run will run the tests only.
func (vm *VM) RunTests(verbose bool, filter *regexp.Regexp) error {
	for _, t := range vm.tests {
		if !filter.MatchString(t.TestName) {
			continue
		}

		if verbose {
			fmt.Println("#", t.TestName)
		}

		vm.CurrentTestPassed = true
		err := vm.runTest(t, map[string]*ast.Literal{})
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

func (vm *VM) appendStack(stackDescription string, parentScope map[string]*ast.Literal, returnType *types.Type) {
	registers := map[Register]*ast.Literal{
		Register(StackRegister): asttest.NewLiteralString(stackDescription),
	}
	vm.Stack = append(vm.Stack, registers)

	// Setup a new state, even we don't use it.
	vm.Set(StateRegister, &ast.Literal{
		Kind: returnType,
		Map: map[string]*ast.Literal{
			StateRegister: {
				Kind: types.Any,
				Map:  parentScope,
			},
		},
	})
}

func (vm *VM) call(
	name string,
	arguments []Register,
	parentScope map[string]*ast.Literal,
	returnType *types.Type,
) ([]Register, error) {
	// TODO(elliot): Check function exists, especially main.

	// Copy the registers of this context into the new call context.
	fn := vm.fns[name]
	if fn == nil {
		panic("no such function: " + name)
	}

	stackDesc := stackDescription(fn.Pos, fn.Name)
	vm.appendStack(stackDesc, parentScope, returnType)

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

	returns, err := vm.runInstructions(fn.Name, fn.Instructions, false)
	if err != nil {
		return nil, err
	}

	for _, fb := range finallyBlocks {
		if fb.Run {
			// Ignore returns here.
			// TODO(elliot): The compiler must disallow return
			//  statements within a finally block.
			_, err := vm.runInstructions(fn.Name, fb.Instructions, true)
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

	var keys []string
	for key := range vm.Stack[len(vm.Stack)-1] {
		keys = append(keys, string(key))
	}
	sort.Strings(keys)

	for _, n := range keys {
		if n == StateRegister {
			fmt.Printf("  %s: <State>\n", n)
		} else {
			fmt.Printf("  %s: %v\n", n, vm.Stack[len(vm.Stack)-1][Register(n)])
		}
	}

	fmt.Printf("\nState:\n")

	keys = nil
	for key := range vm.Stack[len(vm.Stack)-1][StateRegister].Map {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, n := range keys {
		fmt.Printf("  %s: %v\n", n, vm.Stack[len(vm.Stack)-1][StateRegister].Map[n])
	}

	fmt.Printf("\nStack:\n")

	for _, v := range vm.Stack {
		fmt.Printf("  %s\n", v[StackRegister].Value)
	}
}

func (vm *VM) recoverPanic(funcName string, ins []Instruction, i *int) func() {
	return func() {
		if r := recover(); r != nil {
			// i+1 because the first instruction shown in "ok asm" is #1.
			fmt.Printf("VM panicked in function %s at instruction #%d: %s\n\n",
				funcName, *i+1, ins[*i].String())
			vm.dumpMemory()

			fmt.Println()
			fmt.Println(r)
			fmt.Println(string(debug.Stack()))
			os.Exit(1)
		}
	}
}

func (vm *VM) runInstructions(funcName string, ins []Instruction, inFinally bool) ([]Register, error) {
	i := 0
	defer vm.recoverPanic(funcName, ins, &i)()

	totalInstructions := len(ins)
	for ; i < totalInstructions; i++ {
		ins := ins[i]

		// If we are in an error state, we keep moving forward until we find an
		// appropriate error handler.
		//
		// TODO(elliot): This can not differentiate a handler in a lower scope
		//  that should be ignored. It would be best for raise to provide a jump
		//  to the first (or each) of the handlers, ideally.
		if vm.ErrType != nil && !inFinally {
			if on, ok := ins.(*On); ok {
				if on.Type == NoTypeRegister {
					// An empty type signals the end of the error handlers.
					// Making it to here means none of the error handlers worked
					// for us. We need to return now and let the parent scope
					// try to handle it.
					return nil, nil
				}

				switch vm.Types[on.Type].Name {
				case vm.ErrType.Name,
					// TODO(elliot): This is a stupid hack for now. This was
					//  created before interfaces could determine this properly.
					"error.Error", "Error":

					// TODO(elliot): The err register might be better as a
					//  fixed position register rather than a variable?
					vm.Set("err", vm.ErrValue)

					// We found the handler. Remove the error state and continue
					// as normal. There is a jump at the end of the handler that
					// will launch us out when it's done.
					vm.ErrType = nil
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

func (vm *VM) catchUnhandledError() {
	if vm.ErrType != nil {
		// TODO(elliot): This needs to be handled much more gracefully.
		panic(fmt.Sprintf("unhandled %s: %+v", vm.ErrType, vm.ErrValue.Map["Error"]))
	}
}

func stackDescription(pos, funcName string) string {
	return fmt.Sprintf("%s|%s", pos, funcName)
}

func (vm *VM) runTest(test *CompiledTest, parentScope map[string]*ast.Literal) error {
	vm.CurrentTestName = test.TestName

	stackDesc := stackDescription(test.Pos,
		fmt.Sprintf("test \"%s\"", test.TestName))
	vm.appendStack(stackDesc, parentScope, types.Any)
	_, err := vm.runInstructions(test.TestName, test.Instructions, false)

	vm.catchUnhandledError()

	// Make sure we roll the stack back after each test. This is normally
	// handled in Call for functions that are not tests.
	vm.Stack = vm.Stack[:len(vm.Stack)-1]

	return err
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
	switch {
	case register[0] == '^':
		parentScope := vm.Stack[len(vm.Stack)-1][StateRegister].Map[StateRegister].Map
		parentScope[string(register[1:])] = val

	case isRegister(register):
		vm.Stack[len(vm.Stack)-offset][register] = val

	default:
		vm.Stack[len(vm.Stack)-offset][StateRegister].Map[string(register)] = val
	}
}

// Get will get a register.
func (vm *VM) get(register Register, offset int) (lit *ast.Literal) {
	if register[0] == '^' {
		parentScope := vm.Stack[len(vm.Stack)-1][StateRegister].Map[StateRegister].Map

		return parentScope[string(register[1:])]
	}

	if isRegister(register) {
		return vm.Stack[len(vm.Stack)-offset][register]
	}

	return vm.Stack[len(vm.Stack)-offset][StateRegister].Map[string(register)]
}

// Get will get a register.
func (vm *VM) Get(register Register) *ast.Literal {
	return vm.get(register, 1)
}

func (vm *VM) Raise(message string) {
	vm.ErrType = types.ErrorInterface
	vm.ErrValue = &ast.Literal{
		Kind:  types.ErrorInterface,
		Array: []*ast.Literal{asttest.NewLiteralString("Error")},
		Map: map[string]*ast.Literal{
			"Error": asttest.NewLiteralString(message),
		},
	}
}

func (vm *VM) LoadPackage(pkgVariable, packageName string) error {
	file, err := Load(packageName)
	if err != nil {
		return err
	}

	return vm.LoadFile(pkgVariable, file)
}

func (vm *VM) LoadFile(pkgVariable string, file *File) error {
	for k, v := range file.Funcs {
		vm.fns[k] = v
	}

	for k, v := range file.Types {
		vm.Types[k] = v
	}

	vm.tests = append(vm.tests, file.Tests...)

	for packageName := range file.Imports {
		err := vm.LoadPackage(path.Base(packageName), packageName)
		if err != nil {
			return err
		}
	}

	vm.packageFunctions[pkgVariable] = file.PackageFunc

	return nil
}
