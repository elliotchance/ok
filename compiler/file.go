package compiler

import (
	"ok/parser"
	"ok/vm"
)

// CompileFile translates a single file into a set of instructions. The number
// of instructions returned may be zero.
func CompileFile(f *parser.File) (map[string]*vm.CompiledFunc, error) {
	fns := map[string]*vm.CompiledFunc{}
	for name, fn := range f.Funcs {
		compiledFn, err := CompileFunc(fn, f.Funcs)
		if err != nil {
			return nil, err
		}

		fns[name] = compiledFn
	}

	return fns, nil
}
