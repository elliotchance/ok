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
	constants map[string]*ast.Literal,
	imports map[string]*types.Type,
	scopeOverrides map[string]*types.Type,
) (*vm.CompiledFunc, error) {
	compiled := vm.NewCompiledFunc(fn, parentFunc, constants, file)

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
		resolvedTypeRegister, err := file.Types.Add(arg.Type)
		if err != nil {
			return nil, err
		}

		compiled.NewVariable(arg.Name, file.Types.Get(resolvedTypeRegister))

		compiled.Arguments = append(compiled.Arguments, arg.Name)
	}

	err := compileBlock(compiled, fn.Statements, nil, nil,
		file, scopeOverrides)
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
	// resolved the type of all variables) we can now compile all the deferred
	// function literals that might reference variables in this scope.
	instructions := len(compiled.Instructions.Instructions)
	for _, fn := range compiled.DeferredFuncsToCompile {
		cf, err := CompileFunc(fn.Func, file, compiled, constants, imports, scopeOverrides)
		if err != nil {
			return nil, err
		}

		file.AddSymbolFunc(cf)

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
	compiled.Instructions.Instructions = append(
		compiled.Instructions.Instructions[instructions:],
		compiled.Instructions.Instructions[:instructions]...)

	// Since we moved instructions to the top we also need to adjust the jump
	// positions down accordingly.
	if instructions > 0 {
		count := len(compiled.Instructions.Instructions) - instructions

		for index, ins := range compiled.Instructions.Instructions {
			switch actual := ins.(type) {
			case *vm.Jump:
				actual.To += count

			case *vm.JumpUnless:
				actual.To += count
			}

			compiled.Instructions.Instructions[index] = ins
		}
	}

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
