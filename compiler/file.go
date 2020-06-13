package compiler

import (
	"ok/parser"
)

// CompileFile translates a single file into a set of instructions. The number
// of instructions returned may be zero.
func CompileFile(f *parser.File) (*CompiledFunc, error) {
	if f.Root != nil {
		return CompileFunc(f.Root)
	}

	return nil, nil
}
