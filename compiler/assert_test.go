package compiler_test

import (
	"errors"
	"ok/ast"
	"ok/compiler"
	"ok/lexer"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssert(t *testing.T) {
	for testName, test := range map[string]struct {
		node     *ast.Assert
		expected []vm.Instruction
		err      error
	}{
		"not-a-bool": {
			node: &ast.Assert{
				Expr: ast.NewBinary(
					ast.NewLiteralNumber("1"),
					lexer.TokenPlus,
					ast.NewLiteralNumber("2"),
				),
			},
			err: errors.New("assert condition must be a bool but is number"),
		},
		"success": {
			node: &ast.Assert{
				Expr: ast.NewBinary(
					ast.NewLiteralNumber("1"),
					lexer.TokenEqual,
					ast.NewLiteralNumber("2"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("2"),
				},
				&vm.EqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&vm.Assert{
					Left:  "1",
					Op:    "==",
					Right: "2",
					Final: "3",
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
