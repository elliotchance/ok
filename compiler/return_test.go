package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReturn(t *testing.T) {
	for testName, test := range map[string]struct {
		node     ast.Node
		expected []vm.Instruction
		err      error
	}{
		"return-zero": {
			node: &ast.Return{},
			expected: []vm.Instruction{
				&vm.Return{},
			},
		},
		"return-number": {
			node: &ast.Return{
				Exprs: []ast.Node{
					ast.NewLiteralNumber("123"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("123"),
				},
				&vm.Return{
					Results: []string{"1"},
				},
			},
		},
		"return-multiple-values": {
			node: &ast.Return{
				Exprs: []ast.Node{
					ast.NewLiteralNumber("123"),
					ast.NewLiteralString("foo"),
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("123"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("foo"),
				},
				&vm.Return{
					Results: []string{"1", "2"},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(&ast.Func{
				Statements: []ast.Node{
					test.node,
				},
			}, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}