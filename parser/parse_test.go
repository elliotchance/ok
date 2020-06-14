package parser_test

import (
	"errors"
	"ok/ast"
	"ok/lexer"
	"ok/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	for testName, test := range map[string]struct {
		str      string
		expected *ast.Func
		comments []*ast.Comment
		errs     []error
	}{
		"empty": {
			str: "",
		},
		"func-paren-close": {
			str: "func)",
			errs: []error{
				errors.New("expecting identifier after func, but found )"),
			},
		},
		"func-curly-open": {
			str: "func {",
			errs: []error{
				errors.New("expecting identifier after func, but found {"),
			},
		},
		"func-name-paren-close": {
			str: "func main)",
			errs: []error{
				errors.New(`expecting ( after identifier, but found )`),
			},
		},
		"func-name-paren-open": {
			str: "func main (",
			errs: []error{
				errors.New("expecting ) after (, but found end of file"),
			},
		},
		"func-name-paren-open-close": {
			str: "func main ()",
			errs: []error{
				errors.New("expecting { after ), but found end of file"),
			},
		},
		"func-name-paren-open-close-open": {
			str: "func main () {",
			errs: []error{
				errors.New("expecting } after {, but found end of file"),
			},
		},
		"func-empty": {
			str:      "func main() {\n}",
			expected: &ast.Func{Name: "main"},
		},
		"func-empty-2": {
			str:      "func\nmain() {\n}",
			expected: &ast.Func{Name: "main"},
		},
		"unterminated-string": {
			str: `func "`,
			errs: []error{
				errors.New("unterminated string found after 'func'"),
			},
		},
		"unterminated-string-first-token": {
			str: `"`,
			errs: []error{
				errors.New("unterminated string found at the start"),
			},
		},
		"hello-world": {
			str:      `func main() {print("hello world")}`,
			expected: newFuncPrint(ast.NewLiteralString("hello world")),
		},
		"hello-world-2": {
			str: `func main() {print("hello") print("world")}`,
			expected: &ast.Func{
				Name: "main",
				Statements: []ast.Node{
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							ast.NewLiteralString("hello"),
						},
					},
					&ast.Call{
						FunctionName: "print",
						Arguments: []ast.Node{
							ast.NewLiteralString("world"),
						},
					},
				},
			},
		},
		"extra-token": {
			str:      "func main() {\n} (",
			expected: newFunc(),
			errs: []error{
				errors.New("found extra '(' at the end of the file"),
			},
		},
		"only-comment": {
			str: "// nothing to see here",
			comments: []*ast.Comment{
				{Comment: " nothing to see here"},
			},
		},
		"comments-everywhere": {
			str:      "// foo\n //bar\nfunc main() {\n// baz\nprint(\"hello\") // qux\n// quux\n}//corge\n//grault",
			expected: newFuncPrint(ast.NewLiteralString("hello")),
			comments: []*ast.Comment{
				{Comment: " foo"},
				{Comment: "bar"},
				{Comment: " baz"},
				{Comment: " qux"},
				{Comment: " quux"},
				{Comment: "corge"},
				{Comment: "grault"},
			},
		},
		"literal-true": {
			str: `func main() { print(true) }`,
			expected: newFuncPrint(&ast.Literal{
				Kind:  lexer.TokenBool,
				Value: "true",
			}),
		},
		"literal-false": {
			str: `func main() { print(false) }`,
			expected: newFuncPrint(&ast.Literal{
				Kind:  lexer.TokenBool,
				Value: "false",
			}),
		},
		"literal-char": {
			str: `func main() { print('a') }`,
			expected: newFuncPrint(&ast.Literal{
				Kind:  lexer.TokenCharacter,
				Value: "a",
			}),
		},
		"literal-zero-length-char": {
			str: `func main() { print('') }`,
			expected: newFuncPrint(&ast.Literal{
				Kind:  lexer.TokenCharacter,
				Value: "",
			}),
			errs: []error{
				errors.New("character literal cannot be empty"),
			},
		},
		"literal-number-zero": {
			str: `func main() { print(0) }`,
			expected: newFuncPrint(&ast.Literal{
				Kind:  lexer.TokenNumber,
				Value: "0",
			}),
		},
		"call-identifier-close": {
			str: `func main() { print) }`,
			errs: []error{
				errors.New("expecting } after identifier, but found )"),
			},
		},
		"call-identifier-without-literal": {
			str: `func main() { print( }`,
			errs: []error{
				errors.New("expecting } after identifier, but found ("),
			},
		},
		"call-identifier-missing-close": {
			str: `func main() { print("hello" }`,
			errs: []error{
				errors.New("expecting } after identifier, but found ("),
			},
		},
		"literal-number-negative": {
			str: `func main() { print(-3.20) }`,
			expected: newFuncPrint(&ast.Unary{
				Op:   lexer.TokenMinus,
				Expr: ast.NewLiteralNumber("3.20"),
			}),
		},
		"numbers-plus": {
			str: `func main() { print(3 + 2) }`,
			expected: newFuncPrint(&ast.Binary{
				Left:  ast.NewLiteralNumber("3"),
				Op:    lexer.TokenPlus,
				Right: ast.NewLiteralNumber("2"),
			}),
		},
		"numbers-minus": {
			str: `func main() { print(3 - 2) }`,
			expected: newFuncPrint(&ast.Binary{
				Left:  ast.NewLiteralNumber("3"),
				Op:    lexer.TokenMinus,
				Right: ast.NewLiteralNumber("2"),
			}),
		},
		"numbers-times": {
			str: `func main() { print(3.0*2.1) }`,
			expected: newFuncPrint(&ast.Binary{
				Left:  ast.NewLiteralNumber("3.0"),
				Op:    lexer.TokenTimes,
				Right: ast.NewLiteralNumber("2.1"),
			}),
		},
		"numbers-divide": {
			str: `func main() { print(3/2.0) }`,
			expected: newFuncPrint(&ast.Binary{
				Left:  ast.NewLiteralNumber("3"),
				Op:    lexer.TokenDivide,
				Right: ast.NewLiteralNumber("2.0"),
			}),
		},
		"numbers-remainder": {
			str: `func main() { print(3 % 2) }`,
			expected: newFuncPrint(&ast.Binary{
				Left:  ast.NewLiteralNumber("3"),
				Op:    lexer.TokenRemainder,
				Right: ast.NewLiteralNumber("2"),
			}),
		},
		"expr-3-linear-order": {
			str: `func main() { print(1 + 2 - 3) }`,
			expected: newFuncPrint(&ast.Binary{
				Left: ast.NewLiteralNumber("1"),
				Op:   lexer.TokenPlus,
				Right: &ast.Binary{
					Left:  ast.NewLiteralNumber("2"),
					Op:    lexer.TokenMinus,
					Right: ast.NewLiteralNumber("3"),
				},
			}),
		},
		"expr-3-non-linear-order": {
			str: `func main() { print(1 * 2 - 3) }`,
			expected: newFuncPrint(&ast.Binary{
				Left: &ast.Binary{
					Left:  ast.NewLiteralNumber("1"),
					Op:    lexer.TokenTimes,
					Right: ast.NewLiteralNumber("2"),
				},
				Op:    lexer.TokenMinus,
				Right: ast.NewLiteralNumber("3"),
			}),
		},
		"expr-3-grouping": {
			str: `func main() { print(1 * (2 - 3)) }`,
			expected: newFuncPrint(&ast.Binary{
				Left: ast.NewLiteralNumber("1"),
				Op:   lexer.TokenTimes,
				Right: &ast.Group{
					Expr: &ast.Binary{
						Left:  ast.NewLiteralNumber("2"),
						Op:    lexer.TokenMinus,
						Right: ast.NewLiteralNumber("3"),
					},
				},
			}),
		},
		"group-1": {
			str: `func main() { print((2 - 3)) }`,
			expected: newFuncPrint(
				&ast.Group{
					Expr: &ast.Binary{
						Left:  ast.NewLiteralNumber("2"),
						Op:    lexer.TokenMinus,
						Right: ast.NewLiteralNumber("3"),
					},
				},
			),
		},
		"group-2": {
			str: `func main() { print((2)) }`,
			expected: newFuncPrint(
				&ast.Group{
					Expr: ast.NewLiteralNumber("2"),
				},
			),
		},
		"bool-and-bool": {
			str: `func main() { print(true and false) }`,
			expected: newFuncPrint(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenAnd,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"bool-or-bool": {
			str: `func main() { print(true or false) }`,
			expected: newFuncPrint(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenOr,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"not-bool": {
			str: `func main() { print(not true) }`,
			expected: newFuncPrint(
				&ast.Unary{
					Op:   lexer.TokenNot,
					Expr: ast.NewLiteralBool(true),
				},
			),
		},
		"not-not-bool": {
			str: `func main() { print(not not false) }`,
			expected: newFuncPrint(
				&ast.Unary{
					Op: lexer.TokenNot,
					Expr: &ast.Unary{
						Op:   lexer.TokenNot,
						Expr: ast.NewLiteralBool(false),
					},
				},
			),
		},
		"bool-equal-bool": {
			str: `func main() { print(true==false) }`,
			expected: newFuncPrint(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenEqual,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"bool-not-equal-bool": {
			str: `func main() { print(true != false) }`,
			expected: newFuncPrint(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenNotEqual,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"bool-greater-than-bool": {
			str: `func main() { print(true>false) }`,
			expected: newFuncPrint(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenGreaterThan,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"bool-greater-than-equal-bool": {
			str: `func main() { print(true>=false) }`,
			expected: newFuncPrint(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenGreaterThanEqual,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"bool-less-than-bool": {
			str: `func main() { print(true < false) }`,
			expected: newFuncPrint(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenLessThan,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"bool-less-than-equal-bool": {
			str: `func main() { print(true <= false) }`,
			expected: newFuncPrint(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenLessThanEqual,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"print-2": {
			str: `func main() { print(true, false) }`,
			expected: newFuncPrint(
				ast.NewLiteralBool(true),
				ast.NewLiteralBool(false),
			),
		},
		"assign-string": {
			str: `func main() { foo = "bar" }`,
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralString("bar"),
				},
			),
		},
		"read-variable": {
			str: `func main() { foo = bar }`,
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "foo"},
					Op:    lexer.TokenAssign,
					Right: &ast.Identifier{Name: "bar"},
				},
			),
		},
		"end-of-line-1": {
			str: "func main() { a = 1\nprint(a) }",
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("1"),
				},
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Identifier{Name: "a"},
					},
				},
			),
		},
		"end-of-line-2": {
			str: "func main() { b = true\nprint(b)\nc = 1.23 }",
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "b"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralBool(true),
				},
				&ast.Call{
					FunctionName: "print",
					Arguments: []ast.Node{
						&ast.Identifier{Name: "b"},
					},
				},
				&ast.Binary{
					Left:  &ast.Identifier{Name: "c"},
					Op:    lexer.TokenAssign,
					Right: ast.NewLiteralNumber("1.23"),
				},
			),
		},
		"for-empty-body": {
			str: "func main() { for true {} }",
			expected: newFunc(
				&ast.For{
					Condition: ast.NewLiteralBool(true),
				},
			),
		},
		"for-1": {
			str: "func main() { for true { print(a)\nprint(b) } }",
			expected: newFunc(
				&ast.For{
					Condition: ast.NewLiteralBool(true),
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "print",
							Arguments: []ast.Node{
								&ast.Identifier{Name: "a"},
							},
						},
						&ast.Call{
							FunctionName: "print",
							Arguments: []ast.Node{
								&ast.Identifier{Name: "b"},
							},
						},
					},
				},
			),
		},
		"for-empty-without-condition": {
			str: "func main() { for {} }",
			expected: newFunc(
				&ast.For{},
			),
		},
		"for-2": {
			str: "func main() { for { print(a)\nprint(b) } }",
			expected: newFunc(
				&ast.For{
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "print",
							Arguments: []ast.Node{
								&ast.Identifier{Name: "a"},
							},
						},
						&ast.Call{
							FunctionName: "print",
							Arguments: []ast.Node{
								&ast.Identifier{Name: "b"},
							},
						},
					},
				},
			),
		},
		"break-1": {
			str: "func main() { for { break\nprint(a) } }",
			expected: newFunc(
				&ast.For{
					Statements: []ast.Node{
						&ast.Break{},
						&ast.Call{
							FunctionName: "print",
							Arguments: []ast.Node{
								&ast.Identifier{Name: "a"},
							},
						},
					},
				},
			),
		},
		"continue-1": {
			str: "func main() { for { print(a)\ncontinue } }",
			expected: newFunc(
				&ast.For{
					Statements: []ast.Node{
						&ast.Call{
							FunctionName: "print",
							Arguments: []ast.Node{
								&ast.Identifier{Name: "a"},
							},
						},
						&ast.Continue{},
					},
				},
			),
		},
		"increment": {
			str: "func main() { ++a }",
			expected: newFunc(
				&ast.Unary{
					Op:   "++",
					Expr: &ast.Identifier{Name: "a"},
				},
			),
		},
		"decrement": {
			str: "func main() { --a }",
			expected: newFunc(
				&ast.Unary{
					Op:   "--",
					Expr: &ast.Identifier{Name: "a"},
				},
			),
		},
		"add-assign": {
			str: "func main() { a += 3 }",
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenPlusAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
		},
		"minus-assign": {
			str: "func main() { a -= 3 }",
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenMinusAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
		},
		"times-assign": {
			str: "func main() { a *= 3 }",
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenTimesAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
		},
		"divide-assign": {
			str: "func main() { a /= 3 }",
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenDivideAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
		},
		"remainder-assign": {
			str: "func main() { a %= 3 }",
			expected: newFunc(
				&ast.Binary{
					Left:  &ast.Identifier{Name: "a"},
					Op:    lexer.TokenRemainderAssign,
					Right: ast.NewLiteralNumber("3"),
				},
			),
		},
		"if-1": {
			str: "func main() { if a == 3 { } }",
			expected: newFunc(
				&ast.If{
					Condition: &ast.Binary{
						Left:  &ast.Identifier{Name: "a"},
						Op:    lexer.TokenEqual,
						Right: ast.NewLiteralNumber("3"),
					},
				},
			),
		},
		"if-2": {
			str: "func main() { if a == 3 { a = 1\n++a } }",
			expected: newFunc(
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
			),
		},
		"if-else-1": {
			str: "func main() { if a == 3 { a = 1 } else { } }",
			expected: newFunc(
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
				},
			),
		},
		"if-else-2": {
			str: "func main() { if a == 3 { a = 1 } else { ++a } }",
			expected: newFunc(
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
			),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			p := parser.ParseString(test.str)

			assertEqualErrors(t, test.errs, p.Errors)
			assert.Equal(t, test.expected, p.File.Root)
			assert.Equal(t, test.comments, p.File.Comments)
		})
	}
}

func assertEqualErrors(t *testing.T, expected, actual []error) {
	var e []string
	for _, err := range expected {
		e = append(e, err.Error())
	}

	var a []string
	for _, err := range actual {
		a = append(a, err.Error())
	}

	assert.Equal(t, e, a)
}

func newFuncPrint(args ...ast.Node) *ast.Func {
	return &ast.Func{
		Name: "main",
		Statements: []ast.Node{
			&ast.Call{
				FunctionName: "print",
				Arguments:    args,
			},
		},
	}
}

func newFunc(statements ...ast.Node) *ast.Func {
	return &ast.Func{
		Name:       "main",
		Statements: statements,
	}
}
