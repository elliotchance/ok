package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

// CompileFunc translates a single function into a set of instructions. The
// number of instructions returned may be zero.
func CompileFunc(fn *ast.Func, file *vm.File) (*vm.CompiledFunc, error) {
	compiled := &vm.CompiledFunc{
		Variables:  map[string]*types.Type{},
		Type:       fn.Type(),
		Name:       fn.Name,
		UniqueName: fn.UniqueName,
	}

	// All variables in a function are stored internally as a map right now. So
	// the first thing we need to do is initialise the map. Whether we return it
	// at the end will depend on the return type.
	//
	// It is important that we always use a map (even if its not returns)
	// because we may need it for a subscope (closure).
	isObject := len(fn.Returns) == 1 && fn.Returns[0].Name == fn.Name

	// Load the arguments from the registers.
	for _, arg := range fn.Arguments {
		compiled.NextRegister()
		compiled.NewVariable(arg.Name, arg.Type)

		compiled.Arguments = append(compiled.Arguments, arg.Name)
	}

	err := compileBlock(compiled, fn.Statements, nil, nil, file)
	if err != nil {
		return nil, err
	}

	// If this is an object, always returns its state.
	if isObject {
		compiled.Append(&vm.Return{
			Results: []vm.Register{vm.StateRegister},
		})
	}

	return compiled, nil
}
