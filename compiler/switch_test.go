package compiler_test

import (
	"testing"

	"github.com/elliotchance/ok/ast"
	"github.com/elliotchance/ok/ast/asttest"
	"github.com/elliotchance/ok/compiler"
	"github.com/elliotchance/ok/lexer"
	"github.com/elliotchance/ok/vm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSwitch(t *testing.T) {
	for testName, test := range map[string]struct {
		fn       *ast.Func
		expected []vm.Instruction
		err      error
	}{
		"switch-1": {
			fn: newFunc(
				&ast.Switch{},
			),
			expected: nil,
		},
		"switch-6": {
			fn: newFunc(
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("1"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										asttest.NewLiteralString("ONE"),
									},
								},
							},
						},
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("2"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										asttest.NewLiteralString("TWO"),
									},
								},
							},
						},
					},
				},
			),
			expected: []vm.Instruction{
				// a = 0
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// a == 1
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
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

				// print("ONE")
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralString("ONE"),
				},
				&vm.Print{
					Arguments: []vm.Register{"4"},
				},
				&vm.Jump{
					To: 13,
				},

				// a == 2
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "5",
					Result: "6",
				},
				&vm.JumpUnless{
					Condition: "6",
					To:        13,
				},

				// print("TWO")
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralString("TWO"),
				},
				&vm.Print{
					Arguments: []vm.Register{"7"},
				},
				&vm.Jump{
					To: 13,
				},
			},
		},
		"switch-else-1": {
			fn: newFunc(
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("1"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										asttest.NewLiteralString("ONE"),
									},
								},
							},
						},
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("2"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										asttest.NewLiteralString("TWO"),
									},
								},
							},
						},
					},
					Else: []ast.Node{
						&ast.Call{
							FunctionName: "print",
							Arguments: []ast.Node{
								asttest.NewLiteralString("NO MATCH"),
							},
						},
					},
				},
			),
			expected: []vm.Instruction{
				// a = 0
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// a == 1
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
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

				// print("ONE")
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralString("ONE"),
				},
				&vm.Print{
					Arguments: []vm.Register{"4"},
				},
				&vm.Jump{
					To: 15,
				},

				// a == 2
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "5",
					Result: "6",
				},
				&vm.JumpUnless{
					Condition: "6",
					To:        13,
				},

				// print("TWO")
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralString("TWO"),
				},
				&vm.Print{
					Arguments: []vm.Register{"7"},
				},
				&vm.Jump{
					To: 15,
				},

				// print("NO MATCH")
				&vm.Assign{
					VariableName: "8",
					Value:        asttest.NewLiteralString("NO MATCH"),
				},
				&vm.Print{
					Arguments: []vm.Register{"8"},
				},
			},
		},
		"switch-7": {
			fn: newFunc(
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("1"),
								),
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("2"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										asttest.NewLiteralString("ONE OR TWO"),
									},
								},
							},
						},
						{
							Conditions: []ast.Node{
								asttest.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									asttest.NewLiteralNumber("3"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										asttest.NewLiteralString("THREE"),
									},
								},
							},
						},
					},
				},
			),
			expected: []vm.Instruction{
				// a = 0
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// a == 1
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
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

				// print("ONE OR TWO")
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralString("ONE OR TWO"),
				},
				&vm.Print{
					Arguments: []vm.Register{"4"},
				},
				&vm.Jump{
					To: 19,
				},

				// a == 2
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "5",
					Result: "6",
				},
				&vm.JumpUnless{
					Condition: "6",
					To:        13,
				},

				// print("ONE OR TWO")
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralString("ONE OR TWO"),
				},
				&vm.Print{
					Arguments: []vm.Register{"7"},
				},
				&vm.Jump{
					To: 19,
				},

				// a == 3
				&vm.Assign{
					VariableName: "8",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "8",
					Result: "9",
				},
				&vm.JumpUnless{
					Condition: "9",
					To:        19,
				},

				// print("THREE")
				&vm.Assign{
					VariableName: "10",
					Value:        asttest.NewLiteralString("THREE"),
				},
				&vm.Print{
					Arguments: []vm.Register{"10"},
				},
				&vm.Jump{
					To: 19,
				},
			},
		},
		"switch-value-1": {
			fn: newFunc(
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Switch{
					Expr: &ast.Identifier{Name: "a"},
				},
			),
			expected: []vm.Instruction{
				// a = 0
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "1",
				},
			},
		},
		"switch-value-2": {
			fn: newFunc(
				&ast.Assign{
					Lefts: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
					Rights: []ast.Node{
						asttest.NewLiteralNumber("0"),
					},
				},
				&ast.Switch{
					Expr: &ast.Identifier{Name: "a"},
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								asttest.NewLiteralNumber("1"),
								asttest.NewLiteralNumber("2"),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										asttest.NewLiteralString("ONE OR TWO"),
									},
								},
							},
						},
						{
							Conditions: []ast.Node{
								asttest.NewLiteralNumber("3"),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										asttest.NewLiteralString("THREE"),
									},
								},
							},
						},
					},
				},
			),
			expected: []vm.Instruction{
				// a = 0
				&vm.Assign{
					VariableName: "1",
					Value:        asttest.NewLiteralNumber("0"),
				},
				&vm.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// case 1
				&vm.Assign{
					VariableName: "2",
					Value:        asttest.NewLiteralNumber("1"),
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

				// print("ONE OR TWO")
				&vm.Assign{
					VariableName: "4",
					Value:        asttest.NewLiteralString("ONE OR TWO"),
				},
				&vm.Print{
					Arguments: []vm.Register{"4"},
				},
				&vm.Jump{
					To: 19,
				},

				// case 2
				&vm.Assign{
					VariableName: "5",
					Value:        asttest.NewLiteralNumber("2"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "5",
					Result: "6",
				},
				&vm.JumpUnless{
					Condition: "6",
					To:        13,
				},

				// print("ONE OR TWO")
				&vm.Assign{
					VariableName: "7",
					Value:        asttest.NewLiteralString("ONE OR TWO"),
				},
				&vm.Print{
					Arguments: []vm.Register{"7"},
				},
				&vm.Jump{
					To: 19,
				},

				// case 3
				&vm.Assign{
					VariableName: "8",
					Value:        asttest.NewLiteralNumber("3"),
				},
				&vm.EqualNumber{
					Left:   "a",
					Right:  "8",
					Result: "9",
				},
				&vm.JumpUnless{
					Condition: "9",
					To:        19,
				},

				// print("THREE")
				&vm.Assign{
					VariableName: "10",
					Value:        asttest.NewLiteralString("THREE"),
				},
				&vm.Print{
					Arguments: []vm.Register{"10"},
				},
				&vm.Jump{
					To: 19,
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(test.fn,
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
