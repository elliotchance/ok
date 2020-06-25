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

func TestTest(t *testing.T) {
	for testName, test := range map[string]struct {
		fn       *ast.Test
		expected []vm.Instruction
		err      error
	}{
		"no-statements": {
			fn: &ast.Test{},
		},
		"one-statement": {
			fn: &ast.Test{
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
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileTest(test.fn, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
