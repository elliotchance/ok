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
		"true-and-true": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenAnd,
				ast.NewLiteralBool(true),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"true-and-false": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenAnd,
				ast.NewLiteralBool(false),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"true-or-true": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenOr,
				ast.NewLiteralBool(true),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"true-or-false": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenOr,
				ast.NewLiteralBool(false),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"not-true": {
			fn: newFunc(&ast.Unary{
				Op:   lexer.TokenNot,
				Expr: ast.NewLiteralBool(true),
			}),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"not-false": {
			fn: newFunc(&ast.Unary{
				Op:   lexer.TokenNot,
				Expr: ast.NewLiteralBool(true),
			}),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"bool-greater-than-bool": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenGreaterThan,
				ast.NewLiteralBool(true),
			)),
			err: errors.New("cannot perform bool > bool"),
		},
		"bool-equals-bool": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenEqual,
				ast.NewLiteralBool(true),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"char-equals-char": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralChar('a'),
				lexer.TokenEqual,
				ast.NewLiteralChar('B'),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"data-equals-data": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralData([]byte("a")),
				lexer.TokenEqual,
				ast.NewLiteralData([]byte("B")),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"number-equals-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenEqual,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"string-equals-string": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralString("foo"),
				lexer.TokenEqual,
				ast.NewLiteralString("bar"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"bool-not-equals-bool": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenNotEqual,
				ast.NewLiteralBool(true),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"char-not-equals-char": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralChar('a'),
				lexer.TokenNotEqual,
				ast.NewLiteralChar('B'),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"data-not-equals-data": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralData([]byte("a")),
				lexer.TokenNotEqual,
				ast.NewLiteralData([]byte("B")),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"number-not-equals-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenNotEqual,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"string-not-equals-string": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralString("foo"),
				lexer.TokenNotEqual,
				ast.NewLiteralString("bar"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"number-greater-than-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenGreaterThan,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"number-greater-than-equal-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenGreaterThanEqual,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"number-less-than-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenLessThan,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"number-less-than-equal-number": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenLessThanEqual,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"string-greater-than-string": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralString("1.0"),
				lexer.TokenGreaterThan,
				ast.NewLiteralString("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"string-greater-than-equal-string": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralString("1.0"),
				lexer.TokenGreaterThanEqual,
				ast.NewLiteralString("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(true),
					},
				},
			},
		},
		"string-less-than-string": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralString("1.0"),
				lexer.TokenLessThan,
				ast.NewLiteralString("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"string-less-than-equal-string": {
			fn: newFunc(ast.NewBinary(
				ast.NewLiteralString("1.0"),
				lexer.TokenLessThanEqual,
				ast.NewLiteralString("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralBool(false),
					},
				},
			},
		},
		"print-2": {
			fn: newFunc(
				ast.NewLiteralString("total is"),
				ast.NewBinary(
					ast.NewLiteralNumber("1.5"),
					lexer.TokenPlus,
					ast.NewLiteralNumber("0.8"),
				),
			),
			expected: []instruction.Instruction{
				&instruction.Print{
					Stdout: os.Stdout,
					Values: []*ast.Literal{
						ast.NewLiteralString("total is"),
						ast.NewLiteralNumber("2.3"),
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

func newFunc(args ...ast.Node) *ast.Func {
	return &ast.Func{
		Statements: []ast.Node{
			&ast.Call{
				FunctionName: "print",
				Arguments:    args,
			},
		},
	}
}
