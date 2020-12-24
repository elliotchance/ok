package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/types"
	"github.com/elliotchance/ok/vm"
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
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.If{
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenEqual,
						Right: asttest.NewLiteralNumber("3"),
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.Assign{
					Result:   "a",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "1",
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
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.If{
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenEqual,
						Right: asttest.NewLiteralNumber("3"),
					},
					True: []ast.Node{
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "a"},
							},
							Rights: []ast.Node{
								asttest.NewLiteralNumber("1"),
							},
						},
						&ast.Unary{
							Op:   lexer.TokenIncrement,
							Expr: &ast.Identifier{Name: "a"},
						},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.Assign{
					Result:   "a",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "1",
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
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "2",
				},
				&vm.Assign{
					Result:   "a",
					Register: "4",
				},
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "3",
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
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.If{
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenEqual,
						Right: asttest.NewLiteralNumber("3"),
					},
					True: []ast.Node{
						&ast.Assign{
							Lefts: []ast.Node{
								&ast.Identifier{Name: "a"},
							},
							Rights: []ast.Node{
								asttest.NewLiteralNumber("1"),
							},
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
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.Assign{
					Result:   "a",
					Register: "1",
				},

				// if a == 3
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "1",
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
				&vm.AssignSymbol{
					Result: "4",
					Symbol: "2",
				},
				&vm.Assign{
					Result:   "a",
					Register: "4",
				},
				&vm.Jump{
					To: 9,
				},

				// ++a
				&vm.AssignSymbol{
					Result: "5",
					Symbol: "3",
				},
				&vm.Add{
					Left:   "a",
					Right:  "5",
					Result: "a",
				},
			},
		},
		"if-is-1": {
			nodes: []ast.Node{
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.If{
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenIs,
						Right: &ast.Identifier{Name: "number"},
					},
				},
			},
			expected: []vm.Instruction{
				&vm.AssignSymbol{
					Result: "1",
					Symbol: "0",
				},
				&vm.Assign{
					Result:   "a",
					Register: "1",
				},
				&vm.AssignSymbol{
					Result: "2",
					Symbol: "1",
				},
				&vm.Is{
					Value:  "a",
					Type:   "2",
					Result: "3",
				},
				&vm.JumpUnless{
					Condition: "3",
					To:        4,
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(newFunc(test.nodes...),
				&vm.File{
					Symbols: map[vm.SymbolRegister]*vm.Symbol{},
					Types:   types.Registry{},
				}, nil, nil, nil, nil)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions.Instructions)
			}
		})
	}
}
