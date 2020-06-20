package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/vm"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFunc(t *testing.T) {
	for testName, test := range map[string]struct {
		fn       *ast.Func
		expected []vm.Instruction
		err      error
	}{
		"no-statements": {
			fn: &ast.Func{},
		},
		"one-statement-no-args": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{
					Stdout: os.Stdout,
				},
			},
		},
		"two-statements-with-args": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							ast.NewLiteralString("hello"),
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Print{
					Stdout: os.Stdout,
				},
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("hello"),
				},
				&vm.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"1"},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(test.fn, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}

func newFunc(statements ...ast.Node) *ast.Func {
	return &ast.Func{
		Statements: statements,
	}
}
