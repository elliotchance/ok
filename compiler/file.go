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
		Imports: imports,
		Types:   map[vm.TypeRegister]*types.Type{},
		Symbols: map[vm.SymbolRegister]*vm.Symbol{},
	}

	for _, fn := range funcs {
		file.AddSymbolFunc(&vm.CompiledFunc{
			Type:       fn.Type(),
			Name:       fn.Name,
			UniqueName: fn.UniqueName,
		})
	}

	for _, fn := range funcs {
		compiledFn, err := CompileFunc(fn, file, nil, constants)
		if err != nil {
			return nil, err
		}

		file.AddSymbolFunc(compiledFn)
	}

	for _, fn := range tests {
		compiledFn, err := CompileTest(fn, file, constants)
		if err != nil {
			return nil, err
		}

		file.Tests = append(file.Tests, compiledFn)
	}

	return file, nil
}
