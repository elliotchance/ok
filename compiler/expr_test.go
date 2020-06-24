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

func TestExpr(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []vm.Instruction
		err      error
	}{
		"group-1": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Group{
						Expr: ast.NewBinary(
							ast.NewLiteralNumber("3"),
							lexer.TokenPlus,
							ast.NewLiteralNumber("4"),
						),
					},
					lexer.TokenTimes,
					ast.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("4"),
				},
				&vm.Add{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("5"),
				},
				&vm.Multiply{
					Left:   "3",
					Right:  "4",
					Result: "5",
				},
			},
		},
		"group-2": {
			nodes: []ast.Node{
				ast.NewBinary(
					ast.NewLiteralNumber("5"),
					lexer.TokenMinus,
					&ast.Group{
						Expr: ast.NewBinary(
							ast.NewLiteralNumber("3"),
							lexer.TokenPlus,
							ast.NewLiteralNumber("4"),
						),
					},
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("5"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("4"),
				},
				&vm.Add{
					Left:   "2",
					Right:  "3",
					Result: "4",
				},
				&vm.Subtract{
					Left:   "1",
					Right:  "4",
					Result: "5",
				},
			},
		},
		"group-3": {
			nodes: []ast.Node{
				ast.NewBinary(
					&ast.Group{
						Expr: &ast.Group{
							Expr: ast.NewLiteralNumber("5"),
						},
					},
					lexer.TokenPlus,
					&ast.Group{
						Expr: ast.NewLiteralNumber("3"),
					},
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("5"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.Add{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
			},
		},
		"group-4": {
			nodes: []ast.Node{
				&ast.Group{
					Expr: &ast.Group{
						Expr: ast.NewLiteralNumber("5"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("5"),
				},
			},
		},
		"variable-undefined": {
			nodes: []ast.Node{
				&ast.Identifier{Name: "foo"},
			},
			err: errors.New("undefined variable: foo"),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...), nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}