package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/instruction"
	"ok/parser"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompileFile(t *testing.T) {
	for testName, test := range map[string]struct {
		f        *parser.File
		expected []instruction.Instruction
	}{
		"no-functions": {
			&parser.File{},
			nil,
		},
		"one-function": {
			&parser.File{
				Root: &ast.Func{
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "print",
						},
					},
				},
			},
			[]instruction.Instruction{
				&instruction.Print{Stdout: os.Stdout},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			instructions, err := compiler.CompileFile(test.f)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, instructions)
		})
	}
}
