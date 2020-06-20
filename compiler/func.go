package compiler

import (
	"ok/ast"
	"ok/vm"
)

// CompileFunc translates a single function into a set of instructions. The
// number of instructions returned may be zero.
func CompileFunc(fn *ast.Func, fns map[string]*ast.Func) (*vm.CompiledFunc, error) {
	compiled := &vm.CompiledFunc{
		Variables: map[string]string{},
	}

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

	return compiled, nil
}
