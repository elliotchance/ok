package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/vm"
)

// CompiledFile contains all the compiled entities in a file.
type Compiled struct {
	Funcs      map[string]*vm.CompiledFunc
	FuncDefs   map[string]*ast.Func
	Tests      []*vm.CompiledTest
	Interfaces map[string]map[string]string
}

// CompileFile translates a single file into a set of instructions. The number
// of instructions returned may be zero.
func CompileFile(f *parser.File, interfaces map[string]map[string]string) (*Compiled, error) {
	return compile(f.Funcs, f.Tests, interfaces)
}

func compile(
	funcs map[string]*ast.Func,
	tests []*ast.Test,
	interfaces map[string]map[string]string,
) (*Compiled, error) {
	file := &Compiled{
		Funcs:      map[string]*vm.CompiledFunc{},
		FuncDefs:   funcs,
		Interfaces: interfaces,
	}

	for name, fn := range funcs {
		compiledFn, err := CompileFunc(fn, file)
		if err != nil {
			return nil, err
		}

		file.Funcs[name] = compiledFn
	}

	for _, fn := range tests {
		compiledFn, err := CompileTest(fn, file)
		if err != nil {
			return nil, err
		}

		file.Tests = append(file.Tests, compiledFn)
	}

	return file, nil
}
