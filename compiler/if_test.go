package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/lexer"
	"ok/vm"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIf(t *testing.T) {
	for testName, test := range map[string]struct {
		nodes    []ast.Node
		expected []vm.Instruction
		err      error
	}{
		"if-1": {
			nodes: []ast.Node{
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.If{
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenEqual,
						Right: ast.NewLiteralNumber("3"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&vm.JumpUnless{
					Condition: "3",
					To:        4,
				},
			},
		},
		"if-2": {
			nodes: []ast.Node{
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.If{
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenEqual,
						Right: ast.NewLiteralNumber("3"),
					},
					True: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "a"},
							Op:    lexer.TokenAssign,
							Right: ast.NewLiteralNumber("1"),
						},
						&ast.Unary{
							Op:   lexer.TokenIncrement,
							Expr: &ast.Identifier{Name: "a"},
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "1",
				},
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&vm.JumpUnless{
					Condition: "3",
					To:        8,
				},
				&vm.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("1"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "4",
				},
				&vm.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("1"),
				},
				&vm.Add{
					Left:   "a",
					Right:  "5",
					Result: "a",
				},
			},
		},
		"if-else-2": {
			nodes: []ast.Node{
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.If{
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenEqual,
						Right: ast.NewLiteralNumber("3"),
					},
					True: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "a"},
							Op:    lexer.TokenAssign,
							Right: ast.NewLiteralNumber("1"),
						},
					},
					False: []ast.Node{
						&ast.Unary{
							Op:   lexer.TokenIncrement,
							Expr: &ast.Identifier{Name: "a"},
						},
					},
				},
			},
			expected: []vm.Instruction{
				// a = 0
				&vm.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// if a == 3
				&vm.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&vm.JumpUnless{
					Condition: "3",
					To:        7,
				},

				// a = 1
				&vm.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("1"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "4",
				},
				&vm.Jump{
					To: 9,
				},

				// ++a
				&vm.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("1"),
				},
				&vm.Add{
					Left:   "a",
					Right:  "5",
					Result: "a",
				},
			},
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
