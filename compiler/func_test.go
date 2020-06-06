package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/instruction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompileFunc(t *testing.T) {
	for testName, test := range map[string]struct {
		fn       *ast.Func
		expected []instruction.Instruction
	}{
		"no-statements": {
			&ast.Func{},
			nil,
		},
		"one-statement-no-args": {
			&ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
				},
			},
			[]instruction.Instruction{
				&instruction.Print{},
			},
		},
		"two-statements-with-args": {
			&ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
					&ast.Call{
						FunctionName: "print",
						Arguments:    []string{"hello world"},
					},
				},
			},
			[]instruction.Instruction{
				&instruction.Print{},
				&instruction.Print{Value: "hello world"},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			instructions := compiler.CompileFunc(test.fn)
			assert.Equal(t, test.expected, instructions)
		})
	}
}
