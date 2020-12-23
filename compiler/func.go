package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

// CompileFunc translates a single function into a set of instructions. The
// number of instructions returned may be zero.
func CompileFunc(
	fn *ast.Func,
	file *vm.File,
	parentFunc *vm.CompiledFunc,
) (*vm.CompiledFunc, error) {
	compiled := vm.NewCompiledFunc(fn, parentFunc)

	// Make sure we clear state that shouldn't be serialized.
	defer func() {
		compiled.Parent = nil
		compiled.DeferredFuncsToCompile = nil
	}()

	// All variables in a function are stored internally as a map right now. So
	// the first thing we need to do is initialise the map. Whether we return it
	// at the end will depend on the return type.
	//
	// It is important that we always use a map (even if it's not returned)
	// because we may need it for a subscope (closure).
	isObject := len(fn.Returns) == 1 && fn.Returns[0].Name == fn.Name

	// Load the arguments from the registers.
	for _, arg := range fn.Arguments {
		compiled.NextRegister()
		compiled.NewVariable(arg.Name, arg.Type)

		compiled.Arguments = append(compiled.Arguments, arg.Name)
	}

	err := compileBlock(compiled, fn.Statements, nil, nil,
		file, nil)
	if err != nil {
		return nil, err
	}

	// If this is an object, always returns its state.
	if isObject {
		compiled.Append(&vm.Return{
			Results: []vm.Register{vm.StateRegister},
		})
	}

	// Now we have finished compiling this scope (and so have discovered and
	// resolved the type of all variables) we can no compile all the deferred
	// function literals that might reference variables in this scope.
	instructions := len(compiled.Instructions)
	for _, fn := range compiled.DeferredFuncsToCompile {
		cf, err := CompileFunc(fn.Func, file, compiled)
		if err != nil {
			return nil, err
		}

		file.Funcs[fn.Func.UniqueName] = cf
		compiled.Append(&vm.AssignFunc{
			Result:     fn.Register,
			Type:       file.AddType(fn.Func.Type()),
			UniqueName: fn.Func.UniqueName,
		})

		compiled.Append(&vm.ParentScope{
			X: fn.Register,
		})
	}

	// However... the instructions that were just appended to this scope need to
	// be hoisted to the top otherwise the variables containing the function
	// literals will be missing until the end.
	compiled.Instructions = append(
		compiled.Instructions[instructions:],
		compiled.Instructions[:instructions]...)

	return compiled, nil
}

func compileFunc(
	compiledFunc *vm.CompiledFunc,
	fn *ast.Func,
) ([]vm.Register, []*types.Type, error) {
	// We cannot compile a function literal at this point because we
	// probably haven't finished compiling the current scope. This is
	// important because the function literal may reference variables in
	// this scope that haven't been discovered (or type resolved) yet.
	//
	// Instead we will defer the compilation of this literal to after we're
	// finished. We can reserve the register that will hold the compiled
	// function later so everything lines up without any post correction.
	funcLitRegister := compiledFunc.NextRegister()
	compiledFunc.DeferredFuncsToCompile = append(
		compiledFunc.DeferredFuncsToCompile, vm.DeferredFunc{
			Register: funcLitRegister,
			Func:     fn,
		})

	return []vm.Register{funcLitRegister}, []*types.Type{fn.Type()}, nil
}
