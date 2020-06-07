package compiler_test

import (
	"errors"
	"ok/ast"
	"ok/compiler"
	"ok/instruction"
	"ok/lexer"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompileFunc(t *testing.T) {
	for testName, test := range map[string]struct {
		fn       *ast.Func
		expected []instruction.Instruction
		err      error
	}{
		"no-statements": {
			fn: &ast.Func{},
		},
		"one-statement-no-args": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
				},
			},
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
				},
			},
		},
		"two-statements-with-args": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
					},
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							&ast.Literal{
								Kind:  lexer.TokenString,
								Value: "hello",
							},
						},
					},
				},
			},
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
				},
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{{
						Kind:  lexer.TokenString,
						Value: "hello",
					}},
				},
			},
		},
		"bool-plus-bool": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							&ast.Binary{
								Left: &ast.Literal{
									Kind:  lexer.TokenBool,
									Value: "true",
								},
								Op: lexer.TokenPlus,
								Right: &ast.Literal{
									Kind:  lexer.TokenBool,
									Value: "false",
								},
							},
						},
					},
				},
			},
			err: errors.New("cannot perform bool + bool"),
		},
		"string-divide-number": {
			fn: &ast.Func{
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							&ast.Binary{
								Left: &ast.Literal{
									Kind:  lexer.TokenString,
									Value: "foo",
								},
								Op: lexer.TokenDivide,
								Right: &ast.Literal{
									Kind:  lexer.TokenNumber,
									Value: "123",
								},
							},
						},
					},
				},
			},
			err: errors.New("cannot perform string / number"),
		},
		"data-plus-data": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralData([]byte("foo")),
				lexer.TokenPlus,
				ast.NewLiteralData([]byte("bar")),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralData([]byte("foobar")),
					},
				},
			},
		},
		"number-plus-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.20"),
				lexer.TokenPlus,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("6.20"),
					},
				},
			},
		},
		"string-plus-string": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralString("foo"),
				lexer.TokenPlus,
				ast.NewLiteralString("bar"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralString("foobar"),
					},
				},
			},
		},
		"number-minus-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.20"),
				lexer.TokenMinus,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("-3.80"),
					},
				},
			},
		},
		"number-times-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.20"),
				lexer.TokenTimes,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("6.00"),
					},
				},
			},
		},
		"number-divide-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.20"),
				lexer.TokenDivide,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("0.24"),
					},
				},
			},
		},
		"number-remainder-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("5"),
				lexer.TokenRemainder,
				ast.NewLiteralNumber("1.20"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("0.20"),
					},
				},
			},
		},
		"left-binary-failed": {
			fn: newFunc(ast.NewBinary(
				ast.NewBinary(
					ast.NewLiteralNumber("3"),
					lexer.TokenDivide,
					ast.NewLiteralNumber("0"),
				),
				lexer.TokenPlus,
				ast.NewLiteralNumber("5"),
			)),
			err: errors.New("division by zero"),
		},
		"right-binary-failed": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("5"),
				lexer.TokenPlus,
				ast.NewBinary(
					ast.NewLiteralNumber("3"),
					lexer.TokenDivide,
					ast.NewLiteralNumber("0"),
				),
			)),
			err: errors.New("division by zero"),
		},
		"group-1": {
			fn: newFunc(ast.NewBinary(
				&ast.Group{
					Expr: ast.NewBinary(
						ast.NewLiteralNumber("3"),
						lexer.TokenPlus,
						ast.NewLiteralNumber("4"),
					),
				},
				lexer.TokenTimes,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("35"),
					},
				},
			},
		},
		"group-2": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("5"),
				lexer.TokenMinus,
				&ast.Group{
					Expr: ast.NewBinary(
						ast.NewLiteralNumber("3"),
						lexer.TokenPlus,
						ast.NewLiteralNumber("4"),
					),
				},
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("-2"),
					},
				},
			},
		},
		"group-3": {
			fn: newFunc(ast.NewBinary(
				&ast.Group{
					Expr: &ast.Group{
						Expr: ast.NewLiteralNumber("5"),
					},
				},
				lexer.TokenPlus,
				&ast.Group{
					Expr: ast.NewLiteralNumber("3"),
				},
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("8"),
					},
				},
			},
		},
		"group-4": {
			fn: newFunc(
				&ast.Group{
					Expr: &ast.Group{
						Expr: ast.NewLiteralNumber("5"),
					},
				},
			),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralNumber("5"),
					},
				},
			},
		},
	} {
		t.Run(testName, func(t *testing.T) {
			instructions, err := compiler.CompileFunc(test.fn)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expected, instructions)
		})
	}
}

func newFunc(arg ast.Node) *ast.Func {
	return &ast.Func{
		Statements: []ast.Node{
			&ast.Call{
				FunctionName: "print",
				Arguments:    []ast.Node{arg},
			},
		},
	}
}
