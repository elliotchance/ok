package compiler

import (
	"ok/ast"
	"ok/parser"
	"ok/vm"
)

// CompiledFile contains all the compiled entities in a file.
type Compiled struct {
	Funcs map[string]*vm.CompiledFunc
	Tests []*vm.CompiledTest
}

// CompileFile translates a single file into a set of instructions. The number
// of instructions returned may be zero.
func CompileFile(f *parser.File) (*Compiled, error) {
	return compile(f.Funcs, f.Tests)
}

func compile(funcs map[string]*ast.Func, tests []*ast.Test) (*Compiled, error) {
	file := &Compiled{
		Funcs: map[string]*vm.CompiledFunc{},
	}

	for name, fn := range funcs {
		compiledFn, err := CompileFunc(fn, funcs)
		if err != nil {
			return nil, err
		}

		file.Funcs[name] = compiledFn
	}

	for _, fn := range tests {
		compiledFn, err := CompileTest(fn, funcs)
		if err != nil {
			return nil, err
		}

		file.Tests = append(file.Tests, compiledFn)
	}

	return file, nil
}
