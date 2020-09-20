package compiler

import (
	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
)

func compile(
	funcs map[string]*ast.Func,
	tests []*ast.Test,
	constants map[string]*ast.Literal,
	imports map[string]map[string]*types.Type,
) (*vm.File, error) {
	file := &vm.File{
		Funcs:     map[string]*vm.CompiledFunc{},
		FuncDefs:  funcs,
		Constants: constants,
		Imports:   imports,
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
