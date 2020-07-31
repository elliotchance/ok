package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/vm"
)

// CompileFunc translates a single function into a set of instructions. The
// number of instructions returned may be zero.
func CompileFunc(fn *ast.Func, fns map[string]*ast.Func) (*vm.CompiledFunc, error) {
	compiled := &vm.CompiledFunc{
		Variables: map[string]string{},
	}

	// Objects are stored internally as maps right now. So the first thing we
	// need to do is initialise the map.
	//
	// The parser has already corrected the missing return type.
	isObject := len(fn.Returns) == 1 && fn.Returns[0] == fn.Name

	// Load the arguments from the registers.
	for _, arg := range fn.Arguments {
		compiled.NextRegister()
		compiled.NewVariable(arg.Name, arg.Type)

		compiled.Arguments = append(compiled.Arguments, arg.Name)
	}

	err := compileBlock(compiled, fn.Statements, nil, nil, fns)
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
