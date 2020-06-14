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
	"github.com/stretchr/testify/require"
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
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("hello"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"1"},
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
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralData([]byte("foo")),
				lexer.TokenPlus,
				ast.NewLiteralData([]byte("bar")),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralData([]byte("foo")),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralData([]byte("bar")),
				},
				&instruction.Combine{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-plus-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.20"),
				lexer.TokenPlus,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Add{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"string-plus-string": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralString("foo"),
				lexer.TokenPlus,
				ast.NewLiteralString("bar"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("foo"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("bar"),
				},
				&instruction.Concat{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-minus-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.20"),
				lexer.TokenMinus,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Subtract{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-times-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.20"),
				lexer.TokenTimes,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Multiply{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-divide-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.20"),
				lexer.TokenDivide,
				ast.NewLiteralNumber("5"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Divide{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-remainder-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("5"),
				lexer.TokenRemainder,
				ast.NewLiteralNumber("1.20"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1.20"),
				},
				&instruction.Remainder{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"group-1": {
			fn: newFuncPrint(ast.NewBinary(
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
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("4"),
				},
				&instruction.Add{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Multiply{
					Left:   "3",
					Right:  "4",
					Result: "5",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"5"},
				},
			},
		},
		"group-2": {
			fn: newFuncPrint(ast.NewBinary(
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
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("4"),
				},
				&instruction.Add{
					Left:   "2",
					Right:  "3",
					Result: "4",
				},
				&instruction.Subtract{
					Left:   "1",
					Right:  "4",
					Result: "5",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"5"},
				},
			},
		},
		"group-3": {
			fn: newFuncPrint(ast.NewBinary(
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
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Add{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"group-4": {
			fn: newFuncPrint(
				&ast.Group{
					Expr: &ast.Group{
						Expr: ast.NewLiteralNumber("5"),
					},
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("5"),
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"1"},
				},
			},
		},
		"true-and-true": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenAnd,
				ast.NewLiteralBool(true),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.And{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"true-and-false": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenAnd,
				ast.NewLiteralBool(false),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(false),
				},
				&instruction.And{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"true-or-true": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenOr,
				ast.NewLiteralBool(true),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Or{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"true-or-false": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenOr,
				ast.NewLiteralBool(false),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(false),
				},
				&instruction.Or{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"not-true": {
			fn: newFuncPrint(&ast.Unary{
				Op:   lexer.TokenNot,
				Expr: ast.NewLiteralBool(true),
			}),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Not{
					Left:   "1",
					Result: "2",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"2"},
				},
			},
		},
		"not-false": {
			fn: newFuncPrint(&ast.Unary{
				Op:   lexer.TokenNot,
				Expr: ast.NewLiteralBool(false),
			}),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(false),
				},
				&instruction.Not{
					Left:   "1",
					Result: "2",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"2"},
				},
			},
		},
		"bool-greater-than-bool": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenGreaterThan,
				ast.NewLiteralBool(true),
			)),
			err: errors.New("cannot perform bool > bool"),
		},
		"bool-equals-bool": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenEqual,
				ast.NewLiteralBool(true),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Equal{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"char-equals-char": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralChar('a'),
				lexer.TokenEqual,
				ast.NewLiteralChar('B'),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralChar('a'),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralChar('B'),
				},
				&instruction.Equal{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"data-equals-data": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralData([]byte("a")),
				lexer.TokenEqual,
				ast.NewLiteralData([]byte("B")),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralData([]byte("a")),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralData([]byte("B")),
				},
				&instruction.Equal{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-equals-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenEqual,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.EqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"string-equals-string": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralString("foo"),
				lexer.TokenEqual,
				ast.NewLiteralString("bar"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("foo"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("bar"),
				},
				&instruction.Equal{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"bool-not-equals-bool": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralBool(true),
				lexer.TokenNotEqual,
				ast.NewLiteralBool(true),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.NotEqual{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"char-not-equals-char": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralChar('a'),
				lexer.TokenNotEqual,
				ast.NewLiteralChar('B'),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralChar('a'),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralChar('B'),
				},
				&instruction.NotEqual{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"data-not-equals-data": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralData([]byte("a")),
				lexer.TokenNotEqual,
				ast.NewLiteralData([]byte("B")),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralData([]byte("a")),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralData([]byte("B")),
				},
				&instruction.NotEqual{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-not-equals-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenNotEqual,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.NotEqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"string-not-equals-string": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralString("foo"),
				lexer.TokenNotEqual,
				ast.NewLiteralString("bar"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("foo"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("bar"),
				},
				&instruction.NotEqual{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-greater-than-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenGreaterThan,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.GreaterThanNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-greater-than-equal-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenGreaterThanEqual,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.GreaterThanEqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-less-than-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenLessThan,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.LessThanNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"number-less-than-equal-number": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralNumber("1.0"),
				lexer.TokenLessThanEqual,
				ast.NewLiteralNumber("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.LessThanEqualNumber{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"string-greater-than-string": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralString("1.0"),
				lexer.TokenGreaterThan,
				ast.NewLiteralString("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("1"),
				},
				&instruction.GreaterThanString{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"string-greater-than-equal-string": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralString("1.0"),
				lexer.TokenGreaterThanEqual,
				ast.NewLiteralString("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("1"),
				},
				&instruction.GreaterThanEqualString{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"string-less-than-string": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralString("1.0"),
				lexer.TokenLessThan,
				ast.NewLiteralString("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("1"),
				},
				&instruction.LessThanString{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"string-less-than-equal-string": {
			fn: newFuncPrint(ast.NewBinary(
				ast.NewLiteralString("1.0"),
				lexer.TokenLessThanEqual,
				ast.NewLiteralString("1"),
			)),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("1.0"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("1"),
				},
				&instruction.LessThanEqualString{
					Left:   "1",
					Right:  "2",
					Result: "3",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3"},
				},
			},
		},
		"print-2": {
			fn: newFuncPrint(
				ast.NewLiteralString("total is"),
				ast.NewBinary(
					ast.NewLiteralNumber("1.5"),
					lexer.TokenPlus,
					ast.NewLiteralNumber("0.8"),
				),
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("total is"),
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("0.8"),
				},
				&instruction.Add{
					Left:   "2",
					Right:  "3",
					Result: "4",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"1", "4"},
				},
			},
		},
		"assign-1": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("1.5"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&instruction.Assign{
					VariableName: "foo",
					Register:     "1",
				},
			},
		},
		"assign-print": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("1.5"),
				},
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Identifier{Name: "foo"},
					},
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&instruction.Assign{
					VariableName: "foo",
					Register:     "1",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"foo"},
				},
			},
		},
		"assign-print-2": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("1.5"),
				},
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "foo"},
							Op:    lexer.TokenPlus,
							Right: ast.NewLiteralNumber("2"),
						},
						&ast.Binary{
							Left:  ast.NewLiteralNumber("10"),
							Op:    lexer.TokenTimes,
							Right: &ast.Identifier{Name: "foo"},
						},
					},
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1.5"),
				},
				&instruction.Assign{
					VariableName: "foo",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.Add{
					Left:   "foo",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("10"),
				},
				&instruction.Multiply{
					Left:   "4",
					Right:  "foo",
					Result: "5",
				},
				&instruction.Print{
					Stdout:    os.Stdout,
					Arguments: []string{"3", "5"},
				},
			},
		},
		"variable-undefined": {
			fn:  newFuncPrint(&ast.Identifier{Name: "foo"}),
			err: errors.New("undefined variable: foo"),
		},
		"variable-bad-reassign": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("1.5"),
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralString("1.5"),
				},
			),
			err: errors.New("cannot assign string to variable foo (expecting number)"),
		},
		"empty-for-without-condition": {
			fn: newFunc(
				&ast.For{},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.JumpUnless{
					Condition: "1",
					To:        2,
				},
				&instruction.Jump{
					To: -1,
				},
			},
		},
		"for-true-with-statements": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("1"),
				},
				&ast.For{
					Statements: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "bar"},
							Op:    lexer.TokenAssign,
							Right: ast.NewLiteralNumber("2"),
						},
						&ast.Binary{
							Left:  &ast.Identifier{Name: "baz"},
							Op:    lexer.TokenAssign,
							Right: ast.NewLiteralNumber("3"),
						},
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "qux"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("4"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Assign{
					VariableName: "foo",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.JumpUnless{
					Condition: "2",
					To:        8,
				},
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.Assign{
					VariableName: "bar",
					Register:     "3",
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Assign{
					VariableName: "baz",
					Register:     "4",
				},
				&instruction.Jump{
					To: 1,
				},
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("4"),
				},
				&instruction.Assign{
					VariableName: "qux",
					Register:     "5",
				},
			},
		},
		"for-condition": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.For{
					Condition: ast.NewBinary(
						&ast.Identifier{Name: "i"},
						lexer.TokenLessThan,
						ast.NewLiteralNumber("10"),
					),
					Statements: []ast.Node{
						&ast.Binary{
							Left: &ast.Identifier{Name: "i"},
							Op:   lexer.TokenAssign,
							Right: ast.NewBinary(
								&ast.Identifier{Name: "i"},
								lexer.TokenPlus,
								ast.NewLiteralNumber("1"),
							),
						},
					},
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("10"),
				},
				&instruction.LessThanNumber{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&instruction.JumpUnless{
					Condition: "3",
					To:        8,
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Add{
					Left:   "i",
					Right:  "4",
					Result: "5",
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "5",
				},
				&instruction.Jump{
					To: 2,
				},
			},
		},
		"for-break": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.For{
					Statements: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "i"},
							Op:    lexer.TokenAssign,
							Right: ast.NewLiteralNumber("1"),
						},
						&ast.Break{},
						&ast.Binary{
							Left:  &ast.Identifier{Name: "i"},
							Op:    lexer.TokenAssign,
							Right: ast.NewLiteralNumber("2"),
						},
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.JumpUnless{
					Condition: "2",
					To:        9,
				},
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
				},
				&instruction.Jump{
					To: 9,
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "4",
				},
				&instruction.Jump{
					To: 1,
				},
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "5",
				},
			},
		},
		"for-continue": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.For{
					Statements: []ast.Node{
						&ast.Binary{
							Left:  &ast.Identifier{Name: "i"},
							Op:    lexer.TokenAssign,
							Right: ast.NewLiteralNumber("1"),
						},
						&ast.Continue{},
						&ast.Binary{
							Left:  &ast.Identifier{Name: "i"},
							Op:    lexer.TokenAssign,
							Right: ast.NewLiteralNumber("2"),
						},
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralBool(true),
				},
				&instruction.JumpUnless{
					Condition: "2",
					To:        9,
				},
				&instruction.Assign{
					VariableName: "3",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
				},
				&instruction.Jump{
					To: 1,
				},
				&instruction.Assign{
					VariableName: "4",
					Value:        ast.NewLiteralNumber("2"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "4",
				},
				&instruction.Jump{
					To: 1,
				},
				&instruction.Assign{
					VariableName: "5",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "5",
				},
			},
		},
		"increment-variable": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Unary{
					Op:   "++",
					Expr: &ast.Identifier{Name: "i"},
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Add{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"decrement-variable": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Unary{
					Op:   "--",
					Expr: &ast.Identifier{Name: "i"},
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("1"),
				},
				&instruction.Subtract{
					Left:   "i",
					Right:  "2",
					Result: "i",
				},
			},
		},
		"plus-assign-data": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralData([]byte("foo")),
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: ast.NewLiteralData([]byte("bar")),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralData([]byte("foo")),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralData([]byte("bar")),
				},
				&instruction.Combine{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
				},
			},
		},
		"plus-assign-string": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralString("foo"),
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: ast.NewLiteralString("bar"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralString("foo"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralString("bar"),
				},
				&instruction.Concat{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
				},
			},
		},
		"plus-assign-number": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenPlusAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Add{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
				},
			},
		},
		"minus-assign-number": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenMinusAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Subtract{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
				},
			},
		},
		"times-assign-number": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenTimesAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Multiply{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
				},
			},
		},
		"divide-assign-number": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenDivideAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Divide{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
				},
			},
		},
		"remainder-assign-number": {
			fn: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("0"),
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "i"},
					Op:    lexer.TokenRemainderAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
			expected: []instruction.Instruction{
				&instruction.Assign{
					VariableName: "1",
					Value:        ast.NewLiteralNumber("0"),
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "1",
				},
				&instruction.Assign{
					VariableName: "2",
					Value:        ast.NewLiteralNumber("3"),
				},
				&instruction.Remainder{
					Left:   "i",
					Right:  "2",
					Result: "3",
				},
				&instruction.Assign{
					VariableName: "i",
					Register:     "3",
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

func newFunc(statements ...ast.Node) *ast.Func {
	return &ast.Func{
		Statements: statements,
	}
}

func newFuncPrint(args ...ast.Node) *ast.Func {
	return &ast.Func{
		Statements: []ast.Node{
			&ast.Call{
				FunctionName: "print",
				Arguments:    args,
			},
		},
	}
}
