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
	imports map[string]*types.Type,
) (*vm.File, error) {
	file := &vm.File{
		Funcs:     map[string]*vm.CompiledFunc{},
		Constants: constants,
		Imports:   imports,
		Types:     map[vm.TypeRegister]*types.Type{},
	}

	for _, fn := range funcs {
		file.Funcs[fn.UniqueName] = &vm.CompiledFunc{
			Type:       fn.Type(),
			Name:       fn.Name,
			UniqueName: fn.UniqueName,
		}
	}

	for _, fn := range funcs {
		compiledFn, err := CompileFunc(fn, file, nil)
		if err != nil {
			return nil, err
		}

		file.Funcs[fn.UniqueName] = compiledFn
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
