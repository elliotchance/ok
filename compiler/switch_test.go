package compiler_test

import (
	"ok/ast"
	"ok/compiler"
	"ok/instruction"
	"ok/lexer"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSwitch(t *testing.T) {
	for testName, test := range map[string]struct {
		fn       *ast.Func
		expected []instruction.Instruction
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
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("1"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										ast.NewLiteralString("ONE"),
									},
								},
							},
						},
						{
							Conditions: []ast.Node{
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("2"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										ast.NewLiteralString("TWO"),
									},
								},
							},
						},
					},
				},
			),
			expected: []instruction.Instruction{
				// a = 0
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// a == 1
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&instruction.JumpUnless{
					Condition: "3",
					To:        7,
				},

				// print("ONE")
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralString("ONE"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"4"},
				},
				&instruction.Jump{
					To: 13,
				},

				// a == 2
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "5",
					Result: "6",
				},
				&instruction.JumpUnless{
					Condition: "6",
					To:        13,
				},

				// print("TWO")
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralString("TWO"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"7"},
				},
				&instruction.Jump{
					To: 13,
				},
			},
		},
		"switch-else-1": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("1"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										ast.NewLiteralString("ONE"),
									},
								},
							},
						},
						{
							Conditions: []ast.Node{
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("2"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										ast.NewLiteralString("TWO"),
									},
								},
							},
						},
					},
					Else: []ast.Node{
						&ast.Call{
							FunctionName: "print",
							Arguments: []ast.Node{
								ast.NewLiteralString("NO MATCH"),
							},
						},
					},
				},
			),
			expected: []instruction.Instruction{
				// a = 0
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// a == 1
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&instruction.JumpUnless{
					Condition: "3",
					To:        7,
				},

				// print("ONE")
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralString("ONE"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"4"},
				},
				&instruction.Jump{
					To: 15,
				},

				// a == 2
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "5",
					Result: "6",
				},
				&instruction.JumpUnless{
					Condition: "6",
					To:        13,
				},

				// print("TWO")
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralString("TWO"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"7"},
				},
				&instruction.Jump{
					To: 15,
				},

				// print("NO MATCH")
				&instruction.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralString("NO MATCH"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"8"},
				},
			},
		},
		"switch-7": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Switch{
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("1"),
								),
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("2"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										ast.NewLiteralString("ONE OR TWO"),
									},
								},
							},
						},
						{
							Conditions: []ast.Node{
								ast.NewBinary(
									&ast.Identifier{Name: "a"},
									lexer.TokenEqual,
									ast.NewLiteralNumber("3"),
								),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										ast.NewLiteralString("THREE"),
									},
								},
							},
						},
					},
				},
			),
			expected: []instruction.Instruction{
				// a = 0
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// a == 1
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&instruction.JumpUnless{
					Condition: "3",
					To:        7,
				},

				// print("ONE OR TWO")
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralString("ONE OR TWO"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"4"},
				},
				&instruction.Jump{
					To: 19,
				},

				// a == 2
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "5",
					Result: "6",
				},
				&instruction.JumpUnless{
					Condition: "6",
					To:        13,
				},

				// print("ONE OR TWO")
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralString("ONE OR TWO"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"7"},
				},
				&instruction.Jump{
					To: 19,
				},

				// a == 3
				&instruction.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "8",
					Result: "9",
				},
				&instruction.JumpUnless{
					Condition: "9",
					To:        19,
				},

				// print("THREE")
				&instruction.Assign{
					VariableName: "10",
					Value:        ast.NewLiteralString("THREE"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"10"},
				},
				&instruction.Jump{
					To: 19,
				},
			},
		},
		"switch-value-1": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Switch{
					Expr: &ast.Identifier{Name: "a"},
				},
			),
			expected: []instruction.Instruction{
				// a = 0
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "a",
					Register:     "1",
				},
			},
		},
		"switch-value-2": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Switch{
					Expr: &ast.Identifier{Name: "a"},
					Cases: []*ast.Case{
						{
							Conditions: []ast.Node{
								ast.NewLiteralNumber("1"),
								ast.NewLiteralNumber("2"),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										ast.NewLiteralString("ONE OR TWO"),
									},
								},
							},
						},
						{
							Conditions: []ast.Node{
								ast.NewLiteralNumber("3"),
							},
							Statements: []ast.Node{
								&ast.Call{
									FunctionName: "print",
									Arguments: []ast.Node{
										ast.NewLiteralString("THREE"),
									},
								},
							},
						},
					},
				},
			),
			expected: []instruction.Instruction{
				// a = 0
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "a",
					Register:     "1",
				},

				// case 1
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "2",
					Result: "3",
				},
				&instruction.JumpUnless{
					Condition: "3",
					To:        7,
				},

				// print("ONE OR TWO")
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralString("ONE OR TWO"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"4"},
				},
				&instruction.Jump{
					To: 19,
				},

				// case 2
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "5",
					Result: "6",
				},
				&instruction.JumpUnless{
					Condition: "6",
					To:        13,
				},

				// print("ONE OR TWO")
				&instruction.Assign{
					VariableName: "7",
					Value:        ast.NewLiteralString("ONE OR TWO"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"7"},
				},
				&instruction.Jump{
					To: 19,
				},

				// case 3
				&instruction.Assign{
					VariableName: "8",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.EqualNumber{
					Left:   "a",
					Right:  "8",
					Result: "9",
				},
				&instruction.JumpUnless{
					Condition: "9",
					To:        19,
				},

				// print("THREE")
				&instruction.Assign{
					VariableName: "10",
					Value:        ast.NewLiteralString("THREE"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"10"},
				},
				&instruction.Jump{
					To: 19,
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			compiledFunc, err := compiler.CompileFunc(test.fn)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, compiledFunc.Instructions)
			}
		})
	}
}
