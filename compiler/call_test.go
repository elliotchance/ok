package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/instruction"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCall(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []instruction.Instruction
		err      error
	}{
		"print-0": {
			nodes: []ast.Node{
				&ast.Call{
					FunctionName: "print",
				},
			},
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
				},
			},
		},
		"len-1": {
			nodes: []ast.Node{
				&ast.Call{
					FunctionName: "len",
					Arguments: []ast.Node{
						ast.NewArrayNumbers(nil),
					},
				},
			},
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.ArrayAllocNumber{
					Size:   "1",
					Result: "2",
				},
				&instruction.Len{
					Argument: "2",
					Result:   "3",
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...))
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
