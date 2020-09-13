package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/vm"
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
				&vm.Print{},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileTest(test.fn,
				&vm.File{})
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
