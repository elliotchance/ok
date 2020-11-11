package compiler_test

import (
	"errors"
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/vm"
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
				asttest.NewBinary(
					&ast.Group{
						Expr: asttest.NewBinary(
							asttest.NewLiteralNumber("3"),
							lexer.TokenPlus,
							asttest.NewLiteralNumber("4"),
						),
					},
					lexer.TokenTimes,
					asttest.NewLiteralNumber("5"),
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("4"),
				},
				&vm.Add{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralNumber("5"),
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
				asttest.NewBinary(
					asttest.NewLiteralNumber("5"),
					lexer.TokenMinus,
					&ast.Group{
						Expr: asttest.NewBinary(
							asttest.NewLiteralNumber("3"),
							lexer.TokenPlus,
							asttest.NewLiteralNumber("4"),
						),
					},
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("5"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.Assign{
					VariableName: "3",
					Value:        asttest.NewLiteralNumber("4"),
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
				asttest.NewBinary(
					&ast.Group{
						Expr: &ast.Group{
							Expr: asttest.NewLiteralNumber("5"),
						},
					},
					lexer.TokenPlus,
					&ast.Group{
						Expr: asttest.NewLiteralNumber("3"),
					},
				),
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("5"),
				},
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("3"),
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
						Expr: asttest.NewLiteralNumber("5"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("5"),
				},
			},
		},
		"variable-undefined": {
			nodes: []ast.Node{
				&ast.Identifier{Name: "foo"},
			},
			err: errors.New(" undefined variable: foo"),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...),
				&vm.File{}, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
