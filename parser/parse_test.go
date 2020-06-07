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
		err      error
	}{
		"empty": {
			str: "",
		},
		"func-paren-close": {
			str: "func)",
			err: errors.New("expecting identifier after func, but found )"),
		},
		"func-curly-open": {
			str: "func {",
			err: errors.New("expecting identifier after func, but found {"),
		},
		"func-name-paren-close": {
			str: "func main)",
			err: errors.New(`expecting ( after identifier, but found )`),
		},
		"func-name-paren-open": {
			str: "func main (",
			err: errors.New("expecting ) after (, but found end of file"),
		},
		"func-name-paren-open-close": {
			str: "func main ()",
			err: errors.New("expecting { after ), but found end of file"),
		},
		"func-name-paren-open-close-open": {
			str: "func main () {",
			err: errors.New("expecting } after {, but found end of file"),
		},
		"func-empty": {
			str:      "func main() {\n}",
			expected: &ast.Func{Name: "main"},
		},
		"unterminated-string": {
			str: `func "`,
			err: errors.New("unterminated string found after 'func'"),
		},
		"unterminated-string-first-token": {
			str: `"`,
			err: errors.New("unterminated string found at the start"),
		},
		"hello-world": {
			str:      `func main() {print("hello world")}`,
			expected: newFunc(ast.NewLiteralString("hello world")),
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
			str: "func main() {\n} (",
			err: errors.New("found extra '(' at the end of the file"),
		},
		"only-comment": {
			str: "// nothing to see here",
			comments: []*ast.Comment{
				{Comment: " nothing to see here"},
			},
		},
		"comments-everywhere": {
			str:      "// foo\n //bar\nfunc main() {\n// baz\nprint(\"hello\") // qux\n// quux\n}//corge\n//grault",
			expected: newFunc(ast.NewLiteralString("hello")),
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
			expected: newFunc(&ast.Literal{
				Kind:  lexer.TokenBool,
				Value: "true",
			}),
		},
		"literal-false": {
			str: `func main() { print(false) }`,
			expected: newFunc(&ast.Literal{
				Kind:  lexer.TokenBool,
				Value: "false",
			}),
		},
		"literal-char": {
			str: `func main() { print('a') }`,
			expected: newFunc(&ast.Literal{
				Kind:  lexer.TokenCharacter,
				Value: "a",
			}),
		},
		"literal-zero-length-char": {
			str: `func main() { print('') }`,
			err: errors.New("character literal cannot be empty"),
		},
		"literal-number-zero": {
			str: `func main() { print(0) }`,
			expected: newFunc(&ast.Literal{
				Kind:  lexer.TokenNumber,
				Value: "0",
			}),
		},
		"call-identifier-close": {
			str: `func main() { print) }`,
			err: errors.New("expecting ( after identifier, but found )"),
		},
		"call-identifier-without-literal": {
			str: `func main() { print( }`,
			err: errors.New("expecting expression after (, but found }"),
		},
		"call-identifier-missing-close": {
			str: `func main() { print("hello" }`,
			err: errors.New("expecting ) after string, but found }"),
		},
		"literal-number-negative": {
			str: `func main() { print(-3.20) }`,
			expected: newFunc(&ast.Unary{
				Op:   lexer.TokenMinus,
				Expr: ast.NewLiteralNumber("3.20"),
			}),
		},
		"numbers-plus": {
			str: `func main() { print(3 + 2) }`,
			expected: newFunc(&ast.Binary{
				Left:  ast.NewLiteralNumber("3"),
				Op:    lexer.TokenPlus,
				Right: ast.NewLiteralNumber("2"),
			}),
		},
		"numbers-minus": {
			str: `func main() { print(3 - 2) }`,
			expected: newFunc(&ast.Binary{
				Left:  ast.NewLiteralNumber("3"),
				Op:    lexer.TokenMinus,
				Right: ast.NewLiteralNumber("2"),
			}),
		},
		"numbers-times": {
			str: `func main() { print(3.0*2.1) }`,
			expected: newFunc(&ast.Binary{
				Left:  ast.NewLiteralNumber("3.0"),
				Op:    lexer.TokenTimes,
				Right: ast.NewLiteralNumber("2.1"),
			}),
		},
		"numbers-divide": {
			str: `func main() { print(3/2.0) }`,
			expected: newFunc(&ast.Binary{
				Left:  ast.NewLiteralNumber("3"),
				Op:    lexer.TokenDivide,
				Right: ast.NewLiteralNumber("2.0"),
			}),
		},
		"numbers-remainder": {
			str: `func main() { print(3 % 2) }`,
			expected: newFunc(&ast.Binary{
				Left:  ast.NewLiteralNumber("3"),
				Op:    lexer.TokenRemainder,
				Right: ast.NewLiteralNumber("2"),
			}),
		},
		"expr-3-linear-order": {
			str: `func main() { print(1 + 2 - 3) }`,
			expected: newFunc(&ast.Binary{
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
			expected: newFunc(&ast.Binary{
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
			expected: newFunc(&ast.Binary{
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
			expected: newFunc(
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
			expected: newFunc(
				&ast.Group{
					Expr: ast.NewLiteralNumber("2"),
				},
			),
		},
		"bool-and-bool": {
			str: `func main() { print(true and false) }`,
			expected: newFunc(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenAnd,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"bool-or-bool": {
			str: `func main() { print(true or false) }`,
			expected: newFunc(
				&ast.Binary{
					Left:  ast.NewLiteralBool(true),
					Op:    lexer.TokenOr,
					Right: ast.NewLiteralBool(false),
				},
			),
		},
		"not-bool": {
			str: `func main() { print(not true) }`,
			expected: newFunc(
				&ast.Unary{
					Op:   lexer.TokenNot,
					Expr: ast.NewLiteralBool(true),
				},
			),
		},
		"not-not-bool": {
			str: `func main() { print(not not false) }`,
			expected: newFunc(
				&ast.Unary{
					Op: lexer.TokenNot,
					Expr: &ast.Unary{
						Op:   lexer.TokenNot,
						Expr: ast.NewLiteralBool(false),
					},
				},
			),
		},
	} {
		t.Run(testName, func(t *testing.T) {
			f, err := parser.ParseString(test.str)
			if test.err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.NoError(t, err)
			}

			if f != nil {
				assert.Equal(t, test.expected, f.Root)
				assert.Equal(t, test.comments, f.Comments)
			}
		})
	}
}

func newFunc(arg ast.Node) *ast.Func {
	return &ast.Func{
		Name: "main",
		Statements: []ast.Node{
			&ast.Call{
				FunctionName: "print",
				Arguments:    []ast.Node{arg},
			},
		},
	}
}
