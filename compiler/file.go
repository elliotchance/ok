package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/parser"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

// CompileFile translates a single file into a set of instructions. The number
// of instructions returned may be zero.
func CompileFile(
	f *parser.File,
	interfaces map[string]map[string]*types.Type,
	constants map[string]*ast.Literal,
	imports map[string]map[string]*types.Type,
) (*vm.File, error) {
	return compile(f.Funcs, f.Tests, interfaces, constants, imports)
}

func compile(
	funcs map[string]*ast.Func,
	tests []*ast.Test,
	interfaces map[string]map[string]*types.Type,
	constants map[string]*ast.Literal,
	imports map[string]map[string]*types.Type,
) (*vm.File, error) {
	file := &vm.File{
		Funcs:      map[string]*vm.CompiledFunc{},
		FuncDefs:   funcs,
		Interfaces: interfaces,
		Constants:  constants,
		Imports:    imports,
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
